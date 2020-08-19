package client

import (
	"fmt"
)

// Prepares GraphQL query as string.
func PrepareQuery(userName string) string {
	return fmt.Sprintf(`
query {
    user(login: %q){
        login
        name
        avatarUrl

        starredRepositories {
            totalCount
        }

        followers {
            totalCount
        }

        organizations {
            totalCount
        }

        repositoriesContributedTo {
            totalCount
        }

        repositories(isFork: false, first: 100) {
            totalCount
            nodes {
                stargazers {
                    totalCount
                }

                languages(first: 7) {
                    nodes {
                        name
                    }
                }

                primaryLanguage {
                    name
                }
            }
        }
    }
}
`, userName)
}
