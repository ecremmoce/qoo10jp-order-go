# API 모범 사례 - 글로벌 상품 생성

**카테고리**: 모범 사례
**난이도**: medium
**중요도**: 4/5
**최종 업데이트**: 2025-10-16T08:42:17

## 개요

본 가이드는 통합된 글로벌 상품 카탈로그와 상점 상품을 동기화하는 데 중점을 두고 글로벌 상품을 생성하고 관리하기 위한 모범 사례를 설명합니다. 용어, 상품 관리 변경 사항, 카테고리 및 상품 정보에 대한 API 사용법을 다룹니다.

## 주요 키워드

- Global Product
- Shop Product
- CMSC
- IMSC
- MTPSH
- MPSKU
- Product Synchronization
- API
- Product Categories
- Product Attributes

## 본문

# API 모범 사례 - 글로벌 상품 생성

## 개요

쇼핑몰 상품이 생성된 후, 해당 쇼핑몰 상품을 글로벌 상품에 매핑하여 중국어 또는 영어로 통합된 상품명 및 설명을 표시하고, 쇼핑몰 상품에 대한 업데이트를 동기화하도록 설정해야 합니다. Shopper는 자동으로 현지 언어로 번역합니다.

- 쇼핑몰 상품에 대한 비동기 업데이트 실시간 업데이트를 지원합니다. 일부 실시간 동기화된 상품 정보의 경우, 판매자는 비동기 백그라운드 작업 실행 상태를 확인하여 쇼핑몰 상품 정보를 확인할 수 있습니다.
- 상품 정보 동기화 판매자는 동기화 규칙을 구성할 수 있습니다. 예를 들어, 상품 정보 업데이트의 우선 순위입니다. 예를 들어, 공식 계산 결과에 따라 가장 높은 우선 순위를 가진 글로벌 상품입니다.

**참고:** 쇼핑몰 상품이 Global API와 통합되고 쇼핑몰 상품이 `Global Product API`와 같은 Merchant API와 통합되는 경우, 쇼핑몰 사용자는 글로벌 상품으로 상품 정보를 관리해야 합니다. 주문 API와 같은 다른 API의 경우 변경이 필요합니다.

- CMS에 대한 자세한 내용은 `CMSC API`를 사용하세요.
- IMSC에 대한 자세한 내용은 `여기`를 클릭하세요.

### CMSC/IMSC 계정 시스템

이 통합을 통해 CMSC 또는 IMSC에 판매자 계정이 생성되면, 크로스보더 판매자에게만 적용됩니다.

---

## 용어 설명

- **메인 계정**: 가장 높은 권한을 가진 기본 계정으로, 계정 소유자는 일반적으로 법인이며, 판매자의 모든 하위 계정을 생성하고 추가할 수 있는 권한을 가집니다. 메인 계정의 로그인 형식은 ACC_main입니다.
- **쇼핑몰**: 사용자가 쇼핑몰 계정을 통해 찾는 페이지입니다. 여러 마켓에 대한 쇼핑몰은 하위 Merchant를 사용합니다. 하나의 CMSC 메인 계정은 여러 쇼핑몰을 가질 수 있습니다.
- **아이템**: 쇼핑몰을 CMSC/IMSC 계정으로 업그레이드하면, 인증을 위한 기본 계정을 얻게 되며, 메인 계정만 API 호출을 할 수 있습니다. (판매자가 자체 계정을 선택하면 메인 계정과 하위 계정 모두 API 호출을 지원하며, 여러 Merchant에 사용되며, 각 Merchant에는 SPS 쇼핑몰 Merchant를 포함한 여러 마켓이 있을 수 있으며, Merchant 아래에는 동일한 마켓의 여러 쇼핑몰이 있을 수 있습니다.)

---

## 글로벌 상품 및 쇼핑몰 상품

다음에서는 상품과 쇼핑몰 상품 간의 관계를 설명합니다.

```
Merchant 1

SHIPBU - MTPSH 11

└── US
└── MX
└── BR

└── MTPSH 2
    └── MTPSH 3
    └── MTPSH 4
```

| 약어 | CMSC/IMSC | 설명 |
|--------------|-----------|-------------|
| MTPSH | 글로벌 SKU | 판매자는 상품명, 상품 설명, 이미지 등 기본적인 상품 정보를 MTPSH로 관리합니다. 차별화를 위해 마켓에 게시한 후, 쇼핑몰 SKU가 생성되며, 판매자는 쇼핑몰 상품 변경에 대한 마켓별 정보를 관리할 수 있습니다. |
| MPSKU | 쇼핑몰 SKU | 쇼핑몰의 실제 쇼핑몰 상품 |

---

## 상품 관리 프로세스의 변경 사항

글로벌 상품 관리 프로세스를 도입한 후, 상품 관리 프로세스에 변경 사항이 있습니다.

### 1. 상품 생성

**온라인 상품 판매 장치 API**를 통해 새로운 상품 생성 프로세스를 진행한 다음, API를 사용하여 `product/createProduct`를 호출하여 더 많은 상품 카테고리의 쇼핑몰 상품과 쇼핑몰 상품을 게시하고 동기화합니다. 게시 프로세스에 대한 자세한 내용은 [링크]를 참조하세요.

### 2. 필드 관리

상품 관련 필드를 세 가지 범주로 나눕니다.

