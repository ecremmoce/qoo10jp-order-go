# Product Catalog API - 개발자 가이드

**카테고리**: API 참조
**난이도**: 중간
**중요도**: 4/5
**최종 업데이트**: 2025-10-16T08:35:53

## 개요

본 개발자 가이드는 Product Catalog API 사용 방법에 대한 정보를 제공하며, 카테고리 및 속성에 중점을 둡니다. 글로벌 카테고리 데이터, 카테고리 트리, 추천 카테고리, 속성 유형, 속성 값 데이터 유형 및 속성 값 단위를 다룹니다.

## 주요 키워드

- Product Catalog API
- 카테고리
- 속성
- Category Tree
- Attribute Types
- Attribute Values
- Global Category Data
- 추천 카테고리

## 본문

# 상품 카탈로그 API - 개발자 가이드

## API 루트 리소스 > 상품 인용 속성

참고: product.fetch는 필수 속성을 사용합니다. 하지만 유연한 카테고리 데이터 또는 `us.product.req_category` API에 있을 수 있습니다. 각 매치의 카테고리는 고유한 category_id를 가집니다.

## 1.1 글로벌 카테고리 데이터

Shopee의 카테고리 트리 데이터는 모든 마켓에 적용되지만, 마켓의 현지 정책에 따라 일부 카테고리는 활성화된 속성에 없습니다. 즉, 이 API를 사용하여 카테고리 목록을 가져올 때 API는 전체 목록을 반환하지만 모든 카테고리 목록입니다. 하지만 실제로는 마켓플레이스 상점에서 여기에 보이는 것보다 더 많은 것을 얻을 수 있습니다. 하지만 실제로는 마켓플레이스 상점에서 카테고리의 작은 부분만 사용할 수 있습니다. 요컨대, 글로벌 데이터는 모든 카테고리 목록을 의미하며, 이 카테고리 데이터에 따라서만 의미가 있습니다. 다양한 유형의 배치에 따라 지원되는 카테고리 데이터에도 약간의 차이가 있습니다.

따라서 가장 정확한 카테고리 데이터를 얻으려면 shop_id를 쿼리할 상점을 선택하는 것이 좋습니다.

## 1.2 카테고리 트리

카테고리 트리는 세 가지 레벨과 상점에서 사용할 수 있는 카테고리를 가집니다.

글로벌 카테고리 API를 호출하면 반환되는 트리는 모든 카테고리를 포함하는 완전한 Shopee 카테고리 트리이지만, 특정 마켓의 카테고리에 대한 상점은 이 카테고리 아래에 직접 제품을 추가하는 것을 지원하지 않을 수 있습니다. 즉, children-filled는 마지막 레벨 카테고리를 의미하고, 그렇지 않으면 카테고리에 하위 항목이 있음을 의미합니다. 레벨 1 카테고리 또는 레벨 2 카테고리 아래의 카테고리만 제품을 생성하거나 업로드하는 데 사용할 수 있습니다.

예: 카테고리 트리 경로가 레벨 1 카테고리 -> 레벨 2 카테고리 -> 레벨 3 카테고리인 경우 API는 아래와 같은 카테고리 데이터를 반환합니다.

```
"display_category_name": "레벨 1 카테고리",
"has_children": true,
```

레벨 1 카테고리:

```
"original_category_name": "레벨 1 카테고리",
"display_category_id": 100001
```

레벨 2 카테고리:

```
"display_category_name": "레벨 2 카테고리",
"has_children": true,
```

레벨 2 카테고리:

```
"original_category_name": "레벨 2 카테고리",
"display_category_id": 100002
```

레벨 3 카테고리:

```
"display_category_name": "레벨 3 카테고리",
"has_children": false,
```

레벨 3 카테고리:

```
"original_category_name": "레벨 3 카테고리",
"display_category_id": 100003
```

## 1.3 추천 카테고리

제품 속성을 사용하여 제품을 업로드한 후에는 `v2.product.category_recommend API`를 호출할 수도 있습니다. API는 제품 이름과 제품 이미지를 기반으로 추천 카테고리 목록을 반환합니다.

