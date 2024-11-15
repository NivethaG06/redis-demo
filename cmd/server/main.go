package main

import (
	"1.Redis/internal/handlers"
	"1.Redis/internal/storage"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	redisClient := storage.NewRedisClient()
	defer redisClient.Close()

	router := gin.Default()

	router.GET("/data/:id", handlers.GetDataHandler(redisClient))

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("failed to run router: %v", err)
	}
}
