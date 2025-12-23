package shopee

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// Client represents a Shopee API client
type Client struct {
	BaseURL    string
	PartnerID  int64
	PartnerKey string
	httpClient *http.Client
}

// OrderListRequest represents the request parameters for get_order_list
type OrderListRequest struct {
	TimeRangeField string // "create_time" or "update_time"
	TimeFrom       int64  // Unix timestamp
	TimeTo         int64  // Unix timestamp
	PageSize       int    // 1-100, default 20
	Cursor         string // For pagination
	OrderStatus    string // Optional: filter by order status
	ShopID         int64  // Shop ID
	AccessToken    string // Access token for the shop
}

// OrderListResponse represents the response from get_order_list
type OrderListResponse struct {
	Error     string `json:"error"`
	Message   string `json:"message"`
	Warning   string `json:"warning"`
	RequestID string `json:"request_id"`
	Response  struct {
		More       bool   `json:"more"`
		NextCursor string `json:"next_cursor"`
		OrderList  []struct {
			OrderSN     string `json:"order_sn"`
			OrderStatus string `json:"order_status"`
			CreateTime  int64  `json:"create_time"`
			UpdateTime  int64  `json:"update_time"`
		} `json:"order_list"`
	} `json:"response"`
}

// OrderDetailRequest represents the request parameters for get_order_detail
type OrderDetailRequest struct {
	OrderSNList            []string // List of order SNs (max 50)
	ResponseOptionalFields []string // Optional fields to include
	ShopID                 int64    // Shop ID
	AccessToken            string   // Access token for the shop
}

// OrderDetailResponse represents the response from get_order_detail
type OrderDetailResponse struct {
	Error     string `json:"error"`
	Message   string `json:"message"`
	Warning   string `json:"warning"`
	RequestID string `json:"request_id"`
	Response  struct {
		OrderList []OrderDetail `json:"order_list"`
	} `json:"response"`
}

// OrderDetail represents detailed order information
type OrderDetail struct {
	OrderSN          string `json:"order_sn"`
	OrderStatus      string `json:"order_status"`
	CreateTime       int64  `json:"create_time"`
	UpdateTime       int64  `json:"update_time"`
	BuyerUserID      int64  `json:"buyer_user_id"`
	BuyerUsername    string `json:"buyer_username"`
	RecipientAddress struct {
		Name        string `json:"name"`
		Phone       string `json:"phone"`
		FullAddress string `json:"full_address"`
		District    string `json:"district"`
		City        string `json:"city"`
		State       string `json:"state"`
		Country     string `json:"country"`
		Zipcode     string `json:"zipcode"`
	} `json:"recipient_address"`
	ItemList []struct {
		ItemID                 int64   `json:"item_id"`
		ItemName               string  `json:"item_name"`
		ItemSKU                string  `json:"item_sku"`
		ModelID                int64   `json:"model_id"`
		ModelName              string  `json:"model_name"`
		ModelSKU               string  `json:"model_sku"`
		ModelQuantityPurchased int     `json:"model_quantity_purchased"`
		ModelOriginalPrice     float64 `json:"model_original_price"`
		ModelDiscountedPrice   float64 `json:"model_discounted_price"`
	} `json:"item_list"`
	TotalAmount     float64 `json:"total_amount"`
	Currency        string  `json:"currency"`
	PaymentMethod   string  `json:"payment_method"`
	ShippingCarrier string  `json:"shipping_carrier"`
	TrackingNumber  string  `json:"tracking_no"`
}