## 2. 속성

각 카테고리에는 다른 속성 목록이 있을 수 있습니다. `v2.product.get_attribute_tree` API는 입력 카테고리에 대한 속성 데이터를 반환합니다. 하지만 API는 카테고리 목록을 반환하지 않으므로 카테고리 API를 호출하여 어떤 카테고리가 속성 데이터를 받을 수 있는지 확인해야 합니다. 사용자는 제품 속성 및 속성 값이 규칙 시리즈로 정의되었는지 확인해야 합니다. "category_id" 섹션 아래의 "has_custom" 필드 값은 특정 속성을 사용자 정의할 수 있는지 확인합니다.

## 2.1 필수 및 선택적 속성

제품 내의 일부 속성의 경우 속성을 업로드해야 하는 반면 선택적 속성은 제공할 수 있습니다.

판매자는 선택적 제품 속성을 제품에 추가할 수 있습니다. 예를 들어 반환 결과의 `is_mandatory`는 반드시 업로드해야 하는 반면 보증 관련 속성과 같은 선택적 속성은 판매자 결정에 따라 업로드할 수 있습니다.

즉, "is_mandatory": 선택적 속성은 선택적 속성입니다.

## 2.2 속성 유형

속성이 여러 선택을 지원하는지, 사용자 정의 값을 허용하는지 등의 요소를 기반으로 속성 유형을 분류합니다.

| # | 유형 | 정의 | 설명 | 사용자 정의 값이 필요함 | 속성 값을 입력할 수 있음 |
|---|------|------------|-------------|--------------------------|-------------------------------|
| 1 | int | SINGLE_DROP_DOWN | 단일 선택 드롭다운. 판매자는 API에서 반환된 `attribute_value_list`에서 반환된 허용된 값 목록에서 하나의 값만 선택하여 업로드할 수 있습니다. | 아니요 | 예 |
| 2 | int | SINGLE_COMBO_BOX | 단일 선택 드롭다운 + 텍스트 입력 필드. 판매자는 허용된 속성 값 목록에서 값을 선택하거나 속성 값을 자유롭게 입력할 수 있습니다. API에서 반환된 `attribute_value_list`는 사용자 정의 값 목록을 가져올 수 있습니다. | 예 | 예 |
| 3 | int | FREE_TEXT_FIELD | 자유 텍스트 필드 - 사용자 정의 값을 설정할 수 있습니다. | 예 | 아니요 |
| 4 | int | MULTI_DROP_DOWN | 다중 선택 드롭다운. 판매자는 API에서 반환된 `attribute_value_list`에서 반환된 속성 값 목록에서 여러 옵션을 선택하여 업로드하거나 사용자 정의 값을 설정할 수 있습니다. | 아니요 | 예 |
| 5 | int | MULTI_COMBO_BOX | 다중 선택 드롭다운 + 텍스트 입력 필드. 판매자는 API에서 반환된 `attribute_value_list`에서 반환된 속성 값 목록에서 여러 옵션을 선택하거나 속성 값을 자유롭게 입력하여 업로드하거나 사용자 정의 값을 설정할 수 있습니다. | 예 | 예 |

*다음 그림은 `input_type`이 SINGLE_COMBO_BOX인 예제를 보여줍니다. 판매자는 속성 목록(Upright Trainers)에서 값을 선택하거나 "새 항목 추가" 항목(Long A)을 변경하여 사용자 정의 값을 설정할 수 있습니다.

**예제 UI:**
- Fitness Machine Type [드롭다운 선택]
  - [취소] [확인] 버튼
  - Elliptical Trainers ← 화살표 가리킴
  - **+ 새 항목 추가**
- Fitness Machine Type [드롭다운 선택]
  - [취소] [확인] 버튼
  - Elliptical Trainers
  - **+ 추가**

