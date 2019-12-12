package services

import (
	"github.com/rifqiakrm/golang-microservice-api/config"
	"github.com/rifqiakrm/golang-microservice-api/model"
	"github.com/rifqiakrm/golang-microservice-api/model/repositories"
	"github.com/rifqiakrm/golang-microservice-api/providers/github_provider"
	"github.com/rifqiakrm/golang-microservice-api/utils/errors"
	"strings"
)

type RepoService struct{}

type RepoServiceInterface interface {
	CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
}

var(
	RepositoryService RepoServiceInterface
)

func init()  {
	RepositoryService = &RepoService{}
}

func (s *RepoService) CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
	input.Name = strings.TrimSpace(input.Name)

	if input.Name == ""{
		return nil, errors.NewBadRequestError("invalid repository name")
	}

	request := model.CreateRepoRequest{
		Name:        input.Name,
		Description: input.Description,
		Private:     false,
		HasIssues:   false,
		HasProjects: false,
		HasWiki:     false,
	}

	response, err := github_provider.CreateRepo(config.GithubAccessToken(),request)

	if err != nil {
		return nil, errors.NewApiError(err.StatusCode, err.Message)
	}
	
	result := repositories.CreateRepoResponse{
		Id:    response.Id,
		Owner: response.Owner.Login,
		Name:  response.Name,
	}

	return &result, nil
}


