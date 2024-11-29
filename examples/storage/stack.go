package storage

type Stack struct {
	size     int
	elements [2]int
}

func NewStack() *Stack {
	return &Stack{0, [...]int{0, 0}}
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) Push(element int) {
	s.elements[s.size] = element
	s.size += 1
}

func (s *Stack) Pop() int {
	if s.size == 0 {
		panic("underflow")
	}
	s.size -= 1
	return s.elements[s.size]
}

func (s *Stack) Size() int {
	return s.size
}
