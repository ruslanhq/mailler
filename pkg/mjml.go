package pkg

import (
	"bytes"
	"encoding/base64"
	JSON "encoding/json"
	"fmt"
	"github.com/getsentry/sentry-go"
	"gitlab.com/lawchad/mailler/configs"
	"io/ioutil"
	"log"
	"net/http"
)

type mjmlRequest struct {
	Mjml string `json:"mjml"`
}

type mjmlResponse struct {
	Html string `json:"html"`
	Mjml string `json:"mjml"`
}

type App struct {
	url      string
	name     string
	password string
	request  []byte
}

func NewMjmlApp(templateName string) (a App) {
	mjmlTemplateString, err := GetMjmlTemplateString(templateName)
	if err != nil {
		sentry.CaptureException(err)
		log.Fatal("trouble with template")
	}

	mjml := mjmlRequest{Mjml: mjmlTemplateString}
	byArr, err := JSON.Marshal(mjml)

	if err != nil {
		sentry.CaptureException(err)
		log.Fatal(err)
	}

	app := App{
		url:      "https://api.mjml.io/v1/render",
		name:     configs.MjmlApplicationId,
		password: configs.MjmlSecretKey,
		request:  byArr,
	}

	return app
}

func (a App) GetHtml() (string, error) {
	req, err := http.NewRequest("POST", a.url, bytes.NewBuffer(a.request))
	byteArray := []byte(fmt.Sprintf("%s:%s", a.name, a.password))
	b64Str := base64.StdEncoding.EncodeToString(byteArray)
	reqHeaderStr := fmt.Sprintf("Basic %s", b64Str)
	req.Header.Add("Authorization", reqHeaderStr)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		sentry.CaptureException(err)
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		sentry.CaptureException(err)
		return "", err
	}

	var response mjmlResponse

	err = JSON.Unmarshal([]byte(body), &response)
	if err != nil {
		sentry.CaptureException(err)
		return "", err
	}

	return response.Html, nil
}
