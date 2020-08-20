package client

import (
	"testing"
)

func TestRepository_CalculateStars(t *testing.T) {
	nodes := []RepositoryNode{
		{
			StarNode:        Stargazers{TotalCount: 35},
			PrimaryLanguage: nil,
			Languages:       LanguageNode{},
		},
		{
			StarNode:        Stargazers{TotalCount: 35},
			PrimaryLanguage: nil,
			Languages:       LanguageNode{},
		},
	}

	repo := &Repository{
		Nodes: nodes,
	}

	t.Run("Should calculate stars properly", func(t *testing.T) {
		got := repo.CalculateStars()

		if expected := uint32(70); got != expected {
			t.Errorf("Expected stars is %d but got %d", expected, got)
		}
	})
}

func TestUserInfoResponse_CalculateTotalRepositoryCount(t *testing.T) {
	resp := &UserInfoResponse{
		ContributionCount: Counters{TotalCount: uint32(350)},
		Repositories: Repository{
			TotalCount: uint32(350),
			Nodes:      nil,
		},
	}

	t.Run("Should calculate total repositories correctly", func(t *testing.T) {
		if expected, got := uint32(700), resp.CalculateTotalRepositoryCount(); expected != got {
			t.Errorf("expected repository count %d but got %d", expected, got)
		}
	})
}

func TestRepository_ParseUsedLanguages(t *testing.T) {
	nodes := []RepositoryNode{
		{
			PrimaryLanguage: &LanguagePartial{Name: "Ruby"},
			Languages: LanguageNode{
				[]*LanguagePartial{
					{Name: "Ruby"},
					{Name: "Go"},
					{Name: "HTML"},
					{Name: "JavaScript"},
					{Name: "CSS"},
				},
			},
		},
		{
			PrimaryLanguage: &LanguagePartial{Name: "Go"},
			Languages: LanguageNode{
				[]*LanguagePartial{
					{Name: "Go"},
					{Name: "HTML"},
					{Name: "JavaScript"},
				},
			},
		},
	}

	repo := &Repository{
		Nodes: nodes,
	}

	got := repo.ParseUsedLanguages()
	testCases := map[string]int{
		"Go":         12,
		"Ruby":       11,
		"HTML":       2,
		"JavaScript": 2,
		"CSS":        1,
	}

	t.Run("Should have equal length", func(t *testing.T) {
		if a, b := len(got), len(testCases); a != b {
			t.Errorf("Expected length is %d but got %d", b, a)
		}
	})

	t.Run("Values should match", func(t *testing.T) {
		for _, item := range got {
			val, ok := testCases[item.Name]
			if !ok {
				t.Errorf("Expected to key present but not exists key=%s", item.Name)
			}

			if val != item.Weight {
				t.Errorf("Expected to language have weight %d but got %d", testCases[item.Name], item.Weight)
			}
		}
	})
}
