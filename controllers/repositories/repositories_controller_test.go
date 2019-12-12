package repositories

import (
	"encoding/json"
	"github.com/rifqiakrm/golang-microservice-api/client"
	"github.com/rifqiakrm/golang-microservice-api/model/repositories"
	"github.com/rifqiakrm/golang-microservice-api/utils/errors"
	"github.com/rifqiakrm/golang-microservice-api/utils/test"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M)  {
	client.StartMocks()
	os.Exit(m.Run())
}

func TestCreateRepoInvalidJsonBody(t *testing.T) {
	response := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(``))

	c := test.GetMockedContext(request, response)

	CreateRepo(c)

	assert.EqualValues(t, http.StatusBadRequest, response.Code)

	apiErr, err := errors.NewApiErrorFromBytes(response.Body.Bytes())

	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusBadRequest, apiErr.GetStatus())
	assert.EqualValues(t, "invalid body", apiErr.GetMessage())
}

func TestCreateRepoErrorCreateRepo(t *testing.T) {
	response := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name":"rifqi"}`))
	c := test.GetMockedContext(request, response)

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

	CreateRepo(c)

	assert.EqualValues(t, http.StatusUnauthorized, response.Code)
	apiErr, err := errors.NewApiErrorFromBytes(response.Body.Bytes())

	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusUnauthorized, apiErr.GetStatus())
	assert.EqualValues(t, "Requires authentication", apiErr.GetMessage())
}

func TestCreateRepoNoError(t *testing.T) {
	response := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name":"rifqi"}`))
	c := test.GetMockedContext(request, response)

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

	CreateRepo(c)

	assert.EqualValues(t, http.StatusCreated, response.Code)

	var result repositories.CreateRepoResponse

	err := json.Unmarshal(response.Body.Bytes(),  &result);
	assert.Nil(t, err)
	assert.EqualValues(t, 123, result.Id)
	assert.EqualValues(t, "", result.Name)
	assert.EqualValues(t, "", result.Owner)
}
