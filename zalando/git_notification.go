package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

var httpClient = &http.Client{}

type User struct {
	UserLogin string `json:"login,omitempty"`
}

type GitNotification struct {
	Action string `json:"action,omitempty"`
	Issue  struct {
		Title      string `json:"title,omitempty"`
		User       User   `json:"user,omitempty"`
		State      string `json:"state,omitempty"`
		Created_at string `json:"created_at,omitempty"`
	} `json:"issue,omitempty"`
	PullRequest struct {
		Title     string `json:"title,omitempty"`
		User      User   `json:"user,omitempty"`
		CreatedAt string `json:"created_at,omitempty"`
		State     string `json:"state,omitempty"`
	} `json:"pull_request,omitempty"`
	Repository struct {
		Name  string `json:"name,omitempty"`
		Owner struct {
			Login string `json:"login,omitempty"`
			Type  string `json:"type,omitempty"`
		} `json:"owner,omitempty"`
	} `json:"repository,omitempty"`
}

func main() {
	http.HandleFunc("/git-notifications", handleGitNotification)
	http.ListenAndServe(":8080", nil)
}

func handleGitNotification(w http.ResponseWriter, r *http.Request) {
	//log.Printf("Received a git notification %v", r.Body)
	var gitNotification GitNotification
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer r.Body.Close()

	jsonParseError := json.Unmarshal(body, &gitNotification)
	if jsonParseError != nil {
		fmt.Println(jsonParseError)
		return
	}

	sendGoogleChatNotification(&gitNotification)

}

func sendGoogleChatNotification(gn *GitNotification) {
	if gn.Action != "opened" {
		return
	}
	var message string
	if gn.Issue.Title != "" {
		message = fmt.Sprintf("A New `ISSUE: %s` opened by user `%s` in `%s` repository at `%s`", gn.Issue.Title, gn.Issue.User.UserLogin, gn.Repository.Name, gn.Issue.Created_at)
	} else if gn.PullRequest.Title != "" {
		message = fmt.Sprintf("A New `PR: %s` opened by user `%s` in `%s` repository at `%s`", gn.PullRequest.Title, gn.PullRequest.User.UserLogin, gn.Repository.Name, gn.PullRequest.CreatedAt)
	}

	url := ""

	reqBody, err := json.Marshal(map[string]interface{}{"text": message})
	if err != nil {
		log.Fatal(err)
	}

	bodyReader := bytes.NewReader(reqBody)
	req, err := http.NewRequest(http.MethodPost, url, bodyReader)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var _, reqErr = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(reqErr)
	}

}

// create map to send git notification map by repo to channel
func createConfigMap() (map[string]string, map[string]string) {
	pushNotificationMap := make(map[string]string)
	issueNotificationMap := make(map[string]string)

	return pushNotificationMap, issueNotificationMap
}
