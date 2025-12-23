# API 모범 사례 > First Mile 바인딩

**카테고리**: 모범 사례
**난이도**: 중간
**중요도**: 4/5
**최종 업데이트**: 2025-10-16T08:53:15

## 개요

본 가이드는 Shopee API에서 first-mile 바인딩을 위한 모범 사례를 설명합니다. 픽업, 드롭오프, 자체 배송과 같은 다양한 배송 방법을 다루며, 각 방법에 필요한 단계와 API 호출을 자세히 설명합니다. 또한 자주 묻는 질문에 대한 답변과 추적 번호 상태에 대한 데이터 정의를 제공합니다.

## 주요 키워드

- first mile
- binding
- pick up
- drop off
- self deliver
- tracking number
- shipping methods
- API
- order status

## 본문

# API 모범 사례 > First Mile 바인딩

## 용어

**Pick up (수거)**: 채널의 문전 수거, "logistics_channel_name": "shopeef" ("logistics_channel_id": 81s)만 수거를 지원합니다.

**Drop off (위탁)**: 채널 대리점에 배송, "logistics_channel_name": "shopeef" ("logistics_channel_id": 81s) 및 "logistics_channel_name": "Self Deliver" ("logistics_channel_id": 0)을 예상합니다. 다른 채널은 위탁입니다.

**Self deliver (자체 배송)**: 자체 배송, 타사 물류를 사용하지 않음, "logistics_channel_name": "Self Deliver"

**first_mile_tracking_number**: first-mile 추적 번호

## 모범 사례

다양한 유형의 first-mile 배송 방법은 다음과 같이 분류할 수 있습니다.
- Pick up (수거)
- Drop off (위탁)
- Self deliver (자체 배송)

개발자는 first_mile get_channel_list API를 통해 first-mile 채널 목록과 해당 채널에서 지원하는 배송 방법을 얻을 수 있습니다.

### Pick up 모드의 경우:

1단계: v2.first_mile.generate_first_mile_tracking_number API를 통해 first mile 추적 번호를 생성하고 v2.first_mile.get_tracking_number_list API를 사용하여 생성된 first mile 추적 번호 목록을 쿼리할 수 있습니다.

2단계: v2.first_mile.get_unibind_order_list API를 사용하여 first-mile 추적 번호와 바인딩할 주문 목록을 쿼리할 수 있으며, v2.first_mile.bind_first_mile_tracking_number API를 통해 바인딩하십시오. v2.first_mile.get_detail API를 사용하여 first-mile 추적 번호 세부 정보를 쿼리할 수 있습니다.

3단계: v2.first_mile.get_waybill API를 호출하여 first mile 패키지 라벨을 인쇄합니다.

### Drop off 모드의 경우:

1단계: 판매자는 채널 추적 번호를 오프라인으로 획득합니다.

2단계: v2.first_mile.get_unbind_order_list API를 통해 first-mile 추적 번호와 바인딩할 주문 목록을 가져오고 v2.first_mile.bind_first_mile_tracking_number API를 통해 바인딩합니다. v2.first_mile.get_detail API를 사용하여 first-mile 추적 번호 세부 정보를 쿼리할 수 있습니다.

### Self deliver 모드의 경우:

1단계: 판매자는 v2.first_mile.get_unbind_order_list API를 통해 first-mile 추적 번호와 바인딩할 주문 목록을 가져옵니다.

2단계: v2.first_mile.generate_first_mile_tracking_number API를 통해 first mile 추적 번호를 생성하고 v2.first_mile.get_tracking_number_list API를 사용하여 생성된 first mile 추적 번호 목록을 쿼리할 수 있습니다 (ship_method가 self-deliver인 경우: first_mile_tracking_number=order_sn, self-deliver logistics_channel_id는 null로 업로드해야 합니다).

## FAQ

**Q: first-mile 추적 번호에 대해 여러 상점의 주문을 바인딩하는 것이 허용됩니까?**
A: 예, 하지만 여러 상점의 주문이 동일한 운송/환승/창고를 사용하는지 확인하십시오. API에는 인증이 있으므로 v2.first_mile.bind_first_mile_tracking_number API를 호출할 때 한 번의 호출에 대해 하나의 상점 주문만 호출할 수 있으므로 매번 호출할 때 order_sn과 shop_id가 일치하는지 확인하십시오.

