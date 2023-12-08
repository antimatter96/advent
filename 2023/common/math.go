package common

func GCD(a, b int) int {
	var temp int
	for b != 0 {
		temp = b
		b = a % b
		a = temp
	}
	return a
}

func LCM(a, b int) int {
	return (a * b / GCD(a, b))
}

func LCMs(arr []int, i int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	if len(arr)-i == 2 {
		return LCM(arr[i], arr[i+1])
	} else {
		return LCM(arr[i], LCMs(arr, i+1))
	}
}
