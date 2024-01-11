package main

import (
	"example/db"
	"example/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080") //localhost:8080
}

// bierzemy wiadomo z gin (frameworku)
