package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/yigitsadic/gogithubprofiler/models"
	"os"
)

func main() {
	user := flag.String("user", "", "Github user name")
	token := flag.String("token", "", "Github authentication bearer token")

	flag.Parse()

	if *user == "" {
		fmt.Println("You should provide a username")
		os.Exit(1)
	}

	if *token == "" {
		fmt.Println("You should provide an authentication bearer token")
		os.Exit(1)
	}

	usr, err := models.FetchUser(*user, *token)
	if err != nil {
		fmt.Printf("Unable to fetch given user %q \n", *user)
		os.Exit(1)
	}

	if err = json.NewEncoder(os.Stdout).Encode(&usr); err != nil {
		fmt.Println("Unable to serialize JSON")
		os.Exit(1)
	}
}