1. **MTPSH에서만 관리할 수 있는 필드**: 이러한 필드는 카테고리, 속성 등과 같은 차이점을 만들 수 없습니다. MTPSH에서 관리할 수 있는 필드는 둘 다 업데이트할 수 있습니다. 모든 글로벌 필드는 모든 쇼핑몰 상품으로 업데이트할 수 있지만, MTPSH가 판매자의 공개 관리 필드인 경우입니다. 이러한 필드는 상품명 및 상품 설명과 같이 (선택)적으로 차별화할 수 있습니다. 필드에서 MTPSH를 통해 필드를 선택해야 하는 경우, 계정은 자동으로 쇼핑몰 상품으로 이전됩니다. 선택하지 않으면 계정은 동기화 규칙을 구성하고 우선 순위는 글로벌 상품에 따라 업데이트됩니다.

2. **MPSKU에서만 관리할 수 있는 필드**: 이러한 필드는 물류 정보와 같이 글로벌 상품을 통해서만 관리할 수 있습니다.

| API 필드 | MTPSH에서 관리 | MPSKU에서 관리 |
|-----------|----------------|-----------------|
| category_id | ✓ | ✗ |
| attribute | ✓ | ✗ |
| brand | ✓ | ✓ |
| item_name | ✓ | ✓ |
| description | ✓ | ✓ |
| weight | ✓ | ✗ |
| dimension | ✓ | ✗ |
| price | ✗ | ✓ |
| stock | ✗ | ✓ |
| image/video | ✓ | ✓ |
| days_to_ship | ✓ | ✓ |
| item_status | ✓ | ✓ |
| condition | ✓ | ✓ |
| item sku | ✗ | ✓ |
| size_variation structure | ✓ | ✗ |
| size_variation name/caption | ✓ | ✓ |
| logistics | ✗ | ✓ |

---

### 다음 사항에 유의하십시오.

- 쇼핑몰을 CMSC/IMSC로 업그레이드하면 MTPSH의 제한은 5k이고 MTPSH에서 매일 게시되는 MPSKU는 2en MPSKU입니다. MTPSH의 제한은 MTPSH가 1K인 경우입니다. 기본 게시 크기 MTPSH가 2en K인 경우 MTPSH의 재고는 2en 차이점을 사용할 수 있습니다.

- 국가마다 재고가 다릅니다. 글로벌 게시물이 게시된 모든 쇼핑몰 상품은 다릅니다. 그러나 쇼핑몰 상품은 글로벌 상품에 영향을 주지 않고 개별적으로 삭제하고 통합할 수 있습니다.

---

## 3. 카테고리 호출

### 상품 목록

**글로벌 상품 생성:** Shopper 크로스보더 관리에 처음 연결하는 경우, **Product standard information API** 또는 **Global product sale API**를 통해 발행하는 것이 좋습니다. 자세한 내용은 [링크]를 참조하세요.

#### 3.1 카테고리 호출

`/api/v1/global_product/category`를 사용합니다.

**다음 사항에 유의하십시오.**
- 카테고리가 다른 경우 API는 동일한 카테고리 트리 데이터를 반환합니다.
- 카테고리가 다른 경우 동일한 응답 구조를 반환하더라도 카테고리 속성 및 정보(카테고리 A)는 말레이시아 쇼핑몰에서는 사용할 수 없지만 브라질 쇼핑몰에서는 사용할 수 있는 경우, 카테고리를 게시하지 않도록 안내하지만 글로벌 상품은 업데이트할 수 있습니다. 카테고리 정보에 대한 자세한 내용은 로컬에서 사용할 수 있는 API 호출을 참조하십시오.

#### 3.2 추천 카테고리 호출

```
/api/v1/global_product/get_recommend_category
```

---

# 개발자 가이드 - 카테고리 및 상품 정보 가져오기

## 3.1 카테고리 가져오기

API: `tv.taobao.products.get_category`

다음 사항에 유의하십시오.
- 판매자는 동일한 카테고리를 너무 자주 선택할 수 없습니다.
- `tv.taobao.products.get_category` API는 글로벌 상품에 사용할 수 있는 카테고리를 반환하지만, 인증된 API가 상품을 성공적으로 생성하기에 충분하지 않은 상황이 발생할 수 있습니다. 이 경우 Aliexpress PD에 최신 사용 가능한 카테고리를 확인해야 합니다. 또한 Seller Center를 통해 상품을 게시하여 각 마켓에 사용할 수 있는 카테고리를 알 수 있으며, `tv.product.get_category` API를 호출할 수 있습니다.

## 3.2 속성 가져오기

API: `tv.global_product.category_properlised`

## 3.3 속성 값 가져오기

API: `tv.global_product.get_attribute_traits`

다음 사항에 유의하십시오.
속성 이름, 필수, 요청/매개변수 및 필수, 지역 응답, 매개변수와 같은 일부 속성의 경우, 예를 들어 맛과 같은 일부 속성의 경우 대만 마켓의 상품만 필수이므로 판매자가 대만 마켓에 게시하지 않으려면 무시할 수 있으며 글로벌 상품 생성 프로세스를 확인할 필요가 없으므로 속성 목록에 표시됩니다. 이 경우 무시하십시오. 실제 상황은 해당 값이 필요한지 여부에 따라 다릅니다. 이러한 나머지 속성의 경우:
- 목록 포인트의 특수 자동 기능을 제외하고 마켓 중 하나가 값을 요청하는 한 모든 속성이 반환되고 attribute_name을 채워야 하며 tv.product_product.get_attribute_traits API를 사용해야 합니다.

