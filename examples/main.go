package main

import (
	"fmt"

	"github.com/nicolito128/go-linkedlist"
)

func main() {
	list := linkedlist.NewList[int](1, 2, 3, 4, 5)
	fmt.Println(list.Head(), list.Tail())

	linkedlist.Remove(list, 3)
	linkedlist.Remove(list, 5)
	fmt.Println(list)

	linkedlist.Filter(list, func(v int) bool {
		return (v % 2) == 0
	})

	list.Unshift(6)
	fmt.Println(list)

	list.Clear()
	fmt.Println(list.Head(), list.Tail())
}
