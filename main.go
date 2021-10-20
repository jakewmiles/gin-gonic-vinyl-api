package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jakewmiles/vinyl-api/controllers"
	"github.com/jakewmiles/vinyl-api/models"
)

func main() {
	router := gin.Default()
	models.ConnectDatabase()

	router.GET("/albums", controllers.GetAlbums)
	router.GET("/albums/:id", controllers.GetAlbumById)
	router.POST("/albums", controllers.PostAlbum)
	router.Run("localhost:8080")
}
