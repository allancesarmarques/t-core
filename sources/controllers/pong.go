package controllers

import "github.com/gin-gonic/gin"

type body struct {
	Message string `json:"message"`
}

// Ping Example
// @Summary Ping Example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} pong
// @Failure 400 {string} models.Error
// @Router /v1/ping [get]
func Pong(context *gin.Context) {
	context.JSON(200, &body{
		Message: "pong",
	})
}
