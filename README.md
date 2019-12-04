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
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{"Body": "Here you go 🌺", "From": "1-800-FLOWERS"}' \
  http://localhost:8080/sms
```

[Ngrok](https://ngrok.com/) can also be helpful for local development

## Docker build and deploy

Build and push the image to Docker hub:

```
docker build -t bufferapp/twilio-to-slack:1.0.0 .
docker push bufferapp/twilio-to-slack:1.0.0 .
```

Test your container locally by running this command:

```
docker run -p 8080:8080 --env-file .env bufferapp/twilio-to-slack:1.0.0
```

### License

MIT
