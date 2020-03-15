package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/grupokindynos/common/logger"
	"strconv"
)

type TelegramBot struct {
	telegramBot tgbotapi.BotAPI
	isWorking   bool
	chatID      string
}

var fileLog logger.FileLogger

func NewTelegramBot(apiKey string, chatId string) TelegramBot {
	bot, err := tgbotapi.NewBotAPI(apiKey)
	if err != nil {
		_ = logger.SingleLog("telegram_log", "telegram", "NewTelegramBot - "+err.Error())
		return TelegramBot{isWorking: false}
	}

	tb := TelegramBot{
		telegramBot: *bot,
		isWorking:   true,
		chatID:      chatId,
	}

	return tb
}

func (tb *TelegramBot) IsWorking() bool {
	return tb.isWorking
}

func (tb *TelegramBot) SendMessage(message string) {
	if !tb.IsWorking() {
		return
	}
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	chatId, _ := strconv.ParseInt(tb.chatID, 10, 64)
	_, _ = tb.telegramBot.Send(tgbotapi.NewMessage(chatId, message))
}

func (tb *TelegramBot) SendError(message string) {
	errorMessage := "**********ERROR**********\n" + message + "\n*************************"
	tb.SendMessage(errorMessage)
}
