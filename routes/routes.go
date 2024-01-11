package routes

import (
	"example/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authenticated := server.Group("/")

	//Dzięki use, nie musimy dodawać middlewares do wszystkich poniżej
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)
	/*

		server.POST("/events", middlewares.Authenticate, createEvent)
		server.PUT("/events/:id", updateEvent)
		server.DELETE("/events/:id", deleteEvent)
	*/
	server.POST("/signup", signup)
	server.POST("/login", login)
	server.GET("/users", getUsers)
}
