package routes

import (
	"github.com/gin-gonic/gin"
	"calendarme-backend/controller"
)

func NoteRoute(router *gin.Engine) {
	router.GET("/api", controller.GetNotes)
	router.POST("/api", controller.CreateNote)
	router.DELETE("/api/:id", controller.DeleteNote)
	router.PUT("/api/:id", controller.UpdateNote)
	router.GET("api/:month", controller.FindNote)
	router.GET("/searchtags/:tag", controller.SearchTags)
	router.GET("/retrievealltags", controller.RetrieveAllTags)
}
