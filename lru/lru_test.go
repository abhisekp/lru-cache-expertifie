package lru_test

import (
	"testing"

	"lru/lru"
)

func TestLRU(t *testing.T) {
	t.Run("Get from cache", func(t *testing.T) {
		lruCache := lru.New[int](lru.Options{
			MaxEntries: 4,
		})

		val, found := lruCache.Get(1)
		if found != false {
			t.Errorf("Expected %v, got %v", false, found)
		}
		if val != 1 {
			t.Errorf("Expected %v, got %v", 1, val)
		}

		val, found = lruCache.Get(2)
		if found != false {
			t.Errorf("Expected %v, got %v", false, found)
		}
		if val != 2 {
			t.Errorf("Expected %v, got %v", 2, val)
		}

		val, found = lruCache.Get(3)
		if found != false {
			t.Errorf("Expected %v, got %v", false, found)
		}
		if val != 3 {
			t.Errorf("Expected %v, got %v", 3, val)
		}

		val, found = lruCache.Get(1)
		if found != true {
			t.Errorf("Expected %v, got %v", true, found)
		}
		if val != 1 {
			t.Errorf("Expected %v, got %v", 1, val)
		}
	})

	t.Run("Get from cache with more than max size", func(t *testing.T) {
		lruCache := lru.New[int](lru.Options{
			MaxEntries: 4,
		})

		val, found := lruCache.Get(1)
		if found != false {
			t.Errorf("Expected %v, got %v", false, found)
		}
		if val != 1 {
			t.Errorf("Expected %v, got %v", 1, val)
		}

		val, found = lruCache.Get(2)
		if found != false {
			t.Errorf("Expected %v, got %v", false, found)
		}
		if val != 2 {
			t.Errorf("Expected %v, got %v", 2, val)
		}

		val, found = lruCache.Get(3)
		if found != false {
			t.Errorf("Expected %v, got %v", false, found)
		}
		if val != 3 {
			t.Errorf("Expected %v, got %v", 3, val)
		}

		val, found = lruCache.Get(4)
		if found != false {
			t.Errorf("Expected %v, got %v", false, found)
		}
		if val != 4 {
			t.Errorf("Expected %v, got %v", 4, val)
		}

		val, found = lruCache.Get(5)
		if found != false {
			t.Errorf("Expected %v, got %v", false, found)
		}
		if val != 5 {
			t.Errorf("Expected %v, got %v", 5, val)
		}

		val, found = lruCache.Get(1)
		if found != false {
			t.Errorf("Expected %v, got %v", false, found)
		}
		if val != 1 {
			t.Errorf("Expected %v, got %v", 1, val)
		}

		val, found = lruCache.Get(2)
		if found != false {
			t.Errorf("Expected %v, got %v", false, found)
		}
		if val != 2 {
			t.Errorf("Expected %v, got %v", 2, val)
		}

		val, found = lruCache.Get(3)
		if found != false {
			t.Errorf("Expected %v, got %v", false, found)
		}
		if val != 3 {
			t.Errorf("Expected %v, got %v", 3, val)
		}

	})
}
