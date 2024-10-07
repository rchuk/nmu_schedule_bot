package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
	"nmu_schedule_bot/schedule"
)

type Bot struct {
	bot             *tgbotapi.BotAPI
	scheduleManager *schedule.ScheduleManager
}

func StartBot(botToken string, scheduleManager *schedule.ScheduleManager) error {
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		return err
	}

	slog.Info("Authorized on account", "username", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	ctx := Bot{
		bot:             bot,
		scheduleManager: scheduleManager,
	}
	ctx.updateLoop(bot.GetUpdatesChan(u))

	return nil
}

func (ctx *Bot) updateLoop(updates tgbotapi.UpdatesChannel) {
	slog.Info("Starting update loop")

	for update := range updates {
		ctx.handleUpdate(&update)
	}
}

func (ctx *Bot) handleUpdate(update *tgbotapi.Update) {
	if update.Message != nil {
		if command := update.Message.Command(); len(command) != 0 {
			ctx.handleCommand(command, update)
		} else {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Я нічого не розумію")
			ctx.bot.Send(msg)
		}
	} else if update.InlineQuery != nil {
		ctx.handleInlineQuery(update.InlineQuery, update)
	} else if update.CallbackQuery != nil {
		ctx.handleCallbackQuery(update.CallbackQuery, update)
	}
}