**참고:**
- 다중 선택 드롭다운을 지원하는 속성이 있는 제품의 경우 "max_value_count" 매개변수를 통해 업로드할 수 있는 최대 속성 값 수를 가져올 수 있습니다.
- 속성 값 목록에서 Shopee에서 제공하는 속성 값을 사용하거나 사용 가능한 경우 "attribute_value" API 매개변수를 통해 추가할 수 있습니다.

## 2.3 속성 값의 데이터 유형

속성 값에는 다른 데이터 유형이 있습니다. 예를 들어 필요한 데이터 유형에 따라 속성 값을 업로드할 수 있습니다. `v2.product.get_attribute_tree` API에서 반환된 `input_validation_type` 필드를 통해 속성 값의 데이터 유형을 얻을 수 있습니다.

| # | input_validation_type | 유형 | 설명 | 샘플 |
|---|-----------------------|------|-------------|---------|
| 0 | int | VALIDATE_NO_VALIDATE_TYPE | | |
| 1 | int | VALIDATE_INT_TYPE | | |
| 2 | int | VALIDATE_STRING_TYPE | | |
| 3 | int | VALIDATE_FLOAT_TYPE | | |
| 4 | int | VALIDATE_DATE_TYPE, 다음 두 가지 형식을 포함합니다. | 제품을 편집하거나 편집하는 동안 예를 들어 아래의 제조 날짜를 사용합니다. 형식 1: YYYY-MM-DD, 예: 31082021 형식 2: YYYY.Y, 예: 2021.3 | 입력한 형식이 숫자 + 단위 형식인 경우. API에서 제공하는 속성 값 목록 속성에서 볼 수 있는 해당 속성에 대한 단위(attribute_unit)를 표시합니다. 예를 들어 속성 값을 사용할 수 있는 경우: |

**참고:**
- 데이터 유형 속성의 경우 두 가지 형식이 있습니다. 하나는 MONTH.DAY(2021.2)이고 다른 하나는 DATE(YYYY-MM-DD)입니다. 이러한 형식은 `v2.product.get_attribute_tree` API에서 반환된 `format_type` API에서 제공합니다.
- 데이터 유형 속성의 경우 숫자 단위 형식은 값이어야 합니다. 그러나 "original_value_name" 필드와 같은 제품 속성의 경우 제품 정보를 비판할 때 "original_value_name" 필드는 입력의 "attribute_value" 필드입니다.

---

# 개발자 가이드 - 제품 속성

## 2.4 속성 값 단위

반환된 속성 유형이 "정량적"인 경우 "cm" 및 "oz"와 같은 단위를 선택할 수 있는 옵션이 있는 속성 값을 의미합니다. 따라서 먼저 어떤 속성 값에 단위가 필요한지, 해당 속성 값에 사용할 수 있는 단위가 무엇인지 이해해야 합니다. `c_product_get_attribute_unit` API를 호출하여 이 정보를 얻을 수 있습니다.

### 정량적 단위 정보

반환된 매개변수 "format_type"이 "FORMAT_QUANTITATIVE_WITH_UNIT"인 경우 속성 값에 단위가 필요함을 의미합니다(예: 속성 값의 "attribute_unit"). "attribute_unit" 객체 아래의 "unit_list"에는 이 속성에 사용할 수 있는 모든 단위가 포함되어 있으며 "default_unit"는 기본적으로 표시되는 단위를 보여줍니다.

Shopee에서 제공하는 속성 값 목록의 경우 값과 단위를 별도로 표시합니다. 특히 "name" 필드는 속성 값이고 "unit" 부분은 표시 용도로만 사용되며 "default_attribute_value_unit"는 실제로 표시된 형식으로 제출되는 단위입니다(예: 킬로그램의 경우 "kg").

---

## 2.5 상위 속성 및 상위 브랜드

예를 들어 "체중 유형" 속성의 경우 관련 속성 값은 "체중" 및 "용량"입니다. "체중" 속성은 "체중 유형" 상위 속성 값입니다. 따라서 판매자가 "체중 유형" 속성에 대해 "체중"을 선택하면 "체중 유형" 속성이 표시되고 판매자가 "선택-맨발" 변형 값을 입력하고 체중 유형 속성을 표시하는 것이 좋습니다.

