package sqs

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

func GetMessages(ctx context.Context, api SQSReceiveMessageAPI, input *sqs.ReceiveMessageInput) (*sqs.ReceiveMessageOutput, error) {
	return api.ReceiveMessage(ctx, input)
}

func HandleReceiveSQSMessages() {

	client, err := GetClient()

	if err != nil {
		panic(err)
	}

	var timeout *int

	_timeout := 12 * 60 * 60
	timeout = &_timeout

	gQInput := GetQInput()
	urlResult, _ := GetQueueURL(context.TODO(), client, gQInput)

	queueURL := urlResult.QueueUrl

	gMInput := &sqs.ReceiveMessageInput{
		MessageAttributeNames: []string{
			string(types.QueueAttributeNameAll),
		},
		QueueUrl:            queueURL,
		MaxNumberOfMessages: 4,
		VisibilityTimeout:   int32(*timeout),
	}

	msgResult, err := GetMessages(context.TODO(), client, gMInput)

	if err != nil {
		fmt.Println("Got an error receiving messages:")
		fmt.Println(err)
		return
	}

	if msgResult.Messages != nil {
		fmt.Println("Message ID:     " + *msgResult.Messages[0].MessageId)
		fmt.Println("Message Handle: " + *msgResult.Messages[0].ReceiptHandle)
	} else {
		fmt.Println("No messages found")
	}
}
