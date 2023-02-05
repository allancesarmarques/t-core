package main

import (
	"t-back-end/config"

	"github.com/gin-gonic/gin"
)

type Body struct {
  Message string `json:"message"`
}

func callback(context *gin.Context) {
	context.JSON(200, &Body{
    Message: "pong",
  })
}

func main() {
  configuration := config.New()

	router := gin.Default()
	router.GET("/ping", callback)
	router.Run(configuration.Port)
}
