# 📋 Shopee API 분석 결과 요약

## 📁 생성된 파일

이 디렉토리에는 Shopee Open Platform API에 대한 완전한 분석 자료가 포함되어 있습니다.

### 주요 문서

1. **`SHOPEE_API_COMPLETE_GUIDE.md`** (⭐ 메인 문서)
   - 362개 전체 API에 대한 완전한 가이드
   - 27개 카테고리별 분류 및 상세 설명
   - HTTP 메서드, 설명, 문서화 상태 포함
   - 사용 사례 및 베스트 프랙티스 제공

2. **`SHOPEE_API_COMPLETE_ANALYSIS.md`** (기본 분석)
   - API 구조 및 카테고리 개요
   - 통계 정보 중심

## 📊 주요 통계

- **총 API 개수**: 362개
- **카테고리 개수**: 27개
- **문서화율**: 
  - 메타데이터: 99% (360개)
  - 마크다운: 0.5% (2개)
  - 한글 번역: 1% (4개)

## 🏆 상위 카테고리 (API 개수 기준)

| 순위 | 카테고리 | API 개수 | 비율 |
|:----:|---------|:--------:|:----:|
| 1 | Product | 55 | 15.2% |
| 2 | Logistics | 41 | 11.3% |
| 3 | Global Product | 34 | 9.4% |
| 4 | Ads | 26 | 7.2% |
| 5 | Livestream | 25 | 6.9% |

## 🎯 카테고리별 주요 용도

### 🛒 Product (55개 API)
상품 관리의 핵심 기능
- 상품 등록/수정/삭제
- 재고 및 가격 관리
- 카테고리 및 속성 관리
- 상품 검색 및 조회

**주요 API**:
- `add_item` - 새 상품 등록
- `update_item` - 상품 정보 수정
- `get_item_list` - 상품 목록 조회
- `update_stock` - 재고 업데이트

### 📦 Order (21개 API)
주문 처리 및 관리
- 주문 조회 및 상세 정보
- 주문 취소 및 분할
- 송장 생성 및 관리
- 배송 준비 처리

**주요 API**:
- `get_order_list` - 주문 목록 조회
- `get_order_detail` - 주문 상세 정보
- `cancel_order` - 주문 취소
- `generate_fbs_invoices` - 송장 생성

### 🚚 Logistics (41개 API)
물류 및 배송 관리
- 배송 방법 설정
- 운송장 번호 생성 및 관리
- 배송 문서 생성
- 배송 상태 추적

**주요 API**:
- `get_channel_list` - 배송 채널 조회
- `get_tracking_number` - 운송장 번호 발급
- `ship_order` - 배송 처리
- `create_shipping_document` - 배송 문서 생성

### 💰 Payment (16개 API)
결제 및 정산 관리
- 에스크로 정보 조회
- 수익 리포트 생성
- 정산 내역 확인
- 지갑 거래 관리

**주요 API**:
- `get_escrow_list` - 에스크로 목록
- `generate_income_report` - 수익 리포트
- `get_payout_info` - 정산 정보

### 🏷️ Discount & Promotion (여러 카테고리)
다양한 프로모션 관리
- **Discount** (9개): 기본 할인
- **Bundle Deal** (10개): 번들 상품
- **Add-on Deal** (14개): 추가 상품
- **Shop Flash Sale** (11개): 플래시 세일
- **Voucher** (6개): 쿠폰/바우처

### 📢 Ads (26개 API)
광고 캠페인 관리
- CPC 광고 생성 및 관리
- 상품 광고 설정
- GMS 캠페인
- 광고 성과 분석

### 📹 Livestream (25개 API)
라이브 커머스
- 라이브 세션 생성 및 관리
- 상품 연동
- 댓글 관리
- 성과 지표 조회

## 🔧 사용 방법

### 1. 필수 인증 정보

```
Partner ID: 파트너 식별자
Partner Key: API 서명 키
Access Token: 샵별 액세스 토큰
Shop ID: 대상 샵 ID
```

### 2. API 호출 형식

```
Base URL: https://partner.shopeemobile.com
API Path: /api/v2/{category}/{method}
```

### 3. 공통 헤더

```
Content-Type: application/json
partner_id: {파트너 ID}
timestamp: {Unix 타임스탬프}
sign: {HMAC-SHA256 서명}
access_token: {액세스 토큰}
shop_id: {샵 ID}
```

## 📖 빠른 시작 가이드

### 상품 등록 플로우

1. **카테고리 조회**: `v2.product.get_category`
2. **속성 조회**: `v2.product.get_attribute_tree`
3. **상품 등록**: `v2.product.add_item`
4. **재고 설정**: `v2.product.update_stock`

### 주문 처리 플로우

1. **주문 목록 조회**: `v2.order.get_order_list`
2. **주문 상세 확인**: `v2.order.get_order_detail`
3. **배송 준비**: `v2.logistics.ship_order`
4. **운송장 발급**: `v2.logistics.get_tracking_number`

## 🔍 API 찾는 방법

### 카테고리별로 찾기
`SHOPEE_API_COMPLETE_GUIDE.md`의 목차에서 카테고리 선택

### 기능별로 찾기
다음 키워드로 문서 내 검색:
- **조회**: `get_*`
- **생성**: `add_*`, `create_*`
- **수정**: `update_*`, `edit_*`
- **삭제**: `delete_*`, `remove_*`

## 📚 참고 자료

### 공식 문서
- [Shopee Open Platform](https://open.shopee.com/)
- [Developer Guide](https://open.shopee.com/developer-guide/4)
- [API Documentation](https://open.shopee.com/documents)

### 로컬 리소스
- `workspaces/` - 각 API별 상세 분석 자료
- `api_specifications/` - OpenAPI 스펙
- `content_registry.json` - API 메타데이터

## ⚠️ 주의사항

1. **Rate Limiting**: API 호출 제한 주의
2. **인증 만료**: 주기적인 토큰 갱신 필요
3. **국가별 차이**: 일부 API는 특정 국가만 지원
4. **버전 관리**: 현재 v2 API 기준

## 🔄 업데이트 이력

- **2025-10-16**: 초기 분석 완료
  - 362개 API 전체 분석
  - 27개 카테고리 분류
  - 완전 가이드 문서 생성

## 📞 지원

문제가 있거나 질문이 있는 경우:
1. 공식 문서 확인
2. Shopee Open Platform 지원팀 문의
3. 로컬 분석 자료 참조

---

**생성 정보**
- 생성일: 2025-10-16
- 분석 도구: `analyze_shopee_apis_detailed.py`
- API 버전: v2