| 체중 유형 | 체중 | 체중 크기 | 허용 오차(kg) | 기타 |
|-------------|-------------|-------------|----------------|-----|
| 체중 | 맨발 | | | |

"체중 선택" 값은 체중 유형 속성을 표시합니다.

| 체중 유형 | 체중 맨발 | 체중 크기 | 허용 오차(kg) | 기타 |
|-------------|---------------|-------------|----------------|-----|
| 체중 맨발 유형 | | | | |

---

## API 응답 예제

```json
"v2.product.get_attribute_time" API 포털 응답

{
  "attribute_id": 100001,
  "mandatory": false,
  "name": "체중 유형",
  "attribute_value_list": [
    {
      "value_id": 0,
      "name": "맨발",
      "unit": "text"
    },
    {
      "language": "en-text",
      "value": "BBB"
    }
  ],
  {
    "value_id": 200,
    "name": "체중",
    "attribute_unit": {
      "attribute_id": 100002,
      "mandatory": false,
      "name": "체중 유형",
      "attribute_value_list": [
        {
          "value_id": 10,
          "level_type": [
            {
              "language": "en-text",
              "value": "BBB"
            }
          ]
        }
      ],
      "attribute_unit": {
        "need_input": true,
        "input_validation_type": 2,
        "format_type": "FORMAT_QUANTITATIVE_WITH_UNIT",
        "support_search_in_value": false
      }
    },
    "level_text": {
      "language": "en-text",
      "value": "請輸入數值BBB"
    }
  },
  "BJJ-text": {
    "language": "en-text",
    "value": "請輸入數值"
  }
}
]
},
{
  "attribute_unit": {
    "need_input": true,
    "input_validation_type": 2,
    "format_type": "",
    "unit_list": [],
    "support_search_in_value": false
  },
  "level_text": {
    "language": "en-text",
    "value": "입력 수치 정보"
  }
}
}
```

이 예에서 "체중 유형"은 상위 속성이고 "체중 유형"은 하위 속성입니다. 하위 속성을 업로드하려면 관련 상위 속성도 제출해야 합니다.

---

## 2.6 추천 속성

판매자가 제품에 대한 추천 속성을 빠르게 찾을 수 있도록 `c_product_list_recommend_attribute` API를 호출하여 검색할 수 있습니다. 일부 카테고리에만 추천 속성이 있습니다. 일반적으로 API는 추천 속성이 있는 카테고리만 나열합니다.

---

## 3. 브랜드

제품에 대한 브랜드 정보를 업로드할 수 있지만 특정 카테고리에서는 브랜드 정보도 특정 카테고리로 분류되므로 그렇게 할 필요가 없습니다. 일부 카테고리에서는 브랜드 선택이 필요합니다. 카테고리에 브랜드 정보가 필요한지 확인하려면 `c_product_get_category` API를 호출하십시오(카테고리 구조 참조).

"mandatory" 필드는 카테고리에 브랜드 정보를 입력해야 함을 의미합니다.

판매자가 업로드하려는 브랜드가 드롭다운 목록에 없는 경우 Shopee에 등록된 브랜드 신청서를 제출할 수 있습니다. `c_product_register_brand` API를 호출하고 QC 프로세스에 대한 이 `App`을 참조하십시오.

> **참고:** 판매자는 모든 카테고리에 대해 브랜드 신청서를 제출할 수 없습니다. 브랜드 중심 브랜드(예: 브랜드가 브랜드 값을 선택할 수 있는 브랜드)는 제외합니다.

---

## 4. 배송 일수

