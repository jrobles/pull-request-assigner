package main

import (
	"encoding/json"
	"io/ioutil"
)

type JSONConfigData struct {
	Fd_Token       string         `json:"fd_token"`
	Users_Git_Flow []UsersGitFlow `json:"users_git_flow"`
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
