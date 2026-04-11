package main

import "fmt"

type Node struct {
	Val  int
	Prev *Node
	Next *Node
}

type DoublyLinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

func newNode(val int) *Node {
	return &Node{Val: val}
}

func (dll *DoublyLinkedList) insertTail(val int) {
	node := newNode(val)

	if dll.Tail == nil { // List is empty
		dll.Head = node
		dll.Tail = node
	} else {
		node.Prev = dll.Tail
		dll.Tail.Next = node
		dll.Tail = node
		dll.Size = dll.Size + 1
	}
}

func (dll *DoublyLinkedList) insertAtHead(val int) {
	node := newNode(val)

	if dll.Head == nil && dll.Tail == nil { //List is empty
		dll.Head = node
		dll.Tail = node
	} else {
		node.Next = dll.Head
		dll.Head.Prev = node
		dll.Head = node
		dll.Size = dll.Size + 1
	}
}

func (dll *DoublyLinkedList) removeFromHead() {

	switch {
	case dll.Head == nil:
		fmt.Println("List is empty")

	case dll.Head == dll.Tail:
		dll.Head = nil
		dll.Tail = nil
		dll.Size--

	default:
		pointer := dll.Head
		dll.Head = dll.Head.Next
		dll.Head.Prev = nil
		pointer.Next = nil
		dll.Size--
	}

}

func (dll *DoublyLinkedList) removeFromTail() {
	switch {
	case dll.Tail == nil:
		fmt.Println("List is Empty")

	case dll.Tail == dll.Head:
		dll.Head = nil
		dll.Tail = nil
		dll.Size--

	default:
		pointer := dll.Tail
		dll.Tail = dll.Tail.Prev
		dll.Tail.Next = nil
		pointer.Prev = nil
		dll.Size--
	}
}

func (dll *DoublyLinkedList) InsertBetweenTwoNodes(val int) {

}

func (dll *DoublyLinkedList) RemovBetweenTwoNodes(val int) {

}
