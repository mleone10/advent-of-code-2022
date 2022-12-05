package stack

type Stack[T any] struct {
	stack []T
}

func (s *Stack[T]) Push(item T) {
	s.stack = append(s.stack, item)
}

func (s *Stack[T]) Pop() T {
	ret := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return ret
}

func (s *Stack[T]) Len() int {
	return len(s.stack)
}
