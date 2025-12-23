# Acsell Fulfillment API 레퍼런스
## 사방넷 호환 RESTful API

**버전**: 1.0.0  
**Base URL**: `https://api.acsell.co.kr/v1`

---

## 📋 목차

1. [개요](#개요)
2. [인증 (Authentication)](#인증-authentication)
3. [사방넷과의 차이점](#사방넷과의-차이점)
4. [빠른 시작](#빠른-시작)
5. [API 엔드포인트](#api-엔드포인트)
   - [주문 수집](#post-orderscollect)
   - [송장 등록](#post-invoices)
   - [클레임 수집](#post-claimscollect)
6. [에러 처리](#에러-처리)
7. [필드 매핑 테이블](#필드-매핑-테이블)

---

## 개요

Acsell Fulfillment API는 Shopee 등 해외 쇼핑몰의 주문 정보를 Spigen WMS 풀필먼트 센터로 전송하기 위한 RESTful API입니다.

### 주요 특징

- ✅ **JSON 기반**: XML 대신 가볍고 파싱하기 쉬운 JSON 형식 사용
- ✅ **RESTful 표준**: 표준 HTTP Methods와 상태 코드 활용
- ✅ **타입 안정성**: 명확한 데이터 타입 정의 및 검증
- ✅ **ISO 8601 날짜**: 표준 날짜/시간 형식 지원
- ✅ **배열 지원**: 파이프 구분자 대신 네이티브 JSON 배열 사용
- ✅ **사방넷 호환**: 기존 사방넷 API와 동일한 기능 제공

### 지원 기능

| 기능 | 설명 | 사방넷 원본 |
|------|------|-------------|
| 주문 수집 | 쇼핑몰 주문 정보 조회 | xml_order_info.html |
| 송장 등록 | 운송장번호 및 택배사 등록 | xml_order_invoice.html |
| 클레임 수집 | 취소/반품/교환 정보 조회 | xml_clm_info.html |

---

## 인증 (Authentication)

API 요청 시 HTTP 헤더에 인증 정보를 포함해야 합니다.

### 필수 헤더

```http
Authorization: Bearer {API_KEY}
X-Company-ID: {COMPANY_ID}
Content-Type: application/json
```

### 인증 방식 비교

#### 사방넷 (XML)

```xml
<HEADER>
  <SEND_COMPAYNY_ID>company123</SEND_COMPAYNY_ID>
  <SEND_AUTH_KEY>auth_key_here</SEND_AUTH_KEY>
  <SEND_DATE>20251016</SEND_DATE>
</HEADER>
```

#### Acsell (RESTful)

```http
Authorization: Bearer auth_key_here
X-Company-ID: company123
Content-Type: application/json
```

### 인증 키 발급

1. Acsell 관리자 페이지 로그인
2. **설정 > API 관리** 메뉴 이동
3. **새 API 키 생성** 버튼 클릭
4. 생성된 API Key와 Company ID 저장

---

## 사방넷과의 차이점

### 1. 데이터 형식

| 구분 | 사방넷 | Acsell |
|------|--------|--------|
| 형식 | XML | JSON |
| 인코딩 | EUC-KR (UTF-8 옵션) | UTF-8 |
| CDATA | 사용 | 불필요 |

### 2. 요청 방식

| 구분 | 사방넷 | Acsell |
|------|--------|--------|
| Method | GET (XML URL 파라미터) | POST/PUT/DELETE |
| 인증 | XML 헤더 노드 | HTTP Authorization Header |
| 파라미터 | XML 노드 | JSON 객체 |

### 3. 날짜 형식

| 구분 | 사방넷 | Acsell |
|------|--------|--------|
| 날짜 | YYYYMMDD | YYYY-MM-DD (ISO 8601) |
| 시간 | YYYYMMDDHHMMSS | YYYY-MM-DDTHH:MM:SSZ |
| 예시 | 20251016 | 2025-10-16 |

### 4. 배열 처리

| 구분 | 사방넷 | Acsell |
|------|--------|--------|
| 구분자 | 파이프(\|) | JSON Array |
| 예시 | `field1\|field2\|field3` | `["field1", "field2", "field3"]` |

### 5. 에러 처리

| 구분 | 사방넷 | Acsell |
|------|--------|--------|
| 형식 | HTML/XML 메시지 | JSON 구조화된 에러 |
| HTTP 코드 | 항상 200 | 표준 HTTP 상태 코드 |

---

## 빠른 시작

### 주문 수집 예제

```bash
curl -X POST https://api.acsell.co.kr/v1/orders/collect \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "X-Company-ID: YOUR_COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{
    "startDate": "2025-10-01",
    "endDate": "2025-10-16",
    "outputFields": ["orderId", "orderDate", "productName", "quantity", "price"]
  }'
```

### 송장 등록 예제

```bash
curl -X POST https://api.acsell.co.kr/v1/invoices \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "X-Company-ID: YOUR_COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d '{
    "orderId": "SB2025001",
    "courierCode": "01",
    "trackingNumber": "123456789012"
  }'
```

### Python 예제

```python
import requests

API_BASE_URL = "https://api.acsell.co.kr/v1"
API_KEY = "YOUR_API_KEY"
COMPANY_ID = "YOUR_COMPANY_ID"

headers = {
    "Authorization": f"Bearer {API_KEY}",
    "X-Company-ID": COMPANY_ID,
    "Content-Type": "application/json"
}

# 주문 수집
response = requests.post(
    f"{API_BASE_URL}/orders/collect",
    headers=headers,
    json={
        "startDate": "2025-10-01",
        "endDate": "2025-10-16",
        "outputFields": ["orderId", "orderDate", "productName"]
    }
)

if response.status_code == 200:
    data = response.json()
    print(f"총 {data['data']['totalCount']}건의 주문 수집")
    for order in data['data']['orders']:
        print(f"- {order['orderId']}: {order['productName']}")
else:
    print(f"오류: {response.json()['error']['message']}")
```

### JavaScript (Node.js) 예제

```javascript
const axios = require('axios');

const API_BASE_URL = 'https://api.acsell.co.kr/v1';
const API_KEY = 'YOUR_API_KEY';
const COMPANY_ID = 'YOUR_COMPANY_ID';

const headers = {
  'Authorization': `Bearer ${API_KEY}`,
  'X-Company-ID': COMPANY_ID,
  'Content-Type': 'application/json'
};

// 주문 수집
async function collectOrders() {
  try {
    const response = await axios.post(
      `${API_BASE_URL}/orders/collect`,
      {
        startDate: '2025-10-01',
        endDate: '2025-10-16',
        outputFields: ['orderId', 'orderDate', 'productName']
      },
      { headers }
    );
    
    const { totalCount, orders } = response.data.data;
    console.log(`총 ${totalCount}건의 주문 수집`);
    orders.forEach(order => {
      console.log(`- ${order.orderId}: ${order.productName}`);
    });
  } catch (error) {
    console.error('오류:', error.response.data.error.message);
  }
}

collectOrders();
```

---

## API 엔드포인트

### POST /orders/collect

주문 정보를 수집합니다.

#### 요청

**Headers**

```http
Authorization: Bearer {API_KEY}
X-Company-ID: {COMPANY_ID}
Content-Type: application/json
```

**Body**

```json
{
  "startDate": "2025-10-01",
  "endDate": "2025-10-16",
  "outputFields": ["orderId", "orderDate", "productName", "quantity", "price"],
  "mallId": "shop001",
  "orderStatus": "01",
  "settlementConfirmed": "ALL"
}
```

**Parameters**

| 필드 | 타입 | 필수 | 설명 | 사방넷 원본 |
|------|------|------|------|-------------|
| startDate | string | Y | 검색 시작일 (YYYY-MM-DD) | ORD_ST_DATE |
| endDate | string | Y | 검색 종료일 (YYYY-MM-DD) | ORD_ED_DATE |
| outputFields | array | Y | 출력할 필드 목록 | ORD_FIELD |
| settlementConfirmed | string | N | 정산대조확인여부 (Y/N/ALL) | JUNG_CHK_YN2 |
| orderId | string | N | 특정 주문번호 조회 | ORDER_ID |
| mallId | string | N | 쇼핑몰 코드 | MALL_ID |
| orderStatus | string | N | 주문 상태 코드 | ORDER_STATUS |
| partnerId | string | N | 매입처 ID | PARTNER_ID |
| mallUserId | string | N | 쇼핑몰 로그인 ID | MALL_USER_ID |
| logisticsId | string | N | 물류처 ID | DPARTNER_ID |
| accountSerial | string | N | 계정등록순번 | ACNT_REGS_SRNO |

#### 응답

**Success (200)**

```json
{
  "success": true,
  "message": "주문 수집이 완료되었습니다",
  "data": {
    "totalCount": 150,
    "orders": [
      {
        "orderId": "SB2025001",
        "shopOrderId": "SHOP-2025-001",
        "mallName": "Shopee Singapore",
        "orderDate": "2025-10-15T10:30:00Z",
        "orderStatus": "주문확인",
        "productName": "스마트폰 케이스",
        "quantity": 2,
        "unitPrice": 15000,
        "totalPrice": 30000,
        "recipientName": "홍길동",
        "recipientPhone": "010-1234-5678",
        "address": "서울시 강남구 테헤란로 123",
        "zipCode": "06234",
        "deliveryMessage": "문 앞에 놓아주세요"
      }
    ]
  }
}
```

#### 주의사항

- API 요청 시 **신규주문이 주문확인 상태로 자동 변경**됩니다
- 동일 조건으로 재요청 시 **중복 데이터**가 반환될 수 있으므로 중복 체크 필요
- 정렬: 수취인명 → 우편번호 → 주소 → 주문번호 오름차순

---

### POST /invoices

송장 정보를 등록합니다.

#### 요청

**Headers**

```http
Authorization: Bearer {API_KEY}
X-Company-ID: {COMPANY_ID}
Content-Type: application/json
```

**Body**

```json
{
  "orderId": "SB2025001",
  "courierCode": "01",
  "trackingNumber": "123456789012",
  "deliveryHopeDate": "2025-10-20",
  "forceUpdate": false
}
```

**Parameters**

| 필드 | 타입 | 필수 | 설명 | 사방넷 원본 |
|------|------|------|------|-------------|
| orderId | string | Y | 주문번호 (사방넷) | SABANGNET_IDX |
| courierCode | string | Y | 택배사 코드 | TAK_CODE |
| trackingNumber | string | Y | 송장번호 | TAK_INVOICE |
| deliveryHopeDate | string | N | 배송희망일 (YYYY-MM-DD) | DELV_HOPE_DATE |
| forceUpdate | boolean | N | 강제 수정 여부 | SEND_INV_EDIT_YN |

**택배사 코드**

| 코드 | 택배사명 |
|------|----------|
| 01 | CJ대한통운 |
| 02 | 우체국택배 |
| 03 | 한진택배 |
| 04 | 롯데택배 |
| 05 | 로젠택배 |
| 06 | 대신택배 |
| 07 | 경동택배 |
| 08 | KGB택배 |
| 09 | CVSnet 편의점택배 |
| 10 | 합동택배 |

#### 응답

**Success (200)**

```json
{
  "success": true,
  "message": "송장이 등록되었습니다",
  "data": {
    "orderId": "SB2025001",
    "courierCode": "01",
    "trackingNumber": "123456789012",
    "registeredAt": "2025-10-16T15:30:00Z"
  }
}
```

#### 주의사항

- **주문확인** 상태에서 송장 입력 시 → **출고대기** 상태로 변경
- **출고대기** 상태에서 쇼핑몰 송장 전송 완료 시 → 수정 불가
- **강제완료** 상태 → 수정 불가
- `forceUpdate: true` 설정 시 출고대기 상태에서도 수정 가능 (단, 송장 전송 대기/실패 건만)

---

### POST /claims/collect

클레임 정보를 수집합니다.

#### 요청

**Headers**

```http
Authorization: Bearer {API_KEY}
X-Company-ID: {COMPANY_ID}
Content-Type: application/json
```

**Body**

```json
{
  "startDate": "2025-10-01",
  "endDate": "2025-10-16",
  "outputFields": ["claimId", "claimType", "claimDate", "orderId", "productName"]
}
```

**Parameters**

| 필드 | 타입 | 필수 | 설명 | 사방넷 원본 |
|------|------|------|------|-------------|
| startDate | string | Y | 검색 시작일 (YYYY-MM-DD) | CLM_ST_DATE |
| endDate | string | Y | 검색 종료일 (YYYY-MM-DD) | CLM_ED_DATE |
| outputFields | array | Y | 출력할 필드 목록 | CLM_FIELD |

#### 응답

**Success (200)**

```json
{
  "success": true,
  "message": "클레임 수집이 완료되었습니다",
  "data": {
    "totalCount": 25,
    "claims": [
      {
        "claimId": "CLM2025001",
        "orderId": "SB2025001",
        "shopOrderId": "SHOP-2025-001",
        "claimType": "RETURN",
        "claimReason": "단순변심",
        "claimDate": "2025-10-15T14:20:00Z",
        "collectedDate": "2025-10-15T14:25:00Z",
        "productName": "스마트폰 케이스",
        "quantity": 1,
        "claimAmount": 15000
      }
    ]
  }
}
```

**claimType 값**

| 값 | 설명 |
|----|------|
| CANCEL | 취소 |
| RETURN | 반품 |
| EXCHANGE | 교환 |

#### 주의사항

- **자동 수집된 클레임만** 조회 가능
- 사방넷 UI에서 수동 입력한 클레임은 조회되지 않음
- 동일 조건 재요청 시 중복 데이터 반환 가능 → 중복 체크 필요

---

## 에러 처리

### 에러 응답 구조

```json
{
  "success": false,
  "error": {
    "code": "ERROR_CODE",
    "message": "에러 메시지",
    "details": {
      "field": "fieldName",
      "reason": "상세 사유"
    }
  }
}
```

### HTTP 상태 코드

| 코드 | 설명 | 예시 |
|------|------|------|
| 200 | 성공 | 정상 처리 |
| 400 | 잘못된 요청 | 필수 파라미터 누락, 잘못된 형식 |
| 401 | 인증 실패 | API 키 누락 또는 유효하지 않음 |
| 403 | 권한 없음 | 접근 권한 없는 리소스 요청 |
| 404 | 리소스 없음 | 존재하지 않는 엔드포인트 |
| 429 | 요청 횟수 초과 | Rate Limit 초과 |
| 500 | 서버 오류 | 내부 서버 오류 |
| 503 | 서비스 이용 불가 | 점검 중 또는 서버 과부하 |

### 에러 코드

| 코드 | 설명 |
|------|------|
| UNAUTHORIZED | 인증 실패 |
| INVALID_REQUEST | 잘못된 요청 형식 |
| MISSING_PARAMETER | 필수 파라미터 누락 |
| INVALID_DATE_FORMAT | 잘못된 날짜 형식 |
| DATE_RANGE_TOO_LARGE | 조회 기간 초과 (최대 90일) |
| ORDER_NOT_FOUND | 주문을 찾을 수 없음 |
| INVOICE_ALREADY_SENT | 이미 송장이 전송됨 (수정 불가) |
| INVALID_ORDER_STATUS | 송장 등록 불가 상태 |
| RATE_LIMIT_EXCEEDED | API 요청 횟수 초과 |
| INTERNAL_ERROR | 서버 내부 오류 |

### 에러 처리 예제

```python
import requests

try:
    response = requests.post(
        "https://api.acsell.co.kr/v1/orders/collect",
        headers=headers,
        json=request_data
    )
    response.raise_for_status()
    data = response.json()
    
    if data['success']:
        print("성공:", data['message'])
    else:
        print("실패:", data['error']['message'])
        
except requests.exceptions.HTTPError as e:
    if e.response.status_code == 401:
        print("인증 오류: API 키를 확인하세요")
    elif e.response.status_code == 400:
        error_data = e.response.json()
        print(f"요청 오류: {error_data['error']['message']}")
    else:
        print(f"HTTP 오류: {e.response.status_code}")
        
except requests.exceptions.RequestException as e:
    print(f"네트워크 오류: {str(e)}")
```

---

## 필드 매핑 테이블

### 주문 수집 (Orders)

| 사방넷 XML | Acsell JSON | 타입 | 필수 | 설명 |
|-----------|-------------|------|------|------|
| ORD_ST_DATE | startDate | string | Y | 검색 시작일 |
| ORD_ED_DATE | endDate | string | Y | 검색 종료일 |
| ORD_FIELD | outputFields | array | Y | 출력 필드 목록 |
| JUNG_CHK_YN2 | settlementConfirmed | string | N | 정산대조확인여부 |
| ORDER_ID | orderId | string | N | 주문번호(쇼핑몰) |
| MALL_ID | mallId | string | N | 쇼핑몰 코드 |
| ORDER_STATUS | orderStatus | string | N | 주문 상태 |
| LANG | - | - | N | 제거됨 (항상 UTF-8) |
| PARTNER_ID | partnerId | string | N | 매입처 ID |
| MALL_USER_ID | mallUserId | string | N | 쇼핑몰 ID |
| DPARTNER_ID | logisticsId | string | N | 물류처 ID |
| ACNT_REGS_SRNO | accountSerial | string | N | 계정등록순번 |

### 송장 등록 (Invoices)

| 사방넷 XML | Acsell JSON | 타입 | 필수 | 설명 |
|-----------|-------------|------|------|------|
| SABANGNET_IDX | orderId | string | Y | 주문번호(사방넷) |
| TAK_CODE | courierCode | string | Y | 택배사 코드 |
| TAK_INVOICE | trackingNumber | string | Y | 송장번호 |
| DELV_HOPE_DATE | deliveryHopeDate | string | N | 배송희망일 |
| SEND_INV_EDIT_YN | forceUpdate | boolean | N | 강제 수정 여부 |

### 클레임 수집 (Claims)

| 사방넷 XML | Acsell JSON | 타입 | 필수 | 설명 |
|-----------|-------------|------|------|------|
| CLM_ST_DATE | startDate | string | Y | 검색 시작일 |
| CLM_ED_DATE | endDate | string | Y | 검색 종료일 |
| CLM_FIELD | outputFields | array | Y | 출력 필드 목록 |
| LANG | - | - | N | 제거됨 (항상 UTF-8) |

---

## 부록

### Rate Limiting

API 요청 횟수 제한:
- **일반**: 분당 60회, 시간당 1,000회
- **대량 처리**: 분당 10회, 시간당 100회 (주문/클레임 수집)

Rate Limit 헤더:
```http
X-RateLimit-Limit: 60
X-RateLimit-Remaining: 45
X-RateLimit-Reset: 1697456789
```

### 페이지네이션

대량 데이터 조회 시 페이지네이션 지원:

```json
{
  "page": 1,
  "pageSize": 100,
  "totalCount": 1500,
  "totalPages": 15
}
```

### Webhook (향후 지원 예정)

주문/클레임 발생 시 실시간 알림을 위한 Webhook 지원 계획 중

---

## 지원

- **기술 문의**: api@acsell.co.kr
- **문서 업데이트**: [GitHub Repository](https://github.com/acsell/api-docs)
- **OpenAPI 스펙**: [openapi.yaml](./openapi.yaml), [openapi.json](./openapi.json)
- **Swagger UI**: https://api.acsell.co.kr/docs

---

**마지막 업데이트**: 2025-10-16  
**API 버전**: 1.0.0

