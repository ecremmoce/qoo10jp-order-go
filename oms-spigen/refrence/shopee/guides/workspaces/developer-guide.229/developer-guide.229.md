# 주문 관리

**카테고리**: 모범 사례
**난이도**: medium
**중요도**: 5/5
**최종 업데이트**: 2025-10-16T08:52:03

## 개요

본 개발자 가이드는 주문 및 패키지 상태 흐름, 주문 및 패키지 검색을 위한 API 엔드포인트, 주문 취소 및 주문 분할을 포함한 주문 관리에 대한 포괄적인 정보를 제공합니다. 또한 배송 API 호출 흐름 및 기본 배송 로직을 다룹니다.

## 주요 키워드

- 주문 관리
- 주문 상태
- 패키지 상태
- API Endpoints
- 주문 분할
- Shipment API
- Shipping Logic

## 본문

# 주문 관리

## 1. 소개

API 모범 사례 > 주문 관리

**아이템:** 주문 내의 개별 제품. 단위 수량 및 기타 세부 정보. 아이템은 배송을 위해 패키지에 포함됩니다.

---

## 2. 주문 상태 흐름

### 주문 상태 워크플로

**주문 진행 상황을 보여주는 순서도:**

```
PLACED
↓
PAYMENT_FAILED
↓
ACCEPTED → IN_PREPARATION
↓
READY_TO_SHIP
↓
SHIPPED
↓
IN_TRANSIT
↓
DELIVERED
↓
COMPLETED
```

**사이드 브랜치:**
- CANCELED (여러 단계에서 발생 가능)
- READY_TO_SHIP 또한 CANCELED로 이어질 수 있음

### 상태 설명

**주요 상태 정의:**

- **PLACED:** 주문이 플랫폼에 생성되었지만 판매자가 아직 처리하지 않음
- **PAYMENT_FAILED:** 결제 승인 또는 캡처에 실패함
- **ACCEPTED:** 판매자가 주문을 수락하고 준비를 시작함
- **IN_PREPARATION:** 주문이 배송 준비 중임
- **READY_TO_SHIP:** 모든 아이템이 포장되어 픽업/배송 준비 완료
- **SHIPPED:** 주문이 배송 업체에 인계됨
- **IN_TRANSIT:** 패키지가 고객에게 배송 중임
- **DELIVERED:** 패키지가 고객에게 배송됨
- **COMPLETED:** 주문이 완료되고 종료됨
- **CANCELED:** 주문이 취소됨

### 주문 상태 테이블

| Status | Description |
|--------|-------------|
| PLACED | 초기 주문 접수 |
| PAYMENT_FAILED | 결제 문제 발생 |
| ACCEPTED | 판매자가 주문을 수락함 |
| IN_PREPARATION | 준비 중 |
| READY_TO_SHIP | 배송 준비 완료 |
| SHIPPED | 배송 업체에 인계됨 |
| IN_TRANSIT | 배송 중 |
| DELIVERED | 고객에게 배송됨 |
| COMPLETED | 주문 완료 |
| CANCELED | 주문 취소 |

---

## 3. 패키지 처리 상태

### 패키지 상태 흐름

**두 개의 병렬 흐름:**

**왼쪽 (표준 흐름):**
```
PLACED
↓
READY_TO_SHIP
↓
SHIPPED
↓
IN_TRANSIT
↓
DELIVERED
↓
COMPLETED
```

**오른쪽 (대체 상태):**
```
READY_TO_SHIP
↓
PICKUP_REQUESTED
↓
PICKUP_FAILED
↓
RETURNED
↓
RETURN_IN_TRANSIT
↓
RETURN_DELIVERED
```

### 패키지 상태 설명

**상태 정의:**

- **PLACED:** 패키지가 생성되었지만 아직 준비되지 않음
- **READY_TO_SHIP:** 패키지가 준비되어 배송 업체 픽업 준비 완료
- **PICKUP_REQUESTED:** 배송 업체 픽업이 예약됨
- **SHIPPED:** 패키지가 배송 업체에 의해 픽업됨
- **IN_TRANSIT:** 패키지가 운송 중임
- **DELIVERED:** 패키지가 성공적으로 배송됨
- **PICKUP_FAILED:** 배송 업체가 패키지를 픽업하지 못함
- **RETURNED:** 패키지가 반송 중임
- **RETURN_IN_TRANSIT:** 반송 패키지가 운송 중임
- **RETURN_DELIVERED:** 반송 패키지가 판매자에게 다시 배송됨
- **COMPLETED:** 패키지 처리가 완료됨

---

## 4. 주문 목록 및 세부 정보 가져오기

**엔드포인트:**

- `GET /api/v1/orders` - 다양한 주문 상태의 주문 목록 가져오기
- `GET /api/v1/orders/{order_id}` - 주문 세부 정보 보기

---

## 5. 주문 취소

**엔드포인트:**

- `POST /api/v1/orders/{order_id}/cancel` - 주문을 취소하려면 이 엔드포인트를 사용하십시오.

**참고:** 예외 및 취소 요청을 처리해야 합니다.

---

## 6. 주문 분할

### 6.1 분할

**엔드포인트:** `POST /api/v1/orders/split`

주문에 여러 아이템이 포함된 경우, 주문 분할 기능을 사용하면 배송 준비 상태에 따라 각 아이템에 대한 배송을 개별적으로 조정할 수 있습니다.

**API 요청 예시:**

이 요청에서 주문은 4개의 아이템을 포함하며 두 개의 패키지로 나뉩니다.

### 코드 예제

```json
{
  "order_id": "ORDER123",
  "packages": [
    {
      "package_id": "PKG001",
      "item_list": [
        {
          "item_id": "ITEM001",
          "quantity": 2
        }
      ]
    },
    {
      "package_id": "PKG002", 
      "item_list": [
        {
          "item_id": "ITEM002",
          "quantity": 1
        },
        {
          "item_id": "ITEM003",
          "quantity": 1
        }
      ]
    }
  ]
}
```

### 추가 분할 예제

```json
{
  "item_id": "ITEM004",
  "quantity": 3,
  "order_item_id": "ORDERITEM001",
  "product_id": "PROD001"
}
```

```json
{
  "item_id": "ITEM005",
  "quantity": 2,
  "order_item_id": "ORDERITEM002",
  "product_id": "PROD002",
  "variant_list": []
}
```

```json
{
  "item_id": "ITEM006",
  "quantity": 1,
  "order_item_id": "ORDERITEM003",
  "product_id": "PROD003"
}
```

```json
{
  "item_id": "ITEM007",
  "quantity": 4,
  "order_item_id": "ORDERITEM004",
  "product_id": "PROD004"
}
```

```json
{
  "item_id": "ITEM008",
  "quantity": 1,
  "order_item_id": "ORDERITEM005",
  "product_id": "PROD005"
}
```

```json
{
  "item_id": "ITEM009",
  "quantity": 2,
  "order_item_id": "ORDERITEM006",
  "product_id": "PROD006"
}
```

---

**BACK** (오른쪽 하단 버튼)

---

# 배송 API 개발자 가이드

