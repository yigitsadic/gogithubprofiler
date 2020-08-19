package models

import (
	"github.com/yigitsadic/gogithubprofiler/client"
	"sort"
)

type UserLanguages struct {
	Name   string `json:"name"`
	Weight int    `json:"weight"`
}

type User struct {
	UserName       string          `json:"userName"`
	Name           string          `json:"name"`
	ProfilePicture string          `json:"profilePicture"`
	TotalPoint     uint32          `json:"totalPoint"`
	Stars          uint32          `json:"stars"`
	Followers      uint32          `json:"followers"`
	Repos          uint32          `json:"repos"`
	Languages      []UserLanguages `json:"languages"`
}

// Fetches user from Github GraphQL API.
func FetchUser(userName, auth string) (*User, error) {
	c := client.NewGraphQLClient(auth, userName)

	res, err := c.FetchUser()
	if err != nil {
		return nil, err
	}

	usr := NewUser(res)

	return usr, nil
}

// Initializes a new user with given GraphQLResponse.
func NewUser(res *client.GraphQLResponse) *User {
	usr := &User{
		Name:           res.Data["user"].Login,
		UserName:       res.Data["user"].Name,
		ProfilePicture: res.Data["user"].AvatarUrl,
		Followers:      res.Data["user"].Followers.TotalCount,
	}

	var langMap = make(map[string]int)
	var langArr []UserLanguages

	var totalRepoCount uint32
	totalRepoCount += res.Data["user"].Repositories.TotalCount
	totalRepoCount += res.Data["user"].ContributionCount.TotalCount
	usr.Repos = totalRepoCount

	var stars uint32
	for _, item := range res.Data["user"].Repositories.Nodes {
		stars += item.StarNode.TotalCount
		for _, lang := range item.Languages.Nodes {
			if val, ok := langMap[lang.Name]; ok {
				langMap[lang.Name] = val + 1
			} else {
				langMap[lang.Name] = 1
			}

			langMap[item.PrimaryLanguage.Name] = langMap[item.PrimaryLanguage.Name] + 10
		}
	}
	usr.Stars = stars

	usr.TotalPoint = stars*100 + totalRepoCount*10 + usr.Followers*3

	for k, v := range langMap {
		langArr = append(langArr, UserLanguages{
			Name:   k,
			Weight: v,
		})
	}

	sort.SliceStable(langArr, func(i, j int) bool {
		return langArr[i].Weight < langArr[j].Weight
	})

	if len(langArr) < 6 {
		usr.Languages = langArr
	} else {
		usr.Languages = langArr[5:len(langArr)]
	}

	return usr
}
