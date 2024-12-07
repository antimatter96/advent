package common

type Stack[K comparable] struct {
	arr []K
}

func CopyStack[K comparable](old Stack[K]) Stack[K] {
	newArr := make([]K, len(old.arr))

	copy(newArr, old.arr)
	return Stack[K]{arr: newArr}
}

func NewStackFromSlice[K comparable](old []K) Stack[K] {
	newArr := make([]K, len(old))

	copy(newArr, old)
	return Stack[K]{arr: newArr}
}

func (stk *Stack[K]) Empty() bool {
	return stk.arr == nil || len(stk.arr) == 0
}

func (stk *Stack[K]) Top() K {
	return stk.arr[len(stk.arr)-1]
}

func (stk *Stack[K]) Pop() K {
	ret := stk.arr[len(stk.arr)-1]

	stk.arr = stk.arr[:len(stk.arr)-1]
	return ret
}

func (stk *Stack[K]) Push(ele K) {
	if stk.arr == nil {
		stk.arr = make([]K, 0)
	}
	stk.arr = append(stk.arr, ele)
}

func (stk *Stack[K]) Reverse() {
	if stk.arr == nil {
		return
	}
	if len(stk.arr) < 2 {
		return
	}
	temp := make([]K, len(stk.arr))
	copy(temp, stk.arr)

	for i := len(stk.arr) - 1; i > -1; i-- {
		stk.arr[i] = temp[len(stk.arr)-1-i]
	}
}

func (stk *Stack[K]) Equals(other *Stack[K]) bool {
	if stk.arr == nil && other.arr == nil {
		return true
	}
	if stk.arr == nil || other.arr == nil {
		return false
	}

	if len(stk.arr) != len(other.arr) {
		return false
	}

	for i := 0; i < len(stk.arr); i++ {
		if stk.arr[i] != other.arr[i] {
			return false
		}
	}

	return true
}

func (stk *Stack[K]) Length() int {
	if stk.arr == nil {
		return 0
	}
	return len(stk.arr)
}
