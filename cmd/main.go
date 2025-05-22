package main

import (
	"log"
	"example/web-service-gin/api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
    
    routes.SetupRoutes(router)
    
    log.Fatal(router.Run("localhost:8080"))
}