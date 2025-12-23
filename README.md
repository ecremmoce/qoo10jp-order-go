# Shopee Order Collector

Shopee 주문 데이터를 수집하고 관리하는 Go 애플리케이션입니다.

## 기술 스택

- **Go 1.21+** - 백엔드 언어
- **Gin** - HTTP 웹 프레임워크
- **Supabase** - 데이터베이스 및 백엔드 서비스
- **Redis** - 캐싱 및 세션 관리
- **N8N** - 워크플로우 자동화
- **Shopee API** - 주문 데이터 소스

## 프로젝트 구조

```
shopee_order_go/
├── cmd/                    # 애플리케이션 진입점
│   └── main.go
├── internal/               # 내부 패키지
│   ├── api/               # API 라우터 및 핸들러
│   ├── config/            # 설정 관리
│   ├── models/            # 데이터 모델
│   └── services/          # 비즈니스 로직
├── pkg/                   # 외부 패키지
│   ├── shopee/           # Shopee API 클라이언트 ✅
│   ├── supabase/         # Supabase 클라이언트
│   ├── redis/            # Redis 클라이언트
│   └── webhook/          # 웹훅 클라이언트
├── scripts/              # 유틸리티 스크립트
├── workflow/             # N8N 워크플로우
├── web/                  # 관리자 페이지
├── go.mod
├── go.sum
└── README.md
```

## 설치 및 실행

### 1. 환경 변수 설정

`.env` 파일을 생성하고 다음 내용을 설정하세요:

```bash
# Supabase Configuration
SUPABASE_URL=your_supabase_url
SUPABASE_ANON_KEY=your_supabase_anon_key
SUPABASE_SERVICE_KEY=your_supabase_service_key

# Redis Configuration
REDIS_URL=redis://localhost:6379
REDIS_PASSWORD=
REDIS_DB=0

# N8N Configuration
N8N_BASE_URL=https://ec.ddns.net
N8N_API_KEY=your_n8n_api_key
N8N_WEBHOOK_URL=your_n8n_webhook_url
N8N_AUTH_ENDPOINT=your_n8n_auth_endpoint

# Shopee API Configuration
SHOPEE_PARTNER_ID=your_shopee_partner_id
SHOPEE_PARTNER_KEY=your_shopee_partner_key
SHOPEE_BASE_URL=https://partner.shopeemobile.com

# Server Configuration
PORT=8080
GIN_MODE=debug

# Worker Configuration
WORKER_COUNT=3

# Webhook Configuration
ORDER_COLLECTION_WEBHOOK_URL=https://n01.acsell.ai/webhook-test/shopee-order-collect-message
```

### 2. 의존성 설치

```bash
go mod tidy
```

### 3. 애플리케이션 실행

```bash
go run cmd/main.go
```

## API 엔드포인트

### Health Check
```
GET /api/v1/health
```

### 주문 수집
```
POST /api/v1/orders/collect
Content-Type: application/json

{
  "start_date": "2024-01-01",
  "end_date": "2024-01-31"
}
```

### 주문 조회
```
GET /api/v1/orders?start_date=2024-01-01&end_date=2024-01-31&status=completed&limit=100&offset=0
```

### 스케줄러 API

#### 작업 생성
```
POST /api/v1/scheduler/job
Content-Type: application/json

{
  "start_date": "2024-01-01",
  "end_date": "2024-01-31"
}
```

#### 스케줄러 상태 조회
```
GET /api/v1/scheduler/status
```

#### 워커 시작
```
POST /api/v1/scheduler/worker/start
```

#### 워커 중지
```
POST /api/v1/scheduler/worker/stop
```

#### 다음 작업 스케줄링
```
POST /api/v1/scheduler/schedule-next
```

## 데이터베이스 스키마

