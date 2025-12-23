# API 문서 추출

**카테고리**: API 참조
**난이도**: medium
**중요도**: 4/5
**최종 업데이트**: 2025-10-16T08:32:06

## 개요

이 문서는 TikTok Shop API 내에서 사용되는 다양한 데이터 정의, 상태 코드, 이유 및 유형에 대한 포괄적인 개요를 제공합니다. 주문 상태, 반품 상태, 물류 상태, 반품 사유, 분쟁 사유, 속성 유형, 취소 사유, 수수료 유형 및 결제 방법에 대한 세부 정보를 포함합니다.

## 주요 키워드

- API
- 문서
- 데이터 정의
- 상태 코드
- 주문 상태
- 반품 상태
- 물류 상태
- 결제 방법
- 반품 사유
- 분쟁 사유

## 본문

# API 문서 추출

## 공지 콘솔

시작하기 > V2.0 데이터 정의

- **uint32**: 32비트 부호 없는 정수
- **uint64**: 64비트 부호 없는 정수
- **timestamp**: unit32
- **string**: UTF-8 코드. 기본 문자열은 2자입니다.

---

## OrderStatus

- **UNPAID**: 주문이 생성되었지만 구매자가 아직 결제하지 않았습니다.
- **READY_TO_SHIP**: 주문이 배송 준비되었습니다.
- **READY_TO_SHIP_PENDING**: 주문이 배송 준비되었으며 3PL로부터 추적 번호를 받기 기다리고 있습니다.
- **RETRY_SHIP**: 주문을 다시 보내야 합니다.
- **IN_TRANSIT**: 소포가 3PL에 전달되었거나 3PL에서 픽업했습니다.
- **SHIPPED** 또는 **DELIVERED**: 주문이 판매자에게 전달되었습니다.
- **COMPLETED**: 주문이 완료 및 종료되었습니다.
- **CANCELLED**: 주문이 취소되었습니다.
- **TO_RETURN**: 구매자가 주문 반품을 요청했으며 반품이 진행 중입니다.
- **LOST**: TikTok Ship 주문이 분실되었습니다.

---

## ReturnStatus

- **REQUESTED**
- **ACCEPTED**
- **IN_TRANSIT**
- **JUDGING**
- **CLOSED**
- **CANCELLED**
- **SELLER_DISPUTE**

---

## ReturnSolution

- **RETURN_REFUND**
- **REFUND_ONLY**

---

## ReturnReason

- **NONRECEIPT**
- **WRONG_ITEM**
- **COUNTERFEIT**
- **GIFT_DESC**
- **MATRL_NOISE**
- **SIZE_SHOES**
- **USED**
- **NO_REASON**
- **RECEIVED_DAMAGED**
- **UNWANTED**
- **ITEM_MISSING**
- **FUNCTIONAL_FAILED**
- **COLOR**
- **PHYSICAL_DMG**
- **FUNCTIONAL_DMG**
- **NOT_AS_DESCRIBED**
- **SURPCODE_PARCEL**
- **EXPIRES_PRODUCT**
- **LOST_IN_DELIVERY**
- **PARTIAL_MISSING**
- **CHANGE_OF_MIND**
- **INCOMPLETE_MISSING_ITEM**
- **ITEM_FAULT**
- **BROKEN_PRODUCTS**
- **DAMAGE_PACKAGE**
- **WRONG_DELIVERY**
- **SIZE_DEVIATION**
- **LOOK_FEEL**
- **POOR_QUALITY**
- **DIFFERENT_DESCRIPTION**

---

## LogisticsStatus

- **LOGISTICS_NOT_STARTED**: 초기 상태, 주문이 이행 준비되지 않았습니다.
- **LOGISTICS_REQUEST_CREATED**: 주문이 픽업 준비되었습니다.
- **LOGISTICS_COD_REJECTED**: 통합 물류 COD. 주문이 COD로 거부되었습니다.
- **LOGISTICS_PICKUP_NOT_DONE**: 주문이 픽업 준비되었습니다. non-COD. 주문 preCOD. 주문이 픽업 SLA를 통과했습니다. 3PL에서 픽업하지 않았습니다.
- **LOGISTICS_REQUEST_CREATED**: 주문이 배송 준비되었습니다.
- **LOGISTICS_PICKUP_DONE**: 주문이 3PL에서 픽업되었거나 추적 실패 또는 소포 분실 - 마지막 구간이 아닙니다.
- **LOGISTICS_IN_TRANSIT**: 주문이 중간 구간을 운송 중입니다.
- **LOGISTICS_ARRIVAL**: 주문 상태가 LOGISTICS_REACH로 계산되었습니다.
- **LOGISTICS_DELIVERY_CANCELLED**: 주문 배송이 취소되었습니다. LOGISTICS_REQUEST_CREATED 또는 LOGISTICS_PICKUP_DONE 상태입니다.
- **LOGISTICS_DELIVERY_FAILED**: TikTok Ship 배송 주문이 실패했습니다.
- **LOGISTICS_ROUTE_RETRY**: 주문이 SPI를 통과합니다. 픽업 재시도
- **LOGISTICS_DELIVERY_FAILED**: 주문이 마지막 구간 배송 실패를 통과했습니다.
- **LOGISTICS_DELIVERED**: TikTok Ship 배송 주문이 완료되었습니다.

---

## PackageFulfillmentStatus

- **LOGISTICS_NOT_START**: 초기 상태, 패키지가 이행 준비되지 않았습니다.
- **LOGISTICS_READY**: 패키지가 이행 준비되었습니다 (반품/제거 가능). non-COD, prdt의 경우 COD. 3PL에서 픽업 SLA를 통과했습니다.
- **LOGISTICS_PICKED**: 패키지가 3PL에서 픽업되었습니다.
- **LOGISTICS_PICUP_DONE**: 패키지가 3PL에 전달되었습니다.
- **LOGISTICS_CANCELLED**: 패키지가 성공적으로 배송되었습니다.
- **LOGISTICS_DELIVERY_FAIL**: TikTok Ship 주문 배송이 실패했습니다. LOGISTICS_READY
- **LOGISTICS_REQUEST_CANCELLED**: 패키지가 LOGISTICS_REQUEST_CREATED에 있을 때 주문이 취소되었습니다.
- **LOGISTICS_PICDUP_FAILED**: 패키지가 LOGISTICS_PICUP_DONE에 있을 때 주문이 취소되었습니다.
- **LOGISTICS_LOST**: 주문이 3PL에서 분실되었습니다.
- **LOGISTICS_DELIVERY_CANCELLED**: TikTok Ship 주문이 취소되었습니다.
- **LOGISTICS_DELIVERY_FAILED**: SPI로 인해 주문이 취소되었거나 배송 실패 또는 분실되었습니다.
- **LOGISTICS_LOST**: SPI로 인해 주문이 취소되었거나 패키지가 분실되었습니다.

---

## ReturnDisputeReasonId

