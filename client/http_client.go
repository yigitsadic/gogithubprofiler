package client

import (
	"encoding/json"
	"fmt"
	"github.com/yigitsadic/gogithubprofiler/shared"
	"log"
	"net/http"
	"strings"
)

const (
	BaseUrl = "https://api.github.com/graphql"
)

type GraphQLClient struct {
	UserName      string
	BaseUrl       string
	Authorization string
}

// Initializes a new graphql client with authentication bearer and desired userName
func NewGraphQLClient(bearer string, userName string) *GraphQLClient {
	return &GraphQLClient{
		BaseUrl:       BaseUrl,
		UserName:      userName,
		Authorization: fmt.Sprintf("Bearer %s", bearer),
	}
}

// Generates GraphQL request payload as JSON string.
func (g *GraphQLClient) generatePayload() (string, error) {
	q := PrepareQuery(g.UserName)

	var payloadMap = map[string]string{"query": q}

	a, err := json.Marshal(payloadMap)
	if err != nil {
		return "", shared.ErrUnableToMarshalToJson
	}

	return string(a), nil
}

// Generates and returns a pointer to http.Request
func (g *GraphQLClient) generateRequest(payload string) (*http.Request, error) {
	req, err := http.NewRequest("POST", g.BaseUrl, strings.NewReader(payload))
	if err != nil {
		log.Println("Unable to generate http request object", err)
		return nil, err
	}

	req.Header.Add("Authorization", g.Authorization)
	req.Header.Add("Content-Type", "application/json")

	return req, nil
}

// Queries Github API with GraphQL and handles response.
func (g *GraphQLClient) FetchUser() (*GraphQLResponse, error) {
	payload, err := g.generatePayload()
	if err != nil {
		return nil, err
	}

	c := &http.Client{}

	req, err := g.generateRequest(payload)
	if err != nil {
		return nil, err
	}

	res, err := c.Do(req)
	if err != nil {
		log.Println("Error occurred during sending HTTP Request", err)
		return nil, err
	}
	defer res.Body.Close()

	var deserialized GraphQLResponse

	err = json.NewDecoder(res.Body).Decode(&deserialized)
	if err != nil {
		log.Println("Unable to read and assign from response", err)
		return nil, err
	}

	if deserialized.Errors != nil {
		return nil, shared.ErrRequestedUserNotFound
	}

	return &deserialized, nil
}
