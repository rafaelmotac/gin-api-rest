package problem

type MessageErrorCode struct {
	Message   string
	ErrorCode int
}

func (e *MessageErrorCode) Error() string {
	return e.Message
}

func (e *MessageErrorCode) Kind() int {
	return e.ErrorCode
}
