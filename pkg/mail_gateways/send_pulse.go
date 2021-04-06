package mail_gateways

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/adityaxdiwakar/go-sendpulse"
	"github.com/getsentry/sentry-go"
	"gitlab.com/lawchad/mailler/configs"
	"io/ioutil"
	"log"
	"net/http"
)

func SpSendEmail(name, email, htmlText string) {
	html := []byte(htmlText)
	text := []byte("Peekaboo!")
	recipients := []sendpulse.Recipient{
		sendpulse.Recipient{
			Name:  name,
			Email: email,
		},
	}
	subject := "Hey There"

	sendpulse.Initialize(
		configs.ClientID,
		configs.ClientSecret,
		"name",
		"email",
	)

	err := sendpulse.SendEmail(
		html,
		text,
		subject,
		recipients,
	)

	sentry.CaptureException(err)
	log.Fatalln(err)
}

func GetBalance() (error, int) {
	req, err := http.NewRequest(
		"GET",
		"https://api.sendpulse.com/user/balance/detail",
		nil,
	)

	if err != nil {
		sentry.CaptureException(err)
		return errors.New("something wrong with string -> request"), 0
	}

	token, err := GetKey()

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		sentry.CaptureException(err)
		return errors.New("reading the response gave an error"), 0
	}

	bytes, err := ioutil.ReadAll(resp.Body)

	var balance BalanceDetailed
	if err := json.Unmarshal(bytes, &balance); err != nil {
		fmt.Println(err)
		sentry.CaptureException(err)
		return errors.New("reading the response gave an error"), 0
	}

	return nil, balance.Email.EmailsLeft
}
