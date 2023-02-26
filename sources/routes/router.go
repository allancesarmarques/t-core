package routes

import "github.com/gin-gonic/gin"

type Middleware interface {
	Handler() gin.HandlerFunc
}

type Router struct {
	engine *gin.Engine
}

func Default() *Router {
	return &Router{gin.Default()}
}

func (router *Router) Middlewares(middlewares ...Middleware) {
	for _, middleware := range middlewares {
		router.engine.Use(middleware.Handler())
	}
}

func (router *Router) Run(port string) {
	router.engine.Run(":" + port)
}
