package main

import (
	"fmt"
	"log"
	"os"
	"shopee-order-go/pkg/webhook"
	"time"
)

func main() {
	// 환경변수에서 웹훅 URL 가져오기
	webhookURL := os.Getenv("ORDER_COLLECTION_WEBHOOK_URL")
	if webhookURL == "" {
		webhookURL = "https://n01.acsell.ai/webhook-test/qoo10-order-collect-message"
		fmt.Printf("환경변수가 설정되지 않아 기본 URL 사용: %s\n", webhookURL)
	}

	// 웹훅 클라이언트 생성
	client := webhook.NewClient(10 * time.Second)

	fmt.Println("=== 웹훅 테스트 시작 ===")

	// 1. 주문 수집 시작 웹훅 테스트
	fmt.Println("\n1. 주문 수집 시작 웹훅 테스트")
	now := time.Now()
	startMessage := fmt.Sprintf("[QOO10JP] 테스트계정 계정 주문수집 시작 - %s", now.Format("2006-01-02 15:04:05"))
	startData := map[string]interface{}{
		"platform":     "QOO10JP",
		"account_name": "테스트계정",
		"status":       "started",
		"action":       "order_collection",
		"timestamp":    now.Format("2006-01-02 15:04:05"),
		"date_range":   "2025-01-01 ~ 2025-03-31 (89일간)",
	}

	err := client.SendWebhook(webhookURL, startMessage, startData)
	if err != nil {
		log.Printf("❌ 시작 웹훅 실패: %v", err)
	} else {
		fmt.Println("✅ 시작 웹훅 성공")
	}

	// 2초 대기
	time.Sleep(2 * time.Second)

	// 2. 주문 수집 완료 웹훅 테스트
	fmt.Println("\n2. 주문 수집 완료 웹훅 테스트")
	now2 := time.Now()
	endMessage := fmt.Sprintf("[QOO10JP] 테스트계정 계정 주문수집 완료 - %s (5건 수집성공)", now2.Format("2006-01-02 15:04:05"))
	endData := map[string]interface{}{
		"platform":     "QOO10JP",
		"account_name": "테스트계정",
		"status":       "completed",
		"action":       "order_collection",
		"saved_count":  5,
		"total_count":  10,
		"timestamp":    now2.Format("2006-01-02 15:04:05"),
		"date_range":   "2025-01-01 ~ 2025-03-31 (89일간)",
	}

	err = client.SendWebhook(webhookURL, endMessage, endData)
	if err != nil {
		log.Printf("❌ 완료 웹훅 실패: %v", err)
	} else {
		fmt.Println("✅ 완료 웹훅 성공")
	}

	// 2초 대기
	time.Sleep(2 * time.Second)

	// 3. 주문 수집 실패 웹훅 테스트
	fmt.Println("\n3. 주문 수집 실패 웹훅 테스트")
	now3 := time.Now()
	failMessage := fmt.Sprintf("[QOO10JP] 테스트계정 계정 주문수집 실패 - %s (오류: API 연결 실패)", now3.Format("2006-01-02 15:04:05"))
	failData := map[string]interface{}{
		"platform":     "QOO10JP",
		"account_name": "테스트계정",
		"status":       "failed",
		"action":       "order_collection",
		"saved_count":  0,
		"total_count":  0,
		"error":        "API 연결 실패",
		"timestamp":    now3.Format("2006-01-02 15:04:05"),
		"date_range":   "2025-01-01 ~ 2025-03-31 (89일간)",
	}

	err = client.SendWebhook(webhookURL, failMessage, failData)
	if err != nil {
		log.Printf("❌ 실패 웹훅 실패: %v", err)
	} else {
		fmt.Println("✅ 실패 웹훅 성공")
	}

	// 4. 재시도 로직 테스트 (잘못된 URL로)
	fmt.Println("\n4. 재시도 로직 테스트 (잘못된 URL)")
	badURL := "https://invalid-url-for-testing.com/webhook"
	err = client.SendWebhookWithRetry(badURL, "테스트 메시지", nil, 2)
	if err != nil {
		fmt.Printf("✅ 예상된 실패: %v\n", err)
	}

	fmt.Println("\n=== 웹훅 테스트 완료 ===")
}
