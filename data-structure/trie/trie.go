package main

import "fmt"

//Alphaber Size is the number of possible characters in the trie
const AlphabetSize = 26

// Node represents each node in the trie
type Node struct {
	children [AlphabetSize]*Node
	isEnd    bool
}

// Trie represents a trie and has a pointer to the root node
type Trie struct {
	root *Node
}

// InitTrie will create new Trie
func InitTrie() *Trie {
	result := &Trie{root: &Node{}}
	return result
}

// Insert will take in a word and add it to the trie
func (t *Trie) Insert(word string) {
	currentNode := t.root
	for i := 0; i < len(word); i++ {
		charIndex := word[i] - 'a' //  replace the code of the number for or own code
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

// Search will take in a word and RETURN true is that that is included in the trie
func (t *Trie) Search(word string) bool {
	currentNode := t.root
	for i := 0; i < len(word); i++ {
		charIndex := word[i] - 'a' //  replace the code of the number for or own code
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	if currentNode.isEnd == true { // it's the end of the world?
		return true
	}

	return false // in case the  is end is not true
}

func main() {
	firstTrie := InitTrie()
	toAdd := []string{
		"language", "lambda", "linux", "linked", "learn",
	}

	for _, word := range toAdd {
		firstTrie.Insert(word)
	}

	fmt.Println(firstTrie.Search("language"))
	fmt.Println(firstTrie.Search("linux"))
	fmt.Println(firstTrie.Search("loop"))
	fmt.Println(firstTrie.Search("long"))
}
