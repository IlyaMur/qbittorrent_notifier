package qbittorrent

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/ilyamur/qbittorrent_notifier/internal/notifier"
)

type Torrent struct {
	Hash  string `json:"hash"`
	Name  string `json:"name"`
	State string `json:"state"`
}

type Client struct {
	url      string
	username string
	password string
}

func NewClient(url, username, password string) *Client {
	return &Client{
		url:      url,
		username: username,
		password: password,
	}
}

func (c *Client) FetchTorrents() ([]Torrent, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v2/torrents/info", c.url), nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.username, c.password)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var torrents []Torrent
	if err := json.NewDecoder(resp.Body).Decode(&torrents); err != nil {
		return nil, fmt.Errorf("не удалось распарсить данные о торрентах: %v", err)
	}

	return torrents, nil
}

func escapeMarkdownV2(str string) string {
	specialChars := []string{
		"_", "*", "[", "]", "(", ")", "~", "`", ">", "#", "+", "-", ".", "!",
	}
	for _, char := range specialChars {
		str = strings.ReplaceAll(str, char, "\\"+char)
	}
	return str
}

func MonitorDownloads(client *Client, notifier notifier.Notifier, interval int) error {
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	defer ticker.Stop()

	downloadingTorrents := make(map[string]Torrent)

	for range ticker.C {
		log.Println("Проверка статуса загрузок...")

		torrents, err := client.FetchTorrents()
		if err != nil {
			log.Printf("Ошибка при получении загрузок: %v", err)
			continue
		}

		for _, torrent := range torrents {
			if torrent.State == "downloading" {
				if _, exists := downloadingTorrents[torrent.Hash]; !exists {
					downloadingTorrents[torrent.Hash] = torrent
					log.Printf("Новая закачка: %s начала загрузку.", torrent.Name)
				}
			} else {
				if prevTorrent, exists := downloadingTorrents[torrent.Hash]; exists {
					if prevTorrent.State == "downloading" && torrent.State != "downloading" {
						if torrent.State == "stalledUP" {
							log.Printf("Загрузка торрента %s завершена, новый статус: %s", torrent.Name, torrent.State)
							message := fmt.Sprintf(
								"*Загрузка торрента* `%s` *завершена!*",
								escapeMarkdownV2(torrent.Name),
							)
							if err := notifier.SendNotification(message); err != nil {
								log.Printf("Ошибка при отправке уведомления: %v", err)
							}
						}
						delete(downloadingTorrents, torrent.Hash)
					}
				}
			}
		}
	}

	return nil
}
