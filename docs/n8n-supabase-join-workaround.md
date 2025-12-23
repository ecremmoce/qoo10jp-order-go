# N8N에서 Supabase JOIN 해결 방법

## 문제 상황

N8N의 기본 Supabase 노드는 **JOIN 연산을 지원하지 않습니다**.

```
❌ N8N Supabase 노드의 한계:
- SELECT 쿼리에 JOIN 불가
- 관계 테이블의 데이터를 함께 조회할 수 없음
- 복잡한 PostgREST 쿼리 문법 미지원
```

### 실제 사례

**요구사항**: `shopee_access_tokens` 테이블에서 활성 토큰을 조회하면서, 관련된 `sales_platform_accounts_shopee` 테이블의 `shop_name`도 함께 가져오기

```sql
-- 원하는 SQL 쿼리
SELECT 
  t.*,
  s.shop_name
FROM shopee_access_tokens t
INNER JOIN sales_platform_accounts_shopee s ON t.shop_id = s.shop_id
WHERE t.is_active = true;
```

**문제**: N8N Supabase 노드로는 이런 JOIN 쿼리를 실행할 수 없음 ❌

---

## ✅ 해결 방법: HTTP Request 노드 + PostgREST API

Supabase는 내부적으로 **PostgREST**를 사용하므로, HTTP Request 노드로 직접 REST API를 호출하면 JOIN이 가능합니다!

### 1단계: HTTP Request 노드 설정

#### 노드 타입
```
n8n-nodes-base.httpRequest (v4.2+)
```

#### 기본 설정
- **Method**: `GET`
- **URL**: `https://YOUR_PROJECT.supabase.co/rest/v1/TABLE_NAME`
- **Authentication**: `Predefined Credential Type`
- **Credential Type**: `Supabase API`

#### Query Parameters 추가

| Parameter Name | Value |
|---------------|-------|
| `is_active` | `eq.true` |
| `select` | `*,sales_platform_accounts_shopee!inner(shop_name)` |

### 2단계: PostgREST JOIN 문법

```
select=*,관계테이블명!inner(컬럼1,컬럼2,...)
```

#### 핵심 문법 요소

1. **`*`**: 메인 테이블의 모든 컬럼
2. **`,`**: AND 조건
3. **`관계테이블명!inner(...)`**: INNER JOIN
   - `!inner`: INNER JOIN (필수 관계)
   - `!left`: LEFT JOIN (선택 관계)
4. **`(컬럼1,컬럼2)`**: 관계 테이블에서 가져올 컬럼

#### 예시

```
# 단일 컬럼
select=*,sales_platform_accounts_shopee!inner(shop_name)

# 여러 컬럼
select=*,sales_platform_accounts_shopee!inner(shop_name,shop_region,seller_name)

# LEFT JOIN
select=*,sales_platform_accounts_shopee!left(shop_name)

# 여러 관계 테이블
select=*,table1!inner(col1),table2!left(col2,col3)
```

### 3단계: 반환 데이터 구조 이해 ⚠️

**중요**: PostgREST JOIN은 **객체**로 반환됩니다 (배열 아님!)

#### 반환 형식
```json
[
  {
    "id": "uuid-here",
    "shop_id": 1140237553,
    "access_token": "token...",
    "is_active": true,
    "sales_platform_accounts_shopee": {  // ← 객체!
      "shop_name": "shoelamode.mx"
    }
  }
]
```

#### N8N 표현식에서 접근

```javascript
// ✅ 올바른 방법 (객체 접근)
$json.sales_platform_accounts_shopee.shop_name

// ❌ 잘못된 방법 (배열 접근)
$json.sales_platform_accounts_shopee[0].shop_name
```

#### 안전한 접근 (Null 체크 포함)

```javascript
// 방법 1: AND 조건
$json.sales_platform_accounts_shopee && 
$json.sales_platform_accounts_shopee.shop_name 
  ? $json.sales_platform_accounts_shopee.shop_name 
  : 'Shop-' + $json.shop_id

// 방법 2: Optional chaining (N8N에서 지원 안 될 수 있음)
$json.sales_platform_accounts_shopee?.shop_name || 'Shop-' + $json.shop_id
```

---

## 완전한 예시: Redis에 메시지 푸시

### 시나리오
Shopee 주문 수집을 위해 활성 계정 정보를 Redis 큐에 푸시

### 워크플로우 구조

```
Cron Trigger 
  → HTTP Request (Supabase JOIN) 
  → IF (계정 존재 확인)
  → Split In Batches
  → Redis Enhanced (메시지 푸시)
```

### HTTP Request 노드 설정

```yaml
노드명: Get Active Accounts
타입: n8n-nodes-base.httpRequest
버전: 4.2

Parameters:
  method: GET
  url: https://cawyuwexdhlgoflckaxv.supabase.co/rest/v1/shopee_access_tokens
  authentication: predefinedCredentialType
  nodeCredentialType: supabaseApi
  sendQuery: true
  queryParameters:
    - name: is_active
      value: eq.true
    - name: select
      value: *,sales_platform_accounts_shopee!inner(shop_name)
```

### Redis Enhanced 노드 설정

