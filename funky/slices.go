package funky

func Map[T any, R any](data []T, transformer func(T) R) []R {
	result := make([]R, len(data))
	for idx, elem := range data {
		result[idx] = transformer(elem)
	}
	return result
}

func MapWithIndex[T any, R any](data []T, transformer func(int, T) R) []R {
	result := make([]R, len(data))
	for idx, elem := range data {
		result[idx] = transformer(idx, elem)
	}
	return result
}

func Find[T any](data []T, matcher func(T) bool) (int, Option[T]) {
	for idx, elem := range data {
		if matcher(elem) {
			return idx, SomeOf(elem)
		}
	}
	return -1, None[T]{}
}

func AnyMatches[T any](data []T, matcher func(T) bool) bool {
	for _, elem := range data {
		if matcher(elem) {
			return true
		}
	}
	return false
}

func NoneMatches[T any](data []T, matcher func(T) bool) bool {
	return !AnyMatches(data, matcher)
}

func Filter[T any](data []T, matcher func(T) bool) []T {
	result := make([]T, 0)
	for idx, elem := range data {
		if matcher(elem) {
			result = append(result, data[idx])
		}
	}
	return result
}
