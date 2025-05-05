package service

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

//go:generate mockgen -source=interfaces.go -destination=mocks/mock_bot.go -package=mocks

// BotAPI определяет интерфейс для взаимодействия с API Telegram
type BotAPI interface {
	Send(c tgbotapi.Chattable) (tgbotapi.Message, error)
	GetUpdatesChan(config tgbotapi.UpdateConfig) tgbotapi.UpdatesChannel
	Request(c tgbotapi.Chattable) (*tgbotapi.APIResponse, error)
}

// ConfigLoader определяет интерфейс для загрузки конфигурации
type ConfigLoader interface {
	LoadConfig() (*ButtonConfig, error)
}

// UserStateManager определяет интерфейс для управления состоянием пользователей
type UserStateManager interface {
	GetLang(userID int64) string
	SetLang(userID int64, lang string)
}
