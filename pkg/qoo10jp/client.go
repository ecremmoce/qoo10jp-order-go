package qoo10jp

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// Client represents a Qoo10JP API client
type Client struct {
	apiID            string
	certificationKey string
	baseURL          string
	httpClient       *http.Client
}

// NewClient creates a new Qoo10JP API client
func NewClient(apiID, certificationKey, baseURL string) *Client {
	return &Client{
		apiID:            apiID,
		certificationKey: certificationKey,
		baseURL:          baseURL,
		httpClient:       &http.Client{Timeout: 60 * time.Second},
	}
}

// ========================================
// API Response Types
// ========================================

// OrderResponse represents the response from GetShippingInfo API
type OrderResponse struct {
	ResultCode   int     `json:"ResultCode"`
	ResultMsg    string  `json:"ResultMsg"`
	ResultObject []Order `json:"ResultObject"`
	TotalCount   int     `json:"TotalCount,omitempty"`
}

// Order represents a single order from Qoo10JP API
type Order struct {
	OrderNo          int64   `json:"OrderNo"`
	PackNo           int64   `json:"PackNo"`
	OrderDate        string  `json:"OrderDate"`
	PaymentDate      string  `json:"PaymentDate"`
	ShippingDate     string  `json:"ShippingDate"`
	DeliveredDate    string  `json:"DeliveredDate"`
	ShippingStatus   string  `json:"ShippingStatus"`
	SellerID         string  `json:"SellerID"`
	Buyer            string  `json:"Buyer"`
	BuyerKana        string  `json:"BuyerKana"`
	BuyerTel         string  `json:"BuyerTel"`
	BuyerMobile      string  `json:"BuyerMobile"`
	BuyerEmail       string  `json:"BuyerEmail"`
	ItemNo           string  `json:"ItemNo"`
	SellerItemCode   string  `json:"SellerItemCode"`
	ItemTitle        string  `json:"ItemTitle"`
	Option           string  `json:"Option"`
	OptionCode       string  `json:"OptionCode"`
	OrderPrice       float64 `json:"OrderPrice"`
	OrderQty         int     `json:"OrderQty"`
	Discount         float64 `json:"Discount"`
	Total            float64 `json:"Total"`
	Receiver         string  `json:"Receiver"`
	ReceiverKana     string  `json:"ReceiverKana"`
	ZipCode          string  `json:"ZipCode"`
	ShippingAddress  string  `json:"ShippingAddress"`
	Address1         string  `json:"Address1"`
	Address2         string  `json:"Address2"`
	ReceiverTel      string  `json:"ReceiverTel"`
	ReceiverMobile   string  `json:"ReceiverMobile"`
	PaymentMethod    string  `json:"PaymentMethod"`
	SellerDiscount   float64 `json:"SellerDiscount"`
	Currency         string  `json:"Currency"`
	ShippingRate     float64 `json:"ShippingRate"`
	DeliveryCompany  string  `json:"DeliveryCompany"`
	PackingNo        string  `json:"PackingNo"`
	SellerDeliveryNo string  `json:"SellerDeliveryNo"`
	SettlePrice      float64 `json:"SettlePrice"`
	TrackingNo       string  `json:"TrackingNo"`
}

// OrderItem represents an order item (for compatibility)
type OrderItem struct {
	ItemCode   string  `json:"ItemCode"`
	ItemName   string  `json:"ItemName"`
	Quantity   int     `json:"Quantity"`
	ItemPrice  float64 `json:"ItemPrice"`
	TotalPrice float64 `json:"TotalPrice"`
}

// OrderListRequest represents request parameters for GetOrders
type OrderListRequest struct {
	StartDate      time.Time
	EndDate        time.Time
	ShippingStatus string // 1:발송대기, 2:발송완료, 3:배송완료, 4:구매결정, 5:전체
	Page           int
	PageSize       int
	SearchType     string // 1:주문일자, 2:발송일자, 3:결제일자
}

// ========================================
// API Methods
// ========================================

// GetOrders retrieves orders using GetShippingInfo_v3 API
func (c *Client) GetOrders(startDate, endDate time.Time, page, pageSize int) (*OrderResponse, error) {
	return c.GetOrdersWithStatus(startDate, endDate, page, pageSize, "5") // 5 = 전체
}

