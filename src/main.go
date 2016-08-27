package main

import (
	"encoding/json"
	"fmt"
	"github.com/josemrobles/robification-go"
	"log"
	"net/http"
)

var (
	configs = getConfigs()
	ghAuth  = githubAuth(configs)
)

func main() {

	http.HandleFunc("/", indexAction)
	err := http.ListenAndServe(":8008", nil)
	if err != nil {
		log.Printf("Could not start API %q", err)
	} else {
		log.Print("Listening on port 8008")
	}
}

func indexAction(res http.ResponseWriter, req *http.Request) {

	decoder := json.NewDecoder(req.Body)
	var p ApiResponse
	err := decoder.Decode(&p)

	if err != nil {
		res.WriteHeader(400)
		res.Header().Set("Content-Type", "application/json")
		b, _ := json.Marshal(p)
		fmt.Fprintf(res, string(b))
	} else {
		if p.Action == "opened" {
			author := string(p.Pull_Request.User.Login)
			prUrl := p.Pull_Request.Html_Url
			org := string(p.Pull_Request.Base.Repo.User.Login)
			repo := string(p.Pull_Request.Head.Repo.Name)
			reviewerA, reviewerB := selectReviewers(author, *configs)

			assignToPullRequest(org, repo, 1234, "johnDoe") // owner, repo, number, reviewer
			log.Print(prUrl)

			// Send robification
			message := fmt.Sprint(prUrl, " To: ", repo, " by: ", author, " review: ", "@", reviewerA, " @", reviewerB)
			post := robification.NewFdChat(string(configs.Fd_Token), string(message))
			err = robification.Send(post)
			if err != nil {
				res.WriteHeader(500)
				log.Printf("ERROR: Could not send robification")
			}

			res.WriteHeader(201)
			log.Printf("*** Robification sent to %s and %s for %s repo ***", reviewerA, reviewerB, repo)
		} else {
			log.Printf("No robification for %s event", string(p.Action))
		}
	}
}
