package algo

type Queue[T any] struct {
	list []T
}

func (q *Queue[T]) PushBack(item T) {
	q.list = append(q.list, item)
}

func (q *Queue[T]) TakeFront() (T, bool) {
	if len(q.list) == 0 {
		var zero T
		return zero, false
	}

	item := q.list[0]
	q.list = q.list[1:]
	return item, true
}

func (q *Queue[T]) Count() int {
	return len(q.list)
}
