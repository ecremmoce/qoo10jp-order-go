# API 모범 사례 - SIP 모범 사례

**카테고리**: 모범 사례
**난이도**: 중간
**중요도**: 4/5
**최종 업데이트**: 2025-10-16T08:57:30

## 개요

본 가이드는 SIP (Shop in Place) API 사용에 대한 모범 사례를 설명하며, 인증, 제품 관리, 가격 로직을 다룹니다. 비 CNSC 및 CNSC 판매자 모두를 위한 인증 프로세스와 토큰 획득 및 상점 관계 관리에 대한 지침을 자세히 설명합니다.

## 주요 키워드

- SIP
- API
- Authorization
- Product Management
- CNSC
- Access Token
- Refresh Token
- Shop Relationship
- Price Logic

## 본문

# API 모범 사례 - SIP 모범 사례

**배송 시간 약 2주**

---

## 용어

- **Primary shop (주요 상점)**: SIP (개방형 운영 모드)에서 판매자가 (임시로) 운영하는 상점을 주요 상점이라고도 합니다.
- **Local SIP (로컬 SIP)**: 로컬 판매자가 주요 상점을 운영하는 경우 해당 SIP 모드를 로컬 SIP라고 부릅니다.
- **CNSC SIP**: CNSC 판매자가 SIP 상점(아파트)을 운영하는 경우, 판매자(또는 제휴 판매자)가 될 보조자 및 SIP 상점 설정도 의미합니다. 사실, 여기서의 SIP 상점(로컬 모드)과 로컬 SIP 모드는 실질적으로 동일하며, 판매자와 판매자가 운영하는 상점이 루트 SIP 및 SIP 아파트 사용자에 배치된다는 점을 제외하고는 명확한 경계가 없습니다.
- **SIP shop product (SIP 상점 상품)**: SIP 상점에서 판매하는 상품을 의미하며, 작년 말 기준 상품 및 SIP 가격(아파트) 사용자를 기반으로 상점 내 상품을 의미합니다. 비 SIP 모드의 경우, 상품의 원가와 작년 말 기준 특가 가격이 상점 상품의 원가입니다.

**참고:** 소비자 혼란을 방지하기 위해 상점 상품 또는 CNSC 상품을 다루든 상점 상품 표시는 항상 상품과 동일한 상태로 유지됩니다.

---

## 1. SIP 상점 권한 부여

### 1.1 non CNSC Seller (비 CNSC 판매자)

#### 1.1.1 Authorization process (권한 부여 프로세스)

로컬 SIP 상점을 성공적으로 생성하면 자동으로 판매자가 되어 상점 인증을 받게 됩니다. 판매자가 SIP 상점 계정에 로그인하면 CNSC 공식 검토를 기다려야 합니다. 검토가 통과되면 SIP 권한 부여 페이지를 받고 권한 부여를 확인하기 위해 클릭합니다. SIP 상점은 이 API에 대한 권한을 부여받게 됩니다.

### Authorization (권한 부여)

**Local (로컬)**

[번호가 매겨진 1-4단계가 표시된 권한 부여 흐름도]

**Login (로그인)**
Seller CNSC account or System Store (판매자 CNSC 계정 또는 시스템 스토어)

Apply for authorization to get the following information: (다음 정보를 얻기 위해 권한 부여를 신청하십시오.)
- **Product (상품)**: Authorization of product and content management for products period (상품 기간에 대한 상품 및 콘텐츠 관리 권한 부여)
- **Payment (결제)**: Information of transactions and refunds to your shop account (상점 계정으로의 거래 및 환불 정보)
- **Marketing (마케팅)**: Information related to discounts in your shops (상점의 할인 관련 정보)
- **Logistics (물류)**: Information related to order logistics (주문 물류 관련 정보)

[주황색으로 표시된 AUTHORIZE 버튼]

#### 1.1.2 Getting code (코드 가져오기)

권한 부여에 성공하면 콜백 URL을 통해 shop_id 및 사용 가능한 코드를 직접 얻을 수 있습니다.

`sip_code_exchange_access_API`를 호출하고 IT 상점의 shop id 및 code (상점 액세스 토큰 및 리프레시 토큰을 가져오는 API)를 업로드할 수 있습니다.

