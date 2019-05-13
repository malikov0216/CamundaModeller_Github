package main

import (
	h "./handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/deployment/create", h.CamundaModeller)
	router.GET("/", )
	router.Run(":8080")
}