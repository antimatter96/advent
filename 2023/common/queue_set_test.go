package common

import (
	"testing"
)

func TestQueueSet(t *testing.T) {
	var isEmpty bool
	var size, popped int

	q := QueueSet[int]{}

	isEmpty = q.Empty()
	if isEmpty != true {
		t.Errorf("expected %v but got %v", true, isEmpty)
	}
	size = q.Size()
	if size != 0 {
		t.Errorf("expected %d but got %d", 0, size)
	}

	q.Push(1)

	isEmpty = q.Empty()
	if isEmpty != false {
		t.Errorf("expected %v but got %v", false, isEmpty)
	}
	size = q.Size()
	if size != 1 {
		t.Errorf("expected %d but got %d", 1, size)
	}

	q.Push(8)

	size = q.Size()
	if size != 2 {
		t.Errorf("expected %d but got %d", 2, size)
	}

	q.Push(8)

	size = q.Size()
	if size != 2 {
		t.Errorf("expected %d but got %d", 2, size)
	}

	popped = q.Pop()

	if popped != 1 {
		t.Errorf("expected %d but got %d", 1, popped)
	}
	size = q.Size()
	if size != 1 {
		t.Errorf("expected %d but got %d", 1, size)
	}

	q.Push(9)

	size = q.Size()
	if size != 2 {
		t.Errorf("expected %d but got %d", 2, size)
	}

	popped = q.Pop()
	if popped != 8 {
		t.Errorf("expected %d but got %d", 8, popped)
	}
	size = q.Size()
	if size != 1 {
		t.Errorf("expected %d but got %d", 1, size)
	}

	popped = q.Pop()
	if popped != 9 {
		t.Errorf("expected %d but got %d", 9, popped)
	}
	isEmpty = q.Empty()
	if isEmpty != true {
		t.Errorf("expected %v but got %v", true, isEmpty)
	}
	size = q.Size()
	if size != 0 {
		t.Errorf("expected %d but got %d", 0, size)
	}

}