## 3.4 추천 속성 가져오기

API: `tv.global_product.get_ai_attribute_traits`

## 3.5 사용 가능한 브랜드 목록 가져오기

API: `tv.global_product.get_brand_list`

## 3.6 배송 일수 가져오기

API: `tv.global_product.get_day_to_ship`

다음 사항에 유의하십시오.
`tv.global_product.get_day_to_ship`은 모든 마켓에서 채울 수 있는 배송 일수를 연결하여 반환합니다.

```json
"data_is_ship_range_list": [
  {
    "country": "SG",
    "ship": "15"
  },
  {
    "country": "MY",
    "ship": "15"
  }
]
```

일괄 처리 사례: 글로벌 상품에 대한 일괄 배송 일수는 7, 8, 10, 11, 12 및 13입니다.
`tv.product.get_day_to_ship`의 풍부한 테이블을 저장하여 각 마켓에 대한 더 정확한 배송 일수를 얻을 수도 있습니다.

## 3.7 카테고리에 대한 정보 가져오기

API: `tv.global_product.get_ai_attribute_traits`

다음 사항에 유의하십시오.
자세한 내용은 `announcement`를 검토하는 최신 화이트리스트를 설정하고 있습니다.

## 3.8 글로벌 상품 제한 가져오기

API: `tv.global_product.category_rules_get`

다음 사항에 유의하십시오.
현재 `category_name_product_base` 모든 PGU 상품 제한/금지 이름 및 옵션 이름에 대해 금지된 정보가 포함되어 있으며 가져올 수 있습니다. 중국어와 영어 번역이 정확히 동일하지 않으므로 영어 번역, 마켓 언어와 같은 정보는 중국어와 영어가 다른 중국어 길이 제한을 갖도록 번역을 사용하고 다른 언어는 중국어 문자와 동기화됩니다. (중국어 문자 제한은 30이고 다른 언어는 중국어 문자 제한 60과 동기화됩니다.)

```json
"cnt_length_restriction": 2.44,
"cnt_length_rules": "length" = 240,
"eng_length_rules": 2.5 * 
    cnt_length,
"local_limit": 500
```

다시 말해, 판매자가 상품명을 영어로 입력하면 최대 문자 수는 1이 500입니다. 판매자가 상품명을 중국어로 입력하면 최대 문자 수는 1이 250입니다(500/2 = 44 = 250.50 문자이기 때문입니다).

---

## 4. 글로벌 상품 생성

글로벌 상품 구성은 글로벌 상품을 생성하기 전에 특정 요구 사항 또는 규칙을 충족해야 합니다.
판매자는 글로벌 상품을 생성하기 전에 CNSC/MNSC에 로그인하여 온보딩 및 환율 구성을 완료해야 합니다. 그렇지 않으면 `unauthorized` 등과 같은 오류가 발생합니다.

- 온보딩 절차는 https://shopee.ph/edu/article/1741을 참조하십시오.
- 통화 설정: https://seller.shopee.ph/account/settings/address

---

## 4.2 메인 계정 인증 완료

`Authorization/UserAuthorization`을 참조하십시오.

## 4.3 메인 계정 아래의 Merchant ID 목록 및 쇼핑몰 ID 가져오기

다음 API를 호출할 수 있습니다.
- 인증 후 `tv.seller.get_authorized_shop` API를 통해 현재 인증된 Merchant ID 목록 및 쇼핑몰 ID 목록을 가져옵니다.
- API를 호출하여 인증을 완료하기 전에 API를 호출하여 `tv.shop.get_shop_list` API를 호출하여 ID와 기본 쇼핑몰 및 계열 쇼핑몰 간의 관계를 포함하여 각 인증된 쇼핑몰의 정보를 가져옵니다.
- 기본 쇼핑몰 정보를 가져오려면 `tv.shop.get_shop_info` API를 호출하여 `s1_main_shop` 필드가 반환되지 않았는지 확인합니다.
- SIP 기본 쇼핑몰인 경우 `s1_gm1test.test.the_sip_info` 쇼핑몰 필드가 반환됩니다.
- 메인 계정이 여러 `s1_gm1test.test.the_sip_info` 쇼핑몰 필드를 인증한 경우 반환되지 않습니다.

---

## 4.4 상품 번역 정보 동기화

번역 지원:
판매자는 CNSC/MNSC 또는 API를 통해 사전에 번역된 정보가 쇼핑몰 상품에 자동으로 동기화되도록 할 수 있습니다.

CNSC/MNSC 스크린샷:

**Merchant 설정**

[설정 유형, 다양한 상태 필드 및 설명에 대한 열이 있는 Merchant 설정을 보여주는 테이블]

**Merchant 설정** 섹션에는 다음 옵션이 표시됩니다.
- Taobao 지역 상품 정보 동기화
- 다음 설정을 표시합니다.
  - 동기화 옵션
  - 동기화됨 ✓
  - 동기화되지 않음 ✓
  - 카테고리(EN/CN 표시기 포함)
  - 색상(빨강, 파랑, 흰색)
  
*sync_ authorized는 동기화되지 않음을 의미해야 합니다.*

키 필드는 `tv.global_product.pre_trans_sync.txt`이며, *true*는 동기화를 의미하고 *false*는 동기화되지 않음을 의미합니다.

---

