package main

import (
	factories "go-api/infra/factories"
	"go-api/infra/mocks"
	inmemorydb "go-api/infra/repositories"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	albums := mocks.GetAlbumsInMemory()

	repo := inmemorydb.NewAlbumRepository(albums)

	getAllAlbumController := factories.GetAllAlbumFactory(repo)
	router.GET("/albums", getAllAlbumController.Handle)

	getAlbumByIdController := factories.GetAlbumByIdFactory(repo)
	router.GET("/albums/:id", getAlbumByIdController.Handle)

	createAlbumController := factories.CreateAlbumFactory(repo)
	router.POST("/albums", createAlbumController.Handle)

	// router.PUT("/albums/:id", updateAlbum)

	router.Run("localhost:8080")
}

// func updateAlbum(c *gin.Context) {
// 	id := c.Param("id")

// 	for i, a := range albums {
// 		if a.ID == id {
// 			var updatedAlbum entities.Album
// 			if err := c.BindJSON(&updatedAlbum); err != nil {
// 				c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 				return
// 			}

// 			albums[i].Title = updatedAlbum.Title

// 			albums[i].Artist = updatedAlbum.Artist

// 			albums[i].Price = updatedAlbum.Price
// 			c.IndentedJSON(http.StatusOK, albums[i])
// 			return
// 		}
// 	}
// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
// }
