package config

import (
	"os"
	"strconv"
)

type Config struct {
	Server   ServerConfig
	Supabase SupabaseConfig
	Redis    RedisConfig
	N8N      N8NConfig
	Qoo10JP  Qoo10JPConfig
	Worker   WorkerConfig
	Webhook  WebhookConfig
}

type ServerConfig struct {
	Port string
	Mode string
}

type SupabaseConfig struct {
	URL        string
	AnonKey    string
	ServiceKey string
}

type RedisConfig struct {
	URL      string
	Password string
	DB       int
}

type N8NConfig struct {
	BaseURL      string
	APIKey       string
	WebhookURL   string
	AuthEndpoint string
}

type Qoo10JPConfig struct {
	APIKey    string
	APISecret string
	BaseURL   string
}

type WorkerConfig struct {
	Count int
}

type WebhookConfig struct {
	OrderCollectionURL string
}

func Load() *Config {
	return &Config{
		Server: ServerConfig{
			Port: getEnv("PORT", "8080"),
			Mode: getEnv("GIN_MODE", "debug"),
		},
		Supabase: SupabaseConfig{
			URL:        getEnv("SUPABASE_URL", ""),
			AnonKey:    getEnv("SUPABASE_ANON_KEY", ""),
			ServiceKey: getEnv("SUPABASE_SERVICE_KEY", ""),
		},
		Redis: RedisConfig{
			URL:      getEnv("REDIS_URL", "redis://localhost:6379"),
			Password: getEnv("REDIS_PASSWORD", ""),
			DB:       0,
		},
		N8N: N8NConfig{
			BaseURL:      getEnv("N8N_BASE_URL", ""),
			APIKey:       getEnv("N8N_API_KEY", ""),
			WebhookURL:   getEnv("N8N_WEBHOOK_URL", ""),
			AuthEndpoint: getEnv("N8N_AUTH_ENDPOINT", ""),
		},
		Qoo10JP: Qoo10JPConfig{
			APIKey:    getEnv("QOO10JP_API_KEY", ""),
			APISecret: getEnv("QOO10JP_API_SECRET", ""),
			BaseURL:   getEnv("QOO10JP_BASE_URL", "https://api.qoo10.jp"),
		},
		Worker: WorkerConfig{
			Count: getEnvInt("WORKER_COUNT", 3),
		},
		Webhook: WebhookConfig{
			OrderCollectionURL: getEnv("ORDER_COLLECTION_WEBHOOK_URL", ""),
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

