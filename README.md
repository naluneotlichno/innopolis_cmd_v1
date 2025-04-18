# Innopolis

# 🏛️ Проект: Innopolis Bot Platform

## 📌 Описание

**Innopolis** — это распределённая система управления цепочками сообщений и блоками мультимедиа, взаимодействующая с Telegram-ботом и имеющая административную панель. В основе лежит архитектура микросервисов, написанных на Go, с использованием PostgreSQL, MinIO и Docker. Проект сопровождается строгим Git Flow и CI-практиками.

---

## 🗂️ Структура проекта

```
repo/
├── bot-service/         # 🤖 Telegram-бот (планируется)
├── chain-service/       # 🔗 Сервис управления цепочками и блоками
│   ├── cmd/             # Точка входа
│   ├── internal/        # Контроллеры, логика, обработчики
│   └── migrations/      # SQL миграции
├── docker-compose.yml   # Docker-оркестрация сервисов
├── dockerfile            # Сборка Go-приложения
├── makefile              # Команды линтинга, сборки, тестов
├── install-linter.sh     # Установка golangci-lint и конфиг
├── .golangci.yml         # Конфиг линтера
├── .gitignore            # Исключения из Git
├── docs/                # Документация по проекту и Git Flow
```

---

## ⚙️ Запуск проекта локально

```bash
make lint-setup         # Установка линтера (один раз)
make build-docker       # Сборка всех контейнеров
make run-local          # Запуск docker-compose
make lint               # Проверка линтерами
make test               # Запуск тестов
```

> ⚠️ Обязательно наличие .env файла с переменными окружения.

---

## 📁 Обзор файлов и их назначение

| Файл               | Расположение                            | Описание                                               |
| ------------------ | --------------------------------------- | ------------------------------------------------------ |
| main.go            | chain-service/cmd/chain-service/        | Точка входа в HTTP-сервер. Запускает обработчик /ping. |
| ping_handler.go    | chain-service/internal/controller/http/ | Обработчик /ping, возвращает 'pong' при GET-запросе.   |
| go.mod             | chain-service/                          | Go-модуль, указывает зависимости и имя проекта.        |
| init.up.sql        | chain-service/migrations/               | Создание таблиц message_blocks и chain_block_links.    |
| init.down.sql      | chain-service/migrations/               | Удаление таблиц (откат миграции).                      |
| message_chain.sql  | chain-service/migrations/               | Таблицы users, message_chains, перечисления и связи.   |
| readme.md (chain)  | chain-service/migrations/               | Заметка о добавлении миграций через Goose.             |
| readme.md (bot)    | bot-service/                            | Заготовка под сервис Telegram-бота (gRPC/HTTP).        |
| docker-compose.yml | repo/                                   | Запускает tg-bot, PostgreSQL и MinIO.                  |
| dockerfile         | repo/                                   | Dockerfile для сборки Go-приложения.                   |
| makefile           | repo/                                   | Команды сборки, тестов, линтинга и запуска.            |
| install-linter.sh  | repo/                                   | Установка golangci-lint и генерация конфига.           |
| .golangci.yml      | repo/                                   | Конфигурация линтера: gofmt, staticcheck, gosec и др.  |
| .gitignore         | repo/                                   | Исключения Git: IDE, .env, .db, сертификаты.           |
| gitflow.md         | docs/                                   | Полное описание ветвления, коммитов и PR.              |
| system-design.md   | docs/                                   | Тех. задание и описание бизнес-логики системы.         |

---

## 🧠 Логика работы

1. **Пользователь** взаимодействует с Telegram-ботом (в процессе разработки).
2. Бот передаёт команды в **chain-service** по HTTP/gRPC.
3. `chain-service` отвечает за:
   - создание/редактирование цепочек сообщений;
   - управление блоками мультимедиа (видео, текст, аудио);
   - сохранение и загрузку данных в PostgreSQL и MinIO.
4. Вся бизнес-логика инкапсулирована внутри `internal/controller/http`, начальная точка входа — `main.go`.
5. Миграции управляются SQL-файлами, позже будет добавлен Goose.

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
