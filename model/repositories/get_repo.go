package repositories

type GetRepoResponse struct {
	Name string `json:"name"`
	FullName string `json:"full_name"`
	Owner RepoOwner `json:"owner"`
}

type RepoOwner struct {
	Id      int64  `json:"id" `
	Login   string `json:"login"`
	Url     string `json:"url"`
	HtmlUrl string `json:"html_url"`
}

type GetAllRepoResponse []GetRepoResponse