package main

import (
	"fmt"
	"net/http"
	"time"
)

func makeAsyncRestCall(url string, resp chan string) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
	}
	//fmt.Println(response.Status)
	resp <- response.Status
}

func makeSyncRestCall(url string) string {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
	}
	//fmt.Println(response.Status)
	return response.Status
}

func main() {
	asyncCalls()
	//syncCall()
}

func syncCall() {
	start := time.Now()
	url1 := "https://tutorialedge.net/golang/consuming-restful-api-with-go/"
	url2 := "https://www.sohamkamani.com/blog/2017/10/18/parsing-json-in-golang/"
	url3 := "https://stackoverflow.com/questions/29366038/looping-iterate-over-the-second-level-nested-json-in-go-lang/29382205"
	url4 := "https://blog.serverbooter.com/post/parsing-nested-json-in-go/"
	url5 := "https://www.google.com/search?client=firefox-b-d&q=go+json+traversal"
	resp1 := makeSyncRestCall(url1)
	resp2 := makeSyncRestCall(url2)
	resp3 := makeSyncRestCall(url3)
	resp4 := makeSyncRestCall(url4)
	resp5 := makeSyncRestCall(url5)
	fmt.Println(resp1)
	fmt.Println(resp2)
	fmt.Println(resp3)
	fmt.Println(resp4)
	fmt.Println(resp5)
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func asyncCalls() {
	start := time.Now()
	resp := make(chan string)
	url1 := "https://tutorialedge.net/golang/consuming-restful-api-with-go/"
	url2 := "https://www.sohamkamani.com/blog/2017/10/18/parsing-json-in-golang/"
	url3 := "https://stackoverflow.com/questions/29366038/looping-iterate-over-the-second-level-nested-json-in-go-lang/29382205"
	url4 := "https://blog.serverbooter.com/post/parsing-nested-json-in-go/"
	url5 := "https://www.google.com/search?client=firefox-b-d&q=go+json+traversal"
	go makeAsyncRestCall(url1, resp)
	go makeAsyncRestCall(url2, resp)
	go makeAsyncRestCall(url3, resp)
	go makeAsyncRestCall(url4, resp)
	go makeAsyncRestCall(url5, resp)
	fmt.Println(<-resp)
	fmt.Println(<-resp)
	fmt.Println(<-resp)
	fmt.Println(<-resp)
	fmt.Println(<-resp)
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
