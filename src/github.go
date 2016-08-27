package main

import (
	"fmt"
	"github.com/google/go-github/github"
	"log"
	"strings"
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
	ID        int     `json:"id,omitempty"`
	Number    int     `json:"number,omitempty"`
	State     string  `json:"state,omitempty"`
	Title     string  `json:"title,omitempty"`
	Body      string  `json:"body,omitempty"`
	User      User    `json:"user,omitempty"`
	Assignee  User    `json:"assignee,omitempty"`
	URL       string  `json:"url,omitempty"`
	HTMLURL   string  `json:"html_url,omitempty"`
	Assignees []*User `json:"assignees,omitempty"`
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
		log.Printf("INFO: Logged in as: %s", *user.Login)
		return client
	}
}

func assignToPullRequest(owner, repo string, number int, reviewer string) error {

	assignees := make([]string, 1)
	assignees[0] = reviewer

	users := &struct {
		Assignees []string `json:"assignees,omitempty"`
	}{Assignees: assignees}

	u := fmt.Sprintf("repos/%v/%v/issues/%v/assignees", owner, repo, number)
	req, err := ghAuth.NewRequest("POST", u, users)
	if err != nil {
		log.Print(err)
		return err
	}

	issue := &Issue{}
	_, err = ghAuth.Do(req, issue)
	if err != nil {
		log.Print(err)
		return err
	}
	log.Printf("INFO: PR %v has been assigned to %s on GitHub", number, reviewer)
	return nil

}