- **"1":** "수령 거부 주장을 거부하고 싶습니다."
- **"2":** "주장에 동의하지 않습니다."
- **"3":** "상품이 배송되었다는 증거가 있습니다. (반품되어야 하는 제품이 수령되지 않은 경우)"
- **"6":** "商品拒收异常"
- **"8":** "买家拒收异常"
- **"9":** "包裹异常被拒收"
- **"9":** "기타"
- **"10":** "상품을 분실했고 배송 증거가 있습니다."
- **"11":** "반품된 상품이 주문한 상품과 일치하지 않습니다."
- **"12":** "상품이 양호하지만 사용된 상태이므로 주장에 동의하지 않습니다."
- **"13":** "상품 가격이 정확하다고 생각합니다."
- **"14":** "반품 상품에는 동의하지만 반품되어야 하는 상품/전체 반품을 받지 못했습니다."
- **"15":** "반품 상품에는 동의하지만 패키지/반품 상품이 손상/변조되었습니다."
- **"16":** "송장에 따라 정확한 상품을 배송했습니다."
- **"17":** "송장에 따라 정확한 상품을 배송했습니다."
- **"18":** "정상 작동 상태로 상품을 공급했습니다."
- **"4":** "제품이 적절한 기간에 있지 않습니다."
- **"5":** "제품이 적절한 기간에 있지 않습니다."
- **"19":** "분쟁 사유에 대해 이의를 제기합니다."
- **"20":** "해당 사항 없음" (개봉 사진/송장/비디오가 누락된 경우 표시됨)
- **"AP:** "잘못된 (결함 있는) 제품을 받았습니다."
- **"32":** "송장과 함께 상품이 반품되었습니다."
- **"33":** "송장과 함께 상품이 반품되었습니다."
- **"34":** "구매자의 주장이 정확하지 않습니다."
- **"35":** "구매자의 주장이 정확하지 않습니다."
- **"36":** "구매자에게 송장과 함께 환불되었습니다."
- **"37":** "구매자에게 송장과 함께 환불되었습니다."
- **"38":** "제품 목록에서 본 제품이 아닙니다."
- **"51":** "목록에 없지만 조항 인스턴스가 있습니다."
- **"52":** "잘못된 손상으로 잘못된 제품을 받았습니다."
- **"53":** "해당 사항 없음" (개봉 사진/송장/비디오가 누락된 경우 표시됨)
- **"54":** "잘못된 (결함 있는) 제품을 받았습니다."
- **"55":** "제품이 적절한 기간에 있지 않습니다."
- **"56":** "해당 사항 없음" (개봉 사진/송장/비디오가 누락된 경우 표시됨)
- **"57":** "반품된 제품은 구매자의 법적 철회 권리에서 제외됩니다."
- **"58":** "판매자의 청구 분쟁에서 구매자가 불합리했습니다."
- **"59":** "제품이 양호한 상태였습니다."

---

## AttributeType

- **INT_TYPE**
- **STRING_TYPE**
- **ENUM_TYPE**
- **FLOAT_TYPE**
- **DATE_TYPE**
- **MULTIPLE_TYPE**

---

## AttributeInputTypeEdit

- **DROP_DOWN**
- **TEXT_FILED**
- **MULTIPLE_PICKER**
- **MATERIE_SELECT**
- **MATERIAL_SELECT_COMBO_BOX**

자세한 내용은 다음을 확인하십시오: https://partner.tiktokshop.com/docv2/page/650dbc2c85f0eg028de51c4c

---

## CancelReason (Seller)

- **OUT_OF_STOCK**
- **UNDELIVERABLE_AREA** (TW 및 MY에만 해당)

---

## FeeType

- **BASE_SELECTION**
- **SHIPPING_FEE**
- **TRANSACTION_FEE**
- **CUSTOMER_PRICE**
- **SELLER_LOGISTICS**

---

# API 문서

## CancelReason (Game)

### 옵션:
- OUT_OF_STOCK
- DISCONTINUED_PRODUCT (판매자용)

---

## FeeType

- BULK_INSERTION
- NICE_INPUT
- FIXED_DEFAULT_PRICE
- CUSTOM_PRICE
- CHARGE_LOGISTICS

---

## PaymentMethod

- Overseas [ID, VN, TW, SG, MY, TH, PH]
- Nespay Credit Card [ID]
- Nespay Installment [ID]
- Nespay Virtual Account [ID]
- Bayar COD [ID]
- Airpay COD [TW]
- Airpay Credit Card [TW, VN, TH, MY, TH, PH]
- Bank BCA (Manual Transfer) [ID]
- Bank BCA (Manual Transfer) [ID]
- Bank BNI (Manual Transfer) [ID]
- Bank BRI (Manual Transfer) [ID]
- Bank CIMB Niaga (Manual Transfer) [ID]
- Bank Mandiri (Manual Transfer) [ID]
- Boost Bank Transfer [MY]
- Boost Wallet [MY]
- Bank Transfer [BG]
- Bank Transfer [TW]
- Bank Transfer [MY]
- Bank Transfer [SG]
- Bank Transfer [TH]
- Bank Transfer [PH]
- ATM Payment (Debit) [TH]
- ATM Payment (Debit) [TH]
- ATM Payment (Credit) [TH]
- Online Payment (Debit) [TH]
- Bank Transfer (Maybank) [MY, TH, SG, MY, TH, PH]
- Shopee Seller Wallet [ID]
- Shopee Seller Wallet [ID]
- Shopee Seller [ID, VN, PH]
- Bank BRI (Virtual Account) [ID]
- Bank BRI (Virtual Account) [ID]
- Bank BCA (Virtual Account) [ID]
- Bank BNI (Virtual Account) [ID]
- Bank Mandiri (Virtual Account) [ID]
- Virtual Account Payment [ID]
- Airpay Pay [ID]
- SPayLater [ID]
- Pay By MTG
- Banking Payment [TH]
- Banking Payment (Alipay) [TH]
- Banking Payment (ATM) [TH]
- Banking Payment (BCA) [TH]
- Banking Payment (Counter) [TH]
- Dragonpay – Online Payment [PH]
- Dragonpay – Online Payment [PH]
- Payout Seller Self-Arrange [ID, VN, TW, SG, MY, TH, PH]
- Kredivo [ID]
- Kredivo – Limits [ID]
- Kredivo – 3 Months Installment [ID]
- Kredivo – 6 Months Installment [ID]
- Kredivo – 12 Months Installment [ID]
- Nespay Credit Card Installment [ID]
- BCA One [ID]
- BCA One [ID]
- Free [TH, SG, MY, TH, PH]
- Paypal CC Installment [PH]
- Shopee Credit Card [ID]
- Eliana Credit Card Installment [BR]
- Eliana Credit Card Installment in installment plan [BR]
- Eliana Credit Card Installment in installment plan [BR]
- Eliana Credit Card Installment in installment plan [BR]
- Eliana Credit Card Installment in installment plan [BR]
- Eliana Credit Card Installment in installment plan [BR]
- Eliana Credit Card Installment in installment plan [BR]
- Eliana Boleto [BR]

---

## CancelReason

### 옵션:
- Out of Stock
- Fraudulent Order
- Undeliverable Area
- COD Uncompleted
- Duplicate Order
- Same Completed
- Illegible Order
- Buyer Cancelled
- Buyer Failed to Make Payment
- Logistics Request to Cancelled
- Invalid Order
- SPL product
- COD Rejection
- Seller did not list
- Seller did not list Cancelled
- Other
- Partner Seller
- Unknown
- Your opponent did not approve order on time
- You are unable to place out of the moment
- Other

---

## ShippingDocumentType

- NORMAL_AIR_WAYBILL
- THERMAL_AIR_WAYBILL
- NORMAL_JOB_AIR_WAYBILL
- THERMAL_JOB_AIR_WAYBILL

---