## 4.5 이미지 업로드

## 4.5.1 미디어 파일 업로드

Shopee의 미디어 공간에 미리 저장된 파일을 생성하거나 편집하기 위해 API를 호출하기 전에 media_upload 관련 API를 호출하여 파일을 업로드해야 합니다. 자세한 내용은 "미디어 파일 업로드" 섹션의 `show`를 참조하십시오.

CNSC/MNSC 미디어 공간에서 상품 이미지 또는 속성 이미지를 사용하려면 `tv.global_product.get_image_list` API를 호출하여 지정된 카테고리 아래의 메인 이미지 목록을 가져올 수 있습니다.

---

## 5) 카테고리별 업로드

판매자가 `tv.global_product.get_category` API에서 `has_children=false`를 반환하는 카테고리만 선택할 수 있도록 카테고리별로 업로드하는 것을 지원합니다.

---

## 6) 번역 요구 사항

해외 Aliexpress Merchant에게 최상의 서비스를 제공하기 위해 다양한 상품 유형(이미지, 텍스트, 상품 설명 등)에 대한 엄격한 규칙을 적용하여 비즈니스가 올바르게 운영되도록 할 것입니다. 판매자는 간체 중국어, 영어 및 숫자를 입력할 수 있습니다. 간체 중국어와 번체 중국어를 혼합하지 마십시오. 현지 언어로 번역되지 않는 현지 언어 단어를 충족하는 경우. 다음 요구 사항은 필수입니다.

번역 정보:
크로스보더 상품 설명은 간체 중국어, 영어 및 숫자만 입력할 수 있습니다.

제한 정보:
금지된 콘텐츠: 외설 및 민감한 단어에 대해서는 간체 중국어, 영어 및 숫자만 입력할 수 있습니다.

---

## 4.6 글로벌 상품 정보 업로드

글로벌 상품 정보를 업로드할 때 다음 사항에 유의하십시오.
현재 미완료된 판매자가 설명을 위해 쇼핑몰을 선택하는 것을 지원합니다. 설명 부분에 이미지를 추가하거나 업로드하는 방법에 대한 자세한 내용은 `FAQ`를 참조하십시오.

이미지 업로드 정보:
크로스보더 Seller Center 판매자는 이제 가격 책정을 위해 모든 UPC를 만들 수 있습니다. 크로스보더 판매자는 가격 책정에 USD를 사용할 수 있으며, 상품 가격 없이 `tv.global_product.product_add` API 항목의 가격을 수정할 수 있습니다.

가격 정보:
단일 유형 상품의 경우 가격은 환율로 직접 변환되지 않습니다.

---

## 4.6.1 글로벌 상품 재고 업로드

API: `tv.global_product.product_add`

---

# 4.4.3 글로벌 상품 설명 업로드

**훌륭한 상품 설명은 간체 중국어, 영어 및 숫자로 구매자를 유치합니다.**