#### 1.1.3 Refresh Access Token (액세스 토큰 갱신)

이때, 앞에서 얻은 액세스 토큰과 리프레시 토큰은 IT 상점에서 사용할 수 있습니다. 그런 다음 `sip_code_refresh_access_API`를 호출하여 IT 상점의 액세스 토큰과 리프레시 토큰을 갱신할 수 있습니다.

**Things to note (주의 사항)**:
- After the new access_token is generated, the old access_token is still valid within 5 minutes (새 access_token이 생성된 후에도 이전 access_token은 5분 동안 유효합니다.)
- If the new refresh_token is generated at the same time, the old refresh_token will become invalid immediately. It must be saved immediately after the new refresh_token is generated, and be used later to refresh. (새 refresh_token이 동시에 생성되면 이전 refresh_token은 즉시 무효화됩니다. 새 refresh_token이 생성된 후 즉시 저장하고 나중에 갱신하는 데 사용해야 합니다.)
- If the new refresh_token and access_token renewed are lost, please check FAQ. (갱신된 새 refresh_token 및 access_token을 분실한 경우 FAQ를 확인하십시오.)

---

### 1.2 CNSC Seller (CNSC 판매자)

#### 1.2.1 Authorization process (권한 부여 프로세스)

단일 "Authorization and Authentication (권한 부여 및 인증)"을 참조하십시오. 판매자가 기본 계정에 로그인하여 권한 부여 페이지에 들어가면 생성한 모든 주요 SIP 상점을 볼 수 있으며 "Local (로컬)"을 추가할 수도 있습니다(SIP IT Shop이 "Confirm Authorization Status (권한 부여 상태 확인)"를 클릭하지 않은 경우 SIP IT Shop을 확인하고 "Confirm (확인)" 권한 부여 상태를 클릭합니다. SIP IT Shop이 IT 상점 시스템에 있는 경우 IT 상점). 권한 부여 후 SIP IT 상점의 액세스 토큰과 리프레시 토큰이 자동으로 생성됩니다.

[다음 내용이 포함된 상점 권한 부여 인터페이스 스크린샷:
- Goshh Seller
- Authorized Shops List (권한 부여된 상점 목록)
- Affiliate Shop List (ID) (제휴 상점 목록 (ID))
- Shop information fields including: (다음 내용을 포함한 상점 정보 필드:)
  - Name (shop name, display name, promotions, show details must fill) (이름 (상점 이름, 표시 이름, 프로모션, 세부 정보 표시는 필수))
  - Shop Postcode (상점 우편 번호)
  - Shop Address (상점 주소)
  - Shop Telephone (상점 전화번호)
  - Description (설명)
  - Contacts (candAllo2) (연락처 (candAllo2))
  - LegalAddress (법적 주소)
  - SortOrder (정렬 순서)
  - And various shop IDs and status indicators (다양한 상점 ID 및 상태 표시기)

**Note (참고)**: At this time, the account bound to a merchant is a main account (let be bound under the "Unregistered Shop Group" list. It does in the [Unregistered Shop Group] module, it will show all the shops that merchants have). This status will be marked under it. (이때, 판매자에게 바인딩된 계정은 기본 계정입니다("Unregistered Shop Group (미등록 상점 그룹)" 목록에 바인딩됩니다. [Unregistered Shop Group (미등록 상점 그룹)] 모듈에 있는 경우 판매자가 보유한 모든 상점을 표시합니다). 이 상태는 그 아래에 표시됩니다.)

[다음 내용이 포함된 추가 인터페이스 스크린샷:
- Unassigned007 shop (미할당007 상점)
- Unregistered Shop Group - Suspended (미등록 상점 그룹 - 일시 중단됨)
- Multiple shop entries with status indicators (상태 표시기가 있는 여러 상점 항목)
- Audit Managed and Reset at SPS markers (SPS 마커에서 감사 관리 및 재설정)]

#### 1.2.2 Getting code (코드 가져오기)

권한 부여에 성공하면 콜백 URL을 통해 main_account_id 및 사용 가능한 코드를 직접 얻을 수 있습니다.

#### 1.2.3 Getting token (토큰 가져오기)

`sip_code_cnsc_exchange_access_API`를 호출하고 main_account_id 및 code (현재 권한이 부여된 판매자를 가져오는 API이며, sicp_id_list는 현재 권한이 부여된 모든 판매자를 반환합니다)를 업로드할 수 있습니다.

**Note (참고)**: The main_account_id of the currently authorized merchant account access token refresh_token of IT shop. (현재 권한이 부여된 판매자 계정의 main_account_id는 IT 상점의 액세스 토큰 refresh_token입니다.)

#### 1.2.4 Getting shop relationship (상점 관계 가져오기)

언급된 API로 main_account_id 및 sicp_id_list를 가져온 후 이 판매자와 연결된 권한 부여된 상점 목록을 가져오고 `sip_cnsc_list_cnsc_shop_API`를 설정합니다. 그러면 각 목록과 연결된 권한 부여된 상점 목록의 1)(상점은 판매자 목록에 속합니다. HTML의 상점 목록은 이 상점 목록에 있으며 sicp_id_list는 해당 목록에 있어야 합니다)과 같은 정보를 얻을 수 있습니다.