## ItemStatus

- NORMAL
- DELETED
- UNLIST
- REVIEWING
- BANNED
- SHOPEE_DELETE

---

## StockType

1. Shopee Warehouse Stock
2. Seller Stock

---

## Language

- zh-hans
- zh-hant
- FILIPINO
- th-my
- id
- vi
- th
- ms
- en-ph
- ja
- pt
- de-CL
- de-co
- de-mx
- es-br

---

## PromotionType

- Campaign
- Discount Promotions
- Bundle Deal
- Group Deal
- Bundle Deal
- Flash Sale
- Add-on Discount
- Brand Sale
- Free Shipping
- Gift with purchase
- Fortunate Prize
- Combo Sale
- Seller Discounting

---

## BuyerCancelReason

- Seller is not Responsive to buyer's Inquiries
- Seller asks Buyer to Cancel
- Modify Existing Order
- Product Has Null Barcode
- Buyer Cancel Before Ship My Order
- Unverified
- Others
- Seller is Input Voucher Code
- Need to Change Delivery Address
- Need to Change Delivery Address
- Buyer would rather buy from another seller
- Buyer Cancelled

---

# BuyerCancelReason

- Seller is not responsive to buyer's inquiries
- Seller is uncooperative
- Modify Existing Order
- Product Has Bid Reviews
- Seller Doesn't Accept The Order
- Seller is untrustworthy
- Others
- Cancel
- Need to Modify Voucher Code
- Need to Change Delivery Address
- Need to Modify Product Voucher Code
- Change of mind
- Payment Procedure too Troublesome
- Found Cheaper Elsewhere
- Don't want to buy anymore
- Your approval rejected the order
- You are unable to push order at the moment
- Seller reject the order
- Change of mind (others)
- Modify existing order (size, test, voucher, etc)
- Change of mind - others

---

# TrackingLogisticsStatus

- RETAIL
- ORDER_SUBMITTED
- ORDER_TRANSMITTED
- ONHOLD
- CANCEL_CREATED
- PICKUP_REQUESTED
- PICKUP_DONE
- PICKED_UP
- SHIPMENT_INBOUND
- INBOUNDED
- TIMEOUT
- LOST
- DAMAGED
- UPDATE_SUBMITTED
- UPDATE_TRANSMITTED
- TRANSFERRED
- RETURNED
- RETURN_INBOUND
- RETURN_HANDED
- DISPOSED
- CANCEL
- DELIVERED
- CANCEL_CREATED
- FAILED_ORDER_INIT
- FAILED_ORDER_SUBMITTED
- FAILED_ORDER_TRANSMITTED
- FAILED_PICKUP_REQUESTED
- FAILED_PICKED_UP
- FAILED_SHIPMENT_INBOUND
- FAILED_SHIPMENT_INBOUNDED
- FAILED_UPDATE_CREATED
- FAILED_UPDATE_SUBMITTED
- FAILED_UPDATE_TRANSMITTED
- FAILED_CANCEL_CREATED
- FAILED_CANCEL

---

# SellerProofStatus

- APPROVED
- PENDING
- UPLOADED
- REJECTED

---

# TransactionType

- **ESCROW_VERIFIED_ADD = B/I**: 에스크로 수수료가 확인되어 판매자에게 지급되었습니다.
- **WITHDRAWAL_VERIFIED_ADD = B/I**: 주문에 에스크로 금액이 인출 금액보다 많은 경우 에스크로 수수료가 확인되어 청구되었습니다.
- **WITHDRAWAL_CREATED = [D]**: 판매자가 인출을 생성하여 잔액에서 차감되었습니다.
- **CANCELLATION_REFUND_ADD = I**: 클로백이 완료되어 에스크로 금액이 구매자에게 환불되었습니다.
- **CANCELLATION_REFUND_DEDUCT = [D]**: 클로백이 완료되어 에스크로 금액이 판매자에게서 차감되었습니다.
- **REFUND_VERIFIED_ADD = S/B**: 반전 주문 환불
- **REVERSAL_FEE_DEDUCT = [D]**: 반전 주문 환불
- **ADJUSTMENT_CREDIT_ADD = I/B**: 하나의 조정이 판매자에게 추가되었습니다.
- **ADJUSTMENT_DEBIT = [D]**: 하나의 조정이 판매자에게서 청구되었습니다.
- **FTA_ADJUSTMENT_ADD = I/B**: 하나의 조정 금액이 OMS에 성공적으로 기록된 후 FTA 조정 금액이 판매자의 지갑에 추가되거나 판매자 잔액 지갑에 추가되었습니다.
- **FTA_ADJUSTMENT_DEDUCT = [D]**: 조정 요청을 통해 청구되는 경우 조정 금액이 판매자의 지갑에서 차감됩니다.
- **ADJUSTMENT_CREDIT_ADD = I/B**: 하나의 조정이 판매자 지갑에 추가되었습니다.
- **ADJUSTMENT_FEE_DEDUCT = [D]**: FTA 조정이 차감되어 에스크로 지갑이 추가되었습니다.
- **INSPECTION_FEE_DEDUCT = [D]**: 실패한 경우 검사 보고서에 대한 요금이 청구됩니다.
- **INSPECTION_VAL_SM_DEDUCT = [D]**: 처방 보고서 VAT 수수료에 대한 실패한 요금이 청구됩니다.
- **ADJUSTMENT_FOR_SPLITTING_ESCROW_OFFSET = I/B**: 조정 제안 송장이 확인된 경우
- **AFFILIATE_COMMISSION_ADD = I/B**: 추가 수수료는 주문 판매자 지갑입니다.
- **AFFILIATE_COMMISSION_FEE_ADD = B/I**: 판매자 지갑 후 재판매 판매자 제휴 수수료
- **CAMPAIGNPROMO_ESCROW_ADD = B/I**: 판매자는 캠페인/프로모션 지갑에 에스크로 지갑 금액이 있는 후 지불합니다.
- **CAMPAIGNPROMO_DISCOUNT_FEE_DEDUCT = [D]**: TRANSAC_DISCOU의 판매자 할인 또는 캠페인 가격
- **SELLER_COMPENSATION_ADD = I**: 이 지갑 금액 삽입에는 FEBpre(특정 총 상점 학습)가 있습니다.
- **SELLER_COMPENSATION_FEE_DEDUCT = [D]**: 반전 주문 환불
- **CAMPAIGNPROMO_VOUCHER_DISCOUNT_DEDUCT = [D]**: 바우처 할인이 적용됩니다.
- **PROMOTIONAL_PACKAGE_FEE_DEDUCT = [D]**: 이것은 결제 패키지 프로모션 수수료입니다. 판매자는 필요한 경우 결제 요청을 합니다. 이 수수료 유형은 해당 판매자를 보상하기 위한 것입니다.
- **CAMPAIGNPROMO_ADD = B/I**: 구매자 프로모션으로 추가 공제하여 과도한 예약(또는 과도한 수수료)을 방지합니다.
- **AFFILIATE_COMMISSION_ADD = I/B**: 추가 수수료가 주문 판매자 지갑에 있습니다.
- **AFFILIATE_COMMISSION_FEE_ADD = B/I**: 판매자 지갑 후 재판매 판매자 제휴 수수료
- **CAMPAIGNPROMO_ESCROW_ADD = B/I**: 판매자는 캠페인/프로모션 지갑에 에스크로 지갑 금액이 있는 후 지불합니다.
- **CAMPAIGNPROMO_DISCOUNT_FEE_DEDUCT = [D]**: TRANSAC_DISCOU의 판매자 할인 또는 캠페인 가격
- **SELLER_COMPENSATION_ADD = I**: 이 지갑 금액 삽입에는 FEBpre(특정 총 상점 학습)가 있습니다.
- **SELLER_COMPENSATION_FEE_DEDUCT = [D]**: 주문 취소 환불
- **CAMPAIGNPROMO_VOUCHER_DISCOUNT_DEDUCT = [D]**: 바우처 할인이 적용됩니다.
- **PROMOTIONAL_PACKAGE_FEE_DEDUCT = [D]**: 이것은 결제 패키지 프로모션 수수료입니다. 판매자는 필요한 경우 결제 요청을 합니다. 이 수수료 유형은 해당 판매자를 보상하기 위한 것입니다.
- **CAMPAIGNPROMO_ADD = B/I**: 과도한 예약(또는 과도한 수수료)을 방지하기 위해 구매자 프로모션으로 추가 공제합니다(캠페인 수수료에만 해당, 3P 역할).
- **CAMPAIGNPROMO_PACKAGE_SELLER = I/B**: 캠페인/프로모션 후 판매자 지갑 금액이 추가되었습니다.
- **CAMPAIGNPROMO_PACKAGE_INWARD = I/B**: 캠페인 수수료가 과소 청구된 판매자로부터 추가 공제합니다(3P 역할).
- **CAMPAIGNPROMO_UNRECORD_DEDUCT = [D]**: 판매자가 금액을 기록하지 않을 때 캠페인 공제
- **CAMPAIGNPROMO_UNRECORD_INWARD_DEDUCT = [D]**: 과소 청구된 판매자로부터 캠페인 공제
- **FAST_ESCROW_DISBURSE = I/D + ADD**: 빠른 에스크로의 빠른 지급이 판매자에게 지급된 경우
- **AFFILIATE_ADV_SELLER_FEE = I/D + DEDUCT**: 제휴 광고는 판매자로부터 참여해야 합니다.
- **AFFILIATE_CONTENT_CREATOR_FEE = I/D + DEDUCT**: 콘텐츠 판매자 수수료 제휴 수수료가 판매자 지갑에서 청구되었습니다.
- **FAST_ESCROW_DEDUCT = I/D**: 반품 및 환불 시 빠른 에스크로가 판매자 잔액에서 차감되는 경우
- **FAST_ESCROW_DISBURSE_BUNDLE = I/D**: 빠른 에스크로의 일반 지급이 판매자에게 지급된 경우
- **AFFILIATE_KOL_FEE = I/D + DEDUCT**: 해당 판매와 관련된 KOL 수수료 제휴
- **AFFILIATE_KOL_SELLER_FEE = I/D + DEDUCT**: KOL 수수료 제휴 수수료가 판매자 지갑에서 청구되었습니다.
- **BPM_DEDUCT = S/D**: P/M 요금 판매자 지갑 또는 결제
- **AFFILIATE_CELEBRITY_FEE = I/D + DEDUCT**: 유명인 수수료 제휴가 판매자로부터 청구되었습니다.
- **AFFILIATE_CELEBRITY_SELLER_FEE = I/D + DEDUCT**: 유명인 수수료 제휴 수수료가 판매자 지갑에서 청구되었습니다.
- **APM_REFUND_ADD = S/B**: 감정 평가 딜러 환불
- **BY_SPLPAY_VERIFIED_ADD = I/B**: 원본 제품 구매 환불이 확인되어 판매자에게 지급되었습니다.
- **BY_SPLPAY_CREATED_DEDUCT = [D]**: 분할 결제가 분할 결제 요구 사항을 생성했으므로 청구된 ID(KODICE)가 차감되었습니다.
- **BPM_DISBURSE_ADD = I/B**: 판매자 지갑에 지급금 지급

