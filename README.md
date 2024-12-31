# qBittorrent Notifier Bot

qBittorrent Notifier Bot — это простой и удобный бот для отправки уведомлений о статусах загрузок в qBittorrent через Telegram.

## Возможности

- Мониторинг загрузок из qBittorrent.
- Уведомления о завершении загрузок.
- Простая настройка через JSON-файл.

---

## Как собрать и запустить

### 1. Требования
- [Go](https://golang.org/) версии 1.19 или выше.
- Аккаунт Telegram для создания бота ([инструкция](https://core.telegram.org/bots#botfather)).
- Установленный и настроенный qBittorrent с включенным API (Web UI).

### 2. Сборка
Склонируйте проект:
```bash
git clone https://github.com/ilyamur/qbittorrent_nofier.git
cd qbittorrent_nofier
go build -o qbnotifier ./cmd/qbnotifier
```

### 3. Настройка

В файле config.json указать URL и порт хоста с qbittorrent, а так же креды для доступа и токен с chat_id для бота в телеграме.