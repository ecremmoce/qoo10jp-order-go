# Quotation API (Entrance Ex.)

**카테고리**: API Reference
**난이도**: medium
**중요도**: 4/5
**최종 업데이트**: 2025-10-16T09:28:30

## 개요

본 가이드는 판매자가 쇼피에서 예상하는 배송비 및 배송 시간을 구매자에게 보낼 수 있도록 하는 Quotation API에 대한 정보를 제공합니다. 승인 및 통합 요구 사항, 요청 및 응답 흐름, 요청 및 응답에 대한 자세한 매개변수 설명을 설명합니다.

## 주요 키워드

- Quotation API
- Shipping Quote
- API Integration
- Request Parameters
- Response Parameters
- Shopee
- Logistics
- Shipping Costs
- Delivery Times

## 본문

# 견적 API (Entrance Ex.)

## API Rest Practice > Brand Upon API Best Prac... > 견적 API (Entrance Ex.

### 배송 견적 API
본 API는 Shoppe 자체 물류 API(배송 및 무게)에서 지원하는 제품 가격을 지정하는 데 사용할 수 있으며, Shoppe 영업팀에서 결정한 규칙에 따릅니다.

---

## 배송 견적 URL

이 도구를 사용하면 판매자는 Shoppe에서 예상한 배송 비용 및 배송 시간을 보낼 수 있으며, 이는 구매 시 구매자에게 표시됩니다. 관련 없는 결제 요구 사항에 따라 추가 판매 및 판매될 수 있습니다.

**참고:** 판매자 도구는 견적 URL을 직접 연결하거나 통합 API를 통해 연결하기 위해 Shoppe와 직접 통합하도록 선택했습니다.

---

## 승인/통합 요구 사항

Entrance Express에서 Shoppe에서 지원하는 저비용 통합을 사용하려면 견적 API에서 검색 및 완료하고 Shoppe 팀의 검토 및 승인이 필요합니다. Shoppe 팀의 승인을 받으려면 다음을 수행해야 합니다.

- Freight-Guard API (Lazada 파트너십)
- 판매자 보증 프로그램에 참여하고 제품 품질 지침을 따르십시오.
- URL은 특정 구조를 가져야 하며 구성에 유연성을 제공해야 합니다.
- 견적 및 주문 API 호출; 수신 및 응답은 3초 미만으로 변환될 수 있습니다.
- 견적 및 partner_id API 호출; 수신 및 응답은 1초 미만으로 변환될 수 있습니다.
- 판매자 센터를 통해 직접 지속적으로 등록하십시오.
- 견적 수정 및 완료된 호출에 대한 추적을 보장하여 나중에 백엔드에서 추적 처리를 수행합니다.
- Shoppe의 지연 배송율(LSR)/지연 배송율(LDR)/유효하지 않은 추적율(ITR)/반품 요구 사항을 충족해야 합니다.
   참고: 위의 요구 사항은 판매자의 "최고 지원" 또는 EXPRESS/통합을 통해 충족되어야 하며, Shoppe 지원 저비용 통합(해당 위치) 시스템에서 지원되는 운송업체를 검토하는 것이 좋습니다.
- 배송 API, 주문(픽업 준비 API) 및 추적 API와 통합하는 것이 좋습니다.
   참고: Entrance Express API가 견적 URL에만 사용되는 경우. API를 사용하려면 먼저 토큰을 생성해야 합니다.
- API를 통해 구매자가 재고를 추적하려는 욕구를 견적할 수 있습니다(모든 기본 전화).

통합을 수행하고 추적할 수 있는 위치: 사이트(필수)
- 수정 및 추적된 견적에 대한 모든 입력을 제출합니다. 그리고 그것; Shoppe 검토 (API에 대한 증거 서류)
- 모든 것(물류 등급을 매길 수 있는 PUT)을 가지고 견적 서류를 작성하고 API에 서류를 배포하시겠습니까?

이러한 모든 비율이 Shoppe에 완전히 전송되면(API 또는 도구를 통해) 또는 최소한 해당 커터가 있으면 추가 주문이 처리로 전송되는 것으로 반영되어 구매자는 결제 중에 이를 볼 수 있습니다.

---

## 견적 API 요청 및 응답 흐름

Shoppe는 요청 및 응답 매개변수(일부 중복)가 다른 두 가지 시나리오에서 견적 API를 호출합니다. 요청 + 응답 > [기본값]에 나열되어 있습니다.

[API 요청 및 응답 흐름을 보여주는 다이어그램]

---

## 견적 API 요청 및 응답 (기본)

### 요청 매개변수 (쿼리):

| 이름 | 예시 | HTTP 주소 |
|------|---------|--------------|
| URL | https://api.folder | TSP 또는 판매자가 제공한 URL |

| 이름 | 유형 | 예시 | 설명 |
|------|------|---------|-------------|
| partner_id | int | 1 | 성공적인 등록 후 할당된 파트너 ID (모든 요청에 필요) |
| timestamp | timestamp | 1670000000 | Shoppe에서 초 형식으로 제공하는 타임스탬프 (또는 모든 요청). 5분 후에 만료됩니다. |
| sign | string | e19d0e52237f91fab0f0a84e7c0311f6f1b4374bf0540a56e04d50689f51060d | 판매자가 생성한 서명. IP="quotation_id: api_key 호출; 타임스탬프 및 partner_key: api 키 (또는 HMAC+SHA256 해싱 알고리즘)는 https://open.shopee.com/documents/v2/v2.logistics.get_tracking?module=94&type=1에서 자세한 단계를 참조하십시오. |

