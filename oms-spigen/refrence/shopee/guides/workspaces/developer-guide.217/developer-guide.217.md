# API 호출 흐름 개요

**카테고리**: Integration
**난이도**: medium
**중요도**: 4/5
**최종 업데이트**: 2025-10-16T08:33:12

## 개요

본 가이드는 제품 생성, 글로벌 제품 생성, 글로벌 제품 게시를 위한 API 호출 흐름에 대한 개요를 제공합니다. 각 흐름에 관련된 필수 및 선택적 프로세스를 설명하여 API 통합에 대한 구조화된 접근 방식을 제공합니다.

## 주요 키워드

- API call flow
- 제품 생성
- 글로벌 제품
- 게시
- workflow
- API integration

## 본문

# API 호출 흐름 개요

*굵은 글씨는 필수 프로세스이며, 일반 글씨는 필수가 아닌 프로세스입니다.*

## 1. 제품 생성

**시작**
↓
**initial_user_webhook_user**
↓
**selling_price.list.user_selling_currency** (선택 사항)
↓
**master_attributes.product_type** (선택 사항)
↓
**selling_location.user_selling_location** (선택 사항)
↓
**channel_product.get.linked.product_field** (선택 사항)
↓
**product.get.category**
↓
**product.get.reference**
↓
**product_search_by.id**
↓
**product.create.V2**
↓
**product.image.push**
↓
**product.create.qc.approval** ← **push.product.qc.result**
↓
**종료**

---

## 2. 글로벌 제품 생성

**시작**
↓
**initial_user.webhook_user**
↓
**selling_price.list.user_selling_currency** (선택 사항)
↓
**Temp.logistics.detail.product.Type** (선택 사항)
↓
**selling_location.user_selling_location** (선택 사항)
↓
**master_attributes.get.global_category_id** (선택 사항)
↓
**global_product.get.global_item.list**
↓
**global_product.get.reference**
↓
**Global_product.get.product.id**
↓
**global.V2.create.product**
↓
**global.V2.product.image.push**
↓
**global.product.qc.result.check**
↓
**global.product.qc.approval** ← **push.product.qc.result**
↓
**종료**

---

## 3. 글로벌 제품 게시

**시작**
↓
**global_product.get_publishable_shop**
↓
**logistics.get_channel_list**
↓
**global_product.create_publish_task**
↓
**global_product.get_publish_task_result**
↓
**global_product.get_published_list**
↓
**종료**

---

# 데이터 정의

## 속성 값 데이터 유형
(input_validation_type)
- INT_TYPE
- STRING_TYPE
- ENUM_TYPE
- FLOAT_TYPE
- TIMESTAMP_TYPE
- DATE_TYPE

## 속성 입력 유형
(input_type)
- DROP_DOWN
- TEXT_FILED
- COMBO_BOX
- MULTIPLE_SELECT
- MULTIPLE_SELECT_COMBO_BOX

## 물류 유형
(fee_type)
- SIZE_SELECTION
- SIZE_INPUT
- FIXED_DEFAULT_PRICE
- CUSTOM_PRICE

## 아이템 상태 유형
(item_status)
- NORMAL
- DELETED
- BANNED
- UNLIST

## 번역 언어
(language)
- zh-hans: 중국어 간체
- zh-hant: 중국어 번체
- ms-my: 말레이어
- en-my: 영어 (말레이시아)
- en: 영어
- id: 인도네시아어
- vi: 베트남어
- th: 태국어
- pt-br: 포르투갈어
- es-mx: 스페인어 (멕시코)
- pl: 폴란드어
- es-CO: 스페인어 (콜롬비아)
- es-CL: 스페인어 (칠레)

## 재고 유형
(stock_type)
- 1: Shopee 창고 재고
- 2: 판매자 재고

## 제품 프로모션 유형
(promotion_type)
- Campaign
- Discount Promotions
- Flash Sale
- Whole Sale
- Group Buy
- Bundle Deal
- Welcome Package
- Add-on Discount
- Brand Sale
- In ShopFlash Sale
- Gift with purchase
- Exclusive Price

## 마켓 코드
- SG: 싱가포르
- MY: 말레이시아
- TW: 대만
- ID: 인도네시아
- VN: 베트남
- TH: 태국
- BR: 브라질
- PH: 필리핀
- MX: 멕시코
- CO: 콜롬비아
- CL: 칠레
- PL: 폴란드

## 사용 사례

1. 제품 생성 API와 통합
2. 글로벌 제품 목록 자동화
3. 제품 게시 워크플로우 오케스트레이션
4. 제품 관리를 위한 API 종속성 이해

