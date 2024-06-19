package main

type Stack struct {
	stackList []int
}

func (s *Stack) isEmpty() bool {
	return len(s.stackList) == 0
}

func (s *Stack) top() int {
	if s.isEmpty() {
		return -1
	}
	return s.stackList[len(s.stackList)-1]
}

func (s *Stack) size() int {
	return len(s.stackList)
}

func (s *Stack) push(value int) {
	s.stackList = append(s.stackList, value)
}

func (s *Stack) pop() int {
	if s.isEmpty() {
		return -1
	}
	topValue := s.stackList[len(s.stackList)-1]
	s.stackList = s.stackList[:len(s.stackList)-1]
	return topValue
}
