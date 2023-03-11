package node

type Noder[T any] interface {
	Prev() *Node[T]
	Next() *Node[T]
	SetPrev(*Node[T])
	SetNext(*Node[T])
}

type Node[T any] struct {
	Val  T
	prev *Node[T]
	next *Node[T]
}

var _ Noder[any] = (*Node[any])(nil)

func New[T any](val T) *Node[T] {
	return &Node[T]{Val: val}
}

func (n *Node[T]) Prev() *Node[T] {
	return n.prev
}

func (n *Node[T]) Next() *Node[T] {
	return n.next
}

func (n *Node[T]) SetPrev(prev *Node[T]) {
	n.prev = prev
}

func (n *Node[T]) SetNext(next *Node[T]) {
	n.next = next
}