**Q: 바인딩에 성공하면 주문 상태가 변경됩니까?**
A: 예, FM 바인딩에 성공하고 스캔되면 바인딩된 모든 주문의 주문 상태가 PROCESSED에서 SHIPPED로 업데이트됩니다.

**Q: 주문을 배송하고 항공 운송장을 먼저 인쇄해야 합니까, 아니면 first mile 바인딩을 먼저 해야 합니까?**
A: 올바른 순서는 주문을 먼저 배송한 다음 항공 운송장을 인쇄하고 주문에 대한 first mile 추적 번호를 바인딩하는 것입니다.

**Q: 주문의 first mile 추적 번호를 어떻게 확인할 수 있습니까?**
A: v2.logistics.get_tracking_number API를 사용하고 요청 매개변수 response_optional_fields: first_mile_tracking_number를 업로드하면 주문의 first-mile 추적 번호를 얻을 수 있습니다.

**Q: first_mile_tracking_number 상태는 무엇으로 바인딩할 수 있습니까?**
A: v2.first_mile.get_detail API를 통해 first-mile 추적 번호 상태를 쿼리할 수 있습니다.

1. ship_method가 수거인 경우:
   - first_mile_tracking_number가 방금 생성된 경우 상태는 NOT_AVAILABLE이며, 이때 주문을 바인딩할 수 있습니다.
   - first_mile_tracking_number 추적 상태가 ORDER_RECEIVED인 경우 이미 바인딩된 주문이 있음을 의미하며 다른 주문을 계속 바인딩할 수 있습니다.
   - first_mile_tracking_number 상태가 PICKED_UP인 경우 소포가 채널에 의해 수거되었음을 나타내며 더 이상 주문을 바인딩할 수 없습니다.
   - first_mile_tracking_number 상태가 DELIVERED인 경우 소포가 창고에 도착했음을 의미하며 더 이상 주문을 바인딩할 수 없습니다.

2. ship_method가 위탁 또는 자체 배송인 경우: 상태 제한이 없습니다.

**Q: 자체 배송도 first mile 바인딩을 위해 API를 호출해야 합니까?**
A: 예, 자체 배송은 판매자의 오프라인 배송 작업이지만 first-mile 바인딩을 위해 v2.first_mile.bind_first_mile_tracking_number API를 호출해야 합니다.

**Q: 주문에 대한 first mile 추적 번호를 어떤 상태로 언바인딩할 수 있습니까?**
A: 1. ship_method가 수거인 경우 first_mile_tracking_number 상태가 ORDER_RECEIVED인 경우에만 언바인딩할 수 있습니다.
   2. ship_method가 위탁 또는 자체 배송인 경우 first_mile_tracking_number 상태 제한이 없습니다.

## 데이터 정의

**First Mile 추적 번호 상태**
- ORDER_RECEIVED
- PICKED_UP
- DELIVERED

## 사용 사례

1. first-mile 추적을 배송 관리 시스템에 통합
2. 주문을 first-mile 추적 번호에 바인딩하는 프로세스 자동화
3. API를 통한 다양한 배송 방법 (픽업, 드롭오프, 자체 배송) 처리
4. first-mile 추적 및 주문 상태 업데이트 관련 문제 해결

## 관련 API

- v2.first_mile.generate_first_mile_tracking_number
- v2.first_mile.get_tracking_number_list
- v2.first_mile.get_unibind_order_list
- v2.first_mile.bind_first_mile_tracking_number
- v2.first_mile.get_detail
- v2.first_mile.get_waybill
- v2.logistics.get_tracking_number

---

## 원문 (English)

### Summary

This guide outlines the best practices for first-mile binding in the Shopee API. It covers different shipping methods like pick-up, drop-off, and self-delivery, detailing the necessary steps and API calls for each. The guide also addresses frequently asked questions and provides data definitions for tracking number statuses.

### Content

# API Best Practices > First Mile Binding

## Terminology

**Pick up**: Channel door-to-door collection, only "logistics_channel_name": "shopeef" ("logistics_channel_id": 81s) support pick up.

**Drop off**: Delivered to channel outlets, expect "logistics_channel_name": "shopeef" ("logistics_channel_id": 81s) In"logistics_channel_name": "Self Deliver" ("logistics_channel_id": 0), other channels you get are drop off