**Note (참고)**:
- If it is an normal shop, (i.e., non-does not the sip_id), shops level will not be returned. (일반 상점인 경우(즉, sip_id가 없는 경우) 상점 수준이 반환되지 않습니다.)
- If it is a SIP shop managed by CNSC (that is, the sip_id), shops level will be returned and ask for the sicp_value. (CNSC에서 관리하는 SIP 상점인 경우(즉, sip_id가 있는 경우) 상점 수준이 반환되고 sicp_value를 요청합니다.)
- If it is a SIP shop managed by merchant (that is, the sip_id), shops level will be returned but the sicp_value will be null. (판매자가 관리하는 SIP 상점인 경우(즉, sip_id가 있는 경우) 상점 수준이 반환되지만 sicp_value는 null입니다.)

#### 1.2.5 Refresh Access Token (액세스 토큰 갱신)

`sip_code_refresh_access_API`를 참조하여 만료될 때 각 sicp_id의 액세스 토큰과 리프레시 토큰을 갱신하십시오. `sip_code_refresh_access_API`를 호출할 때 각 sicp_id(각 판매자)의 액세스 토큰과 리프레시 토큰을 하나씩 갱신하십시오.

---

## 2. Product Management (상품 관리)

### 2.1 Product information logic (상품 정보 로직)

**Info (정보)**

**Sync logic (동기화 로직)**

**Item types info (item) (상품 유형 정보 (상품))**
- Creator (생성자): After the seller creates a IT shop product, Shopee will automatically transcode (판매자가 IT 상점 상품을 생성하면 Shopee가 자동으로 트랜스코딩합니다.)

---

*[End of visible content]*

---

# 2. Product Management (상품 관리)

## 2.1 Product information logic (상품 정보 로직)

### Info (정보)

**Sync logic (동기화 로직)**

**Item base info (Item base + category) (상품 기본 정보 (상품 기본 + 카테고리))**
- **Create (생성)**: After the seller creates a P-Help product, Shoppe will automatically translate the product information to the target shop language via P-Help API. (판매자가 P-Help 상품을 생성하면 Shoppe가 P-Help API를 통해 상품 정보를 대상 상점 언어로 자동 번역합니다.)
- **Update (업데이트)**: If the seller modifies the P-Help product, Shoppe will automatically synchronize the modifications to the target shop. (판매자가 P-Help 상품을 수정하면 Shoppe가 자동으로 수정 사항을 대상 상점에 동기화합니다.)

**Item status (상품 상태)**
- If the P-Help product is unlisted or deleted, the A shop product will also be unlisted or deleted. (P-Help 상품이 목록에서 제거되거나 삭제되면 A 상점 상품도 목록에서 제거되거나 삭제됩니다.)