---

# SellerCompensationStatus

- COMPENSATION_APPLICABLE
- COMPENSATION_NOT_APPLICABLE
- COMPENSATION_PENDING_REQUEST
- COMPENSATION_NOT_RECORDED
- COMPENSATION_RECORDED
- COMPENSATION_REJECTED
- COMPENSATION_APPROVED
- COMPENSATION_PARTIALLY_REFUNDED
- COMPENSATION_REFUNDED
- COMPENSATION_RTS_CLAWBACK

---

# SellerProofType

- PENDING_RESPOND
- PENDING_APPROVAL
- SUBMITTED

---

# 반품 환불 요청 유형

- **0 - 일반 RR**: 구매자가 예상 배송 날짜를 기준으로 소포를 받은 경우 발행됩니다.
- **1 - 배송 중 RR**: (RR은 구매자에게 아직 배송 중인 경우 발행됩니다.)
- **2 - 판매자 반품 RR**: (RR은 배송 시 구매자가 소포를 거부한 후 판매자가 발행합니다.)

---

# 유효성 검사 유형

- **0 - 유효성 검사 유형 없음**: 아직 선택된 유효성 검사 유형이 없습니다. Shopee에서 유효성 검사를 위해 판매자가 반품한 소포
- **1 - 유휴**: 구매자에게 환불할지 분쟁을 제기할지 여부 때문에
- **2 - 무게 없음**: 무게 없는 문제로 판매자에 따라 Shopee 3 PLEDGE에 대한 유효성 검사
- **3 - 무게 일치 대기 중**: 판매자를 위한 3PL PLEDGE의 유효성 검사 유형

---

# 역물류 상태

## [일반 반품]

- **LOGISTICS_PENDING_ARRANGE**: 이제 반품은 판매자가 배송 옵션을 선택하기를 기다리고 있습니다. 대량 통합 물류 및 비통합 물류와 동일
- **LOGISTICS_PICKUP_CREATED**: 물류 소포가 성공적으로 생성되었습니다. 이제 Shopee에서 추적 번호를 제공합니다. 통합 물류와 비통합 물류 모두 동일
- **LOGISTICS_PICKED_UP**: 소포가 구매자로부터 픽업되었습니다. 비통합 물류(판매자가 구매자로부터 직접 픽업)에만 사용할 수 있습니다. Shopee에서는 건너뛰므로 통합 물류에는 적용되지 않습니다.
- **LOGISTICS_PICKED_UP_DONE**: 통합 물류(제3자 물류 제공업체가 Shopee 물류 허브로 다시 배송) 또는 비통합 물류에만 사용할 수 있으며, 이는 판매자가 이미 반품 소포를 수락했음을 의미합니다.
- **LOGISTICS_INBOUND**: 소포가 이제 Shopee 물류 허브에 입고되었습니다. 제3자 물류 제공업체가 Shopee로 다시 배송하므로 통합 물류에만 사용할 수 있습니다.
- **LOGISTICS_PICKUP_DONE**: 통합 물류의 경우 이는 소포가 제3자(통합 물류 제공업체) 또는 판매자가 직접 픽업했음을 의미합니다.
- **LOGISTICS_INBOUND**: 비통합 물류의 경우 이는 사용자가 배송 증거를 입력했음을 의미합니다.
- **LOGISTICS_TO_DELIVER**: 이제 소포를 배송할 준비가 되었습니다. 제3자 물류 제공업체가 Shopee에 업데이트하므로 통합 물류에만 사용할 수 있습니다.
- **LOGISTICS_DELIVERED**: 소포가 판매자에게 배송되었습니다. 제3자 물류 제공업체가 Shopee에 다시 업데이트하므로 통합 물류에만 사용할 수 있습니다.
- **LOGISTICS_LOST**: 소포가 분실되었습니다. 제3자 물류 제공업체가 Shopee에 다시 업데이트하므로 통합 물류에만 사용할 수 있습니다.

