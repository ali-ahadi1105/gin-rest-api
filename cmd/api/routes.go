package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) routes() http.Handler {
	g := gin.Default()

	v1 := g.Group("/api/v1")
	{
		v1.POST("/events", app.CreateEvent)
		v1.GET("/events", app.GetAllEvents)
		v1.GET("/events/:id", app.GetEvent)
		v1.PATCH("/events", app.UpdateEvent)
		v1.DELETE("/events", app.DeleteEvent)

		v1.POST("/auth/users", app.registerUser)
	}

	return g
}
