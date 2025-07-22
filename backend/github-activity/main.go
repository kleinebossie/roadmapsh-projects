package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type GithubEvent struct {
	Type  string `json:"type"`
	Repo struct {
		Name string `json:"name"`
	} `json:"repo"`
	Payload json.RawMessage `json:"payload"`
}

type PushPayload struct {
	Size int    `json:"size"`
	Ref  string `json:"ref"`
}

type PullRequestPayload struct {
	Action      string `json:"action"`
}

type IssuesPayload struct {
	Action string `json:"action"`
}

type ForkPayload struct {
	Forkee struct {
		Name string `json:"name"`
	} `json:"forkee"`
}

type WatchEventPayload struct {
	Action string `json:"action"`
}

type ReleaseEventPayload struct {
	Action  string `json:"action"`
}

type CreateEventPayload struct {
	Ref     string `json:"ref"`
	Reftype string `json:"ref_type"`
}

type DeleteEventPayload struct {
	Ref     string `json:"ref"`
	Reftype string `json:"ref_type"`
}

type MemberEventPayload struct {
	Action string `json:"action"`
	Member struct {
		Name string `json:"name"`
	} `json:"member"`
}

type CommitEventPayload struct {
	Comment struct {
		Body string `json:"body"`
	} `json:"comment"`
}

var events []GithubEvent

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: github-activity <username>")
		return
	}

	if os.Args[1] == "help" || os.Args[1] == "-h" {
		fmt.Println("Usage: github-activity <username>")
		return
	}

	var userName = os.Args[1]
	if userName == "" {
		fmt.Println("Usage: github-activity <username>")
	}

	resp, err := requestUserEvents(userName)
	if err != nil {
		log.Fatalf("Failed to receive http response: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		log.Fatalf("HTTP %d: %s", resp.StatusCode, string(body))
	}

	fmt.Println("Got response")
	fmt.Printf("Showing activity for %v\n\n", userName)

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read http response: %v", err)
	}

	if err = json.Unmarshal(data, &events); err != nil {
		log.Fatalf("Failed to parse json: %v", err)
	}

	for _, v := range events {
		switch v.Type {
		case "PushEvent":
			var payload PushPayload
			if err = json.Unmarshal(v.Payload, &payload); err != nil {
				fmt.Printf("Failed to read payload: %v\n", err)
			}
			fmt.Printf("pushed %v commits to %v\n", payload.Size, v.Repo.Name)
		case "CreateEvent":
			var payload CreateEventPayload
			if err = json.Unmarshal(v.Payload, &payload); err != nil {
				fmt.Printf("Failed to read payload: %v\n", err)
			}
			fmt.Printf("created %v at %v\n", payload.Reftype, v.Repo.Name)
		case "DeleteEvent":
			var payload DeleteEventPayload
			if err = json.Unmarshal(v.Payload, &payload); err != nil {
				fmt.Printf("Failed to read payload: %v\n", err)
			}
			fmt.Printf("deleted %v at %v\n", payload.Reftype, v.Repo.Name)
		case "ForkEvent":
			var payload ForkPayload
			if err = json.Unmarshal(v.Payload, &payload); err != nil {
				fmt.Printf("Failed to read payload: %v\n", err)
			}
			fmt.Printf("forked %v\n", payload.Forkee.Name)
		case "IssuesEvent":
			var payload IssuesPayload
			if err = json.Unmarshal(v.Payload, &payload); err != nil {
				fmt.Printf("Failed to read payload: %v\n", err)
			}
			fmt.Printf("%v issue at %v\n", payload.Action, v.Repo.Name)
		case "MemberEvent":
			var payload MemberEventPayload
			if err = json.Unmarshal(v.Payload, &payload); err != nil {
				fmt.Printf("Failed to read payload: %v\n", err)
			}
			fmt.Printf("%v %v to %v\n", payload.Action, payload.Member.Name, v.Repo.Name)
		case "PullRequestEvent":
			var payload PullRequestPayload
			if err = json.Unmarshal(v.Payload, &payload); err != nil {
				fmt.Printf("Failed to read payload: %v\n", err)
			}
			fmt.Printf("%v pull request at %v\n", payload.Action, v.Repo.Name)
		case "ReleaseEvent":
			var payload PullRequestPayload
			if err = json.Unmarshal(v.Payload, &payload); err != nil {
				fmt.Printf("Failed to read payload: %v\n", err)
			}
			fmt.Printf("%v new release at %v\n", payload.Action, v.Repo.Name)
		case "WatchEvent":
			var payload PullRequestPayload
			if err = json.Unmarshal(v.Payload, &payload); err != nil {
				fmt.Printf("Failed to read payload: %v\n", err)
			}
			fmt.Printf("starred %v\n", v.Repo.Name)
		}
	}
}

func requestUserEvents(username string) (*http.Response, error) {
	requestURL := fmt.Sprintf("https://api.github.com/users/%s/events", username)
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		log.Fatalf("Failed to make http request: %v", err)
	}

	req.Header.Set("accept", "application/vnd.github+json")
	req.Header.Set("username", username)

	client := &http.Client{}
	resp, err := client.Do(req)

	return resp, err
}
