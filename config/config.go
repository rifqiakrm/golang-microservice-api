package config

import (
	"os"
)

const(
	apiGithubAccessToken = "GITHUB_ACCESS_TOKEN"
)

var(
	githubAccessToken = os.Getenv(apiGithubAccessToken)
)

func GithubAccessToken() string {
	return githubAccessToken
}