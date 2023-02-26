package auth

import (
	"t-core/sources/models"

	"github.com/gin-gonic/gin"
)

type config struct {
	masterKey string
}

func Config(masterKey string) *config {
	return &config{masterKey}
}

type auth struct {
	config *config
}

func New(config *config) *auth {
	return &auth{config}
}

func (_auth *auth) Handler() gin.HandlerFunc {
	return func(context *gin.Context) {
		key := context.GetHeader("API-Key")

		if key != _auth.config.masterKey {
			context.AbortWithStatusJSON(401, &models.Error{Message: "Invalid API-Key"})
			return
		}

		context.Next()
	}
}
