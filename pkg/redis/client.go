package redis

import (
	"context"
	"qoo10jp-order-go/internal/config"
	"time"

	"github.com/go-redis/redis/v8"
)

type Client struct {
	rdb *redis.Client
	ctx context.Context
}

func NewClient(cfg config.RedisConfig) (*Client, error) {
	opt, err := redis.ParseURL(cfg.URL)
	if err != nil {
		return nil, err
	}

	if cfg.Password != "" {
		opt.Password = cfg.Password
	}
	opt.DB = cfg.DB

	rdb := redis.NewClient(opt)
	ctx := context.Background()

	// Test connection
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return &Client{
		rdb: rdb,
		ctx: ctx,
	}, nil
}

func (c *Client) Set(key string, value interface{}, expiration time.Duration) error {
	return c.rdb.Set(c.ctx, key, value, expiration).Err()
}

func (c *Client) Get(key string) (string, error) {
	return c.rdb.Get(c.ctx, key).Result()
}

func (c *Client) Del(key string) error {
	return c.rdb.Del(c.ctx, key).Err()
}

func (c *Client) Exists(key string) (bool, error) {
	result, err := c.rdb.Exists(c.ctx, key).Result()
	return result > 0, err
}

func (c *Client) SetNX(key string, value interface{}, expiration time.Duration) (bool, error) {
	return c.rdb.SetNX(c.ctx, key, value, expiration).Result()
}

func (c *Client) Close() error {
	return c.rdb.Close()
}

// Queue operations for order processing
func (c *Client) PushToQueue(queueName string, data interface{}) error {
	return c.rdb.LPush(c.ctx, queueName, data).Err()
}

func (c *Client) PopFromQueue(queueName string) (string, error) {
	// Use 10 second timeout for better efficiency
	result, err := c.rdb.BRPop(c.ctx, 10*time.Second, queueName).Result()
	if err != nil {
		// Check if it's a timeout (no data available)
		if err.Error() == "redis: nil" {
			return "", nil // No data available, not an error
		}
		return "", err
	}
	if len(result) < 2 {
		return "", nil
	}
	return result[1], nil
}

func (c *Client) GetQueueLength(queueName string) (int64, error) {
	return c.rdb.LLen(c.ctx, queueName).Result()
}

// Increment atomically increments a key by 1
func (c *Client) Incr(key string) (int64, error) {
	return c.rdb.Incr(c.ctx, key).Result()
}

// GetMultiple gets multiple keys at once
func (c *Client) GetMultiple(keys []string) ([]interface{}, error) {
	return c.rdb.MGet(c.ctx, keys...).Result()
}