**Item stock (상품 재고)**
- The product stock of P-Help is equal to the product stock of A shop. For example, there are 2 A shops under the seller, each add 4 of buyer, and the stock of P-Help (P-Shop) will be reduced by 8. If the seller wants to modify the stock, they can only modify the P-Help product stock in A shop. (P-Help의 상품 재고는 A 상점의 상품 재고와 같습니다. 예를 들어 판매자 아래에 2개의 A 상점이 있고 각 상점에 구매자가 4명씩 추가되면 P-Help(P-Shop)의 재고가 8개 줄어듭니다. 판매자가 재고를 수정하려면 A 상점에서 P-Help 상품 재고만 수정할 수 있습니다.)

**Item price (상품 가격)**
- When the seller updates the product price of P-shop, Shoppe will synchronize it to the product price in A shop. (판매자가 P-shop의 상품 가격을 업데이트하면 Shoppe가 A 상점의 상품 가격에 동기화합니다.)

---

## 2.2 Price logic (가격 로직)

Please note that all APIs using `sa_currency` and `sa_product_per_service` API Response: `currency: "MYR"`, `buyer_shipping_fee: "119.90"`, `original_price: "129.90"`, `tag_discounted_price: "119.90"`.

**Note (참고)**: If the item is a P shop item and there are currently subsidiaries, please call the `sa_product_per_service_API` to get the sa_currency, item_price, currency of each. (상품이 P 상점 상품이고 현재 자회사가 있는 경우 `sa_product_per_service_API`를 호출하여 각 자회사의 sa_currency, item_price, currency를 가져오십시오.)

- If the item is a P shop item, please call the `sa_product_per_service_API` to get the sa_currency, item_price, currency of each subsidiary store. (상품이 P 상점 상품인 경우 `sa_product_per_service_API`를 호출하여 각 자회사의 sa_currency, item_price, currency를 가져오십시오.)
- If the item is a sub "B" shop item but not `sa_get_pricing_item_price` separately, item_price_currency fields will not return in v2 API response. (상품이 하위 "B" 상점 상품이지만 `sa_get_pricing_item_price`가 별도로 없는 경우 item_price_currency 필드는 v2 API 응답에서 반환되지 않습니다.)

---

## 3. Order management (주문 관리)

### 3.1 Order synchronization logic (주문 동기화 로직)

#### Steps (단계)

Please follow the SIP P shop order list and the SIP A shop order list respectively for fulfillment or cancellation. Please check seller cancellation list to the item cannot of the seller is no SIP A shop under synchronization rules or the seller want to cancel the order, or when the buyer order times out. (이행 또는 취소를 위해 SIP P 상점 주문 목록과 SIP A 상점 주문 목록을 각각 따르십시오. 판매자가 동기화 규칙에 따라 SIP A 상점이 없거나 판매자가 주문을 취소하려는 경우 또는 구매자 주문 시간이 초과된 경우 판매자 취소 목록을 확인하십시오.)

#### 3.1.1 2 Local SIP

Shoppe automatically synchronizes the SIP A shop order to the SIP P shop, so you only need to obtain and fulfill the SIP P shop order list. The seller only needs to operate in the SIP P shop. Shoppe will automatically synchronize it to the SIP A shop. Examples include a Local SG Shopper buys/receives warehouses from Seller P-Shop Singapore to Local SG Shopper buys/receives from Seller A-Shop Singapore. (Shoppe는 자동으로 SIP A 상점 주문을 SIP P 상점에 동기화하므로 SIP P 상점 주문 목록을 얻고 이행하기만 하면 됩니다. 판매자는 SIP P 상점에서만 운영하면 됩니다. Shoppe는 자동으로 SIP A 상점에 동기화합니다. 예를 들어 로컬 SG 쇼핑객이 판매자 P-Shop 싱가포르에서 창고를 구매/수령하여 로컬 SG 쇼핑객이 판매자 A-Shop 싱가포르에서 구매/수령하는 경우가 있습니다.)

**Note A (참고 A)**

For the same order merchandise, The orders of P shops and A shops will be pushed at the same time. If no SIP sellers, the order will only be pushed to the p shop. So sellers only need to pay attention to the P shop orders, they only need to pay attention to the p shop orders when the seller order list. (동일한 주문 상품의 경우 P 상점과 A 상점의 주문이 동시에 푸시됩니다. SIP 판매자가 없는 경우 주문은 p 상점으로만 푸시됩니다. 따라서 판매자는 P 상점 주문에만 주의하면 되며, 판매자 주문 목록에 있을 때 p 상점 주문에만 주의하면 됩니다.)

