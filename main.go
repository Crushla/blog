package main

import (
	"blog/models"
	"blog/routers"
	"blog/utils"
)

func main() {
	models.Init()
	router := routers.Init()
	router.Run(utils.HttpPort)
}
