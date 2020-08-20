package models

import (
	"fmt"
	"github.com/yigitsadic/gogithubprofiler/shared"
	"testing"
)

func Test_reverse(t *testing.T) {
	languages := []shared.UserLanguages{
		{
			Name:   "Go",
			Weight: 15,
		},
		{
			Name:   "Ruby",
			Weight: 10,
		},
		{
			Name:   "JavaScript",
			Weight: 5,
		},
	}

	expected := []string{"JavaScript", "Ruby", "Go"}
	got := reverse(languages)

	for i, language := range expected {
		t.Run(fmt.Sprintf("%d. element should be %s", i, language), func(t *testing.T) {
			if got[i].Name != language {
				t.Errorf("expected %q equal to %q", got[i].Name, language)
			}
		})
	}

}

func TestUser_calculateTotalPoints(t *testing.T) {
	usr := &User{
		Stars:     1,
		Followers: 1,
		Repos:     1,
	}

	expected := usr.Stars*100 + usr.Repos*10 + usr.Followers*3

	t.Run("Should calculate as expected", func(t *testing.T) {
		got := usr.calculateTotalPoints()

		if expected != got {
			t.Errorf("Expected to calculation result %d but got %d", expected, got)
		}
	})
}
