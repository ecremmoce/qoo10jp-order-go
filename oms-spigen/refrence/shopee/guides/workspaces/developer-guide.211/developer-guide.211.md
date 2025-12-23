# API 레퍼런스 - 제품 생성

**카테고리**: API 레퍼런스
**난이도**: medium
**중요도**: 4/5
**최종 업데이트**: 2025-10-16T08:38:38

## 개요

본 개발자 가이드는 사용 가능한 API를 사용하여 제품을 생성하는 방법에 대한 정보를 제공합니다. 이미지, 비디오, 카테고리, 속성, 설명 및 가격 업로드와 각 단계별 특정 요구 사항 및 예제를 다룹니다.

## 주요 키워드

- 제품 생성
- API 레퍼런스
- 이미지 업로드
- 비디오 업로드
- 카테고리 업로드
- 속성 업로드
- 제품 설명
- 가격 업로드

## 본문

# API 참조 - 제품 생성

## 파일 업로드 API

---

## 1 이미지 업로드

API: **/upload_product_image** API

각 제품은 제품 이미지를 가져야 합니다. 또한 제품 설명에서 대부분의 이미지를 지원합니다. 제품 이미지 파일은 다음 요구 사항을 충족해야 합니다.

- 이미지 크기: <=5M
- 이미지 형식: JPG, JPEG, PNG

**참고 사항:**

- **/2.https://seller.aliexpress.com/** API는 이미지 파일 업로드만 지원하며 URL은 지원하지 않습니다. 이미지가 성공적으로 업로드되면 이미지 토큰(fileUrl)을 얻게 되며, 이를 고유한 image_id와 함께 저장해야 합니다. image_id는 비즈니스 의미 + 타임스탬프 + 난수의 조합으로 구성하는 것이 좋습니다(예: main1_20191015121212_123).
- 업로드에 성공하면 이미지를 비동기적으로 확인합니다(일반적으로 10초 이내, 대부분 몇 초).
- 제품 이미지를 업로드하는 경우 요청 매개변수 합계는 정상이어야 합니다. 제품 설명 이미지를 업로드하는 경우 이미지 유형은 description_image여야 합니다. 제품 이미지를 생성하면 정사각형 메인 사진으로 전송됩니다.

---

## 1.1 비디오 업로드

**제공되는 비디오는 선택 사항입니다. 다음 요구 사항을 충족하는 비디오 파일을 지원합니다.**

- 비디오 크기: <=500M
- 비디오 형식: mp4
- 비디오 해상도/권장 사항: 너비와 높이가 1280px * 720px 이하
- 업로드 시간 대략 비용: 1GB는 10초가 소요됩니다. 1분 비디오는 대략 9개로 분할하여 32, 64 또는 100개의 세그먼트로 분할해야 합니다. 각 조각은 4M보다 클 수 없기 때문입니다.

**단계:**

