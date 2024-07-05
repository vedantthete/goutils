package collections

type IndexError struct {
	message string
}

func (i *IndexError) Error() string {
	return i.message
}

type NotFoundError struct {
	message string
}

func (n *NotFoundError) Error() string {
	return n.message
}
