package main

import (
	_ "./docs"
	h "./handlers"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Camunda Modeller GitHub
// @version 1.0
// @description Serves changing in diagram and upload/update it in GitHub
// @termsOfService http://halykbank.kz

// @contact.name Nartay Dembayev
// @contact.url http://instagram.com/nartaymalikov
// @contact.email mastaok02@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @SecurityDefinitions.basic BasicAuth
// @host localhost:8080
// @BasePath /


func main() {
	router := gin.Default()
	config := &ginSwagger.Config{
		URL: "http://localhost:8080/swagger/doc.json", //The url pointing to API definition
	}
	// use ginSwagger middleware to
	router.GET("/swagger/*any", ginSwagger.CustomWrapHandler(config, swaggerFiles.Handler))
	router.POST("/deployment/create", h.CamundaModeller)
	router.GET("/", )
	router.Run(":8080")
}