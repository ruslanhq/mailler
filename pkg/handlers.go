package pkg

import (
	"errors"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"gitlab.com/lawchad/mailler/configs"
	"gitlab.com/lawchad/mailler/pkg/mail_gateways"
	"net/http"
)

var jsonData mail_gateways.Query
var balanceInfo mail_gateways.BalanceInfo

func SendEmail(c *gin.Context) {
	var balance = CheckBalance(
		&balanceInfo.DateCheckBalance, &balanceInfo.Balance,
	)

	if err := c.ShouldBindJSON(&jsonData); err != nil {
		sentry.CaptureException(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dataFromMessage := DataStringFromStruct(jsonData)
	if !ValidMAC(
		dataFromMessage, []byte(jsonData.Mac), []byte(configs.SecretSignMac),
	) {
		err := errors.New("MAC signature does not match")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		sentry.CaptureException(err)
		return
	}
	//Create map with data for templates
	payload := jsonData.Payload
	payload["username"] = jsonData.UserName
	payload["email"] = jsonData.Mail

	mjmlApp := NewMjmlApp(jsonData.TemplateName, payload)
	htmlText, err := mjmlApp.GetHtml()
	if err != nil {
		sentry.CaptureException(err)
		err := errors.New(
			"Can't render template with name: " + jsonData.TemplateName,
		)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if balance >= 100 {
		mail_gateways.SpSendEmail(
			jsonData.UserName, jsonData.Mail, htmlText,
		)
	} else {
		mail_gateways.MgSendEmail(
			jsonData.UserName, jsonData.Mail, htmlText,
		)
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
