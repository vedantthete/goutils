package goutils

import (
	"fmt"
	"sort"
)

type List[T any] struct {
	list []T
}

func NewList[T any]() List[T] {
	return List[T]{}
}

func (l *List[T]) Append(value T) {
	l.list = append(l.list, value)
}

func (l *List[T]) Items() []T {
	return l.list
}

func (l *List[T]) Pop() T {
	if l.Length() == 0 {
		panic("Cannot pop from an empty list")
	}
	poppedValue := l.list[l.Length()-1]
	l.list = l.list[:l.Length()-1]
	return poppedValue
}

func (l *List[T]) Contains(value T, equal func(a, b T) bool) bool {
	for _, listItem := range l.list {
		if equal(listItem, value) {
			return true
		}
	}
	return false
}

func (l *List[T]) Length() int {
	return len(l.list)
}

func (l *List[T]) Get(idx int) T {
	if idx >= l.Length() {
		panic("Index out of range")
	}
	if idx < 0 {
		idx = l.Length() + idx
		if idx < 0 {
			panic("Index out of range")
		}
	}
	return l.list[idx]
}

func (l *List[T]) Index(value T, equal func(a, b T) bool) (int, error) {
	for i, listItem := range l.list {
		if equal(value, listItem) {
			return i, nil
		}
	}
	return 0, &NotFoundError{fmt.Sprintf("%v is not in list", value)}
}

func (l *List[T]) Extend(values []T) {
	l.list = append(l.list, values...)
}

func (l *List[T]) Clear() {
	l.list = []T{}
}

func (l *List[T]) Copy() List[T] {
	copiedList := NewList[T]()
	copiedList.list = append(copiedList.list, l.list...)
	return copiedList
}

func (l *List[T]) Insert(idx int, value T) {
	var item T
	l.list = append(l.list, item)
	if idx < 0 {
		idx = 0
	} else if idx >= l.Length() {
		idx = l.Length() - 1
	}
	for i := l.Length() - 1; i > idx; i-- {
		l.list[i] = l.list[i-1]
	}
	l.list[idx] = value
}

func (l *List[T]) Remove(value T, equal func(a, b T) bool) error {
	removalIdx := -1
	for i, v := range l.list {
		if equal(value, v) {
			removalIdx = i
		}
	}
	if removalIdx < 0 {
		return &NotFoundError{fmt.Sprintf("%v does not exist in list", value)}
	}
	for i := removalIdx; i < l.Length()-1; i++ {
		l.list[i] = l.list[i+1]
	}
	l.list = l.list[:l.Length()-1]
	return nil
}

func (l *List[T]) Reverse() {
	for i := 0; i < l.Length()/2; i++ {
		l.list[i], l.list[l.Length()-1-i] = l.list[l.Length()-1-i], l.list[i]
	}
}

func (l *List[T]) Sort(less func(i, j int) bool) {
	sort.Slice(l.list, less)
}
