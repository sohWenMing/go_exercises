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

func TestDeleteFromBeginning(t *testing.T) {
	// Setup: Create a linked list 1 -> 2 -> 3
	node3 := &Node{data: 3}
	node2 := &Node{data: 2, next: node3}
	node1 := &Node{data: 1, next: node2}
	list := &LinkedList{head: node1, length: 3}

	// Act: Delete from beginning
	deleted := list.DeleteFromBeginning()

	// Assert
	if !deleted {
		t.Errorf("Expected deletion to return true, got false")
	}
	if list.head != node2 {
		t.Errorf("Expected head to be node2, got %+v", list.head)
	}
	if list.length != 2 {
		t.Errorf("Expected length to be 2 after deletion, got %d", list.length)
	}

	// Act: Delete remaining nodes
	list.DeleteFromBeginning() // deletes node2
	list.DeleteFromBeginning() // deletes node3

	// List should now be empty
	if list.head != nil {
		t.Errorf("Expected head to be nil after deleting all nodes")
	}
	if list.length != 0 {
		t.Errorf("Expected length to be 0 after all deletions, got %d", list.length)
	}

	// Try deleting from empty list
	deleted = list.DeleteFromBeginning()
	if deleted {
		t.Errorf("Expected deletion to return false on empty list, got true")
	}
}

func TestDeleteFromEnd(t *testing.T) {
	// --- Case 1: Deleting from a multi-node list ---
	node3 := &Node{data: 3}
	node2 := &Node{data: 2, next: node3}
	node1 := &Node{data: 1, next: node2}
	list := &LinkedList{head: node1, length: 3}

	deleted := list.DeleteFromEnd()
	if !deleted {
		t.Errorf("Expected deletion to return true, got false")
	}
	if list.length != 2 {
		t.Errorf("Expected length to be 2 after deletion, got %d", list.length)
	}
	if list.head.next.next != nil {
		t.Errorf("Expected last node to be nil after deletion, got %+v", list.head.next.next)
	}

	// --- Case 2: Deleting from a single-node list ---
	singleNode := &Node{data: 10}
	singleList := &LinkedList{head: singleNode, length: 1}

	deleted = singleList.DeleteFromEnd()
	if !deleted {
		t.Errorf("Expected deletion to return true for single-node list, got false")
	}
	if singleList.head != nil {
		t.Errorf("Expected head to be nil after deleting the only node, got %+v", singleList.head)
	}
	if singleList.length != 0 {
		t.Errorf("Expected length to be 0 after deletion, got %d", singleList.length)
	}

	// --- Case 3: Deleting from an empty list ---
	emptyList := &LinkedList{}
	deleted = emptyList.DeleteFromEnd()
	if deleted {
		t.Errorf("Expected deletion to return false on empty list, got true")
	}
	if emptyList.head != nil {
		t.Errorf("Expected head to still be nil on empty list")
	}
	if emptyList.length != 0 {
		t.Errorf("Expected length to be 0 on empty list, got %d", emptyList.length)
	}
}
func TestDeleteAtPosition(t *testing.T) {
	// Helper to build linked list: 1 -> 2 -> 3
	makeList := func() *LinkedList {
		n3 := &Node{data: 3}
		n2 := &Node{data: 2, next: n3}
		n1 := &Node{data: 1, next: n2}
		return &LinkedList{head: n1, length: 3}
	}

	t.Run("delete at position 0 (head)", func(t *testing.T) {
		list := makeList()
		deleted := list.DeleteAtPosition(0)
		if !deleted {
			t.Errorf("Expected deletion at position 0 to succeed")
		}
		if list.head.data != 2 {
			t.Errorf("Expected head to now be 2, got %d", list.head.data)
		}
		if list.length != 2 {
			t.Errorf("Expected length to be 2, got %d", list.length)
		}
	})

	t.Run("delete at middle position", func(t *testing.T) {
		list := makeList()
		deleted := list.DeleteAtPosition(1)
		if !deleted {
			t.Errorf("Expected deletion at position 1 to succeed")
		}
		if list.head.next.data != 3 {
			t.Errorf("Expected node at position 1 to be 3, got %d", list.head.next.data)
		}
		if list.length != 2 {
			t.Errorf("Expected length to be 2, got %d", list.length)
		}
	})

	t.Run("delete at last position", func(t *testing.T) {
		list := makeList()
		deleted := list.DeleteAtPosition(2)
		if !deleted {
			t.Errorf("Expected deletion at position 2 to succeed")
		}
		if list.head.next.next != nil {
			t.Errorf("Expected last node to be nil after deletion")
		}
		if list.length != 2 {
			t.Errorf("Expected length to be 2, got %d", list.length)
		}
	})

	t.Run("delete at out-of-bounds position", func(t *testing.T) {
		list := makeList()
		deleted := list.DeleteAtPosition(5)
		if deleted {
			t.Errorf("Expected deletion at out-of-bounds position to fail")
		}
	})

	t.Run("delete from empty list", func(t *testing.T) {
		list := &LinkedList{}
		deleted := list.DeleteAtPosition(0)
		if deleted {
			t.Errorf("Expected deletion from empty list to fail")
		}
	})

	t.Run("delete at negative position", func(t *testing.T) {
		list := makeList()
		deleted := list.DeleteAtPosition(-1)
		if deleted {
			t.Errorf("Expected deletion at negative position to fail")
		}
	})
}

