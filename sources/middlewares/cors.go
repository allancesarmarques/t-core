package middlewares

import (
	CORS "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Cors struct {}

func (cors *Cors) Default() gin.HandlerFunc {
	return CORS.Default()
}