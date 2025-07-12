package server

import (
	"github.com/gin-gonic/gin"
	"zametki/envs"
	"zametki/handlers"
)

func InitRotes() {
	router := gin.Default()
	auth := router.Group("/")
	auth.Use(handlers.AuthMiddleware())
	{
		auth.POST("/note", handlers.CreateNoteHandler)
		auth.GET("/notes", handlers.GetNotesHandler)
		auth.GET("/notes/:id", handlers.GetNoteHandler)
		auth.PUT("/notes/:id", handlers.UpdateNoteHandler)
		auth.DELETE("/notes/:id", handlers.DeleteNoteHandler)
	}
	router.Run(":" + envs.ServerEnvs.NOTES_PORT)
}
