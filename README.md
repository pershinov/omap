# omap

Ordered map implementation using generics, native map and linked list.
! NO cuncurent safe. Use mutex for concurrent rw.

## Usage

### Get the go-lib module

```bash
go get github.com/pershinov/omap@v1.0.1
```

### Example
```go
package main

import (
	"fmt"
	
	"github.com/pershinov/omap"
)

func main() {
	om := omap.New[string, int]().WithCap(10)

	// Set
	om.Set("test", 1)
	om.Set("test2", 2)
	om.Set("test3", 3)

	// Get
	val, ok := om.Get("test")
	fmt.Println(val, ok) // 1, true

	// Reset (to the end of order)
	om.Set("test", 1)

	// Replace (no change order)
	ok = om.Replace("test2", 10)
	fmt.Println(ok) // true

	// Delete
	ok = om.Delete("test")
	fmt.Println(ok) // true

	om.Iter(func(key string, value int) {
		fmt.Println(key, value)
	})

	om.IterBack(func(key string, value int) {
		fmt.Println(key, value)
	})
}

```

## Have a good vibe ^..^