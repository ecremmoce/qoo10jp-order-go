# 알 수 없음

**카테고리**: 분류되지 않음
**난이도**: medium
**중요도**: 3/5
**최종 업데이트**: 2025-10-16T08:23:26

## 개요

요약 정보를 사용할 수 없음

## 주요 키워드



## 본문

# V2 SAR 호출 흐름

## 공지 콘솔

### 시작하기 > V2 SAR 호출 흐름

---

## 주요 흐름도

```
[시작]
    ↓
[www.www-vendor.xhtml]
    ↓
[1단계: 벤더 xhtml 선택] ──┐
[2단계: 제출 xhtml 선택] ─────┤
[3단계: xhtml xhtml 강제 수정]─┤
[xx: 승인 xxx- xxxxxxxxx -xxxx] ─┘
    ↓
[xxxxxxx xxx- xxxxxxx]
    ↓
[xxxxxxxx xxx- xxxxxxxxx]
    ↓
[xxxxxxxx xxx- xxxxxx xxx]
    ↓
[xxxxxxxx xxx- xxxx -xxx]
    ↓
[xxxxxxx xxx- xxxxxxx xx]
    ↓
[xxxxxxx xxxx- xxxx]
    ↓
[xxxxxxxx xxxx- xxx- xxxxxxx] ←→ [xxxxxxxxxxxxxxx xxxxx- xxxx]
    ↓
[종료]
```

**참고: 점선은 선택적 흐름입니다.**

---

## 주문 상태 흐름

### 흐름 설명

[왼쪽 참고 사항:]
- 주문 상태는 현재 이행 상태를 나타냅니다.
- 주문이 접수되면 초기 상태는 "보류 중"입니다.
- 주문이 이행됨에 따라 상태가 변경됩니다.
- 여러 패키지가 서로 다른 상태를 가질 수 있습니다.
- 최종 상태는 완료 또는 취소를 나타냅니다.

[중앙 흐름:]
```
[시작]
    ↓
[보류 중]
    ↓
[처리 중] ←→ [보류]
    ↓
[승인됨]
    ↓
[피킹]
    ↓
[포장]
    ↓
[배송 준비 완료]
    ↓
[배송됨]
    ↓
[배달 완료]
    ↓
[종료]
```

[오른쪽 참고 사항:]
- 보류 상태는 수동 개입이 필요합니다.
- 주문은 다양한 단계에서 취소될 수 있습니다.
- 배송됨 상태에는 추적 정보가 포함됩니다.
- 배달 완료 상태는 수령을 확인합니다.
- 상태 업데이트는 알림을 트리거합니다.

### 상태 전환 테이블

| 현재 상태 | 다음 상태 | 조건 |
|---------------|-------------|-----------|
| 보류 중 | 처리 중 | 모든 프로세스에 의해 |
| 보류 중 | 보류 | 모든 프로세스에 의해 |
| 보류 | 처리 중 | 모든 프로세스에 의해 |
| 처리 중 | 승인됨 | 모든 프로세스에 의해 |
| 처리 중 | 취소됨/거부됨 | 모든 요청에 의해 |
| 승인됨 | 피킹 | 이행 시작됨 |
| 피킹 | 포장 | 품목 수집됨 |
| 포장 | 배송 준비 완료 | 패키지 준비됨 |
| 배송 준비 완료 | 배송됨 | 운송업체 픽업 |
| 배송됨 | 배달 완료 | 배달 확인됨 |

### LeadStatus

| 상태 | 설명 |
|--------|-------------|
| 부적격 | --- |
| 신규 | --- |
| 작업 중 | 모든 처리된 프로세스에 의해 |
| 육성 중 | --- |

**참고: 점선은 선택적 흐름입니다.**

---

## 패키지 이행 상태 흐름

### 흐름 설명

[왼쪽 참고 사항:]
- 패키지는 개별 배송 단위를 나타냅니다.
- 각 패키지에는 독립적인 추적이 있습니다.
- 패키지 상태는 주문 상태와 별개입니다.
- 주문당 여러 패키지가 가능합니다.
- 패키지 무게와 크기가 기록됩니다.
- 운송업체는 패키지 수준에서 할당됩니다.

