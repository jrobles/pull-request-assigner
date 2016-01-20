package main

import (
	"encoding/json"
	"fmt"
	"github.com/josemrobles/robification-go"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/", indexAction)
	err := http.ListenAndServe(":8008", nil)
	if err != nil {
		panic(err)
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
				fmt.Fprintf(res, "ERROR: Could not send robification")
			}
			res.WriteHeader(201)
			fmt.Fprintf(res, "Robification sent")
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
	counter := count_elements - 1 //available positions in the array of users
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
	seed1 := rand.NewSource(time.Now().UnixNano())
	rand1 := rand.New(seed1)
	random1 := rand1.Intn(counter)
	rev1 = fmt.Sprint(users.Users_Git_Flow[random1].FlowdockName)
	swap(users.Users_Git_Flow, random1, counter)
	counter--

	seed2 := rand.NewSource(time.Now().UnixNano() + 30)
	rand2 := rand.New(seed2)
	random2 := rand2.Intn(counter)
	rev2 = fmt.Sprint(users.Users_Git_Flow[random2].FlowdockName)

	return
}

func swap(users []UsersGitFlow, pos1, pos2 int) {
	temp := users[pos1]
	users[pos1] = users[pos2]
	users[pos2] = temp
}
