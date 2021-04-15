package main

import "fmt"

// MaxHeap struct has a slice that holds the array
type MaxHeap struct {
	array []int
}

// Insert adds an element to the heap
func (h *MaxHeap) insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

// Extract returns the largest key, and removes it  from the heap
func (h *MaxHeap) Extract() int {
	extracted := h.array[0]
	l := len(h.array) - 1

	// When the array is empty
	if len(h.array) == 0 {
		fmt.Println("Cannot extract because array length is 0")
		return -1
	}

	//Take out the last index and put it in root
	h.array[0] = h.array[l]
	h.array = h.array[:l]

	h.maxHeapifyDown(0)
	return extracted
}

// maxHeapifyUp will heapify from bottom to the top
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index) // this part is swapping until it finds its place
		index = parent(index)        // updating the current index to its parent index
	}
}

// maxHeapifyDown will heapify from top to the botoom
func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	//Loop whilw index has at least one chidld
	for l <= lastIndex {
		switch {
		case l == lastIndex: // when left child is the only child
			childToCompare = l
		case h.array[l] > h.array[r]: // when left child is larger
			childToCompare = l
		default: // when right child is larger
			childToCompare = r
		}
		// Compare array value of current index to larger child and swap if smaller
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

// Return the parent index
func parent(index int) int {
	parent := (index - 1) / 2
	return parent
}

// get the left child index
func left(parent_index int) int {
	leftChild := (parent_index * 2) + 1
	return leftChild
}

//get the right child index
func right(parent_index int) int {
	rightChild := (parent_index * 2) + 2
	return rightChild
}

//swap keys in the array
func (h *MaxHeap) swap(indexOne, indexTwo int) {
	aux := h.array[indexTwo]
	h.array[indexTwo] = h.array[indexOne]
	h.array[indexOne] = aux

}

func main() {
	maxHeap := &MaxHeap{}
	fmt.Println(maxHeap)

	buildHeap := []int{10, 20, 30, 5, 7, 9, 11, 12, 17}
	for _, item := range buildHeap {
		maxHeap.insert(item)
		fmt.Println(maxHeap)
	}
	fmt.Println()
	for i := 0; i < 5; i++ {
		maxHeap.Extract()
		fmt.Println(maxHeap)
	}

	fmt.Println("\n", maxHeap)
}
