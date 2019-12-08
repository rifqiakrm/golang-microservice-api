package github_provider

import (
	"errors"
	"github.com/rifqiakrm/golang-microservice-api/client"
	"github.com/rifqiakrm/golang-microservice-api/model"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

import "testing"

func TestMain(m *testing.M)  {
	client.StartMocks()
	os.Exit(m.Run())
}

func TestGetAuthorizationHeader(t *testing.T) {
	request := getAuthorizationHeader("abc123")
	assert.EqualValues(t, "token abc123", request)
}

func TestClientRepoErrorResponse(t *testing.T) {
	client.FlushMocks()
	client.AddMock(client.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Err:errors.New("invalid client response"),
	})
	response, err := CreateRepo("", model.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "invalid client response", err.Message)
}

func TestClientRepoInvalidResponseBody(t *testing.T) {
	client.FlushMocks()
	invalidCloser, _ := os.Open(" -asf3")
	client.AddMock(client.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			Status:           "",
			StatusCode:       http.StatusCreated,
			Body:             invalidCloser,
		},
	})
	response, err := CreateRepo("", model.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "invalid body", err.Message)
}

func TestClientRepoInvalidResponse(t *testing.T) {
	client.FlushMocks()
	client.AddMock(client.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			Status:           "",
			StatusCode:       http.StatusUnauthorized,
			Body:             ioutil.NopCloser(strings.NewReader(`{"message": 1}`)),
		},
	})
	response, err := CreateRepo("", model.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "invalid json response", err.Message)
}

func TestClientRepoInvalidInterface(t *testing.T) {
	client.FlushMocks()
	client.AddMock(client.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			Status:           "",
			StatusCode:       http.StatusUnauthorized,
			Body:             ioutil.NopCloser(strings.NewReader(`{"message": "Requires authentication","documentation_url": "https://developer.github.com/v3/repos/#create"}`)),
		},
	})
	response, err := CreateRepo("", model.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.StatusCode)
	assert.EqualValues(t, "Requires authentication", err.Message)
}

func TestClientRepoInvalidSuccessResponse(t *testing.T) {
	client.FlushMocks()
	client.AddMock(client.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			Status:           "",
			StatusCode:       http.StatusCreated,
			Body:             ioutil.NopCloser(strings.NewReader(`{"id": "123"}`)),
		},
	})
	response, err := CreateRepo("", model.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "invalid response", err.Message)
}

func TestClientRepoNoError(t *testing.T) {
	client.FlushMocks()
	client.AddMock(client.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			Status:           "",
			StatusCode:       http.StatusCreated,
			Body:             ioutil.NopCloser(strings.NewReader(`{"id": 226665745,"node_id": "MDEwOlJlcG9zaXRvcnkyMjY2NjU3NDU=","name": "Golang-Excercises","full_name": "rifqiakrm/Golang-Excercises","private": false,"owner": {}}`)),
		},
	})
	response, err := CreateRepo("", model.CreateRepoRequest{})
	assert.NotNil(t, response)
	assert.Nil(t, err)
	//assert.EqualValues(t, http.StatusCreated, response.StatusCode)
}