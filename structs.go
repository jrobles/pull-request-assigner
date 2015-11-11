package main

type gitPayload struct {
	Action       string `json:"action"`
	Number       int    `json:"number"`
	Pull_Request struct {
		Html_Url string `json:"html_url"`
		Head     struct {
			Repo struct {
				Name string `json:"name"`
			} `json:"repo"`
		} `json:"head"`
		User struct {
			Login string `json:"login"`
		} `json:"user"`
	} `json:"pull_request"`
}

type JSONConfigData struct {
	Fd_Token string `json:"fd_token"`
}

type UsersData struct {
	Users_Git_Flow []UsersGitFlow `json:"users_git_flow"`
}

type UsersGitFlow struct {
	GithubName   string `json:"github_name"`
	FlowdockName string `json:"flowdock_name"`
}
