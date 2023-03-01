package controllers

import (
	"t-core/sources/models"

	"github.com/gin-gonic/gin"
)

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
	context.JSON(200, &models.Body{
		Message: "pong",
	})
}
