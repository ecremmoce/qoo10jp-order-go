# API 모범 사례 - 상품 기본 정보 관리

**카테고리**: 모범 사례
**난이도**: medium
**중요도**: 5/5
**최종 업데이트**: 2025-10-16T08:46:41

## 개요

본 가이드는 Shopee API를 사용하여 상품 기본 정보를 관리하는 모범 사례를 설명합니다. 여기에는 글로벌 상품 및 사이즈 차트를 포함하여 상품 데이터 검색, 검색, 업데이트 및 삭제와 같은 다양한 작업이 포함됩니다.

## 주요 키워드

- 상품 정보
- API
- 상품 관리
- 글로벌 상품
- 사이즈 차트
- 상품 프로모션
- 상품 업데이트
- 상품 삭제

## 본문

# API 모범 사례 > 상품 기본 정보 관리

이 API를 사용하면 상점의 모든 상품 목록을 가져오거나 update_time 범위 및 상품 상태별로 필터링할 수 있습니다.

## 1. item_id 검색
API: `v2.product.search_item`

이 API를 사용하면 다음과 같은 특정 조건을 기반으로 item_id 목록을 검색할 수 있습니다.
- 상품 이름 키워드를 포함하는 item_id 목록
- SKU 키워드를 포함하는 item_id 목록
- 문의 속성을 추적하는 item_id 목록
- 선택적 속성을 추적하는 item_id 목록

---

## 2. 상품 정보 가져오기
API: `v2.product.get_item_base_info`, `v2.product.get_model_list`

`v2.product.get_item_base_info`를 호출하여 상품 기본 정보를 가져올 수 있습니다. 그렇지 않으면 `v2.product.get_model_list` API를 호출하여 variant 가격 및 재고를 가져와야 합니다.

`v2.product.get_item_base_info` API의 "has_model" 필드는 상품에 variant가 있는지 여부를 나타냅니다.

---

## 3. 상품 데이터 가져오기
API: `v2.product.get_item_extra_info`

이 API는 상품의 조회수, 좋아요, 판매량, 평점 및 별점을 가져올 수 있습니다.

조회수 데이터는 지난 90일간의 통계에서 가져온 것이며, 판매 데이터는 누적 값입니다.

---

## 4. 상품 프로모션 정보 가져오기
API: `v2.product.get_item_promotion`

이 API를 사용하면 상품이 추가된 모든 진행 중이거나 예정된 프로모션에 대한 정보를 가져올 수 있습니다. 상품이 여러 프로모션에 추가된 경우 `v2.product.get_item_promotion`의 promotion_id 필드는 둘 이상의 promotion_id를 반환하며, 모든 권한 정보를 얻기 위해 `v2.product.get_item_promotion`을 계속 호출하는 것이 좋습니다.

---

## 5. 상품 정보 업데이트
API: `v2.product.update_item`

1. 이 API는 size_chart/price/stock/video 정보를 제외한 상품 정보 업데이트를 지원합니다. 업데이트 가능한 필드는 업데이트되고, 업데이트 불가능한 필드는 업데이트되지 않습니다.

2. item_sku/wholesales/video_upload_id의 경우 삭제 작업을 지원합니다. null 문자열을 업로드하면 삭제됩니다.

item_sku 삭제 예시
```
"item_sku":6001B2459,
"item_sku":""
```

3. 관련 설명 업데이트에 대한 FAQ를 참조하십시오.

4. 일부 필드를 업데이트하지 않았지만 해당 필드가 잘못 채워졌다는 메시지가 표시되는 경우, 업데이트할 때마다 모든 상품 정보의 적법성을 확인하므로 요구 사항을 충족하지 않으면 수정하십시오.

---

## 6. 상품 판매 중지 또는 삭제
API: `v2.product.unlist_item`, `v2.product.delete_item`

"unlist" - true는 상품이 판매 중지됨을 의미하고, "unlist" - false는 상품이 다시 판매됨을 의미합니다.

이 API는 item_status를 "deleted"로 변경할 수 있습니다. 삭제 후에는 상품을 업데이트할 수 없으며 판매자는 판매자 센터를 통해 이 상품을 볼 수 없습니다.

Shopee에서 삭제된 상품 및 판매자가 삭제한 상품의 경우 90일 이내에는 API를 통해 상품 정보를 계속 가져올 수 있으며, 90일 후에는 상품 데이터가 Shopee 데이터베이스에서 영구적으로 삭제되어 이 상품에 대한 정보를 쿼리할 수 없습니다. 필요한 경우 상품 정보를 제때 저장하십시오.

---

## 7. 사이즈 차트 이미지 업데이트
API: `v2.product.update_size_chart`

