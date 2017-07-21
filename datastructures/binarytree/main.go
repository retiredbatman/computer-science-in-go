package main

import (
	"fmt"
)

var a = []int{
	1, 4, 5, 6, 8, 10, 11, 12, 12, 13, 14, 15, 19, 20, 20,
}

type node struct {
	data  int
	left  *node
	right *node
}
type qNode struct {
	data *node
	next *qNode
}

type que struct {
	front *qNode
	rear  *qNode
}

func (q *que) enque(data *node) {
	qn := &qNode{
		data: data,
	}
	//means que is empty
	if q.rear == nil {
		q.front = qn
		q.rear = qn
		return
	}
	q.rear.next = qn
	q.rear = q.rear.next
}

func (q *que) deque() *qNode {
	if q.front == nil {
		return nil
	}
	t := q.front
	q.front = q.front.next
	if q.front == nil {
		q.rear = nil
	}
	return t
}

func newNode(data int) *node {
	nd := &node{
		data: data,
	}
	return nd
}

func inOrder(root *node) {
	if root != nil {
		inOrder(root.left)
		fmt.Printf("%d \n", root.data)
		inOrder(root.right)

	}
}

func preOrder(root *node) {
	if root != nil {
		fmt.Printf("%d\n", root.data)
		preOrder(root.left)
		preOrder(root.right)
	}
}

func postOrder(root *node) {
	if root != nil {
		postOrder(root.left)
		postOrder(root.right)
		fmt.Printf("%d\n", root.data)
	}
}

func height(nd *node) int {
	if nd == nil {
		return 0
	}
	lh := height(nd.left)
	rh := height(nd.right)
	if lh > rh {
		return lh + 1
	}
	return rh + 1
}

func printGivenLabel(nd *node, level int) {
	if nd == nil {
		return
	}
	if level == 1 {
		fmt.Printf("%d\n", nd.data)
	}
	if level > 1 {
		printGivenLabel(nd.left, level-1)
		printGivenLabel(nd.right, level-1)
	}
}

func printGivenOrder(nd *node) {
	h := height(nd)
	fmt.Printf("height %d\n", h)
	for i := 1; i <= h; i++ {
		printGivenLabel(nd, i)
	}

}

func printLevelOrderQue(root *node) {
	q := &que{}
	t := root
	for t != nil {
		fmt.Printf("%d\n", t.data)
		if t.left != nil {
			q.enque(t.left)
		}
		if t.right != nil {
			q.enque(t.right)
		}
		d := q.deque()
		if d == nil {
			t = nil
		} else {
			t = d.data
		}
	}

}

func main() {
	root := newNode(1)
	root.left = newNode(2)
	root.right = newNode(3)
	root.left.left = newNode(4)
	root.left.right = newNode(5)
	//Depth First Algorithms
	fmt.Println("Inorder traversal")
	inOrder(root)
	fmt.Println("preorder traversal")
	preOrder(root)
	fmt.Println("postorder traversal")
	postOrder(root)

	//Breadth first search

	fmt.Println("bfs")
	printGivenOrder(root)

	//Breadth first search queue
	fmt.Println("bfs que ")
	printLevelOrderQue(root)

}
