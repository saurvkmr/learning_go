package main

import (
	"fmt"
	"net/http"
)

func main() {
	response, err := http.Get("https://api.github.com/users/octocat/orgs")
	if err == nil {
		fmt.Println(response.Body)
	}
}
