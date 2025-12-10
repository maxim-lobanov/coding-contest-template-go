package algo

func Filter[T any](input []T, predicate func(T) bool) []T {
	var result []T
	for _, item := range input {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

func Map[T any, U any](input []T, selector func(T) U) []U {
	result := make([]U, len(input))
	for i, item := range input {
		result[i] = selector(item)
	}
	return result
}

func Reduce[T any, U any](input []T, initial U, accumulator func(U, T) U) U {
	result := initial
	for _, item := range input {
		result = accumulator(result, item)
	}
	return result
}

func Unique[T comparable](input []T) []T {
	seen := make(map[T]struct{})
	var result []T
	for _, item := range input {
		if _, ok := seen[item]; !ok {
			seen[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

func Any[T any](input []T, predicate func(T) bool) bool {
	for _, item := range input {
		if predicate(item) {
			return true
		}
	}
	return false
}

func All[T any](input []T, predicate func(T) bool) bool {
	for _, item := range input {
		if !predicate(item) {
			return false
		}
	}
	return true
}

func First[T any](input []T, predicate func(T) bool) (T, bool) {
	for _, item := range input {
		if predicate(item) {
			return item, true
		}
	}
	var zero T
	return zero, false
}
