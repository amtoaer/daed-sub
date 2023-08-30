package internal

import (
	"encoding/json"
	"os"
)

type Config struct {
	BaseUrl        string   `json:"baseUrl" survey:"baseUrl"`
	Username       string   `json:"username" survey:"username"`
	Password       string   `json:"password" survey:"password"`
	SubscriptionId []string `json:"subscriptionId" survey:"subscriptionId"`
	Interval       int      `json:"interval" survey:"interval"`
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Load() error {
	content, err := os.ReadFile("/etc/daed-sub.conf")
	if err != nil {
		return err
	}
	return json.Unmarshal(content, c)
}

func (c *Config) Save() error {
	content, err := json.Marshal(c)
	if err != nil {
		return err
	}
	return os.WriteFile("/etc/daed-sub.conf", content, 0644)
}
