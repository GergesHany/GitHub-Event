package API

import (
	"GitHubEvent/modules"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func GetUrl(username string) string {
	return fmt.Sprintf("https://api.github.com/users/%s/events", username)
}

func ParseJSONResponse(resp *http.Response) ([]modules.Event, error) {
	var events []modules.Event
	err := json.NewDecoder(resp.Body).Decode(&events)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}
	return events, nil
}

func FetchGitHubActivity(username string) ([]modules.Event, error) {
	url := GetUrl(username)

	client := http.Client{
		Timeout: time.Second * 5, // Timeout after 5 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Add("User-Agent", "GitHub-Activity")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}

	defer resp.Body.Close()

	events, err := ParseJSONResponse(resp)
	return events, nil
}
