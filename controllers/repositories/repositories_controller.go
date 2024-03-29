package repositories

import (
	"github.com/gin-gonic/gin"
	"github.com/rifqiakrm/golang-microservice-api/model/repositories"
	"github.com/rifqiakrm/golang-microservice-api/services"
	"github.com/rifqiakrm/golang-microservice-api/utils/errors"
	"net/http"
)

func CreateRepo(c *gin.Context)  {
	var request repositories.CreateRepoRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		apiError := errors.NewBadRequestError("invalid body")
		c.JSON(apiError.GetStatus(), apiError)
		return
	}
	
	result, err := services.RepositoryService.CreateRepo(request)
	if err != nil{
		c.JSON(err.GetStatus(), err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func CreateRepos(c *gin.Context)  {
	var request []repositories.CreateRepoRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		apiError := errors.NewBadRequestError("invalid body")
		c.JSON(apiError.GetStatus(), apiError)
		return
	}

	result, err := services.RepositoryService.CreateRepos(request)
	if err != nil{
		c.JSON(err.GetStatus(), err)
		return
	}

	c.JSON(result.StatusCode, result)
}

func GetRepo(c *gin.Context)  {
	result, err := services.RepositoryService.GetRepo()
	if err != nil{
		c.JSON(err.GetStatus(), err)
		return
	}

	c.JSON(http.StatusOK, result)
}
