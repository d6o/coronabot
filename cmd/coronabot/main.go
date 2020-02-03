package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/disiqueira/coronabot/internal/application"
	"github.com/disiqueira/coronabot/internal/domain/model"
	"github.com/disiqueira/coronabot/internal/domain/service"
	"github.com/disiqueira/coronabot/internal/infrastructure/arcgis"
	"github.com/disiqueira/coronabot/internal/infrastructure/slack"
	"github.com/disiqueira/coronabot/internal/infrastructure/telegram"

	slackLib "github.com/nlopes/slack"
	telegramLib "gopkg.in/tucnak/telebot.v2"
)

func main() {
	ctx := context.Background()

	log.Println("Starting CoronaBot")

	slackToken := os.Getenv("SLACK_TOKEN")
	telegramToken := os.Getenv("TELEGRAM_TOKEN")

	var messageSender model.MessageSender
	switch {
	case slackToken != "":
		slackChannel := os.Getenv("SLACK_CHANNEL_ID")
		if slackChannel == "" {
			log.Print("no slack channel provided. Set SLACK_CHANNEL_ID environment variable")
			os.Exit(1)
		}

		slackCon := slackLib.New(slackToken)
		messageSender = slack.New(slackCon, slackChannel)
	case telegramToken != "":
		telegramBot, err := telegramLib.NewBot(telegramLib.Settings{
			Token: telegramToken,
			Poller: &telegramLib.LongPoller{
				Timeout: 15 * time.Second,
			},
		})
		if err != nil {
			log.Print("failed to connect to telegram bot with given token")
			os.Exit(1)
		}

		messageSender = telegram.New(telegramBot)
	default:
		log.Print("no slack or telegram token provided. Set SLACK_TOKEN or TELEGRAM_TOKEN environment variable")
		os.Exit(1)
	}

	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}
	statusReporter := arcgis.New(httpClient)

	statusPerCountryToMessageService := service.NewStatusListToMessage()

	notify := application.NewNotifyService(messageSender, statusReporter, statusPerCountryToMessageService)

	notifyInterval := notifyInterval()
	notifyTimer := time.NewTicker(notifyInterval)

	for {
		select {
		case <-notifyTimer.C:
			fmt.Println(time.Now())
			if err := notify.Execute(ctx); err != nil {
				log.Printf("error while running notifier. err: %s", err)
				os.Exit(1)
			}
		case <-ctx.Done():
			log.Println("Finishing CoronaBot")
			os.Exit(0)
		}
	}
}

func notifyInterval() time.Duration {
	notifyInterval := os.Getenv("NOTIFY_INTERVAL_MINUTES")
	if notifyInterval == "" {
		log.Println("no custom interval provided, defaulting to 60 minutes. To modify set NOTIFY_INTERVAL_MINUTES environment variable")
		return 60 * time.Minute
	}

	notifyInt, err := strconv.Atoi(notifyInterval)
	if err != nil {
		log.Printf("invalid custom internal provided, defaulting to 60 minutes. value: %s error: %s\n", notifyInterval, err)
		return 60 * time.Minute
	}

	return time.Duration(notifyInt)
}
