package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID				string		`json:"id`
	Title			string		`json:"title"`
	Artist		string		`json:"artist`
	Price			float64		`json:"price"`
}

var albums = []album {
	{ID: "1", Title: "Susan", Artist: "The Soil", Price: 100.05},
	{ID: "2", Title: "Sechaba", Artist: "Lundi", Price: 200.05},
	{ID: "3", Title: "Isandla", Artist: "Rebecca", Price: 500.05},
	{ID: "4", Title: "Bokintsa", Artist: "Solly", Price: 400.05},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}


