# 배리언트 관리

**카테고리**: 모범 사례
**난이도**: 중간
**중요도**: 4/5
**최종 업데이트**: 2025-10-16T08:45:21

## 개요

이 가이드는 제품 배리언트 추가, 삭제 및 재정렬을 포함하여 제품 배리언트를 관리하는 방법에 대한 지침을 제공합니다. 티어 구조가 동일하게 유지되는 경우와 변경되는 경우를 다루며, 각 경우에 대한 API 요청 예제를 제공합니다.

## 주요 키워드

- 배리언트 관리
- 제품 배리언트
- 티어 구조
- API
- v2_product_update_tier_variation
- v2_product_init_tier_variation
- product.update_tier_variation
- id.product.update_list_variation

## 본문

# 베리에이션 관리

## API 모범 사례 - 베리에이션 관리

제품 트리 및 베리에이션 구조에 대한 지침은 [Catalog API 모범 사례](link)를 참조하십시오.

베리에이션의 재고 및 가격 관리는 [재고 및 가격 관리](link) 문서를 참조하십시오.

---

## 1. 동일한 티어 구조로 베리에이션 추가

초기화 시 이미 정의된 제품 트리(tier_index 및 colors 포함)에서 베리에이션 상황에 따라 다시 추가해야 하는 경우:

### 기존 베리에이션 구조:

| var_name | option | model_id |
|----------|--------|----------|
| var_sku[0][0] | red | 10000 |
| var_sku[0][1] (추가 예정) | black (추가 예정) | - |
| var_sku[1][0] | blue | 20000 |

**`v2_product_update_tier_variation` API를 호출하여 먼저 빈 옵션을 추가하십시오.**

**요청 예시:**

```json
{
  "tier_id": [REDACTED],
  "tier_variation": [
    {
      "name": "color",
      "option_list": [
        {
          "option": "red"
        },
        {
          "option": "black"
        },
        {
          "option": "blue"
        }
      ]
    },
    {
      "model_list": [
        {
          "tier_index": [0],
          "model_sku": "MAIN"
        }
      ]
    },
    {
      "tier_index": [1],
      "model_sku": "MAIN"
    }
  ]
}
```

---

**`v2_product_init_tier_variation` API를 호출하여 빈 베리에이션에 대한 가격, 재고, SKU 정보를 추가하십시오.**

**요청 예시:**

```json
{
  "tier_id": [REDACTED],
  "model_list": [
    {
      "tier_index": [0],
      "original_price": 20,
      "option_list": [
        {
          "option": "red"
        },
        {
          "option": "black"
        }
      ],
      "normal_stock": 100,
      "model_sku": "red_blacka"
    },
    {
      "price": 300
    }
  ]
}
```

---

## 2. 티어 구조를 변경하는 베리에이션 추가

제품 트리 구조를 변경해야 하는 경우(베리에이션 티어가 2티어가 되거나 다른 상황), 티어 사양을 추가해야 합니다. 단계는 1과 2를 포함합니다.

### 기존 베리에이션 구조:

| var_name | option | model_id |
|----------|--------|----------|
| var_sku[0][0] | red | 10000 |
| var_sku[0][1] | blue | 20000 |

### 변경 후 베리에이션 상황:

| var_name | option | model_id |
|----------|--------|----------|
| var_sku[0][0] | red L | - |
| var_sku[0][1] | red XL | - |
| var_sku[1][0] | blue L | - |
| var_sku[1][1] | blue XL | - |

**티어 구조가 변경되었으므로 API를 호출한 후 기존 모델 정보가 제거되며, `10000`의 기존 모델이 대체되고 `20000`의 모델은 유효하지 않게 됩니다.**

**API:** `v2_product_init_tier_variation`

**요청 예시:**