[중앙 흐름:]
```
[시작]
    ↓
[생성됨]
    ↓
[처리 중] ←→ [보류]
    ↓
[픽업 준비 완료]
    ↓
[픽업됨] ──→ [이동 중]
    ↓              ↓
[배달 출발]
    ↓              ↓
[배달 완료] ←─────┘
    ↓
[종료]
```

[오른쪽 참고 사항:]
- 보류는 진행하기 전에 해결이 필요합니다.
- 추적 번호는 픽업 시 할당됩니다.
- 이동 중에는 체크포인트 업데이트가 포함됩니다.
- 배달 시도가 기록됩니다.
- 배달 실패는 예외를 생성합니다.
- 배달에 서명이 필요할 수 있습니다.

**참고: 점선은 선택적 흐름입니다.**

---

## 배송 준비 및 추적 번호 가져오기 & 송장 출력/청구

### 흐름 설명

[왼쪽 내용:]
- 배송 준비는 운송업체 예약을 시작합니다.
- 시스템은 배송 주소를 확인합니다.
- 규칙에 따라 운송업체가 선택됩니다.
- 배송비가 계산됩니다.
- 라벨 생성 프로세스가 시작됩니다.
- 운송업체에서 추적 번호를 가져옵니다.
- 항공 운송장/배송 라벨이 인쇄됩니다.
- 필요한 경우 세관에 대한 서류가 준비됩니다.

[중앙 흐름 - 배송 준비:]
```
[시작]
    ↓
[주소 확인]
    ↓
[운송업체 선택] ←→ [비용 계산]
    ↓
[배송 예약]
    ↓
[추적 번호 가져오기]
    ↓
[라벨 생성]
    ↓
[항공 운송장 인쇄]
    ↓
[종료]
```

[중앙 흐름 - 추적 가져오기:]
```
[시작]
    ↓
[운송업체 API 쿼리]
    ↓
[추적 데이터 검색]
    ↓
[시스템 업데이트]
    ↓
[고객에게 알림]
    ↓
[종료]
```

[오른쪽 내용:]
- 추적 번호 형식은 운송업체에 따라 다릅니다.
- API를 통한 실시간 추적 업데이트
- 고객 알림에는 추적 링크가 포함됩니다.
- 라벨은 운송업체 사양을 충족해야 합니다.
- 바코드 확인이 중요합니다.
- 국제 상업 송장
- 포장 명세서가 패키지에 포함됩니다.
- 배송 증명 보관

**참고: 점선은 선택적 흐름입니다.**

---

## [국경 간] 판매자 배송 추적

### 흐름 설명

[왼쪽 참고 사항:]
- 국경 간 배송에는 통관이 필요합니다.
- 수출 서류는 정확해야 합니다.
- 제품에 필요한 HS 코드
- 신고 가격은 관세에 영향을 미칩니다.
- 판매자는 수출 규정 준수 책임이 있습니다.

[중앙 흐름:]
```
[시작]
    ↓
[수출 서류 준비]
    ↓
[세관 신고서 제출] ←→ [세관 검토]
    ↓
[통관]
    ↓
[목적지 국가로 배송]
    ↓
[수입 통관 절차]
    ↓
[고객에게 배달]
    ↓
[종료]
```

[오른쪽 참고 사항:]
- 수입 관세는 목적지에서 적용될 수 있습니다.
- 추적에는 세관 검문소가 포함됩니다.
- 세관 검사 중 지연 가능성
- 고객은 관세를 지불해야 할 수 있습니다.
- 국경 간 반품이 더 복잡합니다.

**참고: 점선은 선택적 흐름입니다.**

---

## [국경 간-CNSC] 글로벌 상품 흐름

[이 섹션은 아래에 계속되는 다른 흐름도의 제목/머리글인 것 같습니다.]

---

# [국경 간-CNSC] 글로벌 상품 흐름

