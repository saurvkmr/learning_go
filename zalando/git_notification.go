package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	Url = "https://chat.google.com"
)

type (
	GoogleApi struct {
		Text string `json:text`
	}

	User struct {
		UserLogin string `json:"login,omitempty"`
	}

	GitNotification struct {
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

	googleChatNotificationSender func(*GitNotification)
)

func main() {

	sender := newGoogleChatNotificationSender(*http.DefaultClient)
	handler := newGitNotificationHandler(sender)

	http.HandleFunc("/git-notifications", handler)
	http.ListenAndServe(":8080", nil)
}

func newGitNotificationHandler(sender googleChatNotificationSender) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		defer r.Body.Close()
		var gitNotification GitNotification

		if err := json.NewDecoder(r.Body).Decode(&gitNotification); err != nil {
			log.Printf("Cannot decode request payload\n")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// fire and forget (?)
		go sender(&gitNotification)

		w.WriteHeader(http.StatusOK)
	}
}

func newGoogleChatNotificationSender(httpClient http.Client) googleChatNotificationSender {

	return func(gn *GitNotification) {
		var err error

		if gn.Action != "opened" {
			return
		}

		message := prepareMessage(gn)

		buff := new(bytes.Buffer)
		if err = json.NewEncoder(buff).Encode(&GoogleApi{Text: message}); err != nil {
			log.Printf("Encountered error while encoding GoogleApi message: %v\n", err.Error())
			return
		}

		var req *http.Request

		if req, err = http.NewRequest(http.MethodPost, Url, buff); err != nil {
			log.Printf("Encountered error while creating request object: %v\n", err.Error())
			return
		}

		req.Header.Set("Content-Type", "application/json; charset=UTF-8")

		resp, err := httpClient.Do(req)

		if err != nil {
			log.Printf("Encountered error while sending request to Google: %v\n", err.Error())
			return
		}

		defer resp.Body.Close()

		if resp.StatusCode < 200 || resp.StatusCode > 299 {
			log.Printf("Received non 2XX status code from Google: %d\n", resp.StatusCode)
		}
	}
}

func prepareMessage(gn *GitNotification) string {
	var message string

	if gn.Issue.Title != "" {
		message = fmt.Sprintf("A New `ISSUE: %s` opened by user `%s` in `%s` repository at `%s`", gn.Issue.Title, gn.Issue.User.UserLogin, gn.Repository.Name, gn.Issue.Created_at)
	} else if gn.PullRequest.Title != "" {
		message = fmt.Sprintf("A New `PR: %s` opened by user `%s` in `%s` repository at `%s`", gn.PullRequest.Title, gn.PullRequest.User.UserLogin, gn.Repository.Name, gn.PullRequest.CreatedAt)
	}

	return message
}

// create map to send git notification map by repo to channel
func createConfigMap() (map[string]string, map[string]string) {
	pushNotificationMap := make(map[string]string)
	issueNotificationMap := make(map[string]string)

	return pushNotificationMap, issueNotificationMap
}
