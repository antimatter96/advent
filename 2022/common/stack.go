package common

type Stack[K comparable] struct {
	arr []K
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
