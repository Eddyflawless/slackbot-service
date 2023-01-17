package controllers

import (
	"net/http"

	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/aws/aws-sdk-go/aws"
	sqsHelper "github.com/eddyflawless/slack-service/api/helpers/sqs"
	"github.com/gin-gonic/gin"
)

type Message struct {
	//Id      primitive.ObjectID `bson:"_id"`
	Message *string `json:"message" validate:"required"`
}

func ProcessMessages(c *gin.Context) {
	// TODO:
}

func GetMessages(c *gin.Context) {

}

func SendSlackMessage(c *gin.Context) {
	var message Message
	var err error
	if err = c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// TODO: store message into mongoDB

	// TODO: send to queue service
	input := sqsHelper.QueueMessageInput{
		MessageBody: aws.String(*message.Message),
		MessageAttributes: map[string]types.MessageAttributeValue{
			"Title": {
				DataType:    aws.String("String"),
				StringValue: aws.String("The Whistler"),
			},
		},
	}

	go func() {
		sqsHelper.HandleSendSQSMessage(input)
	}()

	c.JSON(http.StatusOK, nil)

}