## [비정상 RR]

- **LOGISTICS_CANCEL**: 반품 요청이 취소되었습니다.

---

# 물류 상태 문서

## 통합 물류 상태

- **LOGISTICS_PICKUP_DONE**: 통합 물류의 경우 이는 소포가 제3자 물류 제공업체에 의해 픽업되었음을 의미합니다. 비통합 물류의 경우 이는 사용자가 배송 증거를 입력했음을 의미합니다.

- **LOGISTICS_DELIVERY_FAILED**: 판매자에게 소포 배송이 실패했습니다. 제3자 물류 제공업체가 Shopee에 다시 업데이트하므로 통합 물류에만 사용할 수 있습니다.

- **LOGISTICS_LOST**: 소포가 분실된 것으로 표시되었습니다. 제3자 물류 제공업체가 Shopee에 다시 업데이트하므로 통합 물류에만 사용할 수 있습니다.

- **LOGISTICS_DELIVERY_DONE**: 소포가 판매자에게 성공적으로 배송되었습니다. 제3자 물류 제공업체가 Shopee에 다시 업데이트하므로 통합 물류에만 사용할 수 있습니다.

## [배송 중 RR]

- 준비 중
- 배송 완료
- 배송 실패
- 분실

## [현장 반품]

- 준비 중
- 배송 완료
- 배송 실패
- 분실

## 반품 후 물류 상태

참고: 이는 창고에서 판매자에게 다시 보내는 반품 소포에만 적용됩니다.

- **POST_RETURN_LOGISTICS_REQUEST_CREATED**: 추적 번호로 물류 요청이 성공적으로 생성되었습니다.

- **POST_RETURN_LOGISTICS_REQUEST_CANCELED**: 창고 팀에서 물류 요청을 취소했습니다.

- **POST_RETURN_LOGISTICS_PICKUP_FAILED**: 소포 픽업에 실패했습니다.

- **POST_RETURN_LOGISTICS_PICKUP_RETRY**: 소포 픽업을 다시 시도합니다.

- **POST_RETURN_LOGISTICS_PICKUP_DONE**: 픽업 성공; 목적지로 가는 중입니다.

- **POST_RETURN_LOGISTICS_DELIVERY_FAILED**: 소포 배송에 실패했습니다. 운전자는 소포를 창고로 다시 반환합니다.

- **POST_RETURN_LOGISTICS_DELIVERY_DONE**: 소포 배송 성공

- **POST_RETURN_LOGISTICS_LOST**: 소포가 분실된 것으로 표시되었습니다.

## 사용 사례

1. TikTok Shop API와 통합
2. 주문 처리 워크플로우 이해
3. 반품 및 분쟁 처리
4. 물류 및 이행 관리
5. 결제 처리

## 관련 API

- Announcement Console
- OrderStatus
- ReturnStatus
- ReturnSolution
- ReturnReason
- LogisticsStatus
- PackageFulfillmentStatus
- ReturnDisputeReasonId
- AttributeType
- AttributeInputTypeEdit
- CancelReason (Seller)
- FeeType
- CancelReason (Game)
- PaymentMethod

---

## 원문 (English)

### Summary

This document provides a comprehensive overview of various data definitions, status codes, reasons, and types used within the TikTok Shop API. It includes details on order status, return status, logistics status, return reasons, dispute reasons, attribute types, cancel reasons, fee types, and payment methods.

### Content

# API Documentation Extract

## Announcement Console

Getting Started > V2.0 Data Definition

- **uint32**: 32 bit unsigned integer
- **uint64**: 64 bit unsigned integer
- **timestamp**: unit32
- **string**: UTF-8 Code. Default string of 2 characters

---

## OrderStatus

- **UNPAID** Order is created, buyer has not paid yet
- **READY_TO_SHIP** Order was arranged shipment
- **READY_TO_SHIP_PENDING** The order is ready to ship and get tracking number from 3PL
- **RETRY_SHIP** The order need resend
- **IN_TRANSIT** The parcel has been drop to 3PL or picked up by 3PL
- **SHIPPED** or **DELIVERED** The order has been received by Seller
- **COMPLETED** The order has been completed and closed
- **CANCELLED** The order has been cancelled
- **TO_RETURN** The buyer requested to return the order and order's return is processing
- **LOST** TikTok Ship order lost

---

## ReturnStatus

- **REQUESTED**
- **ACCEPTED**
- **IN_TRANSIT**
- **JUDGING**
- **CLOSED**
- **CANCELLED**
- **SELLER_DISPUTE**

---

## ReturnSolution

- **RETURN_REFUND**
- **REFUND_ONLY**

---

## ReturnReason

- **NONRECEIPT**
- **WRONG_ITEM**
- **COUNTERFEIT**
- **GIFT_DESC**
- **MATRL_NOISE**
- **SIZE_SHOES**
- **USED**
- **NO_REASON**
- **RECEIVED_DAMAGED**
- **UNWANTED**
- **ITEM_MISSING**
- **FUNCTIONAL_FAILED**
- **COLOR**
- **PHYSICAL_DMG**
- **FUNCTIONAL_DMG**
- **NOT_AS_DESCRIBED**
- **SURPCODE_PARCEL**
- **EXPIRES_PRODUCT**
- **LOST_IN_DELIVERY**
- **PARTIAL_MISSING**
- **CHANGE_OF_MIND**
- **INCOMPLETE_MISSING_ITEM**
- **ITEM_FAULT**
- **BROKEN_PRODUCTS**
- **DAMAGE_PACKAGE**
- **WRONG_DELIVERY**
- **SIZE_DEVIATION**
- **LOOK_FEEL**
- **POOR_QUALITY**
- **DIFFERENT_DESCRIPTION**

---

## LogisticsStatus

- **LOGISTICS_NOT_STARTED** Initial status, order not ready for fulfillment
- **LOGISTICS_REQUEST_CREATED** order arranged for pickup
- **LOGISTICS_COD_REJECTED** Integrated logistics COD. Order rejected for COD
- **LOGISTICS_PICKUP_NOT_DONE** order arranged pickup. non-COD. order preCOD. order passed pickup SLA. not picked up by 3PL
- **LOGISTICS_REQUEST_CREATED** order arranged shipment
- **LOGISTICS_PICKUP_DONE** order picked up by 3PL or failed track or parcel lost-non last mile
- **LOGISTICS_IN_TRANSIT** order in transit middle-mile
- **LOGISTICS_ARRIVAL** order calculated order status at LOGISTICS_REACH
- **LOGISTICS_DELIVERY_CANCELLED** Order delivery cancelled. either LOGISTICS_REQUEST_CREATED or LOGISTICS_PICKUP_DONE status
- **LOGISTICS_DELIVERY_FAILED** TikTok Ship fails delivery order
- **LOGISTICS_ROUTE_RETRY** order passing SPI. retry pickup
- **LOGISTICS_DELIVERY_FAILED** order passed last mile delivery failed
- **LOGISTICS_DELIVERED** TikTok Ship delivered order

