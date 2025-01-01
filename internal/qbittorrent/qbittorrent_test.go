package qbittorrent_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ilyamur/qbittorrent_notifier/internal/qbittorrent"
)

func TestClient_FetchTorrents_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"hash": "abc123", "name": "test_torrent", "state": "downloading"}]`))
	}))
	defer server.Close()

	client := qbittorrent.NewClient(server.URL, "user", "pass")
	torrents, err := client.FetchTorrents()
	if err != nil {
		t.Fatalf("ошибка при получении торрентов: %v", err)
	}

	if len(torrents) != 1 || torrents[0].Name != "test_torrent" {
		t.Errorf("некорректный результат: %+v", torrents)
	}
}
