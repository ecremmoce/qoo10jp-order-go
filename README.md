# Qoo10JP Order Collector

Qoo10JP/Shopee 주문 데이터를 수집하고 관리하는 Go 백엔드 서비스입니다.

## 기술 스택

- **Go 1.21+** - 백엔드 언어
- **Gin** - HTTP 웹 프레임워크
- **Supabase** - 데이터베이스 (PostgreSQL)
- **Redis** - 큐 및 캐싱
- **N8N** - 워크플로우 자동화

## 프로젝트 구조

```
qoo10jp-order-go/
├── cmd/                    # 애플리케이션 진입점
│   └── main.go
├── internal/               # 내부 패키지
│   ├── api/               # API 라우터 및 핸들러
│   ├── config/            # 설정 관리
│   ├── models/            # 데이터 모델
│   ├── repository/        # 데이터 접근 계층 (예정)
│   └── services/          # 비즈니스 로직
├── pkg/                   # 외부 패키지
│   ├── qoo10jp/           # Qoo10JP API 클라이언트
│   ├── shopee/            # Shopee API 클라이언트
│   ├── supabase/          # Supabase 클라이언트
│   ├── redis/             # Redis 클라이언트
│   └── webhook/           # 웹훅 클라이언트
├── scripts/               # 유틸리티 스크립트
├── workflow/              # N8N 워크플로우
├── web/                   # 관리자 페이지
├── docs/                  # 문서
├── go.mod
├── go.sum
└── README.md
```

## 빠른 시작

### 1. 환경 변수 설정

`.env` 파일을 생성하세요 (`.env.example` 참조):

```bash
cp .env.example .env
# .env 파일 편집
```

### 2. 의존성 설치

```bash
go mod tidy
```

### 3. 애플리케이션 실행

```bash
# 개발 모드
go run cmd/main.go

# 또는 빌드 후 실행
go build -o qoo10jp-order-go.exe cmd/main.go
./qoo10jp-order-go.exe
```

### 4. 배치 스크립트 (Windows)

```powershell
# 전체 실행
.\run-all.bat

# 워커 시작/중지
.\start-worker.bat
.\stop-worker.bat

# 상태 확인
.\check-status.bat
.\worker-status.bat
```

## API 엔드포인트

### Health Check

```
GET /api/v1/health
GET /api/v1/qoo10jp/health
```

### Qoo10JP 주문 (V1 - Legacy)

```
POST /api/v1/qoo10jp/orders/collect    # 주문 수집
GET  /api/v1/qoo10jp/orders            # 주문 조회
GET  /api/v1/qoo10jp/orders/stats      # 주문 통계
```

### Qoo10JP 주문 (V2 - 고도화)

```
GET  /api/v2/qoo10jp/health                    # 헬스체크
GET  /api/v2/qoo10jp/shops                     # 활성 상점 목록
GET  /api/v2/qoo10jp/shops/:seller_id          # 상점 상세
POST /api/v2/qoo10jp/orders/collect            # 전체 상점 주문 수집
POST /api/v2/qoo10jp/orders/collect/:seller_id # 특정 상점 주문 수집
GET  /api/v2/qoo10jp/orders                    # 주문 조회 (필터링)
GET  /api/v2/qoo10jp/orders/stats              # 주문 통계
GET  /api/v2/qoo10jp/orders/:order_no          # 주문 상세
```

### Shopee 주문

```
POST /api/v1/shopee/orders/collect     # 주문 수집
GET  /api/v1/shopee/orders             # 주문 조회
GET  /api/v1/shopee/orders/:order_sn   # 주문 상세
```

### 스케줄러

```
POST /api/v1/scheduler/job             # 작업 생성
GET  /api/v1/scheduler/status          # 상태 조회
POST /api/v1/scheduler/worker/start    # 워커 시작
POST /api/v1/scheduler/worker/stop     # 워커 중지
POST /api/v1/scheduler/worker/count    # 워커 수 조정
```

## 아키텍처

### Worker 시스템

Redis Queue 기반의 분산 워커 시스템을 사용합니다:

