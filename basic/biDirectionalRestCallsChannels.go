package main

import (
	"fmt"
	"net/http"
)

func main() {
	url := make(chan string)
	go asyncCalls(url)

	url <- "https://www.google.com"
	response := <-url
	fmt.Println(response)
}

func asyncCalls(url chan string) {
	response, err := http.Get(<-url)
	if err != nil {
		url <- err.Error()
	} else {
		url <- response.Status
	}

}
