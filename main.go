package main

import (
	"fmt"
	album "go-api/domain/entities"
	factories "go-api/infra/factories"
	"net/http"

	"github.com/gin-gonic/gin"
)

var albums = []album.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()
	getAllAlbumController := factories.GetAllAlbumFactory()
	router.GET("/albums", getAllAlbumController.Handle)

	getAlbumByIdController := factories.GetAlbumByIdFactory()
	router.GET("/albums/:id", getAlbumByIdController.Handle)

	router.POST("/albums", postAlbums)
	router.PUT("/albums/:id", updateAlbum)

	router.Run("localhost:8080")
}

func postAlbums(c *gin.Context) {
	var newAlbum album.Album
	newId := fmt.Sprint(len(albums) + 1)

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	newAlbum.ID = newId
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func updateAlbum(c *gin.Context) {
	id := c.Param("id")

	for i, a := range albums {
		if a.ID == id {
			var updatedAlbum album.Album
			if err := c.BindJSON(&updatedAlbum); err != nil {
				c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			albums[i].Title = updatedAlbum.Title

			albums[i].Artist = updatedAlbum.Artist

			albums[i].Price = updatedAlbum.Price
			c.IndentedJSON(http.StatusOK, albums[i])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}
