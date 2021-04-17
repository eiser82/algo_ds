package linkedlist

import "fmt"

// Node represents an element of the list
type Node struct {
	data int
	next *Node
}

// LinkedList is the list itself
type LinkedList struct {
	head   *Node
	length int
}

// Prepend adds an item to the beginning of the list
func (l *linkedList) Prepend(data int) {
	newNode := Node{data: data, next: nil}
	if l.head != nil {
		newNode.next = l.head
		l.head = &newNode
		l.length++
	} else {
		l.head = &newNode
		l.length++
	}
}

// Append: adds an item to the end of the linked list

// Insert(i, t) adds an item t at position i

// RemoveAt(i) removes a node at position i

// IndexOf() returns the position of the item t

//IsEmpty() returns true if the list is empty

// Size() returns the size of the list

// Print() returns the list
func (l *linkedList) Print() {
	if l.head == nil {
		return
	}

	currentNode := l.head
	for currentNode != nil {
		fmt.Println(currentNode.data)
		currentNode = currentNode.next
	}
}
