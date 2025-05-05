// Тесты для бота — как медитация стоика: проверяют всё, но ничему не удивляются
package tests

import (
	"bot-service/service"
	"bot-service/service/mocks"
	"testing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestBotService_sendMessage(t *testing.T) {
	// Создаем контроллер для мока
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Создаем мок для BotAPI
	mockBot := mocks.NewMockBotAPI(ctrl)

	// Создаем тестовые данные
	chatID := int64(123456789)
	text := "Привет, это тест! 🚀"
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Кнопка", "data"),
		),
	)

	// Настраиваем ожидания для мока
	// Любое сообщение, отправленное в чат с chatID и текстом text
	mockBot.EXPECT().
		Send(gomock.Any()). // Важно! В реальном коде лучше проверить, что мы отправляем именно то, что ожидаем
		Return(tgbotapi.Message{}, nil).
		Times(1)

	// Создаем тестируемый сервис с моком вместо реального бота
	cfg := &service.ButtonConfig{
		Buttons: map[string]map[string][][]string{},
		Texts:   map[string]map[string]string{},
	}
	botService := service.NewBotServiceWithDependencies(mockBot, cfg)

	// Вызываем тестируемый метод
	botService.SendMessage(chatID, text, keyboard)

	// Проверка выполняется автоматически через mockBot.EXPECT()
}

func TestBotService_handleCallback(t *testing.T) {
	// Создаем контроллер для мока
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Создаем мок для BotAPI
	mockBot := mocks.NewMockBotAPI(ctrl)

	// Создаем тестовые данные для колбека "change_language"
	userID := int64(123456789)
	callbackData := "change_language"

	callback := &tgbotapi.CallbackQuery{
		ID: "callback123",
		From: &tgbotapi.User{
			ID:       userID,
			UserName: "testuser",
		},
		Message: &tgbotapi.Message{
			Chat: &tgbotapi.Chat{
				ID: userID,
			},
		},
		Data: callbackData,
	}

	// Настраиваем конфигурацию для теста
	cfg := &service.ButtonConfig{
		Buttons: map[string]map[string][][]string{
			"ru": {
				"main": {{"menu", "change_language"}},
			},
			"en": {
				"main": {{"menu", "change_language"}},
			},
		},
		Texts: map[string]map[string]string{
			"ru": {
				"start_message":   "Привет на русском",
				"menu":            "Меню",
				"change_language": "Сменить язык",
			},
			"en": {
				"start_message":   "Hello in English",
				"menu":            "Menu",
				"change_language": "Change language",
			},
		},
	}

	// Ожидаем две отправки: сообщение с клавиатурой и ответ на колбек
	mockBot.EXPECT().
		Send(gomock.Any()).
		Return(tgbotapi.Message{}, nil).
		Times(1)

	mockBot.EXPECT().
		Request(gomock.Any()).
		Return(&tgbotapi.APIResponse{}, nil).
		Times(1)

	// Создаем тестируемый сервис с моком вместо реального бота
	botService := service.NewBotServiceWithDependencies(mockBot, cfg)

	// Вызываем тестируемый метод
	botService.HandleCallback(callback)

	// Проверки через EXPECT уже произошли
	// Можно добавить проверку, что язык пользователя изменился,
	// но для этого надо бы переработать код, чтобы использовать UserStateManager
}

func TestBotService_handleMessage_Start(t *testing.T) {
	// Создаем контроллер для мока
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Создаем мок для BotAPI
	mockBot := mocks.NewMockBotAPI(ctrl)

	// Создаем тестовое сообщение
	message := &tgbotapi.Message{
		MessageID: 123,
		From: &tgbotapi.User{
			ID:       int64(123456789),
			UserName: "testuser",
		},
		Chat: &tgbotapi.Chat{
			ID: int64(123456789),
		},
		Text:     "/start",
		Entities: []tgbotapi.MessageEntity{{Type: "bot_command", Offset: 0, Length: 6}},
	}

	// Настраиваем конфигурацию для теста
	cfg := &service.ButtonConfig{
		Buttons: map[string]map[string][][]string{
			"ru": {
				"main": {{"menu", "change_language"}},
			},
		},
		Texts: map[string]map[string]string{
			"ru": {
				"start_message":   "Привет на русском",
				"menu":            "Меню",
				"change_language": "Сменить язык",
			},
		},
	}

	// Ожидаем одну отправку сообщения
	mockBot.EXPECT().
		Send(gomock.Any()).
		Return(tgbotapi.Message{}, nil).
		Times(1)

	// Создаем тестируемый сервис с моком вместо реального бота
	botService := service.NewBotServiceWithDependencies(mockBot, cfg)

	// Вызываем тестируемый метод
	botService.HandleMessage(message)

	// Проверки через EXPECT уже произошли
}

func TestBotService_createKeyboard(t *testing.T) {
	// Настраиваем конфигурацию для теста
	cfg := &service.ButtonConfig{
		Buttons: map[string]map[string][][]string{
			"ru": {
				"main": {{"button1", "button2"}, {"button3"}},
			},
		},
		Texts: map[string]map[string]string{
			"ru": {
				"button1": "Кнопка 1",
				"button2": "Кнопка 2",
				"button3": "Кнопка 3",
			},
		},
	}

	// Создаем контроллер для мока (хотя для этого теста он не нужен)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Создаем мок для BotAPI (хотя для этого теста он не используется)
	mockBot := mocks.NewMockBotAPI(ctrl)

	// Создаем тестируемый сервис
	botService := service.NewBotServiceWithDependencies(mockBot, cfg)

	// Вызываем тестируемый метод
	keyboard := botService.CreateKeyboard("ru", "main")

	// Проверяем структуру клавиатуры
	assert.Equal(t, 2, len(keyboard.InlineKeyboard), "Должно быть 2 ряда кнопок")
	assert.Equal(t, 2, len(keyboard.InlineKeyboard[0]), "В первом ряду должно быть 2 кнопки")
	assert.Equal(t, 1, len(keyboard.InlineKeyboard[1]), "Во втором ряду должна быть 1 кнопка")

	// Проверяем только тексты кнопок
	assert.Equal(t, "Кнопка 1", keyboard.InlineKeyboard[0][0].Text, "Текст первой кнопки должен быть 'Кнопка 1'")
	assert.Equal(t, "Кнопка 2", keyboard.InlineKeyboard[0][1].Text, "Текст второй кнопки должен быть 'Кнопка 2'")
	assert.Equal(t, "Кнопка 3", keyboard.InlineKeyboard[1][0].Text, "Текст третьей кнопки должен быть 'Кнопка 3'")
}
