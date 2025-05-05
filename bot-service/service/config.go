package service

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// loadConfig загружает и разбирает конфигурацию кнопок из YAML-файла
func loadConfig() (*ButtonConfig, error) {
	log.Info().
		Str("stage", "init").
		Str("action", "load_config").
		Str("config_path", "./config/buttons.yml").
		Msg("Загрузка конфигурации кнопок")

	viper.SetConfigName("buttons")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		log.Error().
			Str("stage", "init").
			Str("action", "load_config").
			Str("config_path", "./config/buttons.yml").
			Err(err).
			Msg("Ошибка загрузки конфигурации кнопок")
		return nil, err
	}

	var cfg ButtonConfig
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Error().
			Str("stage", "init").
			Str("action", "unmarshal_config").
			Str("config_path", "./config/buttons.yml").
			Err(err).
			Msg("Ошибка разбора конфигурации кнопок")
		return nil, err
	}

	return &cfg, nil
}