### orders 테이블
```sql
CREATE TABLE orders (
    id VARCHAR PRIMARY KEY,
    order_no VARCHAR UNIQUE NOT NULL,
    order_date TIMESTAMP NOT NULL,
    customer_id VARCHAR,
    customer_name VARCHAR,
    customer_email VARCHAR,
    customer_phone VARCHAR,
    total_amount DECIMAL(10,2),
    payment_status VARCHAR,
    order_status VARCHAR,
    shipping_address TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
```

### order_items 테이블
```sql
CREATE TABLE order_items (
    id VARCHAR PRIMARY KEY,
    order_id VARCHAR REFERENCES orders(id),
    product_id VARCHAR,
    product_name VARCHAR,
    quantity INTEGER,
    price DECIMAL(10,2),
    total_price DECIMAL(10,2)
);
```

### job_results 테이블
```sql
CREATE TABLE job_results (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    job_id VARCHAR NOT NULL,
    start_time TIMESTAMP WITH TIME ZONE NOT NULL,
    end_time TIMESTAMP WITH TIME ZONE NOT NULL,
    duration_ms BIGINT NOT NULL,
    orders_count INTEGER DEFAULT 0,
    success BOOLEAN NOT NULL DEFAULT false,
    error_msg TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
```

## 스케줄러 시스템

### 아키텍처
1. **N8N 스케줄러**: 5분마다 트리거 발생
2. **Redis 큐**: 작업 메시지 저장
3. **Go 워커**: Redis에서 작업을 가져와 처리
4. **Supabase**: 주문 데이터 및 작업 결과 저장
5. **순환 구조**: 마지막 실행 시간을 기반으로 다음 작업 스케줄링

### 설정 방법

#### 1. Supabase 테이블 생성
```bash
# Supabase에서 SQL 실행
psql -h your-supabase-host -U postgres -d postgres -f scripts/create-supabase-tables.sql
```

#### 2. N8N 워크플로우 배포

**기본 HTTP 방식:**
```powershell
.\scripts\deploy-n8n-workflow.ps1
```

**향상된 Redis 통합 방식 (권장):**
```powershell
.\scripts\deploy-n8n-workflow-redis.ps1 -RedisHost "localhost" -RedisPort "6379"
```

**기존 Redis 연결 사용 (추천):**
```powershell
.\scripts\deploy-existing-redis.ps1
```

**최고급 Enhanced Redis 방식:**
```powershell
.\scripts\deploy-n8n-enhanced.ps1 -RedisHost "localhost" -RedisPort "6379"
```

Enhanced Redis 방식의 추가 장점:
- 분산 스케줄러 잠금 (동시 실행 방지)
- TTL 기반 자동 잠금 해제
- 원자적 카운터 및 메트릭스
- 재시도 로직 및 Dead Letter Queue
- 실시간 워커 상태 추적
- 대량 작업 지원 (MGET, LPUSH 등)

#### 3. 워커 서비스 관리

**워커 시작:**
```powershell
# PowerShell
Invoke-WebRequest -Uri "http://localhost:8080/api/v1/scheduler/worker/start" -Method POST

# 또는 curl (Git Bash/WSL)
curl -X POST http://localhost:8080/api/v1/scheduler/worker/start
```

**워커 중지:**
```powershell
# PowerShell
Invoke-WebRequest -Uri "http://localhost:8080/api/v1/scheduler/worker/stop" -Method POST

# 또는 curl (Git Bash/WSL)
curl -X POST http://localhost:8080/api/v1/scheduler/worker/stop
```

**워커 상태 확인:**
```powershell
# PowerShell
Invoke-WebRequest -Uri "http://localhost:8080/api/v1/scheduler/status" -Method GET

# 또는 curl (Git Bash/WSL)
curl http://localhost:8080/api/v1/scheduler/status
```

**응답 예시:**
```json
{
  "workers_count": 3,
  "running": true,
  "queue_length": 0,
  "last_execution": "2025-09-26T09:12:09.500Z"
}
```

## Worker 관리

### Worker 시스템 개요

이 애플리케이션은 **Redis Queue 기반 Worker 시스템**을 사용합니다:

