package client

import (
	"github.com/yigitsadic/gogithubprofiler/shared"
)

type PartialResponseError struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

type Counters struct {
	TotalCount uint32 `json:"totalCount"`
}

type LanguagePartial struct {
	Name string `json:"name"`
}

type LanguageNode struct {
	Nodes []*LanguagePartial `json:"nodes"`
}

type Stargazers struct {
	TotalCount uint32 `json:"totalCount"`
}

type RepositoryNode struct {
	StarNode        Stargazers       `json:"stargazers"`
	PrimaryLanguage *LanguagePartial `json:"primaryLanguage"`
	Languages       LanguageNode     `json:"languages"`
}

type Repository struct {
	TotalCount uint32           `json:"totalCount"`
	Nodes      []RepositoryNode `json:"nodes"`
}

// Calculates total stars given to user's repositories.
func (r Repository) CalculateStars() (total uint32) {
	for _, node := range r.Nodes {
		total += node.StarNode.TotalCount
	}

	return
}

// Parses languages used by user in his/her repositories.
func (r Repository) ParseUsedLanguages() []shared.UserLanguages {
	var languages []shared.UserLanguages
	var langMap = make(map[string]int)

	for _, item := range r.Nodes {
		langMap[item.PrimaryLanguage.Name] = langMap[item.PrimaryLanguage.Name] + 10

		for _, x := range item.Languages.Nodes {
			langMap[x.Name] = langMap[x.Name] + 1
		}
	}

	for k, v := range langMap {
		languages = append(languages, shared.UserLanguages{Name: k, Weight: v})
	}

	return languages
}

type UserInfoResponse struct {
	Login             string     `json:"login"`
	Name              string     `json:"name"`
	AvatarUrl         string     `json:"avatarUrl"`
	Followers         Counters   `json:"followers"`
	ContributionCount Counters   `json:"repositoriesContributedTo"`
	Repositories      Repository `json:"repositories"`
}

// Calculates total interacted repository count for user info response.
func (u UserInfoResponse) CalculateTotalRepositoryCount() uint32 {
	return u.Repositories.TotalCount + u.ContributionCount.TotalCount
}

type GraphQLResponse struct {
	Data   map[string]UserInfoResponse `json:"data"`
	Errors []PartialResponseError      `json:"errors"`
}
