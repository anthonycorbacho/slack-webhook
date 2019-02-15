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
