# 반품-환불 관리

**카테고리**: Integration
**난이도**: medium
**중요도**: 4/5
**최종 업데이트**: 2025-10-16T08:54:56

## 개요

본 가이드는 API 호출 흐름 및 데이터 정의를 포함하여 반품 및 환불 프로세스를 관리하는 방법에 대한 정보를 제공합니다. 환불만, 반품 & 환불 (분쟁 없음), 반품 & 환불 (분쟁)과 같은 다양한 반품 시나리오를 다루며 각 시나리오에 대한 API 상호 작용을 자세히 설명합니다.

## 주요 키워드

- 반품
- 환불
- API
- 분쟁
- 반품 상태
- 역물류
- 판매자
- 구매자
- 제안
- 솔루션

## 본문

# 반품-환불 관리

## 2 반품 API 호출 흐름

### 2.1 반품 주문 목록 및 상세 정보 가져오기

상점의 반품 주문 목록 및 요청 상세 정보를 가져오려면 각 애플리케이션은 정수 ID인 return_list를 반환합니다. 구매자는 동일한 주문에 대해 여러 개의 return_list를 제출할 수 있습니다. 반품 파라미터에는 이 반품 환불 신청과 관련된 주문 번호인 order_id가 포함됩니다. 또한 API는 상태, update_time_from, update_time_to, page_size 등을 기준으로 다양한 유형의 반품을 필터링하는 것을 지원합니다. 반품 목록에는 관련 취소 상태도 포함됩니다.

**반품에 대해서는 목록 API를 참조하십시오.** return_list를 얻은 후 반품 상세 정보를 얻게 됩니다.

**반품 상세 정보 API:** 사용 가능한 솔루션(AP)과 구매자에게 제공되는 사용 가능한 솔루션을 참조하십시오.

### 2.2 환불만 해당

```
Start
↓
returns.get_return_list
↓
returns.get_return_detail
↓
returns.get_available_solution
↓
returns.confirm
↓
End
```

**returns.confirm API 사용** = 구매자의 반품 신청에 동의하며, 전체 환불 유형에만 해당됩니다. 구매자는 제품을 반품할 필요가 없습니다. 환불에 대한 종료는 수락됨으로 처리됩니다.

### 2.3 반품 & 환불 (분쟁 없음)

```
Start
↓
returns.get_return_list
↓
returns.get_return_detail
↓
returns.get_available_solution
↓
returns.offer
↓
[Decision Diamond: Buyer Accept] → [Yellow: Seller Accept] 
↓                                    ↓
returns.get_return_detail ← returns.accept_offer
↓
End
```

**returns.offer 사용** - 판매자는 구매자가 선택할 수 있는 반품 계획을 제공합니다.

**returns.accept_offer 사용** - 판매자는 구매자가 제공한 반품 계획을 수락합니다.

### 2.4 반품 & 환불 (분쟁)

```
Start
↓
returns.get_return_list
↓
returns.get_return_detail
↓
returns.get_available_solution
↓
returns.offer ← [Purple: Seller doesn't accept solution or buyer's offer]
↓
[Purple: Buyer raise a dispute]
↓
returns.dispute
↓
returns.convert_image
↓
returns.upload_proof
↓
returns.query_proof
↓
returns.get_return_detail
↓
End
```

판매자 측에서는 먼저 **returns.get_available_solution API**를 호출하여 솔루션 또는 반품 주문을 결정한 다음 구매자가 선택할 수 있도록 **returns.offer API**를 호출합니다.

구매자가 판매자가 제공한 제안을 수락하지 않거나 판매자가 거절하는 경우 **returns.query_offer**를 호출하여 제안을 수락합니다.

---

# API 문서 - 반품 프로세스

## 개요

판매자는 `v2.returns.get_available_solutions` API 링크를 호출하여 반품 주문 옵션을 결정한 다음 구매자가 선택할 수 있도록 `v2.returns.offer` API를 호출할 수 있습니다.

