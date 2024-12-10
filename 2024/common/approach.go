package common

import "flag"

func InitApproachFlags() *int {
	approach := flag.Int("approach", 1, "implementation approach")

	flag.Parse()
	return approach
}