```json
{
  "tier_id": [REDACTED],
  "tier_variation": [
    {
      "option_list": [
        {
          "option": "red"
        }
      ]
    },
    {
      "option": "L"
    },
    {
      "option": "XL"
    },
    {
      "name": "Color"
    },
    {
      "name": "Size"
    },
    {
      "option_list": [
        {
          "option": "blue"
        }
      ]
    },
    {
      "option": "L"
    }
  ]
}
```

---

# 동일한 티어 구조로 베리에이션 삭제

**상황 1**: 제품에 색상 사양(빨강, 파랑, 검정 색상 포함)이 정의되어 있고 파란색 모델을 삭제해야 합니다.

## 기존 베리에이션 상황

| tier_index | option | model_id |
|------------|--------|----------|
| tier_index[0] | red | 10000 |
| tier_index[1] (% 삭제 예정) | blue (% 삭제 예정) | 20000 |
| tier_index[2] | black | 30000 |

**베리에이션 상황:**

파란색을 삭제해야 합니다. 삭제 후 tier_index는 tier_index[0], tier_index[1]만 가능하므로 베리에이션 tier_index[2]의 정보를 검정색 베리에이션의 정보로 변환해야 합니다.

| tier_index | option | model_id |
|------------|--------|----------|
| tier_index[0] | red | 10000 |
| tier_index[1] | black | 30000 |

## API

`product.update_tier_variation`을 호출하여 옵션을 제거하고 모델 정보를 동시에 덮어씁니다(연결 장착).

```json
{
  "tier_variation": [
    {
      "name": "color",
      "option_list": [
        {
          "option": "red"
        },
        {
          "option": "black"
        }
      ]
    }
  ],
  "model_list": [
    {
      "tier_index": [0],
      "model_id": 10000
    },
    {
      "tier_index": [1],
      "model_id": 30000
    }
  ]
}
```

---

# 상황 2

제품에 색상 및 크기 사양이 정의되어 있고, 색상은 빨강, 파랑, 빨강이며, 모든 파란색 베리에이션을 삭제해야 합니다.

## 기존 베리에이션 상황

| tier_index | option | model_id |
|------------|--------|----------|
| tier_index[0] | red 5L | 10000 |
| tier_index[1] | blue 5L (% 삭제 예정) | 20000 |
| tier_index[2] (% 삭제 예정) | black 5L | 30000 |

---

# 티어 구조를 변경하는 베리에이션 삭제

## 기존 베리에이션 상황

| tier_index | option | model_id |
|------------|--------|----------|
| tier_index[0] | null | 10000 |
| tier_index[1] | not XL | 25000 |
| tier_index[2] [5] (삭제 예정) | blue-L (5) (삭제 예정) | 30000 |
| tier_index[3] [1] (삭제 예정) | blue-XL [2] (삭제 예정) | 40000 |

**API:** `id-product-variant_http_product()`은 옵션과 모델 정보를 동시에 삭제합니다.

**요청 예시:**

```json
{
  "img_id": "[12345678]",
  
  "tier_var":[
    {
      "name": "color",
      "opt":[
        {
          "option": "red"
        }
      ]
    },
    {
      "name": "size",
      "opt":[
        {
          "option": "L"
        },
        {
          "option": "XL"
        }
      ]
    }
  ],
  
  "model_list": [
    {
      "tier_index": [0,0],
      "model_id": 10000
    },
    {
      "tier_index": [0,1],
      "model_id": 25000
    }
  ]
}
```

---

## 삭제 후 베리에이션 상황

| tier_index | option | model_id |
|------------|--------|----------|
| tier_index[0] | null | 10000 |
| tier_index[1] | not XL | 25000 |
| tier_index[2] | - | - |
| tier_index[3] | - | - |

---

## 4. 티어 구조를 변경하는 베리에이션 삭제

**시나리오 3:** 제품에 색상 및 크기 사양이 정의되어 있고, 색상은 빨강, 파랑, XL이며, 빨강 및 칼로리를 삭제해야 합니다.

### 기존 베리에이션 상황