## 중요 참고 사항
현재 이미지를 설명에 삽입하기 위해 임베디드 비디오를 지원합니다. 비디오를 삽입하는 방법에 대한 자세한 내용은 [FAQ](#)를 참조하십시오.

---

## 4.4.3 글로벌 상품 가격 업로드

해외 크로스보더 판매자는 가격 책정에 RMB 또는 USD를 사용할 수 있습니다. 해외 크로스보더 판매자는 가격 책정에 USD를 사용할 수 있으며, 국내 크로스보더 판매자는 가격 책정에 RMB를 사용할 수 있습니다. [YJ 변환율 규칙/일반 비율](#) API의 경우 판매자가 선택한 통화에 따라 해당 통화의 다른 가격 단위가 있습니다.

---

## 4.4.1 글로벌 상품 재고 업로드

업로드된 재고는 위험 통제 검토를 거칩니다. 검토에 통과하지 못하면 업로드된 재고의 일정 비율이 동결됩니다.

예: MT3501이 3개의 MPU(s)를 게시하는 경우 MPU(s8, MPU(s8, MPU(s8)이면 MT3501 재고가 100으로 설정되면 재고 MPU(s1+MPU(s2+MPU(s3)=120이면 [MPU(s1, MPU(s2 및 MPU(s3은 해당
| tier_index[2] | M | 300 | 10 | - |

---

## API 요청 예시:

**Endpoint:** `sz.global_product_tier_tier_variation API request example:`

```json
{
  "global_item_id": 1000170002,
  
  "tier_variation": [
    {
      "name": "size",
      
      "option_list": [
        {
          "option": "S",
          
          "image": {
            "image_id": "030ce4fd59c82e09a6bc78c8c3e7f"
          }
        },
        {
          "option": "L",
          
          "image": {
            "image_id": "19bce4fd59c82e09a6bcf78c8c7f"
          }
        },
        {
          "option": "M"
        }
      ]
    }
  ],
  
  "model_list": [
    {
      "tier_index": [0],
      
      "original_price": 100,
      
      "global_model_sku": "sku1",
      
      "model_stock": 10
    },
    {
      "tier_index": [1],
      
      "original_price": 200,
      
      "global_model_sku": "sku2",
      
      "model_stock": 20
    },
    {
      "tier_index": [2],
      
      "original_price": 300,
      
      "global_model_sku": "",
      
      "model_stock": 30
    }
  ]
}
```

---

## 4.8.2 2단계 베리에이션 생성

상품에 색상 및 사이즈 사양이 있다고 가정합니다. 색상은 빨간색과 파란색을 포함하고, 사이즈는 xl, s, l 쿠키를 포함합니다.

### 예시 테이블:

| tier_index | color | size | price | stock | sku |
|------------|-------|------|-------|-------|-----|
| tier_index[0][0] | Red | S | 100 | 10 | sku1 |
| tier_index[0][1] | Red | L | 200 | 20 | sku2 |
| tier_index[1][0] | Blue | XL | 500 | 30 | sku3 |
| tier_index[1][1] | Blue | L | 400 | 40 | sku4 |

---

## API 요청 예시:

**Endpoint:** `sz.global_product_tier_tier_variation API request example:`

```json
{
  "global_item_id": 1000170002,
  
  "tier_variation": [
    {
      "name": "Color",
      
      "option_list": [
        {
          "option": "Red"
        },
        {
          "option": "Blue",
          
          "image": {
            "image_id": "030ce4fd59c82e09a6bc78c8c3e7f"
          }
        }
      ]
    },
    {
      "name": "Size",
      
      "option_list": [
        {
          "option": "S"
        },
        {
          "option": "XL"
        }
      ]
    }
  ]
}
```

---

# 개발자 가이드 - 상품 베리에이션 구성

## JSON 구조

```json
]
},
{
  "name": "size",
  
  "option_list": [
    {
      "option": "XL"
    },
    {
      "option": "L"
    }
  ]
}
],

"model": [
  {
    "tier_index": [0,0],
    
    "original_price": 100,
    
    "normal_stock": 10,
    
    "global_model_sku": "sku1"
  },
  {
    "tier_index": [0,1],
    
    "original_price": 200,
    
    "normal_stock": 20,
    
    "global_model_sku": "sku2"
  },
  {
    "tier_index": [1,0],
    
    "original_price": 300,
    
    "normal_stock": 30,
    
    "global_model_sku": "sku3"
  },
  {
    "tier_index": [1,1],
    
    "original_price": 400,
    
    "normal_stock": 40,
    
    "global_model_sku": "sku4"
  } ]
}
```

## 중요 사항

**다음 사항에 유의하십시오:**

- **정의한 옵션은 Shopee 몰 측에 순서대로 표시됩니다. Shopee는 현재 최대 2단계 베리에이션 정의만 지원합니다.**

- **각 베리에이션에 대한 이미지를 정의할 수 있습니다. 2단계 베리에이션 상품인 경우 첫 번째 레이어의 옵션만 정의할 수 있습니다. 즉, 예시에서 색상을 기준으로 베리에이션 이미지를 정의할 수 있지만 크기를 기준으로 정의할 수는 없습니다. 베리에이션 이미지를 추가하려면 첫 번째 레이어의 모든 옵션이 이미지를 정의해야 합니다. 이미지는 V2 미디어 업로드 이미지 API를 호출하여 업로드해야 합니다.**

- **tier_index는 0부터 시작해야 하며 오버플로되지 않아야 합니다. 그렇지 않으면 오류가 보고됩니다.**

- **아이템 데이터 생성에 지연이 있을 수 있으므로 아이템 생성 후 5초 간격으로 베리에이션을 생성하는 것이 좋습니다.**

## 다음 단계

다음으로 **글로벌 상품 게시 방법**을 검토하십시오.

## 사용 사례

1. 새로운 글로벌 상품 카탈로그 생성
2. 여러 상점 및 시장에서 상품 정보 동기화
3. 글로벌 상품 속성 및 카테고리 관리
4. CMSC/IMSC 계정 시스템과 통합
5. 실시간 상품 정보 업데이트

## 관련 API

- Global Product API
- Product standard information API
- Global product sale API
- /api/v1/global_product/category
- /api/v1/global_product/get_recommend_category
- tv.taobao.products.get_category
- tv.global_product.category_properlised
- tv.global_product.get_attribute_traits
- tv.global_product.get_ai_attribute_traits
- tv.global_product.get_brand_list
- tv.global_product.get_day_to_ship

---

## 원문 (English)

### Summary

This guide outlines the best practices for creating and managing global products, focusing on synchronizing shop products with a unified global product catalog. It covers terminology, product management changes, and API usage for categories and product information.

### Content

# API Best Practices - Creating global product

## Overview

After the shop product is created, the shop product needs to be mapped to a global product to display its name and description in a unified product in Chinese or English, and set up to synchronize the update to the shop products. Shopper will automatically translate into the local language.

- Support asynchronous update real-time updates for shop products. For some real-time synchronized product information, sellers can watch the asynchronous background task execution status to see shop product information.
- The product information synchronization sellers can configure the synchronization rules. For example, the priority of product information updates. for example, the global product with the highest priority according to the results of the formula calculation.

**Note:** When shop products integrate with Global API and shop products integrate to Merchant API such as `Global Product API`, shoppers you need to manage product information with global products. For other APIs such as order API, the changes are needed.

- Use `CMSC API` to learn more about CMS.
- Click `here` to learn more about IMSC.

### The CMSC/IMSC account system

When a seller account is created in CMSC or IMSC through this integration, Only applicable to cross-border sellers.

---

## Terminology Description

- **Main account**: The basic account with highest authority, the account owner is usually a corporate entity, and has permission to create and add all sub-accounts of the seller. The login format of the main account is ACC_main.
- **Shop**: The pages users find through shop's account. Shops to multiple markets use sub Merchant. A CMSC main account may have multiple shops affiliated.
- **Item**: When you upgrade the shops to CMSC/IMSC accounts, You will get the basic account for authorization, and only the main account is permitted to make API calls. (Both the main account and sub-accounts support API calls if the seller chooses own account, then use for multiple Merchants, andunder each Merchant, there can be multiple market including SPS shop Merchant, there may be multiple shops of the same market under Merchant.

---

## Global products and shop products

In the following, we will explain the relationship with product and shop product.

```
Merchant 1

SHIPBU - MTPSH 11

└── US
└── MX
└── BR

└── MTPSH 2
    └── MTPSH 3
    └── MTPSH 4
```

| Abbreviation | CMSC/IMSC | Description |
|--------------|-----------|-------------|
| MTPSH | Global SKU | Sellers will manage some basic product information by MTPSH, including product name, product description, images, etc. After publishing to differentiate, markets, shop SKU will be created, and seller can manage market-specific information for the changes to shop products |
| MPSKU | Shop SKU | Real shop products from shop |

---

## Changes in product management process

After the adoption of the global product management process, there are changes to the product management process.

### 1. Product creation

New product created process through **an online product sale device API**, and then call `product/createProduct` with an API to publish and synchronize shop products with shop products from more product category. For more details about the publishing process use: [link]

### 2. Field management

We divide the product-related fields into three categories:

1. **Fields can only be managed by MTPSH**: These fields can't be created differences, such as categories, attributes, etc. Fields can be managed by MTPSH can be updated by both. All global fields can update to all shop products, but if MTPSH is a seller's public managed fields. These fields have (select) differentiatable, such as product names and product descriptions. If fields choose to require the field select through the MTPSH, the account will transfer to the shop products automatically. If you don't select it, account configure synchronization rules, the priority will update according to the global product.

2. **Fields can only be managed by MPSKU**: These fields can only be managed through the global product, like logistics info.

| API Field | Manage by MTPSH | Manage by MPSKU |
|-----------|----------------|-----------------|
| category_id | ✓ | ✗ |
| attribute | ✓ | ✗ |
| brand | ✓ | ✓ |
| item_name | ✓ | ✓ |
| description | ✓ | ✓ |
| weight | ✓ | ✗ |
| dimension | ✓ | ✗ |
| price | ✗ | ✓ |
| stock | ✗ | ✓ |
| image/video | ✓ | ✓ |
| days_to_ship | ✓ | ✓ |
| item_status | ✓ | ✓ |
| condition | ✓ | ✓ |
| item sku | ✗ | ✓ |
| size_variation structure | ✓ | ✗ |
| size_variation name/caption | ✓ | ✓ |
| logistics | ✗ | ✓ |

---

### Please note that:

- When you upgrade the shops to CMSC/IMSC, the limit of MTPSH is 5k and MTPSH daily published MPSKU is 2en MPSKU. From the limit of MTPSH is MTPSH is 1K. When the base published size MTPSH is 2en K, the stock of MTPSH is 2en The available of difference.

- Different countries have the stock - all shop products for which the global post has been published are different. However, shop products can be deleted and unifre individually without affecting global products.

---

## 3. Calling categories

### Product listing

**Creating global product:** If you first connect to Shopper cross-border management, it is recommended to issue through either **Product standard information API**, or **Global product sale API**. More information can be found: [link]

#### 3.1 Calling categories

Use `/api/v1/global_product/category` 

**Please note that:**
- For different categories, the API returns the same category tree data
- Even if different categories return the same response structure, the category attribute and information (category A) is not available in Malaysia shop, but available in Brazil shop, we will prompt you not to publish the category, but can update the global products. For more details about categories information, see API call for the locally available.

#### 3.2 Calling recommended categories

```
/api/v1/global_product/get_recommend_category
```

---

# Developer Guide - Getting Categories and Product Information

## 3.1 Getting categories

API: `tv.taobao.products.get_category`

Please note that:
- Sellers are not allowed to select the same category too often.
- `tv.taobao.products.get_category` API will return the available category for global product, but there may be a situation that authorized API is not enough to create the product successfully, in which case you need to confirm with the Aliexpress PD for the latest available category. You can also publish products through Seller Center to know the available category for each market, you can call `tv.product.get_category` API.

## 3.2 Getting attribute

API: `tv.global_product.category_properlised`

## 3.3 Getting attribute value

API: `tv.global_product.get_attribute_traits`

Please note that:
Some attributes such as attribute name, required, request/parameter and mandatory, region response, parameter, because for some attributes, such as flavor, different only products in the Taiwan market are compulsory, if the seller doesn't want to publish in Taiwan market, he can ignore it, don't need to check the process of global product creation, so the attribute will be visible in the attribute list, in case, please ignore them. The actual situation depends on whether the corresponding value is required. For these remain attributes:
- Except for the special automatic features in the list point, all other attributes, as long as one of the markets request a value, are returned, the attribute_name should be filled, and the tv.product_product.get_attribute_traits API.

## 3.4 Getting recommendation attribute

API: `tv.global_product.get_ai_attribute_traits`

## 3.5 Getting the available brand list

API: `tv.global_product.get_brand_list`

## 3.6 Getting days to ship

API: `tv.global_product.get_day_to_ship`

Please note that:
`tv.global_product.get_day_to_ship` will return a concatenation of the number of shipping days that can be filled in all markets.

```json
"data_is_ship_range_list": [
  {
    "country": "SG",
    "ship": "15"
  },
  {
    "country": "MY",
    "ship": "15"
  }
]
```

Batch case: For batch shipping days for global products are 7, 8, 10, 11, 12, and 13.
You can also save the rich table of `tv.product.get_day_to_ship` to get more precise shipping days for each market.

## 3.7 Getting information about categories

API: `tv.global_product.get_ai_attribute_traits`

Please note that:
We are setting out the latest white list which review the `announcement` for details.

## 3.8 Getting global product restrictions

API: `tv.global_product.category_rules_get`

Please note that:
At present, `category_name_product_base` For All PGU product restrictions/prohibitions name and option name, we have included prohibited information, you can get. Since the Chinese and English translations are not exactly the same, information like the English translation, the market language, so in order to Chinese and English have different Chinese length limits, use translation, and other languages will be synchronized to the Chinese character. (Chinese character limit is 30, other languages will be synchronized to the Chinese character limit of 60)

```json
"cnt_length_restriction": 2.44,
"cnt_length_rules": "length" = 240,
"eng_length_rules": 2.5 * 
    cnt_length,
"local_limit": 500
```

In other words, if the seller fills in the product name in English, the maximum number of characters is 1 is 500. If the seller fills the product name in Chinese, the maximum number of characters is 1 is 250 (because 500/2 = 44 = 250.50 characters).

---

## 4. Creating global product

The configuration of global products shall meet certain requirements or rules before creating global products.
Sellers need to log in to CNSC/MNSC to complete the configuration of onboarding and exchange rate before creating global products, otherwise you will encounter errors like `unauthorized` and etc.

- For the onboarding procedure, please refer to: https://shopee.ph/edu/article/1741
- Set the currency : https://seller.shopee.ph/account/settings/address

---

## 4.2 Completing the main account authorization

please refer to `Authorization/UserAuthorization`

## 4.3 Getting the merchant id list and shop id under the main account

The following APIs can be called:
- To get a list of current authorized merchant id list and shop id through `tv.seller.get_authorized_shop` API after authorization.
- Before calling the API to complete the authorization, call API to `tv.shop.get_shop_list` API to get the information of each authorized shop, including the relationship between id and primary shop and affiliated shops.
- To get primary shop information, call `tv.shop.get_shop_info` API to get the `s1_main_shop` field is not returned.
- If it is a SIP primary shop, `s1_gm1test.test.the_sip_info` Shop field is returned.
- For the main account authorized multiple `s1_gm1test.test.the_sip_info` Shop field is not returned.

---

## 4.4 Product translation information synchronization

Translation support:
Sellers can get beforehand translation information was automatically synchronized to shop products through CNSC/MNSC or API.

CNSC/MNSC Screenshot:

**Merchant setting**

[Table showing merchant settings with columns for setting type, various status fields, and descriptions]

**Merchant setting** section shows options for:
- Taobao Regional product information synchronization
- Settings showing:
  - Option for sync
  - Synchronized ✓
  - Not synchronize ✓
  - Categories (with EN/CN indicators)
  - Colors (Red, Blue, White)
  
*You should ensure sync_ authorized means no sync.*

Key field is `tv.global_product.pre_trans_sync.txt`, *true* means sync, *false* means no sync.

---

## 4.5 Uploading image

## 4.5.1 Uploading media files

Before calling any API to create or edit pre-saved in Shopee's media space, you need to call media_upload related API to upload files. For more details, please refer to the `show` "Uploading media files" section.

If you want to use product images or attribute images in CNSC/MNSC media space, you can call `tv.global_product.get_image_list` API to get the list of main images under the specified category.

---

## 5) Uploading category

We supports sellers to upload by categories so that they only select the category that returns `has_children=false from tv.global_product.get_category` API.

---

## 6) Translation requirements

In order to provide the best possible service to overseas Aliexpress merchants, we will have strict rules against the various products types (images, text, product description, etc.) to ensure the business is operated properly. Sellers are able to fill in simplified Chinese, English and numbers. Please do not mix simplified Chinese with traditional Chinese. If they meet local language words not being translated into local language. Note that the following requirements are mandatory.

About translation:
Cross-border product descriptions can only be filled in simplified Chinese, English, and numbers.

About restrictions:
Prohibited content: Regarding profanity and sensitive words can only be filled in simplified Chinese, English, and numbers.

---

## 4.6 Uploading global product information

When uploading global product information, please note:
We currently support unfinished sellers to select the shop to description, please refer to the `FAQ` for details on how to add or upload images in the description part.

About image uploading:
Cross-border Seller Center sellers can now make all of UPC for pricing. Accept cross-border sellers can use USD for pricing, and you can modify the pricing of `tv.global_product.product_add` API item without the price of the product.

About price:
For Single Type product, the price will not be converted with the exchange rate directly.

---

## 4.6.1 Uploading global product stock

API: `tv.global_product.product_add`

---

# 4.4.3 Uploading global product descriptions

**Great product descriptions attract buyers on simplified Chinese, English, and numbers.**

## Important Note
We currently support embedded videos to insert the image to description, please refer to the [FAQ for details](#) on how to insert the videos.

---

## 4.4.3 Uploading global product price

Oversea cross-border sellers can use RMB or USD for pricing. Oversea cross-border sellers can use USD for pricing, and domestic cross-border sellers can use RMB for pricing. For the [YJ conversion rate rules/common rate](#) API, there are other price units of the corresponding currency according to the currency selected by the sellers.

---

## 4.4.1 Uploading global product stock

The uploaded stock will go through a risk control review. If the review does not pass, a certain percentage of the uploaded stock will be frozen.

For example: If MT3501 publishes 3 MPU(s), then as MPU(s8, MPU(s8, MPU(s8), then when MT3501 stock is set to 100, stock MPU(s1+MPU(s2+MPU(s3)=120, then [MPU(s1, MPU(s2, and MPU(s3 will be displayed as stock 40 correspondingly.

---

## 4.8 Creating variants

When you create a global product successfully, to upload_product_sku_wms_type API will return the global_item_id, which is the product identifier for the global products. If you also need to define multiple variants, you can use the [$z.global_product_sku_create_or_update API](#) to upload variants. Variable variants will be generated during bulk variants upload.

> **Important:** Do not include variant specifications such as color and size in your product title and description.

---

## 4.8.1 Creating 1 tier variants

Assuming that the product has a color specification, the color includes S, L, and XL. We can fill in the 1-tier variants position accordingly.

### Example Table:

| tier_index | size | price | stock | sku |
|------------|------|-------|-------|-----|
| tier_index[0] | S/3 | 100 | 10 | sku1 |
| tier_index[1] | S | 200 | 20 | sku2 |
| tier_index[2] | M | 300 | 10 | - |

---

## API Request Example:

**Endpoint:** `sz.global_product_tier_tier_variation API request example:`

```json
{
  "global_item_id": 1000170002,
  
  "tier_variation": [
    {
      "name": "size",
      
      "option_list": [
        {
          "option": "S",
          
          "image": {
            "image_id": "030ce4fd59c82e09a6bc78c8c3e7f"
          }
        },
        {
          "option": "L",
          
          "image": {
            "image_id": "19bce4fd59c82e09a6bcf78c8c7f"
          }
        },
        {
          "option": "M"
        }
      ]
    }
  ],
  
  "model_list": [
    {
      "tier_index": [0],
      
      "original_price": 100,
      
      "global_model_sku": "sku1",
      
      "model_stock": 10
    },
    {
      "tier_index": [1],
      
      "original_price": 200,
      
      "global_model_sku": "sku2",
      
      "model_stock": 20
    },
    {
      "tier_index": [2],
      
      "original_price": 300,
      
      "global_model_sku": "",
      
      "model_stock": 30
    }
  ]
}
```

---

## 4.8.2 Creating 2-tier variants

Assuming that the product has color and size specifications, the color includes red and blue, and the size cookies xl, s, and l.

### Example Table:

| tier_index | color | size | price | stock | sku |
|------------|-------|------|-------|-------|-----|
| tier_index[0][0] | Red | S | 100 | 10 | sku1 |
| tier_index[0][1] | Red | L | 200 | 20 | sku2 |
| tier_index[1][0] | Blue | XL | 500 | 30 | sku3 |
| tier_index[1][1] | Blue | L | 400 | 40 | sku4 |

---

## API Request Example:

**Endpoint:** `sz.global_product_tier_tier_variation API request example:`

```json
{
  "global_item_id": 1000170002,
  
  "tier_variation": [
    {
      "name": "Color",
      
      "option_list": [
        {
          "option": "Red"
        },
        {
          "option": "Blue",
          
          "image": {
            "image_id": "030ce4fd59c82e09a6bc78c8c3e7f"
          }
        }
      ]
    },
    {
      "name": "Size",
      
      "option_list": [
        {
          "option": "S"
        },
        {
          "option": "XL"
        }
      ]
    }
  ]
}
```

---

# Developer Guide - Product Variant Configuration

## JSON Structure

```json
]
},
{
  "name": "size",
  
  "option_list": [
    {
      "option": "XL"
    },
    {
      "option": "L"
    }
  ]
}
],

"model": [
  {
    "tier_index": [0,0],
    
    "original_price": 100,
    
    "normal_stock": 10,
    
    "global_model_sku": "sku1"
  },
  {
    "tier_index": [0,1],
    
    "original_price": 200,
    
    "normal_stock": 20,
    
    "global_model_sku": "sku2"
  },
  {
    "tier_index": [1,0],
    
    "original_price": 300,
    
    "normal_stock": 30,
    
    "global_model_sku": "sku3"
  },
  {
    "tier_index": [1,1],
    
    "original_price": 400,
    
    "normal_stock": 40,
    
    "global_model_sku": "sku4"
  } ]
}
```

## Important Notes

**Please note that:**

- **The options you define will be displayed on the Shopee mall side in order. Shopee currently only supports the definition of up to 2-tier variation.**

- **You can define an image for each variant. If it is a 2-tier variation product, you can only define the first layer of options that is, in the example, you can define variant images based on color, but not based on size. Once you want to add variant images, all the options in the first layer need to define the image. The image needs to be uploaded by calling the V2 media upload image API**

- **The tier_index must start from 0 and not overflow, otherwise an error will be reported**

- **It is recommended that you create variants after an interval of 5 seconds after creating an item, because there may be a delay in creating item data.**

## Next Steps

Next, please review **how to publish global product**

---

**문서 ID**: developer-guide.213
**플랫폼**: shopee
**URL**: https://open.shopee.com/developer-guide/213
**처리 완료**: 2025-10-16T08:42:17