## 创建类目MTSKU

### Start

↓

**global_product.get_authorize_shop**

↓

**创建全球商品的第一步是准备授权店铺**
通过授权店铺，可以获得关联的类目信息，从而进行后续商品的发布。关于授权店铺的详细内容，请参考"授权店铺"章节。

↓

**global_product.get_category_list**

↓

**通过全球商品创建的第一步，选择类目开始创建商品流程。如果不确定类目信息，可以通过类目查询接口获取。通过类目信息，可以获取该类目下可用的品牌列表、属性列表等信息。选择好类目后，就可以进入商品信息填写阶段。**

↓

**global_product.get_category_detail**

↓

**global_product.get_brand_list**

↓

**global_product.get_global_product_category_attribute**

↓

**global_product.get_global_product_category_rule**

↓

**global_product_sku.global_product_sku_add**

↓

**global_product.add_global_product**

↓

**获取全球商品创建后，通过全球商品ID，调用全球商品查询接口，可以查询到创建的全球商品信息。**

↓

**global_product.get_global_product_list**

↓

**全球商品创建完成后，就可以查询到相关的全球商品信息。接下来需要将全球商品发布到店铺。**

↓

**global_product.get_global_product_detail**

↓

**全球商品创建完成后，就可以通过全球商品详情接口，查询全球商品的详细信息。**

↓

**global_product_publish.add_publish_task**

↓

**global_product_publish.get_publish_task_detail**

↓

**global_product_publish.get_publish_task_result**

↓

**下发发布任务后，系统会异步处理发布请求。可以通过查询发布任务结果接口，获取发布结果。发布成功后，就可以在店铺中看到商品了。**

↓

### End

---

## 添加MIPSKU产品流

### Start

↓

**global_product.add_global_product**

↓

**创建MIPSKU商品时，通过全球商品创建接口，可以直接创建MIPSKU商品。创建时需要指定商品类型为MIPSKU，并且填写相关的商品信息。创建完成后，系统会自动生成MIPSKU商品ID。**

↓

**global_product.get_global_product_detail**

↓

**global_product_sku.global_product_sku_add**

↓

**global_product_sku.global_product_sku_edit**

↓

### End

---

## 发布MIPSKU产品

### Start

↓

**global_product.add_global_product**

↓

**创建MIPSKU商品时需要注意以下几点：
1. 商品类型需要设置为MIPSKU
2. 需要填写完整的商品信息，包括标题、描述、图片等
3. SKU信息需要完整，包括价格、库存等
4. 发布前需要确保商品信息符合平台规范**

↓

**mipsku.get_account_list**

↓

**mipsku.get_category_tree**

↓

**mipsku.get_category_attribute**

↓

**mipsku.get_brand_list**

↓

**发布MIPSKU商品时，需要先获取MIPSKU账号列表，然后选择要发布的账号。接下来需要获取类目树、类目属性、品牌列表等信息，填写完整的商品信息后，就可以发布商品了。**

↓

### End

---

## [국경 간-KRSC] 글로벌 상품 흐름

### MTSKU 추가

#### Start

↓

**global_product.get_authorize_shop**

↓

**제품을 생성하기 전에 먼저 authorize shop을 가져와야 합니다. authorize shop은 카테고리 정보를 가져오고 제품을 게시하는 데 사용됩니다. authorize shop에 대한 자세한 내용은 "Authorize Shop" 섹션을 참조하십시오.**

↓

**global_product.get_global_product_category_tree**

↓

**판매자 센터 정보를 얻은 후 global_product.get_global_product_category_tree API를 사용하여 전체 카테고리 트리 목록을 가져올 수 있습니다.**

↓

**global_product.get_category_detail**

↓

**글로벌 상품 생성의 첫 번째 단계는 카테고리 속성 및 규칙을 포함하는 카테고리 세부 정보를 가져오는 것입니다.**

↓

**global_product.get_brand_list**

↓

