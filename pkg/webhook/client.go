package webhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type Client struct {
	httpClient       *http.Client
	timeout          time.Duration
	rateLimiter      map[string]time.Time // 메시지 타입별 마지막 호출 시간
	rateLimiterMutex sync.Mutex
	minInterval      time.Duration // 최소 호출 간격 (기본: 5초)
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
		timeout:     timeout,
		rateLimiter: make(map[string]time.Time),
		minInterval: 5 * time.Second, // 최소 5초 간격
	}
}

// shouldSkipWebhook checks if the webhook should be skipped due to rate limiting
func (c *Client) shouldSkipWebhook(messageType string) bool {
	c.rateLimiterMutex.Lock()
	defer c.rateLimiterMutex.Unlock()
	
	lastCall, exists := c.rateLimiter[messageType]
	if !exists {
		// 첫 호출인 경우
		c.rateLimiter[messageType] = time.Now()
		return false
	}
	
	// 마지막 호출 이후 경과 시간 확인
	elapsed := time.Since(lastCall)
	if elapsed < c.minInterval {
		log.Printf("⏭️  웹훅 스킵 (Rate Limit): '%s' - 마지막 호출 후 %.1f초 경과 (최소: %.0f초)", 
			messageType, elapsed.Seconds(), c.minInterval.Seconds())
		return true
	}
	
	// 충분한 시간이 지났으면 호출 허용
	c.rateLimiter[messageType] = time.Now()
	return false
}

// SendWebhook sends a webhook message to the specified URL
func (c *Client) SendWebhook(webhookURL, message string, data map[string]interface{}) error {
	if webhookURL == "" {
		log.Println("웹훅 URL이 설정되지 않아 웹훅을 건너뜁니다")
		return nil
	}
	
	// Rate limiting 체크 (data에서 action 추출)
	messageType := "default"
	if data != nil {
		if action, ok := data["action"].(string); ok {
			messageType = action
		}
		if status, ok := data["status"].(string); ok {
			messageType = messageType + "_" + status
		}
	}
	
	// Rate limit 체크
	if c.shouldSkipWebhook(messageType) {
		return nil // 스킵하지만 에러는 아님
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










