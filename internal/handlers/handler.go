package handlers

import (
	"1.Redis/internal/models"
	"1.Redis/internal/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func GetDataHandler(redisClient *storage.RedisClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var data models.Data
		err := redisClient.Get(c.Request.Context(), id, &data)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get data from cache"})
			return
		}
		if data.ID != "" {
			c.JSON(http.StatusOK, gin.H{"data": data})
			return
		}

		dataDb, err := storage.FetchDataFromDb(id)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get data from db"})
			return
		}

		err = redisClient.Set(c.Request.Context(), id, dataDb, 5*time.Minute)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set data to cache"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": dataDb})
	}
}