**제품을 만드는 데 필요한 글로벌 제품 TWO를 가져옵니다. 카테고리 ID 및 키워드를 사용하여 API를 사용하여 브랜드 목록을 검색할 수 있습니다.**

↓

**global_product.get_global_product_category_attribute**

↓

**이 API를 사용하여 이 카테고리에 필요한 모든 속성을 가져옵니다. 속성은 다른 그룹으로 나뉩니다. 일부는 필수이고 일부는 선택 사항입니다.**

↓

**global_product.get_global_product_category_rule**

↓

**위의 모든 것을 준비한 후 글로벌 제품 생성을 시작할 수 있습니다. "항상 그런 것은 아닙니다"는 이 API 호출을 완료하기 위한 검증으로 사용됩니다. 그러나 이 API를 호출하기 전에 속성의 제약 조건을 확인하는 것이 좋습니다. 그렇지 않으면 검증으로 인해 API 호출이 실패할 수 있습니다.**

↓

**global_product_sku.global_product_sku_add**

↓

**global_product.add_global_product**

↓

**global_product.get_global_product_list**

↓

**global_product.get_global_product_detail**

↓

### End

---

### MTRSKU 모델 추가

#### Start

↓

**global_product.add_global_product**

↓

**MTRSKU가 성공적으로 생성된 후 API 문서에서 자세한 정보를 참조하십시오. 이는 MTRSKU가 생성되었지만 아직 게시되지 않았음을 의미합니다. MTRSKU의 상태는 "초안"이 됩니다. 실수를 발견하면 제품을 편집할 수 있습니다. 모든 세부 정보가 올바른지 확인한 후 MTRSKU를 게시하여 마켓플레이스에서 사용할 수 있도록 해야 합니다.**

↓

**global_product.edit_global_product**

↓

**1. MTRSKU 제품의 경우 판매자는 하나의 SPU에 여러 SKU를 만들 수 있습니다.
2. 각 SKU는 서로 다른 사양(예: 색상, 크기 등)을 가져야 합니다.
3. SKU 번호는 판매자 SKU라고도 하며 고유해야 합니다.
4. 각 상점은 하나의 제품에 대해 하나의 outer_product_id만 가질 수 있습니다. 즉, 한 상점에서 다른 제품에 대해 동일한 outer_product_id를 사용할 수 없습니다.
5. 편집 API의 shop_id는 제품을 만들 때의 shop_id와 일치해야 합니다.**

↓

**global_product_sku.global_product_sku_add**

↓

**global_product.edit_global_product**

↓

### End

---

## MTRSKU를 MIPSKU로 게시

### Start

↓

**global_product.add_global_product**

↓

**MTRSKU가 성공적으로 생성된 후 게시자 API를 사용하려면 다음 단계를 따르십시오.
1. 모든 정보를 포함하여 MIPSKU 세부 정보를 가져옵니다.
2. "mipPublisherCreate API"를 호출하고 MTRSKU에서 세부 정보를 전달합니다. 이 엔드포인트는 MTRSKU의 데이터를 표준 MIPSKU 데이터 형식으로 매핑합니다.
3. product_id를 편집해야 하는 경우 "getListApi"를 호출하여 product_id를 가져옵니다.
4. 응답에서 오류 코드 0 또는 성공을 받으면 MTRSKU가 MIPSKU에 성공적으로 게시되었음을 의미합니다.
5. 게시가 성공하면 MIPSKU 관리 포털을 통해 이 MIPSKU를 직접 관리하고 유지할 수 있습니다.**

↓

**mipsku.get_category_tree**

↓

**mipsku.get_category_attribute**

↓

**mipsku.mip_publisher_create**

↓

**global_product_publish.get_publish_task_result**

↓

### End

---

**참고: 데이터는 간단한 게시 서비스를 사용하는 것처럼 유지 관리됩니다. 게시된 목록 제품의 경우.**

## 사용 사례



## 관련 API



---

## 원문 (English)

### Summary

No summary available

### Content

# V2 SAR Call Flow

## Announcement Console

### Getting Started > V2 SAR Call Flow

---

## Main Flow Diagram

