package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jakewmiles/vinyl-api/models"
)

func GetAlbums(c *gin.Context) {
	var albums []models.Album
	models.DB.Find(&albums)

	c.JSON(http.StatusOK, gin.H{"data": albums})
}

// func getAlbumByID(c *gin.Context) {
// 	id := c.Param("id")

// 	for _, a := range albums {
// 		if a.ID == id {
// 			c.JSON(http.StatusOK, a)
// 			return
// 		}
// 	}
// 	c.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
// }

// func postAlbums(c *gin.Context) {
// 	var newAlbum album
// 	err := c.BindJSON(&newAlbum)
// 	if err != nil {
// 		return
// 	}
// 	albums = append(albums, newAlbum)
// 	c.JSON(http.StatusCreated, newAlbum)
// }
