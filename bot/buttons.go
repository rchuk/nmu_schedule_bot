package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"time"
)

func makeFullWeekKeyboard() tgbotapi.InlineKeyboardMarkup {
	return makeWeekKeyboard(nil)
}

func makeWeekKeyboard(exceptWeekday *time.Weekday) tgbotapi.InlineKeyboardMarkup {
	rows := make([][]tgbotapi.InlineKeyboardButton, 0, 2)
	days := []time.Weekday{
		time.Monday,
		time.Tuesday,
		time.Wednesday,
		time.Thursday,
		time.Friday,
		time.Saturday,
	}

	var row []tgbotapi.InlineKeyboardButton
	for _, day := range days {
		if exceptWeekday != nil && day == *exceptWeekday {
			continue
		}

		if row == nil || len(row) == 3 {
			if row != nil {
				rows = append(rows, row)
			}

			row = make([]tgbotapi.InlineKeyboardButton, 0, 3)
		}

		dayStr := strconv.Itoa(int(day))
		button := tgbotapi.InlineKeyboardButton{
			Text:         getDayString(day),
			CallbackData: &dayStr,
		}
		row = append(row, button)
	}

	if len(row) != 0 {
		rows = append(rows, row)
	}

	return tgbotapi.NewInlineKeyboardMarkup(rows...)
}

func getDayString(weekday time.Weekday) string {
	switch weekday {
	case time.Monday:
		return "Пн"
	case time.Tuesday:
		return "Вт"
	case time.Wednesday:
		return "Ср"
	case time.Thursday:
		return "Чт"
	case time.Friday:
		return "Пт"
	case time.Saturday:
		return "Сб"
	case time.Sunday:
		return "Нд"
	default:
		return ""
	}
}
