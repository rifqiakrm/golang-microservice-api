package model

type GithubErrorResponse struct {
	StatusCode       int           `json:"status_code"`
	Message          string        `json:"message"`
	Errors           []GithubError `json:"errors"`
	DocumentationURL string        `json:"documentation_url"`
}

type GithubError struct {
	Resource string `json:"resource"`
	Code     string `json:"code"`
	Fields   string `json:"fields`
	Message  string `json:"message"`
}