구매자도 판매자에게 솔루션을 제공합니다. 판매자가 수락하면 `v2.returns.accept_offer`를 호출하여 제안을 수락할 수 있습니다. 수락할 수 없는 경우 `v2.returns.dispute`를 호출하여 분쟁 센터에 분쟁을 제기할 수 있습니다.

현재 OpenAPI는 REQUESTED 및 PROCESSING의 두 가지 상태만 분쟁을 지원합니다.

## 2.4.1 분쟁

판매자가 구매자의 반품 요청에 이의가 있는 경우 분쟁 센터에 가서 요청을 처리할 수 있습니다. `v2.returns.dispute`는 판매자가 반품 주문을 분쟁 센터로 에스컬레이션하는 데 사용됩니다.

### 2.4.1.1 이미지

분쟁이 제기된 후 판매자는 API를 통해 증거 이미지를 업로드할 수 있지만 현재 비디오 업로드는 지원되지 않습니다.
- `v2.returns.upload_image` : 이미지 업로드
- `v2.returns.bind_cancel` : 분쟁이 있는 이미지 연결

---

## 3 데이터 정의

### ReturnStatus
- REQUESTED
- ACCEPTED
- CANCELLED
- SHIPPED
- CLOSED
- PROCESSING
- SELLER_DISPUTE

### ReturnReason
- OK
- NOT_RECEIPT
- WRONG_ITEM
- ITEM_DAMAGED
- INCORRECT_DESCRIPTION
- MUTUAL_AGREE
- OTHER
- SIZE_FIT_UNMATCHING (베트남 전용)
- CHANGE_MIND
- ITEM_MISSING
- ITEM_INCOMPLETE_FAILED
- ITEM_FAKE
- PHYSICAL_DMG
- FUNCTIONAL_DMG

### ReturnDisputeReason
Reason:
- NOT_RECEIPT: 미수령 주장을 거부하고 싶습니다.
- OTHER: 요청을 거부하고 싶습니다.
- NOT_RECEIVED: 반품 요청에 동의하지만 제품을 받지 못했습니다.
- UNKNOWN

### ReturnSolution
- RETURN_REFUND
- REFUND

### NegotiationStatus
- PENDING_RESPOND
- PENDING_BUYER_RESPOND
- RESOLVED

### SellerProofStatus
- PENDING
- UPLOADED
- OVERDUE

### SellerCommunicationStatus
- COMPENSATION_NOT_APPLICABLE
- COMPENSATION_INITIAL_STAGE
- COMPENSATION_PENDING_REQUEST
- COMPENSATION_NOT_REQUIRED
- COMPENSATION_WAITING_APPROVAL
- COMPENSATION_APPROVED
- COMPENSATION_REJECTED
- COMPENSATION_PAID_TO_BUYER
- COMPENSATION_NOT_ELIGIBLE

### Return Refund Request Type
- 0: Normal RR - RR은 예상 배송 날짜를 기준으로 구매자가 소포를 받은 후에 제기됩니다.
- 1: In-transit RR
  - 1: In-transit RR-RR은 품목이 아직 구매자에게 배송 중인 동안 구매자가 제기합니다.
  - 2: Return-on-the-Spot RR은 구매자가 배송 시 소포를 받은 후 구매자가 제기합니다.

### Validation Type
- seller_validation - 반품 소포가 있는 반품 및 관련 요청의 경우 판매자에게 배송되어 유효성 검사 및 구매자 환불 여부 또는 분쟁 제기 여부를 결정합니다.
- seller_auto_decision - 반품 소포가 없는 반품 요청의 경우 판매자 유효성 검사를 통해 결정하고 구매자 환불 여부 또는 분쟁 제기 여부를 결정해야 합니다.

---

## 역물류 상태