구매자가 구매한 후 판매자는 해당 기간 내에 주문을 이행하는 데 필요한 1-5 카테고리 판매 제한을 제공해야 합니다. 다른 카테고리 및 다른 판매 유형의 경우 판매자는 최대 7일(선주문 품목 제외)까지 지정한 "배송 일수" 내에 주문을 이행해야 합니다. 구매자는 제품 페이지에서 판매자가 주문을 이행하는 데 얼마나 걸리는지 확인할 수 있습니다. `c_product_list_dts_limit` API를 통해 배송 일수 정보를 얻을 수 있습니다.

또는 카테고리의 경우 "dts_limit"를 지정하면 해당 카테고리가 사전 판매를 지원하지 않음을 의미합니다. "dts" 필드가 1이면 판매자는 7일 이하의 "days_to_ship"을 선택할 수 있습니다.

> **테이블 유형의 낮은 체인을 출시하고 있으므로** 일부 화이트리스트에 등록된 판매자만 "days_to_ship" 매개변수를 사용할 수 있습니다. 자세한 내용은 https://open.shopee.com/documents/v2/v2.product.add_item?module=89&type=1을 참조하십시오.

---

## 6. 제품 제한

제품 정보에는 입력할 수 있는 문자 길이 또는 이미지 크기 제한과 같은 특정 제한이 있습니다. 이러한 제한은 Shopee 플랫폼 전체에서 통합됩니다. 자세한 내용은 product_tag_list API `shopee.com/1.0.0/product/get_item_limit`를 참조하여 정보를 얻으십시오.

---

# 5. 사이즈 차트

일부 카테고리의 경우 판매자가 제품에 대한 하나의 사이즈 차트 이미지를 업로드할 수 있도록 지원합니다. v2.product.support_size_chart API를 사용하여 카테고리(리프 카테고리)가 사이즈 차트를 지원하는지 확인할 수 있습니다.

*테이블 유형의 사이즈 차트를 출시하고 있으며 현재 일부 화이트리스트에 등록된 판매자만 판매자 센터를 통해 추가할 수 있습니다. 오픈 API는 테이블 사이즈 차트를 지원하는 카테고리 반환 및 업로드를 지원하지 않습니다. 따라서 화이트리스트에 등록된 판매자는 v2.product.support_size_chart API의 결과를 무시하십시오.

# 6. 제품 제한

제품 이름에 입력할 수 있는 문자 길이, 제품 가격 범위 등 제품 정보에는 특정 제한이 있습니다. 마켓 및 판매자 유형에 따라 다른 제한이 있습니다. v2.product.get_item_limit API를 통해 설정한 제한을 얻을 수 있습니다.

# 7. 제품 물류

v2.logistics.get_channel_list API를 통해 logistics_channel_id를 얻을 수 있으며, 제품에 대해 enabled=true 및 mask_channel_id=0인 채널만 선택할 수 있습니다.

## 사용 사례

1. 마켓플레이스를 위한 제품 카테고리 가져오기.
2. 특정 카테고리에 대한 속성 정보 검색.
3. 제품 정보를 기반으로 카테고리 추천.
4. 올바른 데이터 유형 및 단위를 사용하여 제품 속성 업로드.
5. 적절한 입력 유형(드롭다운, 텍스트 필드 등)으로 제품 속성 표시.

## 관련 API

- product.fetch
- us.product.req_category
- v2.product.category_recommend API
- v2.product.get_attribute_tree
- attribute_value
- c_product_get_attribute_unit

---

## 원문 (English)

### Summary

This developer guide provides information on how to use the Product Catalog API, focusing on categories and attributes. It covers global category data, category trees, recommended categories, attribute types, attribute value data types, and attribute value units.

### Content

# Product Catalog API - Developer Guide

## API Root Resource > Product citation properties

Note: product.fetch takes a required attribute. But it can be in flexible category data or `us.product.req_category` API. Each match's category for a unique category_id.

## 1.1 Global category data

Shopee's category tree data applies to all markets, but according to the local policies of the market, some of the categories are not in enabled properties. That is, when you use this API to get a list of categories, the API will return the full list, but a list of all categories. But in fact in a Marketplace shop, you can get more than what you see here is the full list. But in fact in a Marketplace shop, you can use only a small part of the categories. In short, global data means all the list of categories, and according to the category data with this only. For different types of batches, there are also some differences in the supported category data.

