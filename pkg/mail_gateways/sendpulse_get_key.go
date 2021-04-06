package mail_gateways

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/getsentry/sentry-go"
	"gitlab.com/lawchad/mailler/configs"
	"io/ioutil"
	"net/http"
)

func GetKey() (string, error) {
	requestBody, err := json.Marshal(map[string]string{
		"grant_type":    configs.GrantType,
		"client_id":     configs.ClientID,
		"client_secret": configs.ClientSecret,
	})

	if err != nil {
		sentry.CaptureException(err)
		return "", errors.New("Marshalling request payload gave an error.")
	}

	resp, err := http.Post("https://api.sendpulse.com/oauth/access_token",
		"application/json",
		bytes.NewBuffer(requestBody))

	if err != nil {
		sentry.CaptureException(err)
		return "", errors.New("Making the request gave an error.")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		sentry.CaptureException(err)
		return "", errors.New("Reading the response gave an error.")
	}

	var response oauthTokenResponse
	err = json.Unmarshal([]byte(body), &response)

	if response.ErrorCode != 0 {
		sentry.CaptureException(err)
		return "", errors.New("SendPulse sent an error.")
	}

	//accessToken = response.AccessToken
	return response.AccessToken, nil
}
