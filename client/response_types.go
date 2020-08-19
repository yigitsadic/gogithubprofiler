package client

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

type UserInfoResponse struct {
	Login             string     `json:"login"`
	Name              string     `json:"name"`
	AvatarUrl         string     `json:"avatarUrl"`
	Followers         Counters   `json:"followers"`
	ContributionCount Counters   `json:"repositoriesContributedTo"`
	Repositories      Repository `json:"repositories"`
}

type GraphQLResponse struct {
	Data   map[string]UserInfoResponse `json:"data"`
	Errors []PartialResponseError      `json:"errors"`
}
