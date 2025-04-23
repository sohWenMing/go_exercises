package linked_list

import "errors"

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
	return nil
}

func (l *LinkedList) InsertAtEnd(n *Node) error {
	if checkErr := checkNilNode(n); checkErr != nil {
		return checkErr
	}
	if isSet := l.setHeadIfHeadEmpty(n); isSet {
		return nil
	}
	curNode := l.GetHead()
	for curNode.next != nil {
		curNode = curNode.next
	}
	curNode.next = n
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
	return false
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
