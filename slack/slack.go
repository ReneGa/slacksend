package slack

import (
	"fmt"
	"net/http"
	"net/url"
)

// Send makes a request to Slack's webhook API (see https://api.slack.com/incoming-webhooks)
func Send(secret, body string) (*http.Response, error) {
	slackURL := fmt.Sprintf("https://hooks.slack.com/services/%s", secret)
	return http.PostForm(slackURL, url.Values{
		"payload": []string{body},
	})
}
