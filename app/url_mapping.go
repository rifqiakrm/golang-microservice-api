package app

import (
	"github.com/gin-gonic/gin"
)

func MapUrls() {
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"message":     "Invalid Route",
			"status_code": 404,
			"code":        "invalid_route",
		})
	})
}
