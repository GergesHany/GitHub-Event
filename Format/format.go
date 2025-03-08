package format

import (
	"GitHubEvent/modules"
	"fmt"
	"time"
	"strings"
)

const (
	ColorReset  = "\033[0m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorCyan   = "\033[36m"
	ColorRed    = "\033[31m"
)

func GetRepoLink(repo string) (string) {
	idx := strings.Index(repo, "(")
	if idx != -1 {
      	repo = repo[:idx]
	}
	return fmt.Sprintf("https://github.com/%s", repo)
}

func DisplayActivity(events []modules.Event) {
	if len(events) == 0 {
		fmt.Println(ColorRed + "No recent activity found." + ColorReset)
		return
	}

	fmt.Println(ColorGreen + "Recent GitHub activity:" + ColorReset)
	fmt.Println(ColorGreen + "----------------------" + ColorReset)

	mapping := make(map[string][]modules.Details)

	for _, event := range events {
		Details := FormatEvent(event)
		if Details.RepoName != "" {
			mapping[Details.RepoName] = append(mapping[Details.RepoName], Details)
		}
	}

	for Type, Formats := range mapping {
		i := 1
		fmt.Println(ColorBlue + Type + ColorReset + " {")
		for _, Format := range Formats {
			fmt.Println(ColorYellow + "  {")
			fmt.Println(ColorCyan + "    Repo Name: " + ColorRed + Format.RepoName + ColorReset)
			fmt.Println(ColorCyan + "    Time Ago: " + ColorRed + Format.TimeAgo + ColorReset)
			fmt.Println(ColorCyan + "    Repo Link: " + ColorRed + Format.RepoLink + ColorReset)
			fmt.Println(ColorCyan + "    Event Type: " + ColorRed + Format.EventType + ColorReset)
			if i < len(mapping[Type]) {
				fmt.Println(ColorYellow + "  },")
			} else {
				fmt.Println(ColorYellow + "  }" + ColorReset)
			}
			i++
		}
		fmt.Println("}")
		fmt.Println("")
	}
}

func FormatEvent(event modules.Event) modules.Details {
	repoName := event.Repo.Name
	timeAgo := formatTimeAgo(event.CreatedAt)

	switch event.Type {
	case "CreateEvent":
		if event.Payload.RefType == "repository" {
			return modules.Details{EventType: "Created repository", RepoName: repoName, TimeAgo: timeAgo, RepoLink: GetRepoLink(repoName)}

		} else if event.Payload.RefType == "branch" {
			return modules.Details{EventType: "Created branch", RepoName: repoName, TimeAgo: timeAgo, RepoLink: GetRepoLink(repoName)}
		}

		default:
			return modules.Details{EventType: capitalize(event.Type), RepoName: repoName, TimeAgo: timeAgo, RepoLink: GetRepoLink(repoName)}
	}

	return modules.Details{}
}

func capitalize(s string) string {
    if s == "" {
        return ""
    }
    return strings.ToUpper(s[:1]) + s[1:]
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
