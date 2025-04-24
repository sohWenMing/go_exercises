package linked_list

import (
	"errors"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func (l *LinkedList) GetHead() *Node {
	return l.head
}

func (l *LinkedList) SetHead(newHead *Node) {
	l.head = newHead
}

func CreateNode(data int) *Node {
	return &Node{
		data: data,
		next: nil,
	}
}

func CreateLinkedList() *LinkedList {
	return &LinkedList{
		head:   nil,
		length: 0,
	}
}

func (l *LinkedList) InsertAtBeginning(n *Node) error {
	if checkErr := checkNilNode(n); checkErr != nil {
		return checkErr
	}
	if isSet := l.setHeadIfHeadEmpty(n); isSet {
		return nil
	}
	n.next = l.head
	l.SetHead(n)
	l.incrementLength()
	return nil
}

func (l *LinkedList) InsertAtEnd(n *Node) error {
	if checkErr := checkNilNode(n); checkErr != nil {
		return checkErr
	}
	if isSet := l.setHeadIfHeadEmpty(n); isSet {
		l.incrementLength()
		return nil
	}
	curNode := l.GetHead()
	for curNode.next != nil {
		curNode = curNode.next
	}
	curNode.next = n
	l.incrementLength()
	return nil
}

func (l *LinkedList) InsertAtPosition(pos int, value int) (isInserted bool) {
	// first thing should be to find whether the position even exists
	newNode := CreateNode(value)
	curNode := l.GetHead()
	// two scenarios: if head is filled, if head is empty
	curPos := 0
	if curPos == pos {
		isInserted = true
		if curNode != nil {
			newNode.next = curNode
		}
		l.SetHead(newNode)
		return
	}
	var prevNode *Node
	for curNode.next != nil {
		prevNode = curNode
		curNode = curNode.next
		curPos++
		switch curPos == pos {
		case false:
			continue
		case true:
			prevNode.next = newNode
			newNode.next = curNode
			isInserted = true
			return
		}
	}
	l.incrementLength()
	return false
}

func (l *LinkedList) DeleteFromBeginning() (isDeleted bool) {
	head := l.GetHead()
	if head == nil {
		isDeleted = false
		return
	}
	isDeleted = true
	l.SetHead(head.next)
	l.decrementLength()
	return
}

func (l *LinkedList) DeleteFromEnd() (isDeleted bool) {
	head := l.GetHead()
	if head == nil {
		isDeleted = false
		return
	}
	isDeleted = true
	l.decrementLength()
	if head.next == nil {
		l.SetHead(nil)
		return
	}
	prev := head
	curr := head.next
	for curr.next != nil {
		prev = curr
		curr = curr.next
	}
	prev.next = nil
	return
}

func (l *LinkedList) DeleteAtPosition(pos int) (isDeleted bool) {
	isDeleted = false
	if pos < 0 {
		return
	}
	// if the pos entered is less than zero, just exit
	if pos > l.length-1 {
		return
	}
	// if the position is out of bounds, return

	curr := l.GetHead()
	if curr == nil {
		return
	}
	// if list is empty, just return

	currPos := 0
	l.decrementLength()
	// at this point it's safe to decrement, because we know that pos is within the list by position
	// and that only max 1 will be decremented
	isDeleted = true
	var prev *Node
	for curr != nil {
		switch currPos == pos {
		case true:
			if prev == nil {
				l.SetHead(curr.next)
				return
				//if prev == nil, then we're at head, so just set the new head to next which can be next or nil
			} else {
				prev.next = curr.next
				return
			}
			//node is deleted, setting prev.next to point to curr.next. exit
		case false:
			prev = curr
			curr = curr.next
			currPos++
		}
	}
	return
}

func (l *LinkedList) DeleteByValue(val int) (isDeleted bool) {
	isDeleted = false
	curr := l.GetHead()
	if curr == nil {
		return
	}
	//empty list, return
	var prev *Node
	for curr != nil {
		switch val == curr.data {
		case true:
			isDeleted = true
			l.decrementLength()
			//either way, safe to decrement and return true

			if curr == l.GetHead() {
				l.SetHead(curr.next)
				return
			}
			//if we're still at head, then just reset the head and decrement list length
			prev.next = curr.next
			return

		case false:
			prev = curr
			curr = curr.next
		}
	}
	return
}

func (l *LinkedList) SearchByValue(val int) (isFound bool, position int) {
	isFound = false
	position = -1
	// init negative return values

	curr := l.GetHead()
	for curr != nil {
		position++
		if curr.data == val {
			isFound = true
			return
		}
		curr = curr.next
	}
	position = -1
	return
}

func (l *LinkedList) GetAt(pos int) (n *Node, isFound bool, err error) {
	isFound = false
	n = nil
	if pos < 0 || pos > l.length-1 {
		return n, isFound, errors.New("position out of bounds")
	}
	curr := l.GetHead()
	if curr == nil {
		return n, isFound, errors.New("list cannot be empty")
	}
	currPos := -1
	for curr != nil {
		currPos++
		if currPos == pos {
			return curr, true, nil
		}
		curr = curr.next
	}
	//check through all nodes, if found, return
	return n, isFound, nil

}
func (l *LinkedList) decrementLength() {
	l.length -= 1
}

func (l *LinkedList) incrementLength() {
	l.length++
}

func checkNilNode(n *Node) error {
	if n == nil {
		return errors.New("nil pointer passed as node input")
	}
	return nil
}

func (l *LinkedList) setHeadIfHeadEmpty(n *Node) (isSet bool) {
	head := l.GetHead()
	if head == nil {
		l.SetHead(n)
		return true
	}
	return false
}
