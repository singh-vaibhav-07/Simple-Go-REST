package handlers

import (
	"example/web-service-gin/internal/models"
	"example/web-service-gin/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AlbumHandler struct {
	service service.AlbumService
}

func NewAlbumHandler(service service.AlbumService) *AlbumHandler {
	return &AlbumHandler{service: service}
}

func (h *AlbumHandler) GetAlbums(c *gin.Context) {
	albums := h.service.GetAllAlbums()
	c.IndentedJSON(http.StatusOK, albums)
}

func (h *AlbumHandler) GetAlbumById(c *gin.Context) {
	id := c.Param("id")
	album, err := h.service.GetAlbumById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, album)
}

func (h *AlbumHandler) CreateAlbum(c *gin.Context) {
	var newAlbum models.Album
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	if err := h.service.CreateAlbum(newAlbum); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error - failed to create album"})
		return
	}

	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func (h *AlbumHandler) UpdateAlbum(c *gin.Context) {
	id := c.Param("id")
	var album models.Album

	if err := c.BindJSON(&album); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	if err := h.service.UpdateAlbum(id, album); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, album)
}

func (h *AlbumHandler) DeleteAlbum(c *gin.Context) {
	id := c.Param("id")

	if err := h.service.DeleteAlbum(id); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "album deleted successfully"})
}







