### 기본 요청 매개변수

| 요청 매개변수 | | | | |
|--------------------|------|------|---------|-------------|
| 이름 | 유형 | 필수 | 예시 | 설명 |
| channel_id | int | TRUE | M006 | Shoppe의 물류 채널에 대한 고유 식별자 |
| entry_id | int | TRUE | 15224678 | 각 주문에 대한 고유 식별자 |
| origin_zip_code | string | TRUE | 12345000 | 판매자의 8자리 우편 번호, 마침표 및 대시 없는 숫자만 |
| destination_zip_code | string | TRUE | 12345000 | 구매자의 8자리 우편 번호, 마침표 및 대시 없는 숫자만 |
| destination_lat_long | object | TRUE | | 구매자 배송 위치의 위도 값은 -90.00 < x < 90.00 범위 내에 있어야 합니다. 이 값은 3PL이 사용할 수 있는 정확한 위치 주소를 확인하기 위해 제공됩니다. |
| longitude | float | TRUE | -23.5745841 | 구매자 배송 위치의 경도 값은 -180.00 < x < 180.00 범위 내에 있어야 합니다. 이 값은 3PL이 사용할 수 있는 정확한 위치 주소를 확인하기 위해 제공됩니다. |
| items | array | TRUE | | 제품 목록 |
| item_id | int | TRUE | 15224678 | Shoppe의 품목 식별 |
| title | string | FALSE | Item_title | 판매자가 등록한 SKU / (있는 경우) 추가 정보 |
| model_id | int | FALSE | 15224678 | Shoppe에 등록된 모델 식별 참고: PVP 공급에서 model_id는 항상 '0'이며 가격, 수량, 길이, 너비, 높이, 무게(숫자)도 마찬가지입니다. |
| model_title | string | FALSE | model_sku | Shoppe에 판매자가 등록한 모델 SKU (있는 경우) |
| category_id | int | TRUE | 15224678 | Shoppe에 등록된 카테고리 참고: PVP 공급에서 model_id는 항상 '0'이며 가격, 수량, 길이, 너비, 높이, 무게(숫자)도 마찬가지입니다. |
| quantity | int | TRUE | 15224678 | 제품 수량 품목 |
| price | float | FALSE | 150.4 | 제품 가격 |
| dimensions | object | TRUE | | 제품 크기 |
| length | int | TRUE | 10 | 1개 수량에 대한 센티미터 단위 길이 |
| width | int | TRUE | 10 | 1개 수량에 대한 센티미터 단위 너비 |
| height | int | TRUE | 10 | 1개 수량에 대한 센티미터 단위 높이 |
| weight | int | TRUE | 100 | 1개 수량에 대한 그램 단위 무게 |

### 응답 매개변수 / 견적 응답 (SP's)

| 이름 | 유형 | 필수 | 예시 | 설명 |
|------|------|-----------|---------|-------------|
| status | string | True | | 판매자가 만든 견적의 식별자 |
| message | string | True | | 구매자의 우편 번호 |
| request_id | string | True | | API 호출 식별자 |

### 응답 매개변수 / 견적 응답 (성공)

| 응답 매개변수 | | | | |
|---------------------|------|----------|---------|-------------|
| 이름 | 유형 | 필수 | 예시 | 설명 |

---

# API 문서

## 응답 매개변수 / 견적 응답 (성공)

### 응답 매개변수

| 이름 | 유형 | 필수 | 예시 | 설명 |
|------|------|-----------|---------|-------------|
| channel_id | int | TRUE | 90028 | 물류 채널에 대한 고유 식별자 |
| quotation_id | int | TRUE | 091234000 | 시스템에서 만든 견적의 식별자 |
| destination_zip_code | string | TRUE | 09123400 | 구매자의 우편 번호 |
| destination_url_lang | double | FALSE | - | "위치 x 경도" 또는 "구매자 위치"는 "-180 < x < 180"을 입력해야 합니다. |
| sellists | float | FALSE | -46.7354220599999 | 구매자 배송 위치의 위도 값은 "-90 < x < 90"을 입력해야 합니다. |
| longitude | float | FALSE | -23.5745491 | 구매자 배송 위치의 경도 값은 "-180 < x < 180"을 입력해야 합니다. |
| packages | array | TRUE | - | 패키지 목록 |
| dimensions | object | TRUE | - | 제품 크기 |
| width | int | TRUE | 10 | 센티미터 단위 너비 |
| length | int | TRUE | 10 | - |
| height | int | TRUE | 10 | 센티미터 단위 높이 |
| weight | int | TRUE | 100 | 그램 단위 무게 |
| sku | string | TRUE | - | 제품 SKU |
| seller_id | int | FALSE | 12345678 | Shopee의 식별 |
| sku | string | FALSE | sku_item | Shopee에 판매자가 등록한 SKU |
| model_id | int | FALSE | 12345678 | Shopee에 등록된 모델 식별 |
| model_sku | string | FALSE | model_sku | Shopee의 물류 채널에 등록된 모델 SKU |
| category_id | int | FALSE | 12345678 | Shopee에 등록된 품목 카테고리 |
| quantity | int | TRUE | 2 | 품목 수량 |
| price | int | TRUE | 130.4 | 제품 가격 |
| dimensions | object | TRUE | - | 제품 크기 |
| width | int | TRUE | 10 | 센티미터 단위 너비 |
| length | int | TRUE | 10 | - |
| height | int | TRUE | 10 | 센티미터 단위 높이 |
| weight | int | TRUE | 100 | 그램 단위 무게 |
| quotations | array | TRUE | - | 배송 견적 목록 |
| price | int | TRUE | 130.4 | 사용자에게 표시되는 배송비 |
| handling_time | int | TRUE | 20 | 분 단위 주문 준비 시간. 24시간(1440분)을 초과하는 경우 이 값에는 날짜 3이 포함됩니다. 날짜에는 시간 3(분 아님)으로 계산된 handling_time이 포함됩니다. |
| shipping_time | int | TRUE | 10 | 분 단위 주문 운송 시간(일) |
| promise_time | int | TRUE | 30 | 분 단위 처리 시간 + 배송 시간의 합계 |
| service_code | string | TRUE | M1020 | 다른 모델의 컨텍스트에서 운송업체를 식별하는 코드. 예: Correios API (SEDEX - AR, PAC - M1010); 외부 통합 API (예: "Correio Expresso - M1007") |
| letback_promise_time | int | TRUE | 30 | 분 단위 평균 주문 준비 및 발송 시간. SLA 작업 API 오류가 있는 경우 구매자에게 표시할 값 |

