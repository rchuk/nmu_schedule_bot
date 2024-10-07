package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"time"
)

func (ctx *Bot) handleCallbackQuery(query *tgbotapi.CallbackQuery, update *tgbotapi.Update) {
	data, err := strconv.Atoi(query.Data)
	if err != nil {
		ctx.bot.Send(tgbotapi.CallbackConfig{CallbackQueryID: query.ID})
		return
	}

	weekday := time.Weekday(data)
	if weekday < time.Monday || weekday > time.Saturday {
		ctx.bot.Send(tgbotapi.CallbackConfig{CallbackQueryID: query.ID})
		return
	}

	timeNow := time.Now()
	deltaDays := int(weekday) - int(timeNow.Weekday())
	schedule := ctx.scheduleManager.DaySchedule(timeNow.Add(time.Duration(deltaDays) * 24 * time.Hour))

	resp := tgbotapi.NewEditMessageTextAndMarkup(query.Message.Chat.ID, query.Message.MessageID, FormatDay(schedule), makeWeekKeyboard(&weekday))
	resp.ParseMode = tgbotapi.ModeHTML
	ctx.bot.Send(resp)
	ctx.bot.Send(tgbotapi.CallbackConfig{CallbackQueryID: query.ID})
}