```yaml
노드명: Redis Enhanced
타입: @fancyheat/n8n-nodes-redis-enhanced.redisEnhanced

Parameters:
  operation: push
  list: shopee_order_queue
  messageData: |
    ={{ 
      JSON.stringify({ 
        account_id: $json.id,
        account_name: (
          $json.sales_platform_accounts_shopee && 
          $json.sales_platform_accounts_shopee.shop_name 
            ? $json.sales_platform_accounts_shopee.shop_name 
            : 'Shop-' + $json.shop_id
        ),
        platform: 'shopee',
        shop_id: $json.shop_id,
        access_token: $json.access_token,
        timestamp: new Date().toISOString()
      })
    }}
```

### 결과 (Redis 메시지)

```json
{
    "account_id": "cbba41d5-673f-489a-9bd1-a7f2cb92df16",
    "account_name": "shoelamode.mx",
    "platform": "shopee",
    "shop_id": 1140237553,
    "access_token": "eyJhbGciOiJIUzI1NiJ9...",
    "timestamp": "2025-09-30T17:58:21.886Z"
}
```

---

## PostgREST 필터 치트시트

### 비교 연산자
```
eq.value        # =
neq.value       # !=
gt.value        # >
gte.value       # >=
lt.value        # <
lte.value       # <=
```

### 패턴 매칭
```
like.*pattern*      # LIKE %pattern%
ilike.*pattern*     # ILIKE %pattern% (대소문자 무시)
```

### NULL 체크
```
is.null
not.is.null
```

### IN 연산자
```
in.(value1,value2,value3)
```

### 조합 예시
```
# 활성 상태이면서 특정 region
is_active=eq.true&region=eq.SG

# shop_id가 특정 목록에 포함
shop_id=in.(123,456,789)

# 만료되지 않은 토큰
expires_at=gte.2025-09-30T00:00:00Z

# shop_name이 특정 패턴 포함
shop_name=ilike.*ecremmoce*
```

---

## 외래키 설정 (선택사항)

JOIN이 제대로 작동하려면 Supabase에서 **외래키(Foreign Key)** 관계를 설정해야 합니다.

### SQL로 외래키 추가

```sql
-- shopee_access_tokens → sales_platform_accounts_shopee
ALTER TABLE shopee_access_tokens
ADD CONSTRAINT fk_shop_id
FOREIGN KEY (shop_id) 
REFERENCES sales_platform_accounts_shopee(shop_id)
ON DELETE CASCADE;
```

### Supabase 대시보드에서 설정

1. **Table Editor** → 테이블 선택
2. 컬럼 옆 **Edit Column**
3. **Foreign Key Relations** 섹션
4. 관계 설정:
   - Source column: `shop_id`
   - Target table: `sales_platform_accounts_shopee`
   - Target column: `shop_id`
   - On Delete: `CASCADE` or `SET NULL`

---

## 트러블슈팅

### 1. "relation does not exist" 에러
```
✅ 해결: 테이블 이름과 컬럼 이름 확인
✅ 대소문자 구분 (snake_case 권장)
```

### 2. JOIN 결과가 비어있음
```
✅ !inner → !left 로 변경 시도
✅ 외래키 관계 확인
✅ 실제 데이터 존재 여부 확인
```

### 3. "null" 값 대신 빈 객체 반환
```javascript
// 안전한 접근
$json.table_name && $json.table_name.column
  ? $json.table_name.column
  : 'default_value'
```

### 4. N8N 표현식 에러
```
✅ Optional chaining (?.) 대신 명시적 조건문 사용
✅ 복잡한 표현식은 Code 노드로 대체
```

---

## 성능 최적화

### 1. 필요한 컬럼만 선택
```
# ❌ 모든 컬럼
select=*,table!inner(*)

# ✅ 필요한 컬럼만
select=id,shop_id,access_token,table!inner(shop_name)
```

### 2. 인덱스 활용
```sql
-- JOIN 컬럼에 인덱스 추가
CREATE INDEX idx_shop_id ON shopee_access_tokens(shop_id);
CREATE INDEX idx_shop_id_platform ON sales_platform_accounts_shopee(shop_id);
```

### 3. 페이지네이션
```
# Limit & Offset
limit=100&offset=0

# Range (더 효율적)
Range: 0-99  (HTTP Header)
```

---

## 참고 자료

- [PostgREST Documentation](https://postgrest.org/en/stable/)
- [Supabase PostgREST Guide](https://supabase.com/docs/guides/api/using-postgrest)
- [N8N HTTP Request Node](https://docs.n8n.io/integrations/builtin/core-nodes/n8n-nodes-base.httprequest/)

---

## 요약

| 항목 | N8N Supabase 노드 | HTTP Request + PostgREST |
|-----|------------------|------------------------|
| JOIN 지원 | ❌ 불가능 | ✅ 가능 |
| 복잡한 쿼리 | ❌ 제한적 | ✅ 자유로움 |
| 사용 난이도 | ⭐ 쉬움 | ⭐⭐ 보통 |
| 유연성 | 낮음 | 높음 |
| 권장 용도 | 단순 CRUD | JOIN, 복잡한 필터 |

**결론**: N8N Supabase 노드로 해결되지 않는 JOIN/복잡한 쿼리는 **HTTP Request 노드 + PostgREST API**를 사용하세요! 🚀

---

**작성일**: 2025-09-30  
**프로젝트**: shopee-order-go  
**작성자**: AI Assistant + User Collaboration

