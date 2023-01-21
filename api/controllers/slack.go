package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/eddyflawless/slack-service/api/database"
	sqsHelper "github.com/eddyflawless/slack-service/api/helpers/sqs"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

func ProcessMessages(c *gin.Context) {
	// TODO:
}

func GetMessages(c *gin.Context) {

	var messages []database.Message

	opts := options.Find()

	ctx, cancel := database.CreateTTLContext()
	defer cancel()

	cursor, err := database.FindMany("users", bson.M{}, opts)

	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &messages); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, messages)
}

func SendSlackMessage(c *gin.Context) {
	var message database.Message
	var err error
	if err = c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// TODO: store message into mongoDB
	ctx, cancel := database.CreateTTLContext()
	defer cancel()

	db := database.OpenCollection("messages")

	resultInsertionNumber, err := db.InsertOne(ctx, message)

	if err != nil {
		msg := fmt.Sprintf("User was not added")
		c.JSON(http.StatusBadRequest, gin.H{"error": msg})
		return
	}

	c.JSON(http.StatusOK, resultInsertionNumber)

	// TODO: send to queue service
	input := sqsHelper.QueueMessageInput{
		MessageBody: aws.String(message.Message),
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
