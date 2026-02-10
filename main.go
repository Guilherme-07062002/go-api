package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string   `json:"id"`
	Title  *string  `json:"title"`
	Artist *string  `json:"artist"`
	Price  *float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: &[]string{"Blue Train"}[0], Artist: &[]string{"John Coltrane"}[0], Price: &[]float64{56.99}[0]},
	{ID: "2", Title: &[]string{"Jeru"}[0], Artist: &[]string{"Gerry Mulligan"}[0], Price: &[]float64{17.99}[0]},
	{ID: "3", Title: &[]string{"Sarah Vaughan and Clifford Brown"}[0], Artist: &[]string{"Sarah Vaughan"}[0], Price: &[]float64{39.99}[0]},
}

func main() {
	router := gin.Default()
	router.GET("/albums/:id", getAlbumBydID)
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.PUT("/albums/:id", updateAlbum)

	router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album
	newId := fmt.Sprint(len(albums) + 1)

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	newAlbum.ID = newId
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumBydID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func updateAlbum(c *gin.Context) {
	id := c.Param("id")

	for i, a := range albums {
		if a.ID == id {
			var updatedAlbum album
			if err := c.BindJSON(&updatedAlbum); err != nil {
				c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			if updatedAlbum.Title != nil {
				albums[i].Title = updatedAlbum.Title
			}
			if updatedAlbum.Artist != nil {
				albums[i].Artist = updatedAlbum.Artist
			}
			if updatedAlbum.Price != nil {
				albums[i].Price = updatedAlbum.Price
			}
			c.IndentedJSON(http.StatusOK, albums[i])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}