| tier_index | option | model_id |
|------------|--------|----------|
| tier_index[0] | null | 10000 |
| tier_index[1] | not XL | 25000 |
| tier_index[2] | blue-L | 30000 |
| tier_index[3] | blue XL | 40000 |

### 삭제 후 베리에이션 상황

| tier_index | option | model_id |
|------------|--------|----------|
| tier_index[0] | - | - |
| tier_index[1] | - | - |

**API:** `id-product-variant_http_update()`

**요청 예시:**

```json
{
  "img_id": "[12345678]",
  
  "tier_var":[
    {
      "opt":[
        {
          "option": "L"
        },
        {
          "option": "XL"
        }
      ],
      "name": "size"
    },
    {
      "name": [
        {
          "id-productpg": [0],
          "model_id": "bnpL1",
          "original_price": 100,
          "tier_index": [0]
        },
        {
          "id-productpg": [0],
          "model_id": "bnpL1",
          "original_price": 100,
          "tier_index": [1]
        }
      ]
    }
  ]
}
```

---

# 베리에이션 순서 변경

## 시나리오
항목에 빨강 및 파랑 색상의 색상 사양이 정의되어 있는 경우 색상 순서를 파랑, 빨강으로 변경해야 합니다.

## 기존 베리에이션 상황

| var_active | option | model_id |
|------------|--------|----------|
| var_active(1) | red | 10000 |
| var_active(2) | blue | 20000 |

## 변경 후 베리에이션 상황:

| var_active | option | model_id |
|------------|--------|----------|
| var_active | option | model_id |
| var_active(2) | blue | 20000 |
| var_active(1) | red | 10000 |

## API

**id.product.update_list_variation**

요청 예시:

```json
{
  "link_id": "RRDDCCFFVV",
  "list_variants": [
    {
      "name": "Color",
      "options_list": [
        {
          "variant": "red A"
        },
        {
          "variant": "Blue"
        }
      ]
    }
  ],
  "model_list": [
    {
      "tier_index": [0],
      "model_id": 20000
    },
    {
      "tier_index": [1],
      "model_id": 10000
    }
  ]
}
```

---

# 베리에이션 옵션 이름 변경

## 시나리오
항목에 빨강 및 파랑 색상의 색상 사양이 정의되어 있는 경우 색상 이름을 빨강 A, 파랑으로 변경해야 합니다.

## 기존 베리에이션 상황

| var_active | option | model_id |
|------------|--------|----------|
| var_active | option | model_id |
| var_active(1) | red | 10000 |
| var_active(2) | Blue | 20000 |

## 변경 후 베리에이션 상황:

| var_active | option | model_id |
|------------|--------|----------|
| var_active | option | model_id |
| var_active(2) | red A | 10000 |
| var_active(1) | Blue | 20000 |

## API

**id.product.update_list_variation**

요청 예시:

```json
{
  "link_id": "RRDDCCFFVV",
  "list_variants": [
    {
      "name": "Color",
      "options_list": [
        {
          "variant": "red A"
        },
        {
          "variant": "Blue"
        }
      ]
    }
  ],
  "model_list": [
    {
      "tier_index": [0],
      "model_id": 10000
    },
    {
      "tier_index": [1],
      "model_id": 20000
    }
  ]
}
```

---

# API 문서

## JSON 구조

```json
{
  "tier_index": [1],
  "model_id": 20000
}
]
}
```

## 요약

1. **v2.product.init_tier_variation** API의 기능은 0티어 베리에이션을 1티어 베리에이션으로 변경/0티어 베리에이션을 2티어 베리에이션으로 변경/1티어 베리에이션을 2티어 베리에이션으로 변경/2티어 베리에이션을 1티어 베리에이션으로 변경하는 것을 포함합니다.

2. **v2.product.update_tier_variation** API의 기능은 옵션 추가/삭제/업데이트를 포함합니다. 기존 베리에이션 정보를 유지해야 하는 경우 model_list 필드를 사용하십시오.

3. **v2.product.add_model** API의 기능은 새 모델 정보를 추가하는 것이고, **v2.product.delete_model** API의 기능은 베리에이션 정보를 삭제하는 것입니다.

