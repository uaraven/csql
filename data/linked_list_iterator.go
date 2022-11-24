package data

import (
	"fmt"
	"reflect"
)

type linkedListIterator[T any] struct {
	current LinkedListElement[T]
}

func (l *linkedListIterator[T]) HasNext() bool {
	v := l.current.Next()
	return !(v == nil || reflect.ValueOf(v).IsNil())
}

func (l *linkedListIterator[T]) Next() T {
	if !l.HasNext() {
		panic(fmt.Sprintf("exhausted iterator"))
	}
	l.current = l.current.Next()
	return l.current.Value()
}

func newLinkedListIterator[T any](element LinkedListElement[T]) Iterator[T] {
	return &linkedListIterator[T]{current: element}
}
