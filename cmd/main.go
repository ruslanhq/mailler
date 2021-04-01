package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/lawchad/mailler"
	"gitlab.com/lawchad/mailler/pkg"
)

func main() {
	r := gin.Default()
	r.POST("/mailler", pkg.SendEmail)
	r.Run(mailler.Port)
}