// NewClient creates a new Shopee API client
func NewClient(baseURL string, partnerID int64, partnerKey string) *Client {
	return &Client{
		BaseURL:    baseURL,
		PartnerID:  partnerID,
		PartnerKey: partnerKey,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// generateSign generates HMAC-SHA256 signature for Shopee API
func (c *Client) generateSign(path string, timestamp int64, accessToken string, shopID int64) string {
	// Base string format: partner_id + path + timestamp + access_token + shop_id
	baseString := fmt.Sprintf("%d%s%d%s%d", c.PartnerID, path, timestamp, accessToken, shopID)

	h := hmac.New(sha256.New, []byte(c.PartnerKey))
	h.Write([]byte(baseString))

	return hex.EncodeToString(h.Sum(nil))
}

// buildAuthParams builds common authentication parameters
func (c *Client) buildAuthParams(path string, shopID int64, accessToken string) url.Values {
	timestamp := time.Now().Unix()
	sign := c.generateSign(path, timestamp, accessToken, shopID)

	params := url.Values{}
	params.Set("partner_id", strconv.FormatInt(c.PartnerID, 10))
	params.Set("timestamp", strconv.FormatInt(timestamp, 10))
	params.Set("sign", sign)
	params.Set("shop_id", strconv.FormatInt(shopID, 10))
	params.Set("access_token", accessToken)

	return params
}

// GetOrderList retrieves a list of orders
func (c *Client) GetOrderList(req OrderListRequest) (*OrderListResponse, error) {
	path := "/api/v2/order/get_order_list"
	params := c.buildAuthParams(path, req.ShopID, req.AccessToken)

	// Add request-specific parameters
	params.Set("time_range_field", req.TimeRangeField)
	params.Set("time_from", strconv.FormatInt(req.TimeFrom, 10))
	params.Set("time_to", strconv.FormatInt(req.TimeTo, 10))

	if req.PageSize > 0 {
		params.Set("page_size", strconv.Itoa(req.PageSize))
	} else {
		params.Set("page_size", "100") // Default to max
	}

	if req.Cursor != "" {
		params.Set("cursor", req.Cursor)
	}

	if req.OrderStatus != "" {
		params.Set("order_status", req.OrderStatus)
	}

	// Build full URL
	fullURL := fmt.Sprintf("%s%s?%s", c.BaseURL, path, params.Encode())

	// Make GET request
	resp, err := c.httpClient.Get(fullURL)
	if err != nil {
		return nil, fmt.Errorf("failed to call get_order_list: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var orderListResp OrderListResponse
	if err := json.Unmarshal(body, &orderListResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	if orderListResp.Error != "" {
		return nil, fmt.Errorf("API error: %s - %s", orderListResp.Error, orderListResp.Message)
	}

	return &orderListResp, nil
}

// GetOrderDetail retrieves detailed information for specific orders
func (c *Client) GetOrderDetail(req OrderDetailRequest) (*OrderDetailResponse, error) {
	path := "/api/v2/order/get_order_detail"
	params := c.buildAuthParams(path, req.ShopID, req.AccessToken)

	// Add order_sn_list
	if len(req.OrderSNList) == 0 {
		return nil, fmt.Errorf("order_sn_list cannot be empty")
	}

	orderSNs := ""
	for i, sn := range req.OrderSNList {
		if i > 0 {
			orderSNs += ","
		}
		orderSNs += sn
	}
	params.Set("order_sn_list", orderSNs)

	// Add optional fields
	if len(req.ResponseOptionalFields) > 0 {
		fields := ""
		for i, field := range req.ResponseOptionalFields {
			if i > 0 {
				fields += ","
			}
			fields += field
		}
		params.Set("response_optional_fields", fields)
	}

	// Build full URL
	fullURL := fmt.Sprintf("%s%s?%s", c.BaseURL, path, params.Encode())

	// Make GET request
	resp, err := c.httpClient.Get(fullURL)
	if err != nil {
		return nil, fmt.Errorf("failed to call get_order_detail: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var orderDetailResp OrderDetailResponse
	if err := json.Unmarshal(body, &orderDetailResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	if orderDetailResp.Error != "" {
		return nil, fmt.Errorf("API error: %s - %s", orderDetailResp.Error, orderDetailResp.Message)
	}

	return &orderDetailResp, nil
}
