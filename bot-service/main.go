package main

import (
	"os"

	"bot-service/service/logger"
	"bot-service/service"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func main() {
	logger.Init()

	// Загружаем конфигурацию
	if err := godotenv.Load("./config/.env"); err != nil {
		log.Fatal().Err(err).Msg("Ошибка загрузки .env файла")
	}

	// Чтение токена из переменной окружения
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal().Msg("BOT_TOKEN не установлен в .env файле")
	}

	// Создаем и запускаем бота
	botService, err := service.NewBotService(token)
	if err != nil {
		log.Fatal().Err(err).Msg("Ошибка создания сервиса бота")
	}

	log.Info().Msg("Бот запущен и готов к работе!")
	
	botService.Start()
}
