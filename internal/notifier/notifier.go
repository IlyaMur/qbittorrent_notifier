package notifier

import (
	"fmt"
	"net/http"
	"net/url"
)

type Notifier interface {
	SendNotification(message string) error
}

type TelegramNotifier struct {
	token  string
	chatID string
}

func NewTelegramNotifier(token, chatID string) *TelegramNotifier {
	return &TelegramNotifier{
		token:  token,
		chatID: chatID,
	}
}

func (n *TelegramNotifier) SendNotification(message string) error {
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", n.token)

	data := url.Values{
		"chat_id":    {n.chatID},
		"text":       {message},
		"parse_mode": {"MarkdownV2"},
	}

	resp, err := http.PostForm(apiURL, data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("ошибка телеграм-уведомления: %s", resp.Status)
	}

	return nil
}
