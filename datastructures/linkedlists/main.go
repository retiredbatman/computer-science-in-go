package main

import "fmt"
import "strconv"
import "strings"

type node struct {
	data int
	next *node
}

type linkedList struct {
	length int
	head   *node
}

func (ll *linkedList) add(data int) {
	nd := &node{
		data: data,
	}
	if ll.head == nil {
		ll.head = nd
	} else {
		current := ll.head
		for current.next != nil {
			current = current.next
		}
		current.next = nd
	}
	ll.length++
}

func (ll *linkedList) remove(index int) int {
	i := 1
	if index < 0 && index > ll.length {
		return -1
	}
	current := ll.head
	var previous *node
	if index == 0 {
		ll.head = current.next
	} else {
		for i < index {
			previous = current
			current = current.next
			i++
		}

		previous.next = current.next
	}
	ll.length--
	return current.data
}

func (ll *linkedList) item(index int) int {
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

func (ll *linkedList) toArray() []int {
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

func (ll *linkedList) toString() string {
	a := ll.toArray()
	as := make([]string, len(a))
	l := len(a)
	for i := 0; i < l; i++ {
		as[i] = strconv.Itoa(a[i])
	}
	return strings.Join(as, ",")
}

func (ll *linkedList) printList() {
	current := ll.head
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
	fmt.Println()
}

func (ll *linkedList) size() int {
	return ll.length
}

func main() {
	ll := &linkedList{}
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
