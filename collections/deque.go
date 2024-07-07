package collections

type Node[T any] struct {
	value T
	next  *Node[T]
	prev  *Node[T]
}

type Deque[T any] struct {
	head   *Node[T]
	tail   *Node[T]
	length int
}

func NewDeque[T any]() Deque[T] {
	return Deque[T]{}
}

func (d *Deque[T]) Length() int {
	return d.length
}

func (d *Deque[T]) Append(item T) {
	newNode := Node[T]{value: item}
	if d.tail == nil {
		d.tail = &newNode
		d.head = &newNode
	} else {
		newNode.prev = d.tail
		d.tail.next = &newNode
		d.tail = d.tail.next
	}
	d.length += 1
}

func (d *Deque[T]) Pop() (T, error) {
	var emptyElem T
	if d.Length() == 0 {
		return emptyElem, &EmptyDequeError{"Cannot pop from empty deque"}
	}
	tailElem := d.tail.value
	if d.head == d.tail {
		d.head = nil
		d.tail = nil
	} else {
		d.tail = d.tail.prev
		d.tail.next = nil
	}
	d.length -= 1
	return tailElem, nil
}

func (d *Deque[T]) AppendLeft(item T) {
	newNode := Node[T]{value: item}
	if d.head == nil {
		d.head = &newNode
		d.tail = &newNode
	} else {
		newNode.next = d.head
		d.head.prev = &newNode
		d.head = d.head.prev
	}
	d.length += 1
}

func (d *Deque[T]) PopLeft() (T, error) {
	var emptyElem T
	if d.Length() == 0 {
		return emptyElem, &EmptyDequeError{"Cannot pop from empty deque"}
	}
	headElem := d.head.value
	if d.head == d.tail {
		d.head = nil
		d.tail = nil
	} else {
		d.head = d.head.next
		d.head.prev = nil
	}
	d.length -= 1
	return headElem, nil
}
