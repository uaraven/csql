package data

import (
	"fmt"
)

type linkedListElement[T any] struct {
	prev  *linkedListElement[T]
	next  *linkedListElement[T]
	value T
}

func (lle *linkedListElement[T]) Next() LinkedListElement[T] {
	return lle.next
}

func (lle *linkedListElement[T]) Previous() LinkedListElement[T] {
	return lle.prev
}

func (lle *linkedListElement[T]) Value() T {
	return lle.value
}

type linkedList[T any] struct {
	first *linkedListElement[T]
	last  *linkedListElement[T]
}

func (l *linkedList[T]) Append(value T) {
	elem := &linkedListElement[T]{
		prev:  nil,
		next:  nil,
		value: value,
	}
	if l.first == nil {
		l.first = elem
		l.last = elem
	} else {
		l.last.next = elem
		elem.prev = l.last
		l.last = elem
	}
}

func (l *linkedList[T]) Prepend(value T) {
	elem := &linkedListElement[T]{
		prev:  nil,
		next:  nil,
		value: value,
	}
	if l.first == nil {
		l.first = elem
		l.last = elem
	} else {
		l.first.prev = elem
		elem.next = l.first
		l.first = elem
	}
}

func (l *linkedList[T]) AppendAll(values []T) {
	for _, v := range values {
		l.Append(v)
	}
}

func (l *linkedList[T]) Clear() {
	l.first = nil
	l.last = nil
}

func (l *linkedList[T]) removeItem(item *linkedListElement[T]) {
	if item == l.first && item == l.last {
		l.first = nil
		l.last = nil
	} else if item == l.first {
		l.first = item.next
		item.next.prev = nil
	} else if item == l.last {
		l.last = item.prev
		item.prev.next = nil
	} else {
		item.prev.next = item.next
		item.next.prev = item.prev
	}
}

func (l *linkedList[T]) RemoveByIndex(index int) T {
	idx := 0
	item := l.first
	for item != nil {
		if idx == index {
			l.removeItem(item)
			return item.value
		}
		idx++
		item = item.next
	}
	panic(fmt.Sprintf("invalid index %d", index))
}

func (l *linkedList[T]) RemoveFirst(matcher func(T) bool) {
	item := l.first
	for item != nil {
		if matcher(item.value) {
			l.removeItem(item)
			return
		}
		item = item.next
	}
}

func (l *linkedList[T]) IndexOf(matcher func(T) bool) int {
	idx := 0
	item := l.first
	for item != nil {
		if matcher(item.value) {
			return idx
		}
		idx++
		item = item.next
	}
	return -1
}

func (l *linkedList[T]) Size() int {
	counter := 0
	item := l.first
	for item != nil {
		counter++
		item = item.next
	}
	return counter
}

func (l *linkedList[T]) Get(index int) T {
	idx := 0
	item := l.first
	for item != nil {
		if idx == index {
			return item.value
		}
		idx++
		item = item.next
	}
	panic(fmt.Sprintf("index %d is out of bounds", index))
}

func (l *linkedList[T]) First() T {
	if l.first == nil {
		panic(fmt.Sprintf("cannot get the first element of an empty list"))
	} else {
		return l.first.value
	}
}

func (l *linkedList[T]) Last() T {
	if l.last == nil {
		panic(fmt.Sprintf("cannot get the last element of an empty list"))
	} else {
		return l.last.value
	}
}

func (l *linkedList[T]) Rest() List[T] {
	if l.first == nil {
		return NewLinkedList[T]()
	} else {
		return newLListFromFirstLast(l.first.next, l.last)
	}
}

func (l *linkedList[T]) AsSlice() []T {
	iter := l.Iterator()
	result := make([]T, 0)
	for iter.HasNext() {
		result = append(result, iter.Next())
	}
	return result
}

func (l *linkedList[T]) Empty() bool {
	return l.first == nil
}

func (l *linkedList[T]) Pop() T {
	if l.Empty() {
		panic(fmt.Sprintf("list is empty"))
	}
	return l.RemoveByIndex(0)
}

func (l *linkedList[T]) Iterator() Iterator[T] {
	return newLinkedListIterator[T](&linkedListElement[T]{next: l.first})
}

func (l *linkedList[T]) FirstElement() LinkedListElement[T] {
	return l.first
}

func (l *linkedList[T]) LastElement() LinkedListElement[T] {
	return l.last
}

func newLListFromFirstLast[T any](first *linkedListElement[T], last *linkedListElement[T]) LinkedList[T] {
	return &linkedList[T]{
		first: first,
		last:  last,
	}
}

func NewLinkedList[T any]() LinkedList[T] {
	return &linkedList[T]{}
}

func NewLinkedListFromSlice[T any](values []T) LinkedList[T] {
	ll := &linkedList[T]{}
	ll.AppendAll(values)
	return ll
}
