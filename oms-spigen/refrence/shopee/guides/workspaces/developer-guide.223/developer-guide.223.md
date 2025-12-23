# 재고 & 가격 관리

**카테고리**: API 관리
**난이도**: medium
**중요도**: 4/5
**최종 업데이트**: 2025-10-16T08:47:59

## 개요

본 가이드는 제품 가격 및 재고 수준 업데이트를 포함하여 제품의 재고 및 가격을 관리하는 방법에 대한 정보를 제공합니다. 여기에는 배리에이션이 있는 제품과 없는 제품, 그리고 국경 간 판매자를 위한 글로벌 제품과 같은 다양한 시나리오가 포함됩니다.

## 주요 키워드

- 재고 관리
- 가격 관리
- 제품 배리에이션
- 글로벌 제품
- API
- 국경 간 판매자
- 판매자 재고
- 재판매 재고

## 본문

# 재고 & 가격 관리

## API 응답

```json
"price_API": {
    "seller_price": 500,
    "original_price": 1000,
    "release_price_of_current_price": 937,
    "release_price_of_original_price": 1663,
    "customer": "COP"
}
```

### 참고 사항
- 제품에 진행 중인 프로모션이 있는 경우 'current_price'는 프로모션 기간 동안 프로모션 가격을 표시합니다. 그렇지 않은 경우 'release_price'(제품의 원래 가격)와 동일한 데이터를 표시합니다.
- 제품에 여러 프로모션이 있는 경우 `id-product-per_item_promotion` API를 통해 각 프로모션 가격을 얻을 수 있습니다.
- 귀하가 BU (CC) / PL 판매자인 경우 'release_price_of_current_price/release_price_of_original_price'는 세금이 포함된 가격을 의미합니다.
- 귀하가 마켓플레이스인 경우 'release_price_of_current_price/release_price_of_original_price'는 세금이 제외된 가격을 의미합니다.

## 2. 제품 가격 업데이트

**API:** `id-product-update_price`

- 제품에 베리에이션이 있는 경우 제품의 여러 베리에이션을 업데이트하여 한 번의 호출로 가격을 업데이트할 수 있습니다.
- 제품에 베리에이션이 없는 경우 하나의 오퍼를 업데이트하는 것은 한 번의 호출이 됩니다. 둘 이상의 오퍼를 업데이트해야 하는 경우 모든 오퍼가 여러 번의 호출이 아닌 한 번의 호출에 있어야 합니다.
- `id-product-per_item_price` API에서 가격 업데이트 규칙을 확인하십시오.

### 1) 베리에이션이 없는 제품의 가격 업데이트 예시

```json
"Price_API": {"original_price": "1110"}
```

### 2) 베리에이션이 있는 제품의 가격 업데이트 예시

```json
"Price_API": [
    {"sku": "sell_br", "sku3_original_price": "1234", "Original_price": "22.52"}
]
```

BU (CC) / PL 판매자인 경우 original_price는 세금이 포함된 가격으로 업데이트됩니다.

## 3. 글로벌 제품 가격 업데이트

*다음은 별도의 CBG/CMREC를 가진 해외 판매자에게만 적용됩니다.*

**API:** `id-global_product-update_price`

- 글로벌 제품에 베리에이션이 있는 경우 한 번의 호출로 이 글로벌 제품의 여러 베리에이션 가격을 업데이트할 수 있습니다. 그렇지 않은 경우 한 번의 호출로 하나의 global_item_id 업데이트만 지원합니다. 둘 이상의 global_item_id를 업데이트해야 하는 경우 API를 여러 번 호출해야 합니다.
- 글로벌 제품의 가격은 `countable-per_exchange_rate` API를 통해 가격 통화를 확인하십시오.
- 글로벌 제품의 가격을 제품에 자동으로 동기화해야 하는 경우 셀러 센터의 제품 목록 페이지에서 "자동 동기화"를 선택하십시오. 이 기능이 있는 제품의 경우 공식을 기반으로 `id-global-product-update_price` API를 통해 직접 가격을 업데이트할 수 있습니다. 그렇지 않은 경우 `id-product-update_price` API를 통해 가격을 업데이트할 수 있습니다.

## 4. 제품 재고 가져오기

**API:** `id-product-per_item_basic_info`

- 제품에 베리에이션이 없는 경우 `id-product-per_item_basic_info` API를 사용하여 재고 정보를 가져오십시오.
- 제품에 베리에이션이 있는 경우 `id-product-per_model_info` API를 사용하여 재고 정보를 가져오십시오.

### API 응답:

```json
"stock_API": {
    "return": {
        "total_reserved_stock": 0,
        "total_available_stock": 100
    },
    "seller_stock": {
        "total_id": "788",
        "Stock": 48
    },
    "resale_stock": {
        "total_id": "788",
        "Stock": 30
    },
    "resale_stock": {
        "total_id": "788",
        "Stock": 30
    }
}
```

