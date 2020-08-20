package models

import (
	"github.com/yigitsadic/gogithubprofiler/client"
	"github.com/yigitsadic/gogithubprofiler/shared"
	"sort"
)

type User struct {
	UserName       string                 `json:"userName"`
	Name           string                 `json:"name"`
	ProfilePicture string                 `json:"profilePicture"`
	TotalPoint     uint32                 `json:"totalPoint"`
	Stars          uint32                 `json:"stars"`
	Followers      uint32                 `json:"followers"`
	Repos          uint32                 `json:"repos"`
	Languages      []shared.UserLanguages `json:"languages"`
}

// Calculates total developer point based on user's
// starred repositories, followers and repository count.
func (u User) calculateTotalPoints() uint32 {
	return u.Stars*100 + u.Repos*10 + u.Repos*3
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
		Stars:          res.Data["user"].Repositories.CalculateStars(),
		Repos:          res.Data["user"].CalculateTotalRepositoryCount(),
	}

	langArr := res.Data["user"].Repositories.ParseUsedLanguages()

	usr.TotalPoint = usr.calculateTotalPoints()

	sort.SliceStable(langArr, func(i, j int) bool {
		return langArr[i].Weight < langArr[j].Weight
	})

	rvr := reverse(langArr)

	if len(langArr) < 6 {
		usr.Languages = rvr
	} else {
		usr.Languages = rvr[:5]
	}

	return usr
}

func reverse(given []shared.UserLanguages) []shared.UserLanguages {
	var ret []shared.UserLanguages
	for x := len(given) - 1; x >= 0; x -= 1 {
		ret = append(ret, given[x])
	}

	return ret
}
