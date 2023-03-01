package routes

import "t-core/sources/controllers"

func (router *Router) Metrics() {
	router.engine.GET("/metrics", controllers.Prometheus())
}
