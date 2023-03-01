package main

import (
	"t-core/configs"
	"t-core/sources/middlewares/auth"
	"t-core/sources/middlewares/cors"
	"t-core/sources/routes"
	"t-core/utils"
)

func main() {
	config := configs.Default()
	router := routes.Default()

	cors := cors.New(cors.Config(config.GetAllowList()))
	auth := auth.New(auth.Config(config.GetMasterKey()))

	router.Middlewares(cors, auth)
	router.V1()
	router.Swagger()

	go router.Run(config.GetPort())

	utils.GracefulShutdown()
}
