package common

type Set[T comparable] struct {
	mp map[T]struct{}
}

func (s *Set[T]) Empty() bool {
	return s.mp == nil || len(s.mp) == 0
}

func (s *Set[T]) Add(ele T) {
	if s.mp == nil {
		s.mp = make(map[T]struct{})
	}
	s.mp[ele] = struct{}{}
}

func (s *Set[T]) Remove(ele T) {
	delete(s.mp, ele)
}

func (s *Set[T]) Size() int {
	n := len(s.mp)
	return n
}
