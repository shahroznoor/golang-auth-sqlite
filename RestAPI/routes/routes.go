package routes

import (
	"auth.com/auth/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	// User Routes

	server.POST("/signup", signup)
	server.POST("/login", login)
	server.GET("/user", getUsers)
	server.GET("/user/getUser", middlewares.Authenticate, getUser)

	// REGISTRATION ROUTES
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

}
