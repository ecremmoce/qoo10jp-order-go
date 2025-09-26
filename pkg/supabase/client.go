package supabase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"qoo10jp-order-go/internal/config"
)

type Client struct {
	baseURL    string
	anonKey    string
	serviceKey string
	httpClient *http.Client
}

func NewClient(cfg config.SupabaseConfig) (*Client, error) {
	if cfg.URL == "" || cfg.AnonKey == "" {
		return nil, fmt.Errorf("supabase URL and anon key are required")
	}

	return &Client{
		baseURL:    cfg.URL,
		anonKey:    cfg.AnonKey,
		serviceKey: cfg.ServiceKey,
		httpClient: &http.Client{},
	}, nil
}

func (c *Client) Insert(table string, data interface{}) error {
	url := fmt.Sprintf("%s/rest/v1/%s", c.baseURL, table)
	
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	c.setHeaders(req)
	req.Header.Set("Prefer", "return=minimal")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("supabase error: %s", string(body))
	}

	return nil
}

func (c *Client) Select(table string, query string, result interface{}) error {
	url := fmt.Sprintf("%s/rest/v1/%s", c.baseURL, table)
	if query != "" {
		url += "?" + query
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	c.setHeaders(req)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("supabase error: %s", string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, result)
}

func (c *Client) Update(table string, query string, data interface{}) error {
	url := fmt.Sprintf("%s/rest/v1/%s", c.baseURL, table)
	if query != "" {
		url += "?" + query
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	c.setHeaders(req)
	req.Header.Set("Prefer", "return=minimal")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("supabase error: %s", string(body))
	}

	return nil
}

func (c *Client) Delete(table string, query string) error {
	url := fmt.Sprintf("%s/rest/v1/%s", c.baseURL, table)
	if query != "" {
		url += "?" + query
	}

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	c.setHeaders(req)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("supabase error: %s", string(body))
	}

	return nil
}

func (c *Client) setHeaders(req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("apikey", c.anonKey)
	if c.serviceKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.serviceKey)
	} else {
		req.Header.Set("Authorization", "Bearer "+c.anonKey)
	}
}
