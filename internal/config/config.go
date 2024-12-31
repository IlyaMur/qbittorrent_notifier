package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Config struct {
	QBittorrent struct {
		URL      string `json:"url"`
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"qbittorrent"`
	Telegram struct {
		Token  string `json:"token"`
		ChatID string `json:"chat_id"`
	} `json:"telegram"`
	CheckIntervalSeconds int `json:"check_interval_seconds"`
}

func LoadConfig(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("не удалось открыть файл конфигурации: %v", err)
	}
	defer file.Close()

	var cfg Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, fmt.Errorf("не удалось распарсить конфигурацию: %v", err)
	}

	if err := validateConfig(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func validateConfig(cfg *Config) error {
	if cfg.CheckIntervalSeconds <= 0 {
		return errors.New("интервал проверки должен быть положительным числом")
	}

	if cfg.QBittorrent.URL == "" || cfg.QBittorrent.Username == "" || cfg.QBittorrent.Password == "" {
		return errors.New("необходимо указать URL, Username и Password для qBittorrent")
	}

	if cfg.Telegram.Token == "" || cfg.Telegram.ChatID == "" {
		return errors.New("необходимо указать Token и ChatID для Telegram")
	}

	return nil
}