## 목차

```python
"rate_id": "123456789"
```

```python
"order_id": "123456789"
```

```python
"order_id_value": "123456789"
```

```python
"parcel_id_group_id": []
```

---

## 팁

1. 상점 프론트에서 주문 분할 권한을 설정하십시오. 주문을 분할할 때 "주문 분할 권한이 없습니다"라는 오류가 발생하면 판매자 센터로 이동하여 주문 분할 권한을 부여해야 합니다.

2. 동일한 번들 거래에 속하고 거래 프로모션 중인 아이템은 다른 패키지로 분할할 수 없습니다. 여러 아이템 프로모션의 경우 판매자는 위험 관리를 위해 매번 운영 팀에 연락하여 분할 정책을 문의해야 합니다.

3. 다중 주소로 주문을 분할한 후에는 동일한 분할 정책, 동일한 번들 거래 및 동일한 add_on_deal_id를 사용하여 분할 주문 요청을 받을 것으로 예상합니다. 이는 동일한 추가 기능 거래에 있음을 나타냅니다. 패키지는 동일한 거래 ID 그룹 일련 ID의 한 아이템에서 나올 수 있으며 이 분할 API를 호출할 때 다른 패키지 요청으로 분할할 수 없습니다.

4. 새 패키지를 업로드하고 패키지 목록에 배송 준비가 완료된 아이템과 배송 준비가 완료되지 않은 아이템이 모두 포함된 경우 (총) 거래에 3개의 주문 아이템이 있고 거래를 구성하면 주문을 두 개의 패키지로 분할할 수 있습니다. 동일한 거래 아이템이 동일한 패키지에 있고 거래 아이템만 분할하는 경우에만 이 API를 통해 분할할 수 있습니다. 배송 준비가 완료된 아이템이 있는 패키지는 READY_TO_SHIP 상태로 변경되고 다른 패키지는 UNPAID 또는 INIT 상태가 됩니다. 그러나 초기화된 상태의 패키지에 있는 아이템은 여전히 주문 시의 아이템으로 간주됩니다.

5. "분할됨"

6. [추가 참고] 주문 상태는 주문을 취소하기 전에 READY_TO_SHIP이어야 합니다. 소포가 INIT/UNPAID 상태인 경우 판매자는 배송을 준비해야 합니다.

---

## 7. 배송을 위한 패키지 목록 및 세부 정보 가져오기

[sv_order_package_list](링크) 및 [배송 API](링크)를 사용하여 필요에 따라 다양한 필터와 함께 주문 및 배송된 패키지를 가져옵니다. 이 API는 배송을 위한 패키지를 가져오는 데 선호됩니다.

**자세한 내용은** [sv_order_package_list](링크) **에서 확인하십시오.**

---

## 8. 배송 API 호출 흐름

[이것은 그림/차트입니다: 여러 결정 지점과 프로세스 단계를 포함하는 배송 API 호출 흐름을 보여주는 자세한 순서도입니다. 순서도에는 배송 프로세스의 다양한 단계를 나타내는 다양한 노란색 상자와 연결선 및 결정 다이아몬드가 포함되어 있습니다.]

---

## 8.1 기본 배송 로직

1. [sv_logistics_package_list API](링크)를 통해 `package_status of 1 (SHIPPED)`를 통해 배송될 패키지 목록을 가져옵니다.

   **참고:** `order_sn` & `package_number`는 패키지를 식별하는 데 사용됩니다.

2. 물류 요금이 결정되면 판매자는 패키지/prom_integrated 배송 방법을 사용하여 배송해야 합니다. [sv_logistics_init_info API](링크) 및 [sv_logistics_shipping_parameter API](링크)를 호출하여 배송 문서를 가져옵니다. 초기화 후 [sv_logistics_channel](링크)별 매개변수를 통해 패키지에는 픽업 주소와 예상 배송 비용이 나열됩니다. 물류 제공업체가 생성한 경우 호출자는 추적 번호를 가져올 수 있습니다. 그런 다음 API는 추적 번호를 반환하고 요청 본문에 업로드합니다. API 호출이 성공하면 픽업/요청 상태의 패키지 처리 상태가 업데이트됩니다. 또는 [sv_logistics_ship API](링크) 또는 [sv_logistics_confirm_order_status API](링크)를 사용하고 물류 제공업체에서 자동으로 생성하거나 판매자가 자체적으로 정의한 필요한 정보를 업로드합니다.

3. Shopee 통합 채널을 사용하여 성공적으로 배송한 후 [sv_logistics_get_tracking_number API](링크)를 호출하여 Shopee에서 생성된 추적 번호를 가져올 수 있습니다. 또는 [sv_logistics_get_tracking_info API](링크)를 호출하여 전체 추적 정보를 가져올 수 있습니다.

4. 추적 번호를 얻은 후 그룹 목록을 인쇄할 수 있습니다. "자체 인쇄 또는 Shopee 생성의 두 가지 방법을 선택할 수 있습니다." Shopee 인쇄 준비 API [sv_logistics_create_shipping_document API](링크)를 사용하면 추적 라벨, 배송 문서 및 배송 주문 양식을 [LOGISTICS_PICKUP_DONE](상태)로 만들 수 있습니다.

5. 배송된 패키지의 Shopee 항공 운송장 PDF를 얻으려면 시간에 따라 API가 얼마나 바쁜지 [sv_logistics_get_airway_bill API](링크)를 호출할 수 있습니다. [sv_logistics_download_shipping_document API](링크), [sv_logistics_download_shipping_document_result](링크)를 사용할 수도 있습니다.

6. [sv_logistics_get_shipping_parameter API](링크)를 호출하여 배송 매개변수를 가져올 때 slug 매개변수는 패키지를 기반으로 반환됩니다. [sv_logistics_shipping_document API](링크)를 호출하여 반환된 값은 API를 통해 이 패키지에 대한 것입니다. 비즈니스가 LOGISTICS(3rd)를 채우도록 하십시오. 픽업 목록을 인쇄할 필요가 없습니다. 3PL은 항공 운송장 목록을 제공하고 배송 프로세스를 완료합니다. [logistics_channel_id](param)의 판매자는 "slug=dropoff"만 있도록 slug=dropoff를 지정합니다. 패키지는 자체 인쇄할 준비가 됩니다.

---

## 8.2 주문 배송 관련 API

| API | Description |
|-----|-------------|
| [sv_order_search_package_list](링크) | 아직 배송되지 않은 패키지 목록 가져오기 |
| [sv_order_get_package_detail](링크) | 패키지 세부 정보 가져오기 |
| [sv_logistics_get_shipping_parameter](링크) | 배송 매개변수 가져오기 |
| [sv_logistics_get_order_list](링크) | 주문 목록 가져오기 |
| [sv_logistics_ship_order](링크) | 배송 준비 |
| [sv_logistics_get_address_list_parcel](링크) | - |
| [sv_logistics_get_tracking_number](링크) | 추적 번호 가져오기 |
| [sv_logistics_get_order_tracking_number](링크) | - |
| [sv_logistics_get_shipping_document_data_info](링크) | 청구서를 자체 인쇄하는 데 필요한 정보 가져오기 |
| [sv_logistics_get_shipping_document_parameter](링크) | 거래 및 권장 청구서 템플릿에 액세스 |
| [sv_logistics_create_shipping_document](링크) | 항공 운송장 작업 생성 |
| [sv_logistics_get_shipping_document_result](링크) | 항공 운송장 작업 결과 가져오기 |
| [sv_logistics_download_shipping_document](링크) | Shopee에서 생성된 항공 운송장 다운로드 |