**오류 테이블에 매핑되지 않은 오류에 대한 기본 응답**

```json
*code#* : "내부 시스템 오류",
*message* : "내부 시스템 오류"
```

**예상되는 오류 코드:**

### HTTP 상태 코드

| HTTP 상태 코드 | 오류 | 메시지 | 조치/설명 |
|-----------------|-------|---------|-------------------|
| 400 | success/API_call_error | success/API_call_error Empty | False |
| 403 | error_forbidden_id | partner_id가 없습니다 (publish_id가 아님) | False |
| 403 | error_partner_id | partner_id가 유효하지 않습니다 | False |
| 404 | error_sign | 서명이 유효하지 않습니다 | False |
| 404 | error_sign | 쿼리에 서명이 없습니다 | False |
| 422 | error_timestamp | 타임스탬프가 유효하지 않습니다 | False |
| 423 | error_shop_id | shop_id가 없습니다 (본문에 있음) | False |
| 424 | error_shop_id | shop_id가 유효하지 않습니다 | False |
| 424 | invalid_arg_url_code | origin_zip_code가 유효하지 않습니다 | False |
| 424 | invalid_destination_PA_code | destination_zip_code가 유효하지 않습니다 | False |
| 424 | error_quotation_result | 견적 배송을 사용할 수 없습니다. | False |
| 400 | Invalid item_id | item_id가 유효하지 않습니다 | False |
| 400 | Invalid model_id | model_id가 유효하지 않습니다 | False |
| 400 | Invalid sku | sku가 유효하지 않습니다 | False |
| 400 | Invalid category_id | category_id가 유효하지 않습니다 | False |
| 400 | Invalid quantity | 수량이 유효하지 않습니다 | False |
| 400 | Invalid price | 가격이 유효하지 않습니다 | False |

---

# 견적 API 요청 및 응답 (Fastback)

## 요청 매개변수 (본문)

| 이름 | 유형 | 필수 | 예시 | 설명 |
|------|------|-----------|---------|-------------|
| msg_id | int | TRUE | 112345678 | 각 주문에 대한 고유 식별자 |

## 응답 매개변수 (본문)

| 이름 | 유형 | 필수 | 예시 | 설명 |
|------|------|-----------|---------|-------------|
| msg_id | int | TRUE | 112345678 | 각 주문에 대한 고유 식별자 |
| fastback_promise_time | int | TRUE | 30 | 분 단위 평균 주문 준비 및 배송 시간. 주문에 다른 물류 설정이 포함되어 있거나 품목 오류가 있는 경우 차이 값이 사용됩니다. |

## 응답 유효성 검사 (Shopee URL 포함)

- 업로드되면 유효성 검사 링크가 판매자에게 전송됩니다. 이는 필요한 수정 사항을 검증하고 보고하는 데 사용해야 하는 링크입니다.
- URL: "https://seller.shopee.ph/gou-chat-api-link?to-secondarycomplication_quotation_anwsers"
- URL에 대한 추가 지원 및 Shopee와의 완전한 구현을 위해 판매자는 응답 형식 계산 URL을 사용하는 방법을 보여주는 예제 프로세스입니다.

## 응답 시간 제한 (견적)

- 결제가 구매 중 사용자 경험에서 가장 중요한 순간임을 고려할 때 구매자에게 최상의 경험을 보장하는 것이 필수적입니다.
- 따라서 이 프로세스의 예상 시간 제한은 8~10초 사이입니다.
- 2000초보다 긴 응답 시간은 허용되지 않으므로 통합이 불가능합니다. 따라서 최상의 사용자 경험을 달성하려면 이 매개변수를 최적화하는 것이 중요합니다.

---

# 비상 테이블

비상 화물은 기본 화물 계산입니다. 화물 견적 API에 문제가 있는 경우 (요청에 대한 상태가 최소 500)

