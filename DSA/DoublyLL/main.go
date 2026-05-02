package main

import "fmt"

// Node represents a single element in the doubly linked list
type Node struct {
	Val  int   // the data stored in this node
	Prev *Node // pointer to the previous node (nil if this is the head)
	Next *Node // pointer to the next node (nil if this is the tail)
}

// DoublyLinkedList holds references to both ends of the list and tracks its size
type DoublyLinkedList struct {
	Head *Node // pointer to the first node
	Tail *Node // pointer to the last node
	Size int   // number of nodes currently in the list
}

// newNode allocates and returns a new node with the given value
func newNode(val int) *Node {
	return &Node{Val: val} // Prev and Next are nil by default
}

func (dll *DoublyLinkedList) insertTail(val int) {
	node := newNode(val) // create the new node to append

	if dll.Tail == nil { // list is empty — new node becomes both head and tail
		dll.Head = node
		dll.Tail = node
	} else {
		node.Prev = dll.Tail    // new node's prev points to current tail
		dll.Tail.Next = node    // current tail's next points to new node
		dll.Tail = node         // advance the tail pointer to the new node
		dll.Size = dll.Size + 1 // increment the list size
	}
}

func (dll *DoublyLinkedList) insertAtHead(val int) {
	node := newNode(val) // create the new node to prepend

	if dll.Head == nil && dll.Tail == nil { // list is empty — new node becomes both head and tail
		dll.Head = node
		dll.Tail = node
	} else {
		node.Next = dll.Head    // new node's next points to current head
		dll.Head.Prev = node    // current head's prev points to new node
		dll.Head = node         // move head pointer back to the new node
		dll.Size = dll.Size + 1 // increment the list size
	}
}

func (dll *DoublyLinkedList) removeFromHead() {

	switch {
	case dll.Head == nil: // list is already empty — nothing to remove
		fmt.Println("List is empty")

	case dll.Head == dll.Tail: // only one node — removing it empties the list
		dll.Head = nil
		dll.Tail = nil
		dll.Size--

	default:
		pointer := dll.Head      // save reference to the node being removed
		dll.Head = dll.Head.Next // advance head to the next node
		dll.Head.Prev = nil      // new head has no predecessor
		pointer.Next = nil       // detach the removed node from the list
		dll.Size--               // decrement the list size
	}

}

func (dll *DoublyLinkedList) removeFromTail() {

	switch {
	case dll.Tail == nil: // list is already empty — nothing to remove
		fmt.Println("List is Empty")

	case dll.Tail == dll.Head: // only one node — removing it empties the list
		dll.Head = nil
		dll.Tail = nil
		dll.Size--

	default:
		pointer := dll.Tail      // save reference to the node being removed
		dll.Tail = dll.Tail.Prev // move tail back to the previous node
		dll.Tail.Next = nil      // new tail has no successor
		pointer.Prev = nil       // detach the removed node from the list
		dll.Size--               // decrement the list size
	}
}

// InsertBetweenTwoNodes inserts a new node with val immediately after previousNode.
// Example: [1 <-> 2 <-> 4], InsertBetweenTwoNodes(3, node2) → [1 <-> 2 <-> 3 <-> 4]
func (dll *DoublyLinkedList) InsertBetweenTwoNodes(val int, previousNode *Node) {

	if previousNode == nil { // no anchor node provided — nothing to do
		return
	}

	nextNode := previousNode.Next

	insertedNode := newNode(val)

	previousNode.Next = insertedNode

	insertedNode.Prev = previousNode
	insertedNode.Next = nextNode

	if nextNode != nil {
		nextNode.Prev = insertedNode
	}
	dll.Size++
}

func (dll *DoublyLinkedList) RemovBetweenTwoNodes(val int) {
	if dll.Size == 0 {
		fmt.Println("Linked list is empty")
		return
	}

	pointer := dll.Head
	for pointer != nil {
		if pointer.Val == val {
			if pointer.Prev == nil || pointer.Next == nil {
				fmt.Println("Node is at head or tail, use removeFromHead/removeFromTail")
				return
			}
			pointer.Prev.Next = pointer.Next
			pointer.Next.Prev = pointer.Prev
			pointer.Prev = nil
			pointer.Next = nil
			dll.Size--
			return
		}
		pointer = pointer.Next
	}

	fmt.Println("Value not found")
}

func (dll *DoublyLinkedList) print() {
	cur := dll.Head  // start traversal from the head
	for cur != nil { // walk forward until we fall off the end
		if cur.Next != nil {
			fmt.Printf("%d <-> ", cur.Val) // not the last node — print with arrow
		} else {
			fmt.Printf("%d", cur.Val) // last node — no trailing arrow
		}
		cur = cur.Next // advance to the next node
	}
	fmt.Println() // newline after the full list
}

func main() {
	dll := &DoublyLinkedList{} // create an empty doubly linked list
	dll.insertTail(1)          // list: 1
	dll.insertTail(2)          // list: 1 <-> 2
	dll.insertTail(4)          // list: 1 <-> 2 <-> 4
	dll.insertTail(5)          // list: 1 <-> 2 <-> 4 <-> 5

	fmt.Print("Before: ")
	dll.print() // 1 <-> 2 <-> 4 <-> 5

	// insert 3 after node(2) so it lands between node(2) and node(4)
	dll.InsertBetweenTwoNodes(3, dll.Head.Next) // Head.Next is node(2)

	fmt.Print("After:  ")
	dll.print() // 1 <-> 2 <-> 3 <-> 4 <-> 5
}