---

## PackageFulfillmentStatus

- **LOGISTICS_NOT_START** Initial status, package not ready for fulfillment
- **LOGISTICS_READY** Package ready for fulfillment (the return/ removable). For non-COD, prdt, the COD. passed pickup SLA by 3PL
- **LOGISTICS_PICKED** Package picked up by 3PL
- **LOGISTICS_PICUP_DONE** Package handed over to 3PL
- **LOGISTICS_CANCELLED** Package successfully delivery
- **LOGISTICS_DELIVERY_FAIL** TikTok Ship order delivery failed. LOGISTICS_READY
- **LOGISTICS_REQUEST_CANCELLED** Order cancelled when package at LOGISTICS_REQUEST_CREATED
- **LOGISTICS_PICDUP_FAILED** Order cancelled when package at LOGISTICS_PICUP_DONE
- **LOGISTICS_LOST** Order lost at 3PL
- **LOGISTICS_DELIVERY_CANCELLED** TikTok Ship cancel order
- **LOGISTICS_DELIVERY_FAILED** Order cancelled due to SPI, delivery failed or lost
- **LOGISTICS_LOST** Order cancelled due to SPI, lost the Package

---

## ReturnDisputeReasonId

- **"1":** "I would like to reject the non-receipt claim"
- **"2":** "I don't agree with the claim"
- **"3":** "I have proof the item was delivered. (when the product(s) which was/were supposed to be returned has not been received)"
- **"6":** "商品拒收异常"
- **"8":** "买家拒收异常"
- **"9":** "包裹异常被拒收"
- **"9":** "其他"
- **"10":** "I have lost the item(s) and have proof of shipment"
- **"11":** "Returned item(s) do not match the item(s) ordered"
- **"12":** "I disagree with the claim as the goods are in good but used condition"
- **"13":** "I believe the item(s) is/are correctly priced"
- **"14":** "I agree with the return item(s) but I have not received the item(s)/ full return(s) supposed to be returned"
- **"15":** "I agree with the return item(s) but the package/ return item(s) has been damaged/tempered"
- **"16":** "I shipped the correct item(s) as per invoice"
- **"17":** "I shipped the correct item(s) as per invoice"
- **"18":** "I supplied the item(s) in good working condition"
- **"4":** "Products are not in the appropriate period"
- **"5":** "Products are not in the appropriate period"
- **"19":** "I dispute the reason for the dispute"
- **"20":** "Not applicable" (shown when missing opening photo/invoice/video)"
- **"AP:** "Received wrong (faulty) product"
- **"32":** "Item has been returned along invoice"
- **"33":** "Item has been returned along invoice"
- **"34":** "Buyer's claim is incorrect"
- **"35":** "Buyer's claim is incorrect"
- **"36":** "Buyer has been refunded along invoice"
- **"37":** "Buyer has been refunded along invoice"
- **"38":** "Not the product as seen in product listing"
- **"51":** "Not list but has the clause instance"
- **"52":** "Received wrong products with incorrect damages"
- **"53":** "Not applicable" (shown when missing opening photo/invoice/video)"
- **"54":** "Received wrong (faulty) product"
- **"55":** "Products are not in the appropriate period"
- **"56":** "Not applicable" (shown when missing opening photo/invoice/video)"
- **"57":** "The product(s) returned is excluded from buyer's statutory right of withdrawal"
- **"58":** "Buyer was unreasonable in Seller's charge dispute"
- **"59":** "Product(s) was/were in good condition"

---

## AttributeType

- **INT_TYPE**
- **STRING_TYPE**
- **ENUM_TYPE**
- **FLOAT_TYPE**
- **DATE_TYPE**
- **MULTIPLE_TYPE**

---

## AttributeInputTypeEdit

- **DROP_DOWN**
- **TEXT_FILED**
- **MULTIPLE_PICKER**
- **MATERIE_SELECT**
- **MATERIAL_SELECT_COMBO_BOX**

For more details, please check: https://partner.tiktokshop.com/docv2/page/650dbc2c85f0eg028de51c4c

---

## CancelReason (Seller)

- **OUT_OF_STOCK**
- **UNDELIVERABLE_AREA** (only for TW and MY)

---

## FeeType

- **BASE_SELECTION**
- **SHIPPING_FEE**
- **TRANSACTION_FEE**
- **CUSTOMER_PRICE**
- **SELLER_LOGISTICS**

---

# API Documentation

## CancelReason (Game)

### Options:
- OUT_OF_STOCK
- DISCONTINUED_PRODUCT (for the seller)

---

## FeeType

- BULK_INSERTION
- NICE_INPUT
- FIXED_DEFAULT_PRICE
- CUSTOM_PRICE
- CHARGE_LOGISTICS

---

## PaymentMethod

- Overseas [ID, VN, TW, SG, MY, TH, PH]
- Nespay Credit Card [ID]
- Nespay Installment [ID]
- Nespay Virtual Account [ID]
- Bayar COD [ID]
- Airpay COD [TW]
- Airpay Credit Card [TW, VN, TH, MY, TH, PH]
- Bank BCA (Manual Transfer) [ID]
- Bank BCA (Manual Transfer) [ID]
- Bank BNI (Manual Transfer) [ID]
- Bank BRI (Manual Transfer) [ID]
- Bank CIMB Niaga (Manual Transfer) [ID]
- Bank Mandiri (Manual Transfer) [ID]
- Boost Bank Transfer [MY]
- Boost Wallet [MY]
- Bank Transfer [BG]
- Bank Transfer [TW]
- Bank Transfer [MY]
- Bank Transfer [SG]
- Bank Transfer [TH]
- Bank Transfer [PH]
- ATM Payment (Debit) [TH]
- ATM Payment (Debit) [TH]
- ATM Payment (Credit) [TH]
- Online Payment (Debit) [TH]
- Bank Transfer (Maybank) [MY, TH, SG, MY, TH, PH]
- Shopee Seller Wallet [ID]
- Shopee Seller Wallet [ID]
- Shopee Seller [ID, VN, PH]
- Bank BRI (Virtual Account) [ID]
- Bank BRI (Virtual Account) [ID]
- Bank BCA (Virtual Account) [ID]
- Bank BNI (Virtual Account) [ID]
- Bank Mandiri (Virtual Account) [ID]
- Virtual Account Payment [ID]
- Airpay Pay [ID]
- SPayLater [ID]
- Pay By MTG
- Banking Payment [TH]
- Banking Payment (Alipay) [TH]
- Banking Payment (ATM) [TH]
- Banking Payment (BCA) [TH]
- Banking Payment (Counter) [TH]
- Dragonpay – Online Payment [PH]
- Dragonpay – Online Payment [PH]
- Payout Seller Self-Arrange [ID, VN, TW, SG, MY, TH, PH]
- Kredivo [ID]
- Kredivo – Limits [ID]
- Kredivo – 3 Months Installment [ID]
- Kredivo – 6 Months Installment [ID]
- Kredivo – 12 Months Installment [ID]
- Nespay Credit Card Installment [ID]
- BCA One [ID]
- BCA One [ID]
- Free [TH, SG, MY, TH, PH]
- Paypal CC Installment [PH]
- Shopee Credit Card [ID]
- Eliana Credit Card Installment [BR]
- Eliana Credit Card Installment in installment plan [BR]
- Eliana Credit Card Installment in installment plan [BR]
- Eliana Credit Card Installment in installment plan [BR]
- Eliana Credit Card Installment in installment plan [BR]
- Eliana Credit Card Installment in installment plan [BR]
- Eliana Credit Card Installment in installment plan [BR]
- Eliana Boleto [BR]

