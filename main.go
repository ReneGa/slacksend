package main

import (
	"flag"
	"log"

	"net/http"

	"github.com/ReneGa/slacksend/flags"
	"github.com/ReneGa/slacksend/payload"
	"github.com/ReneGa/slacksend/slack"
)

func main() {
	// Collect and validate command line flags
	channelFlag := flag.String("channel", "general", "the slack channel to slack to")
	secretFlag := &flags.Required{
		Name:  "secret",
		Usage: "the secret part of the API URL (https://hooks.slack.com/services/<secret>)",
	}
	textFlag := &flags.Required{
		Name:  "text",
		Usage: "the text to send as a message",
	}
	requiredFlags := []*flags.Required{secretFlag, textFlag}
	for _, requiredFlag := range requiredFlags {
		flag.StringVar(&requiredFlag.Value, requiredFlag.Name, "", requiredFlag.Usage)
	}
	flag.Parse()
	flags.AssertRequired(requiredFlags)

	payloadJSON := payload.GenerateBody(payload.MessageParams{
		Channel:  *channelFlag,
		Username: "slacksend",
		Text:     textFlag.Value,
	})

	_, err := http.PostForm(slack.Request(secretFlag.Value, payloadJSON))
	if err != nil {
		log.Fatal(err)
	}
}
