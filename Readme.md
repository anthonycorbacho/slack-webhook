## Slack webhook

**Slack webhook** is a simple library that allows application to seamlessly send Slack Message by using the Incoming Webhooks API from Slack.

Get it:
```go
go get -u github.com/anthonycorbacho/slack-webhook
```

### Quick start
You have to setup your Slack Incoming Webhooks, you can refer to this document https://api.slack.com/incoming-webhooks

```go
package main

import (
	"fmt"

	"github.com/anthonycorbacho/slack-webhook"
)

func main() {
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
```