이 API는 상품의 이미지 사이즈 차트를 추가하거나 업데이트하는 데 사용할 수 있습니다. "Your shop can not edit image size chart" 오류가 발생하면 공지 사항을 확인하십시오.

---

## 8. 브랜드 등록
API: `v2.product.register_brand`

판매자는 이 API를 통해 자신의 브랜드를 등록할 수 있습니다. Shopee가 이 브랜드를 성공적으로 감사하면 상품을 추가하거나 업데이트하기 위한 유효한 brand_id를 얻게 됩니다.

브랜드 등록 상태는 FAQ에서 확인할 수 있습니다.

*다음 내용은 CN/SG/PH/SG 판매자에게만 적용됩니다.

---

## 9. 글로벌 상품 목록 가져오기
API: `v2.global_product.get_global_item_list`

이 API를 사용하면 판매자 아래의 모든 global_item_id 목록을 가져오거나 update_time 범위별로 필터링할 수 있습니다. 이 API는 삭제된 global_item_id 목록을 반환하지 않습니다.

---

## 10. 글로벌 상품 ID 가져오기
API: `v2.global_product.get_global_item_id`

이 API를 호출하면 상점 상품의 global_item_id를 빠르게 찾을 수 있습니다.

---

## 11. 글로벌 상품 정보 가져오기
API: `v2.global_product.get_global_item_info`, `v2.global_product.get_global_model_list`

1. 글로벌 상품에 variant가 포함되어 있지 않으면 `v2.global_product.get_global_item_info`만 호출하여 글로벌 상품 정보를 가져오면 됩니다. 그렇지 않으면 `v2.global_product.get_global_model_list`를 호출하여 variant 재고 및 가격 정보를 가져와야 합니다.

2. global_product 관련 인터페이스는 글로벌 상품에서 제공되지 않습니다.

`v2.product.get_item_extra_info`를 호출하여 상점 상품 데이터를 가져오거나 `v2.product.get_item_promotion`을 호출하여 상점 상품의 프로모션 데이터를 가져올 수 있습니다.

---

## 12. 글로벌 상품 업데이트
API: `v2.global_product.update_global_item`

1. 일부 필드는 글로벌 상품과 상점 상품에서 함께 관리할 수 있으므로 "Updateable or not-updateable fields" 문서를 확인하여 글로벌 상품 수준에서 상점 상품 수준으로 동기화되는 필드를 볼 수 있습니다. API가 나열되고 v2_global_item API를 통해 동기화되면 글로벌 상품을 업데이트한 후 Shopee가 자동으로 상점 상품에 동기화합니다.

---

## 13. 글로벌 상품 삭제
API: `v2.global_product.delete_global_item`

글로벌 상품은 "판매 중지"를 지원하지 않지만 삭제할 수 있습니다. 글로벌 상품이 삭제되면 게시된 모든 상점 상품도 삭제됩니다.

---

## 14. 사이즈 차트 이미지 업데이트
API: `v2.global_product.update_size_chart`

이 API를 사용하면 글로벌 상품의 사이즈 차트를 추가하거나 업데이트할 수 있습니다. 글로벌 상품을 통해 상점 상품의 사이즈 차트를 업데이트하려면 `v2.product.update_size_chart` API를 호출하면 됩니다.

CN/SG/PH/SG 판매자가 개별적으로 보유한 상품 모듈 인터페이스 권한에 대해 자세히 알아보려면 이 FAQ를 확인하십시오.

## 사용 사례

1. 판매자 대시보드에 표시하기 위한 상품 세부 정보 검색.
2. 이름, 설명 또는 카테고리와 같은 상품 정보 업데이트.
3. 상품 프로모션 및 할인 관리.
4. 판매자 시스템과 Shopee 간의 상품 데이터 동기화.
5. Shopee에 브랜드 등록.

## 관련 API

- v2.product.search_item
- v2.product.get_item_base_info
- v2.product.get_model_list
- v2.product.get_item_extra_info
- v2.product.get_item_promotion
- v2.product.update_item
- v2.product.unlist_item
- v2.product.delete_item
- v2.product.update_size_chart
- v2.product.register_brand
- v2.global_product.get_global_item_list
- v2.global_product.get_global_item_id
- v2.global_product.get_global_item_info
- v2.global_product.get_global_model_list
- v2.global_product.update_global_item
- v2.global_product.delete_global_item

---

## 원문 (English)

### Summary

This guide outlines best practices for managing product base information using the Shopee API. It covers various operations such as searching, retrieving, updating, and deleting product data, including global products and size charts.

### Content

# API Best Practices > Product base info manage

This API allows you to get a list of all the products in the shop or filter by update_time range and item status

## 1. Searching for item_id
API: `v2.product.search_item`

