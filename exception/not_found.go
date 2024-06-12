package exception

type NotFound struct {
	Message string
}

func (n *NotFound) Error() string {
	return n.Message
}
