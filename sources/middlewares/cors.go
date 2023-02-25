package middlewares

import (
	CORS "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type cors struct {}

func Cors() *cors {
	return &cors{}
}

func (*cors) Default() gin.HandlerFunc {
	return CORS.Default()
}