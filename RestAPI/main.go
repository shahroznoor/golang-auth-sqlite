package main

import (
	"auth.com/auth/db"
	"auth.com/auth/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	

	server.Run(":3000")
}
