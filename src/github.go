package main

import (
	"github.com/google/go-github/github"
	"log"
	"strings"
	"time"
)

type ApiResponse struct {
	Action       string      `json:"action,omitempty"`
	Number       int         `json:"number,omitempty"`
	Pull_Request PullRequest `json:"pull_request,omitempty"`
}

type PullRequest struct {
	Html_Url string `json:"html_url,omitempty"`
	Head     struct {
		Repo *Repository `json:"repo,omitempty"`
	} `json:"head,omitempty"`
	Base struct {
		User *User       `json:"user,omitempty"`
		Repo *Repository `json:"repo,omitempty"`
	} `json:"base,omitempty"`
	User *User `json:"user,omitempty"`
}

// Issue represents a GitHub issue on a repository.
type Issue struct {
	ID        int       `json:"id,omitempty"`
	Number    int       `json:"number,omitempty"`
	State     string    `json:"state,omitempty"`
	Title     string    `json:"title,omitempty"`
	Body      string    `json:"body,omitempty"`
	User      User      `json:"user,omitempty"`
	Assignee  User      `json:"assignee,omitempty"`
	ClosedAt  time.Time `json:"closed_at,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	URL       string    `json:"url,omitempty"`
	HTMLURL   string    `json:"html_url,omitempty"`
	Assignees []*User   `json:"assignees,omitempty"`
}

type Repository struct {
	ID       int    `json:"id,omitempty"`
	User     User   `json:"owner,omitempty"`
	Name     string `json:"name,omitempty"`
	FullName string `json:"full_name,omitempty"`
	HTMLURL  string `json:"html_url,omitempty"`
}

func githubAuth(configs *Config) *github.Client {
	tp := github.BasicAuthTransport{
		Username: strings.TrimSpace(configs.Github_Login),
		Password: strings.TrimSpace(configs.Github_Password),
	}

	client := github.NewClient(tp.Client())
	user, _, err := client.Users.Get("")
	if err != nil {
		log.Printf("error: %v", err)
		return nil
	} else {
		log.Printf("Logged in as: %s", string(*user.Login))
		return client
	}
}

func assignToPullRequest(owner, repo string, number int, reviewer string) error {
	log.Print(owner, repo, number, reviewer)

	return nil
}
