package qoo10jp

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

type Client struct {
	apiID           string
	certificationKey string
	baseURL         string
	httpClient      *http.Client
}

type OrderResponse struct {
	ResultCode    string  `json:"ResultCode"`
	ResultMessage string  `json:"ResultMessage"`
	ResultObject  []Order `json:"ResultObject"`
}

type Order struct {
	OrderNo       string    `json:"OrderNo"`
	OrderDate     string    `json:"OrderDate"`
	BuyerID       string    `json:"BuyerID"`
	BuyerName     string    `json:"BuyerName"`
	BuyerEmail    string    `json:"BuyerEmail"`
	BuyerPhone    string    `json:"BuyerPhone"`
	TotalAmount   float64   `json:"TotalAmount"`
	PaymentStatus string    `json:"PaymentStatus"`
	OrderStatus   string    `json:"OrderStatus"`
	ShipAddress   string    `json:"ShipAddress"`
	OrderItems    []OrderItem `json:"OrderItems"`
}

type OrderItem struct {
	ItemCode     string  `json:"ItemCode"`
	ItemName     string  `json:"ItemName"`
	Quantity     int     `json:"Quantity"`
	ItemPrice    float64 `json:"ItemPrice"`
	TotalPrice   float64 `json:"TotalPrice"`
}

func NewClient(apiID, certificationKey, baseURL string) *Client {
	return &Client{
		apiID:           apiID,
		certificationKey: certificationKey,
		baseURL:         baseURL,
		httpClient:      &http.Client{Timeout: 30 * time.Second},
	}
}

func (c *Client) GetOrders(startDate, endDate time.Time, page, pageSize int) (*OrderResponse, error) {
	// 테스트용 Mock 데이터 반환
	mockOrders := []Order{
		{
			OrderNo:       "QOO10JP-2024-001",
			OrderDate:     startDate.Format("2006-01-02 15:04:05"),
			BuyerID:       "test_buyer_001",
			BuyerName:     "테스트 구매자",
			BuyerEmail:    "test@example.com",
			BuyerPhone:    "090-1234-5678",
			TotalAmount:   15000.0,
			PaymentStatus: "completed",
			OrderStatus:   "delivered",
			ShipAddress:   "도쿄도 시부야구 테스트 주소 1-1-1",
			OrderItems: []OrderItem{
				{
					ItemCode:   "ITEM001",
					ItemName:   "테스트 상품 1",
					Quantity:   2,
					ItemPrice:  5000.0,
					TotalPrice: 10000.0,
				},
				{
					ItemCode:   "ITEM002",
					ItemName:   "테스트 상품 2",
					Quantity:   1,
					ItemPrice:  5000.0,
					TotalPrice: 5000.0,
				},
			},
		},
		{
			OrderNo:       "QOO10JP-2024-002",
			OrderDate:     startDate.AddDate(0, 0, 1).Format("2006-01-02 15:04:05"),
			BuyerID:       "test_buyer_002",
			BuyerName:     "테스트 구매자 2",
			BuyerEmail:    "test2@example.com",
			BuyerPhone:    "090-9876-5432",
			TotalAmount:   25000.0,
			PaymentStatus: "completed",
			OrderStatus:   "processing",
			ShipAddress:   "오사카부 오사카시 테스트 주소 2-2-2",
			OrderItems: []OrderItem{
				{
					ItemCode:   "ITEM003",
					ItemName:   "테스트 상품 3",
					Quantity:   1,
					ItemPrice:  25000.0,
					TotalPrice: 25000.0,
				},
			},
		},
	}

	return &OrderResponse{
		ResultCode:    "0",
		ResultMessage: "Success",
		ResultObject:  mockOrders,
	}, nil
}

func (c *Client) generateSignature(params map[string]string) string {
	// Sort parameters by key
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Build query string
	var parts []string
	for _, k := range keys {
		parts = append(parts, fmt.Sprintf("%s=%s", k, params[k]))
	}
	queryString := strings.Join(parts, "&")

	// Generate HMAC-SHA256 signature using certification key
	h := hmac.New(sha256.New, []byte(c.certificationKey))
	h.Write([]byte(queryString))
	return hex.EncodeToString(h.Sum(nil))
}

func (c *Client) buildURL(endpoint string, params map[string]string) string {
	u, _ := url.Parse(c.baseURL + endpoint)
	q := u.Query()
	for k, v := range params {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()
	return u.String()
}



