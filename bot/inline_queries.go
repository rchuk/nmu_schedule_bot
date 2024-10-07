package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (ctx *Bot) handleInlineQuery(inlineQuery *tgbotapi.InlineQuery, update *tgbotapi.Update) {
	currWeekSchedule := ctx.scheduleManager.CurrWeekSchedule()
	nextWeekSchedule := ctx.scheduleManager.NextWeekSchedule()
	todaySchedule := ctx.scheduleManager.TodaySchedule()
	tomorrowSchedule := ctx.scheduleManager.TomorrowSchedule()

	results := tgbotapi.InlineConfig{
		InlineQueryID: inlineQuery.ID,
		Results: []interface{}{
			tgbotapi.NewInlineQueryResultArticleHTML("today", "Сьогодні", FormatDay(todaySchedule)),
			tgbotapi.NewInlineQueryResultArticleHTML("tomorrow", "Завтра", FormatDay(tomorrowSchedule)),
			tgbotapi.NewInlineQueryResultArticleHTML("week", "Тиждень", FormatWeek(currWeekSchedule)),
			tgbotapi.NewInlineQueryResultArticleHTML("next_week", "Наступний Тиждень", FormatWeek(nextWeekSchedule)),
		},
	}

	ctx.bot.Send(results)
}