---

## 8.3 API 요청 예제

### 8.3.1 [sv_logistics_get_shipping_parameter](링크)

요청 예제:

```json
{
  "v": "",
  "message": "",
  "request": {
    "info_needed": {
      "dropoff": [],
      "pickup": [],
      "order": []
    },
    "order_sn": "",
    "dropoff": "",
    "count": ""
  }
}
```

---

*참고: 이 추출은 요청된 대로 제목, 코드 블록, 테이블 및 서식을 유지하면서 개발자 가이드 스크린샷의 원래 영어 텍스트 구조를 유지합니다.*

---

# 개발자 가이드 - API 문서

## 설정

```javascript
"shipping": {
  "track_id": null
}
```

## 주문 구조

```javascript
"stage": {
  "address_id": "5206",
  "region": "KR",
  "state": "Seoul/Busan",
  "city": "Incheon/Busan",
  "district": "",
  "addr": "",
  "address": "22319",
  "zipcode": "03037",
  "address_line": {
    "default_address",
    "invoice_address",
    "refund_address"
  }
}
```

## 아이템 목록 구조

```javascript
"line_item_list": [
  {
    "sku": "10WSTCROS",
    "original_id": "10WSTCROS"
  },
  {
    "sku": "10W0000WL",
    "original_id": "10W0000WL"
  }
]
```

## 주소 정보

```javascript
"address_id": "5206",
"region": "KR",
"state": "Seoul/Busan",
  "city": "Incheon/Busan",
  "district": "",
  "addr": "",
  "address": "huga CSR",
  "zipcode": "03122",
  "address_line": []
```

## 거래 목록

```javascript
"line_item_list": [
  {
    "sku": "10WSTCROS",
    "original_id": "10WSTCROS"
  },
  {
    "sku": "10W0000WL",
    "original_id": "10W0000WL"
  }
]
```

## 결제 구성

```javascript
"payment_id": "GARMIN0913FUT-hp2t8n353e6f1807"
```

---

## 참고: ark_needed 필드는 주문에서 지원하는 배송 방법을 나타냅니다.

**참고:** ark_needed 필드는 주문에서 지원하는 배송 방법과 전달해야 하는 매개변수를 나타냅니다. 반환 값이 false이면 요청에서 지정된 아이템을 취소하거나 취소하거나 고객 매개변수를 조정할 때 매개변수를 전달할 필요가 없음을 의미합니다. 픽업을 선택한 경우 address_id 및 place_order_time 매개변수를 업로드해야 합니다. 지연된 처리 금액을 지원하는 주문이 있는 경우 X-Yijing delayed_fulfillment_amount를 사용하여 지연된 매개변수의 금액을 나타냅니다.

### 8.1.2 logistics_apis_global

**이 API 메서드 트리거:**
arfs_needed에서 반환된 ShopOff 매개변수가 logistics_arfs_shipping_parameter에 address_id 또는 pickup_arfs_id를 포함하는 경우

```javascript
"shipper": "SHIPPING_NUM"
```

### 사용법

```javascript
"stage": {
  "address_id": "5206",
  "original_id": "10WSTCROS"
}
```

---

## 2) ShopOff 방법 선택

arfs_needed에서 반환된 ShopOff 매개변수가 logistics_arfs_shipping_parameter에 비어 있는 경우

---

## 팁

```javascript
"shipper": "CONVENIENCE_STORE"
"goods": {
  "pickup_store": "ABC"
}
```

**참고:** 드롭 오프 방법에 대한 일부 채널에는 빈 필드의 상점 반환이 있으므로 예제와 같이 빈 필드를 전달해야 합니다. 다른 방법에 매개변수의 상점 반환이 있는 경우 예를 들어:

---

## 추가 구성

```javascript
"shipper": "CONVENIENCE_STORE"
"goods": {
  "pickup_store": "ABC"
}
```

**참고:** cvs_integrated 방법 선택 시
arfs_needed에서 반환된 cvs_integrated 매개변수가 logistics_arfs_shipping_parameter에 tracking_number에 있는 경우

---

## 최종 설정

```javascript
"shipper": "CONVENIENCE_STORE"
"goods": {
  "tracking_number": "ABCDEFGHIJKLM"
}
```

### 8.1.3 logistics_update_shipping_order에서

---

# API 문서 추출

## 8.3.3 w_logistics update_shipping_order

```json
{
  "order_sn": "2210123456789",
  "sla_timestamp": {
    "pickup": 1,
    "ship": 1672387266
  },
  "tracking_number": "ABC123456789"
}
```

---

## 배송 주문 업데이트

```json
{
  "order_sn": "2112345678901",
  "origin_id": {
    "package_id": "1116"
  },
  "tracking_info": "XXXXXXXX"
}
```

**사용법:** 주소, SLA 및 pickup_time 등을 업데이트하기 위해 픽업 주문에 사용됩니다. READY_TO_SHIP 상태의 주문에 적용됩니다.

---

## 9. FAQ

### 주문 관련

**Q:** `v2.order.get_order_detail` API를 호출하고 "잘못된 매개변수, 중단, 주문을 찾을 수 없습니다" 오류를 보고합니다.

