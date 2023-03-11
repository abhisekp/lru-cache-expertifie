package linkedlist

import "lru/node"

type LinkedLister[T comparable] interface {
	// Head of the linked list
	Head() *node.Node[T]

	// Tail of the linked list
	Tail() *node.Node[T]

	// Insert at the end
	Insert(node *node.Node[T])

	// Search a value
	Search(val T) *node.Node[T]

	// Delete a node
	Delete(node *node.Node[T]) *node.Node[T]

	// Pop a node from start
	Pop() *node.Node[T]

	// Size of the list
	Size() int
}

type LinkedList[T comparable] struct {
	head *node.Node[T]
	tail *node.Node[T]
	size int
}

var _ LinkedLister[any] = (*LinkedList[any])(nil)

func New[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (l *LinkedList[T]) Head() *node.Node[T] {
	return l.head
}

func (l *LinkedList[T]) Tail() *node.Node[T] {
	return l.tail
}

func (l *LinkedList[T]) Insert(node *node.Node[T]) {
	if node == nil {
		return
	}

	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		l.tail.SetNext(node) // *->x
		node.SetPrev(l.tail) // *<->x
		l.tail = node
	}

	l.size++
}

func (l *LinkedList[T]) Delete(node *node.Node[T]) *node.Node[T] {
	if node == nil {
		return nil
	}

	prevNode := node.Prev()
	nextNode := node.Next()

	if prevNode != nil {
		prevNode.SetNext(nextNode)
	} else {
		l.head = nextNode
	}

	if nextNode != nil {
		nextNode.SetPrev(prevNode)
	} else {
		l.tail = prevNode
	}

	node.SetPrev(nil)
	node.SetNext(nil)

	l.size--

	return node
}

func (l *LinkedList[T]) Pop() *node.Node[T] {
	// Use the Delete method
	return l.Delete(l.head)
}

func (l *LinkedList[T]) Size() int {
	return l.size
}

func (l *LinkedList[T]) Search(val T) *node.Node[T] {
	var curr *node.Node[T]

	for curr = l.head; curr != nil && curr.Val != val; curr = curr.Next() {
	}

	return curr
}
