# API 호스트 제품 - Instant Mart 통합-Dev

**카테고리**: 통합
**난이도**: medium
**중요도**: 4/5
**최종 업데이트**: 2025-10-16T09:32:11

## 개요

이 개발자 가이드는 Instant Mart의 Open API 통합에 대한 정보를 제공합니다. 여기에는 샌드박스 계정 설정, 인증, 테스트, 그리고 Instant Mart 생태계 내의 제품, 주문, 반품, 재무, 상점 설정을 관리하기 위한 주요 API 기능(메인 상점 및 아울렛 상점 포함)이 포함됩니다.

## 주요 키워드

- API 통합
- Instant Mart
- Open API
- Sandbox testing
- Shop authorization
- 제품 관리
- 주문 관리
- 반품 & 환불
- 재무
- 아울렛 상점

## 본문

# API 호스트 제품 - Instant Mart 연동 - 개발

## 배경:
Instant Mart는 온라인 관리 및 지점 키 시스템을 통합을 통해 운영하는 소매업체를 지원하도록 설계된 오픈 소스 쇼핑몰입니다.

- 각 Mart Shop(공식 쇼핑몰)은 글로벌 재고 관리, 상품 생성 및 재무 보고서 동기화를 담당하는 본사 역할을 합니다.

- 여러 아울렛(지점 쇼핑몰)(및 사용자)은 각 쇼핑몰 위치를 나타냅니다. 아울렛은 지역 재고, 주문 관리, 재고 유지 및 구매 주문과 같은 일상적인 운영을 담당합니다.

Open API 연동 후 Mart 소매업체는 SKU를 판매자 수준(모든 아울렛에서 공유) 또는 아울렛 수준(지점별 독립)에서 관리할지 선택할 수 있습니다.

Open API와 통합하는 소매업체는 선택한 구조 설계 계획이 RDC 포털에서 판매자와 아울렛 쇼핑몰 모두에 대한 흐름을 정의하는 방식에 영향을 미친다는 점을 이해해야 합니다. (각 통합 계정은 본사 또는 지점 수준에서 RDC 시스템의 단일 엔터티에 연결됩니다. 이는 개발자가 아울렛 쇼핑몰의 전체 기능 목록 반환을 고려하여 API 기능을 구축하는 방법을 결정합니다.)

---

## 1. 샌드박스 테스트 계정 설정

### 1.1 샌드박스 테스트 계정 생성

