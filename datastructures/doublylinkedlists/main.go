package main

import "fmt"
import "strconv"
import "strings"

type node struct {
	data int
	next *node
	prev *node
}

type doublylinkedList struct {
	length int
	head   *node
	tail   *node
}

func (ll *doublylinkedList) add(data int) {
	nd := &node{
		data: data,
	}
	if ll.head == nil {
		ll.head = nd
		ll.tail = nd
	} else {
		ll.tail.next = nd
		nd.prev = ll.tail
		ll.tail = nd
	}
	ll.length++
}

func (ll *doublylinkedList) remove(index int) int {
	i := 1
	if index < 0 && index > ll.length {
		return -1
	}
	current := ll.head

	if index == 0 {
		ll.head = current.next
		if ll.head == nil {
			ll.tail = nil
		} else {
			ll.head.prev = nil
		}
	} else if index == ll.length {
		current = ll.tail
		ll.tail = current.prev
		ll.tail.next = nil
	} else {
		for i < index {
			current = current.next
			i++
		}

		current.next.prev = current.prev
		current.prev.next = current.next
	}
	ll.length--
	return current.data
}

func (ll *doublylinkedList) item(index int) int {
	i := 1

	if index < 0 && index > ll.length {
		return -1
	}
	current := ll.head

	for i < index {
		current = current.next
		i++
	}

	return current.data
}

func (ll *doublylinkedList) toArray() []int {
	a := make([]int, ll.length)
	i := 0
	current := ll.head
	for current != nil {
		a[i] = current.data
		i++
		current = current.next
	}
	return a
}

func (ll *doublylinkedList) toString() string {
	a := ll.toArray()
	as := make([]string, len(a))
	l := len(a)
	for i := 0; i < l; i++ {
		as[i] = strconv.Itoa(a[i])
	}
	return strings.Join(as, ",")
}

func (ll *doublylinkedList) printList() {
	current := ll.head
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
	fmt.Println()
}

func (ll *doublylinkedList) size() int {
	return ll.length
}

func main() {
	ll := &doublylinkedList{}
	ll.add(2)
	ll.printList()
	ll.add(3)
	ll.printList()
	ll.add(4)
	ll.printList()
	ll.add(5)
	ll.printList()
	ll.add(6)
	ll.printList()
	ll.remove(3)
	ll.printList()
	fmt.Printf("item at 3rd position %d ", ll.item(3))
	fmt.Println()
	fmt.Printf("size of linked list %d", ll.size())
	fmt.Println()
	fmt.Printf("to array of linked list %v", ll.toArray())
	fmt.Println()
	fmt.Printf("to string of linked list %v", ll.toString())
}
