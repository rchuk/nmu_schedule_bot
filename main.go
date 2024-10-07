package main

import (
	"github.com/joho/godotenv"
	"log"
	"log/slog"
	"nmu_schedule_bot/api"
	"nmu_schedule_bot/bot"
	"nmu_schedule_bot/schedule"
	"os"
	"strings"
)

func main() {
	_ = godotenv.Load()

	logLevel := getLogLevel("LOG_LEVEL")
	slog.SetLogLoggerLevel(logLevel)

	botToken := getEnvVariable("TELEGRAM_BOT_TOKEN")
	username := getEnvVariable("NMU_USERNAME")
	password := getEnvVariable("NMU_PASSWORD")
	version := getEnvVariable("NMU_VERSION")
	updateCron := getEnvVariable("UPDATE_CRON")

	credentials := api.NewCredentials(username, password, version)

	scheduleManager, err := schedule.NewScheduleManager(&credentials, updateCron)
	if err != nil {
		log.Panicf("Couldn't create schedule manager %v", err)
	}
	scheduleManager.Start()

	err = bot.StartBot(botToken, scheduleManager)
	if err != nil {
		log.Panicf("Couldn't start telegram bot %v", err)
	}
}

func getLogLevel(key string) slog.Level {
	value, isSet := os.LookupEnv(key)
	if !isSet {
		return slog.LevelError
	}

	switch strings.ToLower(value) {
	case "error":
		return slog.LevelError
	case "warn":
		return slog.LevelWarn
	case "info":
		return slog.LevelInfo
	case "debug":
		return slog.LevelDebug
	case "trace":
		return slog.LevelDebug
	default:
		return slog.LevelError
	}
}

func getEnvVariable(key string) string {
	value, isSet := os.LookupEnv(key)
	if !isSet {
		log.Panicf("%v environment variable is not set", key)
	}

	return value
}