### [일반 반품]
- **LOGISTICS_PENDING_ARRANGE**: 반품은 현재 사용자가 배송 옵션을 선택하기를 기다리고 있습니다. 통합 물류와 비통합 물류 모두 동일합니다.
- **LOGISTICS_READY**: 사용자가 배송 옵션을 선택했으며 시스템에서 물류 요청을 생성하기를 기다리고 있습니다. 추적 번호는 아직 사용할 수 없습니다. 통합 물류와 비통합 물류 모두 동일합니다.
- **LOGISTICS_READY_TO_SHIP**: 시스템에서 물류 요청을 생성했습니다. 추적 번호를 사용할 수 있어야 합니다.
- **LOGISTICS_PICKUP_RETRY**: 타사 물류 제공업체가 구매자로부터 소포를 픽업하기 위해 다시 시도합니다. 타사 물류 제공업체가 Shopee에 다시 업데이트하므로 통합 물류입니다.
- **LOGISTICS_PICKUP_FAILED**: 타사 물류 제공업체가 구매자로부터 소포를 픽업하지 못했습니다. 타사 물류 제공업체가 Shopee에 다시 업데이트하므로 통합 물류에 사용할 수 있습니다.
- **LOGISTICS_PICKUP_DONE**: 통합 물류의 경우 이는 소포가 3 PHS 당사자(출발 택배)에 의해 픽업되었음을 의미합니다. 비통합 물류의 경우 이는 사용자가 소포를 드롭오프 지점에 떨어뜨렸음을 의미합니다.
- **LOGISTICS_LOST**: 소포가 분실된 것으로 표시되었습니다. 타사 물류 제공업체가 Shopee에 다시 업데이트하므로 통합 물류에만 사용할 수 있습니다.
- **LOGISTICS_DELIVERY_DONE**: 소포가 판매자에게 성공적으로 배송되었습니다. 타사 물류 제공업체가 Shopee에 다시 업데이트하므로 통합 물류에만 사용할 수 있습니다.

### [배송 중 RR]
- Preparing
- Delivered
- Delivery Failed
- Lost

### [현장 반품]
- Preparing
- Delivered
- Delivery Failed
- Lost

---

## 반품 후 물류 상태

창고에서 판매자에게 다시 돌아오는 역물류를 나타내는 모든 물류 상태입니다.
- **POST_RETURN_LOGISTICS_PICKUP_PENDING**: 추적 번호와 함께 물류 요청이 성공적으로 생성되었습니다.
- **POST_RETURN_LOGISTICS_REQUEST_CANCELLED**: 창고 팀에서 물류 요청을 취소했습니다.
- **POST_RETURN_LOGISTICS_PICKUP_FAILED**: 창고에서 판매자에게 소포를 픽업하지 못했습니다.
- **POST_RETURN_LOGISTICS_PICKUP_RETRY**: 타사 물류가 창고에서 소포를 픽업하기 위해 다시 시도합니다.
- **POST_RETURN_LOGISTICS_PICKUP_DONE**: 목적지로 가는 도중에 성공적으로 픽업되었습니다.
- **POST_RETURN_LOGISTICS_DELIVERY_FAILED**: 소포 배송에 실패했습니다. 운전자가 창고로 다시 시도합니다.
- **POST_RETURN_LOGISTICS_DELIVERY_DONE**: 소포가 성공적으로 배송되었습니다.
- **POST_RETURN_LOGISTICS_LOST**: 물류 제공업체에서 소포를 분실했습니다.

## 사용 사례

1. 구매자의 반품 요청 관리
2. 구매자의 환불 처리
3. 반품 관련 분쟁 처리
4. 반품 및 환불 프로세스를 플랫폼과 통합
5. 반품 물류 자동화

## 관련 API

- returns.get_return_list
- returns.get_return_detail
- returns.get_available_solution
- returns.confirm
- returns.offer
- returns.accept_offer
- returns.dispute
- returns.convert_image
- returns.upload_proof
- returns.query_proof
- v2.returns.get_available_solutions
- v2.returns.offer
- v2.returns.accept_offer
- v2.returns.dispute
- v2.returns.upload_image
- v2.returns.bind_cancel

---

## 원문 (English)

### Summary

This guide provides information on how to manage return and refund processes, including API call flows and data definitions. It covers different return scenarios like refund only, return & refund (no dispute), and return & refund (dispute), detailing the API interactions for each.

### Content

# Return-Refund Management

## 2 Return API Call Flow

