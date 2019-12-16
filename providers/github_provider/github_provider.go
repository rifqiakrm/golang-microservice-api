package github_provider

import (
	"encoding/json"
	"fmt"
	"github.com/rifqiakrm/golang-microservice-api/client"
	"github.com/rifqiakrm/golang-microservice-api/model"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	headerAuthorization       = "Authorization"
	headerAuthorizationFormat = "token %s"

	urlCreateRepo = "https://api.github.com/user/repos"
	urlGetRepo = "https://api.github.com/user/repos"
)

func getAuthorizationHeader(accessToken string) string {
	return fmt.Sprintf(headerAuthorizationFormat, accessToken)
}

func CreateRepo(accessToken string, request model.CreateRepoRequest) (*model.CreateRepoResponse, *model.GithubErrorResponse) {
	headers := http.Header{}

	headers.Set(headerAuthorization, getAuthorizationHeader(accessToken))

	response, err := client.Post(urlCreateRepo, request,headers)

	if err != nil{
		log.Print("Error when trying to connect to github.com : ", err.Error())
		return nil, &model.GithubErrorResponse{
			StatusCode:       http.StatusInternalServerError,
			Message:          err.Error(),
		}
	}

	defer response.Body.Close()

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, &model.GithubErrorResponse{
			StatusCode:       http.StatusInternalServerError,
			Message:          "invalid body",
		}
	}

	if response.StatusCode > 299{
		var errResponse model.GithubErrorResponse
		if err := json.Unmarshal(bytes, &errResponse); err != nil{
			return nil, &model.GithubErrorResponse{
				StatusCode:       http.StatusInternalServerError,
				Message:          "invalid json response",
			}
		}
		errResponse.StatusCode = response.StatusCode
		return nil, &errResponse
	}

	var result model.CreateRepoResponse

	if err := json.Unmarshal(bytes, &result); err != nil {
		log.Print("Error while trying to unmarshal response : ", err.Error())
		return nil, &model.GithubErrorResponse{
			StatusCode:       http.StatusInternalServerError,
			Message:          "invalid response",
		}
	}


	return &result, nil
}

func GetRepo(accessToken string) (*model.GetAllRepoResponse, *model.GithubErrorResponse) {
	headers := http.Header{}

	headers.Set(headerAuthorization, getAuthorizationHeader(accessToken))

	response, err := client.Get(urlGetRepo, headers)

	if err != nil {
		log.Print("Error when trying to connect to github.com : ", err.Error())
		return nil, &model.GithubErrorResponse{
			StatusCode:       http.StatusInternalServerError,
			Message:          err.Error(),
		}
	}

	defer response.Body.Close()

	bytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, &model.GithubErrorResponse{
			StatusCode:       http.StatusInternalServerError,
			Message:          "invalid body",
		}
	}

	if response.StatusCode > 299{
		var errResponse model.GithubErrorResponse
		if err := json.Unmarshal(bytes, &errResponse); err != nil{
			return nil, &model.GithubErrorResponse{
				StatusCode:       http.StatusInternalServerError,
				Message:          "invalid json response",
			}
		}
		errResponse.StatusCode = response.StatusCode
		return nil, &errResponse
	}

	var result model.GetAllRepoResponse
	if err := json.Unmarshal(bytes, &result); err != nil {
		return nil, &model.GithubErrorResponse{
			StatusCode:       http.StatusInternalServerError,
			Message:          "invalid json response",
		}
	}

	return &result, nil
}