# Twilio to Slack

A basic service that relays a Twilio SMS message to a Slack webhook.

## Development

To run the app locally install Go and run this:

```
SLACK_WEBHOOK_URL="https://hooks.slack.com/services/12345/123456/3456789034567" \
BOT_USERNAME="SMS Relay bot" \
PORT="8080" \
go run main.go
```

Testing the app can be done with a curl request:

```
curl -X POST -d '{"Body": "This is the sms body text", "From": "12345"}' http://localhost:8080/sms
```

[Ngrok](https://ngrok.com/) can also be helpful for local development

## Build

Compile the binary then add to the minimal Docker image:

```
CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' .
docker build -t bufferapp/twilio-to-slack .
```

More information about this build [here](https://github.com/kelseyhightower/contributors).

### License

MIT