So, in order for you to get the most accurate category data, it is recommended to select a shop to query shop_id.

## 1.2 Category tree

The category tree has three levels and available categories for the shop.

When you call the global category API, the tree that is returned is the complete Shopee category tree containing all categories, but the shop for a category in a particular market may not support adding products directly under this category. That is, children-filled means that the last level category, otherwise, it means that the category has children. Please note that only categories under Level 1 category or Level 2 category can be used to create or upload products.

For example: If the category tree path is Level 1 category -> Level 2 category -> Level 3 category, then API return like the category data shown below.

```
"display_category_name": "Level 1 category",
"has_children": true,
```

Level 1 category:

```
"original_category_name": "Level 1 category",
"display_category_id": 100001
```

Level 2 category:

```
"display_category_name": "Level 2 category",
"has_children": true,
```

Level 2 category:

```
"original_category_name": "Level 2 category",
"display_category_id": 100002
```

Level 3 category:

```
"display_category_name": "Level 3 category",
"has_children": false,
```

Level 3 category:

```
"original_category_name": "Level 3 category",
"display_category_id": 100003
```

## 1.3 Recommend categories

After you have uploaded your products using the attributes of the product, you can also call `v2.product.category_recommend API`. API will return a list of recommended categories based on the product name and product image.

## 2. Attribute

Each category may have different attribute lists. The `v2.product.get_attribute_tree` API will return the attribute data for the input category. But the API will not return a list of categories, so you will need to call the category API to get which categories can receive attribute data. A user must ensure that product attributes and attribute values are defined as rule series the value of the "has_custom" field under the "category_id" section ensures that certain attributes can be customized.

## 2.1 Required and optional Attributes

For some attributes within products, an attribute must be uploaded, while optional attributes can be provided.

Sellers can add optional product attributes to the product. For example, `is_mandatory` in the return result must be uploaded, while optional attributes such as warranty-related attributes can be uploaded based on seller decision.

That is, "is_mandatory": optional attributes are optional attributes.

## 2.2 Type of attributes

We classify attributes types based on factors such as whether the attribute can supports multiple selections, whether it allows custom values, etc.

| # | Type | Definition | Description | Custom value is required | Attribute values can be input |
|---|------|------------|-------------|--------------------------|-------------------------------|
| 1 | int | SINGLE_DROP_DOWN | Single-select dropdown. Sellers can only choose one value from the list of allowed values returned by `attribute_value_list` returned by the API to upload. | No | Yes |
| 2 | int | SINGLE_COMBO_BOX | Single-select dropdown + text input field. Sellers can select a value from the list of allowed attribute values, or freely enter an attribute value. `attribute_value_list` returned by the API can get the list of custom value. | Yes | Yes |
| 3 | int | FREE_TEXT_FIELD | Free text field - can set a custom value. | Yes | No |
| 4 | int | MULTI_DROP_DOWN | Multi-select dropdown. Sellers can select multiple options from the list of attribute values returned by `attribute_value_list` returned by the API to upload, or set a custom value. | No | Yes |
| 5 | int | MULTI_COMBO_BOX | Multi-select dropdown + text input field. Sellers can select multiple options from the list of attribute values or freely enter attribute values returned by `attribute_value_list` returned by the API to upload, or set a custom value. | Yes | Yes |

*The following figure shows an example with `input_type` is SINGLE_COMBO_BOX. Sellers can select a value from the attribute list (Upright Trainers), or set a custom value by changing the "Add a new button" item (Long A).

**Example UI:**
- Fitness Machine Type [Please select dropdown]
  - [Cancel] [OK] buttons
  - Elliptical Trainers ← Arrow pointing to
  - **+ Add a new item**
- Fitness Machine Type [Please select dropdown]
  - [Cancel] [OK] buttons
  - Elliptical Trainers
  - **+ Add**

