# Pull Request Assigner [![Go Report Card](https://goreportcard.com/badge/github.com/josemrobles/pull-request-assigner)](https://goreportcard.com/report/github.com/josemrobles/git-webhook-api)

Creates an API which accepts a git webhook post for new PR's. Once received, two dev team members are selected as reviewers then a Flowdock notification is sent using my [robification-go](https://github.com/josemrobles/robification-go) library.
