package common

type QueueSet[T comparable] struct {
	arr []T
	set map[T]struct{}
}

func (q *QueueSet[T]) Empty() bool {
	return q.arr == nil || len(q.arr) == 0
}

func (q *QueueSet[T]) Push(ele T) {
	if _, present := q.set[ele]; present {
		return
	}
	q.arr = append(q.arr, ele)
	if q.set == nil {
		q.set = make(map[T]struct{})
	}
	q.set[ele] = struct{}{}
}

func (q *QueueSet[T]) Pop() T {
	old := q.arr
	n := len(old)
	item := old[0]
	q.arr = old[1:n]

	delete(q.set, item)
	return item
}

func (q *QueueSet[T]) Size() int {
	n := len(q.arr)
	return n
}
