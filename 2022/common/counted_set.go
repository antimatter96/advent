package common

type CountedSet[K comparable] struct {
	mp map[K]int
}

func (set *CountedSet[K]) Add(e K) {
	if set.mp == nil {
		set.mp = make(map[K]int)
	}

	set.mp[e]++
}

func (set *CountedSet[K]) Remove(e K) {
	if set.mp == nil {
		set.mp = make(map[K]int)
	}

	set.mp[e]--

	if set.mp[e] == 0 {
		delete(set.mp, e)
	}
}

func (set *CountedSet[K]) UniqueCount() int {
	if set.mp == nil {
		return 0
	}

	return len(set.mp)
}

func (set *CountedSet[K]) Count() int {
	if set.mp == nil {
		return 0
	}

	total := 0

	for _, v := range set.mp {
		total += v
	}

	return total
}
