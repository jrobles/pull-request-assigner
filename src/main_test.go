package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	server   *httptest.Server
	reader   io.Reader
	testUrl  string
	testJson = `{"action":"opened","number":"123","pull_request":{"html_url":"https://github.com/pr/url","head":{"repo":{"id":123,"name":"repo","full_name":"josemrobles/repo","owner":{"login":"josemrobles","id":123,"avatar_url":"https://avatars.githubusercontent.com/","name":"josemrobles","type":"User"},"html_url":"https://github.com/josemrobles/repo"}},"base":{"user":{"login":"creativedrive","id":123,"avatar_url":"https://avatars.githubusercontent.com","name":"josemrobles","type":"Organization"},"repo":{"id":123,"name":"repo","full_name":"josemrobles/repo","owner":{"login":"josemrobles","id":123,"avatar_url":"https: //avatars.githubusercontent.com","name":"josemrobles","type":"Organization"}}},"user":{"login":"josemrobles","id":123,"avatar_url":"https: //avatars.githubusercontent.com","name":"josemrobles","type":"Organization"}}}`
)

func init() {
	server = httptest.NewServer(http.HandlerFunc(indexAction)) //Creating new server with the user handlers
	testUrl = fmt.Sprintf("%s/", server.URL)                   //Grab the address for the API endpoint
}

func TestIndexAction(t *testing.T) {

	reader := strings.NewReader(testJson) //Convert string to reader

	request, err := http.NewRequest("POST", testUrl, reader) //Create request with JSON body
	request.Header.Set("Token", "9543195005")

	res, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != 201 {
		t.Fatal("Expected 201 status code, received: ", res.StatusCode) // test failed
	}
}
