package model

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateRepo(t *testing.T) {
	request := CreateRepoRequest{
		Name:        "Golang Repo Test Github API",
		Description: "Testing github API",
		Homepage:    "github.com",
		Private:     false,
		HasIssues:   true,
		HasProjects: false,
		HasWiki:     false,
	}

	if request.Private {

	}

	//Convert struct to json
	bytes, err := json.Marshal(request)

	assert.Nil(t, err)
	assert.NotNil(t, bytes)

	var target CreateRepoRequest

	err = json.Unmarshal(bytes, &target)
	assert.Nil(t, err)
}
