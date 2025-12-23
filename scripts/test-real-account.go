package main

import (
	"encoding/json"
	"fmt"
	"log"
	"shopee-order-go/internal/config"
	"shopee-order-go/internal/models"
	"shopee-order-go/pkg/redis"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	// 환경변수 로드
	err := godotenv.Load("env")
	if err != nil {
		log.Printf("Warning: Could not load env file: %v", err)
	}

	// 설정 로드
	cfg := config.Load()

	// Redis 클라이언트 생성
	redisClient, err := redis.NewClient(cfg.Redis)
	if err != nil {
		log.Fatalf("Redis 클라이언트 생성 실패: %v", err)
	}

	fmt.Println("=== 실제 계정 테스트 시작 ===")

	// 실제 online94 계정 정보로 메시지 생성
	testMessage := &models.N8NOrderMessage{
		AccountID:        "real-online94-account", // 실제 계정이지만 테스트용으로 다른 ID 사용
		AccountName:      "online94",
		CertificationKey: "real-cert-key",
		Timestamp:        time.Now().Format("2006-01-02 15:04:05"),
	}

	// JSON으로 변환
	messageJSON, err := json.Marshal(testMessage)
	if err != nil {
		log.Fatalf("JSON 마샬링 실패: %v", err)
	}

	fmt.Printf("실제 계정 메시지 생성: %s\n", string(messageJSON))

	// Redis 큐에 메시지 추가
	queueName := "shopee_order_queue"
	err = redisClient.PushToQueue(queueName, string(messageJSON))
	if err != nil {
		log.Fatalf("Redis 큐 푸시 실패: %v", err)
	}

	fmt.Printf("✅ 실제 계정 메시지가 큐 '%s'에 성공적으로 추가되었습니다!\n", queueName)

	// 큐 길이 확인
	length, err := redisClient.GetQueueLength(queueName)
	if err != nil {
		log.Printf("큐 길이 확인 실패: %v", err)
	} else {
		fmt.Printf("📊 현재 큐 길이: %d\n", length)
	}

	fmt.Println("\n이제 워커가 실제 API 호출을 시도하고 실패 시 웹훅을 보낼 것입니다.")
	fmt.Println("워커 로그를 확인하여 실패 웹훅 메시지가 전송되는지 확인하세요.")

	fmt.Println("\n=== 테스트 완료 ===")
}
