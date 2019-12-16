package app

import (
	"github.com/gin-gonic/gin"
	"github.com/rifqiakrm/golang-microservice-api/controllers/repositories"
	"github.com/rifqiakrm/golang-microservice-api/utils/errors"
	"net/http"
)

func MapUrls() {
	router.POST("/repositories", repositories.CreateRepo)
	router.GET("/repositories", repositories.GetRepo)
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, errors.NewApiError(http.StatusNotFound,"invalid route"))
	})
}
