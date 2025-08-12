package main

import (
	"ecommerce/models"
	"ecommerce/server"
)

func main() {
	var repo = models.NewMemoryRepository()
	var httpHandlers = server.NewHttpHandlers(repo)
	server.RunServer(8010, httpHandlers)
}