**A:** 자세한 내용은 [문서](#)를 참조하십시오.

**Q:** `v2.order.cancel_order` API를 호출할 때 응답 필드가 누락된 경우 어떻게 해야 합니까?

**A:** 관련 정보를 해당 필드에 업로드해야 하는지 확인하십시오. 자세한 내용은 [문서](#)를 참조하십시오.

### 배송 관련

**Q:** 왜 enn sno를 얻을 수 없습니까?

**A:** shop_by_info에서 배송된 경우 enn을 얻을 수 없는 경우 전달되었습니다.

**Q:** enn "물류 상태가 배송 준비가 되지 않았습니다"를 얻었습니다. 어떻게 확인해야 합니까?

**A:** `v2.logistics.get_tracking_number` API를 호출하여 'READY_TO_SHIP' 상태의 주문을 물류에 업로드할 수 있습니다.

**Q:** `v2.logistics.ship_order` API를 호출했지만 fist_mile_tracking_number가 반환되지 않았습니다. 어떻게 해야 합니까?

**A:** shipplat에서 물류 수수료가 청구될 때 주문 내에서 패키지에 대한 추적 번호를 업로드하는 방법에 대한 자세한 내용은 [문서](#)를 참조하십시오.

### 항상 아픈 관련

**Q:** `v2.logistics.create_shipping_document` API를 호출하면 "주문 상태가 지원되지 않으며 오류가 발생했습니다"라는 오류가 표시됩니다.

**A:**

---

# API 문서 발췌

## 8.3.3 w_logistics update_shipping_order

```json
{
  "order_sn": "2210123456789",
  "sla_timestamp": {
    "pickup": 1,
    "ship": 1672387266
  },
  "tracking_number": "ABC123456789"
}
```

---

## 배송 주문 업데이트

```json
{
  "order_sn": "2112345678901",
  "origin_id": {
    "package_id": "1116"
  },
  "tracking_info": "XXXXXXXX"
}
```

**용도:** 주소, SLA 및 pickup_time 등을 업데이트하기 위해 픽업 주문에 사용됩니다. READY_TO_SHIP 상태의 주문에 적용됩니다.

---

## 9. FAQ

### 주문 관련

**Q:** `v2.order.get_order_detail` API를 호출하면 "잘못된 매개변수, 중단, 주문을 찾을 수 없습니다"라는 오류가 발생합니다.

**A:** 자세한 내용은 [문서](#)를 참조하십시오.

**Q:** `v2.order.cancel_order` API를 호출할 때 응답 필드가 누락되면 어떻게 해야 합니까?

**A:** 관련 정보를 해당 필드에 업로드해야 하는지 확인하십시오. 자세한 내용은 [문서](#)를 참조하십시오.

### 배송 관련

**Q:** 왜 enn sno를 얻을 수 없습니까?

**A:** shop_by_info가 전달되었을 때 배송된 경우 enn을 얻을 수 없습니다.

**Q:** "logistic status not ready to ship"이라는 enn을 받았습니다. 어떻게 확인해야 합니까?

**A:** `v2.logistics.get_tracking_number` API를 호출하여 'READY_TO_SHIP' 상태의 주문이 물류에 업로드될 수 있는지 확인하십시오.

**Q:** `v2.logistics.ship_order` API를 호출했지만 fist_mile_tracking_number가 반환되지 않으면 어떻게 해야 합니까?

**A:** shipplat에서 물류 수수료가 부과될 때 주문 내에서 패키지의 추적 번호를 업로드하는 방법에 대한 자세한 내용은 [문서](#)를 참조하십시오.

### 항상 아픈 관련

**Q:** `v2.logistics.create_shipping_document` API를 호출하면 "주문 상태가 지원되지 않으며 오류가 발생했습니다"라는 오류가 표시됩니다.

**A:** `v2.order.get_order_detail`을 호출하여 order_status 필드를 쿼리하십시오. 주문이 'READY_TO_SHIP' 상태인 경우.

**Q:** `v2.logistics.get_shipping_document_result` API를 호출했는데 "PROCESSING" 상태가 표시됩니다. 어떻게 처리해야 합니까?

**A:** "READY" 상태가 될 때까지 API 콜백을 기다리는 것이 좋습니다.

**Q:** 내 품목이 배송된 후 며칠 후?

**A:** 두 가지 형식이 있습니다.
- 단일 패키지 주문이 먼저 전송됩니다.
- 여러 패키지 (상점 모드에 따라 다름) 품목 형식의 시간. (SC 배송은 7-14일 (shippo id: 90000) 제외, 가족 배송: 90000), Life File (shipment_id: 90001), Family Frozen Support Pickup (수집 배송으로 전송되지 않음) - 90001 이상으로 업로드된 Family Chilled 파일은 모든 픽업 시간으로 차단되고 구매자는 알림을 받지 못함)

---

## 10. 데이터 정의

### 주문 상태

- **UNPAID:** 주문이 생성되었습니다. 구매자가 아직 지불하지 않았습니다.
- **READY_TO_SHIP:** 구매자가 지불했고, 결제 채널에서 결제가 완료되어 판매자에게 성공적으로 이체되었습니다.
- **PROCESSED:** 주문이 "READY_TO_SHIP"이지만 SPL에서 픽업 번호를 받았습니다.
- **SHIPPED:** 주문이 판매자에 의해 배송되었습니다.
- **TO_CONFIRM_RECEIVE:** 구매자가 주문을 받았습니다.
- **COMPLETED:** 주문이 완료되었습니다.
- **IN_CANCEL:** 구매자 또는 판매자가 주문을 취소하고 있습니다.
- **CANCELLED:** 주문 취소가 처리 중입니다.
- **CANCELLED:** 주문이 취소되었습니다.
- **INVOICE_PENDING:** 주문이 송장 반환을 기다리고 있습니다.

### 패키지 상태

- **Pending:** 보류 중인 패키지가 완료되지 않았습니다. 3pl/물류 또는 처리됨, 상태 0
- **Pending:** 배송 준비가 되지 않은 패키지를 가져옵니다, 상태 1
- **Processed:** 배송을 준비해야 하는 패키지를 가져옵니다, 상태 2
- **Processed:** 픽업 또는 드롭오프를 준비하는 패키지를 가져옵니다, 상태 3

### 판매자 센터의 배송 목록에서 동일한 사무실의 주문 상태에 대한 패키지 상태 테이블 (주문 아님): 패키지 상태 + 패키지 배송 상태와 판매자 센터 주문 상태 간의 매핑 (테이블에 없음):

| 패키지 상태 (판매자) | 주문 상태 (판매자 센터 - 판매자 센터의 배송 목록) |
|------------------------|-------------------------------------------------------------|
| All [0] | LOGISTICS_NOT_START, LOGISTICS_READY, LOGISTICS_PICKUP_DONE, LOGISTICS_PICKUP_RETRY_1, LOGISTICS_PICKUP_RETRY_2 | All |
| Pending [1] | LOGISTICS_NOT_START | Pending |
| 3pl/logistic [2] | LOGISTICS_READY or LOGISTICS_PICKUP_RETRY | To Process |
| Processed [3] | LOGISTICS_REQUEST_CREATED | Processed |

---

### 패키지 처리 상태 / 물류 상태

- **LOGISTICS_NOT_START:** 판매자가 배송을 준비하지 않았습니다. 패키지가 처리되지 않았습니다.
- **LOGISTICS_READY:** 결제 관점에서 패키지가 처리될 준비가 되었습니다. 비 COD의 경우 과거: biz/COD: COD 심사를 통과했습니다.
- **LOGISTICS_REQUEST_CREATED:** 패키지가 성공적으로 업로드되었습니다.
- **LOGISTICS_INVALID:** 패키지 업로드 오류 또는 취소 오류입니다.
- **LOGISTICS_ARRANGE_PICKUP_DONE:** 패키지가 성공적으로 할당되었습니다.
- **LOGISTICS_ARRANGE_PICKUP_FAILED:** 패키지 픽업 준비에 실패했습니다.
- **LOGISTICS_PICKUP_DONE:** 패키지가 픽업되었습니다. COD 주문 비통합의 경우 LOGISTICS_REQUEST_CREATED를 통과한 경우
- **LOGISTICS_PICKUP_FAILED:** 레이블 픽업으로 인해 SPL에서 주문이 취소되었거나 픽업되었지만 배송을 진행할 수 없습니다.
- **LOGISTICS_PICKUP_RETRY:** 패키지가 3PL 재시도 픽업을 기다리고 있습니다.
- **LOGISTICS_PICKUP_RETRY_2:** 패키지가 여전히 3PL 재시도 픽업을 기다리고 있으며, 주문 처리에 2번의 추가 재시도가 실패했습니다.
- **LOGISTICS_PENDING_ARRANGE:** 판매자 물류 제공 업체 준비
- **LOGISTICS_COD_REJECTED:** 통합 물류 COD, COD에 의해 주문이 거부되었습니다.

---

### 주문 취소 사유

- **OUT_OF_STOCK**
- **UNDELIVERABLE_AREA**

---

### 취소 사유

- 품절
- 구매자의 취소 요청
- 고객 요청
- COD 사용 불가
- 처리 실패
- 잘못된 가격 업데이트
- 유효하지 않은 주문
- 서비스되지 않는 코드
- 물류 방법 결제
- 물류 요청이 취소되었습니다.
- SPL 거부됨
- SPL 취소됨
- 제품을 사용할 수 없음
- 판매자가 배송하지 않음
- 운송 창고 취소됨
- 구매자 취소됨
- 비활성 판매자
- 판매자가 배송할 수 없음
- 판매자 취소됨
- 판매자 취소
- 귀하의 승인이 승인되지 않았습니다 (온라인 포함)
- 다른 SPL로의 주문이지만 현재 주문 SN

---

### 구매자 취소 사유

- 판매자가 구매자의 문의에 응답하지 않습니다.
- 판매자가 구매자에게 취소를 요청합니다.
- 품목을 변경하고 싶습니다.
- 제품에 나쁜 리뷰가 있습니다.
- 판매자가 주문을 배송할 수 없습니다.
- 주소를 주문하고 싶습니다.
- 기타
- 운임이 바우처 코드를 초과합니다.
- 전화 번호를 변경해야 합니다.
- 배송 주소를 변경해야 합니다.
- 바우처 코드를 구매 / 변경해야 합니다.
- 결제 방법을 변경해야 합니다.
- 사기 / 스캠
- 중복 주문
- 다른 곳에서 사기 청구
- 주문의 특정
- 주소 / 픽업 위치 변경
- 특정 품목을 구매하고 싶지 않습니다 (제거됨).
- 결제 주소를 변경해야 합니다.
- 과도한
- 세일에서 더 저렴한 품목을 구매하고 싶습니다, XXXXXX_ETC.
- 마음의 변화 / 기타

---

# 주문 취소 사유

## 배송 주소를 변경해야 합니다.
- 배송 주소를 변경해야 합니다.
- 바우처 코드를 입력 / 변경해야 합니다.
- 주문을 수정해야 합니다.
- 결제 절차가 너무 번거롭습니다.
- 다른 곳에서 더 저렴하게 찾았습니다.
- 더 이상 구매하고 싶지 않습니다.
- 귀하의 승인자가 주문을 거부했습니다.
- 현재 주문을 할 수 없습니다.
- 배송 주소를 변경해야 합니다.
- 배송 시간이 너무 깁니다.
- 기존 주문 수정 (색상, 크기, 바우처 등)
- 마음의 변화 / 기타

# 배송 문서 유형

- NORMAL_AIR_WAYBILL
- THERMAL_AIR_WAYBILL
- NORMAL_JOB_AIR_WAYBILL
- THERMAL_JOB_AIR_WAYBILL

# 패키지 물류 추적 상태

(`get_tracking_info` API용)

- INITIAL
- ORDER_INIT
- ORDER_SUBMITTED
- ORDER_FINALIZED
- ORDER_CREATED
- PICKUP_REQUESTED
- PICKUP_PENDING
- PICKED_UP
- DELIVERY_PENDING
- DELIVERED
- PICKUP_RETRY
- TIMEOUT
- LOST
- UPDATE
- UPDATE_SUBMITTED
- UPDATE_CREATED
- RETURN_STARTED
- RETURNED
- RETURN_PENDING
- RETURN_INITIATED
- EXPIRED
- CANCEL
- CANCEL_CREATED
- CANCELED
- FAILED_ORDER_INIT
- FAILED_ORDER_SUBMITTED
- FAILED_ORDER_CREATED
- FAILED_PICKUP_REQUESTED
- FAILED_PICKED_UP
- FAILED_DELIVERED
- FAILED_UPDATE_SUBMITTED
- FAILED_UPDATE_CREATED
- FAILED_RETURN_STARTED
- FAILED_RETURNED
- FAILED_CANCEL_CREATED
- FAILED_CANCELED

## 사용 사례

1. 주문 생성부터 완료까지 주문 라이프사이클 관리
2. 주문 및 패키지 상태 추적
3. 효율적인 배송을 위한 주문 분할
4. 배송 API와 통합
5. 배송 프로세스 자동화

## 관련 API

- GET /api/v1/orders
- GET /api/v1/orders/{order_id}
- POST /api/v1/orders/{order_id}/cancel
- POST /api/v1/orders/split
- sv_order_package_list
- sv_logistics_package_list API
- sv_logistics_init_info API
- sv_logistics_shipping_parameter API
- sv_logistics_channel
- sv_logistics_ship API
- sv_logistics_confirm_order_s

---

## 원문 (English)

### Summary

This developer guide provides comprehensive information on order management, including order and package status flows, API endpoints for order and package retrieval, order cancellation, and order splitting. It also covers shipment API call flows and basic shipping logic.

### Content

# Order Management

## 1. Introduction

API Best Practice > Order Management

**Item:** The individual products within an order. Unit quantity and other details. Items are included in packages for shipment.

---

## 2. Order Status Flow

### Order Status Workflow

**Flowchart showing order progression:**

```
PLACED
↓
PAYMENT_FAILED
↓
ACCEPTED → IN_PREPARATION
↓
READY_TO_SHIP
↓
SHIPPED
↓
IN_TRANSIT
↓
DELIVERED
↓
COMPLETED
```

**Side branches:**
- CANCELED (can occur at multiple stages)
- READY_TO_SHIP can also lead to CANCELED

### Status Descriptions

**Key status definitions:**

- **PLACED:** Order has been created in the platform but not yet processed by the seller
- **PAYMENT_FAILED:** Payment authorization or capture has failed
- **ACCEPTED:** Seller has accepted the order and will begin preparation
- **IN_PREPARATION:** Order is being prepared for shipment
- **READY_TO_SHIP:** All items are packed and ready for pickup/shipment
- **SHIPPED:** Order has been handed over to the shipping carrier
- **IN_TRANSIT:** Package is in transit to the customer
- **DELIVERED:** Package has been delivered to the customer
- **COMPLETED:** Order is complete and closed
- **CANCELED:** Order has been canceled

### Order Status Table

| Status | Description |
|--------|-------------|
| PLACED | Initial order placement |
| PAYMENT_FAILED | Payment issue occurred |
| ACCEPTED | Order accepted by seller |
| IN_PREPARATION | Being prepared |
| READY_TO_SHIP | Ready for shipment |
| SHIPPED | Handed to carrier |
| IN_TRANSIT | In delivery transit |
| DELIVERED | Delivered to customer |
| COMPLETED | Order finalized |
| CANCELED | Order canceled |

---

## 3. Package Fulfillment Status

### Package Status Flow

**Two parallel flows:**

**Left side (Standard flow):**
```
PLACED
↓
READY_TO_SHIP
↓
SHIPPED
↓
IN_TRANSIT
↓
DELIVERED
↓
COMPLETED
```

**Right side (Alternative states):**
```
READY_TO_SHIP
↓
PICKUP_REQUESTED
↓
PICKUP_FAILED
↓
RETURNED
↓
RETURN_IN_TRANSIT
↓
RETURN_DELIVERED
```

### Package Status Descriptions

**Status definitions:**

- **PLACED:** Package created but not yet ready
- **READY_TO_SHIP:** Package is prepared and ready for carrier pickup
- **PICKUP_REQUESTED:** Carrier pickup has been scheduled
- **SHIPPED:** Package picked up by carrier
- **IN_TRANSIT:** Package is being transported
- **DELIVERED:** Package delivered successfully
- **PICKUP_FAILED:** Carrier failed to pick up package
- **RETURNED:** Package is being returned
- **RETURN_IN_TRANSIT:** Return package in transit
- **RETURN_DELIVERED:** Return package delivered back to seller
- **COMPLETED:** Package fulfillment complete

---

## 4. Getting order list and details

**Endpoints:**

- `GET /api/v1/orders` - Get the list of orders with different order status
- `GET /api/v1/orders/{order_id}` - View order details

---

## 5. Canceling Order

**Endpoint:**

- `POST /api/v1/orders/{order_id}/cancel` - Use this to cancel orders

**Note:** Need to handle Exceptions and cancellation requests.

---

## 6. Splitting Order

### 6.1 Splitting

**Endpoint:** `POST /api/v1/orders/split`

When an order contains multiple items, the split order function can help you arrange shipping for each item separately according to its readiness to be delivered.

**API Request example:**

In this request, the order contains 4 items and is divided into two packages.

### Code Example

```json
{
  "order_id": "ORDER123",
  "packages": [
    {
      "package_id": "PKG001",
      "item_list": [
        {
          "item_id": "ITEM001",
          "quantity": 2
        }
      ]
    },
    {
      "package_id": "PKG002", 
      "item_list": [
        {
          "item_id": "ITEM002",
          "quantity": 1
        },
        {
          "item_id": "ITEM003",
          "quantity": 1
        }
      ]
    }
  ]
}
```

### Additional Split Examples

```json
{
  "item_id": "ITEM004",
  "quantity": 3,
  "order_item_id": "ORDERITEM001",
  "product_id": "PROD001"
}
```

```json
{
  "item_id": "ITEM005",
  "quantity": 2,
  "order_item_id": "ORDERITEM002",
  "product_id": "PROD002",
  "variant_list": []
}
```

```json
{
  "item_id": "ITEM006",
  "quantity": 1,
  "order_item_id": "ORDERITEM003",
  "product_id": "PROD003"
}
```

```json
{
  "item_id": "ITEM007",
  "quantity": 4,
  "order_item_id": "ORDERITEM004",
  "product_id": "PROD004"
}
```

```json
{
  "item_id": "ITEM008",
  "quantity": 1,
  "order_item_id": "ORDERITEM005",
  "product_id": "PROD005"
}
```

```json
{
  "item_id": "ITEM009",
  "quantity": 2,
  "order_item_id": "ORDERITEM006",
  "product_id": "PROD006"
}
```

---

**BACK** (button at bottom right)

---

# Shipment API Developer Guide

## Table of Contents

```python
"rate_id": "123456789"
```

```python
"order_id": "123456789"
```

```python
"order_id_value": "123456789"
```

```python
"parcel_id_group_id": []
```

---

## Tips

1. Split order permission by the shop front. If you get the error "You don't have the permission to split order" when splitting orders, you should go to Seller Center to grant the split order permission.

2. Items under the same bundle deal and are on deal promotion cannot be split into different packages. Only for multiple item promotions, sellers have to contact operation team every time to ask for split policy for risk control.

3. After splitting an order with multi-address, we expect to receive split order requests with the same split policy same bundle deal and the add_on_deal_id is the same, indicating that they are in the same add-on deal. Packages may come from one item of the same deal id group serial id, you can not split them into different package request when calling this split api.

4. When you upload a new package and the package list contains items that are ready to ship, and some items are not ready to ship instance, if 3 order items in a (total) deal and a configure a deal, the order can be split into two packages. You can only split through this api when the same deal items are in the same package and you only split deal items. The package with ready to ship item(s) will change to READY_TO_SHIP status and others will be UNPAID or INIT status. However, items in the package with Initialized status are still considered as the items at the order.

5. Once an "Spiltted"

6. [Additional note] The order status is READY_TO_SHIP before the order can be canceled splitting. If any parcel has been INIT/UNPAID, sellers must arrange shipment.

---

## 7. Getting package list and details for shipment

Use [sv_order_package_list](link) and [Shipment API](link) to fetch the order and its package shipped, with various filters as per your need. The api is preferred to fetch packages for shipment.

**View more at** [sv_order_package_list](link) **for more info.**

---

## 8. Shipment API Call Flow

[THIS IS FIGURE/CHART: A detailed flowchart showing the Shipment API call flow with multiple decision points and process steps. The flowchart includes various yellow boxes representing different stages of the shipment process, with connecting lines and decision diamonds.]

---

## 8.1 Basic Shipping Logic

1. Get the list of packages to be shipped via `package_status of 1 (SHIPPED)` through [sv_logistics_package_list API](link)

   **Note:** `order_sn` & `package_number` are used to identify a package.

2. Logistics charges are determined, seller should also see of package/prom_integrated shipping method to ship. Call [sv_logistics_init_info API](link) and [sv_logistics_shipping_parameter API](link) to get the shipping document. After init, parameter by [sv_logistics_channel](link), the package lists the pickup address and estimated shipping cost. The caller can get the tracking number, if the logistics provider has generated that. Then the API will return the tracking number and upload it to the request body. After the API call is successful, the package fulfillment status of pickup / request status will be updated. Alternatively [sv_logistics_ship API](link) or [sv_logistics_confirm_order_status API](link) and upload the required info automatically generated by logistics provider or self-defined by sellers.

3. After successful shipment using the Shopee integration channel, you can call [sv_logistics_get_tracking_number API](link) to get the Shopee generated tracking number. Alternatively you can call [sv_logistics_get_tracking_info API](link) to fetch the full tracking information.

4. After getting the tracking number, you can print the group list "You can choose two ways, self-print or Shopee generated." With Shopee print-ready API use [sv_logistics_create_shipping_document API](link), you will have a tracking label, shipping document, and shipping order form to [LOGISTICS_PICKUP_DONE](status)

5. To get the Shopee airway bill PDF of the shipped package, you can call on how busy APIs are by time [sv_logistics_get_airway_bill API](link). You can also use [sv_logistics_download_shipping_document API](link), [sv_logistics_download_shipping_document_result](link).

6. When calling the [sv_logistics_get_shipping_parameter API](link) to get the shipping parameters, the slug parameter is returned based on the package. The returned values calling the [sv_logistics_shipping_document API](link) is for this package via API. Let the business fill the LOGISTICS (3rd), there is no need to print pickup list. 3PL will provide the airway list and complete the shipping process. Seller in [logistics_channel_id](param), specify slug=dropoff so that only the "slug=dropoff" The package will be ready to self-print out.

---

## 8.2 Related APIs for order shipment

| API | Description |
|-----|-------------|
| [sv_order_search_package_list](link) | Get a list of packages which not shipped yet |
| [sv_order_get_package_detail](link) | Get package details |
| [sv_logistics_get_shipping_parameter](link) | Get shipping parameters |
| [sv_logistics_get_order_list](link) | Get order list |
| [sv_logistics_ship_order](link) | Arrange shipment |
| [sv_logistics_get_address_list_parcel](link) | - |
| [sv_logistics_get_tracking_number](link) | Get tracking number |
| [sv_logistics_get_order_tracking_number](link) | - |
| [sv_logistics_get_shipping_document_data_info](link) | Get the information you need to self-print your bill |
| [sv_logistics_get_shipping_document_parameter](link) | Access to transaction and recommended bill template |
| [sv_logistics_create_shipping_document](link) | Create airway bill task |
| [sv_logistics_get_shipping_document_result](link) | Get airway bill task result |
| [sv_logistics_download_shipping_document](link) | Download Shopee generated airway bill |

---

## 8.3 API request example

### 8.3.1 [sv_logistics_get_shipping_parameter](link)

Request example:

```json
{
  "v": "",
  "message": "",
  "request": {
    "info_needed": {
      "dropoff": [],
      "pickup": [],
      "order": []
    },
    "order_sn": "",
    "dropoff": "",
    "count": ""
  }
}
```

---

*Note: This extraction maintains the original English text structure from the developer guide screenshot, preserving headings, code blocks, tables, and formatting as requested.*

---

# Developer Guide - API Documentation

## Setup

```javascript
"shipping": {
  "track_id": null
}
```

## Order Structure

```javascript
"stage": {
  "address_id": "5206",
  "region": "KR",
  "state": "Seoul/Busan",
  "city": "Incheon/Busan",
  "district": "",
  "addr": "",
  "address": "22319",
  "zipcode": "03037",
  "address_line": {
    "default_address",
    "invoice_address",
    "refund_address"
  }
}
```

## Item List Structure

```javascript
"line_item_list": [
  {
    "sku": "10WSTCROS",
    "original_id": "10WSTCROS"
  },
  {
    "sku": "10W0000WL",
    "original_id": "10W0000WL"
  }
]
```

## Address Information

```javascript
"address_id": "5206",
"region": "KR",
"state": "Seoul/Busan",
"city": "Incheon/Busan",
"district": "",
"addr": "",
"address": "huga CSR",
"zipcode": "03122",
"address_line": []
```

## Transaction List

```javascript
"line_item_list": [
  {
    "sku": "10WSTCROS",
    "original_id": "10WSTCROS"
  },
  {
    "sku": "10W0000WL",
    "original_id": "10W0000WL"
  }
]
```

## Payment Configuration

```javascript
"payment_id": "GARMIN0913FUT-hp2t8n353e6f1807"
```

---

## Note: The ark_needed field indicates the shipping method supported by the order

**Note:** The ark_needed field indicates the shipping method supported by the order and the parameters this need to be passed. When the return value is false, it means that it is not necessary to pass the parameter when cancel or revoke a specified item in the request or adjust customer parameters. If picking is selected, you need to upload address_id and place_order_time parameters. If you have the order that supports delayed fulfillment amount, X-Yijing the delayed_fulfillment_amount to indicate the amount of delayed parameter.

### 8.1.2 logistics_apis_global

**Triggering this API method:**
When the ShopOff parameter returned by arfs_needed is in logistics_arfs_shipping_parameter contains address_id or pickup_arfs_id

```javascript
"shipper": "SHIPPING_NUM"
```

### Usage

```javascript
"stage": {
  "address_id": "5206",
  "original_id": "10WSTCROS"
}
```

---

## 2) Selecting the ShopOff method

When the ShopOff parameter returned by arfs_needed is in logistics_arfs_shipping_parameter is empty

---

## Tips

```javascript
"shipper": "CONVENIENCE_STORE"
"goods": {
  "pickup_store": "ABC"
}
```

**Note:** Some channels for drop off methods have a store return of empty fields, you need to pass in the empty field, such as the example. If other methods have a store return of parameters, for example:

---

## Additional Configuration

```javascript
"shipper": "CONVENIENCE_STORE"
"goods": {
  "pickup_store": "ABC"
}
```

**Note:** In Selecting the cvs_integrated method
When the cvs_integrated parameter returned by arfs_needed is in logistics_arfs_shipping_parameter in tracking_number

---

## Final Setup

```javascript
"shipper": "CONVENIENCE_STORE"
"goods": {
  "tracking_number": "ABCDEFGHIJKLM"
}
```

### 8.1.3 in logistics_update_shipping_order

---

# API Documentation Extract

## 8.3.3 w_logistics update_shipping_order

```json
{
  "order_sn": "2210123456789",
  "sla_timestamp": {
    "pickup": 1,
    "ship": 1672387266
  },
  "tracking_number": "ABC123456789"
}
```

---

## Update shipping order

```json
{
  "order_sn": "2112345678901",
  "origin_id": {
    "package_id": "1116"
  },
  "tracking_info": "XXXXXXXX"
}
```

**Usage:** Used for pickup order to update address, sla and pickup_time etc. Applicable to orders w/ READY_TO_SHIP status.

---

## 9. FAQ

### Order related

**Q:** Call `v2.order.get_order_detail` API and report error "Wrong parameters, abort, the order is not found"

**A:** For details, refer to the [documentation](#) for details.

**Q:** Call `v2.order.cancel_order` API, when response fields are missing, what should I do?

**A:** Please check whether the related information needs to upload the corresponding field, please refer to the [documentation](#) for details.

### Shipping related

**Q:** Why I can't get the enn sno?

**A:** If you can't get the enn's when there been shipped at shop_by_info has passed.

**Q:** I got the enn "logistic status not ready to ship" How to check?

**A:** Please call `v2.logistics.get_tracking_number` API to get the orders with 'READY_TO_SHIP' status can be uploaded for logistics.

**Q:** Call `v2.logistics.ship_order` API but no fist_mile_tracking_number is returned, what do I need to do?

**A:** For details on upload tracking number for a package within the order when logistics fee's charged by shipplat, refer to the [documentation](#) for details.

### Always ill related

**Q:** Call `v2.logistics.create_shipping_document` API, prompt error "Order status does not support and error" 

**A:** Please call `v2.order.get_order_detail` Query order_status field. When order is 'READY_TO_SHIP' status.

**Q:** Call `v2.logistics.get_shipping_document_result` API and I get status of "PROCESSING" How to deal with it?

**A:** It is recommended to wait for API callback until you get to the "READY" status.

**Q:** After several days after my item(s) are ship?

**A:** There are two formats:
- Single package order will be sent first
- Multiple packages (Depending on shop mode) time in item format. (SC shipments except 7-14 DAYS (shippo id: 90000), Family shipment: 90000), Life File (shipment_id: 90001), Family Frozen Support Pickup (not sent to colleting shipment) - Family Chilled file uploaded to 90001 and above will be blocked by all pickup times and buyer will receive no notification)

---

## 10. Data Definition

### Order Status

- **UNPAID:** Order is created. Buyer has not paid yet
- **READY_TO_SHIP:** Buyer has paid, payment is done by payment channel and successfully transferred to the seller
- **PROCESSED:** Orders are "READY_TO_SHIP" but get picking number from SPL
- **SHIPPED:** Order been shipped out by seller
- **TO_CONFIRM_RECEIVE:** The order has been received by buyer
- **COMPLETED:** The order has been completed
- **IN_CANCEL:** The buyer or seller is canceling the order
- **CANCELLED:** The order's cancellation is under processing
- **CANCELLED:** The order has been cancelled
- **INVOICE_PENDING:** The order is waiting for the invoice return is processing

### Package Status

- **Pending:** Pending package has not finished. 3pl/logistic or Processed, state 0
- **Pending:** Fetch package that are not ready for shipment, state 1
- **Processed:** Fetch's package that need to arrange shipment, state 2
- **Processed:** Fetch package that arrange pick-up or dropping off, state 3

### Package Status table for the same office in Order Status (Not order) To Ship List in Seller Center, the mapping between Package Status + Package's Shipment Status and Seller Center Order Status (Not in a table):

| Package Status (Seller) | Order Status (Seller Center - To Ship List in Seller Center) |
|------------------------|-------------------------------------------------------------|
| All [0] | LOGISTICS_NOT_START, LOGISTICS_READY, LOGISTICS_PICKUP_DONE, LOGISTICS_PICKUP_RETRY_1, LOGISTICS_PICKUP_RETRY_2 | All |
| Pending [1] | LOGISTICS_NOT_START | Pending |
| 3pl/logistic [2] | LOGISTICS_READY or LOGISTICS_PICKUP_RETRY | To Process |
| Processed [3] | LOGISTICS_REQUEST_CREATED | Processed |

---

### Package Fulfillment Status / Logistics Status

- **LOGISTICS_NOT_START:** seller hasn't arranged shipment. Package not fulfill
- **LOGISTICS_READY:** Package ready for fulfillment from payment perspective. For non-COD, past: biz/COD: passed COD screening
- **LOGISTICS_REQUEST_CREATED:** Package uploaded successfully
- **LOGISTICS_INVALID:** Error on Package upload error or cancellation
- **LOGISTICS_ARRANGE_PICKUP_DONE:** Package successfully allocated
- **LOGISTICS_ARRANGE_PICKUP_FAILED:** Package arrange pickup failed
- **LOGISTICS_PICKUP_DONE:** Package picked-up. For COD order non-integrated, if passed LOGISTICS_REQUEST_CREATED
- **LOGISTICS_PICKUP_FAILED:** Order cancelled by SPL due to label pickup or picked-up but not able to proceed with shipment
- **LOGISTICS_PICKUP_RETRY:** Package pending 3PL retry pickup
- **LOGISTICS_PICKUP_RETRY_2:** Package still pending 3PL retry pickup, order failed fulfillment 2 additional retries
- **LOGISTICS_PENDING_ARRANGE:** seller logistics provider arrangement
- **LOGISTICS_COD_REJECTED:** Integrated logistics COD, Order rejected by COD

---

### Order cancellation reason

- **OUT_OF_STOCK**
- **UNDELIVERABLE_AREA**

---

### Cancel reason

- Out of Stock
- Buyer Request to Cancel
- Customer Request
- COD Unavailable
- Failed to Fulfill
- Wrong Price Update
- Invalid Order
- Unserved Code
- Logistics Method Payment
- Logistics Request is Cancelled
- SPL rejected
- SPL cancelled
- Product not available
- Seller did not ship
- Transit Warehouse Cancelled
- Buyer cancelled
- Inactive Seller
- Seller can not ship
- Seller cancelled
- Seller cancel
- Your approval did not approve (with on-line)
- To other SPL but order SN of the moment

---

### Buyer cancel reason

- Seller is not Response to buyer's inquiries
- Seller ask buyer to cancel
- Want to change the item
- Product has Bad Reviews
- Seller is not Able to Ship The Order
- I want to order the address
- Others
- Freight exceed Voucher Code
- Need to change phone Number
- Need to change delivery address
- Need to buy / Change Voucher Code
- Need to Change Payment Method
- Frauds / Scams
- Duplicated order
- Fraud Charges Elsewhere
- Specific of order
- Change address / pickup location
- Do not want to buy certain item(s) (Removed)
- Need to change payment address
- Excessive
- Want to buy cheaper items from Sale, XXXXXX_ETC.
- Change of mind / others

---

# Order Cancellation Reasons

## Need to Change Delivery Address
- Need to Change Delivery Address
- Need to input / Change Voucher Code
- Need to Modify Order
- Payment Procedure too Troublesome
- Found Cheaper Elsewhere
- Don't Want to Buy Anymore
- Your approver rejected the order
- You are unable to place order at the moment
- Need to change delivery address
- Too long delivery time
- Modify existing order (color, size, voucher, etc)
- Change of mind / others

# Shipping document type

- NORMAL_AIR_WAYBILL
- THERMAL_AIR_WAYBILL
- NORMAL_JOB_AIR_WAYBILL
- THERMAL_JOB_AIR_WAYBILL

# Package logistics track status

(for `get_tracking_info` api)

- INITIAL
- ORDER_INIT
- ORDER_SUBMITTED
- ORDER_FINALIZED
- ORDER_CREATED
- PICKUP_REQUESTED
- PICKUP_PENDING
- PICKED_UP
- DELIVERY_PENDING
- DELIVERED
- PICKUP_RETRY
- TIMEOUT
- LOST
- UPDATE
- UPDATE_SUBMITTED
- UPDATE_CREATED
- RETURN_STARTED
- RETURNED
- RETURN_PENDING
- RETURN_INITIATED
- EXPIRED
- CANCEL
- CANCEL_CREATED
- CANCELED
- FAILED_ORDER_INIT
- FAILED_ORDER_SUBMITTED
- FAILED_ORDER_CREATED
- FAILED_PICKUP_REQUESTED
- FAILED_PICKED_UP
- FAILED_DELIVERED
- FAILED_UPDATE_SUBMITTED
- FAILED_UPDATE_CREATED
- FAILED_RETURN_STARTED
- FAILED_RETURNED
- FAILED_CANCEL_CREATED
- FAILED_CANCELED

---

**문서 ID**: developer-guide.229
**플랫폼**: shopee
**URL**: https://open.shopee.com/developer-guide/229
**처리 완료**: 2025-10-16T08:52:03
