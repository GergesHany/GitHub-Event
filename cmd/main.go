package main

import (
	"fmt"
	"time"
	"GitHubEvent/API"
)

func main() {
	
	username := "Gergeshany"
	events, err := API.FetchGitHubActivity(username)

	if err != nil {
		fmt.Printf("Error fetching GitHub activity: %v\n", err)
        return
	}

	for _, event := range events {
		fmt.Printf("%s - %s: %s\n", event.CreatedAt.Format(time.RFC1123Z), event.Type, event.Repo.Name)
        if event.Payload.Issue.Number != 0 {
            fmt.Printf("\tIssue: %d - %s\n", event.Payload.Issue.Number, event.Payload.Issue.Title)
        } else if event.Payload.PullRequest.Number != 0 {
            fmt.Printf("\tPull Request: %d - %s\n", event.Payload.PullRequest.Number, event.Payload.PullRequest.Title)
        }
    }
}
