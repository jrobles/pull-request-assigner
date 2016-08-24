package main

import (
	"fmt"
	"math/rand"
)

type UsersGitFlow struct {
	GithubName   string `json:"github_name"`
	FlowdockName string `json:"flowdock_name"`
}

type User struct {
	Login       *string    `json:"login,omitempty"`
	ID          *int       `json:"id,omitempty"`
	AvatarURL   *string    `json:"avatar_url,omitempty"`
	Name        *string    `json:"name,omitempty"`
	CreatedAt   *Timestamp `json:"created_at,omitempty"`
	UpdatedAt   *Timestamp `json:"updated_at,omitempty"`
	SuspendedAt *Timestamp `json:"suspended_at,omitempty"`
	Type        *string    `json:"type,omitempty"`
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
