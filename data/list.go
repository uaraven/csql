package data

type List[T any] interface {
	Append(T)
	AppendAll([]T)
	Clear()
	RemoveByIndex(int) T
	RemoveFirst(func(T) bool)
	IndexOf(func(T) bool) int
	Size() int
	Get(int) T
	Empty() bool

	First() T
	Last() T
	Rest() List[T]
	AsSlice() []T
	Iterator() Iterator[T]
}

type LinkedListElement[T any] interface {
	Next() LinkedListElement[T]
	Previous() LinkedListElement[T]
	Value() T
}

type LinkedList[T any] interface {
	List[T]
	FirstElement() LinkedListElement[T]
	LastElement() LinkedListElement[T]
	Prepend(T)
	Pop() T
}

type Iterator[T any] interface {
	HasNext() bool
	Next() T
}
