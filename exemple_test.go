// Copyright 2019 The Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
