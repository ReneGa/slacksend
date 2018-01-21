package slack

import (
	"fmt"
	"net/url"
)

// Request returns the URL and body to use for sending an http request to Slack's incoming webhook
func Request(secret, body string) (string, url.Values) {
	slackURL := fmt.Sprintf("https://hooks.slack.com/services/%s", secret)
	formData := url.Values{
		"payload": []string{body},
	}
	return slackURL, formData
}
