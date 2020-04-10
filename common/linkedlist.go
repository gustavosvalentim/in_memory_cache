package common

import (
	"fmt"
)

// LinkedList : 
type LinkedList struct {
	Head	*Node
	Tail	*Node
	Count	int
}

// Node : 
type Node struct {
	Next	*Node
	Key		string
	Data	interface{}
}

// Append : 
func (ll *LinkedList) Append(key string, value interface{}) *LinkedList {
	emptyNode := &Node{
		Next: nil,
		Key: key,
		Data: value,
	}

	if (ll.Head.Data == nil) {
		ll.Head = emptyNode
	} else {
		ll.Tail.Next = emptyNode
	}

	ll.Tail = emptyNode
	ll.Count++

	return ll
}

// Get : 
func (ll *LinkedList) Get(key string) (*Node, bool) {
	node := ll.Head

	for {
		if (node.Key == key) {
			return node, true
		}

		if (node.Next == nil) {
			break
		}

		node = node.Next
	}

	return nil, false
}

// Scan : 
func (ll *LinkedList) Scan() []*Node {
	var items []*Node
	node := ll.Head

	for {
		if (node.Data != nil) {
			items = append(items, node)
		}

		if (node.Next == nil) {
			return items
		}

		node = node.Next
	}
}

// Pop : 
func (ll *LinkedList) Pop(key string) *LinkedList {
	emptyNode := &Node{
		Next: nil,
		Key: "",
		Data: nil,
	}
	prev := ll.Head
	node := ll.Head

	if (node.Key == key) {
		if (node.Next != nil) {
			ll.Head = node.Next
		} else {
			ll.Head = emptyNode
		}
		return ll
	}

	node = node.Next

	for {
		if (node.Key == key) {
			if (node.Next == nil) {
				ll.Tail = prev
			} else {
				prev.Next = node.Next
			}

			return ll
		}

		// Set previous to this node
		prev = node
		// Set node to the next node in the chain
		node = node.Next
	}
}

// Print : 
func (ll *LinkedList) Print() {
	node := ll.Head
	for {
		fmt.Println(node.Data)

		if node.Next == nil {
			return
		}

		node = node.Next
	}
}

// NewLinkedList : 
func NewLinkedList() *LinkedList {
	emptyNode := &Node{
		Next: nil,
		Key: "",
		Data: nil,
	}
	return &LinkedList{
		Head: emptyNode,
		Tail: emptyNode,
		Count: 0,
	}
}