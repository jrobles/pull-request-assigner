package main

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Fd_Token        string        `json:"fd_token"`
	Github_Login    string        `json:"github_login"`
	Github_Password string        `json:"github_password"`
	Users_Git_Flow  []UserGitFlow `json:"users_git_flow"`
}

func getConfigs() (configs *Config) {

	configs = &Config{}
	J, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	json.Unmarshal([]byte(J), &configs)
	return
}
