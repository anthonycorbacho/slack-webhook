package slack_test

import (
	"fmt"

	"github.com/anthonycorbacho/slack-webhook"
)

func ExemplBasic() {
	hookURL := "<slack_hook_url>"

	attachment1 := slack.Attachment{}
	attachment1.AddField(slack.Field{Title: "Field test", Value: "Field value"})

	msg := slack.Message{
		Text:        "This is a slack message content",
		Attachments: []slack.Attachment{attachment1},
	}
	err := slack.Send(hookURL, msg)
	if err != nil {
		fmt.Printf("failed to send message: %v\n", err)
	}
}
