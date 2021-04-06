package pkg

import (
	"errors"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"gitlab.com/lawchad/mailler/configs"
	"gitlab.com/lawchad/mailler/pkg/mail_gateways"
	"net/http"
)

var json mail_gateways.Query
var balanceInfo mail_gateways.BalanceInfo

func SendEmail(c *gin.Context) {
	var balance = CheckBalance(
		&balanceInfo.DateCheckBalance, &balanceInfo.Balance,
	)

	if err := c.ShouldBindJSON(&json); err != nil {
		sentry.CaptureException(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dataFromMessage := DataStringFromStruct(json)
	if !ValidMAC(
		dataFromMessage, []byte(json.Mac), []byte(configs.SecretSignMac),
	) {
		err := errors.New("MAC signature does not match")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		sentry.CaptureException(err)
		return
	}

	templateName := json.TemplateName
	mjmlApp := NewMjmlApp(templateName)
	htmlText, err := mjmlApp.GetHtml()
	if err != nil{
		sentry.CaptureException(err)
		err := errors.New("Can't render template with name: " + templateName)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if balance >= 100 {
		mail_gateways.SpSendEmail(json.Name, json.Mail, htmlText)
	} else {
		mail_gateways.MgSendEmail(json.Name, json.Mail, htmlText)
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
