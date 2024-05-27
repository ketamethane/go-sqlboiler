package handlers

import (
	"net/http"
	"go-sqlboiler/internal/models"
	"go-sqlboiler/internal/service"

	"github.com/gin-gonic/gin"
)


func GetAlbumsHandler(albumService *service.AlbumService) func(c *gin.Context) {
	return func(c *gin.Context) {
		albums := albumService.GetAllAlbums()
		c.JSON(http.StatusOK, albums)
	}
}

func GetAlbumByIDHandler(albumService *service.AlbumService) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		album, err := albumService.GetAlbumByID(id)
		if err!= nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "album not found"})
			return
		}
		c.JSON(http.StatusOK, album)
	}
}

func PostAlbumHandler(albumService *service.AlbumService) func(c *gin.Context) {
	return func(c *gin.Context) {
		var newAlbum models.Album
		if err := c.ShouldBindJSON(&newAlbum); err!= nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		albumService.AddNewAlbum(newAlbum)
		c.JSON(http.StatusCreated, newAlbum)
	}
}

func PutAlbumHandler(albumService *service.AlbumService) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		var updatedAlbum models.Album
		if err := c.ShouldBindJSON(&updatedAlbum); err!= nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := albumService.UpdateAlbumByID(id, updatedAlbum)
		if err!= nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "album not found"})
			return
		}
		c.JSON(http.StatusOK, updatedAlbum)
	}
}

func DeleteAlbumHandler(albumService *service.AlbumService) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		err := albumService.DeleteAlbumByID(id)
		if err!= nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "album not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "album deleted"})
	}
}