```
[Start]
    ↓
[www.www-vendor.xhtml]
    ↓
[Step1: select the vendor xhtml] ──┐
[Step2: select submit xhtml] ─────┤
[Step3: force to modify xhtml xhtml]─┤
[xx: Approval xxx- xxxxxxxxx -xxxx] ─┘
    ↓
[xxxxxxx xxx- xxxxxxx]
    ↓
[xxxxxxxx xxx- xxxxxxxxx]
    ↓
[xxxxxxxx xxx- xxxxxx xxx]
    ↓
[xxxxxxxx xxx- xxxx -xxx]
    ↓
[xxxxxxx xxx- xxxxxxx xx]
    ↓
[xxxxxxx xxxx- xxxx]
    ↓
[xxxxxxxx xxxx- xxx- xxxxxxx] ←→ [xxxxxxxxxxxxxxx xxxxx- xxxx]
    ↓
[End]
```

**Note: Dotted line is optional flow.**

---

## Order Status Flow

### Flow Description

[Left Side Notes:]
- Order status indicates the current state of fulfillment
- When order is placed, initial status will be "Pending"
- Status changes as order moves through fulfillment
- Multiple packages can have different statuses
- Final status indicates completion or cancellation

[Center Flow:]
```
[Start]
    ↓
[Pending]
    ↓
[Processing] ←→ [On Hold]
    ↓
[Approved]
    ↓
[Picking]
    ↓
[Packing]
    ↓
[Ready to Ship]
    ↓
[Shipped]
    ↓
[Delivered]
    ↓
[End]
```

[Right Side Notes:]
- On Hold status requires manual intervention
- Order can be cancelled at various stages
- Shipped status includes tracking information
- Delivered status confirms receipt
- Status updates trigger notifications

### Status Transition Table

| Current Status | Next Status | Condition |
|---------------|-------------|-----------|
| Pending | Processing | By Any Process |
| Pending | On Hold | By Any Process |
| On Hold | Processing | By Any Process |
| Processing | Approved | By Any Process |
| Processing | Cancelled/Rejected | By Any Request |
| Approved | Picking | Fulfillment Started |
| Picking | Packing | Items Collected |
| Packing | Ready to Ship | Package Prepared |
| Ready to Ship | Shipped | Carrier Pickup |
| Shipped | Delivered | Delivery Confirmed |

### LeadStatus

| Status | Description |
|--------|-------------|
| Unqualified | --- |
| New | --- |
| Working | By Any Processed |
| Nurturing | --- |

**Note: Dotted line is optional flow.**

---

## Package Fulfillment Status Flow

### Flow Description

[Left Side Notes:]
- Package represents individual shipment unit
- Each package has independent tracking
- Package status separate from order status
- Multiple packages per order possible
- Package weight and dimensions recorded
- Carrier assigned at package level

[Center Flow:]
```
[Start]
    ↓
[Created]
    ↓
[Processing] ←→ [On Hold]
    ↓
[Ready for Pickup]
    ↓
[Picked Up] ──→ [In Transit]
    ↓              ↓
[Out for Delivery]
    ↓              ↓
[Delivered] ←─────┘
    ↓
[End]
```

[Right Side Notes:]
- On Hold requires resolution before proceeding
- Tracking number assigned at pickup
- In Transit includes checkpoint updates
- Delivery attempts recorded
- Failed delivery creates exception
- Signature may be required for delivery

**Note: Dotted line is optional flow.**

---

## Arrange Shipment & Get Tracking№ & Print Away/Bill

### Flow Description

[Left Side Content:]
- Shipment arrangement initiates carrier booking
- System validates shipping address
- Carrier selection based on rules
- Shipping cost calculated
- Label generation process starts
- Tracking number obtained from carrier
- Airway bill/shipping label printed
- Documentation prepared for customs if needed

[Center Flow - Arrange Shipment:]
```
[Start]
    ↓
[Validate Address]
    ↓
[Select Carrier] ←→ [Calculate Cost]
    ↓
[Book Shipment]
    ↓
[Get Tracking Number]
    ↓
[Generate Label]
    ↓
[Print Airway Bill]
    ↓
[End]
```