4. CNSC 및 KRSC를 업그레이드한 판매자는 글로벌 제품을 통해서만 베리에이션을 추가/삭제할 수 있으므로 글로벌 제품 API를 사용하여 관리하십시오. 그렇지 않으면 오류가 보고됩니다.

## API 참조

| Global Product API | Product API |
|-------------------|-------------|
| v2.global_product.init_tier_variation | v2.product.init_tier_variation |
| v2.global_product.update_tier_variation | v2.product.update_tier_variation |
| v2.global_product.add_global_model | v2.product.add_model |
| v2.global_product.delete_global_model | v2.product.delete_model |

## 사용 사례

1. 기존 제품에 새로운 제품 배리언트 추가.
2. 더 이상 제공되지 않는 특정 제품 배리언트 삭제.
3. 제품 배리언트의 티어 구조 변경.
4. 제품 배리언트의 표시 순서 재정렬.
5. 배리언트 변경 후 제품 정보 업데이트.

## 관련 API

- v2_product_update_tier_variation
- v2_product_init_tier_variation
- product.update_tier_variation
- id-product-variant_http_product()
- id-product-variant_http_update()
- id.product.update_list_variation

---

## 원문 (English)

### Summary

This guide provides instructions on how to manage product variants, including adding, deleting, and reordering them. It covers scenarios where the tier structure remains the same and when it changes, providing API request examples for each case.

### Content

# Variant Management

## API Best Practice - Variant management

To view guidance about the structure of the product tree and variant structure in this article, please visit the [Catalog API Best Practices](link).

To manage the stock and price of a variation, please refer to the [Stock and Price Management](link) article.

---

## 1. Adding variants with same tier structure

Variants in the product tree already defined upon initialization (with tier_index and colors), when you need to add back for the variant situation:

### Original Variant structure:

| var_name | option | model_id |
|----------|--------|----------|
| var_sku[0][0] | red | 10000 |
| var_sku[0][1] (to be added) | black (to be added) | - |
| var_sku[1][0] | blue | 20000 |

**Call the** `v2_product_update_tier_variation` **API to add a blank option first**

**Request example:**

```json
{
  "tier_id": [REDACTED],
  "tier_variation": [
    {
      "name": "color",
      "option_list": [
        {
          "option": "red"
        },
        {
          "option": "black"
        },
        {
          "option": "blue"
        }
      ]
    },
    {
      "model_list": [
        {
          "tier_index": [0],
          "model_sku": "MAIN"
        }
      ]
    },
    {
      "tier_index": [1],
      "model_sku": "MAIN"
    }
  ]
}
```

---

**Call the** `v2_product_init_tier_variation` **API to add price, stock, SKU information for the blank variant**

**Request example:**

```json
{
  "tier_id": [REDACTED],
  "model_list": [
    {
      "tier_index": [0],
      "original_price": 20,
      "option_list": [
        {
          "option": "red"
        },
        {
          "option": "black"
        }
      ],
      "normal_stock": 100,
      "model_sku": "red_blacka"
    },
    {
      "price": 300
    }
  ]
}
```

---

## 2. Adding variants that change tier structure

When the product tree structure needs to be changed (the variant tier becomes 2-tier or other situations, you would need to add tier specifications, the steps include 1 and 2).

### Original Variant structure:

| var_name | option | model_id |
|----------|--------|----------|
| var_sku[0][0] | red | 10000 |
| var_sku[0][1] | blue | 20000 |

### Variant situation that change after:

| var_name | option | model_id |
|----------|--------|----------|
| var_sku[0][0] | red L | - |
| var_sku[0][1] | red XL | - |
| var_sku[1][0] | blue L | - |
| var_sku[1][1] | blue XL | - |

**Because the tier structure has been changed, the original model information will be removed after calling API, which will replace the original model at `10000` and model at `20000` will be invalid**

