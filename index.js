const http = require('http')
const connect = require('connect')
const bodyParser = require('body-parser')
const request = require('request')

const { SLACK_WEBHOOK_URL, BOT_USERNAME, PORT } = process.env

if (!SLACK_WEBHOOK_URL || !BOT_USERNAME || !PORT) {
  console.log('Missing environment variables')
  process.exit(0)
}

const app = connect()

app.use(bodyParser.urlencoded({ extended: false }))
app.use(bodyParser.json())

app.use((req, res, next) => {
  if (req.url !== '/sms')
    return next()

  console.log(`SMS received from ${req.body.From}`)

  request.post({
    url: SLACK_WEBHOOK_URL,
    json: {
      username: BOT_USERNAME,
      text: `*From ${req.body.From}:* ${req.body.Body}`
    }
  }, (err, httpRes, body) => {
    console.log(!err ? 'SMS relayed to Slack' : 'Error: Failed to relay sms to Slack')
    res.setHeader('Content-type', 'text/xml')
    res.end(`<?xml version="1.0" encoding="UTF-8"?><Response></Response>`)
  })
})

app.use((req, res) => res.end('OK'))

http.createServer(app).listen(PORT)

console.log('Started application successfully')
