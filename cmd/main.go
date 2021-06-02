package main

import (
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"gitlab.com/lawchad/mailler/configs"
	"gitlab.com/lawchad/mailler/pkg"
	"log"
	"unit.nginx.org/go"
)

func main() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: configs.SentryDsn,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	r := gin.Default()
	r.POST("/mailler", pkg.SendEmail)
	unit.ListenAndServe(configs.Port, &*r)
}
