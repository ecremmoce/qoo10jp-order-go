# 글로벌 상품 게시

**카테고리**: 모범 사례
**난이도**: 중간
**중요도**: 4/5
**최종 업데이트**: 2025-10-16T08:43:14

## 개요

본 가이드는 게시 가능한 상점 목록 가져오기, 상점 채널, 상품 게시, 게시된 상점 상품 목록 검색을 포함하여 글로벌 상품을 게시하는 단계를 설명합니다. 또한 Shopee가 선택적 필드와 제휴 상점에 대한 자동 게시를 처리하는 방법에 대해서도 자세히 설명합니다.

## 주요 키워드

- 글로벌 상품
- 게시
- 상점 인증
- 상점 채널
- SIP
- API
- 번역
- 가격
- 이미지

## 본문

# 글로벌 상품 게시

## 탐색
- 공지사항
- 콘솔

**Breadcrumb:** API 모범 사례 > 글로벌 상품 게시

---

## 1단계: 게시 가능한 상점 목록 가져오기

**API:** `v2.global_product.get_publishable_shop`

이 API를 통해 글로벌 상품에 대해 게시 가능한 상점 목록을 가져올 수 있지만, 다음과 같은 경우에는 해당 상점을 반환하지 않습니다.

- 상점 인증을 완료하지 않은 상점
- SIP 제휴 상점
- 게시된 마켓의 상점

---

## 2단계: 상점 채널 가져오기

**API:** `v2.logistics.get_channel_list`

`v2.logistics.get_channel_list` API를 호출하여 enable=true 및 mask_channel_id=0인 채널을 상점 상품에 대해 선택해야 합니다. 상점 채널을 업로드하지 않고 글로벌 상품을 게시하는 경우, 활성화되어 사용 가능한 상점 채널을 기본적으로 선택합니다.

---

## 3단계: 글로벌 상품 게시

**API:** `v2.global_product.create_publish_task`

`v2.global_product.create_publish_task`의 선택적 필드의 경우, 업로드하지 않으면 Shopee에서 일부 처리를 수행하여 상점 상품에 업로드합니다. 업로드하면 사용자 정의 값이 되며 Shopee는 어떠한 처리도 수행하지 않습니다. 구체적인 로직은 다음과 같습니다.

| 필드 이름 | 사용자 정의 값 업로드 |
|------------|------------------------|
| item_name | 예: 판매자는 현지 언어로 번역하여 업로드해야 합니다.<br>아니요: Shopee에서 현지 언어로 번역하는 데 도움을 줍니다. |
| description_info/description | 예: 판매자는 현지 언어로 번역하여 업로드해야 합니다.<br>아니요: Shopee에서 현지 언어로 번역하는 데 도움을 줍니다. |
| original_price | 예: 판매자는 현지 통화로 가격을 업로드해야 합니다.<br>아니요: Shopee는 글로벌 상품 가격 및 계산 공식을 기반으로 현지 가격을 계산합니다. |
| tier_variation--name | 예: 판매자는 현지 언어로 번역하여 업로드해야 합니다.<br>아니요: Shopee에서 현지 언어로 번역하는 데 도움을 줍니다. |
| tier_variation--option_list--option | 예: 판매자는 현지 언어로 번역하여 업로드해야 합니다.<br>아니요: Shopee에서 현지 언어로 번역하는 데 도움을 줍니다. |
| image | 예: 판매자가 사용자 정의 이미지를 업로드합니다.<br>아니요: Shopee에서 글로벌 상품의 이미지를 복사합니다. |

API 호출이 성공하면 publish_task_id를 얻게 됩니다.

**API:** `v2.global_product.get_publish_task_result`

이 API는 게시 작업이 성공했는지 여부를 반환합니다. 성공하면 item_id, shop_id 및 지역 정보를 반환하고, 실패하면 실패에 대한 구체적인 이유를 반환합니다.

