package common

type RuneStack struct {
	arr []rune
}

func (stk *RuneStack) Push(a rune) {
	stk.arr = append(stk.arr, a)
}

func (stk *RuneStack) Pop() rune {
	n := len(stk.arr)
	a := stk.arr[n-1]
	stk.arr = stk.arr[:n-1]
	return a
}

func (stk *RuneStack) Empty() bool {
	return len(stk.arr) == 0
}

func (stk *RuneStack) Top() rune {
	return stk.arr[len(stk.arr)-1]
}
