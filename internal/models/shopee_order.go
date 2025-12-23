package models

import (
	"time"
)

// ShopeeOrder represents a Shopee order stored in the database
type ShopeeOrder struct {
	ID                string    `json:"id" db:"id"`
	OrderSN           string    `json:"order_sn" db:"order_sn"`
	PlatformAccountID string    `json:"platform_account_id" db:"platform_account_id"`
	OrderStatus       string    `json:"order_status" db:"order_status"`
	CreateTime        time.Time `json:"create_time" db:"create_time"`
	UpdateTime        time.Time `json:"update_time" db:"update_time"`

	// Buyer information
	BuyerUserID   int64  `json:"buyer_user_id" db:"buyer_user_id"`
	BuyerUsername string `json:"buyer_username" db:"buyer_username"`

	// Recipient information
	RecipientName     string `json:"recipient_name" db:"recipient_name"`
	RecipientPhone    string `json:"recipient_phone" db:"recipient_phone"`
	RecipientAddress  string `json:"recipient_address" db:"recipient_address"`
	RecipientDistrict string `json:"recipient_district" db:"recipient_district"`
	RecipientCity     string `json:"recipient_city" db:"recipient_city"`
	RecipientState    string `json:"recipient_state" db:"recipient_state"`
	RecipientCountry  string `json:"recipient_country" db:"recipient_country"`
	RecipientZipcode  string `json:"recipient_zipcode" db:"recipient_zipcode"`

	// Order financial information
	TotalAmount   float64 `json:"total_amount" db:"total_amount"`
	Currency      string  `json:"currency" db:"currency"`
	PaymentMethod string  `json:"payment_method" db:"payment_method"`

	// Shipping information
	ShippingCarrier string `json:"shipping_carrier" db:"shipping_carrier"`
	TrackingNumber  string `json:"tracking_number" db:"tracking_number"`

	// Items (stored as JSON)
	ItemsJSON string `json:"items_json" db:"items_json"`

	// Metadata
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// ShopeeOrderItem represents an item in a Shopee order
type ShopeeOrderItem struct {
	ItemID                 int64   `json:"item_id"`
	ItemName               string  `json:"item_name"`
	ItemSKU                string  `json:"item_sku"`
	ModelID                int64   `json:"model_id"`
	ModelName              string  `json:"model_name"`
	ModelSKU               string  `json:"model_sku"`
	ModelQuantityPurchased int     `json:"model_quantity_purchased"`
	ModelOriginalPrice     float64 `json:"model_original_price"`
	ModelDiscountedPrice   float64 `json:"model_discounted_price"`
}

// ShopeeOrderFilter represents filter parameters for querying Shopee orders
type ShopeeOrderFilter struct {
	PlatformAccountID string     `json:"platform_account_id"`
	OrderStatus       string     `json:"order_status"`
	StartDate         *time.Time `json:"start_date"`
	EndDate           *time.Time `json:"end_date"`
	BuyerUsername     string     `json:"buyer_username"`
	OrderSN           string     `json:"order_sn"`
	Limit             int        `json:"limit"`
	Offset            int        `json:"offset"`
}
