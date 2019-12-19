package services

import (
	"encoding/json"
	"github.com/rifqiakrm/golang-microservice-api/config"
	"github.com/rifqiakrm/golang-microservice-api/model"
	"github.com/rifqiakrm/golang-microservice-api/model/repositories"
	"github.com/rifqiakrm/golang-microservice-api/providers/github_provider"
	"github.com/rifqiakrm/golang-microservice-api/utils/errors"
	"net/http"
	"sync"
)

type RepoService struct{}

type RepoServiceInterface interface {
	CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
	CreateRepos(input []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError)
	GetRepo() (*repositories.GetAllRepoResponse, errors.ApiError)
}

var(
	RepositoryService RepoServiceInterface
)

func init()  {
	RepositoryService = &RepoService{}
}

func (s *RepoService) CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
	if err := input.Validate(); err != nil {
		return nil, err
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

func (s *RepoService) GetRepo() (*repositories.GetAllRepoResponse, errors.ApiError)  {
	response, err := github_provider.GetRepo(config.GithubAccessToken())

	if err != nil {
		return nil, errors.NewApiError(err.StatusCode, err.Message)
	}

	var result repositories.GetAllRepoResponse
	bytes, _ := json.Marshal(response)

	if err := json.Unmarshal(bytes, &result); err != nil{
		return nil, errors.NewApiError(http.StatusInternalServerError, "invalid response")
	}

	return &result, nil

}

func (s *RepoService) CreateRepos(input []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError) {
	i := make(chan repositories.CreateRepositoriesResult)
	o := make(chan repositories.CreateReposResponse)
	defer close(o)

	var wg sync.WaitGroup

	go s.handleRepoResults(&wg, i, o)

	for _, current := range input {
		wg.Add(1)
		go s.createRepoConcurrent(current, i)
	}

	wg.Wait()
	close(i)

	result := <-o
	success := 0
	for _, current := range result.Results {
		if current.Response != nil {
			success++
		}
	}

	if success == 0 {
		result.StatusCode = result.Results[0].Error.GetStatus()
	}

	if success == len(input) {
		result.StatusCode = http.StatusCreated
	} else {
		result.StatusCode = http.StatusPartialContent
	}

	return result, nil
}

func (s *RepoService) handleRepoResults(wg *sync.WaitGroup, input chan repositories.CreateRepositoriesResult, output chan repositories.CreateReposResponse){
	var results repositories.CreateReposResponse

	for incomingEvent := range input{
		repoResult := repositories.CreateRepositoriesResult{
			Response: incomingEvent.Response,
			Error:    incomingEvent.Error,
		}
		results.Results = append(results.Results, repoResult)
		wg.Done()
	}
	output <- results
}

func (s *RepoService) createRepoConcurrent(input repositories.CreateRepoRequest, output chan repositories.CreateRepositoriesResult) {
	result, err := s.CreateRepo(input)

	if err != nil {
		output <- repositories.CreateRepositoriesResult{Error:err}
		return
	}

	output <- repositories.CreateRepositoriesResult{Response:result}
}
