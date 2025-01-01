package config_test

import (
	"os"
	"testing"

	"github.com/ilyamur/qbittorrent_notifier/internal/config"
)

func TestLoadConfig_Success(t *testing.T) {
	configContent := `{
		"qbittorrent": {
			"url": "http://localhost:8080",
			"username": "admin",
			"password": "password"
		},
		"telegram": {
			"token": "test_token",
			"chat_id": "123456"
		},
		"check_interval_seconds": 30
	}`
	tmpFile, err := os.CreateTemp("", "config.json")
	if err != nil {
		t.Fatalf("не удалось создать временный файл: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	_, _ = tmpFile.Write([]byte(configContent))
	tmpFile.Close()

	cfg, err := config.LoadConfig(tmpFile.Name())
	if err != nil {
		t.Fatalf("ошибка загрузки конфигурации: %v", err)
	}

	if cfg.CheckIntervalSeconds != 30 {
		t.Errorf("ожидалось 30 секунд, получили: %d", cfg.CheckIntervalSeconds)
	}
}

func TestLoadConfig_InvalidFile(t *testing.T) {
	_, err := config.LoadConfig("non_existed.json")
	if err == nil {
		t.Fatal("ожидалась ошибка при загрузке несуществующего файла")
	}
}
