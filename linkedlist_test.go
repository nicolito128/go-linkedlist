package linkedlist_test

import (
	"testing"

	"github.com/nicolito128/go-linkedlist"
)

func TestAppendEmptyList(t *testing.T) {
	list := linkedlist.NewList[int]()

	list.Append(10)
	got := list.Head().Value()
	went := 10

	if got != went {
		t.Errorf("Output %q not equal to expected %q", got, went)
	}

	list.Append(5)
	got = list.Tail().Value()
	went = 5
	if got != went {
		t.Errorf("Output %q not equal to expected %q", got, went)
	}
}

func TestAppend(t *testing.T) {
	list := linkedlist.NewList[int]()

	list.Append(10)
	list.Append(5)
	list.Append(1)

	got := list.Tail().Value()
	went := 1
	if got != went {
		t.Errorf("Output %q not equal to expected %q", got, went)
	}
}

func TestUnshiftEmptyList(t *testing.T) {
	list := linkedlist.NewList[int]()

	list.Unshift(10)
	got := list.Head().Value()
	went := 10

	if got != went {
		t.Errorf("Output %q not equal to expected %q", got, went)
	}
}

func TestUnshift(t *testing.T) {
	list := linkedlist.NewList[int]()

	list.Unshift(8)
	list.Unshift(4)
	list.Unshift(2)
	got := list.Head().Value()
	went := 2

	if got != went {
		t.Errorf("Output %q not equal to expected %q", got, went)
	}
}

func TestPop(t *testing.T) {
	list := linkedlist.NewList[int]()
	list.Append(1)
	got, err := list.Pop()
	went := 1

	if err != nil {
		t.Errorf("Error popping an unique element in the list")
	}

	if got != went {
		t.Errorf("Output %q not equal to expected %q", got, went)
	}
}

func TestClear(t *testing.T) {
	list := linkedlist.NewList[int](2, 4, 8, 16)
	list.Clear()

	if list.Head() != nil {
		t.Errorf("Clear did not dereference the head of the list")
	}

	if list.Tail() != nil {
		t.Errorf("Clear did not dereference the tail of the list")
	}
}

func TestEach(t *testing.T) {
	list := linkedlist.NewList[int](1, 2, 3)

	count := 0
	linkedlist.Each(list, func(node *linkedlist.Node[int]) {
		if node != nil {
			count++
		}
	})

	if count != list.Len() {
		t.Errorf("Output %q not equal to expected %q", count, list.Len())
	}
}

func TestRemoveUnique(t *testing.T) {
	list := linkedlist.NewList[int](1)
	linkedlist.Remove(list, 1)

	if list.Head() != nil || list.Tail() != nil {
		t.Errorf("Remove did not delete the item")
	}
}

func TestRemoveBetween(t *testing.T) {
	list := linkedlist.NewList[int](1, 2, 3)
	linkedlist.Remove(list, 2)

	got := list.Head().Next().Value()
	went := 3

	if got != went {
		t.Errorf("Output %q not equal to expected %q", got, went)
	}
}

func TestRemoveLast(t *testing.T) {
	list := linkedlist.NewList[int](1, 2, 3)
	linkedlist.Remove(list, 3)

	got := list.Head().Next().Value()
	went := 2

	if got != went {
		t.Errorf("Output %q not equal to expected %q", got, went)
	}
}
