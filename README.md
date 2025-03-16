# GitHub Activity CLI

A command-line tool that fetches and displays a GitHub user's recent activity with color-coded organization and visualization.

<!-- ![GitHub Activity CLI Demo] -->

<hr>

## Features

- **Activity Tracking**: Fetch recent GitHub events for any user
- **Color-Coded Output**: View GitHub activities with intuitive color highlighting
- **Repository Grouping**: Activities organized by repository
- **Event Categorization**: Different GitHub events (pushes, stars, issues, PRs) clearly labeled
- **Activity Summary**: Visual bar chart showing distribution of different activity types
- **Time-Relative Formatting**: Events displayed with human-readable timestamps (e.g., "2 days ago")
- **Repository Links**: Quick access to relevant GitHub repositories

## Installation

```bash
# Clone the repository
git clone git@github.com:GergesHany/GitHub-Event.git
```


## Project Structure

```
GitHubEvent/
├── cmd/
│   └── main.go         # Entry point
├── api/
│   └── api.go          # GitHub API interaction
├── modules/
│   └── modules.go      # Data structures
├── format/
│   └── format.go       # Output formatting
├── go.mod              # Go module definition
└── README.md           # Documentation
```

## Requirements

- Go 1.16 or higher
- Internet connection to access GitHub API

## Limitations

- GitHub API has rate limiting (60 requests per hour for unauthenticated requests)
- Fetches only the most recent activities (up to 30 events)
