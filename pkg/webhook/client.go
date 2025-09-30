package webhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	httpClient *http.Client
	timeout    time.Duration
}

type WebhookMessage struct {
	Message   string            `json:"message"`
	Timestamp time.Time         `json:"timestamp"`
	Data      map[string]interface{} `json:"data,omitempty"`
}

func NewClient(timeout time.Duration) *Client {
	if timeout == 0 {
		timeout = 10 * time.Second
	}
	
	return &Client{
		httpClient: &http.Client{
			Timeout: timeout,
		},
		timeout: timeout,
	}
}

// SendWebhook sends a webhook message to the specified URL
func (c *Client) SendWebhook(webhookURL, message string, data map[string]interface{}) error {
	if webhookURL == "" {
		log.Println("웹훅 URL이 설정되지 않아 웹훅을 건너뜁니다")
		return nil
	}

	// URL 파싱 및 검증
	parsedURL, err := url.Parse(webhookURL)
	if err != nil {
		return fmt.Errorf("잘못된 웹훅 URL: %v", err)
	}

	// 메시지를 URL 쿼리 파라미터로 추가하는 방식
	if message != "" {
		query := parsedURL.Query()
		query.Set("message", message)
		parsedURL.RawQuery = query.Encode()
	}

	log.Printf("🔗 웹훅 호출: %s", parsedURL.String())

	// HTTP GET 요청 (쿼리 파라미터 방식)
	resp, err := c.httpClient.Get(parsedURL.String())
	if err != nil {
		return fmt.Errorf("웹훅 호출 실패: %v", err)
	}
	defer resp.Body.Close()

	// 응답 읽기
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("웹훅 응답 읽기 실패: %v", err)
	}

	// 상태 코드 확인
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		log.Printf("✅ 웹훅 성공 (상태: %d): %s", resp.StatusCode, string(body))
		return nil
	} else {
		return fmt.Errorf("웹훅 실패 (상태: %d): %s", resp.StatusCode, string(body))
	}
}

// SendWebhookJSON sends a webhook with JSON payload
func (c *Client) SendWebhookJSON(webhookURL string, payload WebhookMessage) error {
	if webhookURL == "" {
		log.Println("웹훅 URL이 설정되지 않아 웹훅을 건너뜁니다")
		return nil
	}

	// JSON 페이로드 생성
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("JSON 마샬링 실패: %v", err)
	}

	log.Printf("🔗 웹훅 JSON 호출: %s", webhookURL)

	// HTTP POST 요청
	resp, err := c.httpClient.Post(webhookURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("웹훅 호출 실패: %v", err)
	}
	defer resp.Body.Close()

	// 응답 읽기
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("웹훅 응답 읽기 실패: %v", err)
	}

	// 상태 코드 확인
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		log.Printf("✅ 웹훅 JSON 성공 (상태: %d): %s", resp.StatusCode, string(body))
		return nil
	} else {
		return fmt.Errorf("웹훅 JSON 실패 (상태: %d): %s", resp.StatusCode, string(body))
	}
}

// SendWebhookWithRetry sends a webhook with retry logic
func (c *Client) SendWebhookWithRetry(webhookURL, message string, data map[string]interface{}, maxRetries int) error {
	var lastErr error
	
	for i := 0; i <= maxRetries; i++ {
		err := c.SendWebhook(webhookURL, message, data)
		if err == nil {
			return nil
		}
		
		lastErr = err
		if i < maxRetries {
			waitTime := time.Duration(i+1) * time.Second
			log.Printf("⚠️ 웹훅 실패 (%d/%d), %v 후 재시도: %v", i+1, maxRetries+1, waitTime, err)
			time.Sleep(waitTime)
		}
	}
	
	return fmt.Errorf("웹훅 최대 재시도 횟수 초과: %v", lastErr)
}










