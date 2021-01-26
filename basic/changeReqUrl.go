package main

import (
	"fmt"
	"net/http"
)

func main() {
	url := "http://google.com"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)

	res, _ := client.Do(req)
	fmt.Println(url, res.StatusCode)

	url = "https://play.golang.org/"
	q := req.URL
	fmt.Println(q)
	res, _ = client.Do(req)
	fmt.Println(req.URL, url, res.StatusCode)

}
