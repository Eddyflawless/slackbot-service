## Description
This application acts as 
- A command-line interface that can be integrated as CLI to send messages to a Slack channel.
- A webserver that intercepts messages with a valid API token to send slack messages on its behalf.

## Instrunctions to run application
- cd repo and run `go mod download` to get package dependencies

## COmmands
Running integration tests
-  cd folder and run `go test -v -tags=integration `
Running unit tests
-  cd folder and run `go test -v`


## Environmental Variables

- export MONGODB_URL= your-mongo-db-url
- export SLACK_BOT_OAUTH_TOKEN= your-slack-auth-token
- export NOTIFICATION_CHANNEL_ID= your-notification-channelId
- export AWS_ACCESS_TOKEN= your-aws-access-token
- export AWS_SECRET_KEY= your-aws-secret-key
- export AWS_REGION= your-aws-region

## SQS
1. Install terraform `brew tap tap hashicorp/tap && brew install hashicorp/tap/terraform`
2. Verify if its installed `terraform -help`

## Test Workspace
https://app.slack.com/client/T04GKRTK6Q1/C04G73T8D3M
## Resources used
1.  https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/sqs_queue
2. https://github.com/awsdocs/aws-doc-sdk-examples/blob/main/go/sqs/DeleteMessage/DeleteMessage.go
3. https://www.youtube.com/watch?v=LF3bVLiMj0I&list=PLrSqqHFS8XPaeJ71OKLoEkoBsAVUOQduP&index=6