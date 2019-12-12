package services

import (
	"github.com/rifqiakrm/golang-microservice-api/client"
	"github.com/rifqiakrm/golang-microservice-api/model/repositories"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M)  {
	client.StartMocks()
	os.Exit(m.Run())
}

func TestCreateRepoInvalidName(t *testing.T)  {
	result, err := RepositoryService.CreateRepo(repositories.CreateRepoRequest{
		Name:        "",
	})

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.GetStatus())
	assert.EqualValues(t, "invalid repository name", err.GetMessage())
}

func TestCreateRepoErrorFromGithub(t *testing.T)  {
	client.FlushMocks()
	client.AddMock(client.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			Status:           "",
			StatusCode: http.StatusUnauthorized,
			Body:             ioutil.NopCloser(strings.NewReader(`{"message": "Requires authentication","documentation_url": "https://developer.github.com/v3/repos/#create"}`)),
		},
	})

	response, err := RepositoryService.CreateRepo(repositories.CreateRepoRequest{
		Name:        "name",
	})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.GetStatus())
	assert.EqualValues(t, "Requires authentication", err.GetMessage())
}

func TestCreateRepoNoErrorFromGithub(t *testing.T)  {
	client.FlushMocks()
	client.AddMock(client.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			Status:           "",
			StatusCode: http.StatusCreated,
			Body:             ioutil.NopCloser(strings.NewReader(`{"id": 123}`)),
		},
	})

	response, err := RepositoryService.CreateRepo(repositories.CreateRepoRequest{
		Name:        "name",
	})

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.EqualValues(t, 123, response.Id)
}