This API allows you to search for a list of item_id based on some specific conditions, including
- A list of item_id containing the product name keyword
- A list of item_id containing the sku keyword
- A list of item_id tracking inquiries attributes
- A list of item_id tracking optional attributes

---

## 2. Getting product information
API: `v2.product.get_item_base_info`, `v2.product.get_model_list`

You can call `v2.product.get_item_base_info` to get the product base information, otherwise you also need to call `v2.product.get_model_list` API to get the variants price and stock

The field "has_model" in `v2.product.get_item_base_info` API indicates whether the product has variants or not.

---

## 3. Getting the data of product
API: `v2.product.get_item_extra_info`

This API can get the data of views, likes, sales, ratings, and star rating from a product.

The date of views is from the last 90 days' statistics, the sales data is the cumulative value.

---

## 4. Getting product promotion information
API: `v2.product.get_item_promotion`

This API allows you to get information about all ongoing or upcoming promotions that the product is added in. If the product is added into multiple promotions, the promotion_id field of `v2.product.get_item_promotion` will return more than one promotion_id, and we suggest you to continue to call `v2.product.get_item_promotion` to get all the permissions information.

---

## 5. Updating product information
API: `v2.product.update_item`

1. This API supports updating product information except for the size_chart/price/stock/video information. Fields that are updateable will be updated, and fields that are not updateable will not be updated.

2. For item_sku/wholesales/video_upload_id, we support the delete operation. If you upload the null string then we will delete it.

Example of deleting item_sku
```
"item_sku":6001B2459,
"item_sku":""
```

3. Please refer to the FAQ about updating related descriptions.

4. If you do not update some fields but encountered a prompt that those fields are filled in incorrectly, this situation is normal because every time you update, we will verify the legitimacy of all the product information, so if it does not meet the requirements, please modify it.

---

## 6. Unlisting or deleting product
API: `v2.product.unlist_item`, `v2.product.delete_item`

"unlist" - true means the product will be unlist,"unlist" - false, means the product will be re-listed

This API can change item_status to be "deleted", please note that after the deletion, you will not be able to update the product and the seller can not view this item through the Seller Center.

For Shopee deleted and Seller deleted products, within 90 days, you can still get the product information through API, after 90 days, the product data will be permanently deleted in Shopee database, you can not query any information about this product. If you need, please save the product information in time.

---

## 7. Updating size chart image
API: `v2.product.update_size_chart`

This API can be used to add or update the image size chart of the product. If you encounter the error "Your shop can not edit image size chart", please check the announcement.

---

## 8. Registering Brand
API: `v2.product.register_brand`

Sellers can register their own brands through this API. If Shopee audits this brand successfully, you will get a valid brand_id for adding or updating products.

You can check the brand registration status check the FAQ

*The following content is only applicable to CN/SG/PH/SG sellers.

---

## 9. Getting global product list
API: `v2.global_product.get_global_item_list`

This API allows you to get a list of all global_item_id or filter by update_time range under the merchant. This API will not return a list of deleted global_item_id

---

## 10. Getting a global product ID
API: `v2.global_product.get_global_item_id`

By calling this API, you can quickly find the global_item_id of a shop product

---

## 11. Getting global product information
API: `v2.global_product.get_global_item_info`, `v2.global_product.get_global_model_list`

1. If the global product does not contain variants, you only need to call `v2.global_product.get_global_item_info` to get the global product information, otherwise you also need to call `v2.global_product.get_global_model_list` to get the variants stock and price information

2. global_product-related interfaces are not served on global products

You can call `v2.product.get_item_extra_info` to get the shop product data or call `v2.product.get_item_promotion` to get the promotion data of shop products.

---

## 12. Updating global products
API: `v2.global_product.update_global_item`

1. Since some fields can be managed by global products and shop products together, you can check the article "Updateable or not-updateable fields" to view the fields synchronized by the global product level to the shop product level. The API will be listed, synchronized through v2_global_item APIs, both API, then Shopee will automatically synchronize to shop products after you update the global product

---

## 13. Deleting global product
API: `v2.global_product.delete_global_item`

Global products do not support "unlist", but can be deleted. After the global products are deleted, all published shop products will also be deleted

---

## 14. Updating the size chart image
API: `v2.global_product.update_size_chart`

This API allows you to add or update the size chart of global products. If you want to update the size chart of a shop product through a global product, you can call the `v2.product.update_size_chart` API.

You can check this FAQ to learn more about which product module interface permissions CN/SG/PH/SG sellers have individually.

---

**문서 ID**: developer-guide.221
**플랫폼**: shopee
**URL**: https://open.shopee.com/developer-guide/221
**처리 완료**: 2025-10-16T08:46:41
