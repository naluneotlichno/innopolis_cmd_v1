// bot_service.go — реализация Telegram-бота с YAML-кнопками, Viper и мультиязычностью

package service

import (
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// ButtonConfig описывает структуру YAML-файла с кнопками и текстами
// buttons: lang -> menu -> кнопки
// texts: lang -> key -> текст кнопки
type ButtonConfig struct {
	Buttons map[string]map[string][][]string `mapstructure:"buttons"`
	Texts   map[string]map[string]string     `mapstructure:"texts"`
}

// BotService — основной сервис для работы с Telegram-ботом
// Включает сам бот, конфиг с кнопками, и карту для хранения языков пользователей
type BotService struct {
	bot       BotAPI
	cfg       *ButtonConfig
	langState sync.Map // потокобезопасная мапа: userID -> язык
}

// SyncMapUserStateManager — реализация UserStateManager с использованием sync.Map
type SyncMapUserStateManager struct {
	langState *sync.Map
}

// NewSyncMapUserStateManager создает новый менеджер пользовательских состояний
func NewSyncMapUserStateManager() *SyncMapUserStateManager {
	return &SyncMapUserStateManager{
		langState: &sync.Map{},
	}
}

// GetLang возвращает язык юзера, если установлен, иначе — язык по умолчанию
func (s *SyncMapUserStateManager) GetLang(userID int64) string {
	// Пытаемся загрузить язык пользователя из потокобезопасной мапы
	if val, ok := s.langState.Load(userID); ok {
		// Если язык найден, возвращаем его из мапы
		return val.(string)
	}
	// Если язык не найден, возвращаем язык по умолчанию из конфигурации
	return viper.GetString("bot.language")
}

// SetLang сохраняет выбранный пользователем язык в потокобезопасную мапу
func (s *SyncMapUserStateManager) SetLang(userID int64, lang string) {
	// Сохраняем язык пользователя в потокобезопасную мапу
	s.langState.Store(userID, lang)
}

// NewBotService создаёт новый экземпляр BotService и загружает конфигурацию из YAML через Viper
func NewBotService(token string) (*BotService, error) {
	log.Info().
		Str("stage", "init").
		Str("action", "create_bot_api").
		Msg("Создание экземпляра Telegram Bot API")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Error().
			Str("stage", "init").
			Str("action", "create_bot_api").
			Err(err).
			Msg("Ошибка создания Telegram Bot API")
		return nil, err
	}

	cfg, err := loadConfig()
	if err != nil {
		return nil, err
	}

	log.Info().
		Str("stage", "init").
		Str("action", "ready").
		Msg("BotService успешно инициализирован")

	return &BotService{
		bot:       bot,
		cfg:       cfg,
		langState: sync.Map{},
	}, nil
}

// NewBotServiceWithDependencies создает новый BotService с указанными зависимостями
func NewBotServiceWithDependencies(bot BotAPI, cfg *ButtonConfig) *BotService {
	return &BotService{
		bot:       bot,
		cfg:       cfg,
		langState: sync.Map{},
	}
}

// Start запускает основной цикл обработки обновлений от Telegram
func (s *BotService) Start() {
	log.Info().
		Str("stage", "start").
		Msg("Бот начал получать обновления")

	u := tgbotapi.NewUpdate(0)
	u.Timeout = viper.GetInt("bot.update_timeout")
	updates := s.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			s.HandleMessage(update.Message)
		} else if update.CallbackQuery != nil {
			s.HandleCallback(update.CallbackQuery)
		}
	}
}

// GetLang возвращает язык юзера, если установлен, иначе — язык по умолчанию
func (s *BotService) GetLang(userID int64) string {
	// Пытаемся загрузить язык пользователя из потокобезопасной мапы
	if val, ok := s.langState.Load(userID); ok {
		// Если язык найден, возвращаем его из мапы
		return val.(string)
	}
	// Если язык не найден, возвращаем язык по умолчанию из конфигурации
	return viper.GetString("bot.language")
}

// SetLang сохраняет выбранный пользователем язык в потокобезопасную мапу
func (s *BotService) SetLang(userID int64, lang string) {
	// Сохраняем язык пользователя в потокобезопасную мапу
	s.langState.Store(userID, lang)
}