---

## CancelReason

### Options:
- Out of Stock
- Fraudulent Order
- Undeliverable Area
- COD Uncompleted
- Duplicate Order
- Same Completed
- Illegible Order
- Buyer Cancelled
- Buyer Failed to Make Payment
- Logistics Request to Cancelled
- Invalid Order
- SPL product
- COD Rejection
- Seller did not list
- Seller did not list Cancelled
- Other
- Partner Seller
- Unknown
- Your opponent did not approve order on time
- You are unable to place out of the moment
- Other

---

## ShippingDocumentType

- NORMAL_AIR_WAYBILL
- THERMAL_AIR_WAYBILL
- NORMAL_JOB_AIR_WAYBILL
- THERMAL_JOB_AIR_WAYBILL

---

## ItemStatus

- NORMAL
- DELETED
- UNLIST
- REVIEWING
- BANNED
- SHOPEE_DELETE

---

## StockType

1. Shopee Warehouse Stock
2. Seller Stock

---

## Language

- zh-hans
- zh-hant
- FILIPINO
- th-my
- id
- vi
- th
- ms
- en-ph
- ja
- pt
- de-CL
- de-co
- de-mx
- es-br

---

## PromotionType

- Campaign
- Discount Promotions
- Bundle Deal
- Group Deal
- Bundle Deal
- Flash Sale
- Add-on Discount
- Brand Sale
- Free Shipping
- Gift with purchase
- Fortunate Prize
- Combo Sale
- Seller Discounting

---

## BuyerCancelReason

- Seller is not Responsive to buyer's Inquiries
- Seller asks Buyer to Cancel
- Modify Existing Order
- Product Has Null Barcode
- Buyer Cancel Before Ship My Order
- Unverified
- Others
- Seller is Input Voucher Code
- Need to Change Delivery Address
- Need to Change Delivery Address
- Buyer would rather buy from another seller
- Buyer Cancelled

---

# BuyerCancelReason

- Seller is not responsive to buyer's inquiries
- Seller is uncooperative
- Modify Existing Order
- Product Has Bid Reviews
- Seller Doesn't Accept The Order
- Seller is untrustworthy
- Others
- Cancel
- Need to Modify Voucher Code
- Need to Change Delivery Address
- Need to Modify Product Voucher Code
- Change of mind
- Payment Procedure too Troublesome
- Found Cheaper Elsewhere
- Don't want to buy anymore
- Your approval rejected the order
- You are unable to push order at the moment
- Seller reject the order
- Change of mind (others)
- Modify existing order (size, test, voucher, etc)
- Change of mind - others

---

# TrackingLogisticsStatus

- RETAIL
- ORDER_SUBMITTED
- ORDER_TRANSMITTED
- ONHOLD
- CANCEL_CREATED
- PICKUP_REQUESTED
- PICKUP_DONE
- PICKED_UP
- SHIPMENT_INBOUND
- INBOUNDED
- TIMEOUT
- LOST
- DAMAGED
- UPDATE_SUBMITTED
- UPDATE_TRANSMITTED
- TRANSFERRED
- RETURNED
- RETURN_INBOUND
- RETURN_HANDED
- DISPOSED
- CANCEL
- DELIVERED
- CANCEL_CREATED
- FAILED_ORDER_INIT
- FAILED_ORDER_SUBMITTED
- FAILED_ORDER_TRANSMITTED
- FAILED_PICKUP_REQUESTED
- FAILED_PICKED_UP
- FAILED_SHIPMENT_INBOUND
- FAILED_SHIPMENT_INBOUNDED
- FAILED_UPDATE_CREATED
- FAILED_UPDATE_SUBMITTED
- FAILED_UPDATE_TRANSMITTED
- FAILED_CANCEL_CREATED
- FAILED_CANCEL

---

# SellerProofStatus

- APPROVED
- PENDING
- UPLOADED
- REJECTED

---

# TransactionType

- **ESCROW_VERIFIED_ADD = B/I**: Escrow fee been verified and paid to seller
- **WITHDRAWAL_VERIFIED_ADD = B/I**: If Escrow fee been verified and charged after order has escrow amount is more than withdrawal amount
- **WITHDRAWAL_CREATED = [D]**: The seller has created a withdrawal, so it is deducted from balance
- **CANCELLATION_REFUND_ADD = I**: If is a clawback is has been completed, so the escrow amount is refund to buyer
- **CANCELLATION_REFUND_DEDUCT = [D]**: If a clawback is has been completed, so the escrow amount is deducted from seller
- **REFUND_VERIFIED_ADD = S/B**: A reversal Order Refund
- **REVERSAL_FEE_DEDUCT = [D]**: A reversal Order Refund
- **ADJUSTMENT_CREDIT_ADD = I/B**: One adjustment has been added to seller
- **ADJUSTMENT_DEBIT = [D]**: One adjustment has been charged from seller
- **FTA_ADJUSTMENT_ADD = I/B**: FTA Adjustment amount has been added to seller's wallet or added to seller balance wallet after one adjustment amount is successfully recorded in OMS
- **FTA_ADJUSTMENT_DEDUCT = [D]**: An adjustment amount is deducted from seller's wallet if it takes charge via adjustment request
- **ADJUSTMENT_CREDIT_ADD = I/B**: One adjustment has been added to seller wallet
- **ADJUSTMENT_FEE_DEDUCT = [D]**: FTA adjustment has been subtracted, so the escrow wallet has been added
- **INSPECTION_FEE_DEDUCT = [D]**: If failed is charge for inspection report has to chargable
- **INSPECTION_VAL_SM_DEDUCT = [D]**: If failed charge for prescription report VAT fee chargable
- **ADJUSTMENT_FOR_SPLITTING_ESCROW_OFFSET = I/B**: If the adjustment offer invoice invoice has been verified
- **AFFILIATE_COMMISSION_ADD = I/B**: Additional commission is order seller wallet
- **AFFILIATE_COMMISSION_FEE_ADD = B/I**: To resale seller affiliate commission after seller wallet
- **CAMPAIGNPROMO_ESCROW_ADD = B/I**: The seller pays after Campaign/Promo wallet has escrow wallet amount
- **CAMPAIGNPROMO_DISCOUNT_FEE_DEDUCT = [D]**: Seller discount or campaign price in TRANSAC_DISCOU
- **SELLER_COMPENSATION_ADD = I**: In this wallet amount insert has the FEBpre (lear a a specific Gross Store)
- **SELLER_COMPENSATION_FEE_DEDUCT = [D]**: A reversal Order Refund
- **CAMPAIGNPROMO_VOUCHER_DISCOUNT_DEDUCT = [D]**: Voucher discount is applicable
- **PROMOTIONAL_PACKAGE_FEE_DEDUCT = [D]**: This is a payment package promotional fee. Seller has the payment request if needed This fee type is to compensate such sellers
- **CAMPAIGNPROMO_ADD = B/I**: To further deduct with buyer promo to prevent overbooking (or overcharge fees, only for campaigns fees, role for 3P)
- **CAMPAIGNPROMO_PACKAGE_SELLER = I/B**: Seller wallet amount has been added after Campaign/Promo
- **CAMPAIGNPROMO_PACKAGE_INWARD = I/B**: To further deduct from sellers who have been undercharged for their campaigns fees, role for 3P
- **CAMPAIGNPROMO_UNRECORD_DEDUCT = [D]**: Campaign deduct when seller unrecord amount
- **CAMPAIGNPROMO_UNRECORD_INWARD_DEDUCT = [D]**: Campaign deduct from sellers who are undercharged
- **FAST_ESCROW_DISBURSE = I/D + ADD**: If the fast disbursement of fast escrow has been paid to seller
- **AFFILIATE_ADV_SELLER_FEE = I/D + DEDUCT**: Affiliate ads have the to engaged from seller
- **AFFILIATE_CONTENT_CREATOR_FEE = I/D + DEDUCT**: Content seller FEE Affiliate fee has charged from seller wallet
- **FAST_ESCROW_DEDUCT = I/D**: If fast escrow is deducted from seller balance in the event of return and refund
- **FAST_ESCROW_DISBURSE_BUNDLE = I/D**: If the normal disbursement of fast escrow has been paid to seller
- **AFFILIATE_KOL_FEE = I/D + DEDUCT**: KOL Fee affiliate in seller for to compensation for those associated sales
- **AFFILIATE_KOL_SELLER_FEE = I/D + DEDUCT**: KOL Fee Affiliate fee has charged from seller wallet
- **BPM_DEDUCT = S/D**: a P/M Charge seller wallet or payment
- **AFFILIATE_CELEBRITY_FEE = I/D + DEDUCT**: Celebrity fee affiliate has charged from seller
- **AFFILIATE_CELEBRITY_SELLER_FEE = I/D + DEDUCT**: Celebrity Fee Affiliate fee has charged from seller wallet
- **APM_REFUND_ADD = S/B**: a Appraise Deler Refund
- **BY_SPLPAY_VERIFIED_ADD = I/B**: Original product purchase refund verified and paid to seller
- **BY_SPLPAY_CREATED_DEDUCT = [D]**: By Split-payment has created a Split-Pay requirement, so deducted charged ID (KODICE)
- **BPM_DISBURSE_ADD = I/B**: a disbursement pay out to seller wallet

