package main

import (
	"encoding/json"
	"fmt"
	"github.com/josemrobles/robification-go"
	"log"
	"net/http"
)

func main() {

	//	githubAuth()

	http.HandleFunc("/", indexAction)
	err := http.ListenAndServe(":8008", nil)
	if err != nil {
		log.Printf("Could not start API %q", err)
	} else {
		log.Printf("Listening on port 8008")
	}
}

func indexAction(res http.ResponseWriter, req *http.Request) {
	config := getConfig("config.json")

	decoder := json.NewDecoder(req.Body)
	var p ApiResponse
	err := decoder.Decode(&p)

	if err != nil {
		res.WriteHeader(400)
		res.Header().Set("Content-Type", "application/json")
		b, _ := json.Marshal(p)
		fmt.Fprintf(res, string(b))
	} else {

		fmt.Println(p)

		if p.Action == "opened" {
			pullRequestAuthor := p.Pull_Request.User.Login
			rev1, rev2 := selectReviewers(*pullRequestAuthor, *config)
			message := fmt.Sprint(p.Pull_Request.Html_Url, " To: ", p.Pull_Request.Head.Repo.Name, " by: ", p.Pull_Request.User.Login, " review: ", "@", rev1, " @", rev2)
			post := robification.NewFdChat(string(config.Fd_Token), string(message))
			err = robification.Send(post)
			if err != nil {
				res.WriteHeader(500)
				log.Printf("ERROR: Could not send robification")
			}
			res.WriteHeader(201)
			log.Printf("*** Robification sent to %s and %s for %s repo ***", rev1, rev2, string(*p.Pull_Request.Head.Repo.Name))
		} else {
			log.Printf("No robification for %s", string(p.Action))
		}
	}

}
