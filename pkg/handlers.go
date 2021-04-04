package pkg

import (
	"errors"
	"github.com/gin-gonic/gin"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dataFromMessage := DataStringFromStruct(json)
	if !ValidMAC(dataFromMessage, []byte(json.Mac), []byte("123321")) {
		err := errors.New("MAC signature does not match")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if balance >= 100 {
		mail_gateways.SpSendEmail(json.Name, json.Mail)
	} else {
		mail_gateways.MgSendEmail(json.Name, json.Mail)
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
