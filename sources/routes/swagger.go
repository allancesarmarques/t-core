package routes

import (
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "t-core/docs"
)

// @Title t-core API
// @Version 1.0
// @Description t-core API
func (router *Router) Swagger() {
	group := router.engine.Group("/swagger")

	group.GET("/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
