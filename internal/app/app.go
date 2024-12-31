package app

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/ilyamur/qbittorrent_notifier/internal/config"
	"github.com/ilyamur/qbittorrent_notifier/internal/notifier"
	"github.com/ilyamur/qbittorrent_notifier/internal/qbittorrent"
)

func Run() error {
	executablePath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("не удалось получить путь к исполняемому файлу: %v", err)
	}

	dir := filepath.Dir(executablePath)

	absConfigPath := filepath.Join(dir, "config.json")

	cfg, err := config.LoadConfig(absConfigPath)

	if err != nil {
		return fmt.Errorf("не удалось загрузить конфигурацию: %v", err)
	}

	qbtClient := qbittorrent.NewClient(cfg.QBittorrent.URL, cfg.QBittorrent.Username, cfg.QBittorrent.Password)
	tgNotifier := notifier.NewTelegramNotifier(cfg.Telegram.Token, cfg.Telegram.ChatID)

	log.Println("Запуск мониторинга загрузок...")
	if err := qbittorrent.MonitorDownloads(qbtClient, tgNotifier, cfg.CheckIntervalSeconds); err != nil {
		return err
	}

	return nil
}
