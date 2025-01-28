package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) Add(i int) {
	newNode := &Node{data: i}
	if ll.head == nil {
		ll.head = newNode
	} else {
		curr := ll.head
		for curr.next != nil {
			curr = curr.next
		}
		curr.next = newNode
	}
}

func (ll *LinkedList) Display() {
	curr := ll.head
	for curr != nil {
		fmt.Printf("%d -> ", curr.data)
		curr = curr.next
	}
	fmt.Println("nil")
}

func (ll *LinkedList) Delete(i int) {

	if ll.head == nil {
		fmt.Println("Linked List is empty")
		return
	}

	if ll.head.data == i {
		ll.head = ll.head.next
		return
	}
	current := ll.head
	for current.next != nil && current.next.data != i {
		current = current.next
	}

	if current.next == nil {
		fmt.Println("Value not found in the list")
	} else {
		current.next = current.next.next
	}
}

func main() {
	ll := &LinkedList{}
	ll.Add(2)
	ll.Add(3)
	ll.Add(4)
	ll.Add(5)
	ll.Add(6)
	ll.Add(7)
	ll.Display()
	ll.Delete(4)
	ll.Display()

}
