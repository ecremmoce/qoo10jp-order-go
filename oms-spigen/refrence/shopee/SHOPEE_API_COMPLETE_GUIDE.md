# 🛍️ Shopee Open API 완전 분석 가이드

**생성일시**: 2025년 10월 16일 16:04:45
**분석 버전**: v2.0

> Shopee Open Platform의 362개 API를 카테고리별로 완전 분석한 문서입니다.

---

## 📊 전체 통계

- 🔢 **총 API 개수**: 362개
- 📁 **카테고리 개수**: 27개
- 📝 **마크다운 문서**: 2개 (0%)
- 📋 **메타데이터**: 360개 (99%)
- 🌐 **한글 번역**: 4개 (1%)

### 📦 카테고리별 API 분포

| 순위 | 카테고리 | API 개수 | 비율 |
|:----:|---------|:--------:|:----:|
| 1 | **product** | 55 | 15.2% |
| 2 | **logistics** | 41 | 11.3% |
| 3 | **global_product** | 34 | 9.4% |
| 4 | **ads** | 26 | 7.2% |
| 5 | **livestream** | 25 | 6.9% |
| 6 | **order** | 21 | 5.8% |
| 7 | **first_mile** | 16 | 4.4% |
| 8 | **payment** | 16 | 4.4% |
| 9 | **returns** | 15 | 4.1% |
| 10 | **add_on_deal** | 14 | 3.9% |
| 11 | **shop_flash_sale** | 11 | 3.0% |
| 12 | **bundle_deal** | 10 | 2.8% |
| 13 | **discount** | 9 | 2.5% |
| 14 | **shop_category** | 7 | 1.9% |
| 15 | **account_health** | 6 | 1.7% |
| 16 | **follow_prize** | 6 | 1.7% |
| 17 | **media_space** | 6 | 1.7% |
| 18 | **merchant** | 6 | 1.7% |
| 19 | **public** | 6 | 1.7% |
| 20 | **shop** | 6 | 1.7% |
| 21 | **voucher** | 6 | 1.7% |
| 22 | **sbs** | 5 | 1.4% |
| 23 | **fbs** | 4 | 1.1% |
| 24 | **push** | 4 | 1.1% |
| 25 | **top_picks** | 4 | 1.1% |
| 26 | **v2** | 2 | 0.6% |
| 27 | **media** | 1 | 0.3% |

---

## 📑 목차