Opened Open Platform Console의 [SandGuide](#) 페이지에 접속하여 시작할 수 있습니다. Mart/아울렛 쇼핑몰 모듈을 선택하여 샌드박스 쇼핑몰을 생성합니다.

*참고: Instant Mart 샌드박스 테스트에는 2단계 스케줄러가 필요합니다. Mart 4 아울렛 쇼핑몰을 활성화하기 전에 스케줄러 1/2를 활성화해야 합니다.*

---

## 2. 권한 부여

### 2.1 쇼핑몰 권한 부여
탐색 페이지에서 Instant Mart 파트너 URL을 찾습니다.

*[“앱 목록” 및 쇼핑몰 선택 인터페이스를 보여주는 권한 부여 흐름 스크린샷]*

Authorize를 클릭하고 Mart/아울렛 쇼핑몰 계정으로 로그인하여 권한 부여를 완료합니다.

*[권한 부여 버튼이 있는 "Shoppe OpenPlatform API 권한 부여 로그인" 인터페이스를 보여주는 스크린샷]*

*[권한 부여 프로세스 단계를 보여주는 추가 스크린샷]*

자세한 내용은 [권한 부여 및 권한 부여](#) 문서를 참조하십시오.

---

## 3. 테스트 프로세스

### 3.1 Open API 테스트
API 테스트 도구를 사용하거나 Postman을 사용하여 Open API를 테스트합니다.

*[요청/응답 패널이 있는 API 테스트 인터페이스를 보여주는 스크린샷]*

---

### 3.2 주요 Open API 기능

M-Mart 프로젝트에 필요한 주요 API 기능:

| 기능 | 하위 기능 |
|----------|--------------|
| Sandbox | |
| Product | 신규 SKU 준비 |
| | 신규 SKU 생성 |
| | 기존 SKU 업데이트 |
| | SKU 옵션 생성 |
| | 재고 및 가격 관리 |
| | 일반 SKU 관리 |
| Order | 주문 목록 및 상세 정보 가져오기 |
| | 주문 취소 |
| | 배송 요청 |
| | 추적 번호 생성 |
| Return & Refund | 반품 목록 및 상세 정보 가져오기 |
| | 반품/환불 승인 |
| | 반품/환불 이의 제기 |
| Financials | |

---

**참고:** 이 추출은 개발자 가이드 스크린샷의 원래 영어 텍스트 구조를 유지하여 제목, 목록, 표 및 Markdown 형식의 조직 계층 구조를 보존합니다.

---

# API 문서 추출

## 반품 및 환불
- 추적 번호 생성
- 항공 운송장 생성
- 반품 목록 및 상세 정보 가져오기
- 반품/환불 승인
- 반품/환불 이의 제기

## 재무
- 에스크로 상세 정보 가져오기
- 지갑 거래 내역 가져오기
- 소득 명세서 생성
- 소득 보고서 생성

## 쇼핑몰 설정
- 쇼핑몰 프로필 업데이트
- 쇼핑몰 운영 시간

---

## 3.2.1 MarkOutofStock 관계

관계 확인: `or_shop_gir__shop_info`

input_shop_gir = related_gir을 사용하여 시스템이 Mart Shop이 related_shop_gir(메인 쇼핑몰 ID)과 관련되어 있는지 확인합니다.

**참고:** 쇼핑몰 gir은 메인 쇼핑몰과 관련되어 있으므로 호출 단계에서 (메인 쇼핑몰 ID)를 알아야 합니다.

## 3.2.1.1 or shop gir_shop info 가져오기

`or shop gir_shop_info`를 main_shop_id와 함께 호출하면 해당 메인 쇼핑몰에 속한 모든 아울렛 쇼핑몰 ID 목록이 반환됩니다.

### 응답 예시:

```json
{
  "sells_shop": "1700820216",
  "toplist_shop": "1704576796",
  "or_gir": "fass",
  "or_main_shop": "fass",
  "or_main_shop": "fass",
  "or_next_shop": "fass",
  "or_outlet_shop": "fass",
  "or_shop_grp": "fass",
  "or_shop_info": "fass",
  "items_main_shop_id": 0,
  "overstock_id": null,
  "outlet_shop_add_sett_1": "",
  "outlet_shop_id": 226402201,
  "outlet_shop_id": 226402202,
  "outlet_shop_id": 226402203,
  "outlet_shop_id": 226402204,
  "trigger": "123",
  "trigger": "no",
  "incomin_gir": "vg_K2Ni9xY...",
  "incomin_outlet_shop_id": "17JdA2536GG37",
  "keep_some": "vg_6DaAhG2Y43G9Y8/GqF",
  "outint": "IGz9M0n.."
}
```

---

## 3.2.1.2 or shop gir_shop_info를 아울렛 쇼핑몰과 함께 사용

`or shop gir_shop_info`를 outlet_shop_id와 함께 호출하면 해당 메인 쇼핑몰 ID 및 관련 데이터가 반환됩니다.

### 응답 샘플:

```json
{
  "sells_shop": "1704156903",
  "toplist_shop": "1704723399",
  "or_gir": "fass",
  "or_main_shop": "fass",
  "or_main_shop": "fass",
  "or_next_shop": "fass",
  "or_shop": "fass",
  "or_gir": "fass",
  "or_signature_shop": "",
  "items_main_shop_id": 226402201,
  "items_main_shop_id": 0,
  "overstock_id": 226621997,
  "overstock_id": null,
  "message": "",
  "trigger": "123",
  "trigger": "no",
  "incomin_gir": "vg_K7D3LG4YZ...",
  "main_shop_id": "9Y1Io4G8/G9Y6/10aBt",
  "shop_some": "vg_J1G02k446G17kK6BB57",
  "outint": "IGz9M0n.."
}
```

---

## 3.2.2 상품 관리

[이것은 다이어그램입니다: 화살표로 연결된 상자가 있는 제품 관리 계층 구조를 보여주는 순서도로, "product_db_send", "db product info_in_detail info" 및 "db sent while descending_in_a"와 같은 요소가 포함되어 있습니다.]

### 메인 쇼핑몰과 아울렛 쇼핑몰 간의 필드 관계:

| 모듈 | 필드 | MART SKU 관리 | 아울렛 SKU 관리 |
|--------|-------|--------------------|-----------------------|
| 기본 정보 | title | 예 | 아니요 |
| | video | 예 | 아니요 |
| | image | 예 | 아니요 |
| | status | 예 | 예 |
| | category | 예 | 아니요 |
| 사양 | brand | 예 | 아니요 |
| | attribute | 예 | 아니요 |
| | size chart | 예 | 아니요 |
| 사양 | variation name | 예 | 아니요 |
| | variation option | 예 | 아니요 |
| | variation image | 예 | 아니요 |
| | price | 아니요 | 예 |
| | purchase limit | 아니요 | 예 |
| | subchannel | 아니요 | 예 |
| | discount | 아니요 | 예 |
| | model status | 예 | 예 |
| | module status SKU | 예 | 아니요 |
| | weight | 예 | 아니요 |
| 배송 | dimension | 예 | 아니요 |
| | logistics channel | 아니요 | 예 |
| | preparation channel | 아니요 | 예 |
| 기타 | DTS | 아니요 | 예 |
| | CBS | 아니요 | 예 |
| | condition | 예 | 아니요 |

---

**참고:** 이 추출은 개발자 가이드 스크린샷의 원래 구조, 제목, 코드 블록 및 표 형식을 유지합니다.

---

# 개발자 가이드 - 기능 비교 및 설정

## 기능 비교 표

| 기능 | 대안 | 예 | 아니요 |
|---------|-------------|-----|-----|
| Shipping | - | - | 아니요 |
| - | bundles (일괄 처리) | - | 예 |
| - | installation channel | 아니요 | 예 |
| 기타 | DTS | 아니요 | 예 |
| - | seller SKU | 예 | 아니요 |
| - | condition | - | - |

## 주요 작업 포함

1. 라벨 SKU 추가 --> `v1_product_add_item`
2. 라벨 SKU 업데이트 --> `v1_product_update_item` (가격, 재고, product_price, 아울렛 수준 가격 포함)
3. 코드 SKU 게시 --> `v1_product_publish_item_to_outlet_shop`
4. 판매 SKU 동기화 (+) --> `v1_product_update_item`

---

## 3.2.1 라벨 SKU 추가

main shop_id를 사용하여 `v1_product_add_item`을 호출하여 라벨 SKU를 추가합니다.

자세한 내용은 **[자세한 기술 자료 문서](#)**를 참조하십시오.

---

## 호출 샘플:

```json
{
  "image_price_list": 123.5,
  "description": "cached and item from /getsign API",
  "weight": 1.5,
  "image_price_list": 123.5,
  "description": "cached and item from /getsign API",
  "weight": 1.5,
  "Itemplace_height": 2,
  "Itemplace_length": 3,
  "Itemplace_width": 4,
  "price": 4,
  "category_id": 400491,
  "logistic_id": [
    "stock_id": 0,
    "shipping_fee": 2.5,
    "stock_id": 0,
    "logistic_id": 80016,
    "is_free": false
  ],
  "item_sku": "",
  "stock_list": [
    "stock_id": 0,
    "item_id": "",
    "shop_sku": "",
    "model_id": 0
  ],
  "image": {
    "image_id_list": ["1434556437666206ff3c494568c03da6", "1434556437831203d05186f3ff5674f3d87", "1434556437666206ff3c494568c03da6", "00ac5d6431de5a04b7c1c6ff77340047", "a66d3152406e8a680012305af9880", "1343401733f1da696013e42a60e84fe"]
  },
  "attributes_list": [
    {
      "attribute_id": 100066,
      "attribute_value_list": [
        {
          "value_id": 3,
          "original_value_name": "",
          "value_unit": ""
        }
      ]
    }
  ],
  "item_dangerous": 0,
  "tax_exception_name": "cached and item from /getsign API",
  "original_price": 123,
  "image_id": "0",
  "item_status": "NORMAL",
  "has_variation": false,
  "wholesale_list": []
}
```

```json
{
  "item_type": "image",
  "tax": "",
  "item_name": "",
  "image_list": ["1434556437666206ff3c494568c03da6"]
}
```

```json
{
  "description_type": "extended",
  "seller_stock": [],
  "item_id": 0,
  "brand": {
    "brand_id": 0,
    "original_brand_name": ""
  }
}
```

```json
{
  "condition_id": {
    "id": "NEW",
    "id_pro_outlet": false,
    "data_id_shop": 3
  }
}
```

```json
{
  "logistic_info": [
    {
      "logistic_id": 12345,
      "enabled": true,
      "shipping_fee": 1.25,
      "size_id": 0,
      "is_free": false
    }
  ]
}
```

---

**`v1_product_get_item_base_info`를 호출하여 아울렛 상품이 올바르게 동기화되었는지 확인합니다.**

상품 정보를 확인하려면 **[ProductBase Info GetItemBaseInfo 문서](#)**를 참조하십시오.

---

## 3.2.2 상품 가져오기 (매핑 정보)

---

## 호출 샘플:

```json
{
  "item_list": [
    {
      "seller_stock": [],
      "item_id": 0
    }
  ],
  "need_tax_info": false,
  "need_complaint_policy": false
}
```

---

## 3.2.3 라벨 SKU를 아울렛 SKU로 동기화

메인 쇼핑몰 SKU를 아울렛 SKU로 동기화합니다.

1. `v1_product_update_item`을 호출하여 먼저 메인 상품을 업데이트합니다.
2. `v1_product_get_item_base_info/outlet`을 호출하여 아울렛 상품 정보가 올바르게 동기화되었는지 확인합니다.

메인 SKU로 관리되는 메인 아울렛 쇼핑몰은 아울렛 SKU로 동기화할 수 있습니다.

---

## 3.2.2 라벨 SKU 업데이트

## 3.2.3 채널 소개 개요

---

## 스택 분류기 vs 물류 분류기

| 상품 - 연락처 | SKOS | 2채널 |
|----------------|------|-----------|
| SKOS | SKOS | GlobalExpress 간접 |
| SKOS | SKOS | SPA 물류 |
| SKOS | SKOS | 3P 플랫폼 |
| SKOS - Instant II | SKOS | SPA 관련 + 2.skm |
| SKOS | SKOS | GlobalExpress + 1.skm |
| | SKOS | SPX 플랫폼 + 1.skm |
| SKOS - Instant II | SKOS | SPA 관련 + 2.skm |
| SKOS | SKOS | GlobalExpress + 2.skm |
| | SKOS | SPX Global 간접 + 2.skm |

---

**매핑된 상품에 대해 항상 동일한 물류 채널을 사용해야 합니다(관련 상품은 (판매되지 않은 경우) 푸시되어야 하고 Instant 2 상품(SKO3)은 항상 함께 태그되어야 합니다).**

---

## 상태 표

| 사례 | 현재 조건 | 작업 | 오류 |
|------|-------------------|--------|--------|
| [원래 표의 빈 셀] |

---

# WAF - 예산 및 세금

## 예산 설정

**80053: SPA Instant - 4 JSON**  
**80062: Grid/Lionel Instant - 4 JSON**

*개발자는 예산, 세금 및 프로모션이 이러한 채널에 어떻게 적용되는지 이해해야 합니다.*

### 3.2.3 RI(Instant 채널)에 대한 제한 사항 요약

Install 2 (xxx) BANG은 6시간 동안 Instant (xxx)되거나 함께 지원해야 합니다.

---

| 사례 | 현재 조건 | 작업 | 오류 |
|------|------------------|---------|---------|
| 1 | 8000 Instant: ON<br>80053: SPA Instant<br>: ON<br>• KNFG Gridland<br>Instant : ON<br>• 80062 : Grid/Lionel<br>Instant : 2 hours, OFF<br>• 80053 : SPA Instant<br>: 4 JSON, OFF<br>• 80061 Grid/Lionel<br>Instant : 4 JSON, OFF<br>• KNFG Grid/Express<br>Instant : 2 JSON, OFF | 끄기<br>80044 | SPA Instant는 끌 수 없습니다. 다음 채널 중 하나 이상을 활성화하십시오: {Grid/and Instant, Grid/Express Instant, Grid/Express Instant : 4 JSON, Grid/Express Instant : 2 JSON, SPA Instant : 4 JSON Grid/and Instant, 4 JSON} Grid/and comes Instant : 4 JSON, Grid/and Instant : 4 JSON} |
| 2 | • 80053 SPA Instant<br>: 4 JSON, OFF<br>• KNFG Gridland<br>Instant : ON<br>• 80062 : Grid/Lionel<br>Instant : 2 JSON, OFF | 끄기<br>80044 | SPA Instant : 4 JSON Gridland는 끌 수 없습니다. 다음 채널 중 하나 이상을 활성화하십시오: {SPA Instant, Grid/Express Instant, Grid/Express Instant, Grid/Express Instant : 4 JSON, Grid/and Instant : 2 JSON, Grid/Express Instant : 4 JSON, Grid/and Instant : 4 JSON} Grid/and Instant : 4 JSON} |
| 3 | 8000 Instant: OFF<br>80053 SPA Instant<br>: ON<br>• KNFG Gridland<br>Instant : OFF<br>• 80062 : Grid/Lionel<br>Instant : 2 hours, ON<br>• 80053 : SPA Instant<br>: 4 JSON, OFF<br>• KNFG Grid/Express<br>Instant : 4 JSON, OFF<br>• 80062 Grid/Lionel<br>Instant : 2 JSON, OFF | | |
| | 8000 Instant: 6 hours | | |
| 4 | • 80053 SPA Instant<br>: 4 JSON, OFF<br>• 80062 : Grid/Lionel<br>Instant : 4 JSON, OFF<br>• KNFG Grid/Express<br>Instant : 2 JSON, OFF | | |
| 5 | Seller is now unselected for<br>8000 instant (seller)<br>• 80053 SPA Instant<br>: 4 JSON, OFF<br>• KNFG Gridland<br>Instant : OFF<br>• 80053 SPA Instant<br>: 4 JSON, OFF<br>• KNFG Grid/Express<br>Instant : 2 JSON, OFF | 끄기<br>80044 | SPA Instant 채널을 끌 수 없습니다. 다음 채널을 활성화하십시오: {Grid/and Instant, Grid/Express Instant, Grid/Express Instant} |
| 6 | 8000 Instant: ON<br>80053 SPA Instant<br>: 4 JSON, ON<br>• KNFG Gridland<br>Instant : OFF<br>• 80053 SPA Instant<br>: 4 JSON, OFF<br>• 80062 Grid/Lionel<br>Instant : 4 JSON, OFF<br>• KNFG Grid/Express<br>Instant : 2 JSON, OFF | 끄기<br>80053 | SPA Instant : 4 JSON Gridland는 끌 수 없습니다. 다음 채널 중 하나 이상을 활성화하십시오: {SPA Instant, Grid/and Instant, Grid/Express Instant, Grid/Express Instant, 80062 : Grid/Lionel Instant : 4 JSON, SPA Instant : 2 JSON} |
| 7 | 8000 Instant: 6 hours<br>• 80053 SPA Instant<br>: 4 JSON, OFF<br>• KNFG Gridland<br>Instant : ON<br>• 80062 Grid/Lionel<br>Instant :

| 마스크 채널 | 물류 채널 |
|--------------|-------------------|
| 0 | 70124 Instant Delivery - pulled (with no rn-ff) |
| 1 | 70125 Instant Delivery - pulled (with no rn-directions) |

---

## 3.2.4 주문 관리

**쇼퍼 운영 콘솔의 주문 페이지에서 판매자 주문을 이전하고 실제 주문 생성을 선택합니다.**

[THIS IS FIGURE: 다양한 열과 "실제 주문 생성" 버튼이 있는 주문 관리 인터페이스를 보여주는 스크린샷]

---

**상점, 품목 및 배송 옵션을 선택한 다음 생성을 클릭하여 주문을 생성합니다.**

[THIS IS FIGURE: 상점 텍스트 주문, 품목, 상점 ID, 주소 및 배송 옵션 필드가 있는 주문 생성 양식을 보여주는 스크린샷]

---

**주문 정보에는 주문 SN, 품목 ID, 상태, 업데이트 시간 및 스냅 ID가 포함됩니다. 여기에서 주문을 픽업, 배송 또는 삭제할 수도 있습니다.**

[THIS IS FIGURE: 주문 정보, 품목 세부 정보 및 작업 버튼 열이 있는 주문 세부 정보 테이블을 보여주는 스크린샷]

---

**다음 단계에 대한 자세한 내용은 주문 관리 설정을 참조하십시오.**

---

## 3.2.4 반품 및 환불 관리

자세한 내용은 반품 및 환불 관리 설정을 참조하십시오.

---

## 3.2.5 재무: 수입 및 지갑 거래 내역 가져오기

**참고: 판매자가 API를 통해 지갑 명세서를 검색하거나 짧은 시간 내에 API를 호출하여 지갑 거래 내역을 가져오는 경우 또는 긴 시간을 사용하는 경우 괜찮습니다(응답은 모든 지갑 변경 사항을 포함하는 집계된 문서를 반환합니다).**

| 기능 | Open API | 비고 |
|----------|----------|---------|
| generate_income | v2.payment.generate_income_report | 요청 파라미터: release_time_from 및 release_time_to |
| get income | v2.payment.get_income_report | |
| get timeline | v2.payment.get_income_report | 수입 보고서 상태를 쿼리하고 다운로드 URL을 제공합니다. |
| | v2.payment.download_income_report | 수입 보고서 파일을 다운로드합니다. |

---

# API 문서 발췌

## 수입 명세서, 보고서 및 지갑 거래 내역

Mart shop_id를 사용하여 API를 호출하여 수입 명세서, 수입 보고서 또는 지갑 거래 내역을 생성하고 검색할 수 있습니다. Mart shop_id를 사용하는 경우 응답은 Mart 상점 아래의 모든 아울렛 상점을 포함하는 집계된 문서를 반환합니다.

### API 기능

| 기능 | Open API | 비고 |
|----------|----------|--------|
| generate income report | v2.payment.generate_income_report | 요청 파라미터: release_time_from 및 release_time_to. |
| get income report | v2.payment.get_income_report | income_report_id를 전달하여 수입 보고서 상태를 쿼리하고 파일 링크를 제공합니다. |
| generate income statement | v2.payment.generate_income_statement | 요청 파라미터: statement_type=1은 주간 명세서를 의미합니다. statement_type=2는 월간 명세서를 의미합니다. |
| get income statement | v2.payment.get_income_statement | income_statement_id를 전달하여 수입 명세서 상태를 쿼리하고 파일 링크를 제공합니다. |
| get wallet transaction | v2.payment.get_wallet_transaction_list | 지갑의 거래 기록을 가져옵니다. |

## 3.2.6 상점 설정

1. v2.shop.update_profile을 호출하여 상점 이름, 로고 및 설명을 업데이트할 수 있습니다.

2. v2.logistics.update_operating_hours를 호출하여 운영 시간을 업데이트할 수 있습니다.

**참고:**
- 제공된 값은 v2.logistics.get_operating_hour_restrictions에서 검색된 제한 사항을 준수해야 합니다.
- 이 API는 덮어쓰기 업데이트를 수행합니다. 픽업 운영 시간을 업데이트할 때 변경되지 않은 세그먼트까지 포함하여 모든 세그먼트를 포함해야 합니다.

## 4. 푸시 메커니즘

푸시 메커니즘은 Shopee의 시스템 알림 사전으로, 제품, 주문, 반품, 마케팅, 안정성 및 일반 마켓플레이스 업데이트를 다룹니다.

Instant Mart의 경우 order_status_push에 연결하는 것이 좋습니다. 자세한 내용은 order_status_push 문서를 참조하십시오.

## 5. FAQ 및 도움 요청

일반적인 질문은 FAQ를 참조하십시오. API 연결에 문제가 발생하면 [여기]에서 Shopee 제품 지원 팀에 티켓을 제출할 수 있습니다.

## 사용 사례

1. 소매 시스템을 Instant Mart의 Open API와 통합.
2. 여러 Instant Mart 아울렛에서 제품 및 재고 관리.
3. Instant Mart 플랫폼을 통한 주문 및 배송 처리.
4. Instant Mart 생태계 내에서 반품 및 환불 처리.
5. 소매 시스템과 Instant Mart 간의 재무 데이터 동기화.

## 관련 API

- v1_product_add_item
- v1_product_update_item
- v1_product_publish_item_to_outlet_shop

---

## 원문 (English)

### Summary

This developer guide provides information on integrating with Instant Mart's Open API. It covers sandbox account setup, authorization, testing, and key API functions for managing products, orders, returns, financials, and shop settings within the Instant Mart ecosystem, including main shops and outlet shops.

### Content

# API Host Products - Instant Mart Integration-Dev

## Background:
Instant Mart is a open-source shop designed designed to support retailers who operate both online management and branch key system through integrations

- Each Mart Shop (Official Shop) acts as the headquarters, responsible for managing global stock, creating items, and syncing financial reports.

- Multiple Outlets (Branch shops) (and users) the mart, representing individual shop locations. Outlets are responsible for local stock, handling daily operations, like managing orders, maintaining stock, and placing purchase.

After the Open API Integration, Mart retailers can choose whether to manage SKUs at the merchant level (shared across all outlets) or at the outlet level (independent per branch).

Retailers integrating with Open API - retailers must understand their structure design plan they select will affect how they define the flow for both merchant and its outlet shops in the RDC portal. (Each integrating account will link to a single entity in the RDC system, either at headquarters or branch level; this determines how developers to build API functions that respect the full capabilities list of return of Outlet shops.)

---

## 1. Sandbox Test Account Setup

### 1.1 Create Sandbox Test Account

You can start by accessing the [SandGuide](#) page on Opened Open Platform Console. Select Mart/Outlet shops module to create sandbox shops.

*Note: Instant Mart sandbox testing requires a two-step scheduler. Scheduler 1/2 must be enabled before Mart 4 Outlet Shop(s) activated.*

---

## 2. Grant authorization

### 2.1 Shop Authorization
Find Instant Mart Partner url on the nav page

*[Screenshot showing authorization flow with "App List" and shop selection interface]*

Click Authorize — log in with Mart/Outlet shop account to complete authorization

*[Screenshot showing "Login to Authorize Shoppe OpenPlatform API" interface with authorization button]*

*[Additional screenshots showing authorization process steps]*

Refer to the [Authorization and Authorization](#) article for more information.

---

## 3. Testing Process

### 3.1 Open API Testing
Use the API Test Tool at use Postman to test open API

*[Screenshot showing API testing interface with request/response panels]*

---

### 3.2 Key Open API Functions

Key API functionalities required for M-Mart Project:

| Function | Sub_function |
|----------|--------------|
| Sandbox | |
| Product | Preparation for New SKU |
| | Creating New SKU |
| | Update Existing SKU |
| | Creating SKU Variants |
| | Stock & Price Management |
| | General SKU Management |
| Order | Get Order List & Details |
| | Cancelling Order |
| | Request Shipment |
| | Generate Tracking Number |
| Return & Refund | Get Return List & Details |
| | Accept Return/Refund |
| | Dispute Return/Refund |
| Financials | |

---

**Note:** This extraction maintains the original English text structure from the developer guide screenshot, preserving headings, lists, tables, and organizational hierarchy in Markdown format.

---

# API Documentation Extract

## Returns & Refund
- Generate Tracking Number
- Generate Airway Bill
- Get Return List & Details
- Accept Return/Refund
- Dispute Return/Refund

## Financials
- Get Escrow Detail
- Get Wallet Transaction
- Generate Income Statement
- Generate Income Report

## Shop Setting
- Shop Profile Update
- Shop Operational Hours

---

## 3.2.1 MarkOutofStock relationship

Check relationship: `or_shop_gir__shop_info`

With input_shop_gir = related_gir do system do the fulfill the fulfill Mart Shop is related_shop_gir (Main Shop ID)

**Note:** Shop gir is related to main shop so the steps to call will need to know (Main Shop ID).

## 3.2.1.1 Get or shop gir_shop info

When calling `or shop gir_shop_info` with a main_shop_id, it will returns a list of all Outlet shop IDs under that Main shop.

### Response example:

```json
{
  "sells_shop": "1700820216",
  "toplist_shop": "1704576796",
  "or_gir": "fass",
  "or_main_shop": "fass",
  "or_main_shop": "fass",
  "or_next_shop": "fass",
  "or_outlet_shop": "fass",
  "or_shop_grp": "fass",
  "or_shop_info": "fass",
  "items_main_shop_id": 0,
  "overstock_id": null,
  "outlet_shop_add_sett_1": "",
  "outlet_shop_id": 226402201,
  "outlet_shop_id": 226402202,
  "outlet_shop_id": 226402203,
  "outlet_shop_id": 226402204,
  "trigger": "123",
  "trigger": "no",
  "incomin_gir": "vg_K2Ni9xY...",
  "incomin_outlet_shop_id": "17JdA2536GG37",
  "keep_some": "vg_6DaAhG2Y43G9Y8/GqF",
  "outint": "IGz9M0n.."
}
```

---

## 3.2.1.2 Use or shop gir_shop_info with a outlet shop

When calling `or shop gir_shop_info` with an outlet_shop_id, it will return the corresponding main shop ID and related data.

### Response sample:

```json
{
  "sells_shop": "1704156903",
  "toplist_shop": "1704723399",
  "or_gir": "fass",
  "or_main_shop": "fass",
  "or_main_shop": "fass",
  "or_next_shop": "fass",
  "or_shop": "fass",
  "or_gir": "fass",
  "or_signature_shop": "",
  "items_main_shop_id": 226402201,
  "items_main_shop_id": 0,
  "overstock_id": 226621997,
  "overstock_id": null,
  "message": "",
  "trigger": "123",
  "trigger": "no",
  "incomin_gir": "vg_K7D3LG4YZ...",
  "main_shop_id": "9Y1Io4G8/G9Y6/10aBt",
  "shop_some": "vg_J1G02k446G17kK6BB57",
  "outint": "IGz9M0n.."
}
```

---

## 3.2.2 Product Management

[THIS IS DIAGRAM: A flowchart showing product management hierarchy with boxes connected by arrows, including elements like "product_db_send", "db product info_in_detail info", and "db sent while descending_in_a"]

### Fields relationship between Main Shop and Outlet Shop:

| Module | Field | MART SKU Management | Outlet SKU Management |
|--------|-------|--------------------|-----------------------|
| Basic Information | title | Yes | No |
| | video | Yes | No |
| | image | Yes | No |
| | status | Yes | Yes |
| | category | Yes | No |
| Specification | brand | Yes | No |
| | attribute | Yes | No |
| | size chart | Yes | No |
| Specification | variation name | Yes | No |
| | variation option | Yes | No |
| | variation image | Yes | No |
| | price | No | Yes |
| | purchase limit | No | Yes |
| | subchannel | No | Yes |
| | discount | No | Yes |
| | model status | Yes | Yes |
| | module status SKU | Yes | No |
| | weight | Yes | No |
| Shipping | dimension | Yes | No |
| | logistics channel | No | Yes |
| | preparation channel | No | Yes |
| Others | DTS | No | Yes |
| | CBS | No | Yes |
| | condition | Yes | No |

---

**Note:** This extraction maintains the original structure, headings, code blocks, and table formatting from the developer guide screenshot.

---

# Developer Guide - Feature Comparison and Setup

## Feature Comparison Table

| Feature | Alternative | Yes | No |
|---------|-------------|-----|-----|
| Shipping | - | - | No |
| - | bundles (batched) | - | Yes |
| - | installation channel | No | Yes |
| Others | DTS | No | Yes |
| - | seller SKU | Yes | No |
| - | condition | - | - |

## Key Operations Include

1. Add label SKU --> `v1_product_add_item`
2. Update label SKU --> `v1_product_update_item` (includes price, stock, product_price, price_of_outlet level)
3. Publish Code SKU --> `v1_product_publish_item_to_outlet_shop`
4. Sync Sale SKU (+) --> `v1_product_update_item`

---

## 3.2.1 Add label SKU

Add label SKU by calling `v1_product_add_item` using main shop_id.

Refer to the **[detailed knowledgebase article](#)** for more information.

---

## Calling sample:

```json
{
  "image_price_list": 123.5,
  "description": "cached and item from /getsign API",
  "weight": 1.5,
  "image_price_list": 123.5,
  "description": "cached and item from /getsign API",
  "weight": 1.5,
  "Itemplace_height": 2,
  "Itemplace_length": 3,
  "Itemplace_width": 4,
  "price": 4,
  "category_id": 400491,
  "logistic_id": [
    "stock_id": 0,
    "shipping_fee": 2.5,
    "stock_id": 0,
    "logistic_id": 80016,
    "is_free": false
  ],
  "item_sku": "",
  "stock_list": [
    "stock_id": 0,
    "item_id": "",
    "shop_sku": "",
    "model_id": 0
  ],
  "image": {
    "image_id_list": ["1434556437666206ff3c494568c03da6", "1434556437831203d05186f3ff5674f3d87", "1434556437666206ff3c494568c03da6", "00ac5d6431de5a04b7c1c6ff77340047", "a66d3152406e8a680012305af9880", "1343401733f1da696013e42a60e84fe"]
  },
  "attributes_list": [
    {
      "attribute_id": 100066,
      "attribute_value_list": [
        {
          "value_id": 3,
          "original_value_name": "",
          "value_unit": ""
        }
      ]
    }
  ],
  "item_dangerous": 0,
  "tax_exception_name": "cached and item from /getsign API",
  "original_price": 123,
  "image_id": "0",
  "item_status": "NORMAL",
  "has_variation": false,
  "wholesale_list": []
}
```

```json
{
  "item_type": "image",
  "tax": "",
  "item_name": "",
  "image_list": ["1434556437666206ff3c494568c03da6"]
}
```

```json
{
  "description_type": "extended",
  "seller_stock": [],
  "item_id": 0,
  "brand": {
    "brand_id": 0,
    "original_brand_name": ""
  }
}
```

```json
{
  "condition_id": {
    "id": "NEW",
    "id_pro_outlet": false,
    "data_id_shop": 3
  }
}
```

```json
{
  "logistic_info": [
    {
      "logistic_id": 12345,
      "enabled": true,
      "shipping_fee": 1.25,
      "size_id": 0,
      "is_free": false
    }
  ]
}
```

---

**Call `v1_product_get_item_base_info` to check if outlet item sync correctly.**

Refer to the **[ProductBase Info GetItemBaseInfo article](#)** to check the item information.

---

## 3.2.2 Get the item (mapping info)

---

## Calling sample:

```json
{
  "item_list": [
    {
      "seller_stock": [],
      "item_id": 0
    }
  ],
  "need_tax_info": false,
  "need_complaint_policy": false
}
```

---

## 3.2.3 Sync label SKU to Outlet SKU

Sync the main shop SKU to the outlet sku.

1. Call `v1_product_update_item` to update the main item first
2. Call `v1_product_get_item_base_info/outlet` to check that the outlet item info syncs correctly

Main Outlet shop managed by main sku can be sync to outlet sku

---

## 3.2.2 Update label SKU

## 3.2.3 Channel Introduce Overview

---

## Stack Classifier vs Logistics Classifier

| Item - contact | SKOS | 2-channel | 
|----------------|------|-----------|
| SKOS | SKOS | GlobalExpress Indirect |
| SKOS | SKOS | SPA Logistics |
| SKOS | SKOS | 3P platform |
| SKOS - Instant II | SKOS | SPA Related + 2.skm |
| SKOS | SKOS | GlobalExpress + 1.skm |
| | SKOS | SPX platform + 1.skm |
| SKOS - Instant II | SKOS | SPA Related + 2.skm |
| SKOS | SKOS | GlobalExpress + 2.skm |
| | SKOS | SPX Global Indirect + 2.skm |

---

**We suppose should always use the same logistics channel(s) for those mapped items together (related items should be pushed (when not sold?) and instant 2 items (SKO3) should always be tagged until together)**

---

## Status Table

| Case | Current Condition | Action | Errors |
|------|-------------------|--------|--------|
| [Empty cells in original table] |

---

# WAF - Budget & Taxes

## Budget Settings

**80053: SPA Instant - 4 JSON**  
**80062: Grid/Lionel Instant - 4 JSON**

*Developers should understand how Budget, Taxes and Promotions lay these channels.*

### 3.2.3 Summary of the limitation about RI (Instant Channels)

Install 2 (xxx) BANG will instant 6 hours (xxx) or should support together.

---

| Case | Current Condition | Action | Errors |
|------|------------------|---------|---------|
| 1 | 8000 Instant: ON<br>80053: SPA Instant<br>: ON<br>• KNFG Gridland<br>Instant : ON<br>• 80062 : Grid/Lionel<br>Instant : 2 hours, OFF<br>• 80053 : SPA Instant<br>: 4 JSON, OFF<br>• 80061 Grid/Lionel<br>Instant : 4 JSON, OFF<br>• KNFG Grid/Express<br>Instant : 2 JSON, OFF | Turn off<br>80044 | SPA Instant cannot be turned off. Please enable at least 1 of the following channel(s): {Grid/and Instant, Grid/Express Instant, Grid/Express Instant : 4 JSON, Grid/Express Instant : 2 JSON, SPA Instant : 4 JSON Grid/and Instant, 4 JSON} Grid/and comes Instant : 4 JSON, Grid/and Instant : 4 JSON} |
| 2 | • 80053 SPA Instant<br>: 4 JSON, OFF<br>• KNFG Gridland<br>Instant : ON<br>• 80062 : Grid/Lionel<br>Instant : 2 JSON, OFF | Turn off<br>80044 | SPA Instant : 4 JSON Gridland cannot be turned off. Please enable at least 1 of the following channel(s): {SPA Instant, Grid/Express Instant, Grid/Express Instant, Grid/Express Instant : 4 JSON, Grid/and Instant : 2 JSON, Grid/Express Instant : 4 JSON, Grid/and Instant : 4 JSON} Grid/and Instant : 4 JSON} |
| 3 | 8000 Instant: OFF<br>80053 SPA Instant<br>: ON<br>• KNFG Gridland<br>Instant : OFF<br>• 80062 : Grid/Lionel<br>Instant : 2 hours, ON<br>• 80053 : SPA Instant<br>: 4 JSON, OFF<br>• KNFG Grid/Express<br>Instant : 4 JSON, OFF<br>• 80062 Grid/Lionel<br>Instant : 2 JSON, OFF | | |
| | 8000 Instant: 6 hours | | |
| 4 | • 80053 SPA Instant<br>: 4 JSON, OFF<br>• 80062 : Grid/Lionel<br>Instant : 4 JSON, OFF<br>• KNFG Grid/Express<br>Instant : 2 JSON, OFF | | |
| 5 | Seller is now unselected for<br>8000 instant (seller)<br>• 80053 SPA Instant<br>: 4 JSON, OFF<br>• KNFG Gridland<br>Instant : OFF<br>• 80053 SPA Instant<br>: 4 JSON, OFF<br>• KNFG Grid/Express<br>Instant : 2 JSON, OFF | Turn off<br>80044 | SPA Instant channel cannot be turned off. Please enable the following channel(s): {Grid/and Instant, Grid/Express Instant, Grid/Express Instant} |
| 6 | 8000 Instant: ON<br>80053 SPA Instant<br>: 4 JSON, ON<br>• KNFG Gridland<br>Instant : OFF<br>• 80053 SPA Instant<br>: 4 JSON, OFF<br>• 80062 Grid/Lionel<br>Instant : 4 JSON, OFF<br>• KNFG Grid/Express<br>Instant : 2 JSON, OFF | Turn off<br>80053 | SPA Instant : 4 JSON Gridland cannot be turned off. Please enable at least 1 of the following channel(s): {SPA Instant, Grid/and Instant, Grid/Express Instant, Grid/Express Instant, 80062 : Grid/Lionel Instant : 4 JSON, SPA Instant : 2 JSON} |
| 7 | 8000 Instant: 6 hours<br>• 80053 SPA Instant<br>: 4 JSON, OFF<br>• KNFG Gridland<br>Instant : ON<br>• 80062 Grid/Lionel<br>Instant : 4 JSON, OFF<br>• KNFG Grid/Express<br>Instant : 2 JSON, OFF | | |

---

## 3.2.5 Instant Channel Dependency in TH

### 3.2.1 Indexed numbers

| Mask Channel | Logistics Channel |
|--------------|-------------------|
| 0 | 70124 Instant Delivery - pulled (with no rn-ff) |
| 1 | 70125 Instant Delivery - pulled (with no rn-directions) |

---

## 3.2.4 Order Management

**Transfer a seller order on the Orders page on the Shopper Operation Console and select Create Real Order**

[THIS IS FIGURE: Screenshot showing order management interface with various columns and a "Create Real Order" button]

---

**Choose the shop, items, and shipping option, then click Create to generate the order.**

[THIS IS FIGURE: Screenshot showing order creation form with fields for shop text order, items, shop id, address, and shipping options]

---

**The order information includes Order SN, Item ID, Status, Update Time and Snap ID. You can also pick up, deliver, or delete the order from here.**

[THIS IS FIGURE: Screenshot showing order details table with columns for order information, item details, and action buttons]

---

**Refer to the Order Management setup for more information about following steps.**

---

## 3.2.4 Return & Refund Management

Refer to the Return & Refund Management setup for more information.

---

## 3.2.5 Finance: Get income and wallet transactions

**Note: If a merchant retrieves their wallet statement via the API or wallet transactions by calling the API with a short time, or when using a long time, it is fine (response will return an aggregated document that includes all wallet change.**

| Function | Open API | Remark |
|----------|----------|---------|
| generate_income | v2.payment.generate_income_report | Request Parameter: release_time_from and release_time_to |
| get income | v2.payment.get_income_report | |
| get timeline | v2.payment.get_income_report | To query income report status and provide the download URL |
| | v2.payment.download_income_report | To download the income report file |

---

# API Documentation Extract

## Income Statement, Report, and Wallet Transactions

You can generate and retrieve the income statement, income report, or wallet transactions by calling the APIs with a Mart shop_id. When using a Mart shop_id, the response will return an aggregated document that includes all outlet shops under the Mart shop.

### API Functions

| Function | Open API | Remark |
|----------|----------|--------|
| generate income report | v2.payment.generate_income_report | Request Parameter: release_time_from and release_time_to. |
| get income report | v2.payment.get_income_report | To query income report status and provide file link by passing income_report_id. |
| generate income statement | v2.payment.generate_income_statement | Request Parameter:statement_type=1 means weakly statement;statement_type=2 means monthly statement |
| get income statement | v2.payment.get_income_statement | To query income statement status and provide a file link by passing income_statement_id |
| get wallet transaction | v2.payment.get_wallet_transaction_list | Get the transaction records of the wallet. |

## 3.2.6 Shop Setting

1. You can update the shop name, logo, and description by calling v2.shop.update_profile.

2. You can update the operating hours by calling v2.logistics.update_operating_hours.

**Note:**
- The values provided must comply with the restrictions retrieved from v2.logistics.get_operating_hour_restrictions.
- This API performs overwriting updates. When updating pickup operating hours, you must include all segments, even those that remain unchanged.

## 4. Push Mechanism

The push mechanism is a dictionary of system notifications from Shopee, covering product, order, return, marketing, stability, and general marketplace updates.

For Instant Mart, it is recommended to connect to order_status_push. For more details, please refer to the order_status_push article.

## 5. FAQ & Raise for Help

For common questions, please refer to the FAQs. If you encounter any issues with API connection, you can raise a ticket to the Shopee Product Support team [here].

---

**문서 ID**: developer-guide.643
**플랫폼**: shopee
**URL**: https://open.shopee.com/developer-guide/643
**처리 완료**: 2025-10-16T09:32:11
