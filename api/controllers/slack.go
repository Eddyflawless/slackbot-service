package controllers

import (
	"net/http"

	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/aws/aws-sdk-go/aws"
	sqsHelper "github.com/eddyflawless/slack-service/api/helpers/sqs"
	"github.com/gin-gonic/gin"
)

func SendSlackMessage(c *gin.Context) {
	var message string
	var err error
	if err = c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input := sqsHelper.QueueMessageInput{
		MessageBody: aws.String("Information about the NY Times fiction bestseller for the week of 12/11/2016."),
		MessageAttributes: map[string]types.MessageAttributeValue{
			"Title": {
				DataType:    aws.String("String"),
				StringValue: aws.String("The Whistler"),
			},
			"Author": {
				DataType:    aws.String("String"),
				StringValue: aws.String("John Grisham"),
			},
			"WeeksOn": {
				DataType:    aws.String("Number"),
				StringValue: aws.String("6"),
			},
		},
	}

	go func() {
		sqsHelper.HandleSendSQSMessage(input)
	}()

}
