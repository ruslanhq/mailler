package pkg

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"gitlab.com/lawchad/mailler/pkg/mail_gateways"
	"reflect"
	"strings"
	"time"
)

func CheckBalance(date *string, balance *int) int {
	if *date != time.Now().Format("02-01-2006") {
		*date = time.Now().Format("02-01-2006")
		_, *balance = mail_gateways.GetBalance()
	}
	return *balance
}

func ValidMAC(payload string, messageMAC, key []byte) bool {
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(payload))
	expectedMAC := []byte(base64.StdEncoding.EncodeToString(mac.Sum(nil)))
	return hmac.Equal(messageMAC, expectedMAC)
}

func DataStringFromStruct(query mail_gateways.Query) string {
	var dataSlice []string
	v := reflect.ValueOf(query)
	typeOfS := v.Type()

	for i := 0; i < v.NumField(); i++ {
		if typeOfS.Field(i).Name != "Mac" {
			dataSlice = append(dataSlice, fmt.Sprintf(
				"%s=%v", typeOfS.Field(i).Name, v.Field(i).Interface()),
			)
		}
	}
	return strings.Join(dataSlice, "-") + ";"
}
