package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Router struct {
	engine *gin.Engine
}

func Default() *Router {
	return &Router{ gin.Default() }
}

func (router *Router) Middlewares() {
	router.engine.Use(cors.Default())
}

func (router *Router) Run(port string) {
	router.engine.Run(":" + port)
}