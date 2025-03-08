package format

import (
	"GitHubEvent/modules"
	"fmt"
	"time"
)

const (
	ColorReset  = "\033[0m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorCyan   = "\033[36m"
	ColorRed    = "\033[31m"
)

func DisplayActivity(events []modules.Event) {
	if len(events) == 0 {
		fmt.Println(ColorRed + "No recent activity found." + ColorReset)
		return
	}

	fmt.Println(ColorGreen + "Recent GitHub activity:" + ColorReset)
	fmt.Println(ColorGreen + "----------------------" + ColorReset)

	mapping := make(map[string][]string)

	for _, event := range events {
		Type, Format := FormatEvent(event)
		if Type != "" {
			mapping[Type] = append(mapping[Type], Format)
		}
	}

	for Type, Formats := range mapping {
		i := 1
		fmt.Println(ColorBlue + Type + ColorReset + " {")
		for _, Format := range Formats {
			fmt.Printf(ColorCyan+" %d. %s\n"+ColorReset, i, Format)
			i++
		}
		fmt.Println("}")
	}
}

func FormatEvent(event modules.Event) (string, string) {
	repoName := event.Repo.Name
	timeAgo := formatTimeAgo(event.CreatedAt)

	switch event.Type {
	case "PushEvent":
		return "Pushed to", fmt.Sprintf("%s (%s)", repoName, timeAgo)

	case "CreateEvent":
		if event.Payload.RefType == "repository" {
			return "Created repository", fmt.Sprintf("%s (%s)", repoName, timeAgo)
		} else if event.Payload.RefType == "branch" {
			return "Created branch", fmt.Sprintf("%s in %s (%s)", event.Payload.Ref, repoName, timeAgo)
		}

	case "IssuesEvent":
		return "IssuesEvent", fmt.Sprintf("- %s issue #%d in %s: %s (%s)",
			capitalize(event.Payload.Action),
			event.Payload.Issue.Number,
			repoName,
			event.Payload.Issue.Title,
			timeAgo)

	case "IssueCommentEvent":
		return "IssueCommentEvent", fmt.Sprintf("- Commented on issue in %s (%s)", repoName, timeAgo)

	case "PullRequestEvent":
		return "PullRequestEvent", fmt.Sprintf("- %s pull request #%d in %s (%s)",
			capitalize(event.Payload.Action),
			event.Payload.PullRequest.Number,
			repoName,
			timeAgo)

	case "WatchEvent":
		return "WatchEvent", fmt.Sprintf("- Starred %s (%s)", repoName, timeAgo)

	case "ForkEvent":
		return "ForkEvent", fmt.Sprintf("- Forked %s (%s)", repoName, timeAgo)
	}

	return "", ""
}

func capitalize(s string) string {
	if s == "" {
		return ""
	}
	return fmt.Sprintf("%s%s", string(s[0]-32), s[1:])
}

func formatTimeAgo(t time.Time) string {
	duration := time.Since(t)

	seconds := int(duration.Seconds())
	minutes := seconds / 60
	hours := minutes / 60
	days := hours / 24

	if days > 0 {
		return fmt.Sprintf("%d days ago", days)
	} else if hours > 0 {
		return fmt.Sprintf("%d hours ago", hours)
	} else if minutes > 0 {
		return fmt.Sprintf("%d minutes ago", minutes)
	} else {
		return fmt.Sprintf("%d seconds ago", seconds)
	}
}