1. **N8N 워크플로우**가 5분마다 Supabase에서 활성 계정을 조회
2. 각 계정 정보를 **Redis Queue** (`shopee_order_queue`)에 추가
3. **Go Worker들**이 Queue에서 작업을 가져와 주문 수집 실행
4. 여러 Worker가 동시 실행되어도 **중복 처리 방지** 보장

### Worker 명령어

#### 시작/중지 명령어

**방법 1: PowerShell 터미널에서 실행**
```powershell
# 워커 시작
Invoke-WebRequest -Uri "http://localhost:8080/api/v1/scheduler/worker/start" -Method POST

# 워커 중지  
Invoke-WebRequest -Uri "http://localhost:8080/api/v1/scheduler/worker/stop" -Method POST

# 상태 확인
Invoke-WebRequest -Uri "http://localhost:8080/api/v1/scheduler/status" -Method GET
```

**방법 2: 웹 브라우저에서 실행**
- **Postman**, **Insomnia**, 또는 **Thunder Client** 같은 API 테스트 도구 사용
- 또는 브라우저 개발자 도구의 Console에서:
```javascript
// 워커 시작
fetch('http://localhost:8080/api/v1/scheduler/worker/start', {method: 'POST'})
  .then(r => r.json()).then(console.log);

// 워커 중지
fetch('http://localhost:8080/api/v1/scheduler/worker/stop', {method: 'POST'})
  .then(r => r.json()).then(console.log);

// 상태 확인 (GET 요청은 브라우저 주소창에서도 가능)
// http://localhost:8080/api/v1/scheduler/status
```

**방법 3: Git Bash/WSL에서 curl 사용**
```bash
# 워커 시작
curl -X POST http://localhost:8080/api/v1/scheduler/worker/start

# 워커 중지
curl -X POST http://localhost:8080/api/v1/scheduler/worker/stop

# 상태 확인
curl http://localhost:8080/api/v1/scheduler/status
```

#### 응답 메시지
- **시작 성공**: `{"message":"Worker service started"}`
- **이미 실행중**: `{"error":"Worker service is already running"}`
- **중지 성공**: `{"message":"Worker service stopped"}`
- **이미 중지됨**: `{"error":"Worker service is not running"}`

### Worker 설정

- **Worker 개수**: 3개 (기본값)
- **Queue 이름**: `shopee_order_queue`
- **처리 방식**: FIFO (First In, First Out)
- **재시도**: 최대 3회
- **타임아웃**: 30초

### 모니터링

#### Queue 상태 확인
```bash
# Redis CLI로 직접 확인
redis-cli LLEN shopee_order_queue

# API로 확인 (상태 응답에 포함)
curl http://localhost:8080/api/v1/scheduler/status
```

#### 로그 확인
- 콘솔 창에서 실시간 로그 확인
- Worker 시작/중지 메시지
- 작업 처리 진행 상황
- 오류 및 재시도 로그

### 문제 해결

#### Worker가 작업을 처리하지 않는 경우
1. **Worker 상태 확인**: `GET /api/v1/scheduler/status`
2. **Worker 재시작**: `POST /api/v1/scheduler/worker/stop` → `POST /api/v1/scheduler/worker/start`
3. **Queue 확인**: Redis에서 `shopee_order_queue` 길이 확인
4. **로그 확인**: 콘솔 창에서 오류 메시지 확인

#### 일반적인 오류
- **Redis 연결 실패**: `.env` 파일의 Redis 설정 확인
- **Supabase 연결 실패**: API 키 및 URL 확인
- **Shopee API 오류**: 파트너 ID/KEY 확인

## 개발

### 코드 스타일
- Go 표준 포맷팅 사용 (`go fmt`)
- Linting: `golangci-lint`

### 테스트 실행
```bash
go test ./...
```

### 빌드
```bash
go build -o shopee-order-go.exe cmd/main.go
```

## 배포

Docker를 사용한 배포:

```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o main cmd/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
CMD ["./main"]
```

## 웹훅 기능

