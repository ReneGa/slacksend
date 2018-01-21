package slack_test

import (
	"testing"

	"github.com/ReneGa/slacksend/slack"
)

func TestRequest(t *testing.T) {
	// Given
	secret := "mySecret"
	payload := `{"field":"value"}`

	// When
	actualURL, actualBody := slack.Request(secret, payload)

	// Then
	expectedURL := "https://hooks.slack.com/services/mySecret"
	if actualURL != expectedURL {
		t.Errorf("expected URL to be %s, but got %s", expectedURL, actualURL)
	}

	expectedPayload := payload
	actualPayload := actualBody.Get("payload")
	if actualPayload != expectedPayload {
		t.Errorf(`expected key "payload" to have value %s, but got %s`, expectedPayload, actualPayload)
	}
}
