package models

import (
	"time"
)

type Order struct {
	ID              string    `json:"id" db:"id"`
	OrderNo         string    `json:"order_no" db:"order_no"`
	OrderDate       time.Time `json:"order_date" db:"order_date"`
	CustomerID      string    `json:"customer_id" db:"customer_id"`
	CustomerName    string    `json:"customer_name" db:"customer_name"`
	CustomerEmail   string    `json:"customer_email" db:"customer_email"`
	CustomerPhone   string    `json:"customer_phone" db:"customer_phone"`
	TotalAmount     float64   `json:"total_amount" db:"total_amount"`
	PaymentStatus   string    `json:"payment_status" db:"payment_status"`
	OrderStatus     string    `json:"order_status" db:"order_status"`
	ShippingAddress string    `json:"shipping_address" db:"shipping_address"`
	Items           []OrderItem `json:"items"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}

type OrderItem struct {
	ID        string  `json:"id" db:"id"`
	OrderID   string  `json:"order_id" db:"order_id"`
	ProductID string  `json:"product_id" db:"product_id"`
	ProductName string `json:"product_name" db:"product_name"`
	Quantity  int     `json:"quantity" db:"quantity"`
	Price     float64 `json:"price" db:"price"`
	TotalPrice float64 `json:"total_price" db:"total_price"`
}

type OrderFilter struct {
	StartDate *time.Time `json:"start_date"`
	EndDate   *time.Time `json:"end_date"`
	Status    string     `json:"status"`
	Limit     int        `json:"limit"`
	Offset    int        `json:"offset"`
}

