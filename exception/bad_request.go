package exception

import "strings"

type BadRequest struct {
	Message []string
}

func (b *BadRequest) Error() string {
	return strings.Join(b.Message, ",")
}

func NewBadRequestError(message []string) error {
	return &BadRequest{
		Message: message,
	}
}
