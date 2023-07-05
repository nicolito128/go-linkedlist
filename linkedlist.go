/*
The package offers some useful implementations for using and working with linked lists.

However, it is usually easier to work with slices in the vast majority of cases rather than a linked list.
This package only has the advantage of being written using generics.
*/
package linkedlist

import (
	"errors"
	"fmt"
	"sync"
)

var ErrEmptyList = errors.New("empty list")

// Represents a node structure.
type Node[T interface{}] struct {
	value T
	next  *Node[T]
}

func NewNode[T interface{}](initValue T) *Node[T] {
	return &Node[T]{initValue, nil}
}

func (nd *Node[T]) Value() T {
	return nd.value
}

func (nd *Node[T]) Next() *Node[T] {
	return nd.next
}

// List is a LinkedList implementation.
type List[T interface{}] struct {
	head *Node[T]
	tail *Node[T]
	size int
	mu   *sync.Mutex
}

func NewList[T interface{}](args ...T) *List[T] {
	res := new(List[T])
	res.mu = new(sync.Mutex)

	if len(args) > 0 {
		for _, v := range args {
			res.Append(v)
		}
	}

	return res
}

func (l *List[T]) Head() *Node[T] {
	return l.head
}

func (l *List[T]) Tail() *Node[T] {
	return l.tail
}

func (l List[T]) Len() int {
	return l.size
}

func (l *List[T]) IsEmpty() bool {
	return IsNil(l.head)
}

func (l *List[T]) Append(value T) {
	l.mu.Lock()
	defer l.mu.Unlock()

	newNode := NewNode(value)
	l.size++

	if l.head == nil {
		l.head = newNode
		l.tail = l.head
		return
	}

	l.tail.next = newNode
	l.tail = newNode
}

func (l *List[T]) Unshift(value T) {
	l.mu.Lock()
	defer l.mu.Unlock()

	newNode := NewNode(value)
	l.size++

	if l.head == nil {
		l.head = newNode
		l.tail = l.head
		return
	}

	newNode.next = l.head
	l.head = newNode
}

func (l *List[T]) Pop() (T, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	var res T

	if l.head != nil {
		res = l.head.value
		l.head = l.head.next
		l.size--

		if l.head == nil {
			l.tail = nil
		}
	} else {
		return res, ErrEmptyList
	}

	return res, nil
}

func (l *List[T]) Clear() {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.head = nil
	l.tail = nil
}

func (l *List[T]) String() string {
	var res string
	Each(l, func(node *Node[T]) {
		res += fmt.Sprintf("%v", node.Value())

		if node.Next() != nil {
			res += " -> "
		}
	})

	return res
}