1. **N8N 워크플로우**가 주기적으로 Supabase에서 활성 계정 조회
2. 각 계정 정보를 **Redis Queue** (`shopee_order_queue`)에 추가
3. **Go Worker들**이 Queue에서 작업을 가져와 주문 수집 실행
4. 수집 결과를 웹훅으로 알림

### 주요 특징

- **멀티 플랫폼 지원**: Qoo10JP, Shopee
- **자동 재시도**: 실패 시 최대 3회 재시도
- **웹훅 알림**: 수집 시작/완료/실패 알림
- **페이지네이션**: 대량 주문 처리
- **캐싱**: 중복 수집 방지

## 환경 변수

```bash
# Supabase
SUPABASE_URL=your_supabase_url
SUPABASE_ANON_KEY=your_anon_key
SUPABASE_SERVICE_KEY=your_service_key

# Redis
REDIS_URL=redis://localhost:6379
REDIS_PASSWORD=your_password

# Qoo10JP API
QOO10JP_BASE_URL=https://api.qoo10.jp

# Shopee API
SHOPEE_PARTNER_ID=your_partner_id
SHOPEE_PARTNER_KEY=your_partner_key
SHOPEE_BASE_URL=https://partner.shopeemobile.com

# Server
PORT=8080
GIN_MODE=debug

# Worker
WORKER_COUNT=3

# Webhook
ORDER_COLLECTION_WEBHOOK_URL=your_webhook_url

# Encryption (V2 - 선택사항, 토큰 암호화 시 필요)
ENCRYPTION_KEY=your_base64_encoded_256bit_key
```

## 데이터베이스 스키마

### orders_qoo10jp (V1 - Legacy)

```sql
CREATE TABLE orders_qoo10jp (
    id UUID PRIMARY KEY,
    order_no VARCHAR UNIQUE NOT NULL,
    order_date TIMESTAMP NOT NULL,
    order_status VARCHAR,
    payment_status VARCHAR,
    buyer_id VARCHAR,
    buyer_name VARCHAR,
    buyer_email VARCHAR,
    buyer_phone VARCHAR,
    total_amount DECIMAL(10,2),
    currency VARCHAR,
    shipping_address TEXT,
    platform_account_id VARCHAR,
    raw_data JSONB,
    synced_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
```

### V2 테이블 스키마

V2 테이블은 `scripts/create-qoo10jp-v2-tables.sql`에 정의되어 있습니다:

- **qoo10jp_shops_v2**: 상점 계정 관리 (암호화된 인증키 지원)
- **orders_qoo10jp_v2**: 확장된 주문 정보 (OMS 상태, 상세 필드)
- **order_items_qoo10jp_v2**: 주문 상품 정보
- **qoo10jp_sync_jobs**: 동기화 작업 추적 (중복 방지, 진행률)

```sql
-- V2 테이블 생성
\i scripts/create-qoo10jp-v2-tables.sql
```

### shopee_orders

```sql
CREATE TABLE shopee_orders (
    id UUID PRIMARY KEY,
    order_sn VARCHAR UNIQUE NOT NULL,
    platform_account_id VARCHAR,
    order_status VARCHAR,
    create_time TIMESTAMP,
    update_time TIMESTAMP,
    buyer_user_id BIGINT,
    buyer_username VARCHAR,
    total_amount DECIMAL(10,2),
    currency VARCHAR,
    items_json JSONB,
    created_at TIMESTAMP DEFAULT NOW()
);
```

## 개발

### 코드 스타일

```bash
go fmt ./...
golangci-lint run
```

### 테스트

```bash
go test ./...
```

### 빌드

```bash
# Windows
go build -o qoo10jp-order-go.exe cmd/main.go

# Linux
CGO_ENABLED=0 GOOS=linux go build -o qoo10jp-order-go cmd/main.go
```

## Docker 배포

```bash
docker build -t qoo10jp-order-go .
docker run -d --env-file .env -p 8080:8080 qoo10jp-order-go
```

## 라이센스

MIT License
