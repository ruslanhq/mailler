package pkg

import (
	"gitlab.com/lawchad/mailler/pkg/mail_gateways"
	"time"
)

func CheckBalance(date *string, balance *int) int {
	if *date != time.Now().Format("02-01-2006") {
		*date = time.Now().Format("02-01-2006")
		_, *balance = mail_gateways.GetBalance()
	}
	return *balance
}