### 2.1 Getting the list of return orders and details

To get list of return orders and request details for a shop, Each application will return a return_list. as an integer ID. Buyers may submit multiple return_list for the same order. The return parameter contains order_id where is the order number associated with this return refund application. In addition, the API supports filtering different types of returns based on status, update_time_from, update_time_to, page_size etc. The return list also contains the related Cancellation status.

**For returns, see listing API.** After return_list, we will get return details.

**For return detail API:** See the available solutions (AP) and the available solutions offered to buyers.

### 2.2 Refund Only

```
Start
↓
returns.get_return_list
↓
returns.get_return_detail
↓
returns.get_available_solution
↓
returns.confirm
↓
End
```

**By returns.confirm API** = Agree to the buyer's return application, only for the Full Refund type. the buyer does not need to return the product. Close against the refund will be credited to Accepted.

### 2.3 Return & Refund (No Dispute)

```
Start
↓
returns.get_return_list
↓
returns.get_return_detail
↓
returns.get_available_solution
↓
returns.offer
↓
[Decision Diamond: Buyer Accept] → [Yellow: Seller Accept] 
↓                                    ↓
returns.get_return_detail ← returns.accept_offer
↓
End
```

**By returns.offer** - The seller provides a return plan for the Buyer to choose.

**By returns.accept_offer** - The seller accepts the return plan provided by the Buyer

### 2.4 Return & Refund (Dispute)

```
Start
↓
returns.get_return_list
↓
returns.get_return_detail
↓
returns.get_available_solution
↓
returns.offer ← [Purple: Seller doesn't accept solution or buyer's offer]
↓
[Purple: Buyer raise a dispute]
↓
returns.dispute
↓
returns.convert_image
↓
returns.upload_proof
↓
returns.query_proof
↓
returns.get_return_detail
↓
End
```

The seller side call **returns.get_available_solution API** first, to determine the solution or the return order, and then call **returns.offer API** for the buyer to choose.

If the buyer doesn't accept the offer provided by the seller or the seller declines, then call **returns.query_offer** to accept the offer.

---

# API Documentation - Return Process

## Overview

The seller can call `v2.returns.get_available_solutions` API link to determine the options of the return order, and then call `v2.returns.offer` API for the buyer to choose.

Buyers will also provide solutions to sellers. If the seller accepts, they can call `v2.returns.accept_offer` to accept the offer. If you cannot accept it, then you can call `v2.returns.dispute` to file a dispute with the Dispute Center.

At present, OpenAPI only supports two statuses of REQUESTED and PROCESSING to dispute.

## 2.4.1 Dispute

If the seller has a dispute with the buyer's return request, he can go to the Dispute Center to handle the request. `v2.returns.dispute` Used by sellers to escalate return orders to the dispute center.

### 2.4.1.1 Image

After a dispute is raised, the seller can upload evidence images through API, but uploading videos is not currently supported
- `v2.returns.upload_image` : Upload images
- `v2.returns.bind_cancel` : Obtlr.y disputed images

---

## 3 Data Definition

### ReturnStatus
- REQUESTED
- ACCEPTED
- CANCELLED
- SHIPPED
- CLOSED
- PROCESSING
- SELLER_DISPUTE

### ReturnReason
- OK
- NOT_RECEIPT
- WRONG_ITEM
- ITEM_DAMAGED
- INCORRECT_DESCRIPTION
- MUTUAL_AGREE
- OTHER
- SIZE_FIT_UNMATCHING (Private for Vietnam)
- CHANGE_MIND
- ITEM_MISSING
- ITEM_INCOMPLETE_FAILED
- ITEM_FAKE
- PHYSICAL_DMG
- FUNCTIONAL_DMG

### ReturnDisputeReason
Reason:
- NOT_RECEIPT: I would like to reject the non-receipt claim
- OTHER: I would like to reject the request
- NOT_RECEIVED: I agree with the return request, but I did not receive product(s)
- UNKNOWN

### ReturnSolution
- RETURN_REFUND
- REFUND

### NegotiationStatus
- PENDING_RESPOND
- PENDING_BUYER_RESPOND
- RESOLVED

