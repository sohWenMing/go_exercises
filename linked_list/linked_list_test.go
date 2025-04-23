package linked_list

import (
	"testing"
)

func TestCreateNode(t *testing.T) {
	n := CreateNode(10)
	if n == nil {
		t.Fatal("Expected new node, got nil")
	}
	if n.data != 10 {
		t.Errorf("Expected node data to be 10, got %d", n.data)
	}
	if n.next != nil {
		t.Error("Expected new node's next to be nil")
	}
}

func TestCreateLinkedList(t *testing.T) {
	list := CreateLinkedList()
	if list == nil {
		t.Fatal("Expected new linked list, got nil")
	}
	if list.head != nil {
		t.Error("Expected head to be nil")
	}
	if list.length != 0 {
		t.Errorf("Expected length to be 0, got %d", list.length)
	}
}

func TestSetAndGetHead(t *testing.T) {
	list := CreateLinkedList()
	node := CreateNode(5)
	list.SetHead(node)

	if list.GetHead() != node {
		t.Error("GetHead did not return the node set by SetHead")
	}
}

func TestInsertAtBeginning_EmptyList(t *testing.T) {
	list := CreateLinkedList()
	node := CreateNode(7)

	err := list.InsertAtBeginning(node)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if list.GetHead() != node {
		t.Error("Expected inserted node to be head of the list")
	}
}

func TestInsertAtBeginning_NonEmptyList(t *testing.T) {
	list := CreateLinkedList()
	node1 := CreateNode(1)
	node2 := CreateNode(2)

	_ = list.InsertAtBeginning(node1)
	_ = list.InsertAtBeginning(node2)

	if list.GetHead() != node2 {
		t.Error("Expected head to be the most recently inserted node")
	}
	if list.GetHead().next != node1 {
		t.Error("Expected the new head to point to the old head")
	}
}

func TestInsertAtBeginning_NilNode(t *testing.T) {
	list := CreateLinkedList()
	err := list.InsertAtBeginning(nil)
	if err == nil {
		t.Error("Expected error when inserting nil node, got none")
	}
}

func TestInsertAtEnd_EmptyList(t *testing.T) {
	list := CreateLinkedList()
	node := CreateNode(10)

	err := list.InsertAtEnd(node)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if list.GetHead() != node {
		t.Error("Expected inserted node to be head of the list")
	}
	if node.next != nil {
		t.Error("Expected last node's next to be nil")
	}
}

func TestInsertAtEnd_NonEmptyList(t *testing.T) {
	list := CreateLinkedList()
	node1 := CreateNode(1)
	node2 := CreateNode(2)

	_ = list.InsertAtEnd(node1)
	_ = list.InsertAtEnd(node2)

	head := list.GetHead()
	if head != node1 {
		t.Error("Expected first inserted node to remain head")
	}
	if head.next != node2 {
		t.Error("Expected second node to be appended to the end")
	}
	if node2.next != nil {
		t.Error("Expected last node's next to be nil")
	}
}

func TestInsertAtEnd_NilNode(t *testing.T) {
	list := CreateLinkedList()
	err := list.InsertAtEnd(nil)

	if err == nil {
		t.Error("Expected error when inserting nil node, got none")
	}
}

func TestInsertAtPosition_EmptyListAtZero(t *testing.T) {
	list := CreateLinkedList()
	inserted := list.InsertAtPosition(0, 42)

	if !inserted {
		t.Error("Expected insertion at position 0 to succeed")
	}
	if list.GetHead() == nil || list.GetHead().data != 42 {
		t.Error("Expected head to be the inserted node with data 42")
	}
}

func TestInsertAtPosition_NonEmptyListAtZero(t *testing.T) {
	list := CreateLinkedList()
	_ = list.InsertAtEnd(CreateNode(99))

	inserted := list.InsertAtPosition(0, 77)
	if !inserted {
		t.Error("Expected insertion at position 0 to succeed")
	}
	if list.GetHead().data != 77 {
		t.Errorf("Expected new head to be 77, got %d", list.GetHead().data)
	}
	if list.GetHead().next == nil || list.GetHead().next.data != 99 {
		t.Error("Expected original node to come after newly inserted node")
	}
}

func TestInsertAtPosition_MiddleOfList(t *testing.T) {
	list := CreateLinkedList()
	_ = list.InsertAtEnd(CreateNode(1))
	_ = list.InsertAtEnd(CreateNode(3))

	inserted := list.InsertAtPosition(1, 2)
	if !inserted {
		t.Error("Expected insertion at position 1 to succeed")
	}

	first := list.GetHead()
	second := first.next
	third := second.next

	if first.data != 1 || second.data != 2 || third.data != 3 {
		t.Errorf("Expected list order 1 -> 2 -> 3, got %d -> %d -> %d", first.data, second.data, third.data)
	}
}

func TestInsertAtPosition_OutOfBounds(t *testing.T) {
	list := CreateLinkedList()
	_ = list.InsertAtEnd(CreateNode(1))
	_ = list.InsertAtEnd(CreateNode(2))

	inserted := list.InsertAtPosition(5, 99)
	if inserted {
		t.Error("Expected insertion to fail due to out-of-bounds position")
	}
}
