package models

import (
	"encoding/json"
	"time"
)

// ========================================
// Qoo10JP V2 Models - 고도화된 데이터 모델
// ========================================

// Qoo10JPShopV2 represents a Qoo10JP shop account (v2)
type Qoo10JPShopV2 struct {
	ID                string     `json:"id" db:"id"`
	SellerID          string     `json:"seller_id" db:"seller_id"`
	ShopName          string     `json:"shop_name" db:"shop_name"`
	APIID             string     `json:"api_id" db:"api_id"`
	CertificationKey  string     `json:"certification_key" db:"certification_key"` // 암호화된 키
	Region            string     `json:"region" db:"region"`
	IsActive          bool       `json:"is_active" db:"is_active"`
	LastSyncAt        *time.Time `json:"last_sync_at" db:"last_sync_at"`
	PlatformAccountID string     `json:"platform_account_id" db:"platform_account_id"`
	TokenExpireAt     *time.Time `json:"token_expire_at" db:"token_expire_at"`
	CreatedAt         time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at" db:"updated_at"`
}

// Qoo10JPOrderV2 represents an order in the v2 schema
type Qoo10JPOrderV2 struct {
	ID string `json:"id" db:"id"`

	// 주문 기본 정보
	OrderNo  string `json:"order_no" db:"order_no"`
	PackNo   string `json:"pack_no" db:"pack_no"`
	SellerID string `json:"seller_id" db:"seller_id"`
	Region   string `json:"region" db:"region"`

	// 상태 정보
	Qoo10Status string `json:"qoo10_status" db:"qoo10_status"` // Qoo10JP ShippingStatus
	OmsStatus   string `json:"oms_status" db:"oms_status"`     // 내부 OMS 상태

	// 날짜 정보
	OrderDate       *time.Time `json:"order_date" db:"order_date"`
	PaymentDate     *time.Time `json:"payment_date" db:"payment_date"`
	ShippingDate    *time.Time `json:"shipping_date" db:"shipping_date"`
	DeliveredDate   *time.Time `json:"delivered_date" db:"delivered_date"`
	OrderCreateTime *time.Time `json:"order_create_time" db:"order_create_time"`
	OrderUpdateTime *time.Time `json:"order_update_time" db:"order_update_time"`

	// 금액 정보
	OrderPrice     float64 `json:"order_price" db:"order_price"`
	TotalAmount    float64 `json:"total_amount" db:"total_amount"`
	Discount       float64 `json:"discount" db:"discount"`
	SellerDiscount float64 `json:"seller_discount" db:"seller_discount"`
	ShippingRate   float64 `json:"shipping_rate" db:"shipping_rate"`
	SettlePrice    float64 `json:"settle_price" db:"settle_price"`
	Currency       string  `json:"currency" db:"currency"`

	// 구매자 정보
	BuyerID     string `json:"buyer_id" db:"buyer_id"`
	BuyerName   string `json:"buyer_name" db:"buyer_name"`
	BuyerKana   string `json:"buyer_kana" db:"buyer_kana"`
	BuyerEmail  string `json:"buyer_email" db:"buyer_email"`
	BuyerPhone  string `json:"buyer_phone" db:"buyer_phone"`
	BuyerMobile string `json:"buyer_mobile" db:"buyer_mobile"`

	// 수령인 정보
	RecipientName     string `json:"recipient_name" db:"recipient_name"`
	RecipientKana     string `json:"recipient_kana" db:"recipient_kana"`
	RecipientPhone    string `json:"recipient_phone" db:"recipient_phone"`
	RecipientMobile   string `json:"recipient_mobile" db:"recipient_mobile"`
	RecipientZipcode  string `json:"recipient_zipcode" db:"recipient_zipcode"`
	RecipientAddress  string `json:"recipient_address" db:"recipient_address"`
	RecipientAddress1 string `json:"recipient_address1" db:"recipient_address1"`
	RecipientAddress2 string `json:"recipient_address2" db:"recipient_address2"`

	// 배송 정보
	ShippingMethod   string `json:"shipping_method" db:"shipping_method"`
	DeliveryCompany  string `json:"delivery_company" db:"delivery_company"`
	TrackingNumber   string `json:"tracking_number" db:"tracking_number"`
	PackingNo        string `json:"packing_no" db:"packing_no"`
	SellerDeliveryNo string `json:"seller_delivery_no" db:"seller_delivery_no"`

	// 결제 정보
	PaymentMethod string `json:"payment_method" db:"payment_method"`
	PaymentStatus string `json:"payment_status" db:"payment_status"`

	// 상품 정보 (단일 상품 - 호환성)
	ItemNo         string `json:"item_no" db:"item_no"`
	ItemTitle      string `json:"item_title" db:"item_title"`
	SellerItemCode string `json:"seller_item_code" db:"seller_item_code"`
	OptionName     string `json:"option_name" db:"option_name"`
	OptionCode     string `json:"option_code" db:"option_code"`
	OrderQty       int    `json:"order_qty" db:"order_qty"`

	// 메타데이터
	Flag            string     `json:"flag" db:"flag"`
	FulfillmentType string     `json:"fulfillment_type" db:"fulfillment_type"`
	Notes           string     `json:"notes" db:"notes"`
	LabelPrintedAt  *time.Time `json:"label_printed_at" db:"label_printed_at"`

	// 원본 데이터 (JSON)
	RawData json.RawMessage `json:"raw_data" db:"raw_data"`

	// 시스템 정보
	SyncedAt  *time.Time `json:"synced_at" db:"synced_at"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`

	// 관계 (DB에는 저장 안됨)
	Items []Qoo10JPOrderItemV2 `json:"-" db:"-"`
}

