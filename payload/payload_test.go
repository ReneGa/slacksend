package payload_test

import (
	"fmt"
	"slacksend/payload"
	"testing"
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
			fmt.Printf("%s != %s", bodyResult, bodyExpected)
			t.Fail()
		}
	}
}
