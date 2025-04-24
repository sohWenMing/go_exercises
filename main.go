package main

import (
	"fmt"

	"github.com/sohWenMing/go_exercises/linked_list"
)

func main() {
	list := []int{1, 2, 3}
	ll := linked_list.CreateLinkedList()
	for _, idx := range list {
		node := linked_list.CreateNode(idx)
		ll.InsertAtEnd(node)
	}
	node, isFound, err := ll.GetAt(2)
	fmt.Printf("\nnode: %v\n", node)
	fmt.Printf("\nisFound: %v\n", isFound)
	fmt.Printf("\nerr: %v\n", err)
}
