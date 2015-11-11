package main

import "testing"

func TestSwap(t *testing.T) {
	u0 := UsersGitFlow{GithubName: "Eric0", FlowdockName: "eric test0"}
	u1 := UsersGitFlow{GithubName: "Eric1", FlowdockName: "eric test1"}
	u2 := UsersGitFlow{GithubName: "Eric2", FlowdockName: "eric test2"}

	temp := []UsersGitFlow{
		u0,
		u1,
		u2,
	}

	swap(temp, 1, 2)

	if temp[1] == u1 {
		t.Fatal("errr")
	}
	if temp[2] == u2 {
		t.Fatal("errr")
	}
}

func TestSelectReviewers(t *testing.T) {
	u0 := UsersGitFlow{GithubName: "User0", FlowdockName: "user test0"}
	u1 := UsersGitFlow{GithubName: "User1", FlowdockName: "user test1"}
	u2 := UsersGitFlow{GithubName: "User2", FlowdockName: "user test2"}
	u3 := UsersGitFlow{GithubName: "User3", FlowdockName: "user test3"}
	u4 := UsersGitFlow{GithubName: "User4", FlowdockName: "user test4"}
	u5 := UsersGitFlow{GithubName: "User5", FlowdockName: "user test5"}
	u6 := UsersGitFlow{GithubName: "User6", FlowdockName: "user test6"}
	u7 := UsersGitFlow{GithubName: "User7", FlowdockName: "user test7"}
	u8 := UsersGitFlow{GithubName: "User8", FlowdockName: "user test8"}
	u9 := UsersGitFlow{GithubName: "User9", FlowdockName: "user test9"}
	u10 := UsersGitFlow{GithubName: "User10", FlowdockName: "user test10"}
	u11 := UsersGitFlow{GithubName: "User11", FlowdockName: "user test11"}
	u12 := UsersGitFlow{GithubName: "User12", FlowdockName: "user test12"}

	temp := []UsersGitFlow{
		u0,
		u1,
		u2,
		u3,
		u4,
		u5,
		u6,
		u7,
		u8,
		u9,
		u10,
		u11,
		u12,
	}
	users := UsersData{
		Users_Git_Flow: temp,
	}
	prOwner := "User0"
	rev1, rev2 := selectReviewers(prOwner, users)
	if rev1 == "" {
		t.Fatal("rev1 can not be empty")
	}
	if rev2 == "" {
		t.Fatal("rev2 can not be empty")
	}
	if rev1 == rev2 {
		t.Fatal("rev2 and rev1 can not be the same")
	}
	if rev1 == prOwner || rev2 == prOwner {
		t.Fatal("revs can not be equal to the prOwner")
	}
}

func TestGetUsers(t *testing.T) {
	user1 := &UsersGitFlow{}
	user1.FlowdockName = "User1Flowdock"
	user1.GithubName = "User1Github"

	users := getUsers("users.json")

	u := append(users.Users_Git_Flow, *user1)
	count_elements := len(u)
	last_user := u[count_elements-1]
	if last_user.GithubName != "User1Github" {
		t.Fatal("Reading users from json has failed")
	}
}