### SellerProofStatus
- PENDING
- UPLOADED
- OVERDUE

### SellerCommunicationStatus
- COMPENSATION_NOT_APPLICABLE
- COMPENSATION_INITIAL_STAGE
- COMPENSATION_PENDING_REQUEST
- COMPENSATION_NOT_REQUIRED
- COMPENSATION_WAITING_APPROVAL
- COMPENSATION_APPROVED
- COMPENSATION_REJECTED
- COMPENSATION_PAID_TO_BUYER
- COMPENSATION_NOT_ELIGIBLE

### Return Refund Request Type
- 0: Normal RR - RR is raised by the buyer after they have received the parcel, based on estimated delivery date
- 1: In-transit RR
  - 1: In-transit RR-RR is raised by the buyer while item is still in-transit to buyer)
  - 2: Return-on-the-Spot RR is raised by the buyer after buyer received parcel at delivery)

### Validation Type
- seller_validation - For Return & Related requests with return parcel that will be delivered to the seller for validation and decision whether to refund buyer or to raise dispute
- seller_auto_decision - For return request where there will be no return parcel and seller will need to decide via seller validation and decision whether to refund buyer or to raise Dispute

---

## Reverse Logistics Status

### [Normal Return]
- **LOGISTICS_PENDING_ARRANGE**: Return is now pending user to select shipping option. Same for both integrated logistics and non-integrated logistics.
- **LOGISTICS_READY**: User has selected shipping option, and pending system to create logistics request. Tracking number is not yet available. Same for both integrated logistics and non-integrated logistics.
- **LOGISTICS_READY_TO_SHIP**: System has created logistics request. Tracking number should be available.
- **LOGISTICS_PICKUP_RETRY**: Third party logistics provider will make another attempt to pick up parcel from buyer. Integrated logistics since this is updated by third party logistics provider back to Shopee.
- **LOGISTICS_PICKUP_FAILED**: Third party logistics provider has failed to pickup parcel from buyer. Available for integrated logistics since this is updated by third party logistics provider back to Shopee.
- **LOGISTICS_PICKUP_DONE**: For integrated logistics, this implies the parcel has been picked up by a 3 PHS party (origin courier). For non-integrated logistics, this means the user has dropped the parcel at dropoff point.
- **LOGISTICS_LOST**: Parcel has been marked as lost. Only available for integrated logistics since this is updated by third party logistics provider back to Shopee.
- **LOGISTICS_DELIVERY_DONE**: Parcel has been successfully delivered to seller. Only available for integrated logistics since this is updated by third party logistics provider back to Shopee.

### [In-transit RR]
- Preparing
- Delivered
- Delivery Failed
- Lost

### [Return-on-the-Spot]
- Preparing
- Delivered
- Delivery Failed
- Lost

---

## Post Return Logistics Status

All logistics status that represent reverse logistics from warehouse back to seller.
- **POST_RETURN_LOGISTICS_PICKUP_PENDING**: Logistics request created successfully with tracking number.
- **POST_RETURN_LOGISTICS_REQUEST_CANCELLED**: Logistics request cancelled by warehouse team.
- **POST_RETURN_LOGISTICS_PICKUP_FAILED**: Failed to pickup parcel from warehouse back to seller.
- **POST_RETURN_LOGISTICS_PICKUP_RETRY**: Third party logistics will retry to pickup the parcel from warehouse.
- **POST_RETURN_LOGISTICS_PICKUP_DONE**: Successful pickup on the way to destination.
- **POST_RETURN_LOGISTICS_DELIVERY_FAILED**: Failed delivery of parcel. Driver will make another attempt back to warehouse.
- **POST_RETURN_LOGISTICS_DELIVERY_DONE**: Successful delivery of parcel.
- **POST_RETURN_LOGISTICS_LOST**: Parcel lost by logistics provider.

---

**문서 ID**: developer-guide.227
**플랫폼**: shopee
**URL**: https://open.shopee.com/developer-guide/227
**처리 완료**: 2025-10-16T08:54:56