**API:** `v2_product_init_tier_variation`

**Request example:**

```json
{
  "tier_id": [REDACTED],
  "tier_variation": [
    {
      "option_list": [
        {
          "option": "red"
        }
      ]
    },
    {
      "option": "L"
    },
    {
      "option": "XL"
    },
    {
      "name": "Color"
    },
    {
      "name": "Size"
    },
    {
      "option_list": [
        {
          "option": "blue"
        }
      ]
    },
    {
      "option": "L"
    }
  ]
}
```

---

# Deleting Variants with Same Tier Structure

**Situation 1**: The product has defined color specifications (with colors red, blue, and black) and has the blue color model to be deleted.

## Original Variant Situation

| tier_index | option | model_id |
|------------|--------|----------|
| tier_index[0] | red | 10000 |
| tier_index[1] (% to be deleted) | blue (% to be deleted) | 20000 |
| tier_index[2] | black | 30000 |

**Variants situation:**

We need to delete the blue color, the tier_index can only be tier_index[0], tier_index[1] after deletion, so we need to convert the information of the variant tier_index[2] into the information of the black variant.

| tier_index | option | model_id |
|------------|--------|----------|
| tier_index[0] | red | 10000 |
| tier_index[1] | black | 30000 |

## API

Call `product.update_tier_variation` to remove option and override the model information at the same time (equipped linkage)

```json
{
  "tier_variation": [
    {
      "name": "color",
      "option_list": [
        {
          "option": "red"
        },
        {
          "option": "black"
        }
      ]
    }
  ],
  "model_list": [
    {
      "tier_index": [0],
      "model_id": 10000
    },
    {
      "tier_index": [1],
      "model_id": 30000
    }
  ]
}
```

---

# Situation 2

The product has defined color and size specifications, color is red, blue, and red, and all the blue variants need to be deleted.

## Original variant situation

| tier_index | option | model_id |
|------------|--------|----------|
| tier_index[0] | red 5L | 10000 |
| tier_index[1] | blue 5L (% to be deleted) | 20000 |
| tier_index[2] (% to be deleted) | black 5L | 30000 |

---

# Deleting Variants that Change Tier Structure

## Original Variant Situation

| tier_index | option | model_id |
|------------|--------|----------|
| tier_index[0] | null | 10000 |
| tier_index[1] | not XL | 25000 |
| tier_index[2] [5] (to be deleted) | blue-L (5) (to be deleted) | 30000 |
| tier_index[3] [1] (to be deleted) | blue-XL [2] (to be deleted) | 40000 |

**API:** `id-product-variant_http_product()` is deleted both option and model information at the same time.

**Request Example:**

```json
{
  "img_id": "[12345678]",
  
  "tier_var":[
    {
      "name": "color",
      "opt":[
        {
          "option": "red"
        }
      ]
    },
    {
      "name": "size",
      "opt":[
        {
          "option": "L"
        },
        {
          "option": "XL"
        }
      ]
    }
  ],
  
  "model_list": [
    {
      "tier_index": [0,0],
      "model_id": 10000
    },
    {
      "tier_index": [0,1],
      "model_id": 25000
    }
  ]
}
```

---

## Variants Situation After Deletion

| tier_index | option | model_id |
|------------|--------|----------|
| tier_index[0] | null | 10000 |
| tier_index[1] | not XL | 25000 |
| tier_index[2] | - | - |
| tier_index[3] | - | - |

---

## 4. Deleting Variants that Change Tier Structure

**Scenario 3:** The product has defined color and size specifications, color is red, blue, XL, and red and calory need to be deleted.

### Original Variant Situation

| tier_index | option | model_id |
|------------|--------|----------|
| tier_index[0] | null | 10000 |
| tier_index[1] | not XL | 25000 |
| tier_index[2] | blue-L | 30000 |
| tier_index[3] | blue XL | 40000 |

### Variant Situation After Deletion

| tier_index | option | model_id |
|------------|--------|----------|
| tier_index[0] | - | - |
| tier_index[1] | - | - |

