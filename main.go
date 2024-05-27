package main

import (
	"github.com/gin-gonic/gin"
	"go-sqlboiler/internal/handlers"
	"go-sqlboiler/internal/service"
)

func main() {
	router := gin.Default()

	albumService := service.NewAlbumService()

	router.GET("/albums", handlers.GetAlbumsHandler(albumService))
	router.GET("/albums/:id", handlers.GetAlbumByIDHandler(albumService))
	router.POST("/albums", handlers.PostAlbumHandler(albumService))
	router.DELETE("/albums/:id", handlers.DeleteAlbumHandler(albumService))
	router.PUT("/albums/:id", handlers.PutAlbumHandler(albumService))

	router.Run("localhost:8080")
}