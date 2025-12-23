# Shipper Service - 패키지

**카테고리**: 모범 사례
**난이도**: medium
**중요도**: 4/5
**최종 업데이트**: 2025-10-16T09:00:44

## 개요

본 가이드는 Shipper Service 패키지 기능 사용에 대한 모범 사례를 제공합니다. 그룹 정렬, 비즈니스 운영 로직, 포장되지 않은 SKU ID 조회, 주문 정보 획득, 기본 배송 로직 및 물류 채널 내 패키지 처리와 관련된 기타 프로세스와 같은 주제를 다룹니다.

## 주요 키워드

- Shipper Service
- Package
- 물류 채널
- Unpackaged SKU ID
- 배송
- 정렬 그룹
- Drop Shipment
- API
- 주문 관리

## 본문

# API 모범 사례 > 배송사 서비스 - 패키지

## 1 정렬 그룹
정렬 그룹: 이 물류 채널로 중단할 때 팩 목록 흐름을 함께 선택하는 데 사용되는 이름입니다.

## 2 비즈니스 운영 로직

### 2.1 자동 드롭 배송
물류 채널 50029를 사용하는 주문의 경우, 주문 상태가 "READY_TO_SHIP"이 되면 시스템은 드롭 투 클라이언트 배송이 존재하는지 확인하기 위해 원산지 정보(주문) 패키지 드롭 정보를 자동으로 가져옵니다.
- 드롭 투 클라이언트 배송이 존재하는 경우, 패키지 드롭이 실행되고 1개의 패키지에 3개의 드롭 투 클라이언트 유형의 상태가 업데이트되며, 각 유형은 1개의 수량을 포함합니다.

### 2.2 평균 배송 드롭
물류 채널 50029를 사용하여 품목을 드롭하는 경우, 판매자는 더 이상 주문 라인별로 포장할 필요가 없습니다. 단계는 다음과 같습니다.
- 패키지 SKU ID Label은 여전히 경로 및 재고 SKU입니다. 배송 소포에 SKU Label을 사용하지 마십시오.
- 중단 시, 배치 재고 SKU에 따라 포장하십시오.

자세한 내용은 판매자 교육 허브 문서 [提报发货包, 智能发货包], [三方发货报包]를 참조하십시오.

## 3 품목의 미포장 SKU ID 조회
"미포장 SKU ID-U Label"은 물류 채널 내의 SKU 품목을 기록합니다. 다음 API를 통해 정보를 얻을 수 있습니다.
- V3 [logistics/get_shipping_parameter](/open-api/01_logistics_get_shipping_parameter)
- V2 [product:search_unpackage_master_sku](/open-api/02_product_search_unpackage_master_sku)

## 4 주문 정보 및 세부 정보 가져오기
이 물류 채널은 자동으로 주문을 가져오지 않습니다. 다음 API를 사용하는 동안 패키지 수준 정보를 가져와야 합니다.
- V3 [order:search_package_list](/open-api/03_order_search_package_list) 패키지 목록을 조회합니다.
- V3 [order:get_package_detail](/open-api/04_order_get_package_detail) 패키지 세부 정보를 조회합니다.

주문 정보를 얻은 후, "shipping_carrier" 필드가 값 "package_id"와 일치하는지 확인하고 요청 매개변수 "response_optional_fields"에 포함되어 있는지 확인하십시오.

## 5 기본 배송 로직

### 5.1 필수 API 호출 흐름
이 물류 채널은 자동으로 주문을 분할하지 않습니다. 다음 API의 요청 매개변수에 "package_number"가 포함되어 있는지 확인하십시오.
- V3 [order:ship_package](/open-api/05_order_ship_package) 주문 배송 시 ("package_list"에 배송 매개변수 "tracking_number"도 필요합니다.)
또한:
- V2 [logistics:init_and_order](/open-api/06_logistics_init_and_order) 배송을 준비할 때 "sorting_group" 필드에서 값을 얻을 수 있습니다.