## 비상 사태에 대한 배송 값 설정

비상 사태가 작동하려면 판매자는 배송 값을 등록하고 Foreign Express 채널을 활성화해야 합니다.
배송 값을 구성하려면 메뉴에 액세스하십시오.

**판매자 센터로 이동하여 배송 설정으로 이동하고 채널 설정 버튼을 클릭하십시오.**

[배송 설정 구성 패널을 보여주는 인터페이스 스크린샷]

---

# 비상 화물 값 입력

다음으로 화면에서 배송비를 입력하라는 메시지가 표시됩니다. 배송비는 판매 채널 값 또는 도시별로 입력할 수 있습니다. 이 설정은 각 내 도시/지역마다 다를 수 있으며, 주문이 필요한 장소에서 활성화될 때 비상 차트에 설정된 품목을 제외하고 항상 일반 품목을 보내는 모든 주문에 자동으로 적용됩니다.

## 상태별 비용을 설정하는 단계:

**상태별 배송비:**

1. 상태를 선택하십시오
2. 배송비 필드를 채우십시오
3. 저장을 클릭하십시오

[상태별 배송비 구성을 보여주는 인터페이스 스크린샷]

**도시별 배송비:**

1. 상태를 선택하십시오
2. '이 지역에 대한 값 설정' 버튼을 선택하십시오
3. 새 화면이 나타납니다. 배송비를 알리고 싶은 도시를 선택하십시오
4. 해당 도시 옆에 값을 채우십시오
5. 저장을 클릭하십시오

[도시별 배송비 구성을 보여주는 인터페이스 스크린샷]

---

# 동기화 영역 설정

이 기능은 구매자의 주소가 판매자의 배송 지역 내에 있는지 여부를 식별하고 주문이 실패한 주문으로 표시되는 경우를 포함하여 다른 문제가 있는지 여부를 식별하기 위해 판매자의 서비스 가능 지역 설정을 지원하기 위한 것입니다. 배송지 기능(내 배송 > 배송 설정의 Shopee 판매자 센터에서 찾을 수 있음)을 채우고 항상 구매자에게 지원되지 않음을 표시하십시오. 서비스 가능 설정을 업로드하려면:

1. https://partner.shopeemobile.com/api/v1/top에 로그인하십시오
2. 물류 -> 내 배송 > 물류 설정으로 이동하십시오
3. 업로드되면 https://partner.shopeemobile.com/api/v1/me를 통해 서비스 가능 지역 설정 업로드 상태를 확인하십시오.

---

# 배송 준비

구매자가 주문하면 판매자 APP입니다. 배송을 준비한 후 주문이 항상 PROCESSED 상태로 업데이트되고 get_package_detail API, get_order_detail order_status로 반환 상태 "READY_TO_SHIP"으로 반환될 수 있습니다.

API 문서에 대한 자세한 단계:
- /api/v2/logistics/get_order_logistics_info
- /api/v2/logistics/ship_order
- /api/v2/logistics/get_tracking_info

---

# 주문 상태 및 추적

주문 상태 및 추적 업데이트는 추적 업데이트 API를 통해 수행되며 위의 API 전체에서 사용되는 동일한 식별자인 주문 번호(주문 SN)를 기반으로 합니다. Shopee는 주문 및 배송 상태에 대한 엔드 투 엔드 모니터링을 수행합니다.

---

# 주문 상태 및 추적

## 배송 준비

판매자는 ship_order API를 통해 배송을 준비합니다. 배송을 준비한 후 order_status가 PROCESSED 상태로 업데이트되고 get_package_detail API에서 반환됩니다. get_order_detail order_status는 여전히 READY_TO_SHIP을 반환합니다.

API 문서에 대한 링크는 다음과 같습니다.

- /api/v2/logistics/ship_order
- /api/v2/order/get_order_detail
- /api/v2/order/get_package_detail

---

## 주문 상태 및 추적

주문 상태 및 추적 업데이트는 추적 업데이트 API를 통해 수행되며 Shopee에서 생성된 주문 번호(주문 ID, OrderSN이라고도 함)를 기반으로 합니다. 추적 번호도 필요합니다.

추적 번호와 주문 상태를 업데이트하려면 **v2 logistics update_tracking_status (OpenAPI) 엔드포인트**를 사용하십시오. 가능한 상태는 다음과 같습니다.

1. **주문 발송됨 (logistics_pickup_done)**
   - 주문 상태를 발송됨으로 업데이트할 때 URL과 추적 번호를 보낼 수 있습니다.

2. **주문 배송 완료 (logistics_delivery_done)**
   - 배송 완료 상태는 주문이 이미 발송됨 상태인 경우에만 수신됩니다.
   - OTP 코드 확인: Entrega Expressa 채널의 경우 OTP 코드 확인이 필수입니다. Shopee는 판매자가 배송을 준비한 후 OTP 코드를 생성합니다. 판매자가 구매자에게 소포를 배송할 때 판매자는 구매자로부터 OTP 코드를 검색하고 추적 상태를 logistics_delivery_done으로 업데이트하려고 할 때 API 요청으로 보내야 합니다.

3. **배송 실패 (logistics_delivery_failed)**
   - 배송 실패 상태는 주문이 이미 발송됨 상태인 경우에만 수신됩니다.
   - 실패 이유: Entrega Expressa 채널의 경우 판매자는 추적 상태를 logistics_delivery_failed로 업데이트할 때 failed_reason을 보내야 합니다.

