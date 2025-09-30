package models

import (
	"encoding/json"
	"time"
)

// OrderJob represents a job to be processed by workers
type OrderJob struct {
	ID           string    `json:"id"`
	StartDate    time.Time `json:"start_date"`
	EndDate      time.Time `json:"end_date"`
	CreatedAt    time.Time `json:"created_at"`
	Priority     int       `json:"priority"`     // Higher number = higher priority
	RetryCount   int       `json:"retry_count"`  // Current retry count
	MaxRetries   int       `json:"max_retries"`  // Maximum retry attempts
	StartDatetime time.Time `json:"start_datetime"` // Full datetime for precise filtering
	EndDatetime   time.Time `json:"end_datetime"`   // Full datetime for precise filtering
}

// N8NOrderMessage represents the message from n8n workflow
type N8NOrderMessage struct {
	AccountID        string `json:"account_id"`
	CertificationKey string `json:"certification_key"`
	AccountName      string `json:"account_name"`
	Timestamp        string `json:"timestamp"`
}

// JobResult represents the result of a processed job
type JobResult struct {
	JobID       string    `json:"job_id"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Duration    int64     `json:"duration_ms"` // Duration in milliseconds
	OrdersCount int       `json:"orders_count"`
	Success     bool      `json:"success"`
	ErrorMsg    string    `json:"error_msg,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}

// SchedulerState tracks the last execution time for scheduling
type SchedulerState struct {
	LastExecutionTime time.Time `json:"last_execution_time"`
	NextScheduleTime  time.Time `json:"next_schedule_time"`
	IsRunning         bool      `json:"is_running"`
	UpdatedAt         time.Time `json:"updated_at"`
}

func (oj *OrderJob) ToJSON() (string, error) {
	data, err := json.Marshal(oj)
	return string(data), err
}

func OrderJobFromJSON(data string) (*OrderJob, error) {
	var job OrderJob
	err := json.Unmarshal([]byte(data), &job)
	return &job, err
}

func N8NOrderMessageFromJSON(data string) (*N8NOrderMessage, error) {
	var msg N8NOrderMessage
	err := json.Unmarshal([]byte(data), &msg)
	return &msg, err
}

func (msg *N8NOrderMessage) ToJSON() (string, error) {
	data, err := json.Marshal(msg)
	return string(data), err
}
