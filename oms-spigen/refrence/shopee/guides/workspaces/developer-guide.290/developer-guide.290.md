# 콘솔 공지 | Shopee Entrega Direta

**카테고리**: 통합
**난이도**: 중간
**중요도**: 4/5
**최종 업데이트**: 2025-10-16T09:06:34

## 개요

본 가이드는 Shopee Entrega Direta 물류 채널에 대한 정보(API 변경 사항, 물류 허브 통합 세부 정보, 자주 묻는 질문 포함)를 제공합니다. Shopee Open Platform을 통해 Shopee Entrega Direta 채널에 액세스하고 활용하는 데 필요한 단계를 설명합니다.

## 주요 키워드

- Shopee Entrega Direta
- 물류 채널
- OpenAPI
- API integration
- 주문 관리
- 배송
- 물류 허브

## 본문

# 공지 | 콘솔

## API 모범 사례 > Shopee Entrega Direta

오후 3시 이전에 결제된 주문은 당일 배송되며, 오후 3시 이후에 결제된 주문은 다음 날 배송됩니다.

Shopee Entrega Direta는 현재 SP 시티에서만 서비스되므로, 판매자는 Shopee Entrega Direta가 활성화되면 항상 2개의 채널을 사용할 수 있습니다.

**참고:** *이 채널은 Shopee 상업 팀에서 정한 규칙에 따라 선택된 판매자에게만 제공됩니다. 자세한 내용은* [article](article) *을 참조하십시오.*

---

## OpenAPI 변경 사항

### 1 - 주문 식별 (새로운 물류 채널):

- v2.order.get_order_detail API, "shipping_carrier" 파라미터: "Shopee Entrega Direta"

### 2 - 판매자가 사용할 수 있는 채널 식별 및 상품 생성:

- v2.logistics.get_channel_list, 파라미터 "logistic_channel_id": 90022

### 3 - 주문 마감일 (ship_by_date):

- 주문 배송 마감일은 결제 확인 시점에 따라 결정되므로, 판매자가 v2.order.get_order_detail API의 "ship_by_date" 파라미터를 통해 제공되는 마감일에 접근할 수 있도록 해야 합니다.

### 4 - 상품 생성:

- 이 채널은 판매자에게 유일하게 제공되는 채널이 아니므로, 상품을 생성할 때 사용 가능한 채널을 식별할 때 항상 사용 가능한 모든 채널(예: Shopee Direct Delivery 및 Standard Delivery)을 언급하십시오.
- Shopee Direct Delivery라는 하나의 활성 물류 채널만 사용하여 상품을 생성하려고 시도하는 것은 불가능합니다.

---

## 물류 허브

파트너 및 물류 허브가 판매자가 물류 채널을 통해 주문을 관리하도록 지원하는 경우(예: Tracken, Log Manager) 통합을 위한 주요 API 및 흐름은 다음과 같습니다.

### 1 - Shopee Open Platform에서 계정 및 APP 생성:

- 판매자의 API 및 주문 데이터에 액세스하려면 Open Platform에서 계정을 생성해야 합니다.
- 계정이 생성되면 APP(모든 OpenAPI 기능에 액세스하려면 ERP 시스템 유형이 바람직함)을 생성하고 API 호출을 시작하기 전에 판매자를 계정에 연결합니다.
- 이러한 흐름은 다음 문서에 자세히 설명되어 있습니다.
  - [Developer account registration](Developer%20account%20registration)
  - [App management](App%20management)
  - [Authorization and Authentication](Authorization%20and%20Authentication)

### 2 - 권장 API 및 웹훅 (Push):

- v2.order.get_order_list - 주문 식별용;
- v2.order.get_order_detail API - 주문 세부 정보 및 주문의 물류 채널을 확인하는 유일한 방법 ("shipping_channel" 파라미터를 통해 "Shopee Entrega Direta"가 반환됨);
- v2.logistics.get_tracking_number API - 주문 추적 번호 식별용;
- v2.logistics.create_shipping_document - 라벨 생성용 (판매자의 ERP가 이미 이 호출을 수행하므로 필요하지 않을 수 있음);
- v2.logistics.get_shipping_document_result - 라벨이 이미 성공적으로 생성되었는지 확인하고 다운로드 가능한지 확인;
- v2.logistics.download_shipping_document API - 배송 라벨 다운로드용;
- order_status_push - 새 주문 및 주문 상태 업데이트 시 알림을 받기 위해;
- order_tracking_push - 추적 번호가 생성되었을 때 알림을 받기 위해;
- shipping_document_status_push - 라벨을 다운로드할 준비가 되었을 때 알림을 받기 위해;

다음은 Shopee의 주문 흐름에 대한 자세한 문서 링크입니다.
- [OpenAPI Logistics API Step by Step](OpenAPI%20Logistics%20API%20Step%20by%20Step)
- [API Call Flows](API%20Call%20Flows)

---

## FAQ:

### 1 - 모든 판매자가 Shopee Entrega Direta 채널에 액세스할 수 있습니까?

A: 관리되는 판매자만 채널에 액세스할 수 있습니다. 자세한 내용은 계정 관리자에게 문의하십시오.

### 2 - Shopee Entrega Direta 채널로 송장을 보내는 것이 필수입니까?

A: 이 채널은 송장 전송, 배송 준비 및 라벨 생성을 정상적으로 수행해야 합니다.

### 3 - v2.order.get_order_detail API에서 마스킹된 데이터, 왜 이런 일이 발생할 수 있습니까?

