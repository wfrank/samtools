package isam

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type CPUStats struct {
	Idle   float32 `json:"idle_cpu,string"`
	User   float32 `json:"user_cpu,string"`
	System float32 `json:"system_cpu,string"`
}

type MemoryStats struct {
	Free  float32 `json:"free,string"`
	Used  float32 `json:"used,string"`
	Total float32 `json:"total,string"`
}

type PartitionInfo struct {
	Size  float32 `json:"size,string"`
	Used  float32 `json:"used,string"`
	Avail float32 `json:"avail,string"`
}

type StorageStats struct {
	Boot PartitionInfo `json:"boot"`
	Root PartitionInfo `json:"root"`
}

type SystemStats struct {
	CPU     CPUStats     `json:"cpu"`
	Memory  MemoryStats  `json:"memory"`
	Storage StorageStats `json:"storage"`
}

func (c *Client) PollCPUStats(stats *SystemStats) error {
	req, err := http.NewRequest("GET", "https://"+c.Host+"/statistics/systems/cpu.json", nil)
	if err != nil {
		return fmt.Errorf("error creating http request: %v", err)
	}

	q := url.Values{}
	q.Add("timespan", "1h")
	req.URL.RawQuery = q.Encode()

	req.Header.Add("Accept", "application/json")

	req.SetBasicAuth(c.User, c.Pass)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error sending http request: %v", err)
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&stats.CPU); err != nil {
		return fmt.Errorf("error decoding response body: %v", err)
	}

	return nil
}

func (c *Client) PollMemoryStats(stats *SystemStats) error {
	req, err := http.NewRequest("GET", "https://"+c.Host+"/statistics/systems/memory.json", nil)
	if err != nil {
		return fmt.Errorf("error creating http request: %v", err)
	}

	q := url.Values{}
	q.Add("timespan", "1h")
	req.URL.RawQuery = q.Encode()

	req.Header.Add("Accept", "application/json")

	req.SetBasicAuth(c.User, c.Pass)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error sending http request: %v", err)
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&stats.Memory); err != nil {
		return fmt.Errorf("error decoding response body: %v", err)
	}

	return nil
}

func (c *Client) PollStorageStats(stats *SystemStats) error {
	req, err := http.NewRequest("GET", "https://"+c.Host+"/statistics/systems/storage.json", nil)
	if err != nil {
		return fmt.Errorf("error creating http request: %v", err)
	}

	q := url.Values{}
	q.Add("timespan", "1h")
	req.URL.RawQuery = q.Encode()

	req.Header.Add("Accept", "application/json")

	req.SetBasicAuth(c.User, c.Pass)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error sending http request: %v", err)
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&stats.Storage); err != nil {
		return fmt.Errorf("error decoding response body: %v", err)
	}

	return nil
}

func (c *Client) PollSystemStats() (*SystemStats, error) {

	var stats SystemStats

	var err error
	if err = c.PollCPUStats(&stats); err != nil {
		return nil, fmt.Errorf("error polling CPU stats from %s: %v", c.Host, err)
	}
	if err = c.PollMemoryStats(&stats); err != nil {
		return nil, fmt.Errorf("error polling Memory stats from %s: %v", c.Host, err)
	}
	if err = c.PollStorageStats(&stats); err != nil {
		return nil, fmt.Errorf("error polling Storage stats from %s: %v", c.Host, err)
	}

	return &stats, nil
}
