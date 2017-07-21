package main

import "fmt"
import "math"

type maxHeap struct {
	harr     []int
	capacity int
	heapSize int
}

func new(cap int) *maxHeap {
	a := make([]int, cap)
	return &maxHeap{
		capacity: cap,
		harr:     a,
	}
}

func (mh *maxHeap) insertKey(key int) {
	if mh.heapSize == mh.capacity {
		fmt.Println("heap is already at max capacity")
		return
	}
	//put it in the last position
	mh.heapSize++
	i := mh.heapSize - 1
	mh.harr[i] = key
	//move up to keep heap property
	for mh.harr[parent(i)] < mh.harr[i] {
		swap(mh.harr, parent(i), i)
		i = parent(i)
	}
}

func (mh *maxHeap) deleteKey(i int) {
	if i >= mh.heapSize {
		fmt.Println("index not present in the heap")
		return
	}

	mh.increaseKey(i, int(math.MaxInt64))
	mh.extractMax()
}

func (mh *maxHeap) increaseKey(i int, newVal int) {
	mh.harr[i] = newVal
	for mh.harr[parent(i)] < mh.harr[i] {
		swap(mh.harr, parent(i), i)
		i = parent(i)
	}
}

func (mh *maxHeap) extractMax() int {
	if mh.heapSize <= 0 {
		fmt.Println("no data in heap")
		return int(math.Inf(-1))
	}
	if mh.heapSize == 1 {
		mh.heapSize--
		return mh.harr[0]
	}
	t := mh.harr[0]
	mh.harr[0] = mh.harr[mh.heapSize-1]
	mh.heapSize--
	mh.maxHeapify(0)
	return t
}

func (mh *maxHeap) maxHeapify(i int) {
	l := left(i)
	r := right(i)
	greatest := i
	if l < mh.heapSize && mh.harr[l] > mh.harr[i] {
		greatest = l
	}
	if r < mh.heapSize && mh.harr[r] > mh.harr[greatest] {
		greatest = r
	}
	if greatest != i {
		swap(mh.harr, i, greatest)
		mh.maxHeapify(greatest)
	}
}

func (mh *maxHeap) getMax() int {
	return mh.harr[0]
}

func (mh *maxHeap) printHeap() {
	for i := 0; i < mh.heapSize; i++ {
		fmt.Printf("%d\n", mh.harr[i])
	}
}

func swap(a []int, i int, j int) {
	temp := a[i]
	a[i] = a[j]
	a[j] = temp
}

func parent(i int) int {
	return i / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func main() {
	mh := new(10)
	mh.insertKey(1)
	mh.insertKey(23)
	mh.insertKey(12)
	mh.insertKey(9)
	mh.insertKey(30)
	mh.insertKey(2)
	mh.insertKey(50)
	k := 3
	for i := 0; i < k; i++ {
		fmt.Printf("%d\n", mh.extractMax())
	}
}
