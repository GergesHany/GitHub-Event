package format

import (
	"GitHubEvent/modules"
	"fmt"
	"sort"
	"strings"
	"time"
)

const (
	ColorReset   = "\033[0m"
	ColorGreen   = "\033[32m"
	ColorYellow  = "\033[33m"
	ColorBlue    = "\033[34m"
	ColorCyan    = "\033[36m"
	ColorRed     = "\033[31m"
	ColorMagenta = "\033[35m"
	ColorPurple  = "\033[0;35m"
)

func GetRepoLink(repo string) string {
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

	eventTypeCounts := make(map[string]int)
	mapping := make(map[string][]modules.Details)

	for _, event := range events {
		Details := FormatEvent(event)
		if Details.RepoName != "" {
			eventTypeCounts[Details.EventType]++
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

	displayActivitySummary(eventTypeCounts, len(events))
	ActivityDays(events)
}

func ActivityDays(events []modules.Event) {

	days := make(map[string]int)
	for _, event := range events {
		date := event.CreatedAt.Format("2006-01-02")
		days[date]++
	}

	fmt.Println(ColorPurple + "Activity Days:" + ColorReset)
	fmt.Println(ColorPurple + "----------------" + ColorReset)

	maxCount := 0
	maxBarWidth := 50
	for _, count := range days {
		if count > maxCount {
			maxCount = count
		}
	}

	// Print diagrams
	for date, count := range days {
		barWidth := int(float64(maxBarWidth) * float64(count) / float64(maxCount))
		bar := strings.Repeat("█", barWidth)
		fmt.Printf(ColorCyan+"%-10s"+ColorReset+" ["+ColorGreen+"%s"+ColorReset+"] %d\n", date, bar, count)
	}
	fmt.Println("")
}

func displayActivitySummary(eventTypeCounts map[string]int, totalEvents int) {
	fmt.Println(ColorPurple + "Activity Summary:" + ColorReset)
	fmt.Println(ColorPurple + "----------------" + ColorReset)

	type eventCount struct {
		Type  string
		Count int
	}

	var counts []eventCount
	for eventType, count := range eventTypeCounts {
		counts = append(counts, eventCount{eventType, count})
	}

	sort.Slice(counts, func(i, j int) bool {
		return counts[i].Count > counts[j].Count
	})

	maxBarWidth := 50
	for _, ec := range counts {
		percentage := float64(ec.Count) / float64(totalEvents) * 100
		barWidth := int(float64(maxBarWidth) * float64(ec.Count) / float64(totalEvents))
		bar := strings.Repeat("█", barWidth)
		fmt.Printf(ColorCyan+"%-20s"+ColorReset+" ["+ColorGreen+"%s"+ColorReset+"] %d (%.1f%%)\n", ec.Type, bar, ec.Count, percentage)
	}

	fmt.Printf(ColorCyan+"Total Events: "+ColorReset+"%d\n", totalEvents)
	fmt.Println("")
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
