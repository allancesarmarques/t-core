package controllers

import "github.com/gin-gonic/gin"

type body struct {
  Message string `json:"message"`
}

func Pong(context *gin.Context) {
	context.JSON(200, &body{
	  Message: "pong",
  })
}