func TestDeleteByValue(t *testing.T) {
	t.Run("delete from empty list", func(t *testing.T) {
		ll := CreateLinkedList()
		deleted := ll.DeleteByValue(10)
		if deleted {
			t.Error("expected false, got true when deleting from empty list")
		}
	})

	t.Run("delete head value in list", func(t *testing.T) {
		ll := CreateLinkedList()
		ll.InsertAtEnd(CreateNode(1))
		ll.InsertAtEnd(CreateNode(2))
		ll.InsertAtEnd(CreateNode(3))
		deleted := ll.DeleteByValue(1)
		if !deleted {
			t.Error("expected true, got false when deleting head")
		}
		if ll.GetHead().data != 2 {
			t.Errorf("expected head to be 2, got %d", ll.GetHead().data)
		}
	})

	t.Run("delete middle value", func(t *testing.T) {
		ll := CreateLinkedList()
		ll.InsertAtEnd(CreateNode(1))
		ll.InsertAtEnd(CreateNode(2))
		ll.InsertAtEnd(CreateNode(3))
		deleted := ll.DeleteByValue(2)
		if !deleted {
			t.Error("expected true, got false when deleting middle node")
		}
		if ll.GetHead().next.data != 3 {
			t.Errorf("expected second node to be 3, got %d", ll.GetHead().next.data)
		}
	})

	t.Run("delete last value", func(t *testing.T) {
		ll := CreateLinkedList()
		ll.InsertAtEnd(CreateNode(1))
		ll.InsertAtEnd(CreateNode(2))
		ll.InsertAtEnd(CreateNode(3))
		deleted := ll.DeleteByValue(3)
		if !deleted {
			t.Error("expected true, got false when deleting last node")
		}
		if ll.GetHead().next.next != nil {
			t.Errorf("expected tail to be nil, got %v", ll.GetHead().next.next)
		}
	})

	t.Run("delete non-existent value", func(t *testing.T) {
		ll := CreateLinkedList()
		ll.InsertAtEnd(CreateNode(1))
		ll.InsertAtEnd(CreateNode(2))
		deleted := ll.DeleteByValue(42)
		if deleted {
			t.Error("expected false, got true when value does not exist")
		}
	})

	t.Run("delete only element in list", func(t *testing.T) {
		ll := CreateLinkedList()
		ll.InsertAtEnd(CreateNode(99))
		deleted := ll.DeleteByValue(99)
		if !deleted {
			t.Error("expected true, got false when deleting only node")
		}
		if ll.GetHead() != nil {
			t.Errorf("expected head to be nil, got %v", ll.GetHead())
		}
	})

	t.Run("delete one of multiple same values", func(t *testing.T) {
		ll := CreateLinkedList()
		ll.InsertAtEnd(CreateNode(7))
		ll.InsertAtEnd(CreateNode(7))
		ll.InsertAtEnd(CreateNode(7))
		deleted := ll.DeleteByValue(7)
		if !deleted {
			t.Error("expected true, got false when deleting from list with duplicates")
		}
		if ll.GetHead().data != 7 || ll.GetHead().next.data != 7 || ll.GetHead().next.next != nil {
			t.Error("expected only first occurrence of value 7 to be deleted")
		}
	})
}
