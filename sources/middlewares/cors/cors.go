package cors

import (
	CORS "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type config struct {
	allowList []string
}

func Config(allowList []string) *config {
	return &config{allowList}
}

type cors struct {
	config *config
}

func New(config *config) *cors {
	return &cors{config}
}

func (_cors *cors) Handler() gin.HandlerFunc {
	config := CORS.DefaultConfig()
	config.AllowOrigins = _cors.config.allowList
	return CORS.New(config)
}