---

# SellerCompensationStatus

- COMPENSATION_APPLICABLE
- COMPENSATION_NOT_APPLICABLE
- COMPENSATION_PENDING_REQUEST
- COMPENSATION_NOT_RECORDED
- COMPENSATION_RECORDED
- COMPENSATION_REJECTED
- COMPENSATION_APPROVED
- COMPENSATION_PARTIALLY_REFUNDED
- COMPENSATION_REFUNDED
- COMPENSATION_RTS_CLAWBACK

---

# SellerProofType

- PENDING_RESPOND
- PENDING_APPROVAL
- SUBMITTED

---

# Return Refund Request Type

- **0 - Normal RR**: This is issued by the buyer because they have received the parcel, based on estimated delivery date
- **1 - In transit RR**: (RR is issued by the buyer either duo to a yet in transit to buyer)
- **2 - Return to Seller RR**: (RR is issued by the seller after buyer rejected parcel at delivery)

---

# Validation Type

- **0 - No Validation Type**: There's no validation type selected yet. Parcel that was returned by the seller for validation at Shopee
- **1 - Idle**: Because whether to refund buyer for to issue dispute
- **2 - Weightless**: Validation for Shopee 3 PLEDGE pursuant to sellers with weightless issue
- **3 - Pending Weight Match**: Validation type for 3PL PLEDGE for sellers

---

# Reverse Logistics Status

## [Normal Return]

- **LOGISTICS_PENDING_ARRANGE**: Return is now pending seller to select shipping option. Same for bulk integrated logistics and non-integrated logistics
- **LOGISTICS_PICKUP_CREATED**: About to has been logistics parcel has been created successfully. Tracking number is now given by Shopee. Same for both integrated logistics and non-integrated logistics
- **LOGISTICS_PICKED_UP**: Parcel has been picked up from buyer. Only available for non-integrated logistics (seller pick-up directly from buyer). Not applicable for integrated logistics since this is skipped for Shopee
- **LOGISTICS_PICKED_UP_DONE**: Only available for integrated logistics (since this is shipped for third party logistics provider back to Shopee Logistics hub) or for non-integrated logistics, this means the seller has already accepted return parcel
- **LOGISTICS_INBOUND**: Parcel is now inbounded at Shopee Logistics Hub. Only available for integrated logistics since this is shipped for third-party logistics provider back to Shopee
- **LOGISTICS_PICKUP_DONE**: For integrated logistics, this means the parcel has been picked up by a third party (integrated logistics provider) or the seller directly
- **LOGISTICS_INBOUND**: For non-integrated logistics, this means the user has entered shipping proof
- **LOGISTICS_TO_DELIVER**: Parcel is now ready for delivery. Only available for integrated logistics since this is updated by third-party logistics provider to Shopee
- **LOGISTICS_DELIVERED**: Parcel has been delivered to seller. Only available for integrated logistics since this is updated by third-party logistics provider back to Shopee
- **LOGISTICS_LOST**: Parcel has been lost. Only available for integrated logistics since this is updated by third-party logistics provider back to Shopee

## [Abnormal RR]

- **LOGISTICS_CANCEL**: Return Request has been canceled

---

# Logistics Status Documentation

## Integrated Logistics Statuses

- **LOGISTICS_PICKUP_DONE**: For integrated logistics, this means the parcel has been picked up by a third party logistics provider. For non-integrated logistics, this means the user has entered shipping proof.

- **LOGISTICS_DELIVERY_FAILED**: Parcel delivery to seller has failed. Only available for integrated logistics since this is updated by third party logistics provider back to Shopee.

- **LOGISTICS_LOST**: Parcel has been marked as lost. Only available for integrated logistics since this is updated by third party logistics provider back to Shopee.

- **LOGISTICS_DELIVERY_DONE**: Parcel has been successfully delivered to seller. Only available for integrated logistics since this is updated by third party logistics provider back to Shopee.

## [In-transit RR]

- Preparing
- Delivered
- Delivery Failed
- Lost

## [Return-on-the-Spot]

- Preparing
- Delivered
- Delivery Failed
- Lost

## Post Return Logistics Status

Note this is only applicable to return parcels sent from warehouse back to seller

- **POST_RETURN_LOGISTICS_REQUEST_CREATED**: Logistics request generated successfully with tracking number

- **POST_RETURN_LOGISTICS_REQUEST_CANCELED**: Logistics request cancelled by warehouse team

- **POST_RETURN_LOGISTICS_PICKUP_FAILED**: Failed to pickup parcel

- **POST_RETURN_LOGISTICS_PICKUP_RETRY**: Subsequent attempt to pickup parcel

- **POST_RETURN_LOGISTICS_PICKUP_DONE**: Successful pickup; on the way to destination.

- **POST_RETURN_LOGISTICS_DELIVERY_FAILED**: Failed delivery of parcel. Driver will return parcel back to warehouse.

- **POST_RETURN_LOGISTICS_DELIVERY_DONE**: Successful delivery of parcel

- **POST_RETURN_LOGISTICS_LOST**: Parcel marked as Lost

---

**문서 ID**: developer-guide.31
**플랫폼**: shopee
**URL**: https://open.shopee.com/developer-guide/31
**처리 완료**: 2025-10-16T08:32:06
