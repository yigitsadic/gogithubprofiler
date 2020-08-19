package main

import (
	"encoding/json"
	"flag"
	"github.com/yigitsadic/gogithubprofiler/models"
	"log"
	"os"
)

func main() {
	user := flag.String("user", "", "Github user name")
	token := flag.String("token", "", "Github authorization token")

	flag.Parse()

	usr, err := models.FetchUser(*user, *token)
	if err != nil {
		log.Fatalf("Unable to fetch given user %q", *user)
	}

	json.NewEncoder(os.Stdout).Encode(&usr)
}
