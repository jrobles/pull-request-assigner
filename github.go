package main

import (
	"bufio"
	"fmt"
	"github.com/google/go-github/github"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strings"
	"syscall"
)

func githubAuth() {
	r := bufio.NewReader(os.Stdin)
	fmt.Print("GitHub Username: ")
	username, _ := r.ReadString('\n')

	fmt.Print("GitHub Password: ")
	bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	password := string(bytePassword)

	tp := github.BasicAuthTransport{
		Username: strings.TrimSpace(username),
		Password: strings.TrimSpace(password),
	}

	client := github.NewClient(tp.Client())
	user, _, err := client.Users.Get("")
	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return
	}
	fmt.Printf("\n\nLogged in as: %v\n\n", string(*user.Login))
}
