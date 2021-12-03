package main

import (
	"blog/models"
	"blog/routers"
	"blog/utils"
)

// @title blog
// @version 1.0
// @termsOfService http://swagger.io/terms/

// @BasePath localhost:8081
func main() {
	models.Init()
	router := routers.Init()
	router.Run(utils.HttpPort)
}
