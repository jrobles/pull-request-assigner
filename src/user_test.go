package main

import (
	"testing"
)

func TestSwap(t *testing.T) {
	early := Reviewer{Github: "EarlyCuyler", Flowdock: "bootyHunter"}
	rusty := Reviewer{Github: "RustyCuyler", Flowdock: "rustyRules"}
	durwood := Reviewer{Github: "DurwoodCuyler", Flowdock: "notASquid"}

	reviewers := []Reviewer{
		early,
		rusty,
		durwood,
	}

	swap(reviewers, 1, 2)

	if reviewers[1] == early {
		t.Fatal("ERROR: user.swap failed")
	}
	if reviewers[2] == durwood {
		t.Fatal("ERROR: user.swap failed")
	}
}

func TestSelectReviewers(t *testing.T) {
	u0 := Reviewer{Github: "User0", Flowdock: "user test0"}
	u1 := Reviewer{Github: "User1", Flowdock: "user test1"}
	u2 := Reviewer{Github: "User2", Flowdock: "user test2"}
	u3 := Reviewer{Github: "User3", Flowdock: "user test3"}
	u4 := Reviewer{Github: "User4", Flowdock: "user test4"}
	u5 := Reviewer{Github: "User5", Flowdock: "user test5"}

	temp := []Reviewer{
		u0,
		u1,
		u2,
		u3,
		u4,
		u5,
	}
	users := Config{
		Users_Git_Flow: temp,
	}
	prOwner := "User0"
	rev1, rev2 := selectReviewers(prOwner, users)
	if rev1.Flowdock == "" {
		t.Fatal("rev1 can not be empty")
	}
	if rev2.Github == "" {
		t.Fatal("rev2 can not be empty")
	}
	if rev1 == rev2 {
		t.Fatal("rev2 and rev1 can not be the same")
	}
	if rev1.Github == prOwner || rev2.Github == prOwner {
		t.Fatal("revs can not be equal to the prOwner")
	}
}
