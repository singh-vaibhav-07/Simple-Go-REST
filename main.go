package main

import (
	"net/http"

    "github.com/gin-gonic/gin"
)

// data to store in album record
type album struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float32 `json:"price"`
}

// albums slice to store record album data.
var albums = []album {
	{Id: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{Id: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{Id: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{Id: "4", Title: "Angels for each other", Artist: "Arijit Singh", Price: 69.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.POST("albums", postAlbums)
	router.PUT("/albums/:id", updateAlbum)
	router.DELETE("/albums/:id", deleteAlbum)

	router.Run("localhost:8080")
}

// to get list of albums
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// to add an album from json received in request body
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the receiced json to newalbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// get album by Id -> parameter sent by client, then returns that album as a response.
func getAlbumById(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of albums and look for album whose Id matches the parameter.
	for _, a := range albums {
		if a.Id == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message":"album not found!"})
}

func updateAlbum(c *gin.Context) {
	id := c.Param("id")
	var updatedAlbum album


	// Bind the received JSON to updatedAlbum
	if err := c.BindJSON(&updatedAlbum); err != nil {
		return
	}

	// find and update the album
	for i, a := range albums {
		if a.Id == id {
			// Ensure the ID remains the same
			updatedAlbum.Id = id
			albums[i] = updatedAlbum
			c.IndentedJSON(http.StatusOK, updatedAlbum)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found!"})
}

func deleteAlbum(c *gin.Context) {
	id := c.Param("id")

	// Find and delete the album
	for i, a := range albums {
		if a.Id == id {
			// Remove the album from slice
			albums = append(albums[:i], albums[i+1:]..., )
			c.IndentedJSON(http.StatusOK, gin.H{"message": "album deleted."})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found!"})
}