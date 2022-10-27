package main

import "fmt"

func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

// example of generics with single linked list

type List[T any] struct {
	head, tail *Node[T]
}

type Node[T any] struct {
	next  *Node[T]
	value T
}

func (list *List[T]) Push(val T) {
	if list.tail == nil {
		list.head = &Node[T]{
			value: val,
		}

		list.tail = list.head
	} else {
		list.tail.next = &Node[T]{value: val}
		list.tail = list.tail.next
	}
}

func (list *List[T]) Query() []T {
	var nodes []T
	for n := list.head; n != nil; n = n.next {
		nodes = append(nodes, n.value)
	}

	return nodes

}

func main() {
	mk := map[string]int{
		"1": 1,
	}

	fmt.Println(MapKeys(mk))

	LIST := List[int]{}
	LIST.Push(10)
	LIST.Push(20)
	LIST.Push(30)
	LIST.Push(40)
	res := LIST.Query()
	fmt.Println(res)
}