### 5.2 평균 배송 드롭
단계:
1) 미포장 SKU ID 데드 콜 흐름 가져오기
미포장 SKU ID는 패키지에 해당하는 미포장 SKU ID의 고유 식별자입니다.
- V3 [logistics:get_shipping_document_parameter](/open-api/07_logistics_get_shipping_document_parameter)
- V2 [logistics:get_airway_bill_for_seller](/open-api/08_logistics_get_airway_bill_for_seller)
- V2 [logistics:get_tracking_number](/open-api/09_logistics_get_tracking_number)

물류 채널 ID 필드 "response_OPTIONAL_UNPACKAGED_LABEL"의 값은 "shipping_document_type"입니다.

2) To Label ID 가져오기
정렬 그룹에 따라 패키지를 완료한 후, 다음 API를 통해 TO 레이블 파일을 얻을 수 있습니다.
- V3 [logistics:download_shipping_document](/open-api/10_logistics_download_shipping_document) TO 레이블 다운로드 (드롭 배송 정보를 포함하며, 드롭 오프 프로세스에 영향을 미치는 필수 레이블입니다.)

## 6 기타 프로세스

### 6.1 주문 생성 전 U-Label 생성 또는 업데이트
1. 주문 생성 전에 미포장 SKU ID-Label을 인쇄하는 경우
   판매자가 주문을 생성할 수 있도록 지원 및 업로드하여 고급 라벨링 및 효율성을 보장합니다.

다음 API를 호출하여 파일을 얻을 수 있습니다.
- V3 [logistics:batch_ship_package](/open-api/11_logistics_batch_ship_package) "인쇄 작업 생성" 호출
- V3 [logistics:download_shipping_document_job](/open-api/12_logistics_download_shipping_document_job) 레이블 파일 다운로드

데이터는 소포 채널 50029부터 시작해야 이 고급 레이블 인쇄 프로세스를 효율적으로 사용할 수 있습니다.

### 6.2 맞춤형 U-Label
물류 시스템에서는 시스템 기본 설정에 따라 "미포장 SKU ID Label"을 디자인하는 것이 좋습니다. FAQ 첨부 파일의 "판매자 도움말" 섹션 사양을 참조하여 레이블을 만들 수 있습니다(반품 배송사에서 배송).
- API를 통해 "Unpackage_sku_id" 및 "Unpackaged_sku_id_group" 필드가 있는지 확인하십시오.
- 참조용 지침이며, 공식 사양은 FAQ 첨부 파일에 있습니다(영문 버전만 해당).
  * 10cm x 10cm를 초과할 수 없습니다.
  * EC 모드는 허용되지 않습니다.
  * QR 코드를 가리는 QR 코드 또는 기타 장애물을 피하십시오.
  * 글꼴 크기 또는 텍스트는 최소한 읽을 수 있어야 합니다.
  * 페인트 도구 및 막대 끊기 지침이 제거되었습니다.
  * 사용자 정의 레이블이 사용자 정의되지 않은 영역의 표시 또는 레이아웃에 영향을 미치지 않도록 하십시오.

## 7 권장 API 호출 흐름

### 7.1 주문 생성 전 미포장 SKU ID-Label 인쇄

[이것은 순서도입니다: 주문 생성 전 미포장 SKU ID Label 인쇄를 위한 프로세스 흐름을 보여주며, 여러 의사 결정 지점과 노란색 상자의 프로세스 단계를 포함합니다.]

### 7.2 주문 생성 전 미포장 SKU ID-Label 인쇄하는 경우

[이것은 순서도입니다: 주문 생성 및 레이블 인쇄 프로세스를 포함하여 여러 단계와 의사 결정 지점이 있는 대체 프로세스 흐름을 보여줍니다.]

## 8 FAQ

**Q: 물류 채널에 여러 패키지가 있는 주문의 경우, 패키지의 물류 상태가 다른 경우 주문 상태는 어떻게 표시됩니까?**

A: [주문 시스템 장(/open-faq/01)]에 대한 FAQ를 참조하십시오.

