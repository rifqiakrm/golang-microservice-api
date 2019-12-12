package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConstant (t *testing.T)  {
	assert.EqualValues(t, "GITHUB_ACCESS_TOKEN", apiGithubAccessToken)
}

func TestGithubAccessToken(t *testing.T) {
	assert.EqualValues(t, "", GithubAccessToken())
}
