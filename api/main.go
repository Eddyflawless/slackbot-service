package main

import (
	"os"

	hlp "github.com/eddyflawless/slack-service/api/helpers"
	// "github.com/slack-go/slack"
	// mw "go-jwt/pkg/middleware"
)

var (
	// SlackBotToken is the token for the Slack bot
	slackBotToken string
	// SlackBotChannel is the channel for the Slack bot
	slackBotChannel string
)

func init() {

	hlp.LoadEnv()

	slackBotToken = os.Getenv("SLACK_BOT_OAUTH_TOKEN")
	slackBotChannel = os.Getenv("NOTIFICATION_CHANNEL_ID")
}

func startApp() {

}

func main() {

	startApp()

}
