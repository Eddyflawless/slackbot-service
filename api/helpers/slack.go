package helpers

import (
	"github.com/slack-go/slack"
)

func FormatText(txt string) string {
	return "*" + txt + "*"
}

func CreateActionBlock(details slack.MsgOption) slack.MsgOption {

	// implement block with Approve and Deny button
	return details

}

func CreateMessageBlock(message string, preText ...string) slack.MsgOption {

	dividerSectionOne := slack.NewDividerBlock()
	messageDetails := message + "\n"
	preTextField := slack.NewTextBlockObject("mrkdwn", preText[0]+"\n\n", false, false)
	mainTextField := slack.NewTextBlockObject("mrkdwn", messageDetails, false, false)

	preTextSection := slack.NewSectionBlock(preTextField, nil, nil)
	mainDetailsSection := slack.NewSectionBlock(mainTextField, nil, nil)

	msg := slack.MsgOptionBlocks(
		preTextSection,
		dividerSectionOne,
		mainDetailsSection,
	)

	//preText :=  formatText("Hello! Your have a new message")
	return msg
}