**Please note:**
- For products with attributes that support multi-select dropdown, you can get the maximum number of attribute values that can be uploaded through the "max_value_count" parameter.
- You can use the attribute values provided by Shopee in the attribute values list, or add through the "attribute_value" API parameter if it's available.

## 2.3 Data types of attribute values

Attribute values have different data types. For example: You can upload the attribute values according to the required data type. You can obtain the data type of the attribute value through the `input_validation_type` field returned by the `v2.product.get_attribute_tree` API.

| # | input_validation_type | Type | Description | Samples |
|---|-----------------------|------|-------------|---------|
| 0 | int | VALIDATE_NO_VALIDATE_TYPE | | |
| 1 | int | VALIDATE_INT_TYPE | | |
| 2 | int | VALIDATE_STRING_TYPE | | |
| 3 | int | VALIDATE_FLOAT_TYPE | | |
| 4 | int | VALIDATE_DATE_TYPE, including two formats: | While editing or editing a product, for example, use the Manufacture Date below: Format1: YYYY-MM-DD, such as: 31082021 Format2: YYYY.Y, e.g.,2021.3 | When the format you entered is a number + unit format. Displays units (attribute_unit) for corresponding attributes which you can view in the attribute values list attribute provided by the API. For example, if the attribute value is in available, then: |

**Please note:**
- For the data type attribute, there are two formats: one is MONTH.DAY (2021.2) and the other is DATE (YYYY-MM-DD). These formats are provided by the `format_type` API returned by the `v2.product.get_attribute_tree` API.
- For the data type attribute, the number unit format must be a value. For product attributes like "original_value_name" field, however, when critiquing product information, the "original_value_name" field is the input's "attribute_value" field.

---

# Developer Guide - Product Attributes

## 2.4 Attribute value units

When the returned attribute type is "quantitative", it means the attribute values with options to choose units such as "cm" and "oz". Therefore, you need to first understand which attribute values require a unit and which units are available for those attribute values. You can call the `c_product_get_attribute_unit` API to obtain this information.

### Quantitative Unit Information

When the returned parameter "format_type" is "FORMAT_QUANTITATIVE_WITH_UNIT" is present, it means the attribute value requires a unit, such as "attribute_unit" for attribute value. Under the "attribute_unit" object, "unit_list" contains all units available for this attribute, and the "default_unit" shows which unit is displayed by default.

For the attribute value list provided by Shopee, we'll display the value and the unit separately. Specifically, the "name" field is the attribute value, the "unit" part is for display purposes only, and the "default_attribute_value_unit" is the unit that is actually submitted in the format shown, e.g., "kg" for kilograms.

---

## 2.5 Parent attribute and parent brand

For example, to the "weight Type" attribute, the associated attribute values are "body-Weight" and "capacity". The "body-Weight" attribute is the "Weight Type" parent attribute value. Therefore, where the seller selects "body-Weight" for the "Weight Type" attribute, the "Body-Weight-Type" attribute will be displayed, and we recommend that the seller fill in the "Selecting-Barefoot" variant value and show the Body-Weight-Type attribute.

| Weight Type | Body Weight | Weight Size | Tolerance (kg) | etc |
|-------------|-------------|-------------|----------------|-----|
| Body Weight | Barefoot | | | |

"Selecting 'Body Weight'" value displays the Body Weight Type Attribute

| Weight Type | Body Barefoot | Weight Size | Tolerance (kg) | etc |
|-------------|---------------|-------------|----------------|-----|
| Body Barefoot Type | | | | |

---

## Example API Response

