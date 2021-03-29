package mail_gateways

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/adityaxdiwakar/go-sendpulse"
	"io/ioutil"
	"log"
	"net/http"
)

func SendEmail(name, email string) {
	html := []byte("<strong>Peekaboo!</strong>")
	text := []byte("Peekaboo!")
	recipients := []sendpulse.Recipient{
		sendpulse.Recipient{
			Name:  name,
			Email: email,
		},
	}
	subject := "Hey There"

	sendpulse.Initialize(
		"ClientID",
		"ClientSecret",
		"name",
		"email",
	)

	err := sendpulse.SendEmail(
		html,
		text,
		subject,
		recipients,
	)

	log.Fatalln(err)
}

func GetBalance() error {
	req, err := http.NewRequest(
		"GET",
		"https://api.sendpulse.com/user/balance/detail",
		nil,
	)

	if err != nil {
		return errors.New("something wrong with string -> request")
	}

	token, err := GetKey()

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.New("reading the response gave an error")
	}

	bytes, err := ioutil.ReadAll(resp.Body)

	var balance BalanceDetailed
	if err := json.Unmarshal(bytes, &balance); err != nil {
		return errors.New("reading the response gave an error")
	}

	fmt.Println(balance.Email.EmailsLeft)
	fmt.Println(balance.Balance.Main)

	return nil
}