package linkedlist_test

import (
	"testing"

	"lru/linkedlist"
	"lru/node"
)

func TestLinkedList(t *testing.T) {
	t.Run("Insert", func(t *testing.T) {
		ll := linkedlist.New[int]()
		ll.Insert(node.New(1))

		if ll.Head().Val != 1 {
			t.Errorf("expected %v, got %v", 1, ll.Head().Val)
		}
		if ll.Size() != 1 {
			t.Errorf("expected %v, got %v", 1, ll.Size())
		}

		ll.Insert(node.New(2))
		if ll.Head().Next().Val != 2 {
			t.Errorf("expected %v, got %v", 2, ll.Head().Next().Val)
		}
		if ll.Head().Next().Prev() != ll.Head() {
			t.Errorf("expected %v, got %v", ll.Head(), ll.Head().Next().Prev())
		}
		if ll.Size() != 2 {
			t.Errorf("expected %v, got %v", 2, ll.Size())
		}
	})

	t.Run("Search", func(t *testing.T) {
		t.Run("Search for a non-existent value", func(t *testing.T) {
			ll := linkedlist.New[int]()
			ll.Insert(node.New(1))
			ll.Insert(node.New(2))
			ll.Insert(node.New(3))
			ll.Insert(node.New(4))

			n := ll.Search(5)
			if n != nil {
				t.Errorf("expected %v, got %v", nil, n)
			}
		})

		t.Run("Search for an existing value", func(t *testing.T) {
			ll := linkedlist.New[int]()
			ll.Insert(node.New(1))
			ll.Insert(node.New(2))
			ll.Insert(node.New(3))
			ll.Insert(node.New(4))

			n := ll.Search(3)
			if n.Val != 3 {
				t.Errorf("expected %v, got %v", 3, n.Val)
			}
			if n.Next().Val != 4 {
				t.Errorf("expected %v, got %v", 4, n.Next().Val)
			}
			if n.Prev().Val != 2 {
				t.Errorf("expected %v, got %v", 2, n.Prev().Val)
			}
		})

		t.Run("Search for first value", func(t *testing.T) {
			ll := linkedlist.New[int]()
			ll.Insert(node.New(1))
			ll.Insert(node.New(2))
			ll.Insert(node.New(3))
			ll.Insert(node.New(4))

			n := ll.Search(1)
			if n.Val != 1 {
				t.Errorf("expected %v, got %v", 1, n.Val)
			}
			if n.Next().Val != 2 {
				t.Errorf("expected %v, got %v", 2, n.Next().Val)
			}
			if n.Prev() != nil {
				t.Errorf("expected %v, got %v", nil, n.Prev())
			}
		})

		t.Run("Search for first value with one node", func(t *testing.T) {
			ll := linkedlist.New[int]()
			ll.Insert(node.New(1))

			n := ll.Search(1)
			if n.Val != 1 {
				t.Errorf("expected %v, got %v", 1, n.Val)
			}
			if n.Next() != nil {
				t.Errorf("expected %v, got %v", nil, n.Next())
			}
			if n.Prev() != nil {
				t.Errorf("expected %v, got %v", nil, n.Prev())
			}
		})

		t.Run("Search for last value", func(t *testing.T) {
			ll := linkedlist.New[int]()
			ll.Insert(node.New(1))
			ll.Insert(node.New(2))
			ll.Insert(node.New(3))
			ll.Insert(node.New(4))

			n := ll.Search(4)
			if n.Val != 4 {
				t.Errorf("expected %v, got %v", 4, n.Val)
			}
			if n.Next() != nil {
				t.Errorf("expected %v, got %v", nil, n.Next())
			}
			if n.Prev().Val != 3 {
				t.Errorf("expected %v, got %v", 3, n.Prev().Val)
			}
		})

		t.Run("Search for last value with one node", func(t *testing.T) {
			ll := linkedlist.New[int]()
			ll.Insert(node.New(1))

			n := ll.Search(1)
			if n.Val != 1 {
				t.Errorf("expected %v, got %v", 1, n.Val)
			}
			if n.Next() != nil {
				t.Errorf("expected %v, got %v", nil, n.Next())
			}
			if n.Prev() != nil {
				t.Errorf("expected %v, got %v", nil, n.Prev())
			}
		})

		t.Run("Search for first value with two nodes", func(t *testing.T) {
			ll := linkedlist.New[int]()
			ll.Insert(node.New(1))
			ll.Insert(node.New(2))

			n := ll.Search(1)
			if n.Val != 1 {
				t.Errorf("expected %v, got %v", 1, n.Val)
			}
			if n.Next().Val != 2 {
				t.Errorf("expected %v, got %v", 2, n.Next().Val)
			}
			if n.Prev() != nil {
				t.Errorf("expected %v, got %v", nil, n.Prev())
			}
		})

		t.Run("Search for last value with two nodes", func(t *testing.T) {
			ll := linkedlist.New[int]()
			ll.Insert(node.New(1))
			ll.Insert(node.New(2))

			n := ll.Search(2)
			if n.Val != 2 {
				t.Errorf("expected %v, got %v", 2, n.Val)
			}
			if n.Next() != nil {
				t.Errorf("expected %v, got %v", nil, n.Next())
			}
			if n.Prev().Val != 1 {
				t.Errorf("expected %v, got %v", 1, n.Prev().Val)
			}
		})

		t.Run("Search for a value with deleted nodes", func(t *testing.T) {
			t.Skip("not implemented")
		})
	})

	t.Run("Delete", func(t *testing.T) {
		t.Run("Delete the only node in the list", func(t *testing.T) {
			ll := linkedlist.New[int]()
			n := node.New(1)
			ll.Insert(n)

			out := ll.Delete(n)
			if out != n {
				t.Errorf("expected %v, got %v", n, out)
			}
			if ll.Head() != nil {
				t.Errorf("expected %v, got %v", nil, ll.Head())
			}
			if ll.Tail() != nil {
				t.Errorf("expected %v, got %v", nil, ll.Tail())
			}
			if n.Prev() != nil {
				t.Errorf("expected %v, got %v", nil, n.Prev())
			}
			if n.Next() != nil {
				t.Errorf("expected %v, got %v", nil, n.Next())
			}
		})

		t.Run("Delete the first node", func(t *testing.T) {
			ll := linkedlist.New[int]()
			n1 := node.New(1)
			n2 := node.New(2)
			ll.Insert(n1)
			ll.Insert(n2)

			out := ll.Delete(n1)
			if out != n1 {
				t.Errorf("expected %v, got %v", n1, out)
			}
			if ll.Head() != n2 {
				t.Errorf("expected %v, got %v", n2, ll.Head())
			}
			if ll.Tail() != n2 {
				t.Errorf("expected %v, got %v", n2, ll.Tail())
			}
			if n1.Prev() != nil {
				t.Errorf("expected %v, got %v", nil, n1.Prev())
			}
			if n1.Next() != nil {
				t.Errorf("expected %v, got %v", nil, n1.Next())
			}
			if n2.Prev() != nil {
				t.Errorf("expected %v, got %v", nil, n2.Prev())
			}
			if n2.Next() != nil {
				t.Errorf("expected %v, got %v", nil, n2.Next())
			}
		})

		t.Run("Delete the last node", func(t *testing.T) {
			ll := linkedlist.New[int]()
			n1 := node.New(1)
			n2 := node.New(2)
			ll.Insert(n1)
			ll.Insert(n2)

			out := ll.Delete(n2)
			if out != n2 {
				t.Errorf("expected %v, got %v", n2, out)
			}
			if ll.Head() != n1 {
				t.Errorf("expected %v, got %v", n1, ll.Head())
			}
			if ll.Tail() != n1 {
				t.Errorf("expected %v, got %v", n1, ll.Tail())
			}
			if n1.Prev() != nil {
				t.Errorf("expected %v, got %v", nil, n1.Prev())
			}
			if n1.Next() != nil {
				t.Errorf("expected %v, got %v", nil, n1.Next())
			}
			if n2.Prev() != nil {
				t.Errorf("expected %v, got %v", nil, n2.Prev())
			}
			if n2.Next() != nil {
				t.Errorf("expected %v, got %v", nil, n2.Next())
			}
		})

		t.Run("Delete a node in middle of list", func(t *testing.T) {
			ll := linkedlist.New[int]()
			n1 := node.New(1)
			n2 := node.New(2)
			n3 := node.New(3)
			ll.Insert(n1)
			ll.Insert(n2)
			ll.Insert(n3)

			out := ll.Delete(n2)
			if out != n2 {
				t.Errorf("expected %v, got %v", n2, out)
			}
			if ll.Head() != n1 {
				t.Errorf("expected %v, got %v", n1, ll.Head())
			}
			if ll.Tail() != n3 {
				t.Errorf("expected %v, got %v", n3, ll.Tail())
			}
			if n1.Prev() != nil {
				t.Errorf("expected %v, got %v", nil, n1.Prev())
			}
			if n1.Next() != n3 {
				t.Errorf("expected %v, got %v", n3, n1.Next())
			}
			if n2.Prev() != nil {
				t.Errorf("expected %v, got %v", nil, n2.Prev())
			}
			if n2.Next() != nil {
				t.Errorf("expected %v, got %v", nil, n2.Next())
			}
			if n3.Prev() != n1 {
				t.Errorf("expected %v, got %v", n1, n3.Prev())
			}
			if n3.Next() != nil {
				t.Errorf("expected %v, got %v", nil, n3.Next())
			}
		})

		t.Run("Delete all nodes in the list", func(t *testing.T) {
			ll := linkedlist.New[int]()
			n1 := node.New(1)
			n2 := node.New(2)
			n3 := node.New(3)
			ll.Insert(n1)
			ll.Insert(n2)
			ll.Insert(n3)

			out := ll.Delete(n1)
			if out != n1 {
				t.Errorf("expected %v, got %v", n1, out)
			}
			if ll.Head() != n2 {
				t.Errorf("expected %v, got %v", n2, ll.Head())
			}
			if ll.Tail() != n3 {
				t.Errorf("expected %v, got %v", n3, ll.Tail())
			}
			if n1.Prev() != nil {
				t.Errorf("expected %v, got %v", nil, n1.Prev())
			}
			if n1.Next() != nil {
				t.Errorf("expected %v, got %v", nil, n1.Next())
			}
			if n2.Prev() != nil {
				t.Errorf("expected %v, got %v", nil, n2.Prev())
			}
			if n2.Next() != n3 {
				t.Errorf("expected %v, got %v", n3, n2.Next())
			}
			if n3.Prev() != n2 {
				t.Errorf("expected %v, got %v", n2, n3.Prev())
			}
			if n3.Next() != nil {
				t.Errorf("expected %v, got %v", nil, n3.Next())
			}

			out = ll.Delete(n2)
			if out != n2 {
				t.Errorf("expected %v, got %v", n2, out)
			}
			if ll.Head() != n3 {
				t.Errorf("expected %v, got %v", n3, ll.Head())
			}
			if ll.Tail() != n3 {
				t.Errorf("expected %v, got %v", n3, ll.Tail())
			}
			if n2.Prev() != nil {
				t.Errorf("expected %v, got %v", nil, n2.Prev())
			}
			if n2.Next() != nil {
				t.Errorf("expected %v, got %v", nil, n2.Next())
			}

			out = ll.Delete(n3)
			if out != n3 {
				t.Errorf("expected %v, got %v", n3, out)
			}
			if ll.Head() != nil {
				t.Errorf("expected %v, got %v", nil, ll.Head())
			}
			if ll.Tail() != nil {
				t.Errorf("expected %v, got %v", nil, ll.Tail())
			}
			if n3.Prev() != nil {
				t.Errorf("expected %v, got %v", nil, n3.Prev())
			}
			if n3.Next() != nil {
				t.Errorf("expected %v, got %v", nil, n3.Next())
			}
		})

		t.Run("Delete a node that is not in the list", func(t *testing.T) {
			t.Skip("not applicable for this implementation")

			ll := linkedlist.New[int]()
			n1 := node.New(1)
			n2 := node.New(2)
			n3 := node.New(3)
			ll.Insert(n1)
			ll.Insert(n2)

			out := ll.Delete(n3)
			if out != n3 {
				t.Errorf("expected %v, got %v", nil, out)
			}
			if ll.Head() != n1 {
				t.Errorf("expected %v, got %v", n1, ll.Head())
			}
			if ll.Tail() != n2 {
				t.Errorf("expected %v, got %v", n2, ll.Tail())
			}
			if n1.Prev() != nil {
				t.Errorf("expected %v, got %v", nil, n1.Prev())
			}
			if n1.Next() != n2 {
				t.Errorf("expected %v, got %v", n2, n1.Next())
			}
			if n2.Prev() != n1 {
				t.Errorf("expected %v, got %v", n1, n2.Prev())
			}
			if n2.Next() != nil {
				t.Errorf("expected %v, got %v", nil, n2.Next())
			}
		})
	})

	t.Run("Pop", func(t *testing.T) {
		t.Run("Pop the only node in the list", func(t *testing.T) {
			ll := linkedlist.New[int]()
			n := node.New(1)
			ll.Insert(n)

			out := ll.Pop()
			if out != n {
				t.Errorf("expected %v, got %v", n, out)
			}
			if ll.Head() != nil {
				t.Errorf("expected %v, got %v", nil, ll.Head())
			}
			if ll.Tail() != nil {
				t.Errorf("expected %v, got %v", nil, ll.Tail())
			}
			if n.Prev() != nil {
				t.Errorf("expected %v, got %v", nil, n.Prev())
			}
			if n.Next() != nil {
				t.Errorf("expected %v, got %v", nil, n.Next())
			}
		})

		t.Run("Pop the first node", func(t *testing.T) {
			ll := linkedlist.New[int]()
			n1 := node.New(1)
			n2 := node.New(2)
			ll.Insert(n1)
			ll.Insert(n2)

			out := ll.Pop()
			if out != n1 {
				t.Errorf("expected %v, got %v", n1, out)
			}
			if ll.Head() != n2 {
				t.Errorf("expected %v, got %v", n2, ll.Head())
			}
			if ll.Tail() != n2 {
				t.Errorf("expected %v, got %v", n2, ll.Tail())
			}
			if n1.Prev() != nil {
				t.Errorf("expected %v, got %v", nil, n1.Prev())
			}
			if n1.Next() != nil {
				t.Errorf("expected %v, got %v", nil, n1.Next())
			}
			if n2.Prev() != nil {
				t.Errorf("expected %v, got %v", nil, n2.Prev())
			}
			if n2.Next() != nil {
				t.Errorf("expected %v, got %v", nil, n2.Next())
			}
		})

		t.Run("Pop the last node", func(t *testing.T) {
			ll := linkedlist.New[int]()
			n1 := node.New(1)
			n2 := node.New(2)
			n3 := node.New(3)
			ll.Insert(n1)
			ll.Insert(n2)
			ll.Insert(n3)

			out := ll.Pop()
			if out != n1 {
				t.Errorf("expected %v, got %v", n1, out)
			}
			if ll.Head() != n2 {
				t.Errorf("expected %v, got %v", n2, ll.Head())
			}
			if ll.Tail() != n3 {
				t.Errorf("expected %v, got %v", n3, ll.Tail())
			}
			if n1.Prev() != nil {
				t.Errorf("expected %v, got %v", nil, n1.Prev())
			}
			if n1.Next() != nil {
				t.Errorf("expected %v, got %v", nil, n1.Next())
			}
			if n2.Prev() != nil {
				t.Errorf("expected %v, got %v", nil, n2.Prev())
			}
			if n2.Next() != n3 {
				t.Errorf("expected %v, got %v", n3, n2.Next())
			}
			if n3.Prev() != n2 {
				t.Errorf("expected %v, got %v", n2, n3.Prev())
			}
			if n3.Next() != nil {
				t.Errorf("expected %v, got %v", nil, n3.Next())
			}
		})

		t.Run("Pop all nodes in the list", func(t *testing.T) {
			ll := linkedlist.New[int]()
			n1 := node.New(1)
			n2 := node.New(2)
			n3 := node.New(3)
			ll.Insert(n1)
			ll.Insert(n2)
			ll.Insert(n3)

			out := ll.Pop()
			if out != n1 {
				t.Errorf("expected %v, got %v", n1, out)
			}
			if ll.Head() != n2 {
				t.Errorf("expected %v, got %v", n2, ll.Head())
			}
			if ll.Tail() != n3 {
				t.Errorf("expected %v, got %v", n3, ll.Tail())
			}
			if n1.Prev() != nil {
				t.Errorf("expected %v, got %v", nil, n1.Prev())
			}
			if n1.Next() != nil {
				t.Errorf("expected %v, got %v", nil, n1.Next())
			}
			if n2.Prev() != nil {
				t.Errorf("expected %v, got %v", nil, n2.Prev())
			}
			if n2.Next() != n3 {
				t.Errorf("expected %v, got %v", n3, n2.Next())
			}
			if n3.Prev() != n2 {
				t.Errorf("expected %v, got %v", n2, n3.Prev())
			}
			if n3.Next() != nil {
				t.Errorf("expected %v, got %v", nil, n3.Next())
			}

			out = ll.Pop()
			if out != n2 {
				t.Errorf("expected %v, got %v", n2, out)
			}
			if ll.Head() != n3 {
				t.Errorf("expected %v, got %v", n3, ll.Head())
			}
			if ll.Tail() != n3 {
				t.Errorf("expected %v, got %v", n3, ll.Tail())
			}
			if n2.Prev() != nil {
				t.Errorf("expected %v, got %v", nil, n2.Prev())
			}
			if n2.Next() != nil {
				t.Errorf("expected %v, got %v", nil, n2.Next())
			}
			if n3.Prev() != nil {
				t.Errorf("expected %v, got %v", nil, n3.Prev())
			}
			if n3.Next() != nil {
				t.Errorf("expected %v, got %v", nil, n3.Next())
			}

			out = ll.Pop()
			if out != n3 {
				t.Errorf("expected %v, got %v", n3, out)
			}
			if ll.Head() != nil {
				t.Errorf("expected %v, got %v", nil, ll.Head())
			}
			if ll.Tail() != nil {
				t.Errorf("expected %v, got %v", nil, ll.Tail())
			}
			if n3.Prev() != nil {
				t.Errorf("expected %v, got %v", nil, n3.Prev())
			}
			if n3.Next() != nil {
				t.Errorf("expected %v, got %v", nil, n3.Next())
			}
		})

		t.Run("Pop an empty list", func(t *testing.T) {
			ll := linkedlist.New[int]()
			out := ll.Pop()
			if out != nil {
				t.Errorf("expected %v, got %v", nil, out)
			}
			if ll.Head() != nil {
				t.Errorf("expected %v, got %v", nil, ll.Head())
			}
			if ll.Tail() != nil {
				t.Errorf("expected %v, got %v", nil, ll.Tail())
			}
		})
	})

	t.Run("Size", func(t *testing.T) {
		t.Run("Size of an empty list", func(t *testing.T) {
			ll := linkedlist.New[int]()
			if ll.Size() != 0 {
				t.Errorf("expected %v, got %v", 0, ll.Size())
			}
		})

		t.Run("Size of a list with one node", func(t *testing.T) {
			ll := linkedlist.New[int]()
			n1 := node.New(1)
			ll.Insert(n1)
			if ll.Size() != 1 {
				t.Errorf("expected %v, got %v", 1, ll.Size())
			}
		})

		t.Run("Size of a list with multiple nodes", func(t *testing.T) {
			ll := linkedlist.New[int]()
			n1 := node.New(1)
			n2 := node.New(2)
			n3 := node.New(3)
			ll.Insert(n1)
			ll.Insert(n2)
			ll.Insert(n3)
			if ll.Size() != 3 {
				t.Errorf("expected %v, got %v", 3, ll.Size())
			}
		})

		t.Run("Size of a list after popping all nodes", func(t *testing.T) {
			ll := linkedlist.New[int]()
			n1 := node.New(1)
			n2 := node.New(2)
			n3 := node.New(3)
			ll.Insert(n1)
			ll.Insert(n2)
			ll.Insert(n3)
			ll.Pop()
			ll.Pop()
			ll.Pop()
			if ll.Size() != 0 {
				t.Errorf("expected %v, got %v", 0, ll.Size())
			}
		})

		t.Run("Size of a list after deleting nodes", func(t *testing.T) {
			ll := linkedlist.New[int]()
			n1 := node.New(1)
			n2 := node.New(2)
			n3 := node.New(3)
			ll.Insert(n1)
			ll.Insert(n2)
			ll.Insert(n3)
			ll.Delete(n2)
			if ll.Size() != 2 {
				t.Errorf("expected %v, got %v", 2, ll.Size())
			}
		})
	})
}