**Note B (참고 B)**

For more order fulfillment procedures, please refer to: https://open.shopee.com/documents/v2/v2.order.get_order_list (자세한 주문 이행 절차는 https://open.shopee.com/documents/v2/v2.order.get_order_list를 참조하십시오.)

---

## 4. Order income (주문 수입)

1. `sa_payment_list_income_API` will return the order amount of A shop currency and the corresponding P shop currency at the same time. (`sa_payment_list_income_API`는 A 상점 통화의 주문 금액과 해당 P 상점 통화를 동시에 반환합니다.)

Please focus on the parameters of `sa_payment_list_income_detail` API as below: (다음과 같이 `sa_payment_list_income_detail` API의 매개변수에 집중하십시오.)

- `int_escrow_amount` represent the A shop order, so it is A shop currency. (`int_escrow_amount`는 A 상점 주문을 나타내므로 A 상점 통화입니다.)
- `int_self` will return this parameter. (`int_self`는 이 매개변수를 반환합니다.)
- `int_self` don't return this parameter. (`int_self`는 이 매개변수를 반환하지 않습니다.)

### Fields (필드)

| Field (필드) | SIP P shop order (SIP P 상점 주문) |
|-------|-----------------|
| escrow_amount | ✓ |
| escrow_amount/final_currency: This field shows the final_currency/local_currency. If the seller_rebate, tax_code, buyer_tax, commission, total_released, total_escrow_amount, seller_voucher, buyer_service, int_buyer_service, commission, tax_fee_payable, pay, sub_affiliate, int_sub_affiliate, original_escrow_amount_coin, original_escrow_amount_from_shipping_fee/local_x | ✓ |
| buyer_total_amount | ✓ |
| actual_shipping_fee | ✓ |
| buyer_paid_shipping_fee | ✓ |
| buyer_transaction_fee | ✓ |
| estimated_shipping_fee | ✓ |
| cod/agent_fee | ✓ |
| credit | ✓ |
| coins_transfer_tax | ✓ |
| escrow_tax | ✓ |
| final_product_protection | ✓ |
| final_product_int_tax | ✓ |
| final_shipping_fee | ✓ |
| final_shipping_int_tax | ✓ |
| seller_transaction_fee | ✓ |
| seller_rebate/seller_weight | ✓ |
| original_cost_of_goods_sold | ✓ |
| original_shopee_discount | ✓ |
| payment_promotion | ✓ |
| reverse_shipping_fee | ✓ |
| rtr_seller_protection_fee_basic_amount | ✓ |
| rtr_seller_protection_fee_premium_amount | ✓ |
| seller_coin_cash_back | ✓ |
| seller_discount | ✓ |
| seller_lost_compensation | ✓ |
| seller_shipping_discount | ✓ |
| seller_transaction_fee | ✓ |
| shopping_fee_discount_e/from_ipi | ✓ |
| shopee_service_fee | ✓ |
| shopee_shipping_rebate | ✓ |
| voucher_from_seller | ✓ |
| voucher_from_shopee | ✓ |
| int_of_currency | ✓ |
| commission_fee_int | ✓ |
| prv_allowable_refund_prt | ✓ |

---

# 5. API permission (API 권한)

Currently, SIP A shop can only call some APIs. please check the detailed list. (현재 SIP A 상점은 일부 API만 호출할 수 있습니다. 자세한 목록을 확인하십시오.)

---

## Available Fields (사용 가능한 필드)

| Field (필드) | Status (상태) |
|-------|--------|
| voucher_from_seller | ✓ |
| voucher_from_shopee | ✓ |
| aff_currency | × |
| commission_fee_pri | × |
| drc_adjustable_refund_pri | × |
| escrow_amount_pri | × |
| original_price_pri | × |
| refund_amount_to_buyer_pri | × |
| seller_return_refund_pri | × |
| service_fee_pri | × |
| sip_subsidy_pri | × |
| pri_currency | × |
| sip_subsidy | × |

