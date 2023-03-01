package routes

import (
	"t-core/sources/controllers"
)

func (router *Router) V1() {
	group := router.Engine.Group("/v1")

	group.GET("/ping", controllers.Pong)
}
