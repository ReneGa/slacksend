package payload_test

import (
	"testing"

	"github.com/ReneGa/slacksend/payload"
)

func TestGenerateBody(t *testing.T) {
	// Given
	testCases := []struct {
		given        payload.MessageParams
		bodyExpected string
	}{
		{
			given: payload.MessageParams{
				Channel:  "alerts",
				Username: "alert",
				Text:     "sample message",
			},
			bodyExpected: `{"channel":"#alerts","username":"alert","text":"sample message"}`,
		},
		{
			given: payload.MessageParams{
				Channel:  "testChannel",
				Username: "testUser",
				Text:     "another sample message",
			},
			bodyExpected: `{"channel":"#testChannel","username":"testUser","text":"another sample message"}`,
		},
	}

	for _, testCase := range testCases {
		// When
		bodyResult := payload.GenerateBody(testCase.given)

		// Then
		bodyExpected := testCase.bodyExpected
		if bodyResult != bodyExpected {
			t.Errorf("%s != %s", bodyResult, bodyExpected)
		}
	}
}

func TestGenerateBodyDoNotAddHashSignIfGiven(t *testing.T) {
	// Given
	messageParams := payload.MessageParams{
		Channel:  "#testChannel",
		Username: "testUser",
		Text:     "another sample message",
	}

	// When
	bodyResult := payload.GenerateBody(messageParams)

	// Then
	bodyExpected := `{"channel":"#testChannel","username":"testUser","text":"another sample message"}`
	if bodyResult != bodyExpected {
		t.Errorf("%s != %s", bodyResult, bodyExpected)
	}
}
