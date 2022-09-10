package main

import "fmt"

func main() {

	head := createlist(5)

	head.insert(10)

	head.insert(25)

	head.insert(25)

	head.insertbyindex(50, 3)

	head.displayall()

}

type node struct {
	data     int
	next     *node
	previous *node
}

func createlist(i int) (head node) {
	node := &node{
		data:     i,
		next:     nil,
		previous: nil,
	}
	return *node
}

func (n *node) insert(i int) {

	pre := n

	newnode := &node{
		data:     i,
		next:     nil,
		previous: pre,
	}

	for {

		if n.next != nil {
			n = n.next
			pre = n
		} else {
			newnode.previous = pre
			n.next = newnode
			break
		}
	}

}

func (n *node) insertbyindex(i int, index int) {

	newnode := &node{
		data:     i,
		next:     nil,
		previous: nil,
	}

	for i := 0; i < index; i++ {

		n = n.next
	}

	if n.next != nil {

		newnode.next = n.next
		newnode.previous = n
		n.next.previous = newnode
		n.next = newnode
	} else {

		newnode.previous = n
		n.next = newnode
	}

}

func (n *node) displayall() {
	var current *node = n
	for {
		fmt.Println(current)
		if current.next != nil {
			current = current.next

		} else {
			break
		}
	}
}
