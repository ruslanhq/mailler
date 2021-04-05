package pkg

import(
	"fmt"
	"bytes"
	"encoding/base64"
	"net/http"
	"io/ioutil"
	"log"
	JSON "encoding/json"
	"gitlab.com/lawchad/mailler"
)

type mjmlRequest struct {
	Mjml string `json:"mjml"`
}

type mjmlResponse struct {
	Html string `json:"html"`
	Mjml string `json:"mjml"`
}

type App struct{
	url string
	name string
	password string
	request []byte
}


func NewMjmlApp(templateName string) (a App) {
	mjmlTemplateString, err := GetMjmlTemplateString(templateName)
	if err != nil {
		log.Fatal("trouble with template")
	}

	mjml := mjmlRequest{Mjml:mjmlTemplateString}
	byArr, err := JSON.Marshal(mjml)

	if err != nil {
		log.Fatal(err)
	}

	app := App{
		url: "https://api.mjml.io/v1/render",
		name: mailler.UserId,
		password: mailler.UserKey,
		request: byArr
	}

	return app
}

func (a App) GetHtml() (string,error) {
	req, err := http.NewRequest("POST", a.url, bytes.NewBuffer(a.request))
    byteArray := []byte(fmt.Sprintf("%s:%s",a.name,a.password))
    b64Str := base64.StdEncoding.EncodeToString(byteArray)
    reqHeaderStr := fmt.Sprintf("Basic %s", b64Str)
    req.Header.Add("Authorization", reqHeaderStr)

    client := &http.Client{}
    resp,err := client.Do(req)
    if err != nil {
        panic(err)
    }

    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    var response mjmlResponse

    err = JSON.Unmarshal([]byte(body), &response)
    if err != nil{
        return "",err
    }

    return response.Html,nil
}

