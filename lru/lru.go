package lru

import (
	"lru/linkedlist"
	"lru/node"
)

type Cacher[T comparable] interface {
	Get(key T) (value T, exists bool)

	spaceAvailable() bool
}

type LRU[T comparable] struct {
	dll     *linkedlist.LinkedList[T]
	maxSize int
	data    map[T]*node.Node[T]
}

type Options struct {
	MaxEntries int
}

func New[T comparable](options Options) *LRU[T] {
	if options.MaxEntries < 0 {
		panic("max entries must be non-negative")
	}

	return &LRU[T]{
		dll:     linkedlist.New[T](),
		maxSize: options.MaxEntries,
		data:    make(map[T]*node.Node[T]),
	}
}

var _ Cacher[any] = (*LRU[any])(nil)

func (l *LRU[T]) Get(value T) (T, bool) {
	// Search for value in map
	n, found := l.data[value]

	if !found {
		// Create a new node
		n = node.New(value)
	}

	if !l.spaceAvailable() {
		if found {
			// Remove the node from the list
			l.dll.Delete(n)
		} else {
			// Remove the least recently used data
			deletedNode := l.dll.Pop()
			if deletedNode != nil {
				delete(l.data, deletedNode.Val)
			}
		}
	} else if found {
		// Remove the node from the list
		l.dll.Delete(n)
	}

	// Insert at the end
	l.dll.Insert(n)

	// Add the node to map
	l.data[value] = n

	return n.Val, found
}

func (l *LRU[T]) spaceAvailable() bool {
	return l.maxSize > l.dll.Size() && l.maxSize > len(l.data) && l.dll.Size() == len(l.data)
}