**Legend (범례):**
- ✓ = Available/Permitted (사용 가능/허용됨)
- × = Not available/Not permitted (사용 불가/허용되지 않음)

## 사용 사례

1. 판매자를 위한 SIP 상점 인증.
2. 여러 상점 간의 제품 정보 관리.
3. P-shop과 A-shop 간의 제품 가격 동기화.
4. SIP API 액세스를 위한 액세스 토큰 갱신.

## 관련 API

- sip_code_exchange_access_API
- sip_code_refresh_access_API
- sip_code_cnsc_exchange_access_API
- sip_cnsc_list_cnsc_shop_API
- sa_currency
- sa_product_per_service_API

---

## 원문 (English)

### Summary

This guide outlines the best practices for using the SIP (Shop in Place) API, covering authorization, product management, and price logic. It details the authorization process for both non-CNSC and CNSC sellers, along with instructions for obtaining tokens and managing shop relationships.

### Content

# API Best Practices - SIP best practices

**delivery time approximately 2 weeks**

---

## Terminology

- **Primary shop**: In the SIP (open operation mode, the shop (temporarily) operated by the seller is also called the primary shop.
- **Local SIP**: If local sellers operated the primary shop, we call those SIP mode.
- **CNSC SIP**: If CNSC sellers operated the SIP shop (apartment) also refers to the assistant who will be the seller (or the affiliated merchant) and set up SIP Shop. In fact, the SIP Shop (local mode) here and the local SIP mode are essentially the same in terms of entity, and there are no obvious boundaries, except that the seller and the seller-operated shops are placed on the root SIP and SIP apartment user.
- **SIP shop product**: refers to the goods sold by the SIP shop, and the shop product within the shop based on the last year end the product and the SIP price (apartment) user. For non-SIP mode, the cost price of the product and the last year-end the bargain price of the base is the cost price of the shop product.

**Note:** To prevent consumer confusion, whether dealing with shop products or CNSC products, the display of shop products will always remain on the same status as the product.

---

## 1.SIP Shop Authorization

### 1.1 non CNSC Seller

#### 1.1.1 Authorization process

When you successfully create a local SIP shop, you will automatically become a seller and obtain shop authentication. When the seller enters the SIP Shop account, you need to wait for the CNSC official to review. After the review is passed, you will receive the SIP authorization page and clicks to confirm the authorization. SIP If Shop will be authorized this API.

### Authorization

**Local**

[Authorization flow diagram showing numbered steps 1-4]

**Login**
Seller CNSC account or System Store

Apply for authorization to get the following information:
- **Product**: Authorization of product and content management for products period
- **Payment**: Information of transactions and refunds to your shop account
- **Marketing**: Information related to discounts in your shops
- **Logistics**: Information related to order logistics

[AUTHORIZE button shown in orange]

#### 1.1.2 Getting code

After the authorization is successful, you can directly get the shop_id and if shop use the available code through the callback URL.

You can call `sip_code_exchange_access_API` and upload the shop id and code (API to get the shop access token and refresh token) of IT shop.

#### 1.1.3 Refresh Access Token

At this time, the access token and refresh token obtained in the front have can be used in IT shop. Then you can call `sip_code_refresh_access_API` and refresh the access token and refresh token of IT shop.

**Things to note**:
- After the new access_token is generated, the old access_token is still valid within 5 minutes
- If the new refresh_token is generated at the same time, the old refresh_token will become invalid immediately. It must be saved immediately after the new refresh_token is generated, and be used later to refresh.
- If the new refresh_token and access_token renewed are lost, please check FAQ.

---

### 1.2 CNSC Seller

#### 1.2.1 Authorization process

Please refer to single "Authorization and Authentication" When the seller logs in to the main account and enters the authorization page, they can see all the primary SIP shops they have created, and they can also add "Local" (when the SIP IT Shop not click "Confirm Authorization Status, Check the SIP IT Shop and click "Confirm" Authorization Status. When the SIP IT Shop is under the IT shop system that IT shop). After authorization, the access token and refresh token of SIP IT shop will be automatically generated.

[Screenshot showing shop authorization interface with:
- Goshh Seller
- Authorized Shops List
- Affiliate Shop List (ID)
- Shop information fields including:
  - Name (shop name, display name, promotions, show details must fill)
  - Shop Postcode
  - Shop Address
  - Shop Telephone
  - Description
  - Contacts (candAllo2)
  - LegalAddress
  - SortOrder
  - And various shop IDs and status indicators]

**Note**: At this time, the account bound to a merchant is a main account (let be bound under the "Unregistered Shop Group" list. It does in the [Unregistered Shop Group] module, it will show all the shops that merchants have). This status will be marked under it.

[Additional interface screenshots showing:
- Unassigned007 shop
- Unregistered Shop Group - Suspended
- Multiple shop entries with status indicators
- Audit Managed and Reset at SPS markers]

#### 1.2.2 Getting code

After the authorization is successful, you can directly get the main_account_id and the available code through the callback URL.

#### 1.2.3 Getting token

You can call `sip_code_cnsc_exchange_access_API` and upload the main_account_id and code (API to get the currently authorized merchants, and the sicp_id_list will return all the currently authorized merchants).

**Note**: The main_account_id of the currently authorized merchant account access token refresh_token of IT shop.

#### 1.2.4 Getting shop relationship

After you get the main_account_id and the sicp_id_list by the mentioned API, to get the list of authorized shops associated with this merchant, and set up `sip_cnsc_list_cnsc_shop_API`. Then you will get information like 1)(Under the shop belongs to merchant list. HTML's shop list is under this shop list, and the sicp_id_list should be in that list) of authorized shops associated with each list.

**Note**:
- If it is an normal shop, (i.e., non-does not the sip_id), shops level will not be returned.
- If it is a SIP shop managed by CNSC (that is, the sip_id), shops level will be returned and ask for the sicp_value.
- If it is a SIP shop managed by merchant (that is, the sip_id), shops level will be returned but the sicp_value will be null.

#### 1.2.5 Refresh Access Token

Please refer to `sip_code_refresh_access_API` to refresh the access token and refresh token of each sicp_id when it about to expired. When you call `sip_code_refresh_access_API`, please refresh the access token and refresh token of each sicp_id (of each merchant) one by one.

---

## 2. Product Management

### 2.1 Product information logic

**Info**

**Sync logic**

**Item types info (item)**
- Creator: After the seller creates a IT shop product, Shopee will automatically transcode

---

*[End of visible content]*

---

# 2. Product Management

## 2.1 Product information logic

### Info

**Sync logic**

**Item base info (Item base + category)**
- **Create**: After the seller creates a P-Help product, Shoppe will automatically translate the product information to the target shop language via P-Help API.
- **Update**: If the seller modifies the P-Help product, Shoppe will automatically synchronize the modifications to the target shop.

**Item status**
- If the P-Help product is unlisted or deleted, the A shop product will also be unlisted or deleted.

**Item stock**
- The product stock of P-Help is equal to the product stock of A shop. For example, there are 2 A shops under the seller, each add 4 of buyer, and the stock of P-Help (P-Shop) will be reduced by 8. If the seller wants to modify the stock, they can only modify the P-Help product stock in A shop.

**Item price**
- When the seller updates the product price of P-shop, Shoppe will synchronize it to the product price in A shop.

---

## 2.2 Price logic

Please note that all APIs using `sa_currency` and `sa_product_per_service` API Response: `currency: "MYR"`, `buyer_shipping_fee: "119.90"`, `original_price: "129.90"`, `tag_discounted_price: "119.90"`.

**Note**: If the item is a P shop item and there are currently subsidiaries, please call the `sa_product_per_service_API` to get the sa_currency, item_price, currency of each.

- If the item is a P shop item, please call the `sa_product_per_service_API` to get the sa_currency, item_price, currency of each subsidiary store.
- If the item is a sub "B" shop item but not `sa_get_pricing_item_price` separately, item_price_currency fields will not return in v2 API response.

---

## 3. Order management

### 3.1 Order synchronization logic

#### Steps

Please follow the SIP P shop order list and the SIP A shop order list respectively for fulfillment or cancellation. Please check seller cancellation list to the item cannot of the seller is no SIP A shop under synchronization rules or the seller want to cancel the order, or when the buyer order times out.

#### 3.1.1 2 Local SIP

Shoppe automatically synchronizes the SIP A shop order to the SIP P shop, so you only need to obtain and fulfill the SIP P shop order list. The seller only needs to operate in the SIP P shop. Shoppe will automatically synchronize it to the SIP A shop. Examples include a Local SG Shopper buys/receives warehouses from Seller P-Shop Singapore to Local SG Shopper buys/receives from Seller A-Shop Singapore.

**Note A**

For the same order merchandise, The orders of P shops and A shops will be pushed at the same time. If no SIP sellers, the order will only be pushed to the p shop. So sellers only need to pay attention to the P shop orders, they only need to pay attention to the p shop orders when the seller order list.

**Note B**

For more order fulfillment procedures, please refer to: https://open.shopee.com/documents/v2/v2.order.get_order_list

---

## 4. Order income

1. `sa_payment_list_income_API` will return the order amount of A shop currency and the corresponding P shop currency at the same time.

Please focus on the parameters of `sa_payment_list_income_detail` API as below:

- `int_escrow_amount` represent the A shop order, so it is A shop currency.
- `int_self` will return this parameter.
- `int_self` don't return this parameter.

### Fields

| Field | SIP P shop order |
|-------|-----------------|
| escrow_amount | ✓ |
| escrow_amount/final_currency: This field shows the final_currency/local_currency. If the seller_rebate, tax_code, buyer_tax, commission, total_released, total_escrow_amount, seller_voucher, buyer_service, int_buyer_service, commission, tax_fee_payable, pay, sub_affiliate, int_sub_affiliate, original_escrow_amount_coin, original_escrow_amount_from_shipping_fee/local_x | ✓ |
| buyer_total_amount | ✓ |
| actual_shipping_fee | ✓ |
| buyer_paid_shipping_fee | ✓ |
| buyer_transaction_fee | ✓ |
| estimated_shipping_fee | ✓ |
| cod/agent_fee | ✓ |
| credit | ✓ |
| coins_transfer_tax | ✓ |
| escrow_tax | ✓ |
| final_product_protection | ✓ |
| final_product_int_tax | ✓ |
| final_shipping_fee | ✓ |
| final_shipping_int_tax | ✓ |
| seller_transaction_fee | ✓ |
| seller_rebate/seller_weight | ✓ |
| original_cost_of_goods_sold | ✓ |
| original_shopee_discount | ✓ |
| payment_promotion | ✓ |
| reverse_shipping_fee | ✓ |
| rtr_seller_protection_fee_basic_amount | ✓ |
| rtr_seller_protection_fee_premium_amount | ✓ |
| seller_coin_cash_back | ✓ |
| seller_discount | ✓ |
| seller_lost_compensation | ✓ |
| seller_shipping_discount | ✓ |
| seller_transaction_fee | ✓ |
| shopping_fee_discount_e/from_ipi | ✓ |
| shopee_service_fee | ✓ |
| shopee_shipping_rebate | ✓ |
| voucher_from_seller | ✓ |
| voucher_from_shopee | ✓ |
| int_of_currency | ✓ |
| commission_fee_int | ✓ |
| prv_allowable_refund_prt | ✓ |

---

# 5. API permission

Currently, SIP A shop can only call some APIs. please check the detailed list.

---

## Available Fields

| Field | Status |
|-------|--------|
| voucher_from_seller | ✓ |
| voucher_from_shopee | ✓ |
| aff_currency | × |
| commission_fee_pri | × |
| drc_adjustable_refund_pri | × |
| escrow_amount_pri | × |
| original_price_pri | × |
| refund_amount_to_buyer_pri | × |
| seller_return_refund_pri | × |
| service_fee_pri | × |
| sip_subsidy_pri | × |
| pri_currency | × |
| sip_subsidy | × |

**Legend:**
- ✓ = Available/Permitted
- × = Not available/Not permitted

---

**문서 ID**: developer-guide.261
**플랫폼**: shopee
**URL**: https://open.shopee.com/developer-guide/261
**처리 완료**: 2025-10-16T08:57:30
