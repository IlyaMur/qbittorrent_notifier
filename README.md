![Tests](https://github.com/ilyamur/qbittorrent_notifier/actions/workflows/ci.yml/badge.svg)

# qBittorrent Notifier Bot

qBittorrent Notifier Bot — это легковесный бот для отправки уведомлений о статусах загрузок в qBittorrent через Telegram.  
Бот выполнен для собственных нужд и написан без сторонних зависимостей.

## Возможности

- Мониторинг загрузок из qBittorrent.
- Уведомления о завершении загрузок.
- Простая настройка через JSON-файл.


## Как собрать и запустить

### 1. Требования
- [Go](https://golang.org/) версии 1.19 или выше.
- Аккаунт Telegram для создания бота ([инструкция](https://core.telegram.org/bots#botfather)).
- Установленный и настроенный qBittorrent с включенным API (Web UI).

### 2. Сборка
```bash
git clone https://github.com/ilyamur/qbittorrent_notifier.git
cd qbittorrent_notifier
make build
./notifier_bot
```

### 3. Настройка

Создайте файл `config.json` в той же директории, где будет находиться бинарник бота.  
Структура:

```json
{
    "qbittorrent": {
        "url": "http://<IP_АДРЕС>:<ПОРТ>",
        "username": "<ЛОГИН>",
        "password": "<ПАРОЛЬ>"
    },
    "telegram": {
        "token": "<ТЕЛЕГРАМ_ТОКЕН>",
        "chat_id": "<CHAT_ID>"
    },
    "check_interval_seconds": 20
}
```

В файле `config.json` необходимо указать следующие параметры:

| Параметр                 | Описание                                                                                     |
|--------------------------|---------------------------------------------------------------------------------------------|
| **url**                 | URL и порт хоста с qBittorrent.                                                              |
| **username/password**   | Логин и пароль для доступа к админке qBittorrent.                                            |
| **token**               | Токен для бота в Telegram.                                                                   |
| **chat_id**             | Идентификатор чата, куда бот будет отправлять уведомления.                                    |
| **check_interval_seconds** | Интервал (в секундах), с которым бот будет проверять обновления (qBittorrent не поддерживает веб-сокеты). |

## Пример работы

Вывод в stdout:

```bash
2025/01/01 09:16:11 Новая закачка: ubuntu_pack-22.04-xfce-amd64.iso начала загрузку.
...
2025/01/01 09:33:21 Проверка статуса загрузок...
2025/01/01 09:33:21 Загрузка торрента ubuntu_pack-22.04-xfce-amd64.iso завершена
```

Нотификация в telegram: 
![Нотификация в telegram:](docs/tg_notification.png)


## Запуск тестов

```bash
make test
```