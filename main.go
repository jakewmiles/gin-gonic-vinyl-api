package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jakewmiles/vinyl-api/controllers"
	"github.com/jakewmiles/vinyl-api/models"
)

// type album struct {
// 	ID     string  `json:"id"`
// 	Title  string  `json:"title"`
// 	Artist string  `json:"artist"`
// 	Price  float64 `json:"price"`
// }

// var albums = []album{
// 	{ID: "1", Title: "Geogaddi", Artist: "Boards of Canada", Price: 20.99},
// 	{ID: "2", Title: "Ravedeath, 1972", Artist: "Tim Hecker", Price: 28.99},
// 	{ID: "3", Title: "Kid A", Artist: "Radiohead", Price: 24.99},
// }

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

func main() {
	router := gin.Default()
	models.ConnectDatabase()

	router.GET("/albums", controllers.GetAlbums)
	// router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", controllers.PostAlbum)
	router.Run("localhost:8080")
}
