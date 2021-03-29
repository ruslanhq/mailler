package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/lawchad/mailler/pkg"
)

func main() {
	r := gin.Default()
	r.POST("/mailler", pkg.GetQuery)
	r.Run(":8080")
}