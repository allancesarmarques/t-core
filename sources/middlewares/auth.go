package middlewares

import "github.com/gin-gonic/gin"

type auth struct {
	MasterKey string
}

func Auth(masterKey string) *auth {
	return &auth{masterKey}
}

func (_auth *auth) Default() gin.HandlerFunc {
	return func(context *gin.Context) {
		key := context.GetHeader("API-Key")

		if key != _auth.MasterKey {
			context.AbortWithStatusJSON(401, gin.H{
				"error": "API key inválida",
			})
			return
		}

		context.Next()
	}
}