**Q: V2 (order:get_order_detail) API를 통해 주문 세부 정보를 검색하는 동안 주문 수준의 "shipping_carrier" 필드가 비어 있는 이유는 무엇입니까?**

A: 이 물류 운송업체의 물류 채널에 있지 않으면 API는 값을 표시하지 않습니다. "package_list"의 각 배송에서 "shipping_carrier" 필드를 검색하십시오.

**Q: 배송 (또는 미포장 SKU ID-Label) API를 호출할 때 시스템에서 오류 메시지를 반환하는 이유는 무엇입니까?**

A: 물류 채널이 자동으로 주문을 분할합니다. 물류 관련 API를 호출할 때 "package_number" 필드 매개변수가 API에 포함되도록 요청하십시오.

## 9 데이터 정의

[V3 주문 시스템(https://example.com)] 페이지를 참조하십시오.

## 사용 사례

1. 특정 물류 채널 내에서 패키지 관리.
2. 자동 Drop Shipment 프로세스 구현.
3. 아이템에 대한 포장되지 않은 SKU ID 조회.
4. 주문 정보 및 패키지 세부 정보 획득.
5. 필수 API 호출 흐름을 사용하여 패키지 배송.

## 관련 API

- /open-api/01_logistics_get_shipping_parameter
- /open-api/02_product_search_unpackage_master_sku
- /open-api/03_order_search_package_list
- /open-api/04_order_get_package_detail
- /open-api/05_order_ship_package
- /open-api/06_logistics_init_and_order
- /open-api/07_logistics_get_shipping_document_parameter
- /open-api/08_logistics_get_airway_bill_for_seller
- /open-api/09_logistics_get_tracking_number
- /open-api/10_logistics_download_shipping_document
- /open-api/11_logistics_batch_ship_package
- /open-api/12_logistics_download_shipping_document_job

---

## 원문 (English)

### Summary

This guide provides best practices for using the Shipper Service package functionality. It covers topics such as sorting groups, business operation logic, querying unpackaged SKU IDs, obtaining order information, basic shipping logic, and other processes related to package handling within the logistics channel.

### Content

# API Best Practices > Shipper Service - Package

## 1 Sorting group
Sorting group: The name that is used to select the pack list flow together when stopping with this logistics channel

## 2 Business Operation Logic

### 2.1 Automatic Drop Shipment
For orders with logistics channel 50029, when the order status becomes "READY_TO_SHIP", the system will automatically get the origin info (the order) package drop info to qualify it if drop-to-client shipment exists.
- If drop-to-client shipment exists, package drop will be executed and update the status of that 3 drop-to-client type in 1 package, each containing 1 quantity

### 2.2 Average Shipment Dropping
When using logistics channel 50029 to drop items, sellers no longer need to pack by order lines. The steps are as follows:
- The package SKU ID Label is still a route and stock SKU. Do not use the SKU Label on the shipping parcel.
- When stopping, pack according to the batch stock SKU

Please refer to the Seller Education Hub article for more details: [提报发货包, 智能发货包], [三方发货报包]

## 3 Query item's Unpackaged SKU ID
The "Unpackage SKU ID-U Label" records the SKU item within the logistics channel. You can obtain the information through the API:
- V3 [logistics/get_shipping_parameter](/open-api/01_logistics_get_shipping_parameter)
- V2 [product:search_unpackage_master_sku](/open-api/02_product_search_unpackage_master_sku)

## 4 Obtain order information and detail
This logistics channel does not automatically take orders. Be sure to get package level information while using the following API:
- V3 [order:search_package_list](/open-api/03_order_search_package_list) To query the package list
- V3 [order:get_package_detail](/open-api/04_order_get_package_detail) To query the package detail

After obtaining the order information, confirm that the "shipping_carrier" field matches value "package_id" is included in the request parameter, "response_optional_fields"

## 5 Basic Shipping Logic

### 5.1 Required API Call Flow
This logistics channel does not automatically split orders, please ensure "package_number" is included in the request parameter of the following API:
- V3 [order:ship_package](/open-api/05_order_ship_package) When shipping orders (also shipping parameter "tracking_number" is required in the "package_list")
Also:
- V2 [logistics:init_and_order](/open-api/06_logistics_init_and_order) When arrange shipments, you can obtain the value in the "sorting_group" field

### 5.2 Average Shipment Dropping
Steps:
1) Obtain Unpackage sku id dead call flow
The Unpackage SKU ID is a unique identifier for the Unpackage sku id used corresponding to the package.
- V3 [logistics:get_shipping_document_parameter](/open-api/07_logistics_get_shipping_document_parameter)
- V2 [logistics:get_airway_bill_for_seller](/open-api/08_logistics_get_airway_bill_for_seller)
- V2 [logistics:get_tracking_number](/open-api/09_logistics_get_tracking_number)

The value for the logistics channel ID field "response_OPTIONAL_UNPACKAGED_LABEL" is the "shipping_document_type"

2) Obtain To Label ID
After completing the package according to the sorting group, you can obtain the TO label file through the following API:
- V3 [logistics:download_shipping_document](/open-api/10_logistics_download_shipping_document) Download TO label (including drop-ship info, which are mandatory labels that will affect the drop-off process)