## 관련 API

- initial_user_webhook_user
- selling_price.list.user_selling_currency
- master_attributes.product_type
- selling_location.user_selling_location
- channel_product.get.linked.product_field
- product.get.category
- product.get.reference
- product_search_by.id
- product.create.V2
- product.image.push
- product.create.qc.approval
- push.product.qc.result
- initial_user.webhook_user
- Temp.logistics.detail.product.Type
- master_attributes.get.global_category_id
- global_product.get.global_item.list
- global_product.get.reference
- Global_product.get.product.id
- global.V2.create.product
- global.V2.product.image.push
- global.product.qc.result.check
- global.product.qc.approval
- global_product.get_publishable_shop
- logistics.get_channel_list
- global_product.create_publish_task
- global_product.get_publish_task_result
- global_product.get_published_list

---

## 원문 (English)

### Summary

This guide provides an overview of the API call flows for creating products, creating global products, and publishing global products. It outlines the required and optional processes involved in each flow, offering a structured approach to API integration.

### Content

# API Call Flow Overview

*Bold text is a required process, General text is not required process*

## 1. Creating Product

**Start**
↓
**initial_user_webhook_user**
↓
**selling_price.list.user_selling_currency** (Optional)
↓
**master_attributes.product_type** (Optional)
↓
**selling_location.user_selling_location** (Optional)
↓
**channel_product.get.linked.product_field** (Optional)
↓
**product.get.category**
↓
**product.get.reference**
↓
**product_search_by.id**
↓
**product.create.V2**
↓
**product.image.push**
↓
**product.create.qc.approval** ← **push.product.qc.result**
↓
**End**

---

## 2. Creating Global Product

**Start**
↓
**initial_user.webhook_user**
↓
**selling_price.list.user_selling_currency** (Optional)
↓
**Temp.logistics.detail.product.Type** (Optional)
↓
**selling_location.user_selling_location** (Optional)
↓
**master_attributes.get.global_category_id** (Optional)
↓
**global_product.get.global_item.list**
↓
**global_product.get.reference**
↓
**Global_product.get.product.id**
↓
**global.V2.create.product**
↓
**global.V2.product.image.push**
↓
**global.product.qc.result.check**
↓
**global.product.qc.approval** ← **push.product.qc.result**
↓
**End**

---

## 3. Publishing Global Product

**Start**
↓
**global_product.get_publishable_shop**
↓
**logistics.get_channel_list**
↓
**global_product.create_publish_task**
↓
**global_product.get_publish_task_result**
↓
**global_product.get_published_list**
↓
**End**

---

# Data Definition

## Attribute value data type
(input_validation_type)
- INT_TYPE
- STRING_TYPE
- ENUM_TYPE
- FLOAT_TYPE
- TIMESTAMP_TYPE
- DATE_TYPE

## Attribute input type
(input_type)
- DROP_DOWN
- TEXT_FILED
- COMBO_BOX
- MULTIPLE_SELECT
- MULTIPLE_SELECT_COMBO_BOX

## Logistics type
(fee_type)
- SIZE_SELECTION
- SIZE_INPUT
- FIXED_DEFAULT_PRICE
- CUSTOM_PRICE

## Item status type
(item_status)
- NORMAL
- DELETED
- BANNED
- UNLIST

## Translation language
(language)
- zh-hans: Simplified Chinese
- zh-hant: Traditional Chinese
- ms-my: Malay
- en-my: English (Malaysia)
- en: English
- id: Indonesian
- vi: Vietnamese
- th: Thai
- pt-br: Portuguese
- es-mx: Spanish (Mexican)
- pl: Polish
- es-CO: Spanish (Colombia)
- es-CL: Spanish (Chile)

## Stock type
(stock_type)
- 1: Shopee Warehouse stock
- 2: Seller stock

## Product promotion type
(promotion_type)
- Campaign
- Discount Promotions
- Flash Sale
- Whole Sale
- Group Buy
- Bundle Deal
- Welcome Package
- Add-on Discount
- Brand Sale
- In ShopFlash Sale
- Gift with purchase
- Exclusive Price

## Market Code
- SG: Singapore
- MY: Malaysia
- TW: Taiwan
- ID: Indonesia
- VN: Vietnam
- TH: Thailand
- BR: Brazil
- PH: Philippines
- MX: Mexico
- CO: Colombia
- CL: Chile
- PL: Poland

---

**문서 ID**: developer-guide.217
**플랫폼**: shopee
**URL**: https://open.shopee.com/developer-guide/217
**처리 완료**: 2025-10-16T08:33:12