1. 🛒 [PRODUCT](#product) - 55개 API
   > 상품 관리 - 상품 등록, 수정, 조회, 재고 관리 등 상품 관련 핵심 기능
2. 🚚 [LOGISTICS](#logistics) - 41개 API
   > 물류 및 배송 관리 - 배송 방법, 운송장, 배송 문서 등 물류 전반
3. 📌 [GLOBAL_PRODUCT](#global_product) - 34개 API
   > 글로벌 상품 관리 - 여러 국가 샵에 동시 등록 가능한 글로벌 상품
4. 📢 [ADS](#ads) - 26개 API
   > 광고 캠페인 관리 - CPC 광고, 상품 광고, GMS 캠페인 등 광고 관련 모든 기능
5. 📹 [LIVESTREAM](#livestream) - 25개 API
   > 라이브 스트리밍 관리 - 라이브 커머스 세션, 상품, 댓글 등 관리
6. 📦 [ORDER](#order) - 21개 API
   > 주문 관리 - 주문 조회, 취소, 송장 생성 등 주문 관련 모든 기능
7. 📌 [FIRST_MILE](#first_mile) - 16개 API
   > 퍼스트 마일 배송 관리 - 판매자에서 물류센터까지의 첫 배송 단계
8. 💰 [PAYMENT](#payment) - 16개 API
   > 결제 및 정산 관리 - 에스크로, 수익 보고서, 지갑 거래 등
9. ↩️ [RETURNS](#returns) - 15개 API
   > 반품 및 환불 관리 - 반품 요청, 분쟁, 해결 방안 등
10. 📌 [ADD_ON_DEAL](#add_on_deal) - 14개 API
   > 추가 상품 딜 관리 - 메인 상품과 함께 구매할 수 있는 추가 상품 프로모션
11. 📌 [SHOP_FLASH_SALE](#shop_flash_sale) - 11개 API
   > 샵 플래시 세일 - 샵 자체 플래시 세일 이벤트
12. 📌 [BUNDLE_DEAL](#bundle_deal) - 10개 API
   > 번들 상품 딜 관리 - 여러 상품을 묶어서 판매하는 프로모션
13. 🏷️ [DISCOUNT](#discount) - 9개 API
   > 할인 프로모션 관리 - 상품별 할인 설정 및 관리
14. 📌 [SHOP_CATEGORY](#shop_category) - 7개 API
   > 샵 카테고리 관리 - 샵 내부 카테고리 및 상품 분류
15. 💊 [ACCOUNT_HEALTH](#account_health) - 6개 API
   > 계정 건강도 및 성과 모니터링 - 판매자의 계정 상태, 지연 주문, 문제 있는 리스팅 등을 추적
16. 📌 [FOLLOW_PRIZE](#follow_prize) - 6개 API
   > 팔로우 경품 이벤트 - 샵 팔로우 고객 대상 경품 프로모션
17. 📌 [MEDIA_SPACE](#media_space) - 6개 API
   > 미디어 스페이스 관리 - 비디오 업로드 및 미디어 관리
18. 📌 [MERCHANT](#merchant) - 6개 API
   > 판매자 정보 관리 - 판매자(파트너) 계정 및 창고 정보
19. 📌 [PUBLIC](#public) - 6개 API
   > 공개 API - 인증, 토큰 발급, IP 범위 등 공개적으로 접근 가능한 API
20. 🏪 [SHOP](#shop) - 6개 API
   > 샵 정보 관리 - 샵 프로필, 창고 정보 등
21. 🎟️ [VOUCHER](#voucher) - 6개 API
   > 쿠폰/바우처 관리 - 할인 쿠폰 생성 및 관리
22. 📌 [SBS](#sbs) - 5개 API
   > SBS (Shopee Business Solutions) - 재고 추적, 만료 리포트 등
23. 📌 [FBS](#fbs) - 4개 API
   > FBS (Fulfilled By Shopee) - Shopee가 직접 물류를 담당하는 서비스
24. 📌 [PUSH](#push) - 4개 API
   > 푸시 알림 설정 - 웹훅 및 푸시 알림 설정
25. 📌 [TOP_PICKS](#top_picks) - 4개 API
   > 추천 상품 관리 - 샵의 추천 상품 설정
26. 📌 [V2](#v2) - 2개 API
27. 📌 [MEDIA](#media) - 1개 API
   > 미디어 업로드 - 이미지 업로드 기본 기능

---

## 🛒 PRODUCT

**API 개수**: 55개

**설명**: 상품 관리 - 상품 등록, 수정, 조회, 재고 관리 등 상품 관련 핵심 기능

| API 메서드 | HTTP | 설명 | 문서 |
|-----------|:----:|-----|:----:|
| `add_item` | POST | Documentation - Shopee Open Platform | ❌ |
| `add_kit_item` | POST | Documentation - Shopee Open Platform | ❌ |
| `add_model` | POST | Documentation - Shopee Open Platform | ❌ |
| `add_ssp_item` | POST | Documentation - Shopee Open Platform | ❌ |
| `boost_item` | POST | Documentation - Shopee Open Platform | ❌ |
| `category_recommend` | POST | Documentation - Shopee Open Platform | ❌ |
| `delete_item` | DELETE | Documentation - Shopee Open Platform | ❌ |
| `delete_model` | DELETE | Documentation - Shopee Open Platform | ❌ |
| `generate_kit_image` | POST | Documentation - Shopee Open Platform | ❌ |
| `get_aitem_by_pitem_id` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_all_vehicle_list` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_attribute_tree` | GET | Documentation - Shopee Open Platform | 📝 🌐 |
| `get_boosted_list` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_brand_list` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_category` | GET | Documentation - Shopee Open Platform | 📝 🌐 |
| `get_comment` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_direct_item_list` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_direct_shop_recommended_price` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_item_base_info` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_item_content_diagnosis_result` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_item_extra_info` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_item_limit` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_item_list` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_item_list_by_content_diagnosis` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_item_promotion` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_item_violation_info` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_kit_item_info` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_kit_item_limit` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_main_item_list` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_model_list` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_product_certification_rule` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_recommend_attribute` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_size_chart_detail` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_size_chart_list` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_ssp_info` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_ssp_list` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_variations` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_vehicle_list_by_compatibility_detail` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_weight_recommendation` | GET | Documentation - Shopee Open Platform | ❌ |
| `init_tier_variation` | POST | Documentation - Shopee Open Platform | ❌ |
| `link_ssp` | POST | Documentation - Shopee Open Platform | ❌ |
| `register_brand` | POST | Documentation - Shopee Open Platform | ❌ |
| `reply_comment` | POST | Documentation - Shopee Open Platform | ❌ |
| `search_attribute_value_list` | POST | Documentation - Shopee Open Platform | ❌ |
| `search_item` | POST | Documentation - Shopee Open Platform | ❌ |
| `search_unpackaged_model_list` | POST | Documentation - Shopee Open Platform | ❌ |
| `unlink_ssp` | POST | Documentation - Shopee Open Platform | ❌ |
| `unlist_item` | POST | Documentation - Shopee Open Platform | ❌ |
| `update_item` | PUT | Documentation - Shopee Open Platform | ❌ |
| `update_kit_item` | PUT | Documentation - Shopee Open Platform | ❌ |
| `update_model` | PUT | Documentation - Shopee Open Platform | ❌ |
| `update_price` | PUT | Documentation - Shopee Open Platform | ❌ |
| `update_sip_item_price` | PUT | Documentation - Shopee Open Platform | ❌ |
| `update_stock` | PUT | Documentation - Shopee Open Platform | ❌ |
| `update_tier_variation` | PUT | Documentation - Shopee Open Platform | ❌ |

### 주요 API 상세 정보

#### `add_item`

**전체 이름**: `v2.product.add_item`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.product.add_item?module=89&type=1](https://open.shopee.com/documents/v2/v2.product.add_item?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `add_kit_item`

**전체 이름**: `v2.product.add_kit_item`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.product.add_kit_item?module=89&type=1](https://open.shopee.com/documents/v2/v2.product.add_kit_item?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `add_model`

**전체 이름**: `v2.product.add_model`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.product.add_model?module=89&type=1](https://open.shopee.com/documents/v2/v2.product.add_model?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `add_ssp_item`

**전체 이름**: `v2.product.add_ssp_item`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.product.add_ssp_item?module=89&type=1](https://open.shopee.com/documents/v2/v2.product.add_ssp_item?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `boost_item`

**전체 이름**: `v2.product.boost_item`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.product.boost_item?module=89&type=1](https://open.shopee.com/documents/v2/v2.product.boost_item?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


---

## 🚚 LOGISTICS

**API 개수**: 41개

**설명**: 물류 및 배송 관리 - 배송 방법, 운송장, 배송 문서 등 물류 전반

| API 메서드 | HTTP | 설명 | 문서 |
|-----------|:----:|-----|:----:|
| `batch_ship_order` | POST | Documentation - Shopee Open Platform | ❌ |
| `batch_update_tpf_warehouse_tracking_status` | POST | Documentation - Shopee Open Platform | ❌ |
| `create_booking_shipping_document` | POST | Documentation - Shopee Open Platform | ❌ |
| `create_shipping_document` | POST | Documentation - Shopee Open Platform | ❌ |
| `create_shipping_document_job` | POST | Documentation - Shopee Open Platform | ❌ |
| `delete_address` | DELETE | Documentation - Shopee Open Platform | ❌ |
| `delete_special_operating_hour` | DELETE | Documentation - Shopee Open Platform | ❌ |
| `download_booking_shipping_document` | POST | Documentation - Shopee Open Platform | ❌ |
| `download_shipping_document` | POST | Documentation - Shopee Open Platform | ❌ |
| `download_shipping_document_job` | POST | Documentation - Shopee Open Platform | ❌ |
| `download_to_label` | POST | Documentation - Shopee Open Platform | ❌ |
| `get_address_list` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_booking_shipping_document_data_info` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_booking_shipping_document_parameter` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_booking_shipping_document_result` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_booking_shipping_parameter` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_booking_tracking_info` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_booking_tracking_number` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_channel_list` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_mart_packaging_info` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_mass_shipping_parameter` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_mass_tracking_number` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_operating_hour_restrictions` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_operating_hours` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_shipping_document_data_info` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_shipping_document_job_status` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_shipping_document_parameter` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_shipping_document_result` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_shipping_parameter` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_tracking_info` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_tracking_number` | GET | Documentation - Shopee Open Platform | ❌ |
| `mass_ship_order` | POST | Documentation - Shopee Open Platform | ❌ |
| `set_address_config` | POST | Documentation - Shopee Open Platform | ❌ |
| `set_mart_packaging_info` | POST | Documentation - Shopee Open Platform | ❌ |
| `ship_booking` | POST | Documentation - Shopee Open Platform | ❌ |
| `ship_order` | POST | Documentation - Shopee Open Platform | ❌ |
| `update_channel` | PUT | Documentation - Shopee Open Platform | ❌ |
| `update_operating_hours` | PUT | Documentation - Shopee Open Platform | ❌ |
| `update_self_collection_order_logistics` | PUT | Documentation - Shopee Open Platform | ❌ |
| `update_shipping_order` | PUT | Documentation - Shopee Open Platform | ❌ |
| `update_tracking_status` | PUT | Documentation - Shopee Open Platform | ❌ |

### 주요 API 상세 정보

#### `batch_ship_order`

**전체 이름**: `v2.logistics.batch_ship_order`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.logistics.batch_ship_order?module=89&type=1](https://open.shopee.com/documents/v2/v2.logistics.batch_ship_order?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `batch_update_tpf_warehouse_tracking_status`

**전체 이름**: `v2.logistics.batch_update_tpf_warehouse_tracking_status`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.logistics.batch_update_tpf_warehouse_tracking_status?module=89&type=1](https://open.shopee.com/documents/v2/v2.logistics.batch_update_tpf_warehouse_tracking_status?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `create_booking_shipping_document`

**전체 이름**: `v2.logistics.create_booking_shipping_document`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.logistics.create_booking_shipping_document?module=89&type=1](https://open.shopee.com/documents/v2/v2.logistics.create_booking_shipping_document?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `create_shipping_document`

**전체 이름**: `v2.logistics.create_shipping_document`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.logistics.create_shipping_document?module=89&type=1](https://open.shopee.com/documents/v2/v2.logistics.create_shipping_document?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `create_shipping_document_job`

**전체 이름**: `v2.logistics.create_shipping_document_job`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.logistics.create_shipping_document_job?module=89&type=1](https://open.shopee.com/documents/v2/v2.logistics.create_shipping_document_job?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


---

## 🌍 GLOBAL_PRODUCT

**API 개수**: 34개

**설명**: 글로벌 상품 관리 - 여러 국가 샵에 동시 등록 가능한 글로벌 상품

| API 메서드 | HTTP | 설명 | 문서 |
|-----------|:----:|-----|:----:|
| `add_global_item` | POST | Documentation - Shopee Open Platform | ❌ |
| `add_global_model` | POST | Documentation - Shopee Open Platform | ❌ |
| `category_recommend` | POST | Documentation - Shopee Open Platform | ❌ |
| `create_publish_task` | POST | Documentation - Shopee Open Platform | ❌ |
| `delete_global_item` | DELETE | Documentation - Shopee Open Platform | ❌ |
| `delete_global_model` | DELETE | Documentation - Shopee Open Platform | ❌ |
| `get_attribute_tree` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_brand_list` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_category` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_global_item_id` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_global_item_info` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_global_item_limit` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_global_item_list` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_global_model_list` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_local_adjustment_rate` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_publish_task_result` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_publishable_shop` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_published_list` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_recommend_attribute` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_shop_publishable_status` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_size_chart_detail` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_size_chart_list` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_variations` | GET | Documentation - Shopee Open Platform | ❌ |
| `init_tier_variation` | POST | Documentation - Shopee Open Platform | ❌ |
| `search_global_attribute_value_list` | POST | Documentation - Shopee Open Platform | ❌ |
| `set_sync_field` | POST | Documentation - Shopee Open Platform | ❌ |
| `support_size_chart` | POST | Documentation - Shopee Open Platform | ❌ |
| `update_global_item` | PUT | Documentation - Shopee Open Platform | ❌ |
| `update_global_model` | PUT | Documentation - Shopee Open Platform | ❌ |
| `update_local_adjustment_rate` | PUT | Documentation - Shopee Open Platform | ❌ |
| `update_price` | PUT | Documentation - Shopee Open Platform | ❌ |
| `update_size_chart` | PUT | Shopee Open Platform | ❌ |
| `update_stock` | PUT | Documentation - Shopee Open Platform | ❌ |
| `update_tier_variation` | PUT | Shopee Open Platform | ❌ |

### 주요 API 상세 정보

#### `add_global_item`

**전체 이름**: `v2.global_product.add_global_item`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.global_product.add_global_item?module=89&type=1](https://open.shopee.com/documents/v2/v2.global_product.add_global_item?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `add_global_model`

**전체 이름**: `v2.global_product.add_global_model`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.global_product.add_global_model?module=89&type=1](https://open.shopee.com/documents/v2/v2.global_product.add_global_model?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `category_recommend`

**전체 이름**: `v2.global_product.category_recommend`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.global_product.category_recommend?module=89&type=1](https://open.shopee.com/documents/v2/v2.global_product.category_recommend?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `create_publish_task`

**전체 이름**: `v2.global_product.create_publish_task`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.global_product.create_publish_task?module=89&type=1](https://open.shopee.com/documents/v2/v2.global_product.create_publish_task?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `delete_global_item`

**전체 이름**: `v2.global_product.delete_global_item`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.global_product.delete_global_item?module=89&type=1](https://open.shopee.com/documents/v2/v2.global_product.delete_global_item?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


---

## 📢 ADS

**API 개수**: 26개

**설명**: 광고 캠페인 관리 - CPC 광고, 상품 광고, GMS 캠페인 등 광고 관련 모든 기능

| API 메서드 | HTTP | 설명 | 문서 |
|-----------|:----:|-----|:----:|
| `check_create_gms_product_campaign_eligibility` | POST | Documentation - Shopee Open Platform | ❌ |
| `create_auto_product_ads` | POST | Documentation - Shopee Open Platform | ❌ |
| `create_gms_product_campaign` | POST | Documentation - Shopee Open Platform | ❌ |
| `create_manual_product_ads` | POST | Documentation - Shopee Open Platform | ❌ |
| `edit_auto_product_ads` | PUT | Documentation - Shopee Open Platform | ❌ |
| `edit_gms_item_product_campaign` | PUT | Documentation - Shopee Open Platform | ❌ |
| `edit_gms_product_campaign` | PUT | Documentation - Shopee Open Platform | ❌ |
| `edit_manual_product_ad_keywords` | PUT | Documentation - Shopee Open Platform | ❌ |
| `edit_manual_product_ads` | PUT | Documentation - Shopee Open Platform | ❌ |
| `get_ads_fácil_shop_rate` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_ads_f찼cil_shop_rate` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_all_cpc_ads_daily_performance` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_all_cpc_ads_hourly_performance` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_create_product_ad_budget_suggestion` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_gms_campaign_performance` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_gms_item_performance` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_product_campaign_daily_performance` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_product_campaign_hourly_performance` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_product_level_campaign_id_list` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_product_level_campaign_setting_info` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_product_recommended_roi_target` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_recommended_item_list` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_recommended_keyword_list` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_shop_toggle_info` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_total_balance` | GET | Documentation - Shopee Open Platform | ❌ |
| `list_gms_user_deleted_item` | POST | Documentation - Shopee Open Platform | ❌ |

### 주요 API 상세 정보

#### `check_create_gms_product_campaign_eligibility`

**전체 이름**: `v2.ads.check_create_gms_product_campaign_eligibility`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.ads.check_create_gms_product_campaign_eligibility?module=89&type=1](https://open.shopee.com/documents/v2/v2.ads.check_create_gms_product_campaign_eligibility?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `create_auto_product_ads`

**전체 이름**: `v2.ads.create_auto_product_ads`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.ads.create_auto_product_ads?module=89&type=1](https://open.shopee.com/documents/v2/v2.ads.create_auto_product_ads?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `create_gms_product_campaign`

**전체 이름**: `v2.ads.create_gms_product_campaign`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.ads.create_gms_product_campaign?module=89&type=1](https://open.shopee.com/documents/v2/v2.ads.create_gms_product_campaign?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `create_manual_product_ads`

**전체 이름**: `v2.ads.create_manual_product_ads`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.ads.create_manual_product_ads?module=89&type=1](https://open.shopee.com/documents/v2/v2.ads.create_manual_product_ads?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `edit_auto_product_ads`

**전체 이름**: `v2.ads.edit_auto_product_ads`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.ads.edit_auto_product_ads?module=89&type=1](https://open.shopee.com/documents/v2/v2.ads.edit_auto_product_ads?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


---

## 📹 LIVESTREAM

**API 개수**: 25개

**설명**: 라이브 스트리밍 관리 - 라이브 커머스 세션, 상품, 댓글 등 관리

| API 메서드 | HTTP | 설명 | 문서 |
|-----------|:----:|-----|:----:|
| `add_item_list` | POST | Documentation - Shopee Open Platform | ❌ |
| `apply_item_set` | POST | Documentation - Shopee Open Platform | ❌ |
| `ban_user_comment` | POST | Documentation - Shopee Open Platform | ❌ |
| `create_session` | POST | Documentation - Shopee Open Platform | ❌ |
| `delete_item_list` | DELETE | Documentation - Shopee Open Platform | ❌ |
| `delete_show_item` | DELETE | Documentation - Shopee Open Platform | ❌ |
| `end_session` | POST | Documentation - Shopee Open Platform | ❌ |
| `get_item_count` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_item_list` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_item_set_item_list` | GET | Shopee Open Platform | ❌ |
| `get_item_set_list` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_latest_comment_list` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_like_item_list` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_recent_item_list` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_session_detail` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_session_item_metric` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_session_metric` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_show_item` | GET | Documentation - Shopee Open Platform | ❌ |
| `post_comment` | POST | Documentation - Shopee Open Platform | ❌ |
| `start_session` | POST | Documentation - Shopee Open Platform | ❌ |
| `unban_user_comment` | POST | Documentation - Shopee Open Platform | ❌ |
| `update_item_list` | PUT | Documentation - Shopee Open Platform | ❌ |
| `update_session` | PUT | Documentation - Shopee Open Platform | ❌ |
| `update_show_item` | PUT | Documentation - Shopee Open Platform | ❌ |
| `upload_image` | POST | Documentation - Shopee Open Platform | ❌ |

### 주요 API 상세 정보

#### `add_item_list`

**전체 이름**: `v2.livestream.add_item_list`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.livestream.add_item_list?module=89&type=1](https://open.shopee.com/documents/v2/v2.livestream.add_item_list?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `apply_item_set`

**전체 이름**: `v2.livestream.apply_item_set`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.livestream.apply_item_set?module=89&type=1](https://open.shopee.com/documents/v2/v2.livestream.apply_item_set?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `ban_user_comment`

**전체 이름**: `v2.livestream.ban_user_comment`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.livestream.ban_user_comment?module=89&type=1](https://open.shopee.com/documents/v2/v2.livestream.ban_user_comment?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `create_session`

**전체 이름**: `v2.livestream.create_session`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.livestream.create_session?module=89&type=1](https://open.shopee.com/documents/v2/v2.livestream.create_session?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `delete_item_list`

**전체 이름**: `v2.livestream.delete_item_list`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.livestream.delete_item_list?module=89&type=1](https://open.shopee.com/documents/v2/v2.livestream.delete_item_list?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


---

## 📦 ORDER

**API 개수**: 21개

**설명**: 주문 관리 - 주문 조회, 취소, 송장 생성 등 주문 관련 모든 기능

| API 메서드 | HTTP | 설명 | 문서 |
|-----------|:----:|-----|:----:|
| `cancel_order` | POST | Documentation - Shopee Open Platform | ❌ |
| `download_fbs_invoices` | POST | Documentation - Shopee Open Platform | ❌ |
| `download_invoice_doc` | POST | Documentation - Shopee Open Platform | ❌ |
| `generate_fbs_invoices` | POST | Documentation - Shopee Open Platform | ❌ |
| `get_booking_detail` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_booking_list` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_buyer_invoice_info` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_fbs_invoices_result` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_order_detail` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_order_list` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_package_detail` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_pending_buyer_invoice_order_list` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_shipment_list` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_warehouse_filter_config` | GET | Documentation - Shopee Open Platform | ❌ |
| `handle_buyer_cancellation` | POST | Documentation - Shopee Open Platform | ❌ |
| `handle_prescription_check` | POST | Documentation - Shopee Open Platform | ❌ |
| `search_package_list` | POST | Documentation - Shopee Open Platform | ❌ |
| `set_note` | POST | Documentation - Shopee Open Platform | ❌ |
| `split_order` | POST | Documentation - Shopee Open Platform | ❌ |
| `unsplit_order` | POST | Documentation - Shopee Open Platform | ❌ |
| `upload_invoice_doc` | POST | Documentation - Shopee Open Platform | ❌ |

### 주요 API 상세 정보

#### `cancel_order`

**전체 이름**: `v2.order.cancel_order`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.order.cancel_order?module=89&type=1](https://open.shopee.com/documents/v2/v2.order.cancel_order?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `download_fbs_invoices`

**전체 이름**: `v2.order.download_fbs_invoices`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.order.download_fbs_invoices?module=89&type=1](https://open.shopee.com/documents/v2/v2.order.download_fbs_invoices?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `download_invoice_doc`

**전체 이름**: `v2.order.download_invoice_doc`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.order.download_invoice_doc?module=89&type=1](https://open.shopee.com/documents/v2/v2.order.download_invoice_doc?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `generate_fbs_invoices`

**전체 이름**: `v2.order.generate_fbs_invoices`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.order.generate_fbs_invoices?module=89&type=1](https://open.shopee.com/documents/v2/v2.order.generate_fbs_invoices?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `get_booking_detail`

**전체 이름**: `v2.order.get_booking_detail`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.order.get_booking_detail?module=89&type=1](https://open.shopee.com/documents/v2/v2.order.get_booking_detail?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


---

## 🚀 FIRST_MILE

**API 개수**: 16개

**설명**: 퍼스트 마일 배송 관리 - 판매자에서 물류센터까지의 첫 배송 단계

| API 메서드 | HTTP | 설명 | 문서 |
|-----------|:----:|-----|:----:|
| `bind_courier_delivery_first_mile_tracking_number` | POST | Documentation - Shopee Open Platform | ❌ |
| `bind_first_mile_tracking_number` | POST | Documentation - Shopee Open Platform | ❌ |
| `generate_and_bind_first_mile_tracking_number` | POST | Documentation - Shopee Open Platform | ❌ |
| `generate_first_mile_tracking_number` | POST | Documentation - Shopee Open Platform | ❌ |
| `get_channel_list` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_courier_delivery_channel_list` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_courier_delivery_detail` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_courier_delivery_tracking_number_list` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_courier_delivery_waybill` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_detail` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_tracking_number_list` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_transit_warehouse_list` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_unbind_order_list` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_waybill` | GET | Documentation - Shopee Open Platform | ❌ |
| `unbind_first_mile_tracking_number` | POST | Documentation - Shopee Open Platform | ❌ |
| `unbind_first_mile_tracking_number_all` | POST | Documentation - Shopee Open Platform | ❌ |

### 주요 API 상세 정보

#### `bind_courier_delivery_first_mile_tracking_number`

**전체 이름**: `v2.first_mile.bind_courier_delivery_first_mile_tracking_number`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.first_mile.bind_courier_delivery_first_mile_tracking_number?module=89&type=1](https://open.shopee.com/documents/v2/v2.first_mile.bind_courier_delivery_first_mile_tracking_number?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `bind_first_mile_tracking_number`

**전체 이름**: `v2.first_mile.bind_first_mile_tracking_number`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.first_mile.bind_first_mile_tracking_number?module=89&type=1](https://open.shopee.com/documents/v2/v2.first_mile.bind_first_mile_tracking_number?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `generate_and_bind_first_mile_tracking_number`

**전체 이름**: `v2.first_mile.generate_and_bind_first_mile_tracking_number`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.first_mile.generate_and_bind_first_mile_tracking_number?module=89&type=1](https://open.shopee.com/documents/v2/v2.first_mile.generate_and_bind_first_mile_tracking_number?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `generate_first_mile_tracking_number`

**전체 이름**: `v2.first_mile.generate_first_mile_tracking_number`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.first_mile.generate_first_mile_tracking_number?module=89&type=1](https://open.shopee.com/documents/v2/v2.first_mile.generate_first_mile_tracking_number?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `get_channel_list`

**전체 이름**: `v2.first_mile.get_channel_list`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.first_mile.get_channel_list?module=89&type=1](https://open.shopee.com/documents/v2/v2.first_mile.get_channel_list?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


---

## 💰 PAYMENT

**API 개수**: 16개

**설명**: 결제 및 정산 관리 - 에스크로, 수익 보고서, 지갑 거래 등

| API 메서드 | HTTP | 설명 | 문서 |
|-----------|:----:|-----|:----:|
| `generate_income_report` | POST | Documentation - Shopee Open Platform | ❌ |
| `generate_income_statement` | POST | Documentation - Shopee Open Platform | ❌ |
| `get_billing_transaction_info` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_escrow_detail` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_escrow_detail_batch` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_escrow_list` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_income_report` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_income_statement` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_item_installment_status` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_payment_method_list` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_payout_detail` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_payout_info` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_shop_installment_status` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_wallet_transaction_list` | GET | Documentation - Shopee Open Platform | ❌ |
| `set_item_installment_status` | POST | Documentation - Shopee Open Platform | ❌ |
| `set_shop_installment_status` | POST | Documentation - Shopee Open Platform | ❌ |

### 주요 API 상세 정보

#### `generate_income_report`

**전체 이름**: `v2.payment.generate_income_report`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.payment.generate_income_report?module=89&type=1](https://open.shopee.com/documents/v2/v2.payment.generate_income_report?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `generate_income_statement`

**전체 이름**: `v2.payment.generate_income_statement`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.payment.generate_income_statement?module=89&type=1](https://open.shopee.com/documents/v2/v2.payment.generate_income_statement?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `get_billing_transaction_info`

**전체 이름**: `v2.payment.get_billing_transaction_info`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.payment.get_billing_transaction_info?module=89&type=1](https://open.shopee.com/documents/v2/v2.payment.get_billing_transaction_info?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `get_escrow_detail`

**전체 이름**: `v2.payment.get_escrow_detail`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.payment.get_escrow_detail?module=89&type=1](https://open.shopee.com/documents/v2/v2.payment.get_escrow_detail?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `get_escrow_detail_batch`

**전체 이름**: `v2.payment.get_escrow_detail_batch`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.payment.get_escrow_detail_batch?module=89&type=1](https://open.shopee.com/documents/v2/v2.payment.get_escrow_detail_batch?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


---

## ↩️ RETURNS

**API 개수**: 15개

**설명**: 반품 및 환불 관리 - 반품 요청, 분쟁, 해결 방안 등

| API 메서드 | HTTP | 설명 | 문서 |
|-----------|:----:|-----|:----:|
| `accept_offer` | POST | Documentation - Shopee Open Platform | ❌ |
| `cancel_dispute` | POST | Documentation - Shopee Open Platform | ❌ |
| `confirm` | POST | Documentation - Shopee Open Platform | ❌ |
| `convert_image` | POST | Documentation - Shopee Open Platform | ❌ |
| `dispute` | POST | Documentation - Shopee Open Platform | ❌ |
| `get_available_solutions` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_return_detail` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_return_dispute_reason` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_return_list` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_reverse_tracking_info` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_shipping_carrier` | GET | Documentation - Shopee Open Platform | ❌ |
| `offer` | POST | Documentation - Shopee Open Platform | ❌ |
| `query_proof` | POST | Documentation - Shopee Open Platform | ❌ |
| `upload_proof` | POST | Documentation - Shopee Open Platform | ❌ |
| `upload_shipping_proof` | POST | Documentation - Shopee Open Platform | ❌ |

### 주요 API 상세 정보

#### `accept_offer`

**전체 이름**: `v2.returns.accept_offer`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.returns.accept_offer?module=89&type=1](https://open.shopee.com/documents/v2/v2.returns.accept_offer?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `cancel_dispute`

**전체 이름**: `v2.returns.cancel_dispute`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.returns.cancel_dispute?module=89&type=1](https://open.shopee.com/documents/v2/v2.returns.cancel_dispute?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `confirm`

**전체 이름**: `v2.returns.confirm`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.returns.confirm?module=89&type=1](https://open.shopee.com/documents/v2/v2.returns.confirm?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `convert_image`

**전체 이름**: `v2.returns.convert_image`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.returns.convert_image?module=89&type=1](https://open.shopee.com/documents/v2/v2.returns.convert_image?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `dispute`

**전체 이름**: `v2.returns.dispute`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.returns.dispute?module=89&type=1](https://open.shopee.com/documents/v2/v2.returns.dispute?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


---

## ➕ ADD_ON_DEAL

**API 개수**: 14개

**설명**: 추가 상품 딜 관리 - 메인 상품과 함께 구매할 수 있는 추가 상품 프로모션

| API 메서드 | HTTP | 설명 | 문서 |
|-----------|:----:|-----|:----:|
| `add_add_on_deal` | POST | Documentation - Shopee Open Platform | ❌ |
| `add_add_on_deal_main_item` | POST | Documentation - Shopee Open Platform | ❌ |
| `add_add_on_deal_sub_item` | POST | Documentation - Shopee Open Platform | ❌ |
| `delete_add_on_deal` | DELETE | Documentation - Shopee Open Platform | ❌ |
| `delete_add_on_deal_main_item` | DELETE | Documentation - Shopee Open Platform | ❌ |
| `delete_add_on_deal_sub_item` | DELETE | Documentation - Shopee Open Platform | ❌ |
| `end_add_on_deal` | POST | Documentation - Shopee Open Platform | ❌ |
| `get_add_on_deal` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_add_on_deal_list` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_add_on_deal_main_item` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_add_on_deal_sub_item` | GET | Documentation - Shopee Open Platform | ❌ |
| `update_add_on_deal` | PUT | Documentation - Shopee Open Platform | ❌ |
| `update_add_on_deal_main_item` | PUT | Documentation - Shopee Open Platform | ❌ |
| `update_add_on_deal_sub_item` | PUT | Documentation - Shopee Open Platform | ❌ |

### 주요 API 상세 정보

#### `add_add_on_deal`

**전체 이름**: `v2.add_on_deal.add_add_on_deal`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.add_on_deal.add_add_on_deal?module=89&type=1](https://open.shopee.com/documents/v2/v2.add_on_deal.add_add_on_deal?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `add_add_on_deal_main_item`

**전체 이름**: `v2.add_on_deal.add_add_on_deal_main_item`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.add_on_deal.add_add_on_deal_main_item?module=89&type=1](https://open.shopee.com/documents/v2/v2.add_on_deal.add_add_on_deal_main_item?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `add_add_on_deal_sub_item`

**전체 이름**: `v2.add_on_deal.add_add_on_deal_sub_item`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.add_on_deal.add_add_on_deal_sub_item?module=89&type=1](https://open.shopee.com/documents/v2/v2.add_on_deal.add_add_on_deal_sub_item?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `delete_add_on_deal`

**전체 이름**: `v2.add_on_deal.delete_add_on_deal`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.add_on_deal.delete_add_on_deal?module=89&type=1](https://open.shopee.com/documents/v2/v2.add_on_deal.delete_add_on_deal?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `delete_add_on_deal_main_item`

**전체 이름**: `v2.add_on_deal.delete_add_on_deal_main_item`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.add_on_deal.delete_add_on_deal_main_item?module=89&type=1](https://open.shopee.com/documents/v2/v2.add_on_deal.delete_add_on_deal_main_item?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


---

## 📌 SHOP_FLASH_SALE

**API 개수**: 11개

**설명**: 샵 플래시 세일 - 샵 자체 플래시 세일 이벤트

| API 메서드 | HTTP | 설명 | 문서 |
|-----------|:----:|-----|:----:|
| `add_shop_flash_sale_items` | POST | Documentation - Shopee Open Platform | ❌ |
| `create_shop_flash_sale` | POST | Documentation - Shopee Open Platform | ❌ |
| `delete_shop_flash_sale` | DELETE | Documentation - Shopee Open Platform | ❌ |
| `delete_shop_flash_sale_items` | DELETE | Documentation - Shopee Open Platform | ❌ |
| `get_item_criteria` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_shop_flash_sale` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_shop_flash_sale_items` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_shop_flash_sale_list` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_time_slot_id` | GET | Documentation - Shopee Open Platform | ❌ |
| `update_shop_flash_sale` | PUT | Documentation - Shopee Open Platform | ❌ |
| `update_shop_flash_sale_items` | PUT | Documentation - Shopee Open Platform | ❌ |

### 주요 API 상세 정보

#### `add_shop_flash_sale_items`

**전체 이름**: `v2.shop_flash_sale.add_shop_flash_sale_items`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.shop_flash_sale.add_shop_flash_sale_items?module=89&type=1](https://open.shopee.com/documents/v2/v2.shop_flash_sale.add_shop_flash_sale_items?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `create_shop_flash_sale`

**전체 이름**: `v2.shop_flash_sale.create_shop_flash_sale`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.shop_flash_sale.create_shop_flash_sale?module=89&type=1](https://open.shopee.com/documents/v2/v2.shop_flash_sale.create_shop_flash_sale?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `delete_shop_flash_sale`

**전체 이름**: `v2.shop_flash_sale.delete_shop_flash_sale`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.shop_flash_sale.delete_shop_flash_sale?module=89&type=1](https://open.shopee.com/documents/v2/v2.shop_flash_sale.delete_shop_flash_sale?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `delete_shop_flash_sale_items`

**전체 이름**: `v2.shop_flash_sale.delete_shop_flash_sale_items`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.shop_flash_sale.delete_shop_flash_sale_items?module=89&type=1](https://open.shopee.com/documents/v2/v2.shop_flash_sale.delete_shop_flash_sale_items?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `get_item_criteria`

**전체 이름**: `v2.shop_flash_sale.get_item_criteria`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.shop_flash_sale.get_item_criteria?module=89&type=1](https://open.shopee.com/documents/v2/v2.shop_flash_sale.get_item_criteria?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


---

## 📦🎁 BUNDLE_DEAL

**API 개수**: 10개

**설명**: 번들 상품 딜 관리 - 여러 상품을 묶어서 판매하는 프로모션

| API 메서드 | HTTP | 설명 | 문서 |
|-----------|:----:|-----|:----:|
| `add_bundle_deal` | POST | Documentation - Shopee Open Platform | ❌ |
| `add_bundle_deal_item` | POST | Documentation - Shopee Open Platform | ❌ |
| `delete_bundle_deal` | DELETE | Documentation - Shopee Open Platform | ❌ |
| `delete_bundle_deal_item` | DELETE | Documentation - Shopee Open Platform | ❌ |
| `end_bundle_deal` | POST | Documentation - Shopee Open Platform | ❌ |
| `get_bundle_deal` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_bundle_deal_item` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_bundle_deal_list` | GET | Documentation - Shopee Open Platform | ❌ |
| `update_bundle_deal` | PUT | Documentation - Shopee Open Platform | ❌ |
| `update_bundle_deal_item` | PUT | Documentation - Shopee Open Platform | ❌ |

### 주요 API 상세 정보

#### `add_bundle_deal`

**전체 이름**: `v2.bundle_deal.add_bundle_deal`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.bundle_deal.add_bundle_deal?module=89&type=1](https://open.shopee.com/documents/v2/v2.bundle_deal.add_bundle_deal?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `add_bundle_deal_item`

**전체 이름**: `v2.bundle_deal.add_bundle_deal_item`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.bundle_deal.add_bundle_deal_item?module=89&type=1](https://open.shopee.com/documents/v2/v2.bundle_deal.add_bundle_deal_item?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `delete_bundle_deal`

**전체 이름**: `v2.bundle_deal.delete_bundle_deal`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.bundle_deal.delete_bundle_deal?module=89&type=1](https://open.shopee.com/documents/v2/v2.bundle_deal.delete_bundle_deal?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `delete_bundle_deal_item`

**전체 이름**: `v2.bundle_deal.delete_bundle_deal_item`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.bundle_deal.delete_bundle_deal_item?module=89&type=1](https://open.shopee.com/documents/v2/v2.bundle_deal.delete_bundle_deal_item?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `end_bundle_deal`

**전체 이름**: `v2.bundle_deal.end_bundle_deal`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.bundle_deal.end_bundle_deal?module=89&type=1](https://open.shopee.com/documents/v2/v2.bundle_deal.end_bundle_deal?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


---

## 🏷️ DISCOUNT

**API 개수**: 9개

**설명**: 할인 프로모션 관리 - 상품별 할인 설정 및 관리

| API 메서드 | HTTP | 설명 | 문서 |
|-----------|:----:|-----|:----:|
| `add_discount` | POST | Documentation - Shopee Open Platform | ❌ |
| `add_discount_item` | POST | Documentation - Shopee Open Platform | ❌ |
| `delete_discount` | DELETE | Documentation - Shopee Open Platform | ❌ |
| `delete_discount_item` | DELETE | Documentation - Shopee Open Platform | ❌ |
| `end_discount` | POST | Documentation - Shopee Open Platform | ❌ |
| `get_discount` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_discount_list` | GET | Documentation - Shopee Open Platform | ❌ |
| `update_discount` | PUT | Documentation - Shopee Open Platform | ❌ |
| `update_discount_item` | PUT | Documentation - Shopee Open Platform | ❌ |

### 주요 API 상세 정보

#### `add_discount`

**전체 이름**: `v2.discount.add_discount`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.discount.add_discount?module=89&type=1](https://open.shopee.com/documents/v2/v2.discount.add_discount?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `add_discount_item`

**전체 이름**: `v2.discount.add_discount_item`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.discount.add_discount_item?module=89&type=1](https://open.shopee.com/documents/v2/v2.discount.add_discount_item?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `delete_discount`

**전체 이름**: `v2.discount.delete_discount`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.discount.delete_discount?module=89&type=1](https://open.shopee.com/documents/v2/v2.discount.delete_discount?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `delete_discount_item`

**전체 이름**: `v2.discount.delete_discount_item`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.discount.delete_discount_item?module=89&type=1](https://open.shopee.com/documents/v2/v2.discount.delete_discount_item?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `end_discount`

**전체 이름**: `v2.discount.end_discount`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.discount.end_discount?module=89&type=1](https://open.shopee.com/documents/v2/v2.discount.end_discount?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


---

## 📌 SHOP_CATEGORY

**API 개수**: 7개

**설명**: 샵 카테고리 관리 - 샵 내부 카테고리 및 상품 분류

| API 메서드 | HTTP | 설명 | 문서 |
|-----------|:----:|-----|:----:|
| `add_item_list` | POST | Documentation - Shopee Open Platform | ❌ |
| `add_shop_category` | POST | Documentation - Shopee Open Platform | ❌ |
| `delete_item_list` | DELETE | Documentation - Shopee Open Platform | ❌ |
| `delete_shop_category` | DELETE | Documentation - Shopee Open Platform | ❌ |
| `get_item_list` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_shop_category_list` | GET | Documentation - Shopee Open Platform | ❌ |
| `update_shop_category` | PUT | Documentation - Shopee Open Platform | ❌ |

### 주요 API 상세 정보

#### `add_item_list`

**전체 이름**: `v2.shop_category.add_item_list`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.shop_category.add_item_list?module=89&type=1](https://open.shopee.com/documents/v2/v2.shop_category.add_item_list?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `add_shop_category`

**전체 이름**: `v2.shop_category.add_shop_category`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.shop_category.add_shop_category?module=89&type=1](https://open.shopee.com/documents/v2/v2.shop_category.add_shop_category?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `delete_item_list`

**전체 이름**: `v2.shop_category.delete_item_list`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.shop_category.delete_item_list?module=89&type=1](https://open.shopee.com/documents/v2/v2.shop_category.delete_item_list?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `delete_shop_category`

**전체 이름**: `v2.shop_category.delete_shop_category`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.shop_category.delete_shop_category?module=89&type=1](https://open.shopee.com/documents/v2/v2.shop_category.delete_shop_category?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `get_item_list`

**전체 이름**: `v2.shop_category.get_item_list`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.shop_category.get_item_list?module=89&type=1](https://open.shopee.com/documents/v2/v2.shop_category.get_item_list?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


---

## 💊 ACCOUNT_HEALTH

**API 개수**: 6개

**설명**: 계정 건강도 및 성과 모니터링 - 판매자의 계정 상태, 지연 주문, 문제 있는 리스팅 등을 추적

| API 메서드 | HTTP | 설명 | 문서 |
|-----------|:----:|-----|:----:|
| `get_late_orders` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_listings_with_issues` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_metric_source_detail` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_penalty_point_history` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_punishment_history` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_shop_performance` | GET | Documentation - Shopee Open Platform | ❌ |

### 주요 API 상세 정보

#### `get_late_orders`

**전체 이름**: `v2.account_health.get_late_orders`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.account_health.get_late_orders?module=103&type=1](https://open.shopee.com/documents/v2/v2.account_health.get_late_orders?module=103&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `get_listings_with_issues`

**전체 이름**: `v2.account_health.get_listings_with_issues`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.account_health.get_listings_with_issues?module=89&type=1](https://open.shopee.com/documents/v2/v2.account_health.get_listings_with_issues?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `get_metric_source_detail`

**전체 이름**: `v2.account_health.get_metric_source_detail`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.account_health.get_metric_source_detail?module=89&type=1](https://open.shopee.com/documents/v2/v2.account_health.get_metric_source_detail?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `get_penalty_point_history`

**전체 이름**: `v2.account_health.get_penalty_point_history`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.account_health.get_penalty_point_history?module=89&type=1](https://open.shopee.com/documents/v2/v2.account_health.get_penalty_point_history?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `get_punishment_history`

**전체 이름**: `v2.account_health.get_punishment_history`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.account_health.get_punishment_history?module=89&type=1](https://open.shopee.com/documents/v2/v2.account_health.get_punishment_history?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


---

## 📌 FOLLOW_PRIZE

**API 개수**: 6개

**설명**: 팔로우 경품 이벤트 - 샵 팔로우 고객 대상 경품 프로모션

| API 메서드 | HTTP | 설명 | 문서 |
|-----------|:----:|-----|:----:|
| `add_follow_prize` | POST | Documentation - Shopee Open Platform | ❌ |
| `delete_follow_prize` | DELETE | Documentation - Shopee Open Platform | ❌ |
| `end_follow_prize` | POST | Documentation - Shopee Open Platform | ❌ |
| `get_follow_prize_detail` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_follow_prize_list` | GET | Documentation - Shopee Open Platform | ❌ |
| `update_follow_prize` | PUT | Documentation - Shopee Open Platform | ❌ |

### 주요 API 상세 정보

#### `add_follow_prize`

**전체 이름**: `v2.follow_prize.add_follow_prize`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.follow_prize.add_follow_prize?module=89&type=1](https://open.shopee.com/documents/v2/v2.follow_prize.add_follow_prize?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `delete_follow_prize`

**전체 이름**: `v2.follow_prize.delete_follow_prize`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.follow_prize.delete_follow_prize?module=89&type=1](https://open.shopee.com/documents/v2/v2.follow_prize.delete_follow_prize?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `end_follow_prize`

**전체 이름**: `v2.follow_prize.end_follow_prize`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.follow_prize.end_follow_prize?module=89&type=1](https://open.shopee.com/documents/v2/v2.follow_prize.end_follow_prize?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `get_follow_prize_detail`

**전체 이름**: `v2.follow_prize.get_follow_prize_detail`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.follow_prize.get_follow_prize_detail?module=89&type=1](https://open.shopee.com/documents/v2/v2.follow_prize.get_follow_prize_detail?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `get_follow_prize_list`

**전체 이름**: `v2.follow_prize.get_follow_prize_list`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.follow_prize.get_follow_prize_list?module=89&type=1](https://open.shopee.com/documents/v2/v2.follow_prize.get_follow_prize_list?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


---

## 🎬 MEDIA_SPACE

**API 개수**: 6개

**설명**: 미디어 스페이스 관리 - 비디오 업로드 및 미디어 관리

| API 메서드 | HTTP | 설명 | 문서 |
|-----------|:----:|-----|:----:|
| `cancel_video_upload` | POST | Documentation - Shopee Open Platform | ❌ |
| `complete_video_upload` | POST | Documentation - Shopee Open Platform | ❌ |
| `get_video_upload_result` | GET | Documentation - Shopee Open Platform | ❌ |
| `init_video_upload` | POST | Documentation - Shopee Open Platform | ❌ |
| `upload_image` | POST | Documentation - Shopee Open Platform | ❌ |
| `upload_video_part` | POST | Documentation - Shopee Open Platform | ❌ |

### 주요 API 상세 정보

#### `cancel_video_upload`

**전체 이름**: `v2.media_space.cancel_video_upload`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.media_space.cancel_video_upload?module=89&type=1](https://open.shopee.com/documents/v2/v2.media_space.cancel_video_upload?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `complete_video_upload`

**전체 이름**: `v2.media_space.complete_video_upload`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.media_space.complete_video_upload?module=89&type=1](https://open.shopee.com/documents/v2/v2.media_space.complete_video_upload?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `get_video_upload_result`

**전체 이름**: `v2.media_space.get_video_upload_result`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.media_space.get_video_upload_result?module=89&type=1](https://open.shopee.com/documents/v2/v2.media_space.get_video_upload_result?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `init_video_upload`

**전체 이름**: `v2.media_space.init_video_upload`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.media_space.init_video_upload?module=89&type=1](https://open.shopee.com/documents/v2/v2.media_space.init_video_upload?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `upload_image`

**전체 이름**: `v2.media_space.upload_image`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.media_space.upload_image?module=89&type=1](https://open.shopee.com/documents/v2/v2.media_space.upload_image?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


---

## 👤 MERCHANT

**API 개수**: 6개

**설명**: 판매자 정보 관리 - 판매자(파트너) 계정 및 창고 정보

| API 메서드 | HTTP | 설명 | 문서 |
|-----------|:----:|-----|:----:|
| `get_merchant_info` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_merchant_prepaid_account_list` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_merchant_warehouse_list` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_merchant_warehouse_location_list` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_shop_list_by_merchant` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_warehouse_eligible_shop_list` | GET | Documentation - Shopee Open Platform | ❌ |

### 주요 API 상세 정보

#### `get_merchant_info`

**전체 이름**: `v2.merchant.get_merchant_info`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.merchant.get_merchant_info?module=89&type=1](https://open.shopee.com/documents/v2/v2.merchant.get_merchant_info?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `get_merchant_prepaid_account_list`

**전체 이름**: `v2.merchant.get_merchant_prepaid_account_list`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.merchant.get_merchant_prepaid_account_list?module=89&type=1](https://open.shopee.com/documents/v2/v2.merchant.get_merchant_prepaid_account_list?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `get_merchant_warehouse_list`

**전체 이름**: `v2.merchant.get_merchant_warehouse_list`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.merchant.get_merchant_warehouse_list?module=89&type=1](https://open.shopee.com/documents/v2/v2.merchant.get_merchant_warehouse_list?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `get_merchant_warehouse_location_list`

**전체 이름**: `v2.merchant.get_merchant_warehouse_location_list`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.merchant.get_merchant_warehouse_location_list?module=89&type=1](https://open.shopee.com/documents/v2/v2.merchant.get_merchant_warehouse_location_list?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `get_shop_list_by_merchant`

**전체 이름**: `v2.merchant.get_shop_list_by_merchant`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.merchant.get_shop_list_by_merchant?module=89&type=1](https://open.shopee.com/documents/v2/v2.merchant.get_shop_list_by_merchant?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


---

## 📌 PUBLIC

**API 개수**: 6개

**설명**: 공개 API - 인증, 토큰 발급, IP 범위 등 공개적으로 접근 가능한 API

| API 메서드 | HTTP | 설명 | 문서 |
|-----------|:----:|-----|:----:|
| `get_access_token` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_merchants_by_partner` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_shopee_ip_ranges` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_shops_by_partner` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_token_by_resend_code` | GET | Documentation - Shopee Open Platform | ❌ |
| `refresh_access_token` | POST | Documentation - Shopee Open Platform | ❌ |

### 주요 API 상세 정보

#### `get_access_token`

**전체 이름**: `v2.public.get_access_token`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.public.get_access_token?module=89&type=1](https://open.shopee.com/documents/v2/v2.public.get_access_token?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `get_merchants_by_partner`

**전체 이름**: `v2.public.get_merchants_by_partner`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.public.get_merchants_by_partner?module=89&type=1](https://open.shopee.com/documents/v2/v2.public.get_merchants_by_partner?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `get_shopee_ip_ranges`

**전체 이름**: `v2.public.get_shopee_ip_ranges`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.public.get_shopee_ip_ranges?module=89&type=1](https://open.shopee.com/documents/v2/v2.public.get_shopee_ip_ranges?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `get_shops_by_partner`

**전체 이름**: `v2.public.get_shops_by_partner`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.public.get_shops_by_partner?module=89&type=1](https://open.shopee.com/documents/v2/v2.public.get_shops_by_partner?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `get_token_by_resend_code`

**전체 이름**: `v2.public.get_token_by_resend_code`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.public.get_token_by_resend_code?module=89&type=1](https://open.shopee.com/documents/v2/v2.public.get_token_by_resend_code?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


---

## 🏪 SHOP

**API 개수**: 6개

**설명**: 샵 정보 관리 - 샵 프로필, 창고 정보 등

| API 메서드 | HTTP | 설명 | 문서 |
|-----------|:----:|-----|:----:|
| `get_authorised_reseller_brand` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_profile` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_shop_info` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_shop_notification` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_warehouse_detail` | GET | Documentation - Shopee Open Platform | ❌ |
| `update_profile` | PUT | Documentation - Shopee Open Platform | ❌ |

### 주요 API 상세 정보

#### `get_authorised_reseller_brand`

**전체 이름**: `v2.shop.get_authorised_reseller_brand`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.shop.get_authorised_reseller_brand?module=89&type=1](https://open.shopee.com/documents/v2/v2.shop.get_authorised_reseller_brand?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `get_profile`

**전체 이름**: `v2.shop.get_profile`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.shop.get_profile?module=89&type=1](https://open.shopee.com/documents/v2/v2.shop.get_profile?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `get_shop_info`

**전체 이름**: `v2.shop.get_shop_info`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.shop.get_shop_info?module=89&type=1](https://open.shopee.com/documents/v2/v2.shop.get_shop_info?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `get_shop_notification`

**전체 이름**: `v2.shop.get_shop_notification`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.shop.get_shop_notification?module=89&type=1](https://open.shopee.com/documents/v2/v2.shop.get_shop_notification?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `get_warehouse_detail`

**전체 이름**: `v2.shop.get_warehouse_detail`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.shop.get_warehouse_detail?module=89&type=1](https://open.shopee.com/documents/v2/v2.shop.get_warehouse_detail?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


---

## 🎟️ VOUCHER

**API 개수**: 6개

**설명**: 쿠폰/바우처 관리 - 할인 쿠폰 생성 및 관리

| API 메서드 | HTTP | 설명 | 문서 |
|-----------|:----:|-----|:----:|
| `add_voucher` | POST | Documentation - Shopee Open Platform | ❌ |
| `delete_voucher` | DELETE | Documentation - Shopee Open Platform | ❌ |
| `end_voucher` | POST | Documentation - Shopee Open Platform | ❌ |
| `get_voucher` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_voucher_list` | GET | Documentation - Shopee Open Platform | ❌ |
| `update_voucher` | PUT | Documentation - Shopee Open Platform | ❌ |

### 주요 API 상세 정보

#### `add_voucher`

**전체 이름**: `v2.voucher.add_voucher`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.voucher.add_voucher?module=89&type=1](https://open.shopee.com/documents/v2/v2.voucher.add_voucher?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `delete_voucher`

**전체 이름**: `v2.voucher.delete_voucher`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.voucher.delete_voucher?module=89&type=1](https://open.shopee.com/documents/v2/v2.voucher.delete_voucher?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `end_voucher`

**전체 이름**: `v2.voucher.end_voucher`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.voucher.end_voucher?module=89&type=1](https://open.shopee.com/documents/v2/v2.voucher.end_voucher?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `get_voucher`

**전체 이름**: `v2.voucher.get_voucher`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.voucher.get_voucher?module=89&type=1](https://open.shopee.com/documents/v2/v2.voucher.get_voucher?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `get_voucher_list`

**전체 이름**: `v2.voucher.get_voucher_list`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.voucher.get_voucher_list?module=89&type=1](https://open.shopee.com/documents/v2/v2.voucher.get_voucher_list?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


---

## 📌 SBS

**API 개수**: 5개

**설명**: SBS (Shopee Business Solutions) - 재고 추적, 만료 리포트 등

| API 메서드 | HTTP | 설명 | 문서 |
|-----------|:----:|-----|:----:|
| `get_bound_whs_info` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_current_inventory` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_expiry_report` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_stock_aging` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_stock_movement` | GET | Documentation - Shopee Open Platform | ❌ |

### 주요 API 상세 정보

#### `get_bound_whs_info`

**전체 이름**: `v2.sbs.get_bound_whs_info`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.sbs.get_bound_whs_info?module=89&type=1](https://open.shopee.com/documents/v2/v2.sbs.get_bound_whs_info?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `get_current_inventory`

**전체 이름**: `v2.sbs.get_current_inventory`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.sbs.get_current_inventory?module=89&type=1](https://open.shopee.com/documents/v2/v2.sbs.get_current_inventory?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `get_expiry_report`

**전체 이름**: `v2.sbs.get_expiry_report`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.sbs.get_expiry_report?module=89&type=1](https://open.shopee.com/documents/v2/v2.sbs.get_expiry_report?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `get_stock_aging`

**전체 이름**: `v2.sbs.get_stock_aging`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.sbs.get_stock_aging?module=89&type=1](https://open.shopee.com/documents/v2/v2.sbs.get_stock_aging?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `get_stock_movement`

**전체 이름**: `v2.sbs.get_stock_movement`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.sbs.get_stock_movement?module=89&type=1](https://open.shopee.com/documents/v2/v2.sbs.get_stock_movement?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


---

## 📌 FBS

**API 개수**: 4개

**설명**: FBS (Fulfilled By Shopee) - Shopee가 직접 물류를 담당하는 서비스

| API 메서드 | HTTP | 설명 | 문서 |
|-----------|:----:|-----|:----:|
| `query_br_shop_block_status` | POST | Documentation - Shopee Open Platform | ❌ |
| `query_br_shop_enrollment_status` | POST | Documentation - Shopee Open Platform | ❌ |
| `query_br_shop_invoice_error` | POST | Documentation - Shopee Open Platform | ❌ |
| `query_br_sku_block_status` | POST | Documentation - Shopee Open Platform | ❌ |

### 주요 API 상세 정보

#### `query_br_shop_block_status`

**전체 이름**: `v2.fbs.query_br_shop_block_status`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.fbs.query_br_shop_block_status?module=89&type=1](https://open.shopee.com/documents/v2/v2.fbs.query_br_shop_block_status?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `query_br_shop_enrollment_status`

**전체 이름**: `v2.fbs.query_br_shop_enrollment_status`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.fbs.query_br_shop_enrollment_status?module=89&type=1](https://open.shopee.com/documents/v2/v2.fbs.query_br_shop_enrollment_status?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `query_br_shop_invoice_error`

**전체 이름**: `v2.fbs.query_br_shop_invoice_error`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.fbs.query_br_shop_invoice_error?module=89&type=1](https://open.shopee.com/documents/v2/v2.fbs.query_br_shop_invoice_error?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `query_br_sku_block_status`

**전체 이름**: `v2.fbs.query_br_sku_block_status`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.fbs.query_br_sku_block_status?module=89&type=1](https://open.shopee.com/documents/v2/v2.fbs.query_br_sku_block_status?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


---

## 📌 PUSH

**API 개수**: 4개

**설명**: 푸시 알림 설정 - 웹훅 및 푸시 알림 설정

| API 메서드 | HTTP | 설명 | 문서 |
|-----------|:----:|-----|:----:|
| `confirm_consumed_lost_push_message` | POST | Documentation - Shopee Open Platform | ❌ |
| `get_app_push_config` | GET | Documentation - Shopee Open Platform | ❌ |
| `get_lost_push_message` | GET | Documentation - Shopee Open Platform | ❌ |
| `set_app_push_config` | POST | Documentation - Shopee Open Platform | ❌ |

### 주요 API 상세 정보

#### `confirm_consumed_lost_push_message`

**전체 이름**: `v2.push.confirm_consumed_lost_push_message`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.push.confirm_consumed_lost_push_message?module=89&type=1](https://open.shopee.com/documents/v2/v2.push.confirm_consumed_lost_push_message?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `get_app_push_config`

**전체 이름**: `v2.push.get_app_push_config`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.push.get_app_push_config?module=89&type=1](https://open.shopee.com/documents/v2/v2.push.get_app_push_config?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `get_lost_push_message`

**전체 이름**: `v2.push.get_lost_push_message`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.push.get_lost_push_message?module=89&type=1](https://open.shopee.com/documents/v2/v2.push.get_lost_push_message?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `set_app_push_config`

**전체 이름**: `v2.push.set_app_push_config`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.push.set_app_push_config?module=89&type=1](https://open.shopee.com/documents/v2/v2.push.set_app_push_config?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


---

## 📌 TOP_PICKS

**API 개수**: 4개

**설명**: 추천 상품 관리 - 샵의 추천 상품 설정

| API 메서드 | HTTP | 설명 | 문서 |
|-----------|:----:|-----|:----:|
| `add_top_picks` | POST | Documentation - Shopee Open Platform | ❌ |
| `delete_top_picks` | DELETE | Documentation - Shopee Open Platform | ❌ |
| `get_top_picks_list` | GET | Documentation - Shopee Open Platform | ❌ |
| `update_top_picks` | PUT | Documentation - Shopee Open Platform | ❌ |

### 주요 API 상세 정보

#### `add_top_picks`

**전체 이름**: `v2.top_picks.add_top_picks`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.top_picks.add_top_picks?module=89&type=1](https://open.shopee.com/documents/v2/v2.top_picks.add_top_picks?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `delete_top_picks`

**전체 이름**: `v2.top_picks.delete_top_picks`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.top_picks.delete_top_picks?module=89&type=1](https://open.shopee.com/documents/v2/v2.top_picks.delete_top_picks?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `get_top_picks_list`

**전체 이름**: `v2.top_picks.get_top_picks_list`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.top_picks.get_top_picks_list?module=89&type=1](https://open.shopee.com/documents/v2/v2.top_picks.get_top_picks_list?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


#### `update_top_picks`

**전체 이름**: `v2.top_picks.update_top_picks`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.top_picks.update_top_picks?module=89&type=1](https://open.shopee.com/documents/v2/v2.top_picks.update_top_picks?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


---

## 📌 V2

**API 개수**: 2개

| API 메서드 | HTTP | 설명 | 문서 |
|-----------|:----:|-----|:----:|
| `get_attribute_tree` | GET | Get Attribute Tree | 🌐 |
| `get_category` | GET | Get Category | 🌐 |

### 주요 API 상세 정보

#### `get_attribute_tree`

**전체 이름**: `v2.v2.product.get_attribute_tree`

**포함 자료**: 한글 번역, 스크린샷


#### `get_category`

**전체 이름**: `v2.v2.product.get_category`

**포함 자료**: 한글 번역, 스크린샷


---

## 📌 MEDIA

**API 개수**: 1개

**설명**: 미디어 업로드 - 이미지 업로드 기본 기능

| API 메서드 | HTTP | 설명 | 문서 |
|-----------|:----:|-----|:----:|
| `upload_image` | POST | Documentation - Shopee Open Platform | ❌ |

### 주요 API 상세 정보

#### `upload_image`

**전체 이름**: `v2.media.upload_image`

**공식 문서**: [https://open.shopee.com/documents/v2/v2.media.upload_image?module=89&type=1](https://open.shopee.com/documents/v2/v2.media.upload_image?module=89&type=1)

**제목**: Documentation - Shopee Open Platform

**포함 자료**: 메타데이터, 스크린샷, 이미지 청크


---

## 📖 API 사용 가이드

### 🔐 인증 및 권한

Shopee Open API 사용을 위한 필수 요소:

1. **Partner ID**: Shopee에서 발급받은 파트너 식별자
2. **Partner Key**: API 서명에 사용되는 비밀 키
3. **Access Token**: OAuth 인증을 통해 획득한 샵별 액세스 토큰
4. **Shop ID**: 작업 대상 샵의 고유 식별자

### 🌐 API 엔드포인트

```
Base URL: https://partner.shopeemobile.com
API Version: v2
Format: /api/v2/{category}/{method}
```

### 📝 요청 형식

```
HTTP Method: GET/POST/PUT/DELETE
Content-Type: application/json
```

**공통 헤더**:
- `partner_id`: 파트너 ID
- `timestamp`: Unix 타임스탬프
- `sign`: HMAC-SHA256 서명
- `access_token`: 액세스 토큰
- `shop_id`: 샵 ID

### 🎯 주요 사용 사례

#### 상품 관리

- 새 상품 등록 (`add_item`)
- 상품 정보 수정 (`update_item`)
- 상품 재고 업데이트 (`update_stock`)
- 상품 가격 변경 (`update_price`)
- 상품 목록 조회 (`get_item_list`)
- 카테고리 속성 조회 (`get_attribute_tree`)

#### 주문 처리

- 주문 목록 조회 (`get_order_list`)
- 주문 상세 정보 (`get_order_detail`)
- 주문 취소 (`cancel_order`)
- 패키지 정보 조회 (`get_package_detail`)
- 송장 생성 (`generate_fbs_invoices`)

#### 물류 관리

- 배송 방법 조회 (`get_channel_list`)
- 운송장 번호 생성 (`get_tracking_number`)
- 배송 문서 생성 (`create_shipping_document`)
- 배송 상태 업데이트 (`update_tracking_status`)
- 대량 출고 처리 (`mass_ship_order`)

#### 결제 및 정산

- 에스크로 내역 조회 (`get_escrow_list`)
- 수익 리포트 생성 (`generate_income_report`)
- 정산 정보 조회 (`get_payout_info`)
- 지갑 거래 내역 (`get_wallet_transaction_list`)

#### 프로모션 운영

- 할인 생성 (`add_discount`)
- 할인 상품 추가 (`add_discount_item`)
- 번들 딜 생성 (`add_bundle_deal`)
- 플래시 세일 설정 (`create_shop_flash_sale`)
- 쿠폰 발행 (`add_voucher`)

---

## 📚 부록

### ⚠️ 주의사항

1. **Rate Limiting**: 각 API마다 호출 제한이 있습니다. 과도한 요청 시 일시적으로 차단될 수 있습니다.
2. **인증 만료**: Access Token은 주기적으로 갱신이 필요합니다.
3. **데이터 동기화**: 일부 API는 실시간이 아닌 배치로 처리됩니다.
4. **국가별 차이**: 일부 API는 특정 국가에서만 사용 가능합니다.

### 🔗 참고 자료

- [Shopee Open Platform 공식 사이트](https://open.shopee.com/)
- [Developer Guide](https://open.shopee.com/developer-guide/4)
- [API 문서](https://open.shopee.com/documents)
- [샌드박스 테스트 환경](https://open.shopee.com/developer-guide/12)

### 📊 문서 범례

| 아이콘 | 의미 |
|:-----:|-----|
| 📝 | 마크다운 문서 있음 |
| 🌐 | 한글 번역 있음 |
| ❌ | 문서화 안됨 |

### 📅 변경 이력

- **2025-10-16**: v2.0 완전 분석 문서 생성
  - 362개 API 전체 분석
  - 27개 카테고리 분류
  - 메타데이터 및 번역 정보 통합

---

**본 문서는 자동 생성되었습니다.**

Generated by `analyze_shopee_apis_detailed.py` at 2025-10-16 16:04:45