### 참고 사항
- 제품은 seller_stock과 resale_stock을 모두 가질 수도 있고 여러 위치의 재고를 가질 수도 있습니다.
- 해외 재고 계산 로직은 FAQ를 참조하십시오.

## 5. 제품 재고 업데이트

**API:** `id-product-update_stock`

- 제품에 베리에이션이 있는 경우 제품의 여러 베리에이션을 업로드하여 한 번의 호출로 재고를 업데이트할 수 있습니다.
- 제품에 베리에이션이 없는 경우 하나의 오퍼를 업데이트하는 것은 한 번의 호출이 됩니다. 둘 이상의 오퍼를 업데이트해야 하는 경우 모든 오퍼가 여러 번의 호출이 아닌 한 번의 호출에 있어야 합니다.
- 재고를 업데이트해야 하는 경우 seller_stock update choose_stock, 재고 유형과 함께 API를 호출해야 합니다. 선택하는 경우 업데이트할 수 있는 모든 유형의 재고, 전체 목록을 설정할 수도 있습니다(모든 유형의 재고를 한 번의 호출로 보낼 수 있음).

### 1) 베리에이션이 없는 제품의 재고 업데이트 예시

```json
"Stock_API": {"resale_stock": {"Product": {"Stock"}}}
```

### 2) 베리에이션이 있는 제품의 재고 업데이트 예시

```json
"Stock_API": {"resale_API": {"Add", "resale_stock": {"product": {"10"}}, "resale_API": {"101", "resale_stock": {"product": {"Stock"}}}}}
```

### 참고 사항
- 제품에 베리에이션이 있는 경우 베리에이션 간의 가격 차이가 특정 배수를 초과할 수 없습니다. 예를 들어 가장 비싼 베리에이션 중 하나의 재판매 가격을 가장 저렴한 베리에이션의 가격으로 나눈 값이 특정 숫자를 초과해서는 안 됩니다.

| Region | multiple |
|--------|----------|
| BR | 4 |
| SG/VN/PH/TH/MY/HA | 5 |
| ID | 5 |
| GLOB | 5 |
| CN/EC | 3 |

*제품이 특정 프로모션에 참여하는 경우 판매자는 제품 가격을 수정할 수 없습니다. 자세한 내용은 FAQ를 확인하십시오.* https://open-shopee-partners-faq

## 6. 글로벌 제품 재고 업데이트

*다음은 별도의 CBG/CMREC를 가진 해외 판매자에게만 적용됩니다.*

**API:** `id-global_product-update_stock`

- 글로벌 제품에 베리에이션이 있는 경우 한 번의 호출로 글로벌 제품의 여러 베리에이션 재고를 업데이트할 수 있습니다. 그렇지 않은 경우 한 번의 호출로 하나의 global_item_id 업데이트만 지원합니다. 둘 이상의 global_item_id를 업데이트해야 하는 경우 API를 여러 번 호출해야 합니다.
- 베리에이션이 없는 글로벌 제품의 경우 `id-global_product-update_stock` API를 사용하여 재고를 업데이트할 수 있습니다. 글로벌 재고를 업데이트한 후 글로벌 제품에 해당하는 제품의 재고가 자동으로 업데이트됩니다. 베리에이션이 있는 제품의 경우 `id-product-update_stock` API만 사용하여 재고를 직접 업데이트할 수 있습니다.
- 자동 동기화 재고가 활성화되지 않은 제품의 경우 `id-global_product-update_stock` API를 사용하여 global_item_stock을 업데이트하십시오. 둘 이상의 global_item_id를 업데이트해야 하는 경우 여러 번 요청할 수 있습니다.

## 사용 사례

1. 배리에이션이 있거나 없는 제품의 가격 업데이트.
2. 배리에이션이 있거나 없는 제품의 재고 수준 업데이트.
3. 글로벌 제품을 사용하는 국경 간 판매자의 가격 및 재고 관리.
4. 제품 재고 정보 검색.

## 관련 API

- price_API
- id-product-update_price
- id-product-per_item_price
- id-global_product-update_price
- countable-per_exchange_rate
- id-product-per_item_basic_info
- id-product-per_model_info
- stock_API
- id-product-update_stock
- id-global_product-update_stock

---

## 원문 (English)

### Summary

This guide provides information on how to manage stock and prices for products, including updating product prices and stock levels. It covers different scenarios such as products with and without variants, as well as global products for cross-border sellers.

### Content

# Stock & Price Management

## API response

```json
"price_API": {
    "seller_price": 500,
    "original_price": 1000,
    "release_price_of_current_price": 937,
    "release_price_of_original_price": 1663,
    "customer": "COP"
}
```

