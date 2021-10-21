package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jakewmiles/vinyl-api/controllers"
	"github.com/jakewmiles/vinyl-api/models"
)

func main() {
	app := gin.Default()
	models.ConnectDatabase()

	app.GET("/albums", controllers.GetAlbums)
	app.GET("/albums/:id", controllers.GetAlbumById)
	app.POST("/albums", controllers.PostAlbum)
	app.PUT("/albums/:id", controllers.UpdateAlbumById)
	app.DELETE("/albums/:id", controllers.DeleteAlbumById)
	app.Run("localhost:8080")
}
