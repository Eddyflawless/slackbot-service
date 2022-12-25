package main

import (
	"fmt"
	"os"

	hlp "github.com/eddyflawless/slack-service/api/helpers"
	"github.com/slack-go/slack"
)

// generate api-token for use

// send slack message direct

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

func sendSlackMessage(args []string, others ...interface{}) {

	// for now, sending a notification to general channel

	api := slack.New(slackBotToken, slack.OptionDebug(true))

	msgBlock := hlp.CreateMessageBlock(args[0], args[1])

	_, _, err := api.PostMessage(slackBotChannel, msgBlock)

	hlp.CheckErr(err)

}

func main() {

	args := os.Args[1:]
	payload := os.Args[2] // payload data
	fmt.Println(args)

	sendSlackMessage(args, payload)

}