```json
"v2.product.get_attribute_time" API portal response

{
  "attribute_id": 100001,
  "mandatory": false,
  "name": "Weight Type",
  "attribute_value_list": [
    {
      "value_id": 0,
      "name": "Barefoot",
      "unit": "text"
    },
    {
      "language": "en-text",
      "value": "BBB"
    }
  ],
  {
    "value_id": 200,
    "name": "Body-Weight",
    "attribute_unit": {
      "attribute_id": 100002,
      "mandatory": false,
      "name": "Body Weight Type",
      "attribute_value_list": [
        {
          "value_id": 10,
          "level_type": [
            {
              "language": "en-text",
              "value": "BBB"
            }
          ]
        }
      ],
      "attribute_unit": {
        "need_input": true,
        "input_validation_type": 2,
        "format_type": "FORMAT_QUANTITATIVE_WITH_UNIT",
        "support_search_in_value": false
      }
    },
    "level_text": {
      "language": "en-text",
      "value": "請輸入數值BBB"
    }
  },
  "BJJ-text": {
    "language": "en-text",
    "value": "請輸入數值"
  }
}
]
},
{
  "attribute_unit": {
    "need_input": true,
    "input_validation_type": 2,
    "format_type": "",
    "unit_list": [],
    "support_search_in_value": false
  },
  "level_text": {
    "language": "en-text",
    "value": "輸入數值資訊"
  }
}
}
```

In this example, "Weight Type" is the parent attribute, and "Body Weight Type" is the child attribute. To upload the child attribute, you must also submit the associated parent attribute.

---

## 2.6 Recommended attributes

To help sellers quickly find the recommended attributes for the product, you can call `c_product_list_recommend_attribute` API to retrieve them. There are only recommended attributes for a minority of categories. Typically, the API only lists categories that have recommend attributes.

---

## 3. Brand

You can upload the brand information for the product, but we not require certain categories to do so, since brand information are also categorized under certain categories. Some categories require brand selection. To check whether or not a category requires brand information, make a call to `c_product_get_category` API (see the category structure). 

A "mandatory" field that means that the category must be filled with brand information.

If the brand that the seller wishes to upload is not listed in the dropdown list, you can submit a registered brand application to Shopee, please call `c_product_register_brand` API and refer to this `App` for QC process.

> **Note:** Sellers cannot submit brand applications for any categories. Excluding for brand-centric brands, e.g., a brand can choose the brand value.

---

## 4. Days to ship

After the buyer makes the purchase, the seller must deliver a 1-5 category sales limit required to fulfill orders within that time period. Different categories and different sales types, sellers need to fulfill orders within the "days to ship" you specify, up to a max of 7 days (excluding pre-order items). The buyer will see how long it will take the seller to fulfill the order from the product page. You can get the shipping days information through the `c_product_list_dts_limit` API.

If you do, or for the category, specify a "dts_limit", it means that the category does not support pre-sales. If the "dts" field is 1, the seller can select a "days_to_ship" of 7 days or less.

> **Please note that we are rolling out the low chain of table type**, and thus only some whitelisted sellers can put through the "days_to_ship" parameter. For more detailed information, please refer to: https://open.shopee.com/documents/v2/v2.product.add_item?module=89&type=1

---

## 6. Product restrictions

We have certain restrictions on product information, such as the length of characters that can be filled in or image size limits. These restrictions are unified across Shopee's platforms. For details, please refer to the product_tag_list API `shopee.com/1.0.0/product/get_item_limit` to obtain information.

---

# 5. Size chart

For some categories, we support sellers uploading one size chart image for a product, you can use v2.product.support_size_chart API to check whether the category (leaf category) supports size chart.

*Please note that we are rolling out the size chart of table type, and now only some whitelisted sellers can add through the seller center, open api does not support returning the categories that support the table size chart and uploading. So whitelisted sellers please ignore the results from the v2.product.support_size_chart API.

# 6. Product restrictions

We have certain restrictions on product information, such as the length of characters that can be filled in the product name, the range of product price, etc. We have different restrictions for different markets and different types of sellers. You can get the limits we set through v2.product.get_item_limit API.

# 7. Product logistics

You can get the logistics_channel_id through v2.logistics.get_channel_list API, you can only choose the channel with enabled=true and mask_channel_id=0 for product.

---

**문서 ID**: developer-guide.209
**플랫폼**: shopee
**URL**: https://open.shopee.com/developer-guide/209
**처리 완료**: 2025-10-16T08:35:53
