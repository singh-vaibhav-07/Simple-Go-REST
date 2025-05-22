package routes

import (
	"example/web-service-gin/internal/handlers"
	"example/web-service-gin/internal/repository"
	"example/web-service-gin/internal/service"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	repo := repository.NewAlbumRepository()
	svc := service.NewAlbumService(repo)
	handler := handlers.NewAlbumHandler(svc)

	// Routes
	router.GET("/albums", handler.GetAlbums)
	router.GET("/albums/:id", handler.GetAlbumById)
	router.POST("/albums", handler.CreateAlbum)
	router.PUT("albums/:id", handler.UpdateAlbum)
	router.DELETE("/albums/:id", handler.DeleteAlbum)
}