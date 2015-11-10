package main

import (
	"encoding/json"
	"fmt"
	"github.com/josemrobles/robification-go"
	"io/ioutil"
	"net/http"
	"math/rand"
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
		println(err)
	} else {
		if p.Action == "opened" {
			prOwner := p.Pull_Request.User.Login
			rev1, rev2 := selectReviewers(prOwner, *config)
			message := fmt.Sprint(p.Pull_Request.Html_Url, " To: ", p.Pull_Request.Head.Repo.Name, " by: ", p.Pull_Request.User.Login," review: ", "@",rev1," @",rev2)
			post := robification.NewFdChat(string(config.Fd_Token), string(message))
			err = robification.Send(post)
			if err != nil {
				println(err)
			}

			res.Header().Set("Content-Type", "application/json")
			b, _ := json.Marshal(p)
			fmt.Fprintf(res, string(b))
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

func selectReviewers(prOwner string, config JSONConfigData) (rev1, rev2 string) {
	counter:= 12 //positions availables in the array of users
	i := 0

	//The positions are going to change in the array to be sure each selected user is different
	//to the other and both of them are different from the PullRequest creator  
	for i < 13 {
		if owner := fmt.Sprint(config.Users_Git_Flow[i].GithubName); owner == prOwner {
			break
		}
		i+=1
	}

	if i < 13 {
		swap(config.Users_Git_Flow,i,counter)
		counter --
	}	
		random1 := rand.Intn(counter)
		rev1 = fmt.Sprint(config.Users_Git_Flow[random1].FlowdockName)
		swap(config.Users_Git_Flow,random1,counter)
		counter --

		random2 := rand.Intn(counter)
		rev2 = fmt.Sprint(config.Users_Git_Flow[random2].FlowdockName)

	return
}

func swap(arrElems []UsersGitFlow, pos1, pos2 int) {
	temp := arrElems[pos1]
	arrElems[pos1] = arrElems[pos2]
	arrElems[pos2] = temp
}
