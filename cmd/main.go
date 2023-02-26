package main

import (
	"t-core/configs"
	"t-core/sources/middlewares"
	"t-core/sources/routes"
)

func main() {
  config := configs.Default()
  router := routes.Default()

	cors := middlewares.Cors()
	auth := middlewares.Auth(config.MasterKey)

	router.Middlewares(cors, auth)
	router.V1()
	router.Run(config.Port)
}
