package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

// StartApp function to initialize app
func StartApp() {
	MapUrls()

	if err := router.Run(":8000"); err != nil {
		panic(err)
	}
}
