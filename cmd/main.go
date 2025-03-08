package main

import (
	"GitHubEvent/API"
	format "GitHubEvent/Format"
	"fmt"
)

func main() {

	var username string
	fmt.Print(format.ColorBlue + "Enter GitHub username: " + format.ColorReset)

	fmt.Scanln(&username)
	fmt.Println(format.ColorBlue + "Fetching GitHub activity for " + username + "..." + format.ColorReset)

	events, err := API.FetchGitHubActivity(username)

	if err != nil {
		fmt.Println(format.ColorRed + "Error fetching GitHub activity: " + err.Error() + format.ColorReset)
		return
	}

	format.DisplayActivity(events)
}
