package main

type MyStack struct {
	Queue []int
}

func Constructor() MyStack {
	return MyStack{
		Queue: make([]int, 0),
	}
}

func (s *MyStack) Push(x int) {
	s.Queue = append(s.Queue, x)
}

func (s *MyStack) Pop() int {
	n := len(s.Queue) - 1
	for n != 0 {
		tmp := s.Queue[0]
		s.Queue = s.Queue[1:]
		s.Queue = append(s.Queue, tmp)
		n--
	}
	res := s.Queue[0]
	s.Queue = s.Queue[1:]
	return res
}

func (s *MyStack) Top() int {
	tmp := s.Pop()
	s.Queue = append(s.Queue, tmp)
	return tmp
}

func (s *MyStack) Empty() bool {
	if len(s.Queue) == 0 {
		return true
	}
	return false
}

func main() {

}