**Self deliver**: Self-delivered, not using third-party logistics, "logistics_channel_name": "Self Deliver"

**first_mile_tracking_number**: first-mile tracking number

## Best Practise

For different types of first-mile shipping methods, they can be classified as:
- Pick up
- Drop off
- Self deliver

Developers can get the list of channels for first-mile and the shipping methods supported by the corresponding channels through the first_mile get_channel_list API.

### For pick up mode:

Step 1: Generate the first mile tracking number via v2.first_mile.generate_first_mile_tracking_number API and v2.first_mile.get_tracking_number_list API can be used to query the list of generated first mile tracking numbers

Step 2: v2.first_mile.get_unibind_order_list API can be used to query the list of orders to be bound with the first-mile tracking number, and then please bind them through v2.first_mile.bind_first_mile_tracking_number API. v2.first_mile.get_detail API can be used to query the first-mile tracking number details

Step 3: Call v2.first_mile.get_waybill API to print the first mile package label.

### For drop off mode:

Step 1: The seller obtains the channel tracking number offline.

Step 2: Get the list of orders to be bound with the first-mile tracking number through v2.first_mile.get_unbind_order_list API and binding them through v2.first_mile.bind_first_mile_tracking_number api. v2.first_mile.get_detail API can be used to query the first-mile tracking number details.

### For self deliver mode:

Step 1: The seller gets the list of orders to be bound with the first-mile tracking number through the v2.first_mile.get_unbind_order_list API.

Step 2: Generate the first mile tracking number via v2.first_mile.generate_first_mile_tracking_number API and v2.first_mile.get_tracking_number_list API can be used to query the list of generated first mile tracking numbers (If ship_method is self-deliver: first_mile_tracking_number=order_sn, self-deliver logistics_channel_id need to be uploaded self-deliver.logistics_channel_id need to be uploaded null).

## FAQ

**Q: Is it allowed to bind orders across shops for a first-mile tracking number?**
A: Yes, but make sure that orders across shops use the same transport/transit/warehouse. Because the API has authentication, when calling the v2.first_mile.bind_first_mile_tracking_number API, only one shop's order can be called for one call, so please make sure the order_sn and shop_id match for each time you call

**Q: Will the order status change after successful binding?**
A: Yes, after binding FM successfully and being scanned, the order status of all orders being bound will be updated from PROCESSED to SHIPPED.

**Q: Should I ship the order and print the airway bill first or first mile bind first?**
A: The correct order is to ship the order first, then print the airway bill, then bind the first mile tracking number for order.

**Q: How can I check the first mile tracking number of an order?**
A: You can use v2.logistics.get_tracking_number API and upload the request parameter response_optional_fields: first_mile_tracking_number then you can get the first-mile tracking number of an order

**Q: What can first_mile_tracking_number status be bound?**
A: You can query the first-mile tracking number status through v2.first_mile.get_detail API.

1. If ship_method is a pickup:
   - If the first_mile_tracking_number is just generated, the status is NOT_AVAILABLE, at that time, the order can be bound.
   - If the first_mile_tracking_number tracking status is ORDER_RECEIVED, it means that there are orders that have already been bound, and you can continue to bind other orders.
   - If the first_mile_tracking_number status is PICKED_UP, indicating that the parcels have been collected by the channel, and can no longer bind the order.
   - If the first_mile_tracking_number status is DELIVERED, which means that the parcel has arrived at the warehouse, and orders can no longer be bound.

2. If ship_method is drop off or self-deliver: there is no status restriction.

**Q: Does self-deliver still need to call the API for first mile binding?**
A: Yes, even self-delivery is the offline ship action for sellers, but you still need to call v2.first_mile.bind_first_mile_tracking_number API for the first-mile binding.

**Q: What status can I unbind the first mile tracking number for an order?**
A: 1. If ship_method is a pick up, only if the first_mile_tracking_number status is ORDER_RECEIVED can be unbird
   2. If ship_method is drop off or self-deliver, there is no first_mile_tracking_number status restriction.

## Data definition

**First Mile tracking number status**
- ORDER_RECEIVED
- PICKED_UP
- DELIVERED

---

**문서 ID**: developer-guide.225
**플랫폼**: shopee
**URL**: https://open.shopee.com/developer-guide/225
**처리 완료**: 2025-10-16T08:53:15
