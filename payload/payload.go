package payload

import (
	"encoding/json"
	"log"
)

// MessageParams represent the JSON fields required by Slacks API for sending messages
type MessageParams struct {
	Channel  string `json:"channel"`
	Username string `json:"username"`
	Text     string `json:"text"`
}

// GenerateBody returns MessageParams marshaled to JSON
func GenerateBody(params MessageParams) string {

	params.Channel = "#" + params.Channel

	j, err := json.Marshal(params)
	if err != nil {
		log.Fatal(err)
	}

	return string(j)
}