## 6 Other Process

### 6.1 Create or Update U-Label before order creation
1. If the printing Unpackaged SKU ID-Label before order creation
   Support and uploading to make sure for the seller can create the order created, for advanced labeling and efficiency.

You can obtain the file by calling the following APIs as supposed:
- V3 [logistics:batch_ship_package](/open-api/11_logistics_batch_ship_package) Call "Create a print job"
- V3 [logistics:download_shipping_document_job](/open-api/12_logistics_download_shipping_document_job) Download the label files

Please note that the data must begin with parcel channel 50029 enables to use this advance label print process for efficiency.

### 6.2 Custom-design U-Label
In the logistics system, we recommend that you design the "Unpackaged SKU ID Label according to system preferences. You can refer to the "Seller Help" section specification in the FAQ attachment to create labels (delivery by return shipper).
- Make sure the "Unpackage_sku_id" and "Unpackaged_sku_id_group" fields through the API
- Guidelines for reference only, the official specifications are in the FAQ attachment (only in English version):
  * Cannot exceed 10cm x 10cm
  * EC mode is not allowed
  * Avoid presenting the QR code or other obscuring obstructions to the QR code
  * Font size or text is minimal to readable
  * Paint tool and bar-break instructions have been removed
  * Ensure that custom label has no effect the display or layout of non-custom areas.

## 7 Recommended API Call Flow

### 7.1 Printing Unpackaged SKU ID-Label before order creation

[THIS IS FLOWCHART: Shows the process flow for printing Unpackaged SKU ID Label before order creation, with multiple decision points and process steps in yellow boxes]

### 7.2 If the printing Unpackaged SKU ID-Label before order creation

[THIS IS FLOWCHART: Shows an alternative process flow with multiple steps and decision points, including order creation and label printing processes]

## 8 FAQ

**Q: For orders under the logistics channel with multiple packages, where the logistics statuses of the packages differ, how will the Order status be displayed?**

A: Please to see FAQ for the [Order system chapter(/open-faq/01)]

**Q: Why is the "shipping_carrier" field of the order-level empty while I retrieve order details through the V2 (order:get_order_detail) API?**

A: Unless under the logistics channel of this logistics carrier, the API will not display a value. Please retrieve the "shipping_carrier" field from each shipment in the "package_list."

**Q: Why does the system return an error message when I call the Shipment (or Unpackaged SKU ID-Label) API?**

A: The logistics channel automatically split orders. Please request that the "package_number" field parameter is included in the API when calling logistics-related APIs.

## 9 Data Definition

Please refer to the [V3 Order System(https://example.com)] page.

---

**문서 ID**: developer-guide.677
**플랫폼**: shopee
**URL**: https://open.shopee.com/developer-guide/677
**처리 완료**: 2025-10-16T09:00:44