- 1단계: **/upload_aeop_vod/**video_order** API를 호출합니다. 비디오 업로드 작업을 생성하고 video_upload_id를 가져옵니다. 이 값을 저장해야 합니다. 이를 통해 비디오 상태를 확인해야 하기 때문입니다.
- 2단계: **/upload_video_stream/**video_part** API를 호출합니다. 비디오 파일에 조각이 있는 경우 part_seq(0부터 시작해야 함) 매개변수는 조각 번호를 나타냅니다. 매번 각 조각의 파일 크기를 4M보다 크게 업로드할 수 없습니다. part_seq 연속 호출이 반복되지 않는 것이 좋습니다(동시에, 동시 업로드는 안 됨).
- 3단계: 모든 조각이 성공적으로 업로드된 후에는 **/upload/**video**status** API를 호출하여 업로드한 모든 비디오 조각 파일의 완료를 확인해야 합니다. part_seq는 모든 조각의 시퀀스 번호인 매개변수입니다. 각 업로드 조각 part_seq가 0부터 시작하는 경우 0,1,2로 업로드해야 하며, 확인된 값은 0,1,2여야 합니다.
- 4단계: **/upload/**video/**order/**video** API를 호출합니다. 비디오 트랜스코딩 결과, 상태 값을 가져옵니다.

**비디오 트랜스코딩은 일반적으로 시간이 다소 걸립니다(2분 이내). 이 기간 동안 **/upload/**video/** order**video** API를 호출하여 비디오 상태를 쿼리합니다. 이 인터페이스에서 반환되는 현재 업로드 파일에 따르면 URL 및 비디오 커버 이미지 URL이 아직 처리 중인 경우 완료될 때까지 기다립니다. 기다린 후 video_upload_id를 사용하여 제품을 생성하거나 업데이트할 수 있습니다.

이미지 및 비디오를 업로드한 후 **/product/**add_order** API를 호출하여 제품을 생성할 수 있습니다. 다음은 **/product/**add_order** API를 호출하여 제품을 생성하고 업데이트하는 방법에 대한 자세한 설명을 제공합니다.

---

## 2 카테고리 및 속성 업로드

### 2.1 카테고리 업로드

제품을 업로드할 때 리프 카테고리 ID만 선택할 수 있습니다. 제품에 대한 이 카테고리 상자에서 그렇지 않으면 제품을 성공적으로 만들 수 없으며 오류가 반환됩니다.

### 2.2 속성 업로드

API: **/product/**get/**attribute/**result_order** API

제품에 대한 이 카테고리에서 제품은 여러 속성을 선택할 수 있습니다. 다른 카테고리 속성 및 속성 메타데이터는 필수 사항으로 간주됩니다. **/product/**get/**attribute/**result_order** API에서 반환된 속성 목록을 참조하여 **/product/**order**product** API를 호출할 때 속성 및 속성 유형에 대한 자세한 내용은 "제품 생성 정보" 문서를 참조하십시오.

"Input_type": ("INPUT_DROP_DOWN")으로 속성을 업로드할 때 하나의 "value_id"만 업로드할 수 있으며 가져온 "value_id"는 **/product/**get/**attribute/**result_order** API에서 반환된 속성 목록의 값 중 하나여야 합니다.

**예제 1**

```json
"attributes":{
    "ATTRIBUTE[0]":{
        "ORDER":
            "ATTRIBUTE[0]:INPUT":{
                "value_id": 826
            }
    }
}
```

"Input_type": {"MULTI_DROP_DOWN"}으로 속성을 업로드할 때 여러 "value_id" 값을 업로드할 수 있으며 "value_id" 값은 **/product/**get/**attribute/**result_order** API에서 반환된 속성 목록의 일부여야 합니다.

**예제 2**

```json
"attributes":{
    "ATTRIBUTE[1]":{
        "ATTRIBUTE[1]:ORDERS":
            "ATTRIBUTE[1]:ORDER[0]":{
                "value_id": 826
            },
            "ATTRIBUTE[1]:ORDER[1]":{
                "value_id": 827
            }
    }
}
```

"Input_type": {"FREE_TEXT_FILL2"}으로 속성을 업로드할 때 하나의 "value_A"만 업로드하고 A 사용자 지정 값을 지정할 수 있습니다. 따라서 custom_value_A를 생략해야 합니다. 제공하는 경우 사용자 지정 값은 "customiz_value_A"입니다.

**예제 3**

```json
"attributes":{
    "ATTRIBUTE[2]":{
        "ATTRIBUTE[2]:ORDERS":
            "ATTRIBUTE[2]:ORDER[0]":{
                "value_id": 0,
                "value_name": "92",
                "customiz_value": "customiz1"
            }
    }
}
```

속성에 "input_type": {"FREE_TEXT_FILL2"}가 있고 단위가 필요한 경우 "original_value_name" 필드에 독립 실행형을 업로드해야 합니다(날짜 1, 날짜 1A 및 사용 가능한 경우 unit1="4" 매개변수).

**예제 4**

```json
"attributes":{
    "ATTRIBUTE[3]":{
        "ATTRIBUTE[3]:ORDERS":
            "ATTRIBUTE[3]:ORDER[0]":{
                "value_id": 0,
                "value_name": "92",
                "customiz_value": "customiz2"
            },
            "value_id": 826
    }
}
```

"input_type": {"INPUT_COMBO_BOX"}가 있는 속성의 경우 하나의 "value_id"만 업로드할 수 있습니다. 이 "value_id"는 **/product/**get/**attribute/**result_order** API에서 반환된 속성 목록의 값 중 하나여야 합니다(예제 1을 참조하거나 예제 3에서 하나의 사용자 지정 값을 업로드하십시오. 플래그가 지정되지 않은 경우 예제 4를 참조하십시오. 이 경우 value_id !=0 또는 정확히 사용자 지정 값입니다.

"input_type": {"MULTI_COMBO_BOX"}가 있는 속성의 경우 **/product/**get/**attribute/**result_order** API에서 반환된 속성 목록에서 여러 "value_id" 값을 업로드하거나 정확히 사용자 지정 값을 업로드할 수 있습니다.

**예제 5**

```json
"attributes":{
    "ATTRIBUTE[4]":{
        "ATTRIBUTE[4]:ORDERS":
            "ATTRIBUTE[4]:ORDER[0]":{
                "value_id": 0,
                "value_name": "craft unit89"
            },
            "value_id": 826
    }
}
```

---

# 개발자 가이드 - API 문서

## 예제 1

```ruby
api/data.json: {
  {
    "attribute_id":"AAAA",
    "attribute_value":{}
  }
  {
    "value_id":120
    "original_value_name":"colorful_name"
  }
  ,
  ,
}
```

---

필수 사항을 추가하려면 예제 4를 참조하십시오. 숨김 유형이 필요한 경우 예제 5를 참조하십시오.

### 요약

속성 유형의 경우 "value_id"를 업로드해야 합니다. (사용자 정의 값을 업로드하는 경우 "value_id"는 0이어야 하며 "original_value_name"은 필수입니다).

제품 API의 기본값인 "original_value_name" 필드는 문자열 형식으로 업로드해야 합니다. API가 특수 문자를 사용하도록 설정되지 않은 경우(필요한 경우 "value_id"도 업로드) 예제 가격 필드에 두 가지 옵션이 있는 경우 첫 번째 옵션은 "attribute_value_id"용이고 "original_value_id"는 **to_enabled_site_channel** API의 확장을 위한 것입니다.

---

## 3. 설명 업로드

제품 설명 패턴에 줄 바꿈이 있는 경우 extended_description을 사용하려면 업로드에 대한 자세한 내용은 **FAQ**를 참조하십시오.

---

## 4. 가격 업로드

사우디 상인을 제외하고 판매자가 소수점 두 자리 가격 original_price(s)를 업로드할 것으로 예상되는 다른 마켓플레이스의 경우 판매자는 자체 규칙을 사용할 수 있습니다.

---

## 5. 재고 업로드

판매자는 현재 사용 가능한 사용자에게만 사용되는 재고 창고가 없으므로 재고를 업데이트할 필요가 없습니다.

---

## 6. 배송 채널 및 배송비 업로드

배송 채널을 업로드할 때 판매자는 어떤 배송 채널이 (**to_product_site_channel** API에서 수수료를 받을 수 있는 여러 유형)인지 확인해야 합니다. 수수료 유형은 다음과 같습니다.
- RSE_SELECTION: 배송은 구매자가 선택할 수 있는 무게 단위 기준으로 업로드됩니다(판매자는 제품에 대해 이를 사용해야 함).
- FLAT_FEE: 판매자는 고정 기본 수수료를 업로드합니다.
- FREE: 판매자는 배송비를 지불합니다(판매자 프로모션으로만 사용). 이 판매자에서 수수료와 (판매자) 수수료는 배송비에 대한 수수료를 업데이트하지 않으며 모든 채널 유형은 배송으로 충분히 지원합니다.

판매자는 제품에 대해 여러 채널을 선택할 수 있습니다(선택/수수료는 제품을 이 채널에 개방하고 구매자는 구매자/채널이 결제 시 사용할 채널을 선택할 수 있도록 합니다). 제품이 열려 있는 경우에만 "지원됨"으로 표시되며 구매자는 배송비에 대해 이 도구를 테스트하지 않습니다.

RSE_SELECTION 채널의 경우 size_ad를 계산해야 하며 size_cd는 (필수 키)일 수 있습니다.
- **to_enabled_site_channel API**

### 예제 1:

```json
"height":7
"shipment":"us"
"width":8
"original_value_name"
"size":"Post Base"
"type":"RSE_PICKUP"
```

### 예제 2:

RSE_DEPTH 채널의 경우 무게와 치수를 업로드해야 합니다.

```json
"height":7
"package_weight":71
"package_length":20
"package_width":19
"type":"us"
"width":7
```

### 예제 3:

CUSTOM_PRICE 채널의 경우 이 shipping_fee를 업로드해야 합니다.

```json
"height":7
"shipping_fee":23.72
"width":7
"size":"Post Base"
"type":"CUSTOM_PRICE"
```

---

## 7. 변형 생성

생성할 때(크기/사용 가능, **to_product_site_color**)는 항목의 고유 식별자인 item_id를 반환합니다.

사양을 정의하고 색상 및 크기와 같은 여러 옵션 변형을 생성해야 하는 경우 다음 단계에서 **to_product_site_variant_specified** API를 사용할 수 있습니다. 그러면 변형의 고유한 SKU가 될 선택 조합을 만들 수 있습니다.

**시나리오 1:** 제품에 두 가지 사양이 있고 크기에 A, S 및 B가 포함되어 있습니다. 변형 그룹별로 연결할 수 있습니다.

| bar_white | size | price | stock | SKU |
|-----------|------|-------|-------|-----|
| bar_white(1) | A/S | 100 | H | sku1 |
| bar_white(2) | B | 200 | H | sku2 |
| bar_white(3) | H | 300 | 12 | sku3 |

**to_product_site_tier_variation** API 요청 예제:

```json
{
  "item_id": 835558262,
  "tier_variation": [
    {
      "name": "size",
      "option_list": [
        {
          "option": "65",
          "image": {"value":"f3ec4c833f58ce950e9fea94b8e363e7"}
        }
      ]
    },
    {
      "name": "size",
      "option_list": [
        {
          "option": "8",
          "image": {"value":"785ec833f58cce950e8fea93bbe364e7"}
        }
      ]
    },
    {
      "option": "8",
      "image": {"value":"785ec833f58cce950e8fea93bbe364e7"}
    }
  ],
  "tier_index": [0],
  "original_price": 200,
  "price": [1]
  {
    "tier_index": [0],
    "model_sku": "blue",
    "original_price": 0
  }
}
```

---

# API 문서 추출

## 시나리오 2: Two-Product API

제품에 색상 및 크기 사양이 있습니다. 색상에는 빨간색과 파란색이 포함되고 크기에는 XL과 L이 포함됩니다. 이는 2가지 변형 제품입니다.

### 제품 테이블

| var_index | color | size | price | stock | SKU |
|-----------|-------|------|-------|-------|------|
| var_skuindex[0] | Red | XL | 500 | 50 | sku1 |
| var_skuindex[1] | Red | L | 300 | 20 | sku2 |
| var_skuindex[2] | Blue | XL | 500 | 30 | sku3 |
| var_skuindex[3] | Blue | L | 450 | 40 | sku4 |

### v2.product.sku_list_update API 요청 예제

```json
{
  "list_id": "10001+02",
  "list_add_list": [
    {
      "user": "u1id",
      "get_im_list": [
        {
          "image": ["image_id123c45def6a8bc7f8a9c0d37"],
          "sel_var": "Red"
        },
        {
          "image": ["image_id123c45def6a8bc7f8a9c0d37"],
          "sel_var": "Blue"
        }
      ]
    },
    {
      "user": "s1_id",
      "get_im_list": [
        {
          "sel_var": "XL"
        },
        {
          "sel_var": "L"
        }
      ]
    }
  ],
  "scull": [],
  "list_index": [0, 0],
  "original_price": 115,
  "model_stock": 10,
  "seller_admin_sku": "sku1"
},
{
  "list_index": [0, 1],
  "original_price": 200,
  "model_stock": 20,
  "seller_admin_sku": "sku2"
},
{
  "list_index": [1, 0],
  "original_price": 300,
  "model_stock": 30,
  "seller_admin_sku": "sku3"
},
{
  "list_index": [1, 1],
  "original_price": 400,
  "model_stock": 40,
  "seller_admin_sku": "sku4"
}
```

## 중요 참고 사항

**다음 사항에 유의하십시오.**

- SKU 목록 정보는 Shopee 몰 측에 순서대로 표시됩니다. Shopper는 현재 2가지 변형 제품의 정의만 지원합니다.
- 각 변형에 대한 이미지를 정의할 수 있습니다. 2단계 변형 제품인 경우 첫 번째 항목 또는 변형만 정의할 수 있습니다. 먼저 product.UploadImg API를 사용하여 제품 이미지를 업로드한 다음 product.add_Add 또는 product.UpdateItem을 호출하여 변형의 이미지 목록을 가져오는 것이 좋습니다.
- 2가지 변형 제품이 있는 경우 firstvar_name 아래에 최소 1개의 항목 ID가 있어야 합니다.
- tier_index는 0부터 시작해야 하며 제공되는 인터페이스만 제공될 수 있습니다.
- 제품 이미지를 업로드한 후에는 레거시 항목일 수 있으므로 item_sku에서 사용할 수 있습니다.
- 성공적으로 생성한 후 v2.product.get_model_list API를 호출하여 각 two_attrs에 해당하는 모델 ID를 가져올 수 있습니다.

## 사용 사례

1. 카탈로그에 새 제품 생성
2. 제품 정보 업데이트
3. 제품 이미지 및 비디오 업로드
4. 제품 속성 및 카테고리 설정

## 관련 API

- /upload_product_image
- /upload_aeop_vod/video_order
- /upload_video_stream/video_part
- /upload/video/status
- /upload/video/order/video
- /product/add_order
- /product/get/attribute/result_order

---

## 원문 (English)

### Summary

This developer guide provides information on how to create products using the available APIs. It covers uploading images, videos, categories, attributes, descriptions, and prices, along with specific requirements and examples for each step.

### Content

# API Reference - Creating product

## API to upload file

---

## 1 Uploading image

API: **/upload_product_image** API