참고: 게시된 상점이 SIP 기본 상점인 경우, 성공적으로 게시한 후 Shopee는 SIP 기본 상점 산하의 제휴 상점에 글로벌 상품을 자동으로 게시합니다.

---

## 4단계: 게시된 상점 상품 목록 가져오기

**API:** `v2.global_product.get_published_list`

API는 Shopee가 제휴 상점에 자동으로 게시하는 상점 상품을 포함하여 이 글로벌 상품에 대해 성공적으로 게시된 모든 상점의 item_id 및 shop_id를 반환합니다.

참고: 이 API는 게시되었지만 상점 인증을 완료하지 않은 상점 상품은 반환하지 않습니다.

## 사용 사례

1. 여러 지역으로 제품 도달 범위 확장
2. 여러 상점에서 제품 목록 관리
3. 제품 게시 워크플로 자동화
4. Shopee의 글로벌 상품 기능과 통합

## 관련 API

- v2.global_product.get_publishable_shop
- v2.logistics.get_channel_list
- v2.global_product.create_publish_task
- v2.global_product.get_publish_task_result
- v2.global_product.get_published_list

---

## 원문 (English)

### Summary

This guide outlines the steps to publish global products, including getting the list of publishable shops, shop channels, publishing the product, and retrieving the list of published shop products. It also details how Shopee handles optional fields and automatic publishing to affiliated shops.

### Content

# Publishing Global Product

## Navigation
- Announcement
- Console

**Breadcrumb:** API Best Practices > Publishing global product

---

## Step 1: Getting the list of publishable shops

**API:** `v2.global_product.get_publishable_shop`

You can get the list of publishable shops for global products through this API, but we will not return the corresponding shops for the following cases:

- Shops that have not done the shop authorization
- SIP affiliated shops
- Shops of published market

---

## Step 2: Getting the shop channel

**API:** `v2.logistics.get_channel_list`

You need to call `v2.logistics.get_channel_list` API and select the channel with enable=true and mask_channel_id=0 for shop products. If you publish global products without uploading shop channels, we will choose the enabled and available shop channels for you by default.

---

## Step 3: Publishing global product

**API:** `v2.global_product.create_publish_task`

For the optional fields of `v2.global_product.create_publish_task`, if you do not upload, Shopee will do some processing and upload to shop products. If you upload, it will be a custom value and Shopee will not do any processing. The specific logic is as follows:

| Field Name | Upload the custom value |
|------------|------------------------|
| item_name | YES: Seller need to translate into local language then upload<br>NO: Shopee will help translate into the local language |
| description_info/description | YES: Seller need to translate into local language then upload<br>NO: Shopee will help translate into the local language |
| original_price | YES: Seller need to upload the price in local currency<br>NO: Shopee calculates local prices based on the price of global products and calculation formulas. |
| tier_variation--name | YES: Seller need to translate into local language then upload<br>NO: Shopee will help translate into the local language |
| tier_variation--option_list--option | YES: Seller need to translate into local language then upload<br>NO: Shopee will help translate into the local language |
| image | YES: Seller uploads the custom images<br>NO: Shopee will copy the images of global products |

After a successful API call, you will get a publish_task_id.

**API:** `v2.global_product.get_publish_task_result`

This API will return whether the publish task was successful or not. If it succeeds, it will return the item_id, shop_id, and region information, if it fails, it will return the specific reason for the failure.

Please note: If the published shop is the SIP primary shop, then after the successful publication, Shopee will automatically publish the global product to the affiliated shops under the SIP primary shop.

---

## Step 4: Getting the list of published shop products

**API:** `v2.global_product.get_published_list`

API will return the item_id and shop_id of all the shops that have been successfully published for this global product, including the shop products that Shopee automatically publishes to the affiliated shops.

Please note: This API does not return shop products that have been published but have not done the shop authorization.

---

**문서 ID**: developer-guide.215
**플랫폼**: shopee
**URL**: https://open.shopee.com/developer-guide/215
**처리 완료**: 2025-10-16T08:43:14
