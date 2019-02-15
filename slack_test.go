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

func Test_slack(t *testing.T) {
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