// GetOrdersWithStatus retrieves orders with specific shipping status
func (c *Client) GetOrdersWithStatus(startDate, endDate time.Time, page, pageSize int, shippingStatus string) (*OrderResponse, error) {
	// API 파라미터 설정
	params := map[string]string{
		"QAPIVersion":     "1.0",
		"ShippingStatus":  shippingStatus,
		"SearchStartDate": startDate.Format("20060102"),
		"SearchEndDate":   endDate.Format("20060102"),
		"SearchCondition": "2", // 주문일자 기준
		"v":               "1.0",
		"returnType":      "json",
		"method":          "ShippingBasic.GetShippingInfo_v3",
		"key":             c.certificationKey,
	}

	// 페이지네이션 파라미터 추가 (API 지원 시)
	if page > 0 {
		params["Page"] = strconv.Itoa(page)
	}
	if pageSize > 0 {
		params["PageSize"] = strconv.Itoa(pageSize)
	}

	// API URL 생성
	apiURL := c.buildURL("/GMKT.INC.Front.QAPIService/ebayjapan.qapi", params)

	log.Printf("🔗 Qoo10JP API 호출: %s", apiURL[:min(len(apiURL), 200)])
	log.Printf("📋 API 파라미터: ShippingStatus=%s, 기간=%s~%s",
		shippingStatus, startDate.Format("2006-01-02"), endDate.Format("2006-01-02"))

	// HTTP 요청
	resp, err := c.httpClient.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("API 요청 실패: %w", err)
	}
	defer resp.Body.Close()

	log.Printf("📡 API 응답 상태: %d %s", resp.StatusCode, resp.Status)

	if resp.StatusCode != http.StatusOK {
		body := make([]byte, 1024)
		n, _ := resp.Body.Read(body)
		log.Printf("❌ API 오류 응답: %s", string(body[:n]))
		return nil, fmt.Errorf("API 응답 오류: %d - %s", resp.StatusCode, string(body[:n]))
	}

	// 응답 본문 읽기
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("응답 읽기 실패: %w", err)
	}

	log.Printf("📄 API 응답 길이: %d bytes", len(body))

	// 응답 파싱
	var orderResponse OrderResponse
	if err := json.Unmarshal(body, &orderResponse); err != nil {
		// 에러 시 원본 응답 일부 로깅
		snippet := string(body)
		if len(snippet) > 500 {
			snippet = snippet[:500] + "..."
		}
		log.Printf("⚠️ JSON 파싱 오류, 원본 응답: %s", snippet)
		return nil, fmt.Errorf("응답 파싱 실패: %w", err)
	}

	return &orderResponse, nil
}

// GetOrdersWithRaw retrieves orders and returns raw JSON response
func (c *Client) GetOrdersWithRaw(req OrderListRequest) (*OrderResponse, map[string]interface{}, error) {
	shippingStatus := req.ShippingStatus
	if shippingStatus == "" {
		shippingStatus = "5" // 전체
	}

	searchType := req.SearchType
	if searchType == "" {
		searchType = "2" // 주문일자 기준
	}

	// API 파라미터 설정
	params := map[string]string{
		"QAPIVersion":     "1.0",
		"ShippingStatus":  shippingStatus,
		"SearchStartDate": req.StartDate.Format("20060102"),
		"SearchEndDate":   req.EndDate.Format("20060102"),
		"SearchCondition": searchType,
		"v":               "1.0",
		"returnType":      "json",
		"method":          "ShippingBasic.GetShippingInfo_v3",
		"key":             c.certificationKey,
	}

	if req.Page > 0 {
		params["Page"] = strconv.Itoa(req.Page)
	}
	if req.PageSize > 0 {
		params["PageSize"] = strconv.Itoa(req.PageSize)
	}

	// API URL 생성
	apiURL := c.buildURL("/GMKT.INC.Front.QAPIService/ebayjapan.qapi", params)

	log.Printf("🔗 Qoo10JP API 호출 (with raw): %s...", apiURL[:min(len(apiURL), 150)])

	// HTTP 요청
	resp, err := c.httpClient.Get(apiURL)
	if err != nil {
		return nil, nil, fmt.Errorf("API 요청 실패: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body := make([]byte, 1024)
		n, _ := resp.Body.Read(body)
		return nil, nil, fmt.Errorf("API 응답 오류: %d - %s", resp.StatusCode, string(body[:n]))
	}

	// 응답 본문 읽기
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("응답 읽기 실패: %w", err)
	}

	// Raw JSON 파싱
	var rawJSON map[string]interface{}
	if err := json.Unmarshal(body, &rawJSON); err != nil {
		return nil, nil, fmt.Errorf("raw JSON 파싱 실패: %w", err)
	}

	// 구조화된 응답 파싱
	var orderResponse OrderResponse
	if err := json.Unmarshal(body, &orderResponse); err != nil {
		return nil, rawJSON, fmt.Errorf("응답 파싱 실패: %w", err)
	}

	return &orderResponse, rawJSON, nil
}

