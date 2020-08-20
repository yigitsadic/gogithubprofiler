package client

import (
	"fmt"
	"strings"
	"testing"
)

func TestPrepareQuery(t *testing.T) {
	got := PrepareQuery("yigitsadic")

	expectedKeywords := []string{
		"yigitsadic", "name", "login", "avatarUrl", "starredRepositories",
		"followers", "organizations", "repositoriesContributedTo",
		"repositories", "languages", "primaryLanguage", "stargazers",
	}

	for _, key := range expectedKeywords {
		t.Run(fmt.Sprintf("should contain %q key in prepared query", key), func(t *testing.T) {
			if !strings.Contains(got, key) {
				t.Errorf("Expected to contain %q but not found", key)
			}
		})
	}
}
