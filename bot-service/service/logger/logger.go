package logger

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Init() {
	// Убедимся, что папка logs существует
	if err := os.MkdirAll("logs", os.ModePerm); err != nil {
		panic("Не удалось создать папку logs: " + err.Error())
	}

	// Открываем или создаем файл логов
	logFile, err := os.OpenFile("logs/bot.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic("Не удалось открыть файл логов: " + err.Error())
	}

	// JSON-лог в файл
	log.Logger = zerolog.New(logFile).With().Timestamp().Logger()
}
