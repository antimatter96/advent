package main

import "sort"

type Queue struct {
	arr  []string
	dist *map[string]int
}

func (q *Queue) Add(s string) {
	q.arr = append(q.arr, s)
}

func (q *Queue) Len() int {
	return len(q.arr)
}

func (q *Queue) Empty() bool {
	return len(q.arr) == 0
}

func (q *Queue) Least() string {
	sort.SliceStable(q.arr, func(i, j int) bool {
		return (*q.dist)[q.arr[i]] < (*q.dist)[q.arr[j]]
	})

	a := q.arr[0]
	q.arr = q.arr[1:]

	return a
}