Each product must have product images. In addition, we also support most images in product description. And the product image files need to meet the following requirements:

- Image size: <=5M
- Image format: JPG, JPEG, PNG

**Please note:**

- **/2.https://seller.aliexpress.com/** API we only support image file upload, not support URLs. When the image is uploaded successfully, we will get an image token (fileUrl), you should save it together with a unique image_id. We recommend that the image_id be a combination of business semantics + timestamp + random number, such as main1_20191015121212_123.
- After uploading successfully, we will check the image asynchronously (generally within 10 seconds, mostly just a few seconds).
- If you upload a product image, the request parameter sums should be normal; if you upload a product description image, the image type should be description_image. When you create the product image we will transfer to a square main picture.

---

## 1.1 Uploading video

**Provided video is optional. We support video files with the following requirements:**

- Video size: <=500M
- Video format: mp4
- Video resolution/suggested: your width and height no more than 1280px * 720px
- Upload time approximately costs about: 1GB takes 10 seconds. 1 Minute video approximately is about 9 to split it into 32, 64 of 100 segments, because each fragment cannot be larger than 4M.

**Steps:**

- Step 1: Call the **/upload_aeop_vod/**video_order** API. Create a video upload task and get the video_upload_id. You should save this value, because we need to check the video status through it.
- Step 2: Call the **/upload_video_stream/**video_part** API. When the video file has fragments, the part_seq (should be from 0) parameter indicates the fragment number. Each time you can not upload file sizes larger than 4M of each fragment. It would be better if the part_seq continuous call could not be repeated (the same time, and must not be concurrent upload).
- Step 3: After all fragments have been uploaded successfully, you must call the **/upload/**video**status** API to confirm completion of all video fragments files you uploaded. The part_seq is a parameter that the sequence number of all fragments. If each upload fragment part_seq from 0, which should be uploaded as 0,1,2, and then confirmed the value should be 0,1,2.
- Step 4: Call the **/upload/**video/**order/**video** API. Get the video transcoding result, status value represents the result.

**Transcoding video will generally take some time (within two minutes), during the period will be called **/upload/**video/** order**video** API to query video status. According to the current upload file of this interface returning, when the URLs and video cover image URL are still in processing, they will wait until it is done. Wait and at that time, the video_upload_id can be used to create or update the products.

After uploading images and videos, you can call **/product/**add_order** API to create products. the following will explain the detail to create and update products by calling **/product/**add_order** API.

---

## 2 Uploading category and attributes

### 2.1 Uploading category

Uploading products you can only select the leaf category ID. In this category box for the product, otherwise, the product can't create successfully, returning an error.

### 2.2 Uploading attributes

API: **/product/**get/**attribute/**result_order** API

In this category for the product, the product can select multiple attributes. Different categories attributes and attribute metadatas are considered required. You can refer to the article "Product creation information" to learn more about attributes and attribute types when calling **/product/**order**product** API by reference to the attributes list returned by the **/product/**get/**attribute/**result_order** API.

When uploading the attribute with "Input_type": ("INPUT_DROP_DOWN"), you can only upload one "value_id" and fetch "value_id" must be one of the values from the attribute list returned by the **/product/**get/**attribute/**result_order** API.

**Example 1**

```json
"attributes":{
    "ATTRIBUTE[0]":{
        "ORDER":
            "ATTRIBUTE[0]:INPUT":{
                "value_id": 826
            }
    }
}
```

When uploading the attribute with "Input_type": {"MULTI_DROP_DOWN"}, you can upload multiple "value_id" values, and "value_id" must "value_id" values must be some of the attribute list returned by the **/product/**get/**attribute/**result_order** API.

**Example 2**

```json
"attributes":{
    "ATTRIBUTE[1]":{
        "ATTRIBUTE[1]:ORDERS":
            "ATTRIBUTE[1]:ORDER[0]":{
                "value_id": 826
            },
            "ATTRIBUTE[1]:ORDER[1]":{
                "value_id": 827
            }
    }
}
```

When uploading the attribute with "Input_type": {"FREE_TEXT_FILL2"}, you can only upload one "value_A" and specify A customization value. Therefore, you must omit the custom_value_A. If you provide it the custom value is the "customiz_value_A".

**Example 3**

```json
"attributes":{
    "ATTRIBUTE[2]":{
        "ATTRIBUTE[2]:ORDERS":
            "ATTRIBUTE[2]:ORDER[0]":{
                "value_id": 0,
                "value_name": "92",
                "customiz_value": "customiz1"
            }
    }
}
```

If the attribute has "input_type": {"FREE_TEXT_FILL2"} and requires a unit, you must upload a freestanding in the "original_value_name" field (Date 1, Date 1A, and if available, unit1="4" parameter).

**Example 4**

```json
"attributes":{
    "ATTRIBUTE[3]":{
        "ATTRIBUTE[3]:ORDERS":
            "ATTRIBUTE[3]:ORDER[0]":{
                "value_id": 0,
                "value_name": "92",
                "customiz_value": "customiz2"
            },
            "value_id": 826
    }
}
```

For the attribute with "input_type": {"INPUT_COMBO_BOX"}, you can only upload one "value_id" This "value_id" must be one of the values from the attribute list returned by the **/product/**get/**attribute/**result_order** API (please refer to Example 1 or upload one customization value in Example 3. If it is not a flagged, please refer to Example 4. For this case the value_id !=0 or exactly custom values.

For the attribute with "input_type": {"MULTI_COMBO_BOX"}, you can upload multiple "value_id" values from the attribute list returned by the **/product/**get/**attribute/**result_order** API, or exactly custom values.

**Example 5**

```json
"attributes":{
    "ATTRIBUTE[4]":{
        "ATTRIBUTE[4]:ORDERS":
            "ATTRIBUTE[4]:ORDER[0]":{
                "value_id": 0,
                "value_name": "craft unit89"
            },
            "value_id": 826
    }
}
```

---

# Developer Guide - API Documentation

## Example 1

```ruby
api/data.json: {
  {
    "attribute_id":"AAAA",
    "attribute_value":{}
  }
  {
    "value_id":120
    "original_value_name":"colorful_name"
  }
  ,
  ,
}
```

---

If add a required, please refer to Example 4. If a hide type is required, please refer to Example 5.

### Summary

For the attribute types, you must upload a "value_id". (For uploading a user-defined value, "value_id" must be 0 and the "original_value_name" is required).

The product API's default value, the "original_value_name" field must be uploaded in string format. If the API is not set to use special characters, you (then also upload the "value_id" if this required). If the example price field has two options, the first option must be for the "attribute_value_id", and the "original_value_id" is for the expansion of the **to_enabled_site_channel** API.

---

## 3. Uploading description

When there are line breaks in the product description patterns to use extended_description, please refer to the **FAQ** for details on uploading.

---

## 4. Uploading price

Except for Saudi merchants which expect sellers to upload two decimal price original_price(s), for other marketplaces, sellers can use their own rules.

---

## 5. Uploading stock

The seller does not have a stock warehouse currently only used for available users then there is no need to update the stock.

---

## 6. Uploading shipment channels and shipping fee

When uploading shipment channels, the seller must check which shipment channels will (several types, you can get the fee from **to_product_site_channel** API). fee types including:
- RSE_SELECTION: The shipping is upload base on the weight-unit of the buyer can choose (sellers must use this for the product)
- FLAT_FEE: Sellers upload the fixed base fee
- FREE: Sellers pay the shipping fee (only as a seller promotion). In this seller, the fee and the (sellers) fee will not update the fee for the shipping fee, and all channel types support sufficient as shipping.

The seller can choose multiple channels for the product (selection/fee makes the product it open to this channel and the buyer can choose which buyers/channel choose to use in checkout) then the buyers. If the product is open, only then will it be shown as "supported, the buyer does not test this tool for the shipping fee.

For the RSE_SELECTION channel, the size_ad must be calculated, and size_cd can be (required keys):
- **to_enabled_site_channel API**

### Example 1:

```json
"height":7
"shipment":"us"
"width":8
"original_value_name"
"size":"Post Base"
"type":"RSE_PICKUP"
```

### Example 2:

For the RSE_DEPTH channel, the weight and dimension must be uploaded.

```json
"height":7
"package_weight":71
"package_length":20
"package_width":19
"type":"us"
"width":7
```

### Example 3:

For the CUSTOM_PRICE channel, this shipping_fee must be uploaded.

```json
"height":7
"shipping_fee":23.72
"width":7
"size":"Post Base"
"type":"CUSTOM_PRICE"
```

---

## 7. Creating Variants

When you create (then size/available, **to_product_site_color**) will return the item_id, which is the unique identifier of the item.

If you also need to define specifications and create multiple option variants, such as the color and size, you can use **to_product_site_variant_specified** API for the next step, then we can create choice combinations, which will be the unique skus of variants.

**Scenario 1:** The product has two specifications and the size contains A,S, and B. We can link it per variation groups.

| bar_white | size | price | stock | SKU |
|-----------|------|-------|-------|-----|
| bar_white(1) | A/S | 100 | H | sku1 |
| bar_white(2) | B | 200 | H | sku2 |
| bar_white(3) | H | 300 | 12 | sku3 |

**to_product_site_tier_variation** API request example:

```json
{
  "item_id": 835558262,
  "tier_variation": [
    {
      "name": "size",
      "option_list": [
        {
          "option": "65",
          "image": {"value":"f3ec4c833f58ce950e9fea94b8e363e7"}
        }
      ]
    },
    {
      "name": "size",
      "option_list": [
        {
          "option": "8",
          "image": {"value":"785ec833f58cce950e8fea93bbe364e7"}
        }
      ]
    },
    {
      "option": "8",
      "image": {"value":"785ec833f58cce950e8fea93bbe364e7"}
    }
  ],
  "tier_index": [0],
  "original_price": 200,
  "price": [1]
  {
    "tier_index": [0],
    "model_sku": "blue",
    "original_price": 0
  }
}
```

---

# API Documentation Extract

## Scenario 2: Two-Product API

The product has color and size specifications; the color includes red and blue, and the size includes XL and L. We can this is two-variation product.

### Product Table

| var_index | color | size | price | stock | SKU |
|-----------|-------|------|-------|-------|------|
| var_skuindex[0] | Red | XL | 500 | 50 | sku1 |
| var_skuindex[1] | Red | L | 300 | 20 | sku2 |
| var_skuindex[2] | Blue | XL | 500 | 30 | sku3 |
| var_skuindex[3] | Blue | L | 450 | 40 | sku4 |

### v2.product.sku_list_update API Request Example

```json
{
  "list_id": "10001+02",
  "list_add_list": [
    {
      "user": "u1id",
      "get_im_list": [
        {
          "image": ["image_id123c45def6a8bc7f8a9c0d37"],
          "sel_var": "Red"
        },
        {
          "image": ["image_id123c45def6a8bc7f8a9c0d37"],
          "sel_var": "Blue"
        }
      ]
    },
    {
      "user": "s1_id",
      "get_im_list": [
        {
          "sel_var": "XL"
        },
        {
          "sel_var": "L"
        }
      ]
    }
  ],
  "scull": [],
  "list_index": [0, 0],
  "original_price": 115,
  "model_stock": 10,
  "seller_admin_sku": "sku1"
},
{
  "list_index": [0, 1],
  "original_price": 200,
  "model_stock": 20,
  "seller_admin_sku": "sku2"
},
{
  "list_index": [1, 0],
  "original_price": 300,
  "model_stock": 30,
  "seller_admin_sku": "sku3"
},
{
  "list_index": [1, 1],
  "original_price": 400,
  "model_stock": 40,
  "seller_admin_sku": "sku4"
}
```

## Important Notes

**Please note that:**

- The information of the SKU list will be displayed on the Shopee mall side in order. Shopper currently only supports the definition of two-variation products.
- You can define an image for each variant. If it is a 2-tier variation product, you can only define the first item or variants. We suggest use product.UploadImg API to upload product image first; and then call product.add_Add or product.UpdateItem to get the image list of variants.
- If there are two-variation products, there must be at least 1 item id under the firstvar_name.
- The tier_index must start from 0 and not interface offered can and only will be supplied.
- After uploading product images, you can use them in item_sku because this may be a legacy item.
- After successful creation, you can call v2.product.get_model_list API to get the model id corresponding to each two_attrs.

---

**문서 ID**: developer-guide.211
**플랫폼**: shopee
**URL**: https://open.shopee.com/developer-guide/211
**처리 완료**: 2025-10-16T08:38:38
