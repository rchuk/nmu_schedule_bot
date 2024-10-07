package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (ctx *Bot) handleCommand(command string, update *tgbotapi.Update) {
	if schedule := ctx.getCommandSchedule(command); len(schedule) != 0 {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, schedule)
		msg.ParseMode = tgbotapi.ModeHTML
		ctx.bot.Send(msg)

		return
	}

	switch command {
	case "start":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привіт :)")
		ctx.bot.Send(msg)
	case "help":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Я вмію виводити розклад заданого дня або тижня! Просто перегляньте команди.")
		ctx.bot.Send(msg)
	case "day":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Оберіть день")
		msg.ReplyMarkup = makeFullWeekKeyboard()
		ctx.bot.Send(msg)
	default:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Такої команди не існує :'(")
		ctx.bot.Send(msg)
	}
}

func (ctx *Bot) getCommandSchedule(command string) string {
	switch command {
	case "today":
		schedule := ctx.scheduleManager.TodaySchedule()

		return FormatDay(schedule)
	case "tomorrow":
		schedule := ctx.scheduleManager.TomorrowSchedule()

		return FormatDay(schedule)
	case "week":
		schedule := ctx.scheduleManager.CurrWeekSchedule()

		return FormatWeek(schedule)
	case "next_week":
		schedule := ctx.scheduleManager.NextWeekSchedule()

		return FormatWeek(schedule)
	default:
		return ""
	}
}
