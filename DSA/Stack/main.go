package main

type Stack struct {
	items []int
}

func (s *Stack) Push(value int) {
	s.items = append(s.items, value)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}

	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top, true
}