[Center Flow - Get Tracking:]
```
[Start]
    ↓
[Query Carrier API]
    ↓
[Retrieve Tracking Data]
    ↓
[Update System]
    ↓
[Notify Customer]
    ↓
[End]
```

[Right Side Content:]
- Tracking number format varies by carrier
- Real-time tracking updates via API
- Customer notification includes tracking link
- Label must meet carrier specifications
- Barcode verification important
- Commercial invoice for international
- Packing slip included in package
- Proof of shipment retained

**Note: Dotted line is optional flow.**

---

## [CrossBorder] Seller Ship Out Tracking

### Flow Description

[Left Side Notes:]
- Cross-border shipments require customs clearance
- Export documentation must be accurate
- HS codes required for products
- Declared value affects duties
- Seller responsible for export compliance

[Center Flow:]
```
[Start]
    ↓
[Prepare Export Documents]
    ↓
[Submit Customs Declaration] ←→ [Customs Review]
    ↓
[Customs Clearance]
    ↓
[Ship to Destination Country]
    ↓
[Import Customs Process]
    ↓
[Deliver to Customer]
    ↓
[End]
```

[Right Side Notes:]
- Import duties may apply at destination
- Tracking includes customs checkpoints
- Delays possible during customs inspection
- Customer may need to pay duties
- Returns more complex for cross-border

**Note: Dotted line is optional flow.**

---

## [CrossBorder-CNSC]Global Product Flow

[This section appears to be a title/header for another flow diagram that continues below]

---

# [CrossBorder-CNSC] Global Product Flow

## 创建类目MTSKU

### Start

↓

**global_product.get_authorize_shop**

↓

**创建全球商品的第一步是准备授权店铺**
通过授权店铺，可以获得关联的类目信息，从而进行后续商品的发布。关于授权店铺的详细内容，请参考"授权店铺"章节。

↓

**global_product.get_category_list**

↓

**通过全球商品创建的第一步，选择类目开始创建商品流程。如果不确定类目信息，可以通过类目查询接口获取。通过类目信息，可以获取该类目下可用的品牌列表、属性列表等信息。选择好类目后，就可以进入商品信息填写阶段。**

↓

**global_product.get_category_detail**

↓

**global_product.get_brand_list**

↓

**global_product.get_global_product_category_attribute**

↓

**global_product.get_global_product_category_rule**

↓

**global_product_sku.global_product_sku_add**

↓

**global_product.add_global_product**

↓

**获取全球商品创建后，通过全球商品ID，调用全球商品查询接口，可以查询到创建的全球商品信息。**

↓

**global_product.get_global_product_list**

↓

**全球商品创建完成后，就可以查询到相关的全球商品信息。接下来需要将全球商品发布到店铺。**

↓

**global_product.get_global_product_detail**

↓

**全球商品创建完成后，就可以通过全球商品详情接口，查询全球商品的详细信息。**

↓

**global_product_publish.add_publish_task**

↓

**global_product_publish.get_publish_task_detail**

↓

**global_product_publish.get_publish_task_result**

↓

**下发发布任务后，系统会异步处理发布请求。可以通过查询发布任务结果接口，获取发布结果。发布成功后，就可以在店铺中看到商品了。**

↓

### End

---

## 添加MIPSKU产品流

### Start

↓

**global_product.add_global_product**

↓

**创建MIPSKU商品时，通过全球商品创建接口，可以直接创建MIPSKU商品。创建时需要指定商品类型为MIPSKU，并且填写相关的商品信息。创建完成后，系统会自动生成MIPSKU商品ID。**

↓

**global_product.get_global_product_detail**

↓

**global_product_sku.global_product_sku_add**

↓

**global_product_sku.global_product_sku_edit**

↓

### End

---

## 发布MIPSKU产品

### Start

↓

**global_product.add_global_product**

↓