**중요**: 주문 배송 완료 또는 배송 실패 상태를 보낸 후에는 더 이상 상태 업데이트가 허용되지 않습니다. 둘 다 최종 상태이기 때문입니다.

*tracking_number 및 tracking_url 매개변수는 상태를 logistics_pickup_done으로 업데이트할 때만 보내야 합니다*.

API 문서에 대한 링크는 **/api/v2/logistics/update_tracking_status**입니다.

---

## 채널 주문 흐름 및 상태
판매자는 ship_order API를 통해 배송을 준비합니다. 배송 준비 후 order_status가 PROCESSED 상태로 업데이트되며, get_package_detail API에서 반환됩니다. get_order_detail의 order_status는 여전히 READY_TO_SHIP을 반환합니다.

API 문서에 대한 링크는 다음과 같습니다.

- /api/v2/logistics/ship_order
- /api/v2/order/get_order_detail
- /api/v2/order/get_package_detail

---

## 주문 상태 및 추적

주문 상태 및 추적 업데이트는 Tracking Update API를 통해 수행되며, Shopee에서 생성된 주문 번호(Order ID, OrderSN이라고도 함)를 기준으로 합니다. 추적 번호도 필요합니다.

추적 번호와 주문 상태를 업데이트하려면 **v2 logistics update_tracking_status (OpenAPI) 엔드포인트**를 사용하십시오. 가능한 상태는 다음과 같습니다.

1. **주문 발송 완료 (logistics_pickup_done)**
   - 주문 상태를 Shipped로 업데이트할 때 URL 및 추적 번호를 보낼 수 있습니다.

2. **주문 배송 완료 (logistics_delivery_done)**
   - 배송 완료 상태는 주문이 이미 Shipped 상태인 경우에만 수신됩니다.
   - OTP 코드 확인: Entrega Expressa 채널의 경우 OTP 코드 확인이 필수입니다. Shopee는 판매자가 배송을 준비한 후 OTP 코드를 생성합니다. 판매자가 구매자에게 소포를 배송할 때 구매자로부터 OTP 코드를 받아 tracking_status를 logistics_delivery_done으로 업데이트하려고 할 때 API 요청에 보내야 합니다.

3. **배송 실패 (logistics_delivery_failed)**
   - 배송 실패 상태는 주문이 이미 Shipped 상태인 경우에만 수신됩니다.
   - 실패 사유: Entrega Expressa 채널의 경우 판매자가 tracking_status를 logistics_delivery_failed로 업데이트할 때 failed_reason을 보내야 합니다.

**중요**: 주문 배송 완료 또는 배송 실패 상태를 보낸 후에는 더 이상 상태 업데이트가 허용되지 않습니다. 둘 다 최종 상태이기 때문입니다.

*tracking_number 및 tracking_url 매개변수는 상태를 logistics_pickup_done으로 업데이트할 때만 보내야 합니다*.

API 문서에 대한 링크는 **/api/v2/logistics/update_tracking_status**입니다.

---

## 채널 주문 흐름 및 상태

### REQUESTED FLOW

```
UNPAID
  ↓ 1
READY_TO_SHIP
  ↓ 2
PROCESSED
  ↓ 3
SHIPPED
  ↓ 4
TO_CONFIRM_RECEIVE
  ↓ 5
COMPLETED
```

**단계:**

1. **(1) 구매자가 주문을 결제하고 판매자가 송장을 업로드합니다.**

2. **(2) 판매자가 배송을 준비합니다 (ship_order API).**

3. **(3) 판매자가 (update_tracking_status API)를 통해 주문 상태를 PICK_UP_DONE으로 업데이트합니다.**

4. **(4) 판매자가 (update_tracking_status API)를 통해 주문 상태를 DELIVERY_DONE으로 업데이트합니다.**

5. **(5) 구매자가 패키지 수령을 확인합니다.**

## 사용 사례

1. 구매자의 결제 시 배송비 계산
2. 판매자 시스템과 쇼피의 물류 API 통합
3. 예상 배송 시간을 구매자에게 표시
4. 배송 견적 프로세스 자동화
5. 구매자의 배송 위치 확인

## 관련 API

- Shipping Quote API
- Shipment API
- Order (Arrange Pickup API)
- Tracking API

---

## 원문 (English)

### Summary

This guide provides information about the Quotation API, which allows sellers to send Shopee-estimated shipping costs and delivery times to buyers. It outlines the requirements for approval and integration, request and response flows, and detailed parameter descriptions for both requests and responses.

### Content

# Quotation API (Entrance Ex.)

## API Rest Practice > Brand Upon API Best Prac... > Quotation API (Entrance Ex.

### Shipping Quote API
This API is available to specify where prices products are out supported by Shoppe's own logistics API (a shipments and weight, in accordance with the rules determined by Shoppe's sales team.

---

## Shipping Quote URL

This tool allows sellers to send Shoppe-estimated shipping costs and delivery times, which will be displayed to buyers at the time of purchase. It can be further added sold and sold per the unrelated payment requirements.

**Note:** The seller tool has opted to integrate directly with Shoppe to connected their quote URL or connect via an integration API.

---

## Requirements for approval/integration

