package pkg

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/lawchad/mailler/pkg/mail_gateways"
	"net/http"
)

func GetQuery(c *gin.Context) {
	var json mail_gateways.Query
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	mail_gateways.SendEmail(json.Name, json.Mail)
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}