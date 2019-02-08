package linkedlist

// Node represents an element of the list
type Node struct {
	data int
	next *Node
}

// LinkedList is the list itself
type LinkedList struct {
	head *Node
}
