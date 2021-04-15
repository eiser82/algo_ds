package linkedlist

import "fmt"

// Node represents an element of the list
type node struct {
	data int
	next *node
}

// LinkedList is the list itself
type linkedList struct {
	head   *node
	length int
}

// prepend func inserts a new element
// at the begining of the list
func (l *linkedList) prepend(data int) {
	newNode := node{data: data, next: nil}
	if l.head != nil {
		newNode.next = l.head
		l.head = &newNode
		l.length++
	} else {
		l.head = &newNode
		l.length++
	}
}

// prints the list
func (l *linkedList) print() {
	if l.head == nil {
		return
	}

	currentNode := l.head
	for currentNode != nil {
		fmt.Println(currentNode.data)
		currentNode = currentNode.next
	}
}

// func (l *linkedList) deleteWithValue(value int) {

// }
