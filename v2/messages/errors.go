package messages

type MessageNotReceivedError struct {
	FailedMessageList []FailedMessage
	TotalCount        int
}

func (e *MessageNotReceivedError) Error() string {
	return "all messages failed to be registered"
}
