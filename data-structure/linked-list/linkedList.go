package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

//reciver
func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l linkedList) printListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}

func (l *linkedList) deleteWithValue(value int) {
	if l.length == 0 {
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	previousToDelete := l.head
	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}
func main() {
	mylist := linkedList{}
	node1 := &node{data: 40}
	node2 := &node{data: 18}
	node3 := &node{data: 16}
	node4 := &node{data: 20}
	node5 := &node{data: 50}
	node6 := &node{data: 32}
	mylist.prepend(node1)
	mylist.prepend(node2)
	mylist.prepend(node3)
	mylist.prepend(node4)
	mylist.prepend(node5)
	mylist.prepend(node6)
	mylist.printListData()

	mylist.deleteWithValue(100) //previously: runtime error - this value isn't in the list
	mylist.deleteWithValue(32)  //previously: runtime error- you cannot delete the header
	mylist.deleteWithValue(40)
	mylist.printListData()
	emptylist := linkedList{}
	emptylist.deleteWithValue(10) //previously: this linkedlist is empty, you cannot delete a value in that case
	//Now, go back and handle these especial cases
}
