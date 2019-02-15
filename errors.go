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

// Slack errors.
const (
	ErrSerializeMessage = Error("couldn't serialise Slack Message")
	ErrCreateRequest    = Error("couldn't create the request")
	ErrSendingRequest   = Error("couldn't send Slack Message")
)

// Error represents a Slack error.
type Error string

// Error returns the error message.
func (e Error) Error() string {
	return string(e)
}
