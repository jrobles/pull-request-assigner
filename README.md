# git-webhook-api [![Build Status](https://travis-ci.org/josemrobles/git-webhook-api.svg?branch=master)](https://travis-ci.org/josemrobles/git-webhook-api) [![Go Report Card](https://goreportcard.com/badge/github.com/josemrobles/git-webhook-api)](https://goreportcard.com/report/github.com/josemrobles/git-webhook-api)
Creates an API which accepts a git webhook post for new PR's. Once received, two dev team members are selected as reviewers then a notification is sent using my [robification-go](https://github.com/josemrobles/robification-go) library.

***Example:***
```
curl -X POST -H "token: 37f7f7446d64345dd367744428837fe5" -H "Content-Type: application/json" -d '{{"action": "opened","number": 280,"pull_request": {"html_url": "https://github.com/orgname/repo/pull/280","user": {"login": "josemrobles"}}}}' 'http://localhost:8008/'
```

