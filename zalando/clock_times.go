package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

type MessageQueue struct {
	Messages []struct {
		Body struct {
			PreferredGivenName string      `json:"preferredGivenName,omitempty"`
			FamilyName         string      `json:"familyName,omitempty"`
			HumanResourceID    int         `json:"humanResourceId,omitempty"`
			ActiveTo           interface{} `json:"activeTo,omitempty"`
			ActiveFrom         interface{} `json:"activeFrom,omitempty"`
			From               interface{} `json:"from,omitempty"`
			To                 interface{} `json:"to,omitempty"`
			ShiftID            int         `json:"shiftId,omitempty"`
			TeamID             int         `json:"teamId,omitempty"`
		} `json:"body,omitempty"`
		ID   int    `json:"id,omitempty"`
		Type string `json:"type,omitempty"`
	} `json:"messages,omitempty"`
	NumberOfMessagesLeft int `json:"numberOfMessagesLeft,omitempty"`
}

var client = &http.Client{}

func main() {

	resp := make(chan []byte, 20)
	write := make(chan string)
	wg := sync.WaitGroup{}
	emp, err := os.OpenFile("/Users/saukumar/code/personal_code/isable/msg_q/human_resource.json", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	shiftCreate, err := os.OpenFile("/Users/saukumar/code/personal_code/isable/msg_q/shift_create.json", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	shiftDelete, err := os.OpenFile("/Users/saukumar/code/personal_code/isable/msg_q/shift_delete.json", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}

	last_processed, err := os.ReadFile("/Users/saukumar/code/personal_code/isable/msg_q/last_processed_id.txt")
	if err != nil {
		log.Println(err)
	}
	last_processed_int, _ := strconv.Atoi(string(last_processed))

	defer emp.Close()
	defer shiftCreate.Close()
	defer shiftDelete.Close()

	wg.Add(1)
	go parallelReq(last_processed_int, last_processed_int+100000, resp, write, &wg)
	/*
		go parallelReq(9000001, 9400000, resp, write, &wg)
		go parallelReq(9000001, 9200000, resp, write, &wg)
		go parallelReq(9200001, 9400000, resp, write, &wg)
		go parallelReq(9400001, 9600000, resp, write, &wg)
		go parallelReq(9600001, 9800000, resp, write, &wg)
		go parallelReq(9800001, 1200000, resp, write, &wg)
	*/
	go func() {
		for {
			select {
			case m, ok := <-write:
				v := strings.Split(m, "~")
				switch v[0] {
				case "Shift":
					fmt.Fprintln(shiftCreate, string(v[1]), ",")
				case "ShiftDeleted":
					fmt.Fprintln(shiftDelete, string(v[1]), ",")
				case "HumanResource":
					fmt.Fprintln(emp, string(v[1]), ",")
				}
				if !ok {
					write = nil
					resp = nil
					break
				}
			}
			if write == nil && resp == nil {
				break
			}
		}
	}()

	fmt.Println("wg closeing")
	wg.Wait()
	fmt.Println("wg closed")
	close(write)
	close(resp)

}

func doReq(start string, resp chan []byte) {
	url := "https://warehouse-workforce.logistics.zalan.do/isabel/dataexposure/clocktimesbyid?lastProcessedId=" + start + "&take=100"
	method := "GET"

	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Authorization", "Bearer eyJraWQiOiJwbGF0Zm9ybS1pYW0tdmNlaHloajYiLCJhbGciOiJFUzI1NiJ9.eyJzdWIiOiI5MGVhNDI3Mi1lYzIwLTRhMGUtYTU3NS04NjRjNDA0NjYxOWYiLCJodHRwczovL2lkZW50aXR5LnphbGFuZG8uY29tL3JlYWxtIjoidXNlcnMiLCJodHRwczovL2lkZW50aXR5LnphbGFuZG8uY29tL3Rva2VuIjoiQmVhcmVyIiwiaHR0cHM6Ly9pZGVudGl0eS56YWxhbmRvLmNvbS9tYW5hZ2VkLWlkIjoic2F1a3VtYXIiLCJhenAiOiJ6dG9rZW4iLCJodHRwczovL2lkZW50aXR5LnphbGFuZG8uY29tL2JwIjoiODEwZDFkMDAtNDMxMi00M2U1LWJkMzEtZDgzNzNmZGQyNGM3IiwiYXV0aF90aW1lIjoxNjQzMDIwMDQ0LCJpc3MiOiJodHRwczovL2lkZW50aXR5LnphbGFuZG8uY29tIiwiZXhwIjoxNjQ0NDI0NDMzLCJpYXQiOjE2NDQ0MTAwMjN9.AkrRH8S7RdmqABl-QrmYqHdIw54B5wwyMh1EXrvd45kStv8FmlJZ3vmWf4Dx-OR-KDs-zZgrc1aOh6xVpcYYpQ")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	resp <- body
}

func parallelReq(start, end int, resp chan []byte, write chan string, wg *sync.WaitGroup) {
	oldStart := -1

	for start < end && oldStart != start {
		var msgQ MessageQueue
		if oldStart != start {
			oldStart = start
		}
		func(resp chan []byte, end int) {
			//fmt.Println(start)
			doReq(strconv.Itoa(start), resp)

			body := <-resp
			jsonParseError := json.Unmarshal(body, &msgQ)
			if jsonParseError != nil {
				fmt.Println(jsonParseError)
				return
			}
			for _, msg := range msgQ.Messages {
				start = msg.ID
				m, _ := json.Marshal(msg)
				write <- msg.Type + "~" + string(m)
			}

		}(resp, end)
	}
	fmt.Println("closing wait group")
	wg.Done()
}
