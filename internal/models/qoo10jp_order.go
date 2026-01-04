package models

import (
	"time"
)

// Qoo10JP 주문 정보
type Qoo10JPOrder struct {
	ID string `json:"id" db:"id"`

	// Qoo10JP 고유 정보
	OrderNo       string    `json:"order_no" db:"order_no"`
	OrderDate     time.Time `json:"order_date" db:"order_date"`
	OrderStatus   string    `json:"order_status" db:"order_status"`
	PaymentStatus string    `json:"payment_status" db:"payment_status"`

	// 고객 정보
	BuyerID    string `json:"buyer_id" db:"buyer_id"`
	BuyerName  string `json:"buyer_name" db:"buyer_name"`
	BuyerEmail string `json:"buyer_email" db:"buyer_email"`
	BuyerPhone string `json:"buyer_phone" db:"buyer_phone"`

	// 금액 정보
	TotalAmount    float64 `json:"total_amount" db:"total_amount"`
	PaymentAmount  float64 `json:"payment_amount" db:"payment_amount"`
	DiscountAmount float64 `json:"discount_amount" db:"discount_amount"`
	ShippingFee    float64 `json:"shipping_fee" db:"shipping_fee"`
	Currency       string  `json:"currency" db:"currency"`

	// 배송 정보
	ShippingAddress string `json:"shipping_address" db:"shipping_address"`
	ShippingMethod  string `json:"shipping_method" db:"shipping_method"`
	ShippingStatus  string `json:"shipping_status" db:"shipping_status"`
	TrackingNumber  string `json:"tracking_number" db:"tracking_number"`

	// 플랫폼 계정 정보
	PlatformAccountID string `json:"platform_account_id" db:"platform_account_id"`
	SellerCode        string `json:"seller_code" db:"seller_code"`

	// 원본 데이터
	RawData interface{} `json:"raw_data" db:"raw_data"`

	// 주문 상품들 (별도 테이블에서 로드)
	Items []Qoo10JPOrderItem `json:"-" db:"-"`

	// 시스템 정보
	SyncedAt  time.Time `json:"synced_at" db:"synced_at"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// Qoo10JP 주문 상품 정보
type Qoo10JPOrderItem struct {
	ID      string `json:"id" db:"id"`
	OrderID string `json:"order_id" db:"order_id"`

	// 상품 정보
	ItemCode   string `json:"item_code" db:"item_code"`
	ItemName   string `json:"item_name" db:"item_name"`
	OptionName string `json:"option_name" db:"option_name"`

	// 수량 및 가격
	Quantity   int     `json:"quantity" db:"quantity"`
	UnitPrice  float64 `json:"unit_price" db:"unit_price"`
	TotalPrice float64 `json:"total_price" db:"total_price"`

	// 상품 상태
	ItemStatus string `json:"item_status" db:"item_status"`

	// 원본 데이터
	RawData interface{} `json:"raw_data" db:"raw_data"`

	// 시스템 정보
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// Qoo10JP 주문 필터
type Qoo10JPOrderFilter struct {
	StartDate         *time.Time `json:"start_date"`
	EndDate           *time.Time `json:"end_date"`
	OrderStatus       string     `json:"order_status"`
	PaymentStatus     string     `json:"payment_status"`
	PlatformAccountID string     `json:"platform_account_id"`
	BuyerID           string     `json:"buyer_id"`
	SellerID          string     `json:"seller_id"` // V2 호환성
	Status            string     `json:"status"`    // V2 호환성 (oms_status)
	Limit             int        `json:"limit"`
	Offset            int        `json:"offset"`
}

// Qoo10JP 주문 통계
type Qoo10JPOrderStats struct {
	TotalOrders     int     `json:"total_orders"`
	TotalAmount     float64 `json:"total_amount"`
	PendingOrders   int     `json:"pending_orders"`
	ShippedOrders   int     `json:"shipped_orders"`   // V2 호환성
	CompletedOrders int     `json:"completed_orders"`
	CancelledOrders int     `json:"cancelled_orders"`
}








