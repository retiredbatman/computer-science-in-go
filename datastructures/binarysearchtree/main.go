package main

import (
	"fmt"
)

type node struct {
	data  int
	left  *node
	right *node
}

type binarySearchTree struct {
	root *node
}

func (bst *binarySearchTree) add(data int) {
	nd := &node{
		data: data,
	}
	current := bst.root
	if bst.root == nil {
		bst.root = nd
	} else {
		for {
			//if data less than current value travel to the left side
			if data < current.data {
				if current.left == nil { //if no left node
					current.left = nd
					break
				} else { //go to next left node
					current = current.left
				}
			} else if data > current.data {
				if current.right == nil {
					current.right = nd
					break
				} else {
					current = current.right
				}
			} else {
				break // break because no duplicates
			}
		}
	}
}

func (bst *binarySearchTree) remove(value int) {
	deleteNode(bst.root, value)
}

func deleteNode(nd *node, value int) *node {
	//if the value to delete is smaller than node value it must be in the
	//left sub tree
	if value < nd.data {
		nd.left = deleteNode(nd.left, value)
	} else if value > nd.data { //if the value to delete is greater than node value it must be in the
		//right sub tree
		nd.right = deleteNode(nd.right, value)
	} else { //if value is same as this node , then delete the node
		if nd.left == nil { // if left is  not there return right node to be set as the root node for the sub tree
			temp := nd.right
			nd = nil
			return temp
		} else if nd.right == nil { // if right is  not there return left node to be set as the root node for the sub tree
			temp := nd.left
			nd = nil
			return temp
		}
		//if two children for the node
		temp := minValueNode(nd.right)
		nd.data = temp.data
		nd.right = deleteNode(nd.right, temp.data)
	}
	return nd
}

func (bst *binarySearchTree) contains(value int) bool {
	return search(bst.root, value)
}

func search(nd *node, value int) bool {
	if nd == nil {
		return false
	}
	if value < nd.data {
		return search(nd.left, value)
	} else if value > nd.data {
		return search(nd.right, value)
	} else {
		return true
	}
}

//get minimum value in the right sub tree
func minValueNode(nd *node) *node {
	current := nd
	for current.left != nil {
		current = current.left
	}
	return current
}

func inOrder(root *node) {
	if root != nil {
		inOrder(root.left)
		fmt.Printf("%d \n", root.data)
		inOrder(root.right)
	}
}

func main() {
	bst := &binarySearchTree{}
	bst.add(40)
	fmt.Printf("root element is %d \n", bst.root.data)
	bst.add(30)
	fmt.Printf("root element is %d \n", bst.root.data)
	bst.add(20)
	fmt.Printf("root element is %d \n", bst.root.data)
	bst.add(50)
	fmt.Printf("root element is %d \n", bst.root.data)
	bst.add(70)
	fmt.Printf("root element is %d \n", bst.root.data)
	bst.add(60)
	fmt.Printf("root element is %d \n", bst.root.data)
	bst.add(80)
	fmt.Printf("root element is %d \n", bst.root.data)
	inOrder(bst.root)
	fmt.Printf("root element is %d \n", bst.root.data)
	bst.remove(20)
	inOrder(bst.root)
	fmt.Printf("root element is %d \n", bst.root.data)
	bst.remove(30)
	inOrder(bst.root)
	fmt.Printf("root element is %d \n", bst.root.data)
	bst.remove(50)
	inOrder(bst.root)
	fmt.Printf("root element is %d \n", bst.root.data)
	fmt.Printf("tree contains 40 %v", bst.contains(40))

}
