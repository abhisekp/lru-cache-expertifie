# LRU Cache Implementation in Go
> Uses double linked list and hashmap to implement LRU Cache with generics support

[![GoDoc](https://godoc.org/github.com/abhisekp/lru-cache-expertifie/lru?status.svg)](https://godoc.org/github.com/abhisekp/lru-cache-expertifie/lru)
[![Build Status](https://travis-ci.org/abhisekp/lru-cache-expertifie.svg?branch=master)](https://travis-ci.org/abhisekp/lru-cache-expertifie)
[![Coverage Status](https://coveralls.io/repos/github/abhisekp/lru-cache-expertifie/badge.svg?branch=master)](https://coveralls.io/github/abhisekp/lru-cache-expertifie?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/abhisekp/lru-cache-expertifie)](https://goreportcard.com/report/github.com/abhisekp/lru-cache-expertifie)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/abhisekp/lru-cache-expertifie/master/LICENSE)

---

# Usage
```go
package main
import (
	"fmt"
	"github.com/abhisekp/lru-cache-expertifie/lru"
)

func main() {
	lruCache := lru.New[int](3)
	val, found := lruCache.Get(1)
    if found {
        // do something with val
        fmt.Println(val)
    }
}
```

# LICENSE
[MIT](./LICENSE)

# Author
[Abhisek Pattnaik](https://about.me/abhisekp)