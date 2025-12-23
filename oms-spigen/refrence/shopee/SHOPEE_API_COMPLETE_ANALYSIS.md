# Shopee Open API 전체 분석 보고서

**생성일시**: 2025-10-16 16:02:22

---

## 📊 전체 통계

- **총 API 개수**: 362개
- **카테고리 개수**: 27개

### 카테고리별 API 분포

| 카테고리 | API 개수 |
|---------|----------|
| account_health | 6 |
| add_on_deal | 14 |
| ads | 26 |
| bundle_deal | 10 |
| discount | 9 |
| fbs | 4 |
| first_mile | 16 |
| follow_prize | 6 |
| global_product | 34 |
| livestream | 25 |
| logistics | 41 |
| media | 1 |
| media_space | 6 |
| merchant | 6 |
| order | 21 |
| payment | 16 |
| product | 55 |
| public | 6 |
| push | 4 |
| returns | 15 |
| sbs | 5 |
| shop | 6 |
| shop_category | 7 |
| shop_flash_sale | 11 |
| top_picks | 4 |
| v2 | 2 |
| voucher | 6 |

---

## 📑 목차

1. [ACCOUNT_HEALTH](#account_health)
2. [ADD_ON_DEAL](#add_on_deal)
3. [ADS](#ads)
4. [BUNDLE_DEAL](#bundle_deal)
5. [DISCOUNT](#discount)
6. [FBS](#fbs)
7. [FIRST_MILE](#first_mile)
8. [FOLLOW_PRIZE](#follow_prize)
9. [GLOBAL_PRODUCT](#global_product)
10. [LIVESTREAM](#livestream)
11. [LOGISTICS](#logistics)
12. [MEDIA](#media)
13. [MEDIA_SPACE](#media_space)
14. [MERCHANT](#merchant)
15. [ORDER](#order)
16. [PAYMENT](#payment)
17. [PRODUCT](#product)
18. [PUBLIC](#public)
19. [PUSH](#push)
20. [RETURNS](#returns)
21. [SBS](#sbs)
22. [SHOP](#shop)
23. [SHOP_CATEGORY](#shop_category)
24. [SHOP_FLASH_SALE](#shop_flash_sale)
25. [TOP_PICKS](#top_picks)
26. [V2](#v2)
27. [VOUCHER](#voucher)

---

## ACCOUNT_HEALTH

**API 개수**: 6개

**설명**: 계정 건강도 및 성과 모니터링

| API 이름 | 설명 |
|---------|------|
| `v2.account_health.get_late_orders` | Get Late Orders |
| `v2.account_health.get_listings_with_issues` | Get Listings With Issues |
| `v2.account_health.get_metric_source_detail` | Get Metric Source Detail |
| `v2.account_health.get_penalty_point_history` | Get Penalty Point History |
| `v2.account_health.get_punishment_history` | Get Punishment History |
| `v2.account_health.get_shop_performance` | Get Shop Performance |

### 주요 API 상세

#### `v2.account_health.get_late_orders`

**포함 파일**: 9개


#### `v2.account_health.get_listings_with_issues`

**포함 파일**: 9개


#### `v2.account_health.get_metric_source_detail`

**포함 파일**: 11개


---

## ADD_ON_DEAL

**API 개수**: 14개

**설명**: 추가 상품 딜 관리

| API 이름 | 설명 |
|---------|------|
| `v2.add_on_deal.add_add_on_deal` | Add Add On Deal |
| `v2.add_on_deal.add_add_on_deal_main_item` | Add Add On Deal Main Item |
| `v2.add_on_deal.add_add_on_deal_sub_item` | Add Add On Deal Sub Item |
| `v2.add_on_deal.delete_add_on_deal` | Delete Add On Deal |
| `v2.add_on_deal.delete_add_on_deal_main_item` | Delete Add On Deal Main Item |
| `v2.add_on_deal.delete_add_on_deal_sub_item` | Delete Add On Deal Sub Item |
| `v2.add_on_deal.end_add_on_deal` | End Add On Deal |
| `v2.add_on_deal.get_add_on_deal` | Get Add On Deal |
| `v2.add_on_deal.get_add_on_deal_list` | Get Add On Deal List |
| `v2.add_on_deal.get_add_on_deal_main_item` | Get Add On Deal Main Item |
| `v2.add_on_deal.get_add_on_deal_sub_item` | Get Add On Deal Sub Item |
| `v2.add_on_deal.update_add_on_deal` | Update Add On Deal |
| `v2.add_on_deal.update_add_on_deal_main_item` | Update Add On Deal Main Item |
| `v2.add_on_deal.update_add_on_deal_sub_item` | Update Add On Deal Sub Item |

### 주요 API 상세

#### `v2.add_on_deal.add_add_on_deal`

**포함 파일**: 9개


#### `v2.add_on_deal.add_add_on_deal_main_item`

**포함 파일**: 9개


#### `v2.add_on_deal.add_add_on_deal_sub_item`

**포함 파일**: 9개


---

## ADS

**API 개수**: 26개

**설명**: 광고 캠페인 관리

| API 이름 | 설명 |
|---------|------|
| `v2.ads.check_create_gms_product_campaign_eligibility` | Check Create Gms Product Campaign Eligibility |
| `v2.ads.create_auto_product_ads` | Create Auto Product Ads |
| `v2.ads.create_gms_product_campaign` | Create Gms Product Campaign |
| `v2.ads.create_manual_product_ads` | Create Manual Product Ads |
| `v2.ads.edit_auto_product_ads` | Edit Auto Product Ads |
| `v2.ads.edit_gms_item_product_campaign` | Edit Gms Item Product Campaign |
| `v2.ads.edit_gms_product_campaign` | Edit Gms Product Campaign |
| `v2.ads.edit_manual_product_ad_keywords` | Edit Manual Product Ad Keywords |
| `v2.ads.edit_manual_product_ads` | Edit Manual Product Ads |
| `v2.ads.get_ads_fácil_shop_rate` | Get Ads Fácil Shop Rate |
| `v2.ads.get_ads_f찼cil_shop_rate` | Get Ads F찼Cil Shop Rate |
| `v2.ads.get_all_cpc_ads_daily_performance` | Get All Cpc Ads Daily Performance |
| `v2.ads.get_all_cpc_ads_hourly_performance` | Get All Cpc Ads Hourly Performance |
| `v2.ads.get_create_product_ad_budget_suggestion` | Get Create Product Ad Budget Suggestion |
| `v2.ads.get_gms_campaign_performance` | Get Gms Campaign Performance |
| `v2.ads.get_gms_item_performance` | Get Gms Item Performance |
| `v2.ads.get_product_campaign_daily_performance` | Get Product Campaign Daily Performance |
| `v2.ads.get_product_campaign_hourly_performance` | Get Product Campaign Hourly Performance |
| `v2.ads.get_product_level_campaign_id_list` | Get Product Level Campaign Id List |
| `v2.ads.get_product_level_campaign_setting_info` | Get Product Level Campaign Setting Info |
| `v2.ads.get_product_recommended_roi_target` | Get Product Recommended Roi Target |
| `v2.ads.get_recommended_item_list` | Get Recommended Item List |
| `v2.ads.get_recommended_keyword_list` | Get Recommended Keyword List |
| `v2.ads.get_shop_toggle_info` | Get Shop Toggle Info |
| `v2.ads.get_total_balance` | Get Total Balance |
| `v2.ads.list_gms_user_deleted_item` | List Gms User Deleted Item |

### 주요 API 상세

#### `v2.ads.check_create_gms_product_campaign_eligibility`

**포함 파일**: 9개


#### `v2.ads.create_auto_product_ads`

**포함 파일**: 9개


#### `v2.ads.create_gms_product_campaign`

**포함 파일**: 9개


---

## BUNDLE_DEAL

**API 개수**: 10개

**설명**: 번들 상품 딜 관리

| API 이름 | 설명 |
|---------|------|
| `v2.bundle_deal.add_bundle_deal` | Add Bundle Deal |
| `v2.bundle_deal.add_bundle_deal_item` | Add Bundle Deal Item |
| `v2.bundle_deal.delete_bundle_deal` | Delete Bundle Deal |
| `v2.bundle_deal.delete_bundle_deal_item` | Delete Bundle Deal Item |
| `v2.bundle_deal.end_bundle_deal` | End Bundle Deal |
| `v2.bundle_deal.get_bundle_deal` | Get Bundle Deal |
| `v2.bundle_deal.get_bundle_deal_item` | Get Bundle Deal Item |
| `v2.bundle_deal.get_bundle_deal_list` | Get Bundle Deal List |
| `v2.bundle_deal.update_bundle_deal` | Update Bundle Deal |
| `v2.bundle_deal.update_bundle_deal_item` | Update Bundle Deal Item |

### 주요 API 상세

#### `v2.bundle_deal.add_bundle_deal`

**포함 파일**: 9개


#### `v2.bundle_deal.add_bundle_deal_item`

**포함 파일**: 9개


#### `v2.bundle_deal.delete_bundle_deal`

**포함 파일**: 9개


---

## DISCOUNT

**API 개수**: 9개

**설명**: 할인 프로모션 관리

| API 이름 | 설명 |
|---------|------|
| `v2.discount.add_discount` | Add Discount |
| `v2.discount.add_discount_item` | Add Discount Item |
| `v2.discount.delete_discount` | Delete Discount |
| `v2.discount.delete_discount_item` | Delete Discount Item |
| `v2.discount.end_discount` | End Discount |
| `v2.discount.get_discount` | Get Discount |
| `v2.discount.get_discount_list` | Get Discount List |
| `v2.discount.update_discount` | Update Discount |
| `v2.discount.update_discount_item` | Update Discount Item |

### 주요 API 상세

#### `v2.discount.add_discount`

**포함 파일**: 9개


#### `v2.discount.add_discount_item`

**포함 파일**: 9개


#### `v2.discount.delete_discount`

**포함 파일**: 9개


---

## FBS

**API 개수**: 4개

**설명**: FBS (Fulfilled By Shopee) 관련

| API 이름 | 설명 |
|---------|------|
| `v2.fbs.query_br_shop_block_status` | Query Br Shop Block Status |
| `v2.fbs.query_br_shop_enrollment_status` | Query Br Shop Enrollment Status |
| `v2.fbs.query_br_shop_invoice_error` | Query Br Shop Invoice Error |
| `v2.fbs.query_br_sku_block_status` | Query Br Sku Block Status |

### 주요 API 상세

#### `v2.fbs.query_br_shop_block_status`

**포함 파일**: 9개


#### `v2.fbs.query_br_shop_enrollment_status`

**포함 파일**: 9개


#### `v2.fbs.query_br_shop_invoice_error`

**포함 파일**: 9개


---

## FIRST_MILE

**API 개수**: 16개

**설명**: 퍼스트 마일 배송 관리

| API 이름 | 설명 |
|---------|------|
| `v2.first_mile.bind_courier_delivery_first_mile_tracking_number` | Bind Courier Delivery First Mile Tracking Number |
| `v2.first_mile.bind_first_mile_tracking_number` | Bind First Mile Tracking Number |
| `v2.first_mile.generate_and_bind_first_mile_tracking_number` | Generate And Bind First Mile Tracking Number |
| `v2.first_mile.generate_first_mile_tracking_number` | Generate First Mile Tracking Number |
| `v2.first_mile.get_channel_list` | Get Channel List |
| `v2.first_mile.get_courier_delivery_channel_list` | Get Courier Delivery Channel List |
| `v2.first_mile.get_courier_delivery_detail` | Get Courier Delivery Detail |
| `v2.first_mile.get_courier_delivery_tracking_number_list` | Get Courier Delivery Tracking Number List |
| `v2.first_mile.get_courier_delivery_waybill` | Get Courier Delivery Waybill |
| `v2.first_mile.get_detail` | Get Detail |
| `v2.first_mile.get_tracking_number_list` | Get Tracking Number List |
| `v2.first_mile.get_transit_warehouse_list` | Get Transit Warehouse List |
| `v2.first_mile.get_unbind_order_list` | Get Unbind Order List |
| `v2.first_mile.get_waybill` | Get Waybill |
| `v2.first_mile.unbind_first_mile_tracking_number` | Unbind First Mile Tracking Number |
| `v2.first_mile.unbind_first_mile_tracking_number_all` | Unbind First Mile Tracking Number All |

### 주요 API 상세

#### `v2.first_mile.bind_courier_delivery_first_mile_tracking_number`

**포함 파일**: 9개


#### `v2.first_mile.bind_first_mile_tracking_number`

**포함 파일**: 9개


#### `v2.first_mile.generate_and_bind_first_mile_tracking_number`

**포함 파일**: 9개


---

## FOLLOW_PRIZE

**API 개수**: 6개

**설명**: 팔로우 경품 이벤트

| API 이름 | 설명 |
|---------|------|
| `v2.follow_prize.add_follow_prize` | Add Follow Prize |
| `v2.follow_prize.delete_follow_prize` | Delete Follow Prize |
| `v2.follow_prize.end_follow_prize` | End Follow Prize |
| `v2.follow_prize.get_follow_prize_detail` | Get Follow Prize Detail |
| `v2.follow_prize.get_follow_prize_list` | Get Follow Prize List |
| `v2.follow_prize.update_follow_prize` | Update Follow Prize |

### 주요 API 상세

#### `v2.follow_prize.add_follow_prize`

**포함 파일**: 9개


#### `v2.follow_prize.delete_follow_prize`

**포함 파일**: 9개


#### `v2.follow_prize.end_follow_prize`

**포함 파일**: 9개


---

## GLOBAL_PRODUCT

**API 개수**: 34개

**설명**: 글로벌 상품 관리

| API 이름 | 설명 |
|---------|------|
| `v2.global_product.add_global_item` | Add Global Item |
| `v2.global_product.add_global_model` | Add Global Model |
| `v2.global_product.category_recommend` | Category Recommend |
| `v2.global_product.create_publish_task` | Create Publish Task |
| `v2.global_product.delete_global_item` | Delete Global Item |
| `v2.global_product.delete_global_model` | Delete Global Model |
| `v2.global_product.get_attribute_tree` | Get Attribute Tree |
| `v2.global_product.get_brand_list` | Get Brand List |
| `v2.global_product.get_category` | Get Category |
| `v2.global_product.get_global_item_id` | Get Global Item Id |
| `v2.global_product.get_global_item_info` | Get Global Item Info |
| `v2.global_product.get_global_item_limit` | Get Global Item Limit |
| `v2.global_product.get_global_item_list` | Get Global Item List |
| `v2.global_product.get_global_model_list` | Get Global Model List |
| `v2.global_product.get_local_adjustment_rate` | Get Local Adjustment Rate |
| `v2.global_product.get_publish_task_result` | Get Publish Task Result |
| `v2.global_product.get_publishable_shop` | Get Publishable Shop |
| `v2.global_product.get_published_list` | Get Published List |
| `v2.global_product.get_recommend_attribute` | Get Recommend Attribute |
| `v2.global_product.get_shop_publishable_status` | Get Shop Publishable Status |
| `v2.global_product.get_size_chart_detail` | Get Size Chart Detail |
| `v2.global_product.get_size_chart_list` | Get Size Chart List |
| `v2.global_product.get_variations` | Get Variations |
| `v2.global_product.init_tier_variation` | Init Tier Variation |
| `v2.global_product.search_global_attribute_value_list` | Search Global Attribute Value List |
| `v2.global_product.set_sync_field` | Set Sync Field |
| `v2.global_product.support_size_chart` | Support Size Chart |
| `v2.global_product.update_global_item` | Update Global Item |
| `v2.global_product.update_global_model` | Update Global Model |
| `v2.global_product.update_local_adjustment_rate` | Update Local Adjustment Rate |
| `v2.global_product.update_price` | Update Price |
| `v2.global_product.update_size_chart` | Update Size Chart |
| `v2.global_product.update_stock` | Update Stock |
| `v2.global_product.update_tier_variation` | Update Tier Variation |

### 주요 API 상세

#### `v2.global_product.add_global_item`

**포함 파일**: 11개


#### `v2.global_product.add_global_model`

**포함 파일**: 9개


#### `v2.global_product.category_recommend`

**포함 파일**: 9개


---

## LIVESTREAM

**API 개수**: 25개

**설명**: 라이브 스트리밍 관리

| API 이름 | 설명 |
|---------|------|
| `v2.livestream.add_item_list` | Add Item List |
| `v2.livestream.apply_item_set` | Apply Item Set |
| `v2.livestream.ban_user_comment` | Ban User Comment |
| `v2.livestream.create_session` | Create Session |
| `v2.livestream.delete_item_list` | Delete Item List |
| `v2.livestream.delete_show_item` | Delete Show Item |
| `v2.livestream.end_session` | End Session |
| `v2.livestream.get_item_count` | Get Item Count |
| `v2.livestream.get_item_list` | Get Item List |
| `v2.livestream.get_item_set_item_list` | Get Item Set Item List |
| `v2.livestream.get_item_set_list` | Get Item Set List |
| `v2.livestream.get_latest_comment_list` | Get Latest Comment List |
| `v2.livestream.get_like_item_list` | Get Like Item List |
| `v2.livestream.get_recent_item_list` | Get Recent Item List |
| `v2.livestream.get_session_detail` | Get Session Detail |
| `v2.livestream.get_session_item_metric` | Get Session Item Metric |
| `v2.livestream.get_session_metric` | Get Session Metric |
| `v2.livestream.get_show_item` | Get Show Item |
| `v2.livestream.post_comment` | Post Comment |
| `v2.livestream.start_session` | Start Session |
| `v2.livestream.unban_user_comment` | Unban User Comment |
| `v2.livestream.update_item_list` | Update Item List |
| `v2.livestream.update_session` | Update Session |
| `v2.livestream.update_show_item` | Update Show Item |
| `v2.livestream.upload_image` | Upload Image |

### 주요 API 상세

#### `v2.livestream.add_item_list`

**포함 파일**: 9개


#### `v2.livestream.apply_item_set`

**포함 파일**: 9개


#### `v2.livestream.ban_user_comment`

**포함 파일**: 9개


---

## LOGISTICS

**API 개수**: 41개

**설명**: 물류 및 배송 관리

| API 이름 | 설명 |
|---------|------|
| `v2.logistics.batch_ship_order` | Batch Ship Order |
| `v2.logistics.batch_update_tpf_warehouse_tracking_status` | Batch Update Tpf Warehouse Tracking Status |
| `v2.logistics.create_booking_shipping_document` | Create Booking Shipping Document |
| `v2.logistics.create_shipping_document` | Create Shipping Document |
| `v2.logistics.create_shipping_document_job` | Create Shipping Document Job |
| `v2.logistics.delete_address` | Delete Address |
| `v2.logistics.delete_special_operating_hour` | Delete Special Operating Hour |
| `v2.logistics.download_booking_shipping_document` | Download Booking Shipping Document |
| `v2.logistics.download_shipping_document` | Download Shipping Document |
| `v2.logistics.download_shipping_document_job` | Download Shipping Document Job |
| `v2.logistics.download_to_label` | Download To Label |
| `v2.logistics.get_address_list` | Get Address List |
| `v2.logistics.get_booking_shipping_document_data_info` | Get Booking Shipping Document Data Info |
| `v2.logistics.get_booking_shipping_document_parameter` | Get Booking Shipping Document Parameter |
| `v2.logistics.get_booking_shipping_document_result` | Get Booking Shipping Document Result |
| `v2.logistics.get_booking_shipping_parameter` | Get Booking Shipping Parameter |
| `v2.logistics.get_booking_tracking_info` | Get Booking Tracking Info |
| `v2.logistics.get_booking_tracking_number` | Get Booking Tracking Number |
| `v2.logistics.get_channel_list` | Get Channel List |
| `v2.logistics.get_mart_packaging_info` | Get Mart Packaging Info |
| `v2.logistics.get_mass_shipping_parameter` | Get Mass Shipping Parameter |
| `v2.logistics.get_mass_tracking_number` | Get Mass Tracking Number |
| `v2.logistics.get_operating_hour_restrictions` | Get Operating Hour Restrictions |
| `v2.logistics.get_operating_hours` | Get Operating Hours |
| `v2.logistics.get_shipping_document_data_info` | Get Shipping Document Data Info |
| `v2.logistics.get_shipping_document_job_status` | Get Shipping Document Job Status |
| `v2.logistics.get_shipping_document_parameter` | Get Shipping Document Parameter |
| `v2.logistics.get_shipping_document_result` | Get Shipping Document Result |
| `v2.logistics.get_shipping_parameter` | Get Shipping Parameter |
| `v2.logistics.get_tracking_info` | Get Tracking Info |
| `v2.logistics.get_tracking_number` | Get Tracking Number |
| `v2.logistics.mass_ship_order` | Mass Ship Order |
| `v2.logistics.set_address_config` | Set Address Config |
| `v2.logistics.set_mart_packaging_info` | Set Mart Packaging Info |
| `v2.logistics.ship_booking` | Ship Booking |
| `v2.logistics.ship_order` | Ship Order |
| `v2.logistics.update_channel` | Update Channel |
| `v2.logistics.update_operating_hours` | Update Operating Hours |
| `v2.logistics.update_self_collection_order_logistics` | Update Self Collection Order Logistics |
| `v2.logistics.update_shipping_order` | Update Shipping Order |
| `v2.logistics.update_tracking_status` | Update Tracking Status |

### 주요 API 상세

#### `v2.logistics.batch_ship_order`

**포함 파일**: 11개


#### `v2.logistics.batch_update_tpf_warehouse_tracking_status`

**포함 파일**: 9개


#### `v2.logistics.create_booking_shipping_document`

**포함 파일**: 9개


---

## MEDIA

**API 개수**: 1개

**설명**: 미디어 업로드

| API 이름 | 설명 |
|---------|------|
| `v2.media.upload_image` | Upload Image |

### 주요 API 상세

#### `v2.media.upload_image`

**포함 파일**: 9개


---

## MEDIA_SPACE

**API 개수**: 6개

**설명**: 미디어 스페이스 관리

| API 이름 | 설명 |
|---------|------|
| `v2.media_space.cancel_video_upload` | Cancel Video Upload |
| `v2.media_space.complete_video_upload` | Complete Video Upload |
| `v2.media_space.get_video_upload_result` | Get Video Upload Result |
| `v2.media_space.init_video_upload` | Init Video Upload |
| `v2.media_space.upload_image` | Upload Image |
| `v2.media_space.upload_video_part` | Upload Video Part |

### 주요 API 상세

#### `v2.media_space.cancel_video_upload`

**포함 파일**: 9개


#### `v2.media_space.complete_video_upload`

**포함 파일**: 9개


#### `v2.media_space.get_video_upload_result`

**포함 파일**: 9개


---

## MERCHANT

**API 개수**: 6개

**설명**: 판매자 정보 관리

| API 이름 | 설명 |
|---------|------|
| `v2.merchant.get_merchant_info` | Get Merchant Info |
| `v2.merchant.get_merchant_prepaid_account_list` | Get Merchant Prepaid Account List |
| `v2.merchant.get_merchant_warehouse_list` | Get Merchant Warehouse List |
| `v2.merchant.get_merchant_warehouse_location_list` | Get Merchant Warehouse Location List |
| `v2.merchant.get_shop_list_by_merchant` | Get Shop List By Merchant |
| `v2.merchant.get_warehouse_eligible_shop_list` | Get Warehouse Eligible Shop List |

### 주요 API 상세

#### `v2.merchant.get_merchant_info`

**포함 파일**: 9개


#### `v2.merchant.get_merchant_prepaid_account_list`

**포함 파일**: 9개


#### `v2.merchant.get_merchant_warehouse_list`

**포함 파일**: 9개


---

## ORDER

**API 개수**: 21개

**설명**: 주문 관리

| API 이름 | 설명 |
|---------|------|
| `v2.order.cancel_order` | Cancel Order |
| `v2.order.download_fbs_invoices` | Download Fbs Invoices |
| `v2.order.download_invoice_doc` | Download Invoice Doc |
| `v2.order.generate_fbs_invoices` | Generate Fbs Invoices |
| `v2.order.get_booking_detail` | Get Booking Detail |
| `v2.order.get_booking_list` | Get Booking List |
| `v2.order.get_buyer_invoice_info` | Get Buyer Invoice Info |
| `v2.order.get_fbs_invoices_result` | Get Fbs Invoices Result |
| `v2.order.get_order_detail` | Get Order Detail |
| `v2.order.get_order_list` | Get Order List |
| `v2.order.get_package_detail` | Get Package Detail |
| `v2.order.get_pending_buyer_invoice_order_list` | Get Pending Buyer Invoice Order List |
| `v2.order.get_shipment_list` | Get Shipment List |
| `v2.order.get_warehouse_filter_config` | Get Warehouse Filter Config |
| `v2.order.handle_buyer_cancellation` | Handle Buyer Cancellation |
| `v2.order.handle_prescription_check` | Handle Prescription Check |
| `v2.order.search_package_list` | Search Package List |
| `v2.order.set_note` | Set Note |
| `v2.order.split_order` | Split Order |
| `v2.order.unsplit_order` | Unsplit Order |
| `v2.order.upload_invoice_doc` | Upload Invoice Doc |

### 주요 API 상세

#### `v2.order.cancel_order`

**포함 파일**: 9개


#### `v2.order.download_fbs_invoices`

**포함 파일**: 9개


#### `v2.order.download_invoice_doc`

**포함 파일**: 9개


---

## PAYMENT

**API 개수**: 16개

**설명**: 결제 및 정산 관리

| API 이름 | 설명 |
|---------|------|
| `v2.payment.generate_income_report` | Generate Income Report |
| `v2.payment.generate_income_statement` | Generate Income Statement |
| `v2.payment.get_billing_transaction_info` | Get Billing Transaction Info |
| `v2.payment.get_escrow_detail` | Get Escrow Detail |
| `v2.payment.get_escrow_detail_batch` | Get Escrow Detail Batch |
| `v2.payment.get_escrow_list` | Get Escrow List |
| `v2.payment.get_income_report` | Get Income Report |
| `v2.payment.get_income_statement` | Get Income Statement |
| `v2.payment.get_item_installment_status` | Get Item Installment Status |
| `v2.payment.get_payment_method_list` | Get Payment Method List |
| `v2.payment.get_payout_detail` | Get Payout Detail |
| `v2.payment.get_payout_info` | Get Payout Info |
| `v2.payment.get_shop_installment_status` | Get Shop Installment Status |
| `v2.payment.get_wallet_transaction_list` | Get Wallet Transaction List |
| `v2.payment.set_item_installment_status` | Set Item Installment Status |
| `v2.payment.set_shop_installment_status` | Set Shop Installment Status |

### 주요 API 상세

#### `v2.payment.generate_income_report`

**포함 파일**: 9개


#### `v2.payment.generate_income_statement`

**포함 파일**: 9개


#### `v2.payment.get_billing_transaction_info`

**포함 파일**: 9개


---

## PRODUCT

**API 개수**: 55개

**설명**: 상품 관리

| API 이름 | 설명 |
|---------|------|
| `v2.product.add_item` | Add Item |
| `v2.product.add_kit_item` | Add Kit Item |
| `v2.product.add_model` | Add Model |
| `v2.product.add_ssp_item` | Add Ssp Item |
| `v2.product.boost_item` | Boost Item |
| `v2.product.category_recommend` | Category Recommend |
| `v2.product.delete_item` | Delete Item |
| `v2.product.delete_model` | Delete Model |
| `v2.product.generate_kit_image` | Generate Kit Image |
| `v2.product.get_aitem_by_pitem_id` | Get Aitem By Pitem Id |
| `v2.product.get_all_vehicle_list` | Get All Vehicle List |
| `v2.product.get_attribute_tree` | V2.Product.Get Attribute Tree |
| `v2.product.get_boosted_list` | Get Boosted List |
| `v2.product.get_brand_list` | Get Brand List |
| `v2.product.get_category` | V2.Product.Get Category |
| `v2.product.get_comment` | Get Comment |
| `v2.product.get_direct_item_list` | Get Direct Item List |
| `v2.product.get_direct_shop_recommended_price` | Get Direct Shop Recommended Price |
| `v2.product.get_item_base_info` | Get Item Base Info |
| `v2.product.get_item_content_diagnosis_result` | Get Item Content Diagnosis Result |
| `v2.product.get_item_extra_info` | Get Item Extra Info |
| `v2.product.get_item_limit` | Get Item Limit |
| `v2.product.get_item_list` | Get Item List |
| `v2.product.get_item_list_by_content_diagnosis` | Get Item List By Content Diagnosis |
| `v2.product.get_item_promotion` | Get Item Promotion |
| `v2.product.get_item_violation_info` | Get Item Violation Info |
| `v2.product.get_kit_item_info` | Get Kit Item Info |
| `v2.product.get_kit_item_limit` | Get Kit Item Limit |
| `v2.product.get_main_item_list` | Get Main Item List |
| `v2.product.get_model_list` | Get Model List |
| `v2.product.get_product_certification_rule` | Get Product Certification Rule |
| `v2.product.get_recommend_attribute` | Get Recommend Attribute |
| `v2.product.get_size_chart_detail` | Get Size Chart Detail |
| `v2.product.get_size_chart_list` | Get Size Chart List |
| `v2.product.get_ssp_info` | Get Ssp Info |
| `v2.product.get_ssp_list` | Get Ssp List |
| `v2.product.get_variations` | Get Variations |
| `v2.product.get_vehicle_list_by_compatibility_detail` | Get Vehicle List By Compatibility Detail |
| `v2.product.get_weight_recommendation` | Get Weight Recommendation |
| `v2.product.init_tier_variation` | Init Tier Variation |
| `v2.product.link_ssp` | Link Ssp |
| `v2.product.register_brand` | Register Brand |
| `v2.product.reply_comment` | Reply Comment |
| `v2.product.search_attribute_value_list` | Search Attribute Value List |
| `v2.product.search_item` | Search Item |
| `v2.product.search_unpackaged_model_list` | Search Unpackaged Model List |
| `v2.product.unlink_ssp` | Unlink Ssp |
| `v2.product.unlist_item` | Unlist Item |
| `v2.product.update_item` | Update Item |
| `v2.product.update_kit_item` | Update Kit Item |
| `v2.product.update_model` | Update Model |
| `v2.product.update_price` | Update Price |
| `v2.product.update_sip_item_price` | Update Sip Item Price |
| `v2.product.update_stock` | Update Stock |
| `v2.product.update_tier_variation` | Update Tier Variation |

### 주요 API 상세

#### `v2.product.add_item`

**포함 파일**: 15개


#### `v2.product.add_kit_item`

**포함 파일**: 9개


#### `v2.product.add_model`

**포함 파일**: 11개


---

## PUBLIC

**API 개수**: 6개

**설명**: 공개 API (인증 등)

| API 이름 | 설명 |
|---------|------|
| `v2.public.get_access_token` | Get Access Token |
| `v2.public.get_merchants_by_partner` | Get Merchants By Partner |
| `v2.public.get_shopee_ip_ranges` | Get Shopee Ip Ranges |
| `v2.public.get_shops_by_partner` | Get Shops By Partner |
| `v2.public.get_token_by_resend_code` | Get Token By Resend Code |
| `v2.public.refresh_access_token` | Refresh Access Token |

### 주요 API 상세

#### `v2.public.get_access_token`

**포함 파일**: 9개


#### `v2.public.get_merchants_by_partner`

**포함 파일**: 9개


#### `v2.public.get_shopee_ip_ranges`

**포함 파일**: 9개


---

## PUSH

**API 개수**: 4개

**설명**: 푸시 알림 설정

| API 이름 | 설명 |
|---------|------|
| `v2.push.confirm_consumed_lost_push_message` | Confirm Consumed Lost Push Message |
| `v2.push.get_app_push_config` | Get App Push Config |
| `v2.push.get_lost_push_message` | Get Lost Push Message |
| `v2.push.set_app_push_config` | Set App Push Config |

### 주요 API 상세

#### `v2.push.confirm_consumed_lost_push_message`

**포함 파일**: 9개


#### `v2.push.get_app_push_config`

**포함 파일**: 9개


#### `v2.push.get_lost_push_message`

**포함 파일**: 9개


---

## RETURNS

**API 개수**: 15개

**설명**: 반품 및 환불 관리

| API 이름 | 설명 |
|---------|------|
| `v2.returns.accept_offer` | Accept Offer |
| `v2.returns.cancel_dispute` | Cancel Dispute |
| `v2.returns.confirm` | Confirm |
| `v2.returns.convert_image` | Convert Image |
| `v2.returns.dispute` | Dispute |
| `v2.returns.get_available_solutions` | Get Available Solutions |
| `v2.returns.get_return_detail` | Get Return Detail |
| `v2.returns.get_return_dispute_reason` | Get Return Dispute Reason |
| `v2.returns.get_return_list` | Get Return List |
| `v2.returns.get_reverse_tracking_info` | Get Reverse Tracking Info |
| `v2.returns.get_shipping_carrier` | Get Shipping Carrier |
| `v2.returns.offer` | Offer |
| `v2.returns.query_proof` | Query Proof |
| `v2.returns.upload_proof` | Upload Proof |
| `v2.returns.upload_shipping_proof` | Upload Shipping Proof |

### 주요 API 상세

#### `v2.returns.accept_offer`

**포함 파일**: 9개


#### `v2.returns.cancel_dispute`

**포함 파일**: 9개


#### `v2.returns.confirm`

**포함 파일**: 9개


---

## SBS

**API 개수**: 5개

**설명**: SBS (Shopee Business Solutions)

| API 이름 | 설명 |
|---------|------|
| `v2.sbs.get_bound_whs_info` | Get Bound Whs Info |
| `v2.sbs.get_current_inventory` | Get Current Inventory |
| `v2.sbs.get_expiry_report` | Get Expiry Report |
| `v2.sbs.get_stock_aging` | Get Stock Aging |
| `v2.sbs.get_stock_movement` | Get Stock Movement |

### 주요 API 상세

#### `v2.sbs.get_bound_whs_info`

**포함 파일**: 9개


#### `v2.sbs.get_current_inventory`

**포함 파일**: 9개


#### `v2.sbs.get_expiry_report`

**포함 파일**: 9개


---

## SHOP

**API 개수**: 6개

**설명**: 샵 정보 관리

| API 이름 | 설명 |
|---------|------|
| `v2.shop.get_authorised_reseller_brand` | Get Authorised Reseller Brand |
| `v2.shop.get_profile` | Get Profile |
| `v2.shop.get_shop_info` | Get Shop Info |
| `v2.shop.get_shop_notification` | Get Shop Notification |
| `v2.shop.get_warehouse_detail` | Get Warehouse Detail |
| `v2.shop.update_profile` | Update Profile |

### 주요 API 상세

#### `v2.shop.get_authorised_reseller_brand`

**포함 파일**: 9개


#### `v2.shop.get_profile`

**포함 파일**: 9개


#### `v2.shop.get_shop_info`

**포함 파일**: 9개


---

## SHOP_CATEGORY

**API 개수**: 7개

**설명**: 샵 카테고리 관리

| API 이름 | 설명 |
|---------|------|
| `v2.shop_category.add_item_list` | Add Item List |
| `v2.shop_category.add_shop_category` | Add Shop Category |
| `v2.shop_category.delete_item_list` | Delete Item List |
| `v2.shop_category.delete_shop_category` | Delete Shop Category |
| `v2.shop_category.get_item_list` | Get Item List |
| `v2.shop_category.get_shop_category_list` | Get Shop Category List |
| `v2.shop_category.update_shop_category` | Update Shop Category |

### 주요 API 상세

#### `v2.shop_category.add_item_list`

**포함 파일**: 9개


#### `v2.shop_category.add_shop_category`

**포함 파일**: 9개


#### `v2.shop_category.delete_item_list`

**포함 파일**: 9개


---

## SHOP_FLASH_SALE

**API 개수**: 11개

**설명**: 샵 플래시 세일

| API 이름 | 설명 |
|---------|------|
| `v2.shop_flash_sale.add_shop_flash_sale_items` | Add Shop Flash Sale Items |
| `v2.shop_flash_sale.create_shop_flash_sale` | Create Shop Flash Sale |
| `v2.shop_flash_sale.delete_shop_flash_sale` | Delete Shop Flash Sale |
| `v2.shop_flash_sale.delete_shop_flash_sale_items` | Delete Shop Flash Sale Items |
| `v2.shop_flash_sale.get_item_criteria` | Get Item Criteria |
| `v2.shop_flash_sale.get_shop_flash_sale` | Get Shop Flash Sale |
| `v2.shop_flash_sale.get_shop_flash_sale_items` | Get Shop Flash Sale Items |
| `v2.shop_flash_sale.get_shop_flash_sale_list` | Get Shop Flash Sale List |
| `v2.shop_flash_sale.get_time_slot_id` | Get Time Slot Id |
| `v2.shop_flash_sale.update_shop_flash_sale` | Update Shop Flash Sale |
| `v2.shop_flash_sale.update_shop_flash_sale_items` | Update Shop Flash Sale Items |

### 주요 API 상세

#### `v2.shop_flash_sale.add_shop_flash_sale_items`

**포함 파일**: 9개


#### `v2.shop_flash_sale.create_shop_flash_sale`

**포함 파일**: 9개


#### `v2.shop_flash_sale.delete_shop_flash_sale`

**포함 파일**: 9개


---

## TOP_PICKS

**API 개수**: 4개

**설명**: 추천 상품 관리

| API 이름 | 설명 |
|---------|------|
| `v2.top_picks.add_top_picks` | Add Top Picks |
| `v2.top_picks.delete_top_picks` | Delete Top Picks |
| `v2.top_picks.get_top_picks_list` | Get Top Picks List |
| `v2.top_picks.update_top_picks` | Update Top Picks |

### 주요 API 상세

#### `v2.top_picks.add_top_picks`

**포함 파일**: 9개


#### `v2.top_picks.delete_top_picks`

**포함 파일**: 9개


#### `v2.top_picks.get_top_picks_list`

**포함 파일**: 9개


---

## V2

**API 개수**: 2개

| API 이름 | 설명 |
|---------|------|
| `v2.v2.product.get_attribute_tree` | Get Attribute Tree |
| `v2.v2.product.get_category` | Get Category |

### 주요 API 상세

#### `v2.v2.product.get_attribute_tree`

**포함 파일**: 5개


#### `v2.v2.product.get_category`

**포함 파일**: 5개


---

## VOUCHER

**API 개수**: 6개

**설명**: 쿠폰/바우처 관리

| API 이름 | 설명 |
|---------|------|
| `v2.voucher.add_voucher` | Add Voucher |
| `v2.voucher.delete_voucher` | Delete Voucher |
| `v2.voucher.end_voucher` | End Voucher |
| `v2.voucher.get_voucher` | Get Voucher |
| `v2.voucher.get_voucher_list` | Get Voucher List |
| `v2.voucher.update_voucher` | Update Voucher |

### 주요 API 상세

#### `v2.voucher.add_voucher`

**포함 파일**: 9개


#### `v2.voucher.delete_voucher`

**포함 파일**: 9개


#### `v2.voucher.end_voucher`

**포함 파일**: 9개


---

## 📖 API 사용 가이드

### 인증 및 권한

Shopee Open API를 사용하기 위해서는 다음이 필요합니다:

1. **Partner ID**: 파트너 식별자
2. **Partner Key**: API 인증 키
3. **Access Token**: 샵별 액세스 토큰
4. **Shop ID**: 대상 샵 식별자

### 공통 요청 형식

```
Base URL: https://partner.shopeemobile.com
API Version: v2
```

### 주요 API 카테고리별 사용 사례

#### PRODUCT

- 상품 등록 및 수정
- 상품 정보 조회
- 재고 관리
- 카테고리 및 속성 관리

#### ORDER

- 주문 목록 조회
- 주문 상세 정보 확인
- 주문 취소 처리
- 배송 준비 및 처리

#### LOGISTICS

- 배송 방법 설정
- 운송장 번호 관리
- 배송 문서 생성
- 배송 상태 추적

#### PAYMENT

- 결제 정보 조회
- 정산 내역 확인
- 수익 보고서 생성
- 지갑 거래 내역

#### DISCOUNT

- 할인 프로모션 생성
- 할인 상품 추가/삭제
- 프로모션 기간 관리
- 할인 현황 조회

---

## 📚 부록

### API 호출 제한

- 각 API마다 호출 제한이 다를 수 있습니다
- Rate Limiting 정보는 각 API 문서를 참조하세요

### 참고 자료

- [Shopee Open Platform 공식 문서](https://open.shopee.com/documents)
- [Developer Guide](https://open.shopee.com/developer-guide/4)

### 변경 이력

- 2025-10-16: 초기 문서 생성
