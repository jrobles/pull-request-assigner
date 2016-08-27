package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/josemrobles/robification-go"
	"log"
	"net/http"
)

var (
	configs  = getConfigs()
	ghAuth   = githubAuth(configs)
	testMode = flag.Bool("test", false, "true = no notifications, false = notifications, defaults to false")
)

func init() {
	flag.Parse()

}

func main() {
	if *testMode == true {
		log.Print("INFO: In test mode, messages WILL NOT be sent and reviewers WILL NOT be assigned")
	}

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
			author := p.Pull_Request.User.Login
			prURL := p.Pull_Request.Html_Url
			prNumber := p.Number
			org := p.Pull_Request.Base.Repo.User.Login
			repo := p.Pull_Request.Head.Repo.Name
			reviewerA, reviewerB := selectReviewers(author, *configs)

			if *testMode == false {

				// Assign the PR to each reviewer
				assignToPullRequest(org, repo, prNumber, reviewerA.Github) // owner, repo, number, reviewer
				assignToPullRequest(org, repo, prNumber, reviewerB.Github) // owner, repo, number, reviewer

				// Send robification
				message := fmt.Sprint(prURL, " To: ", repo, " by: ", author, " review: ", "@", reviewerA.Flowdock, " @", reviewerB.Flowdock)
				post := robification.NewFdChat(string(configs.Fd_Token), string(message))
				err = robification.Send(post)
				if err != nil {
					res.WriteHeader(500)
					log.Printf("ERROR: Could not send robificationi: %v", err)
				} else {
					res.WriteHeader(201)
					log.Printf("INFO: Robification sent to %s and %s for %s repo", reviewerA.Flowdock, reviewerB.Flowdock, repo)
				}
			} else {
				res.WriteHeader(201)
				assignToPullRequest(org, repo, prNumber, "josemrobles")
				log.Printf("SIMULATION: Robification sent to %s and %s for %s repo", reviewerA.Flowdock, reviewerB.Flowdock, repo)
			}
		} else {
			log.Printf("No robification for %s event", string(p.Action))
		}
	}
}
