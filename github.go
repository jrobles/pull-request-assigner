package main

import (
	"bufio"
	"fmt"
	"github.com/google/go-github/github"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strings"
	"syscall"
	"time"
)

type Timestamp struct {
	time.Time
}

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
	User *User `json:"user,omitempty"`
}

// Issue represents a GitHub issue on a repository.
type Issue struct {
	ID        *int       `json:"id,omitempty"`
	Number    *int       `json:"number,omitempty"`
	State     *string    `json:"state,omitempty"`
	Title     *string    `json:"title,omitempty"`
	Body      *string    `json:"body,omitempty"`
	User      *User      `json:"user,omitempty"`
	Assignee  *User      `json:"assignee,omitempty"`
	ClosedAt  *time.Time `json:"closed_at,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	URL       *string    `json:"url,omitempty"`
	HTMLURL   *string    `json:"html_url,omitempty"`
	Assignees []*User    `json:"assignees,omitempty"`
}

type Repository struct {
	ID        *int        `json:"id,omitempty"`
	Owner     *User       `json:"owner,omitempty"`
	Name      *string     `json:"name,omitempty"`
	FullName  *string     `json:"full_name,omitempty"`
	CreatedAt *Timestamp  `json:"created_at,omitempty"`
	PushedAt  *Timestamp  `json:"pushed_at,omitempty"`
	UpdatedAt *Timestamp  `json:"updated_at,omitempty"`
	HTMLURL   *string     `json:"html_url,omitempty"`
	Source    *Repository `json:"source,omitempty"`
	Private   *bool       `json:"private"`
}

func githubAuth() {
	r := bufio.NewReader(os.Stdin)
	fmt.Print("GitHub Username: ")
	username, _ := r.ReadString('\n')

	fmt.Print("GitHub Password: ")
	bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	password := string(bytePassword)

	tp := github.BasicAuthTransport{
		Username: strings.TrimSpace(username),
		Password: strings.TrimSpace(password),
	}

	client := github.NewClient(tp.Client())
	user, _, err := client.Users.Get("")
	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return
	}
	fmt.Printf("\n\nLogged in as: %v\n\n", string(*user.Login))
}

//func (s *IssuesService) assignToPr(owner, repo string, number int, assignees []string) error {
func assignToPr(owner, repo string, number int, assignees []string) error {
	/*
			users := &struct {
				Assignees []string `json:"assignees,omitempty"`
			}{Assignees: assignees}
			u := fmt.Sprintf("repos/%v/%v/issues/%v/assignees", owner, repo, number)
			req, err := s.client.NewRequest("POST", u, users)
			if err != nil {
				return err
			}

			issue := &Issue{}
			resp, err := s.client.Do(req, issue)

			fmt.Println(issue, resp)
		return err
	*/
	return nil
}
