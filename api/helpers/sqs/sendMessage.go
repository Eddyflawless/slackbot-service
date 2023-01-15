package sqs

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

// SendMsg sends a message to an Amazon SQS queue.
// Inputs:
//     c is the context of the method call, which includes the AWS Region.
//     api is the interface that defines the method call.
//     input defines the input arguments to the service call.
// Output:
//     If success, a SendMessageOutput object containing the result of the service call and nil.
//     Otherwise, nil and an error from the call to SendMessage.

func SendMsg(ctx context.Context, api SQSSendMessageAPI, input *sqs.SendMessageInput) (*sqs.SendMessageOutput, error) {
	return api.SendMessage(ctx, input)
}

func HandleSendSQSMessage(input QueueMessageInput) {

	env_queue := "slack-service-queue"
	var queue *string

	queue = &env_queue

	client, err := GetClient()

	// Get URL of queue
	gQInput := &sqs.GetQueueUrlInput{
		QueueName: queue,
	}

	result, err := GetQueueURL(context.TODO(), client, gQInput)
	if err != nil {
		fmt.Println("Got an error getting the queue URL:")
		fmt.Println(err)
		return
	}

	queueURL := result.QueueUrl

	sMInput := &sqs.SendMessageInput{
		DelaySeconds:      10,
		MessageAttributes: input.MessageAttributes,
		MessageBody:       input.MessageBody,
		QueueUrl:          queueURL,
	}

	//

	resp, err := SendMsg(context.TODO(), client, sMInput)
	if err != nil {
		fmt.Println("Got an error sending the message:")
		fmt.Println(err)
		return
	}

	fmt.Println("Sent message with ID: " + *resp.MessageId)
}