Seek and complete from the Quotation API to use Shoppe-supported low Cost Integration from Entrance Express (where) requires review and approval from the Shoppe team before it can go live. To be approved by the Shoppe team, you must:

- Freight-Guard API (Lazada Partnership)
- Opt into the Sellers Warranty Program and follow the Product Quality guidelines
- The URL has be specific, structure, providing flexibility in configuration
- Quote and order API calls; receipt and respond can transform less than 3sec(s.)
- Quote and partner_id calls API; receipt and response transform less than 1 sec(s.)
- Continuously keep registered directly through the Seller Center
- Ensure tracking for quote revised and completed call the backend for processing tracking in a later stage
- Must meet Shoppe's Late Ship Rate (LSR)/Late Delivery Rate (LDR)/ Invalid Tracking Rate (ITR)/ Return requirements
   Note: The above requirements must be met by the seller's "best support" or their EXPRESS/Integration, it's recommended to review to review from Shoppe Supported Low Cost Integration (where) system to be supported carrier.
- It is strongly encouraged to integrate with the Shipment API, Order (Arrange Pickup API) and Tracking API
   Note: If the Entrance Express API is only used for quotation URL's. To use any API, it is necessary to create a token first.
- Since the API allows for quoting the Buyer's desire to track inventory (all the native phone).

In an able to perform the integration and track where: it is site (necessary)
- Submit all inputs to quote revised and tracked; and it; Shoppe-review (evidence letters to the API)
- To have all (PUT that is capable of rating Logistics) to quote letters your and distribute letters to the API?)

Once all these rates are completely sent to Shoppe (through the API or a tool) or at least if their cutters, which will be reflected as further order sent to the processing, buyers will be able to see it during checkout.

---

## Quotation API Request and Response Flow

Shoppe will be calling Quotation API in 2 scenarios, which have different request and response parameters (some overlap) as listed in Request + Response > [Default]

[Diagram showing API Request and Response Flow]

---

## Quotation API Request and Response (Default)

### Request parameters (query):

| Name | Example | HTTP Address |
|------|---------|--------------|
| URL | https://api.folder | URL provided by TSP or seller |

| Name | Type | Example | Description |
|------|------|---------|-------------|
| partner_id | int | 1 | The partner ID assigned after successful registration (Required for all requests) |
| timestamp | timestamp | 1670000000 | Timestamp from Shoppe in second format (or all requests). Expires in 5 minutes |
| sign | string | e19d0e52237f91fab0f0a84e7c0311f6f1b4374bf0540a56e04d50689f51060d | Signature generated by the seller. IP="quotation_id: api_key call; timestamp and partner_key: api key (or HMAC+SHA256 hashing algorithm) provided at https://open.shopee.com/documents/v2/v2.logistics.get_tracking?module=94&type=1 for detail steps. |

### Base Request Parameters

| Request Parameters | | | | |
|--------------------|------|------|---------|-------------|
| Name | Type | Mandatory | Example | Description |
| channel_id | int | TRUE | M006 | Unique identifier for the logistic channel in Shoppe |
| entry_id | int | TRUE | 15224678 | Unique identifier for each order |
| origin_zip_code | string | TRUE | 12345000 | Seller's zip code with 8 digits, only numbers without periods and dashes |
| destination_zip_code | string | TRUE | 12345000 | Buyer's zip code with 8 digits, only numbers without periods and dashes |
| destination_lat_long | object | TRUE | | Latitude of buyer's delivery location Value should be within -90.00 < x < 90.00. This value is provided for 3PL to verify the exact location address they may use |
| longitude | float | TRUE | -23.5745841 | Longitude of buyer's delivery location Value should be within -180.00 < x < 180.00. This value is provided for 3PL to verify the exact location address they may use |
| items | array | TRUE | | Product list |
| item_id | int | TRUE | 15224678 | Item Identification on Shoppe |
| title | string | FALSE | Item_title | SKU / registered by the seller (if) additional information |
| model_id | int | FALSE | 15224678 | Identification of the model registered on Shoppe Note: In PVP supply, the model_id is Always '0', As well as price, qty, length, width, height, weight (numerals) |
| model_title | string | FALSE | model_sku | Model SKU registered by the seller (if) on Shoppe |
| category_id | int | TRUE | 15224678 | Category registered on Shoppe Note: In PVP supply, the model_id of is Always '0', As well as price, qty, length, width, height, weight (numerals) |
| quantity | int | TRUE | 15224678 | Product's quant items |
| price | float | FALSE | 150.4 | Product's price |
| dimensions | object | TRUE | | Product dimensions |
| length | int | TRUE | 10 | Length in centimeters for 1 quantity |
| width | int | TRUE | 10 | Width in centimeters for 1 quantity |
| height | int | TRUE | 10 | Height in centimeters for 1 quantity |
| weight | int | TRUE | 100 | Weight in grams for 1 quantity |

### Response Parameters / Quote Response (SP's)

| Name | Type | Mandatory | Example | Description |
|------|------|-----------|---------|-------------|
| status | string | True | | Identifier of the quote made by the seller |
| message | string | True | | Buyer's zip code |
| request_id | string | True | | API call identifier |

### Response Parameters / Quote Response (success)

| Response Parameters | | | | |
|---------------------|------|----------|---------|-------------|
| Name | Type | Mandatory | Example | Description |

---

# API Documentation

## Response Parameters / Quote Response (success)

### Response Parameters

