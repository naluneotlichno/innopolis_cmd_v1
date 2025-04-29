# Innopolis

# 🏛️ Проект: Innopolis Bot Platform

## 📌 Описание

**Innopolis** — это распределённая система управления цепочками сообщений и блоками мультимедиа, взаимодействующая с Telegram-ботом и имеющая административную панель. В основе лежит архитектура микросервисов, написанных на Go, с использованием PostgreSQL, MinIO и Docker. Проект сопровождается строгим Git Flow и CI-практиками.

## 🏗️ Структура проекта

```
Innopolis/
├── bot-service/                  # Сервис Telegram бота
│   ├── go.mod                   # Управление зависимостями
│   ├── main.go                  # Точка входа приложения
│   ├── internal/                # Внутренний код приложения
│   │   ├── handler/             # Обработчики команд и колбэков
│   │   │   ├── start_handler.go    # Обработчик /start
│   │   │   └── callback_handler.go # Обработчик колбэков
│   │   └── service/             # Основная логика сервиса
│   │       └── bot_service.go   # Реализация сервиса бота
│   └── config/                  # Конфигурация
│       └── config.go            # Управление конфигурацией
│
├── chain-service/               # Сервис управления цепочками
│   ├── go.mod                  # Управление зависимостями
│   ├── main.go                 # Точка входа приложения
│   ├── internal/               # Внутренний код приложения
│   │   ├── controller/         # Контроллеры
│   │   │   └── http/           # HTTP-контроллеры
│   │   ├── service/            # Бизнес-логика
│   │   └── repository/         # Работа с данными
│   └── migrations/             # Миграции базы данных
│
├── admin-panel/                # Административная панель
│   ├── frontend/              # Фронтенд часть
│   └── backend/               # Бэкенд часть
│
├── docs/                      # Документация
│   ├── gitflow.md            # Правила Git Flow
│   └── api/                  # API документация
│
└── docker/                    # Docker конфигурации
    ├── docker-compose.yml    # Основной compose файл
    └── services/            # Конфигурации отдельных сервисов
```

## 🧠 Логика работы

1. **Пользователь** взаимодействует с Telegram-ботом (в процессе разработки).
2. Бот передаёт команды в **chain-service** по HTTP/gRPC.
3. `chain-service` отвечает за:
   - создание/редактирование цепочек сообщений;
   - управление блоками мультимедиа (видео, текст, аудио);
   - сохранение и загрузку данных в PostgreSQL и MinIO.
4. Вся бизнес-логика инкапсулирована внутри `internal/controller/http`, начальная точка входа — `main.go`.
5. Миграции управляются SQL-файлами, позже будет добавлен Goose.

## 🛠️ Инструменты разработки

### Git команды

```bash
# Создание новой ветки
git checkout -b feature/RS-123-description

# Просмотр изменений
git status
git diff

# Добавление изменений
git add .
git add path/to/file

# Коммит изменений
git commit -m "[RS-123] Краткое описание задачи"

# Отправка изменений
git push origin feature/RS-123-description

# Обновление ветки
git pull origin main
git rebase main

# Создание PR
gh pr create --title "[RS-123] Описание задачи" --body "Описание изменений"
```

### Тестирование

```bash
# Запуск всех тестов
go test ./...

# Запуск тестов с покрытием
go test -cover ./...

# Запуск тестов с профилированием
go test -cpuprofile cpu.prof -memprofile mem.prof ./...

# Запуск конкретного теста
go test -run TestFunctionName ./...
```

### Линтинг и форматирование

```bash
# Запуск линтера
golangci-lint run

# Форматирование кода
go fmt ./...
goimports -w .

# Проверка зависимостей
go mod tidy
go mod verify
```

### Docker команды

```bash
# Сборка образа
docker build -t innopolis/bot-service .

# Запуск контейнера
docker-compose up -d

# Просмотр логов
docker-compose logs -f service-name

# Остановка контейнеров
docker-compose down
```

### Отладка

```bash
# Запуск с отладкой
dlv debug ./cmd/main.go

# Профилирование CPU
go tool pprof http://localhost:6060/debug/pprof/profile

# Профилирование памяти
go tool pprof http://localhost:6060/debug/pprof/heap
```

---

## 🔄 Git Flow

- `main` — стабильный код
- `feature/RS-<номер>-desc` — новые фичи
- `hotfix/RS-<номер>-desc` — багфиксы

Полное описание правил — в `docs/gitflow.md` 🧷

---

## 🚨 Правила командной работы

- Коммиты: `[RS-123] Краткое описание задачи`
- Ветки: `feature/RS-123-desc`
- PR: только через review и squash-merge
- Всё, что делаешь — записывай или комментируй. Контекст — это 🔑

---

## 📢 Заключение

Проект — не просто код. Это игра на выживание. Всё хрупко. Всё сломается. Всё под вопросом. Ты берёшь задачу. Ты не знаешь, чем она кончится. Но ты идёшь. Коммитишь. Проверяешь. И улыбаешься. Потому что этот хаос — твой дом.

---

**Да пребудет с тобой `go run`, брат.** 🧘‍♂️
