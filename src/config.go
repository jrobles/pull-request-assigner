package main

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Fd_Token        string     `json:"fd_token"`
	Github_Login    string     `json:"github_login"`
	Github_Password string     `json:"github_password"`
	Users_Git_Flow  []Reviewer `json:"users_git_flow"`
}

func getConfigs(configFile string) (configs *Config) {
	configs = &Config{}
	J, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(err)
	}
	json.Unmarshal([]byte(J), &configs)
	return
}
