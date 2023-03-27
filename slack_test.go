package slack

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func handler() http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("/services", testHandler)
	return r
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	body := r.Body
	if body == nil {
		http.Error(w, "missing value", http.StatusBadRequest)
		return
	}
	defer body.Close()

	var msg Message

	b, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(b, &msg)
	if err != nil {
		http.Error(w, "cannot unmarshal body", http.StatusBadRequest)
		return
	}
}

func TestSlack(t *testing.T) {
	srv := httptest.NewServer(handler())
	defer srv.Close()

	slackURL := fmt.Sprintf("%s/services", srv.URL)

	attachment1 := Attachment{}
	attachment1.AddField(Field{Title: "Field", Value: "Field test value"})

	msg := Message{
		Text:        "the is a slack test massage",
		Attachments: []Attachment{attachment1},
	}

	err := Send(slackURL, msg)
	if err != nil {
		t.Error(err)
	}
}
