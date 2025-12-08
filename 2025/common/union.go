package common

type Union[K comparable] struct {
	ParentOf map[K]K
	Rank     map[K]int
}

func (union *Union[K]) Init(arr []K) {
	union.ParentOf = make(map[K]K)
	union.Rank = make(map[K]int)

	for _, ele := range arr {
		union.ParentOf[ele] = ele
		union.Rank[ele] = 0
	}
}

func (union *Union[K]) Find(a K) K {
	parent := union.ParentOf[a]
	for parent != union.ParentOf[parent] {
		union.ParentOf[parent] = union.ParentOf[union.ParentOf[parent]]
		parent = union.ParentOf[parent]
	}
	return parent
}

func (union *Union[K]) Add(a, b K) {
	p1, p2 := union.Find(a), union.Find(b)

	if p1 == p2 {
		return
	}

	if union.Rank[p1] > union.Rank[p2] {
		union.ParentOf[p2] = p1
	} else if union.Rank[p1] < union.Rank[p2] {
		union.ParentOf[p1] = p2
	} else {
		union.ParentOf[p2] = p1
		union.Rank[p1]++
	}
}

func (union *Union[K]) Size() int {
	children := 0
	for ele := range union.ParentOf {
		if ele != union.Find(ele) {
			children++
		}
	}

	return children + 1
}

func (union *Union[K]) Bucketize() map[K][]K {
	buckets := make(map[K][]K)
	for ele := range union.ParentOf {
		buckets[union.Find(ele)] = make([]K, 0)
	}

	for ele := range union.ParentOf {
		buckets[union.Find(ele)] = append(buckets[union.Find(ele)], ele)
	}

	return buckets
}
