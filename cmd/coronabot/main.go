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
	"github.com/disiqueira/coronabot/internal/domain/service"
	"github.com/disiqueira/coronabot/internal/infrastructure/arcgis"
	"github.com/disiqueira/coronabot/internal/infrastructure/slack"
	slackLib "github.com/nlopes/slack"
)

func main() {
	ctx := context.Background()

	log.Println("Starting CoronaBot")

	slackToken := os.Getenv("SLACK_TOKEN")
	if slackToken == "" {
		log.Print("no slack token provided. Set SLACK_TOKEN environment variable")
		os.Exit(1)
	}

	slackChannel := os.Getenv("SLACK_CHANNEL_ID")
	if slackChannel == "" {
		log.Print("no slack channel provided. Set SLACK_CHANNEL_ID environment variable")
		os.Exit(1)
	}

	slackCon := slackLib.New(slackToken)
	slackClient := slack.New(slackCon, slackChannel)

	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}
	statusReporter := arcgis.New(httpClient)

	statusPerCountryToMessageService := service.NewStatusListToMessage()

	notify := application.NewNotifyService(slackClient, statusReporter, statusPerCountryToMessageService)

	notifyInterval := notifyInterval()
	notifyTimer := time.NewTicker(notifyInterval * time.Minute)

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
