package common

type NextFunc func(i int) int

func NoChange(i int) int {
	return i
}
func Inc1(i int) int {
	return i + 1
}
func Dec1(i int) int {
	return i - 1
}

type Directions struct {
	I, J         int
	NextI, NextJ NextFunc
}

func GenerateDirections(i, j int) map[string]*Directions {
	return map[string]*Directions{
		"UP":   {i - 1, j, Dec1, NoChange},
		"DOWN": {i + 1, j, Inc1, NoChange},

		"LEFT":  {i, j - 1, NoChange, Dec1},
		"RIGHT": {i, j + 1, NoChange, Inc1},

		"UP-LEFT":    {i - 1, j - 1, Dec1, Dec1},
		"UP-RIGHT":   {i - 1, j + 1, Dec1, Inc1},
		"DOWN-RIGHT": {i + 1, j + 1, Inc1, Inc1},
		"DOWN-LEFT":  {i + 1, j - 1, Inc1, Dec1},
	}

}

/*
BackSlash  ForwardSlash
\..             ../
.X.      VS     .X.
..\             /..
*/
