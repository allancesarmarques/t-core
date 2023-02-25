package main

import (
	"t-back-end/configs"
	"t-back-end/sources/middlewares"
	"t-back-end/sources/routes"
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
