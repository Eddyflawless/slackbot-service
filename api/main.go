package main

import (
	"os"

	hlp "github.com/eddyflawless/slack-service/api/helpers"
	routes "github.com/eddyflawless/slack-service/api/routes"
	// "github.com/slack-go/slack"
	// mw "go-jwt/pkg/middleware"
)

var (
	// SlackBotToken is the token for the Slack bot
	slackBotToken string
	// SlackBotChannel is the channel for the Slack bot
	slackBotChannel string

	port string
)

func init() {

	hlp.LoadEnv()

	slackBotToken = os.Getenv("SLACK_BOT_OAUTH_TOKEN")
	slackBotChannel = os.Getenv("NOTIFICATION_CHANNEL_ID")
	port = os.Getenv("PORT")

	if port == "" {
		port = "9090"
	}
}

func startApp() {
	router := routes.SetUpRoutes()
	router.Run(":" + port) // listen and serve on 0.0.0.0:{PORT}
}

func main() {

	startApp()

}