### Please note that
- If your product has ongoing promotion, 'current_price' will show the promotion price during the promotion period. If not, it will show the same data as the 'release_price' (the original price of the product).
- If your product has multiple promotions, you can get each promotion price through `id-product-per_item_promotion` API.
- If you are are BU (CC) / PL seller, 'release_price_of_current_price/release_price_of_original_price' means the price with tax included.
- If you are a marketplace, 'release_price_of_current_price/release_price_of_original_price' means the price with tax excluded.

## 2. Updating product price

**API:** `id-product-update_price`

- If a product has variants, you can update multiple variants of the product to update the price in one call.
- If a product does not have variants, updating one offer, will be one call. If you need to update more than one offer, all of the offers must be in one call from multiple times.
- Please check the price updating rule at the `id-product-per_item_price` API.

### 1) Example of updating the price of a product without variants

```json
"Price_API": {"original_price": "1110"}
```

### 2) Example of updating the price of a product with variants

```json
"Price_API": [
    {"sku": "sell_br", "sku3_original_price": "1234", "Original_price": "22.52"}
]
```

Note that if you are are BU (CC) / PL seller, the original_price is updated to be the withtest price.

## 3. Updating global product price

*The following is only applicable to cross-border sellers who have separate CBG/CMREC*

**API:** `id-global_product-update_price`

- If a global product has variants, you can update the price of multiple variants of this global product in one call. If not, only one supports updating one global_item_id in one call. If you need to update more than one global_item_id, you need to call API multiple times.
- For the price of the global product, please check the price currency that through the `countable-per_exchange_rate` API.
- If you need the price of global products automatically synchronized to the product, please tick "Auto-sync" on Seller Center's Product List page. For products with this feature, you can update the price directly through the `id-global-product-update_price` API based on the formula. If not, you can update the price through `id-product-update_price` API.

## 4. Getting product stock

**API:** `id-product-per_item_basic_info`

- If a product has no variants, please use `id-product-per_item_basic_info` API to get the stock information.
- If a product has variants, please use `id-product-per_model_info` API to get the stock information.

### API response:

```json
"stock_API": {
    "return": {
        "total_reserved_stock": 0,
        "total_available_stock": 100
    },
    "seller_stock": {
        "total_id": "788",
        "Stock": 48
    },
    "resale_stock": {
        "total_id": "788",
        "Stock": 30
    },
    "resale_stock": {
        "total_id": "788",
        "Stock": 30
    }
}
```

### Please note
- Products may have both seller_stock and resale_stock, or it may have stock from multiple locations.
- For cross-border stock calculation logic, please refer to the FAQ.

## 5. Updating product stock

**API:** `id-product-update_stock`

- If a product has variants, you can upload multiple variants of the product to update the stock in one call.
- If a product has no variants, updating one offer, will be one call. If you need to update more than one offer, all of the offers must be in one call from multiple times.
- If you need to update stock, seller_stock update choose_stock, you need to call the API with stock type. If you choose, all of the types of stock that can be updated, you also can set the full list (all types of stocks can be sent in one call).

### 1) Example of updating the stock of a product with no variants

```json
"Stock_API": {"resale_stock": {"Product": {"Stock"}}}
```

### 2) Example of updating the stock of a product with variants

```json
"Stock_API": {"resale_API": {"Add", "resale_stock": {"product": {"10"}}, "resale_API": {"101", "resale_stock": {"product": {"Stock"}}}}}
```

### Please note
- If a product has variants, the price difference between the variants cannot exceed a certain multiple. For example, the resale of one of the most expensive variants divided by the price of the cheapest variants should not exceed certain numbers.

| Region | multiple |
|--------|----------|
| BR | 4 |
| SG/VN/PH/TH/MY/HA | 5 |
| ID | 5 |
| GLOB | 5 |
| CN/EC | 3 |

*If your product participates in certain promotion, sellers are not allow to modify either price of the product. More detail please check FAQ:* https://open-shopee-partners-faq

## 6. Updating global product stock

*The following is only applicable to cross-border sellers who have separate CBG/CMREC*

**API:** `id-global_product-update_stock`

- If a global product has variants, you can update the stock of multiple variants of the global product in one call. If not, only one supports updating one global_item_id in one call. If you need to update more than one global_item_id, you need to call API multiple times.
- For a global product without variants, you can use the `id-global_product-update_stock` API to update stock. After updating global stock, the inventory of the products corresponding to the global products will be updated automatically. For a product with variants, you can only use the `id-product-update_stock` API to update the stock directly. 
- For products without auto-sync stock enabled, please use the `id-global_product-update_stock` API to update the global_item_stock. If you need to update more than one global_item_id, you can request it multiple times.

---

**End of extraction**

---

**문서 ID**: developer-guide.223
**플랫폼**: shopee
**URL**: https://open.shopee.com/developer-guide/223
**처리 완료**: 2025-10-16T08:47:59