// GetAllOrderStatuses retrieves orders for all shipping statuses
// ShippingStatus: 1=발송대기, 2=발송완료, 3=배송완료, 4=구매결정
func (c *Client) GetAllOrderStatuses(startDate, endDate time.Time, page, pageSize int) (*OrderResponse, error) {
	// 전체 상태 조회 (5 = 전체)
	return c.GetOrdersWithStatus(startDate, endDate, page, pageSize, "5")
}

// GetPendingOrders retrieves orders with pending shipping status (1 = 발송대기)
func (c *Client) GetPendingOrders(startDate, endDate time.Time, page, pageSize int) (*OrderResponse, error) {
	return c.GetOrdersWithStatus(startDate, endDate, page, pageSize, "1")
}

// GetShippedOrders retrieves orders with shipped status (2 = 발송완료)
func (c *Client) GetShippedOrders(startDate, endDate time.Time, page, pageSize int) (*OrderResponse, error) {
	return c.GetOrdersWithStatus(startDate, endDate, page, pageSize, "2")
}

// GetDeliveredOrders retrieves orders with delivered status (3 = 배송완료)
func (c *Client) GetDeliveredOrders(startDate, endDate time.Time, page, pageSize int) (*OrderResponse, error) {
	return c.GetOrdersWithStatus(startDate, endDate, page, pageSize, "3")
}

// GetConfirmedOrders retrieves orders with purchase confirmed status (4 = 구매결정)
func (c *Client) GetConfirmedOrders(startDate, endDate time.Time, page, pageSize int) (*OrderResponse, error) {
	return c.GetOrdersWithStatus(startDate, endDate, page, pageSize, "4")
}

// ========================================
// Order Detail API (GetClaimInfo for detailed order info)
// ========================================

// OrderDetailResponse represents the response from order detail API
type OrderDetailResponse struct {
	ResultCode   int           `json:"ResultCode"`
	ResultMsg    string        `json:"ResultMsg"`
	ResultObject []OrderDetail `json:"ResultObject"`
}

// OrderDetail represents detailed order information
type OrderDetail struct {
	Order
	// 추가 상세 필드
	ClaimType      string `json:"ClaimType,omitempty"`
	ClaimStatus    string `json:"ClaimStatus,omitempty"`
	ClaimReason    string `json:"ClaimReason,omitempty"`
	ClaimDate      string `json:"ClaimDate,omitempty"`
	RefundAmount   float64 `json:"RefundAmount,omitempty"`
	ExchangeItemNo string `json:"ExchangeItemNo,omitempty"`
}

// GetOrderDetail retrieves detailed information for a specific order
func (c *Client) GetOrderDetail(orderNo int64, packNo int64) (*OrderDetailResponse, error) {
	params := map[string]string{
		"QAPIVersion": "1.0",
		"OrderNo":     strconv.FormatInt(orderNo, 10),
		"PackNo":      strconv.FormatInt(packNo, 10),
		"v":           "1.0",
		"returnType":  "json",
		"method":      "ShippingBasic.GetShippingDetail_v2",
		"key":         c.certificationKey,
	}

	apiURL := c.buildURL("/GMKT.INC.Front.QAPIService/ebayjapan.qapi", params)

	log.Printf("🔍 주문 상세 조회: OrderNo=%d, PackNo=%d", orderNo, packNo)

	resp, err := c.httpClient.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("API 요청 실패: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body := make([]byte, 1024)
		n, _ := resp.Body.Read(body)
		return nil, fmt.Errorf("API 응답 오류: %d - %s", resp.StatusCode, string(body[:n]))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("응답 읽기 실패: %w", err)
	}

	var detailResponse OrderDetailResponse
	if err := json.Unmarshal(body, &detailResponse); err != nil {
		return nil, fmt.Errorf("응답 파싱 실패: %w", err)
	}

	return &detailResponse, nil
}

