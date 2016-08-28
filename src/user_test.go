package main

import (
	"testing"
)

var (
	early   = Reviewer{Github: "EarlyCuyler", Flowdock: "bootyHunter"}
	rusty   = Reviewer{Github: "RustyCuyler", Flowdock: "rustyRules"}
	durwood = Reviewer{Github: "DurwoodCuyler", Flowdock: "notASquid"}
	lil     = Reviewer{Github: "LilCuyler", Flowdock: "tooDrunk2Care"}
	dan     = Reviewer{Github: "DanHalen", Flowdock: "theWorldIsMine"}

	prAuthor = early.Github

	reviewers = []Reviewer{
		early,
		rusty,
		durwood,
		lil,
		dan,
	}

	users = Config{
		Users_Git_Flow: reviewers,
	}
)

func BenchmarkSwap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		swap(reviewers, 1, 2)
	}
}

func BenchmarkSelectReviewers(b *testing.B) {
	for n := 0; n < b.N; n++ {
		selectReviewers(prAuthor, users)
	}
}

func TestSwap(t *testing.T) {
	swap(reviewers, 1, 2)
	if reviewers[1] == early {
		t.Fatal("ERROR: user.swap failed")
	}
	if reviewers[2] == durwood {
		t.Fatal("ERROR: user.swap failed")
	}
}

func TestSelectReviewers(t *testing.T) {
	reviewerA, reviewerB := selectReviewers(prAuthor, users)
	if reviewerA.Flowdock == "" {
		t.Fatal("rev1 can not be empty")
	}
	if reviewerB.Github == "" {
		t.Fatal("rev2 can not be empty")
	}
	if reviewerA == reviewerB {
		t.Fatal("rev2 and rev1 can not be the same")
	}
	if reviewerA.Github == prAuthor || reviewerB.Github == prAuthor {
		t.Fatal("revs can not be equal to the prOwner")
	}
}