A: a) 구매자 데이터는 민감한 데이터이며 주문이 READY_TO_SHIP 및 TO_RETURN 상태일 때만 사용할 수 있습니다.

b) Open Platform Console의 IP Whitelist가 APP 수준에서 채워지지 않은 경우 데이터가 마스킹될 수 있습니다. 채워지면 올바른 상태에 있는 한 자동으로 사용할 수 있어야 합니다.

---

OpenAPI에 대한 추가 질문은 [OP Ticketing Platform](OP%20Ticketing%20Platform)에서 티켓을 생성할 수 있습니다.

## 사용 사례

1. 물류 허브를 Shopee 플랫폼과 통합합니다.
2. Shopee Entrega Direta 채널을 사용하여 주문을 관리합니다.
3. Shopee 판매자를 위한 ERP 시스템을 개발합니다.
4. 배송 라벨 생성 및 추적을 자동화합니다.
5. Shopee Entrega Direta 관련 API 변경 사항을 이해합니다.

## 관련 API

- v2.order.get_order_detail
- v2.logistics.get_channel_list
- v2.order.get_order_list
- v2.logistics.get_tracking_number
- v2.logistics.create_shipping_document
- v2.logistics.get_shipping_document_result
- v2.logistics.download_shipping_document

---

## 원문 (English)

### Summary

This guide provides information about the Shopee Entrega Direta logistics channel, including API changes, integration details for logistics hubs, and frequently asked questions. It outlines the necessary steps for accessing and utilizing the Shopee Entrega Direta channel through the Shopee Open Platform.

### Content

# Announcement | Console

## API Best Practices > Shopee Entrega Direta

paid before 3pm will be delivered on the same day, orders paid after 3pm will be delivered the next day).

The seller will always have 2 channels available when Shopee Entrega Direta is active, as Shopee Entrega Direta currently only serves SP city.

**Note:** *This channel will be made available to selected sellers, in accordance with the rules established by the Shopee commercial team. For more details follow the* [article](article).

---

## OpenAPI changes

### 1 - Order identification (new logistics channel):

- v2.order.get_order_detail API, "shipping_carrier" parameter: "Shopee Entrega Direta"

### 2 - Identification of channels available to the seller and creation of items:

- v2.logistics.get_channel_list, parameter "logistic_channel_id": 90022

### 3 - Order deadline (ship_by_date):

- As the order delivery deadline is defined by the moment payment is confirmed, it is necessary to ensure that the seller has access to the deadline provided by the v2.order.get_order_detail API, "ship_by_date" parameter.

### 4 - Item creation:

- This channel will never be the only one available to a seller, so whenever you create an item, when identifying the available channels, always mention all those available (Shopee Direct Delivery and Standard Delivery for example).
- Attempting to create an item with only one active logistics channel, namely Shopee Direct Delivery, will not be possible;

---

## Logistic HUBs

For partners and logistics HUBs that help sellers manage orders with their logistics channels (e.g. Tracken, Log Manager), below are the main APIs and flows for your integration:

### 1 - Create an account and APP on the Shopee Open Platform:

- To access sellers' APIs and order data, you will need to create an account on the Open Platform.
- Once the account is created, create an APP (preferably an ERP System type to have access to all OpenAPI functionalities) and connect sellers to their account before starting to call the API.
- These flows are further detailed in the following articles:
  - [Developer account registration](Developer%20account%20registration)
  - [App management](App%20management)
  - [Authorization and Authentication](Authorization%20and%20Authentication)

### 2 - Recommended APIs and Webhooks (Pushs):

- v2.order.get_order_list - for identifying orders;
- v2.order.get_order_detail API - for order details and the only way to check the order's logistics channel (via the "shipping_channel" parameter where "Shopee Entrega Direta" will be returned);
- v2.logistics.get_tracking_number API - to identify order tracking number;
- v2.logistics.create_shipping_document - for creating labels (it is worth noting that the seller's ERP will already make this call and it may not be necessary);
- v2.logistics.get_shipping_document_result - to check if the label has already been created successfully and is available for download;
- v2.logistics.download_shipping_document API - to download the shipping label;
- order_status_push - to be notified when a new order and when there is any order status update;
- order_tracking_push - To be notified when the tracking_number has been created;
- shipping_document_status_push - to be notified when the label is ready for download;

Below is a link to more articles about order flow on Shopee:
- [OpenAPI Logistics API Step by Step](OpenAPI%20Logistics%20API%20Step%20by%20Step)
- [API Call Flows](API%20Call%20Flows)

---

## FAQ:

### 1 - Do all sellers have access to the Shopee Entrega Direta channel?

A: Only managed sellers have access to the channel, for more information contact your Account Manager.

### 2 - Is it mandatory to send an invoice to the Shopee Entrega Direta channel?

A: The channel requires sending an invoice, organizing shipping and generating a label normally.

### 3 - Masked data in the v2.order.get_order_detail API, why can this happen?

A: a) buyer data is sensitive data and is only made available when the order is in READY_TO_SHIP and TO_RETURN status.

b) If the IP Whitelist (from the Open Platform Console) is not filled in at the APP level, the data can be masked, once filled in it should automatically be available (as long as it is in the correct status).

---

For further questions about OpenAPI, you can raise a ticket on the [OP Ticketing Platform](OP%20Ticketing%20Platform).

---

**문서 ID**: developer-guide.290
**플랫폼**: shopee
**URL**: https://open.shopee.com/developer-guide/290
**처리 완료**: 2025-10-16T09:06:34