**API:** `id-product-variant_http_update()`

**Request Example:**

```json
{
  "img_id": "[12345678]",
  
  "tier_var":[
    {
      "opt":[
        {
          "option": "L"
        },
        {
          "option": "XL"
        }
      ],
      "name": "size"
    },
    {
      "name": [
        {
          "id-productpg": [0],
          "model_id": "bnpL1",
          "original_price": 100,
          "tier_index": [0]
        },
        {
          "id-productpg": [0],
          "model_id": "bnpL1",
          "original_price": 100,
          "tier_index": [1]
        }
      ]
    }
  ]
}
```

---

# Changing the order of variants

## Scenario
If the item has defined color specification with colors red and blue, then the color order needs to be changed to blue, red, then the original variant information (ordered).

## Original variant situation

| var_active | option | model_id |
|------------|--------|----------|
| var_active(1) | red | 10000 |
| var_active(2) | blue | 20000 |

## Variant situation after change:

| var_active | option | model_id |
|------------|--------|----------|
| var_active | option | model_id |
| var_active(2) | blue | 20000 |
| var_active(1) | red | 10000 |

## API

**id.product.update_list_variation**

Request example:

```json
{
  "link_id": "RRDDCCFFVV",
  "list_variants": [
    {
      "name": "Color",
      "options_list": [
        {
          "variant": "red A"
        },
        {
          "variant": "Blue"
        }
      ]
    }
  ],
  "model_list": [
    {
      "tier_index": [0],
      "model_id": 20000
    },
    {
      "tier_index": [1],
      "model_id": 10000
    }
  ]
}
```

---

# Changing the option name of the variant

## Scenario
If the item has defined a color specification with the colors red and blue. Now the color name needs to be changed to red A, blue.

## Original variant situation

| var_active | option | model_id |
|------------|--------|----------|
| var_active | option | model_id |
| var_active(1) | red | 10000 |
| var_active(2) | Blue | 20000 |

## Variant situation after change:

| var_active | option | model_id |
|------------|--------|----------|
| var_active | option | model_id |
| var_active(2) | red A | 10000 |
| var_active(1) | Blue | 20000 |

## API

**id.product.update_list_variation**

Request example:

```json
{
  "link_id": "RRDDCCFFVV",
  "list_variants": [
    {
      "name": "Color",
      "options_list": [
        {
          "variant": "red A"
        },
        {
          "variant": "Blue"
        }
      ]
    }
  ],
  "model_list": [
    {
      "tier_index": [0],
      "model_id": 10000
    },
    {
      "tier_index": [1],
      "model_id": 20000
    }
  ]
}
```

---

# API Documentation

## JSON Structure

```json
{
  "tier_index": [1],
  "model_id": 20000
}
]
}
```

## Summary

1. The function of **v2.product.init_tier_variation** API, including changing 0-tier variation to 1-tier variation / 0-tier variation to 2-tier variation / 1-tier variation to 2-tier variation / 2-tier variation to 1-tier variation.

2. The function of **v2.product.update_tier_variation** API, including add/delete/update the options, If you still need to keep the original variants information, please use the model_list field.

3. The function of **v2.product.add_model** API is adding the new model information, the function of **v2.product.delete_model** API is deleting the variant information.

4. Sellers who have upgraded CNSC and KRSC can only add/delete variants through global products, so please use the global product API for management, otherwise an error will be reported.

## API Reference

| Global Product API | Product API |
|-------------------|-------------|
| v2.global_product.init_tier_variation | v2.product.init_tier_variation |
| v2.global_product.update_tier_variation | v2.product.update_tier_variation |
| v2.global_product.add_global_model | v2.product.add_model |
| v2.global_product.delete_global_model | v2.product.delete_model |

---

**문서 ID**: developer-guide.219
**플랫폼**: shopee
**URL**: https://open.shopee.com/developer-guide/219
**처리 완료**: 2025-10-16T08:45:21
