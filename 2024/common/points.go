package common

import "fmt"

type Point struct {
	X, Y int
}

func (point *Point) String() string {
	return FormatAsPointString(point.X, point.Y)
}

func (point *Point) FromString(s string) {
	fmt.Sscanf(s, "%d,%d", &point.X, &point.Y)
}

func FormatAsPointString(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}
