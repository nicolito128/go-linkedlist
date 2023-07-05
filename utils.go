package linkedlist

type EachCallback[T interface{}] func(*Node[T])
type FilterCallback[T comparable] func(T) bool

func IsNil[T interface{}](node *Node[T]) bool {
	return node == nil
}

// Each runs through any linked list, and allows you to interact on each node with a callback
func Each[T interface{}](list *List[T], callback EachCallback[T]) {
	temp := list.head
	for temp != nil {
		callback(temp)
		temp = temp.Next()
	}
}

// Filter removes any node if its value does not meet the condition returned in the callback.
// The type of the linked list must be a comparable.
func Filter[T comparable](list *List[T], callback FilterCallback[T]) {
	temp := list.head
	for temp != nil && !list.IsEmpty() {
		ok := callback(temp.Value())
		if ok {
			temp = temp.next
			continue
		}

		aux := temp.value
		temp = temp.next
		Remove(list, aux)
	}
}

// Remove the first occurency of an element.
// The type of the linked list must be a comparable.
func Remove[T comparable](list *List[T], mark T) {
	list.mu.Lock()
	defer list.mu.Unlock()

	if list.head != nil {
		if list.head.value == mark {
			list.head = list.head.next
			list.size--

			if list.head == nil {
				list.tail = nil
			}

			return
		}

		var cur, prev *Node[T]
		cur = list.head

		for cur.next != nil && cur.value != mark {
			prev = cur
			cur = cur.next
		}

		if prev != nil {
			if cur.value == mark {
				// If the element was found in the tail
				if cur.next == nil {
					prev.next = nil
					list.tail = prev
				} else {
					prev.next = cur.next
				}

				list.size--
			}
		}
	}
}