| Name | Type | Mandatory | Example | Description |
|------|------|-----------|---------|-------------|
| channel_id | int | TRUE | 90028 | Unique identifier for the logistic channel |
| quotation_id | int | TRUE | 091234000 | Identifier of the quote made by the system |
| destination_zip_code | string | TRUE | 09123400 | Buyer's zip code |
| destination_url_lang | double | FALSE | - | "Location x Longitude" or "buyers location" should be enter "-180 < x < 180" |
| sellists | float | FALSE | -46.7354220599999 | Latitude of buyer's delivery location. Value should be enter "-90 < x < 90" |
| longitude | float | FALSE | -23.5745491 | Longitude of buyer's delivery location. Value should be enter "-180 < x < 180" |
| packages | array | TRUE | - | Package List |
| dimensions | object | TRUE | - | Product dimensions |
| width | int | TRUE | 10 | Width in centimeters |
| length | int | TRUE | 10 | - |
| height | int | TRUE | 10 | Height in centimeters |
| weight | int | TRUE | 100 | Weight in grams |
| sku | string | TRUE | - | Product sku |
| seller_id | int | FALSE | 12345678 | Identification on Shopee |
| sku | string | FALSE | sku_item | SKU registered by the seller on Shopee |
| model_id | int | FALSE | 12345678 | Identification of the model registered on Shopee |
| model_sku | string | FALSE | model_sku | Model SKU registered by the logistic channel on Shopee |
| category_id | int | FALSE | 12345678 | Item category registered on Shopee |
| quantity | int | TRUE | 2 | Quantity of items |
| price | int | TRUE | 130.4 | Product price |
| dimensions | object | TRUE | - | Product dimensions |
| width | int | TRUE | 10 | Width in centimeters |
| length | int | TRUE | 10 | - |
| height | int | TRUE | 10 | Height in centimeters |
| weight | int | TRUE | 100 | Weight in grams |
| quotations | array | TRUE | - | List of shipping quotes |
| price | int | TRUE | 130.4 | Shipping cost displayed to Users |
| handling_time | int | TRUE | 20 | Order preparation time in minutes. If it exceeds 24 hours (1440 minutes), this value includes the date 3. The date includes the handling_time calculated in hours 3 (not minutes) |
| shipping_time | int | TRUE | 10 | Order transit time in days in minutes |
| promise_time | int | TRUE | 30 | Sum of handling time + shipping time in minutes |
| service_code | string | TRUE | M1020 | Code that identifies the carrier in the context of the other model. Example: Correios API (SEDEX - AR, PAC - M1010); External Integrations API (ex: "Correio Expresso - M1007") |
| letback_promise_time | int | TRUE | 30 | Average order preparation and dispatch time in minutes. Value to be shown to the buyer when there is an SLA operation API error |

**Default response for errors not mapped in the error table**

```json
*code#* : "internal system error",
*message* : "internal system error"
```

**Expected error codes:**

### HTTP status code

