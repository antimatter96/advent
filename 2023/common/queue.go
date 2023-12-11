package common

type Queue[T any] struct {
	arr []T
}

func (q *Queue[T]) Empty() bool {
	return q.arr == nil || len(q.arr) == 0
}

func (q *Queue[T]) Push(ele T) {
	q.arr = append(q.arr, ele)
}

func (q *Queue[T]) Pop() T {
	old := q.arr
	n := len(old)
	item := old[0]
	q.arr = old[1:n]
	return item
}

func (q *Queue[T]) Size() int {
	n := len(q.arr)
	return n
}