// GetOrderDetailWithRaw retrieves detailed order info and returns raw JSON
func (c *Client) GetOrderDetailWithRaw(orderNo int64, packNo int64) (*OrderDetailResponse, map[string]interface{}, error) {
	params := map[string]string{
		"QAPIVersion": "1.0",
		"OrderNo":     strconv.FormatInt(orderNo, 10),
		"PackNo":      strconv.FormatInt(packNo, 10),
		"v":           "1.0",
		"returnType":  "json",
		"method":      "ShippingBasic.GetShippingDetail_v2",
		"key":         c.certificationKey,
	}

	apiURL := c.buildURL("/GMKT.INC.Front.QAPIService/ebayjapan.qapi", params)

	resp, err := c.httpClient.Get(apiURL)
	if err != nil {
		return nil, nil, fmt.Errorf("API 요청 실패: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body := make([]byte, 1024)
		n, _ := resp.Body.Read(body)
		return nil, nil, fmt.Errorf("API 응답 오류: %d - %s", resp.StatusCode, string(body[:n]))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("응답 읽기 실패: %w", err)
	}

	// Raw JSON 파싱
	var rawJSON map[string]interface{}
	if err := json.Unmarshal(body, &rawJSON); err != nil {
		return nil, nil, fmt.Errorf("raw JSON 파싱 실패: %w", err)
	}

	// 구조화된 응답 파싱
	var detailResponse OrderDetailResponse
	if err := json.Unmarshal(body, &detailResponse); err != nil {
		return nil, rawJSON, fmt.Errorf("응답 파싱 실패: %w", err)
	}

	return &detailResponse, rawJSON, nil
}

// ========================================
// Shipping Update API
// ========================================

// SetTrackingNumberRequest represents the request for setting tracking number
type SetTrackingNumberRequest struct {
	OrderNo         int64
	PackNo          int64
	DeliveryCompany string
	TrackingNo      string
}

// SetTrackingNumber sets the tracking number for an order
func (c *Client) SetTrackingNumber(req SetTrackingNumberRequest) error {
	params := map[string]string{
		"QAPIVersion":     "1.0",
		"OrderNo":         strconv.FormatInt(req.OrderNo, 10),
		"PackNo":          strconv.FormatInt(req.PackNo, 10),
		"DeliveryCompany": req.DeliveryCompany,
		"TrackingNo":      req.TrackingNo,
		"v":               "1.0",
		"returnType":      "json",
		"method":          "ShippingBasic.SetSellerDeliveryInfo",
		"key":             c.certificationKey,
	}

	apiURL := c.buildURL("/GMKT.INC.Front.QAPIService/ebayjapan.qapi", params)

	log.Printf("📦 송장번호 등록: OrderNo=%d, TrackingNo=%s", req.OrderNo, req.TrackingNo)

	resp, err := c.httpClient.Get(apiURL)
	if err != nil {
		return fmt.Errorf("API 요청 실패: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body := make([]byte, 1024)
		n, _ := resp.Body.Read(body)
		return fmt.Errorf("API 응답 오류: %d - %s", resp.StatusCode, string(body[:n]))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("응답 읽기 실패: %w", err)
	}

	var result struct {
		ResultCode int    `json:"ResultCode"`
		ResultMsg  string `json:"ResultMsg"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return fmt.Errorf("응답 파싱 실패: %w", err)
	}

	if result.ResultCode != 0 {
		return fmt.Errorf("송장번호 등록 실패: %s (code: %d)", result.ResultMsg, result.ResultCode)
	}

	log.Printf("✅ 송장번호 등록 성공: OrderNo=%d", req.OrderNo)
	return nil
}

// ========================================
// Helper Methods
// ========================================

// buildURL creates the full API URL with query parameters
func (c *Client) buildURL(endpoint string, params map[string]string) string {
	u, _ := url.Parse(c.baseURL + endpoint)
	q := u.Query()
	for k, v := range params {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()
	return u.String()
}

// min returns the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// GetAPIID returns the API ID
func (c *Client) GetAPIID() string {
	return c.apiID
}

// GetCertificationKey returns the certification key
func (c *Client) GetCertificationKey() string {
	return c.certificationKey
}

// GetBaseURL returns the base URL
func (c *Client) GetBaseURL() string {
	return c.baseURL
}

// ========================================
// Shipping Status Constants
// ========================================

const (
	ShippingStatusPending   = "1" // 발송대기
	ShippingStatusShipped   = "2" // 발송완료
	ShippingStatusDelivered = "3" // 배송완료
	ShippingStatusConfirmed = "4" // 구매결정
	ShippingStatusAll       = "5" // 전체
)

// GetShippingStatusName returns the name of shipping status
func GetShippingStatusName(status string) string {
	switch status {
	case ShippingStatusPending:
		return "발송대기"
	case ShippingStatusShipped:
		return "발송완료"
	case ShippingStatusDelivered:
		return "배송완료"
	case ShippingStatusConfirmed:
		return "구매결정"
	case ShippingStatusAll:
		return "전체"
	default:
		return "알수없음"
	}
}
