package client

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestNewGraphQLClient(t *testing.T) {
	bearer := "lorem"
	user := "ipsum"
	got := NewGraphQLClient(bearer, user)

	cases := []struct {
		name     string
		expected string
		result   string
	}{
		{
			name:     fmt.Sprintf("BaseUrl should be %q", BaseUrl),
			expected: BaseUrl,
			result:   got.BaseUrl,
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			if testCase.expected != testCase.result {
				t.Errorf("expected %q equal to %q", testCase.result, testCase.expected)
			}
		})
	}
}

func TestGraphQLClient_generatePayload(t *testing.T) {
	bearer := "lorem"
	user := "ipsum"
	c := NewGraphQLClient(bearer, user)

	got, err := c.generatePayload()

	t.Run("Should not return an error", func(t *testing.T) {
		if err != nil {
			t.Errorf("not expecting error %s", err)
		}
	})

	t.Run("Should return a valid json string", func(t *testing.T) {
		var returnMap = make(map[string]string)
		err = json.Unmarshal([]byte(got), &returnMap)

		if err != nil {
			t.Errorf("not expected to get error from unmarshal json %s", err)
		}

		_, ok := returnMap["query"]

		if !ok {
			t.Errorf("Serialized content should have a \"query\" field")
		}
	})
}

func TestGraphQLClient_generateRequest(t *testing.T) {
	c := NewGraphQLClient("bearer", "user")

	got, err := c.generateRequest("lorems")

	t.Run("Should not return an error", func(t *testing.T) {
		if err != nil {
			t.Errorf("Not expecting to return an error")
		}
	})

	t.Run("Should add authorization header properly", func(t *testing.T) {
		if h := got.Header.Get("Authorization"); h != c.Authorization {
			t.Errorf("Expected to add an authorization token as %q not %q", c.Authorization, h)
		}
	})

	t.Run("Should add content type header properly", func(t *testing.T) {
		if h := got.Header.Get("Content-Type"); h != ContentType {
			t.Errorf("Expected to add a content-type header as %q not %q", ContentType, h)
		}
	})

	t.Run("Should return a http.Request", func(t *testing.T) {
		if typ, exp := fmt.Sprintf("%T", *got), "http.Request"; typ != exp {
			t.Errorf("Expected to get %q but got %q", exp, typ)
		}
	})
}
