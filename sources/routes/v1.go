package routes

import (
	"t-back-end/sources/controllers"
)

func (router *Router) V1() {
	group := router.engine.Group("/v1")

	group.GET("/ping", controllers.Pong)
}