package qoo10jp

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
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
	ResultCode    int     `json:"ResultCode"`
	ResultMsg     string  `json:"ResultMsg"`
	ResultObject  []Order `json:"ResultObject"`
}

type Order struct {
	OrderNo             int64   `json:"OrderNo"`
	PackNo              int64   `json:"PackNo"`
	OrderDate           string  `json:"OrderDate"`
	PaymentDate         string  `json:"PaymentDate"`
	ShippingDate        string  `json:"ShippingDate"`
	DeliveredDate       string  `json:"DeliveredDate"`
	ShippingStatus      string  `json:"ShippingStatus"`
	SellerID            string  `json:"SellerID"`
	Buyer               string  `json:"Buyer"`
	BuyerKana           string  `json:"BuyerKana"`
	BuyerTel            string  `json:"BuyerTel"`
	BuyerMobile         string  `json:"BuyerMobile"`
	BuyerEmail          string  `json:"BuyerEmail"`
	ItemNo              string  `json:"ItemNo"`
	SellerItemCode      string  `json:"SellerItemCode"`
	ItemTitle           string  `json:"ItemTitle"`
	Option              string  `json:"Option"`
	OptionCode          string  `json:"OptionCode"`
	OrderPrice          float64 `json:"OrderPrice"`
	OrderQty            int     `json:"OrderQty"`
	Discount            float64 `json:"Discount"`
	Total               float64 `json:"Total"`
	Receiver            string  `json:"Receiver"`
	ReceiverKana        string  `json:"ReceiverKana"`
	ZipCode             string  `json:"ZipCode"`
	ShippingAddress     string  `json:"ShippingAddress"`
	Address1            string  `json:"Address1"`
	Address2            string  `json:"Address2"`
	ReceiverTel         string  `json:"ReceiverTel"`
	ReceiverMobile      string  `json:"ReceiverMobile"`
	PaymentMethod       string  `json:"PaymentMethod"`
	SellerDiscount      float64 `json:"SellerDiscount"`
	Currency            string  `json:"Currency"`
	ShippingRate        float64 `json:"ShippingRate"`
	DeliveryCompany     string  `json:"DeliveryCompany"`
	PackingNo           string  `json:"PackingNo"`
	SellerDeliveryNo    string  `json:"SellerDeliveryNo"`
	SettlePrice         float64 `json:"SettlePrice"`
	TrackingNo          string  `json:"TrackingNo"`
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
	// API 파라미터 설정 (Postman 예시 기반)
	params := map[string]string{
		"QAPIVersion":      "1.0",
		"ShippingStatus":   "5", // 이미지에서 확인한 값
		"SearchStartDate":  startDate.Format("20060102"),
		"SearchEndDate":    endDate.Format("20060102"),
		"SearchCondition":  "2", // 주문일자 기준
		"v":                "1.0",
		"returnType":       "json",
		"method":           "ShippingBasic.GetShippingInfo_v3",
		"key":              c.certificationKey,
	}

	// API URL 생성 (올바른 엔드포인트 사용)
	apiURL := c.buildURL("/GMKT.INC.Front.QAPIService/ebayjapan.qapi", params)
	
	// 디버깅용 로그
	fmt.Printf("🔗 API URL: %s\n", apiURL)
	fmt.Printf("📋 API 파라미터: %+v\n", params)

	// HTTP 요청
	resp, err := c.httpClient.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("API 요청 실패: %v", err)
	}
	defer resp.Body.Close()

	fmt.Printf("📡 API 응답 상태: %d %s\n", resp.StatusCode, resp.Status)

	if resp.StatusCode != http.StatusOK {
		// 응답 본문 읽어서 오류 내용 확인
		body := make([]byte, 1024)
		n, _ := resp.Body.Read(body)
		fmt.Printf("❌ API 오류 응답: %s\n", string(body[:n]))
		return nil, fmt.Errorf("API 응답 오류: %d - %s", resp.StatusCode, string(body[:n]))
	}

	// 응답 본문을 먼저 읽어서 로그로 출력
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("응답 읽기 실패: %v", err)
	}
	
	fmt.Printf("📄 API 응답 내용: %s\n", string(body))

	// 응답 파싱
	var orderResponse OrderResponse
	if err := json.Unmarshal(body, &orderResponse); err != nil {
		return nil, fmt.Errorf("응답 파싱 실패: %v", err)
	}

	return &orderResponse, nil
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



