package main

import (
	"encoding/json"
	"fmt"
	"log"
	"qoo10jp-order-go/internal/config"
	"qoo10jp-order-go/internal/models"
	"qoo10jp-order-go/pkg/redis"
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
	
	fmt.Println("=== Redis 큐 테스트 시작 ===")

	// 테스트용 N8N 메시지 생성
	testMessage := &models.N8NOrderMessage{
		AccountID:        "test-account-id",
		AccountName:      "online94",
		CertificationKey: "test-cert-key",
		Timestamp:        time.Now().Format("2006-01-02 15:04:05"),
	}

	// JSON으로 변환
	messageJSON, err := json.Marshal(testMessage)
	if err != nil {
		log.Fatalf("JSON 마샬링 실패: %v", err)
	}

	fmt.Printf("테스트 메시지 생성: %s\n", string(messageJSON))

	// Redis 큐에 메시지 추가
	queueName := "qoo10jp_order_queue"
	err = redisClient.PushToQueue(queueName, string(messageJSON))
	if err != nil {
		log.Fatalf("Redis 큐 푸시 실패: %v", err)
	}

	fmt.Printf("✅ 메시지가 큐 '%s'에 성공적으로 추가되었습니다!\n", queueName)

	// 큐 길이 확인
	length, err := redisClient.GetQueueLength(queueName)
	if err != nil {
		log.Printf("큐 길이 확인 실패: %v", err)
	} else {
		fmt.Printf("📊 현재 큐 길이: %d\n", length)
	}

	fmt.Println("\n이제 워커가 이 메시지를 처리하고 웹훅을 보낼 것입니다.")
	fmt.Println("워커 로그를 확인하여 웹훅 메시지가 전송되는지 확인하세요.")
	
	fmt.Println("\n=== 테스트 완료 ===")
}
