package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type SMSMessage struct {
	Body string
	From string
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}

func SMSWebhookHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var msg SMSMessage
	err := decoder.Decode(&msg)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	log.Printf("SMS received from %s", msg.From)

	url := os.Getenv("SLACK_WEBHOOK_URL")
	botname := os.Getenv("BOT_USERNAME")
	jsonBody := fmt.Sprintf(`{"username": "%s", "text": "*From %s:* %s"}`, botname, msg.From, msg.Body)
	postBody := []byte(jsonBody)

	_, err = http.Post(url, "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		errMsg := fmt.Sprintf("Error: Failed to relay sms to Slack")
		log.Println(errMsg)
		log.Println(err)
		http.Error(w, errMsg, 500)
		return
	}

	log.Println("SMS relayed to Slack")

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `<?xml version="1.0" encoding="UTF-8"?><Response></Response>`)
}

func init() {
	pool = x509.NewCertPool()
	pool.AppendCertsFromPEM(pemCerts)
	client = &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{RootCAs: pool}}}
}

func main() {
	port := os.Getenv("PORT")

	log.Printf("Starting server on port %s", port)

	http.HandleFunc("/", HealthCheckHandler)
	http.HandleFunc("/sms", SMSWebhookHandler)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
