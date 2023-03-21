package slack

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
)

// Send sends message (Payload) to the given slack hook URL.
func Send(hookURL string, message Message) error {
	bts, err := json.Marshal(message)
	if err != nil {
		return ErrSerializeMessage
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	client := &http.Client{Transport: tr}
	req, err := http.NewRequest("POST", hookURL, bytes.NewReader(bts))
	if err != nil {
		return ErrCreateRequest
	}

	res, err := client.Do(req)
	if err != nil {
		return ErrSendingRequest
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		return fmt.Errorf("error sending slack message. Status: %v", res.StatusCode)
	}

	return nil
}

// Message represent a Slack message.
type Message struct {
	Parse       string       `json:"parse,omitempty"`
	Username    string       `json:"username,omitempty"`
	IconUrl     string       `json:"icon_url,omitempty"`
	IconEmoji   string       `json:"icon_emoji,omitempty"`
	Channel     string       `json:"channel,omitempty"`
	Text        string       `json:"text,omitempty"`
	LinkNames   string       `json:"link_names,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty"`
	UnfurlLinks bool         `json:"unfurl_links,omitempty"`
	UnfurlMedia bool         `json:"unfurl_media,omitempty"`
	Markdown    bool         `json:"mrkdwn,omitempty"`
}

// Attachment let you add more context to a message, making them more useful and effective.
// See https://api.slack.com/docs/message-attachments
type Attachment struct {
	Fallback     string   `json:"fallback"`
	Color        string   `json:"color"`
	PreText      string   `json:"pretext"`
	AuthorName   string   `json:"author_name"`
	AuthorLink   string   `json:"author_link"`
	AuthorIcon   string   `json:"author_icon"`
	Title        string   `json:"title"`
	TitleLink    string   `json:"title_link"`
	Text         string   `json:"text"`
	ImageUrl     string   `json:"image_url"`
	Fields       []Field  `json:"fields"`
	Footer       string   `json:"footer"`
	FooterIcon   string   `json:"footer_icon"`
	Timestamp    int64    `json:"ts"`
	MarkdownIn   []string `json:"mrkdwn_in"`
	Actions      []Action `json:"actions"`
	CallbackID   string   `json:"callback_id"`
	ThumbnailUrl string   `json:"thumb_url"`
}

// AddField appends a new field to the Attachment
func (attachment *Attachment) AddField(field Field) *Attachment {
	attachment.Fields = append(attachment.Fields, field)
	return attachment
}

// AddAction appends a new Action to the Attachment
func (attachment *Attachment) AddAction(action Action) *Attachment {
	attachment.Actions = append(attachment.Actions, action)
	return attachment
}

// Field is defined as a dictionary with key-value pairs.
type Field struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

// Action make message interactive
type Action struct {
	Type  string `json:"type"`
	Text  string `json:"text"`
	Url   string `json:"url"`
	Style string `json:"style"`
}
