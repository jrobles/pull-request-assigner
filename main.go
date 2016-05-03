package main

import (
	"encoding/json"
	"fmt"
	"github.com/josemrobles/robification-go"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
)

func main() {

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
	var p gitPayload
	err := decoder.Decode(&p)

	if err != nil {
		res.WriteHeader(400)
		res.Header().Set("Content-Type", "application/json")
		b, _ := json.Marshal(p)
		fmt.Fprintf(res, string(b))
	} else {
		if p.Action == "opened" {
			prOwner := p.Pull_Request.User.Login
			rev1, rev2 := selectReviewers(prOwner, *config)
			message := fmt.Sprint(p.Pull_Request.Html_Url, " To: ", p.Pull_Request.Head.Repo.Name, " by: ", p.Pull_Request.User.Login, " review: ", "@", rev1, " @", rev2)
			post := robification.NewFdChat(string(config.Fd_Token), string(message))
			err = robification.Send(post)
			if err != nil {
				res.WriteHeader(500)
				log.Printf("ERROR: Could not send robification")
			}
			res.WriteHeader(201)
			log.Printf("Robification sent to %s and %s for %s repo",rev1,rev2,string(p.Pull_Request.Head.Repo.Name))
		} else {
			log.Printf("No robification for %s",string(p.Action))
		}
	}

}

func getConfig(jsonFile string) (config *JSONConfigData) {
	config = &JSONConfigData{}
	J, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		panic(err)
	}
	json.Unmarshal([]byte(J), &config)
	return
}

func selectReviewers(prOwner string, users JSONConfigData) (rev1, rev2 string) {
	count_elements := len(users.Users_Git_Flow)
	counter := count_elements - 1 //positions availables in the array of users
	i := 0

	//The positions are going to change in the array to be sure each selected user is different
	//to the other and both of them are different from the PullRequest creator
	for i < count_elements {
		if owner := fmt.Sprint(users.Users_Git_Flow[i].GithubName); owner == prOwner {
			break
		}
		i++
	}

	if i < count_elements {
		swap(users.Users_Git_Flow, i, counter)
		counter--
	}
	random1 := rand.Intn(counter)
	rev1 = fmt.Sprint(users.Users_Git_Flow[random1].FlowdockName)
	swap(users.Users_Git_Flow, random1, counter)
	counter--

	random2 := rand.Intn(counter)
	rev2 = fmt.Sprint(users.Users_Git_Flow[random2].FlowdockName)

	return
}

func swap(users []UsersGitFlow, pos1, pos2 int) {
	temp := users[pos1]
	users[pos1] = users[pos2]
	users[pos2] = temp
}
