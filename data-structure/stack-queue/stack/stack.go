package main

import "fmt"

// Stack represents stack that hold a slice
type Stack struct {
	items []int
}

// Push - push will add a value at the end
func (s *Stack) Push(a int) {
	s.items = append(s.items, a)
}

//Pop - pop will remove a value at the end
// and return the removed value
func (s *Stack) Pop() int {
	length := len(s.items) - 1
	toRemove := s.items[length]
	s.items = s.items[:length] //this means it's going start from the beggining and it's going to just leave one out at the end and replace items
	return toRemove
}

func main() {
	myStack := Stack{}
	fmt.Println(myStack)
	myStack.Push(100)
	myStack.Push(200)
	myStack.Push(300)
	fmt.Println(myStack)
	myStack.Pop()
	fmt.Println(myStack)
	myStack.Pop()
	fmt.Println(myStack)
	myStack.Pop()
	fmt.Println(myStack)
}
