package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Geogaddi", Artist: "Boards of Canada", Price: 20.99},
	{ID: "2", Title: "Ravedeath, 1972", Artist: "Tim Hecker", Price: 28.99},
	{ID: "3", Title: "Kid A", Artist: "Radiohead", Price: 24.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.Run("localhost:8080")
}