// Qoo10JPOrderItemV2 represents an order item in the v2 schema
type Qoo10JPOrderItemV2 struct {
	ID string `json:"id" db:"id"`

	// 연관 정보
	OrderNo  string `json:"order_no" db:"order_no"`
	PackNo   string `json:"pack_no" db:"pack_no"`
	SellerID string `json:"seller_id" db:"seller_id"`

	// 상품 기본 정보
	ItemNo    string `json:"item_no" db:"item_no"`
	ItemCode  string `json:"item_code" db:"item_code"`
	ItemName  string `json:"item_name" db:"item_name"`
	ItemTitle string `json:"item_title" db:"item_title"`

	// 옵션 정보
	OptionName string `json:"option_name" db:"option_name"`
	OptionCode string `json:"option_code" db:"option_code"`

	// 수량 및 가격
	Quantity        int     `json:"quantity" db:"quantity"`
	UnitPrice       float64 `json:"unit_price" db:"unit_price"`
	OriginalPrice   float64 `json:"original_price" db:"original_price"`
	DiscountedPrice float64 `json:"discounted_price" db:"discounted_price"`
	TotalPrice      float64 `json:"total_price" db:"total_price"`

	// 상품 상태
	ItemStatus string `json:"item_status" db:"item_status"`

	// 이미지
	ImageURL string `json:"image_url" db:"image_url"`

	// 원본 데이터
	RawData json.RawMessage `json:"raw_data" db:"raw_data"`

	// 시스템 정보
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// Qoo10JPSyncJob represents a synchronization job for order collection
type Qoo10JPSyncJob struct {
	ID string `json:"id" db:"id"`

	// 작업 식별
	SellerID string `json:"seller_id" db:"seller_id"`
	JobType  string `json:"job_type" db:"job_type"`

	// 기간 정보
	StartDate *time.Time `json:"start_date" db:"start_date"`
	EndDate   *time.Time `json:"end_date" db:"end_date"`

	// 상태 정보
	Status string `json:"status" db:"status"` // pending, running, completed, failed, cancelled

	// 진행 정보
	TotalCollected     int     `json:"total_collected" db:"total_collected"`
	TotalSaved         int     `json:"total_saved" db:"total_saved"`
	TotalUpdated       int     `json:"total_updated" db:"total_updated"`
	TotalSkipped       int     `json:"total_skipped" db:"total_skipped"`
	TotalFailed        int     `json:"total_failed" db:"total_failed"`
	ProgressPercentage float64 `json:"progress_percentage" db:"progress_percentage"`

	// 시간 정보
	StartedAt   *time.Time `json:"started_at" db:"started_at"`
	CompletedAt *time.Time `json:"completed_at" db:"completed_at"`
	DurationMs  int64      `json:"duration_ms" db:"duration_ms"`

	// 결과 정보
	ErrorMessage  string          `json:"error_message" db:"error_message"`
	ResultSummary json.RawMessage `json:"result_summary" db:"result_summary"`

	// 시스템 정보
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// ========================================
// OMS Status Constants
// ========================================

const (
	OmsStatusNew        = "NEW"        // 신규 주문
	OmsStatusProcessing = "PROCESSING" // 처리 중
	OmsStatusShipped    = "SHIPPED"    // 발송 완료
	OmsStatusDelivered  = "DELIVERED"  // 배송 완료
	OmsStatusCompleted  = "COMPLETED"  // 완료
	OmsStatusCancelled  = "CANCELLED"  // 취소
	OmsStatusReturning  = "RETURNING"  // 반품 중
	OmsStatusReturned   = "RETURNED"   // 반품 완료
)

// Sync Job Status Constants
const (
	SyncJobStatusPending   = "pending"
	SyncJobStatusRunning   = "running"
	SyncJobStatusCompleted = "completed"
	SyncJobStatusFailed    = "failed"
	SyncJobStatusCancelled = "cancelled"
)

// ========================================
// Helper Methods
// ========================================

// GetOmsStatusFromQoo10Status converts Qoo10 ShippingStatus to OMS status
func GetOmsStatusFromQoo10Status(qoo10Status string) string {
	switch qoo10Status {
	case "1": // 발송대기
		return OmsStatusNew
	case "2": // 발송완료
		return OmsStatusShipped
	case "3": // 배송완료
		return OmsStatusDelivered
	case "4": // 구매결정
		return OmsStatusCompleted
	default:
		return OmsStatusNew
	}
}

// ShouldSkipSync checks if the order should skip synchronization
// Returns true for shipped, completed, or cancelled orders
func (o *Qoo10JPOrderV2) ShouldSkipSync() bool {
	switch o.OmsStatus {
	case OmsStatusShipped, OmsStatusDelivered, OmsStatusCompleted, OmsStatusCancelled:
		return true
	default:
		return false
	}
}

// HasTrackingNumber checks if the order has a tracking number
func (o *Qoo10JPOrderV2) HasTrackingNumber() bool {
	return o.TrackingNumber != ""
}

// ToJSON converts the order to JSON bytes
func (o *Qoo10JPOrderV2) ToJSON() ([]byte, error) {
	return json.Marshal(o)
}

// ParseRawData extracts raw data as a map
func (o *Qoo10JPOrderV2) ParseRawData() (map[string]interface{}, error) {
	if o.RawData == nil {
		return nil, nil
	}
	var result map[string]interface{}
	err := json.Unmarshal(o.RawData, &result)
	return result, err
}

// SetRawData sets raw data from a map
func (o *Qoo10JPOrderV2) SetRawData(data map[string]interface{}) error {
	if data == nil {
		o.RawData = nil
		return nil
	}
	rawBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	o.RawData = rawBytes
	return nil
}

// ========================================
// Collection Result
// ========================================

// Qoo10JPCollectionResult represents the result of order collection
type Qoo10JPCollectionResult struct {
	SellerID       string    `json:"seller_id"`
	StartDate      time.Time `json:"start_date"`
	EndDate        time.Time `json:"end_date"`
	TotalCollected int       `json:"total_collected"`
	TotalSaved     int       `json:"total_saved"`
	TotalUpdated   int       `json:"total_updated"`
	TotalSkipped   int       `json:"total_skipped"`
	TotalFailed    int       `json:"total_failed"`
	DurationMs     int64     `json:"duration_ms"`
	Errors         []string  `json:"errors,omitempty"`
}
