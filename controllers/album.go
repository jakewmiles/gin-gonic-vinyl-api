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

func GetAlbumById(c *gin.Context) {
	var album models.Album
	err := models.DB.Where("id = ?", c.Param("id")).First(&album).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Album not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": album})
}

func PostAlbum(c *gin.Context) {
	var input models.CreateAlbumInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	album := models.Album{Title: input.Title, Artist: input.Artist, Price: input.Price}
	models.DB.Create(&album)
	c.JSON(http.StatusOK, gin.H{"data": album})
}