**创建MIPSKU商品时需要注意以下几点：
1. 商品类型需要设置为MIPSKU
2. 需要填写完整的商品信息，包括标题、描述、图片等
3. SKU信息需要完整，包括价格、库存等
4. 发布前需要确保商品信息符合平台规范**

↓

**mipsku.get_account_list**

↓

**mipsku.get_category_tree**

↓

**mipsku.get_category_attribute**

↓

**mipsku.get_brand_list**

↓

**发布MIPSKU商品时，需要先获取MIPSKU账号列表，然后选择要发布的账号。接下来需要获取类目树、类目属性、品牌列表等信息，填写完整的商品信息后，就可以发布商品了。**

↓

### End

---

## [CrossBorder-KRSC] Global Product Flow

### Add MTSKU

#### Start

↓

**global_product.get_authorize_shop**

↓

**Before creating a product, you will need to get the authorize shop first. The authorize shop is used to get the category information and publish the product. For more details about authorize shop, please refer to the "Authorize Shop" section.**

↓

**global_product.get_global_product_category_tree**

↓

**After getting the seller center info, you can use the global_product.get_global_product_category_tree API to get the complete category tree list.**

↓

**global_product.get_category_detail**

↓

**First step of creating a global product is to get the category detail which includes the category attributes and rules.**

↓

**global_product.get_brand_list**

↓

**Get the TWO of the global product which is required for creating a product. You can use the API with category ID and keyword to search the brand list.**

↓

**global_product.get_global_product_category_attribute**

↓

**Use this API to get all the attributes required for this category. The attributes are divided into different groups. Some are required and some are optional.**

↓

**global_product.get_global_product_category_rule**

↓

**After preparing all above, you can start to create a global product. "Not always" is used as the verification for completing this API call. However, it is good practice to check the constraints of attributes before calling this API, otherwise the API call may fail due to verification.**

↓

**global_product_sku.global_product_sku_add**

↓

**global_product.add_global_product**

↓

**global_product.get_global_product_list**

↓

**global_product.get_global_product_detail**

↓

### End

---

### Add MTRSKU model

#### Start

↓

**global_product.add_global_product**

↓

**After the MTRSKU is created successfully, refer more info in the API documentation. This means the MTRSKU is created but not published yet. The status of the MTRSKU will be "Draft". You can edit the product if you find any mistake. After you ensure all details are correct, you need to publish the MTRSKU to make it available in the marketplace.**

↓

**global_product.edit_global_product**

↓

**1. For MTRSKU product, seller can create multiple SKUs under one SPU.
2. Each SKU should have different spec (like color, size, etc)
3. SKU number is also called seller SKU which should be unique
4. Each shop can only have one outer_product_id for one product, which means you can't use same outer_product_id for different products in one shop
5. The shop_id in edit API should be consistent with the shop_id when you are creating the product**

↓

**global_product_sku.global_product_sku_add**

↓

**global_product.edit_global_product**

↓

### End

---

## Publish MTRSKU to be MIPSKU

### Start

↓

**global_product.add_global_product**

↓

**After the MTRSKU is created successfully, to use the publisher API, you need to follow these steps:
1. Get the MIPSKU detail including all information
2. Call the "mipPublisherCreate API" and pass the details from MTRSKU. This endpoint maps data from MTRSKU into a standard MIPSKU data format
3. If you need to edit product_id, get the product_id by calling "getListApi"
4. If you get an error code 0 or success in the response, that means your MTRSKU has been successfully published to MIPSKU
5. After successful publishing, you can manage and maintain this MIPSKU directly through MIPSKU management portal**

↓

**mipsku.get_category_tree**

↓

**mipsku.get_category_attribute**

↓

**mipsku.mip_publisher_create**

↓

**global_product_publish.get_publish_task_result**

↓

### End

---

**Note: The data will be maintained like you use simple publish services. For the published listing product.**

---

**문서 ID**: developer-guide.27
**플랫폼**: shopee
**URL**: https://open.shopee.com/developer-guide/27
**처리 완료**: 2025-10-16T08:23:26
