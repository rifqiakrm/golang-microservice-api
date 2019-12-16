package model

type GetRepoResponse struct {
	Name string `json:"name"`
	FullName string `json:"full_name"`
	Owner RepoOwner `json:"owner"`
}

type GetAllRepoResponse []GetRepoResponse