### 개요
워커가 주문 수집을 시작하고 완료할 때 자동으로 웹훅을 호출하여 외부 시스템에 알림을 보냅니다.

### 설정
환경변수 `ORDER_COLLECTION_WEBHOOK_URL`에 웹훅 URL을 설정하세요:

```bash
ORDER_COLLECTION_WEBHOOK_URL=https://n01.acsell.ai/webhook-test/shopee-order-collect-message
```

### 웹훅 메시지 형식

#### 1. 주문 수집 시작
```
GET https://your-webhook-url?message=Shopee주문수집시작%20(계정:%20계정명)
```

#### 2. 주문 수집 완료
```
GET https://your-webhook-url?message=Shopee주문수집완료%20(계정:%20계정명,%205/10개%20저장)
```

#### 3. 주문 수집 실패
```
GET https://your-webhook-url?message=Shopee주문수집실패%20(계정:%20계정명,%20오류:%20API%20연결%20실패)
```

### 웹훅 테스트

웹훅 기능을 테스트하려면:

```bash
# Windows
test-webhook.bat

# 또는 직접 실행
go run scripts/test-webhook.go
```

### 웹훅 기능 특징

- **비동기 호출**: 워커 처리 속도에 영향을 주지 않음
- **재시도 로직**: 최대 2회 재시도
- **오류 처리**: 웹훅 실패 시 로그 기록
- **URL 인코딩**: 한글 메시지 자동 인코딩
- **타임아웃**: 10초 타임아웃 설정
- **Rate Limiting**: 동일 메시지 타입당 최소 5초 간격

## 📚 Shopee API 구현 현황

### ✅ 완료된 작업

1. **Shopee API 클라이언트 구현** (`pkg/shopee/client.go`)
   - ✅ HMAC-SHA256 인증 구현
   - ✅ `v2.order.get_order_list` - 주문 목록 조회
   - ✅ `v2.order.get_order_detail` - 주문 상세 조회
   - ✅ 페이지네이션 지원 (cursor 방식)

2. **데이터 모델 정의** (`internal/models/shopee_order.go`)
   - ✅ `ShopeeOrder` - 주문 정보
   - ✅ `ShopeeOrderItem` - 주문 상품
   - ✅ `ShopeeOrderFilter` - 조회 필터

### 🚧 진행 예정

1. **주문 수집 서비스** (`internal/services/shopee_order_service.go`)
   - ⏳ Shopee API를 통한 주문 수집
   - ⏳ Supabase 저장 로직
   - ⏳ 배치 처리 및 중복 체크

2. **데이터베이스 스키마**
   - ⏳ `shopee_orders` 테이블 생성
   - ⏳ 인덱스 설정 (order_sn, platform_account_id)

3. **API 라우트 추가**
   - ⏳ `POST /api/v1/shopee/orders/collect` - 주문 수집
   - ⏳ `GET /api/v1/shopee/orders` - 주문 조회
   - ⏳ 필터링 및 페이지네이션

4. **워커 통합**
   - ⏳ Redis 메시지에서 shop_id, access_token 파싱
   - ⏳ Shopee API 호출 및 주문 저장
   - ⏳ 웹훅 알림 (수집 시작/완료/실패)

### 📋 Shopee API 주요 스펙

#### 인증 방식
```
GET https://partner.shopeemobile.com/api/v2/order/get_order_list
  ?partner_id={partner_id}
  &timestamp={timestamp}
  &sign={hmac_sha256_signature}
  &shop_id={shop_id}
  &access_token={access_token}
  &time_range_field=create_time
  &time_from={unix_timestamp}
  &time_to={unix_timestamp}
  &page_size=100
```

#### 서명 생성
```
base_string = {partner_id}{api_path}{timestamp}
sign = HMAC-SHA256(base_string, partner_key)
```

#### 페이지네이션
- 첫 요청: `cursor` 없이 호출
- 다음 페이지: 응답의 `next_cursor` 사용
- `more=false`일 때까지 반복

## 라이센스

MIT License





