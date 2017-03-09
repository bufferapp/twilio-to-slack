# Twilio to Slack

A basic service that relays a Twilio SMS message to a Slack webhook.

## Development

To run the app locally install Node and [nodemon](https://www.npmjs.com/package/nodemon), then
install all the dependencies:

```
npm install
```

```
SLACK_WEBHOOK_URL="https://hooks.slack.com/services/12345/123456/3456789034567" \
BOT_USERNAME="SMS Relay bot" \
PORT="8080" \
nodemon index.js
```

Testing the app can be done with a curl request:

```
curl -X POST -d '{"Body": "This is the sms body text", "From": "12345"}' http://localhost:8080/sms
```

[Ngrok](https://ngrok.com/) can also be helpful for local development

## Docker build and deploy

Build and push the image to Docker hub:

```
docker build -t bufferapp/twilio-to-slack:1.0.0 .
docker push bufferapp/twilio-to-slack:1.0.0 .
```

### License

MIT
