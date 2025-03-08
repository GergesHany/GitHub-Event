package modules

import (
	"time"
)

type Repo struct {
	Name string `json:"name"`
}

type Issue struct {
	Number int    `json:"number"`
	Title  string `json:"title"`
}

type PullRequest struct {
	Number int    `json:"number"`
	Title  string `json:"title"`
}

type Payload struct {
	Action      string      `json:"action"`
	Issue       Issue       `json:"issue"`
	PullRequest PullRequest `json:"pull_request"`
	Ref         string      `json:"ref"`
	RefType     string      `json:"ref_type"`
	Size        int         `json:"size"`
}

type Event struct {
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	Repo      Repo      `json:"repo"`
	Payload   Payload   `json:"payload"`
}
