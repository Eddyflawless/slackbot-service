## Description
This application acts as 
- A command-line interface that can be integrated as CLI to send messages to a Slack channel.
- A webserver that intercepts messages with a valid API token to send slack messages on its behalf.


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