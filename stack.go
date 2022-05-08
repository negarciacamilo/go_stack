package go_stack

type Stack interface {
	IsEmpty() bool
	Push(elem any)
	Pop() any
	Size() int
	Peek() any
	Clear()
}

type stack struct {
	data []any
}

func NewStack[T any]() Stack {
	return &stack{data: []any{}}
}

func (s *stack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *stack) Push(elem any) {
	s.data = append(s.data, elem)
}

func (s *stack) Pop() any {
	if !s.IsEmpty() {
		elem := s.data[s.Size()-1]
		s.data = s.data[:s.Size()-1]
		return elem
	}
	return nil
}

func (s *stack) Peek() any {
	if !s.IsEmpty() {
		return s.data[s.Size()-1]
	}
	return nil
}

func (s *stack) Clear() {
	if !s.IsEmpty() {
		s.data = []any{}
	}
}

func (s *stack) Size() int {
	return len(s.data)
}