| HTTP status code | Error | Message | Action/Description |
|-----------------|-------|---------|-------------------|
| 400 | success/API_call_error | success/API_call_error Empty | False |
| 403 | error_forbidden_id | there is no partner_id (it didn't publish_id) | False |
| 403 | error_partner_id | partner_id is invalid | False |
| 404 | error_sign | your sign is invalid | False |
| 404 | error_sign | there is no sign in query | False |
| 422 | error_timestamp | your timestamp is invalid | False |
| 423 | error_shop_id | there is no shop_id (in body) | False |
| 424 | error_shop_id | The shop_id is invalid | False |
| 424 | invalid_arg_url_code | The origin_zip_code is invalid | False |
| 424 | invalid_destination_PA_code | The destination_zip_code is invalid | False |
| 424 | error_quotation_result | The quotation shipment is unavailable. | False |
| 400 | Invalid item_id | The item_id is invalid | False |
| 400 | Invalid model_id | The model_id is invalid | False |
| 400 | Invalid sku | The sku is NOT valid | False |
| 400 | Invalid category_id | The category_id is invalid | False |
| 400 | Invalid quantity | The quantity is invalid | False |
| 400 | Invalid price | The price is invalid | False |

---

# Quotation API Request and Response (Fastback)

## Request Parameters (Body)

| Name | Type | Mandatory | Example | Description |
|------|------|-----------|---------|-------------|
| msg_id | int | TRUE | 112345678 | Unique identifier for each order |

## Response Parameters (Body)

| Name | Type | Mandatory | Example | Description |
|------|------|-----------|---------|-------------|
| msg_id | int | TRUE | 112345678 | Unique identifier for each order |
| fastback_promise_time | int | TRUE | 30 | Average order preparation and shipping time in minutes. Difference value will be used if the order contains different logistics settings or when there are item fails. |

## Response validation (with Shopee URL)

- Once uploaded, a validation link will be sent to seller. It is a link that should be used to validate and report any necessary corrections.
- URL: "https://seller.shopee.ph/gou-chat-api-link?to-secondarycomplication_quotation_anwsers"
- For further assist with the URL and fully implemented with Shopee, seller is an example process showing how to use the response format calculation URL.

## Response Time Limit (of the quotation)

- Considering that checkout is the most crucial moment in the user experience during the purchase, it is essential to ensure the best possible experience for the buyer.
- Therefore, the estimated time limit of this process is between 8 to 10 seconds
- Response times longer than 2000s will not be permitted, making integration impossible. Therefore, optimizing this parameter is crucial achieving the best possible user experience.

---

# Contingency Table

The Contingency Freight is the default freight calculation. In case there is a problem with freight quotation out API (Status at least 500 for requests)

## Shipping Value Settings for Contingency

For the contingency to work, the seller must register the shipping values and enable the Foreign Express channel.
Access the menu to configure your Shipping Values.

**To Seller Centre, go to Shipping Settings and click the Channel Settings button**

[Interface screenshots showing shipping settings configuration panels]

---

# Filling in the contingency freight value

Next, the screen will ask to enter the shipping cost. Shipping costs can be entered by sales channel value or by city. This setting can vary for each my city/region which will automatically apply for all orders always send normal items except that are established in the contingency chart when they are activated from the place where orders need to go.

## Steps to set per cost state:

**Stage shipping cost per state:**

1. Select a state
2. Fill in the shipping cost field
3. Click Save

[Interface screenshots showing state-by-state shipping cost configuration]

**Stage shipping cost per city:**

1. Select a state
2. Select the 'I set value button for this area'
3. A new screen will appear. Select the cities you want to inform the shipping cost
4. Fill in the values next to the respective cities
5. Click Save

[Interface screenshots showing city-specific shipping cost configuration]

---

# Syncronization Area Settings

This feature is intended to support seller's serviceable area settings, to identify if the buyer's address is within seller's delivery area and to know if there are any other issues, including when an order is marked as a failed order. Please fill in the Shipping Destination feature (can be found in Shopee Seller Centre on My Shipping > Shipping Settings), and always display an unsupported to buyer. To upload serviceable settings:

1. Login to https://partner.shopeemobile.com/api/v1/top
2. Navigate to Logistics -> My Shipping > Logistics Settings
3. Once uploaded, tweak click check the status of serviceable area settings upload via https://partner.shopeemobile.com/api/v1/me

---

# Arrange Shipment

Once an order is placed by a buyer, it is the Seller APP. Please note the order always will be updated to PROCESSED status after arranging shipment, and can be returned to get_package_detail API, get_order_detail order_status with return status "READY_TO_SHIP"

For detailed steps on API documentation:
- /api/v2/logistics/get_order_logistics_info
- /api/v2/logistics/ship_order
- /api/v2/logistics/get_tracking_info

---

# Order Status and Tracking

Order status and tracking updates will be performed via the Tracking Update API and will be based on the order number (order SN) which is the same identifier used throughout the above API. Shopee takes end-to-end monitoring of the order and delivery status.

---

# Order Status and Tracking

## Shipment Arrangement

Seller will arrange shipment via ship_order API. Please note that order_status will be updated to PROCESSED status after arranging shipment, and will be returned in get_package_detail API. get_order_detail order_status will still return READY_TO_SHIP.

Here are the links to the API documentation:

- /api/v2/logistics/ship_order
- /api/v2/order/get_order_detail
- /api/v2/order/get_package_detail

---

## Order Status and Tracking

Order status and tracking updates will be performed via the Tracking Update API and will be based on the order number created in Shopee (Order ID, also known as OrderSN). The tracking number will also be required.

To update the tracking number and order status, use the **v2 logistics update_tracking_status (OpenAPI) endpoint**. The possible statuses are:

1. **Order Sent (logistics_pickup_done)**
   - The URL and tracking number can be sent when updating the order status to Shipped.

2. **Order Delivered (logistics_delivery_done)**
   - The Delivered status will only be received if the order already has the Shipped status.
   - OTP code verification: For Entrega Expressa channel, it is mandatory to have OTP code verification. Shopee will be generating the OTP code after seller arranges shipment. When seller is delivering the parcel to buyer, it is mandatory for seller to retrieve OTP code from buyer, and send it in API request when attempting to update tracking_status to logistics_delivery_done.

3. **Delivery Failure (logistics_delivery_failed)**
   - The Delivery Failure status will only be received if the order already has the Shipped status.
   - Failed reason: For Entrega Expressa channel, it is mandatory for seller to send failed_reason when updating tracking_status to logistics_delivery_failed.

**IMPORTANT**: after sending the Order Delivered or Delivery Failed statuses, no further status updates will be allowed, as both are finalizing statuses;

*the tracking_number and tracking_url parameters should only be sent when updating the status to logistics_pickup_done*.

Here is a link to the API documentation **/api/v2/logistics/update_tracking_status**.

---

## Channel Order Flow and Status

### REQUESTED FLOW

```
UNPAID
  ↓ 1
READY_TO_SHIP
  ↓ 2
PROCESSED
  ↓ 3
SHIPPED
  ↓ 4
TO_CONFIRM_RECEIVE
  ↓ 5
COMPLETED
```

**Steps:**

1. **(1) Buyer paid the order and seller upload the invoice**

2. **(2) Seller Arrange Shipment (ship_order API)**

3. **(3) Seller updates order status as PICK_UP_DONE via (update_tracking_status API)**

4. **(4) Seller updates order status as DELIVERY_DONE via (update_tracking_status API)**

5. **(5) Buyer confirms receive the package**

---

**문서 ID**: developer-guide.697
**플랫폼**: shopee
**URL**: https://open.shopee.com/developer-guide/697
**처리 완료**: 2025-10-16T09:28:30
