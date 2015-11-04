package main

type gitPayload struct {
	Action       string `json:action`
	Number       int    `json:number`
	Pull_Request struct {
		Html_Url string `json:html_url`
	} `json:pull_request`
	Repo struct {
		Name string `json:name`
	} `json:repo`
}

type JSONConfigData struct {
	Fd_Token string `json:fd_token`
}
