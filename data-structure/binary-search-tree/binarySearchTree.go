package main

import "fmt"

//Node ->> represents the components of a binary search tree
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// Insert will add a node to the tree
// the key to add should not be already in the tree
func (n *Node) InsertInTree(newKey int) {
	switch {
	case n.Key > newKey: //MOVE LEFT
		if n.Left == nil {
			n.Left = &Node{Key: newKey}
		} else {
			n.Left.InsertInTree(newKey)
		}
	case n.Key < newKey: //MOVE RIGHT
		if n.Right == nil {
			n.Right = &Node{Key: newKey}
		} else {
			n.Right.InsertInTree(newKey)
		}
	}
}

//Search will take in a key value
// and RETURN true if there is a node with that value
func (n *Node) SearchInTree(value int) bool {

	if n == nil {
		return false
	}

	switch {
	case n.Key > value: //MOVE LEFT
		return n.Left.SearchInTree(value)
	case n.Key < value: //MOVE RIGHT
		return n.Right.SearchInTree(value)
	}

	return true
}

func main() {
	tree := &Node{Key: 100}
	tree.InsertInTree(200)
	tree.InsertInTree(300)

	fmt.Println(tree)
	fmt.Println(tree.SearchInTree(20))
}
