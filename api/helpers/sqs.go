package helpers

import (
	"github.com/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

var (
	svc                 *session.Session
	delaySeconds        *int64
	maxNumberOfMessages *int64
	timeout             *int64
)

func init() {
	svc = NewSession()
	timeout = aws.Int64(10)
	delaySeconds = aws.Int64(10)
	maxNumberOfMessages = aws.Int64(10)
}

func NewSession() *session.Session {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	return sess

}

func GetQueueURL(queue *string) (*sqs.GetQueueUrlOutput, error) {
	// Create an SQS service client
	svc := sqs.New(NewSession())

	result, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: queue,
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}

func SendSQSMesssage(queueURL *string, messageAttributes map[string]*sqs.MessageAttributeValue, messageBody *string) (err error) {
	_, err = svc.SendMessage(&sqs.SendMessageInput{
		DelaySeconds:      aws.Int64(10),
		MessageAttributes: messageAttributes,
		MessageBody:       messageBody,
		QueueUrl:          queueURL,
	})

	return err

}

func ReceiveSQSMessage(queueURL *string) (msgResult *sqs.ReceiveMessageOutput, err error) {

	msgResult, err = svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl:            queueURL,
		MaxNumberOfMessages: maxNumberOfMessages,
		VisibilityTimeout:   timeout,
	})

	return msgResult, err

}

func DeleteSQSMessage(queueURL *string, messageHandle *string) (err error) {
	_, err = svc.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      queueURL,
		ReceiptHandle: messageHandle,
	})

	return err
}
