# LinkedList
The package offers some useful implementations for using and working with linked lists.

However, it is usually easier to work with slices in the vast majority of cases rather than a linked list. This package only has the advantage of being written using generics.

## Getting Started

Start a new project and get the module:

    go mod init example
    go get github.com/nicolito128/go-linkedlist

Import and use the module:

```go
package main

import (
	"fmt"

	"github.com/nicolito128/go-linkedlist"
)

func main() {
	list := linkedlist.NewList[int](4, 8, 16, 32, 64)

	list.Append(128) // Add to the end
	list.Pop() // Remove the first element (4)
    list.Unshift(2) // Add to the beginning

	fmt.Println(list) // 2 -> 8 -> 16 -> 32 -> 64 -> 128
}
```
