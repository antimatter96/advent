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

var DirectionChanges = map[string]*Directions{
	"UP":   &Directions{NextI: Dec1, NextJ: NoChange},
	"DOWN": &Directions{NextI: Inc1, NextJ: NoChange},

	"LEFT":  &Directions{NextI: NoChange, NextJ: Dec1},
	"RIGHT": &Directions{NextI: NoChange, NextJ: Inc1},

	"UP-LEFT":    &Directions{NextI: Dec1, NextJ: Dec1},
	"UP-RIGHT":   &Directions{NextI: Dec1, NextJ: Inc1},
	"DOWN-RIGHT": &Directions{NextI: Inc1, NextJ: Inc1},
	"DOWN-LEFT":  &Directions{NextI: Inc1, NextJ: Dec1},
}

var RotationRight = map[string]string{
	"UP":    "RIGHT",
	"LEFT":  "UP",
	"RIGHT": "DOWN",
	"DOWN":  "LEFT",
}

func GenerateDirections(i, j int) map[string]*Directions {
	mp := make(map[string]*Directions)

	for dirString, dir := range DirectionChanges {
		mp[dirString] = &Directions{I: dir.NextI(i), J: dir.NextJ(j), NextI: dir.NextI, NextJ: dir.NextJ}
	}

	return mp
}

/*
BackSlash  ForwardSlash
\..             ../
.X.      VS     .X.
..\             /..
*/