// SendMessage — курьер среди функций: доставляет текст и кнопки в чат
func (s *BotService) SendMessage(chatID int64, text string, markup interface{}) {
	// Создаем новое сообщение с текстом
	msg := tgbotapi.NewMessage(chatID, text)
	// Устанавливаем клавиатуру для сообщения
	msg.ReplyMarkup = markup
	// Пытаемся отправить сообщение
	if _, err := s.bot.Send(msg); err != nil {
		log.Error().
			Str("stage", "send message").
			Int64("chat_id", chatID).
			Err(err).
			Msg("Ошибка отправки")
	}
}

// HandleCallback — шаман кнопочного мира: ловит нажатия
func (s *BotService) HandleCallback(cb *tgbotapi.CallbackQuery) {
	userID := cb.From.ID
	lang := s.GetLang(userID)

	var (
		text     string
		keyboard tgbotapi.InlineKeyboardMarkup
	)

	switch cb.Data {
	case "start_tour":
		text = "Добро пожаловать на экскурсию по Иннополису! 🏛️"
		keyboard = s.CreateKeyboard(lang, "excursion")

	case "show_routes":
		text = "Вот список доступных маршрутов 🗺️"
		keyboard = s.CreateKeyboard(lang, "main")

	case "change_language":
		if lang == "ru" {
			lang = "en"
			text = s.cfg.Texts["en"]["start_message"]
		} else {
			lang = "ru"
			text = s.cfg.Texts["ru"]["start_message"]
		}
		s.SetLang(userID, lang)
		keyboard = s.CreateKeyboard(lang, "main")

	case "menu":
		text = s.cfg.Texts[lang]["start_message"]
		keyboard = s.CreateKeyboard(lang, "main")

	case "back", "next":
		text = "Навигация по экскурсии..."
		keyboard = s.CreateKeyboard(lang, "excursion")
	}

	log.Info().
		Str("action", cb.Data).
		Str("username", cb.From.UserName).
		Int64("chat_id", cb.Message.Chat.ID).
		Msg("Нажата кнопка")

	s.SendMessage(cb.Message.Chat.ID, text, keyboard)

	// Убираем "часики" после нажатия кнопки
	_, err := s.bot.Request(tgbotapi.NewCallback(cb.ID, ""))
	if err != nil {
		log.Error().
			Str("stage", "callback response").
			Str("callback_id", cb.ID).
			Str("username", cb.From.UserName).
			Int64("chat_id", cb.Message.Chat.ID).
			Err(err).
			Msg("Ошибка при отправке callback-ответа")
	}
}

// HandleMessage — дирижёр команд
func (s *BotService) HandleMessage(msg *tgbotapi.Message) {
	// Если сообщение является командой и это команда /start
	if msg.IsCommand() && msg.Command() == "start" {
		userID := msg.From.ID
		lang := s.GetLang(userID)

		text := s.cfg.Texts[lang]["start_message"]
		keyboard := s.CreateKeyboard(lang, "main")

		log.Info().
			Str("action", "start").
			Str("username", msg.From.UserName).
			Int64("chat_id", msg.Chat.ID).
			Msg("Пользователь начал экскурсию")

		s.SendMessage(msg.Chat.ID, text, keyboard)
	}
}

// CreateKeyboard — архитектор кнопочного бытия. Собирает клавиатуру из YAML
func (s *BotService) CreateKeyboard(lang, menu string) tgbotapi.InlineKeyboardMarkup {
	// Получаем раскладку кнопок для указанного языка и меню
	layout := s.cfg.Buttons[lang][menu]
	// Получаем тексты кнопок для указанного языка
	txts := s.cfg.Texts[lang]

	// Создаем массив для строк клавиатуры
	var rows [][]tgbotapi.InlineKeyboardButton
	// Проходимся по каждой строке раскладки
	for _, row := range layout {
		// Создаем массив для кнопок в строке
		var btnRow []tgbotapi.InlineKeyboardButton
		// Проходимся по каждому ключу в строке
		for _, key := range row {
			// Получаем текст кнопки по ключу
			label := txts[key]
			// Создаем новую кнопку с текстом и ключом
			btn := tgbotapi.NewInlineKeyboardButtonData(label, key)
			// Добавляем кнопку в строку
			btnRow = append(btnRow, btn)
		}
		// Добавляем строку в массив строк клавиатуры
		rows = append(rows, btnRow)
	}
	// Возвращаем собранную клавиатуру
	return tgbotapi.NewInlineKeyboardMarkup(rows...)
}
