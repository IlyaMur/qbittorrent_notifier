package main

import (
	"log"

	"github.com/ilyamur/qbittorrent_notifier/internal/app"
)

func main() {
	// Запуск приложения
	if err := app.Run(); err != nil {
		log.Fatalf("Ошибка при запуске приложения: %v", err)
	}
}
