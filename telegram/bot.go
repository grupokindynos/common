package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/grupokindynos/common/logger"
	"strconv"
)

type TelegramBot struct {
	telegramBot tgbotapi.BotAPI
	isWorking   bool
}

func NewTelegramBot(apiKey string) TelegramBot {
	bot, err := tgbotapi.NewBotAPI(apiKey)
	if err != nil {
		_ = logger.SingleLog("telegram_log", "telegram", "NewTelegramBot - "+err.Error())
		return TelegramBot{isWorking: false}
	}

	tb := TelegramBot{
		telegramBot: *bot,
		isWorking:   true,
	}

	return tb
}

func (tb *TelegramBot) IsWorking() bool {
	return tb.isWorking
}

func (tb *TelegramBot) SendMessage(message string, chat string) {
	if !tb.IsWorking() {
		return
	}
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	chatId, err := strconv.ParseInt(chat, 10, 64)
	if err != nil {
		return
	}
	_, _ = tb.telegramBot.Send(tgbotapi.NewMessage(chatId, message))
}

func (tb *TelegramBot) SendError(message string, chatId string) {
	errorMessage := "**********ERROR**********\n" + message + "\n*************************"
	tb.SendMessage(errorMessage, chatId)
}

func (tb *TelegramBot) Debug(val bool) {
	tb.telegramBot.Debug = val
}

func (tb *TelegramBot) GetUpdates() (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return tb.telegramBot.GetUpdatesChan(u)
}
