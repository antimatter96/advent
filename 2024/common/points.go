package common

import "fmt"

type Point struct {
	X, Y int
}

func (point *Point) String() string {
	return fmt.Sprintf("%d,%d", point.X, point.Y)
}
