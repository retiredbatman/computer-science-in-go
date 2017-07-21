package main

import "fmt"
import "math"

type minHeap struct {
	harr     []int
	capacity int
	heapSize int
}

func new(cap int) *minHeap {
	a := make([]int, cap)
	return &minHeap{
		capacity: cap,
		harr:     a,
	}
}

func (mh *minHeap) insertKey(key int) {
	if mh.heapSize == mh.capacity {
		fmt.Println("heap is already at max capacity")
		return
	}
	//put it in the last position
	mh.heapSize++
	i := mh.heapSize - 1
	mh.harr[i] = key
	//move up to keep heap property
	for mh.harr[parent(i)] > mh.harr[i] {
		swap(mh.harr, parent(i), i)
		i = parent(i)
	}
}

func (mh *minHeap) deleteKey(i int) {
	if i >= mh.heapSize {
		fmt.Println("index not present in the heap")
		return
	}
	mh.decreaseKey(i, int(math.MinInt64))
	mh.extractMin()
}

func (mh *minHeap) decreaseKey(i int, newVal int) {
	mh.harr[i] = newVal
	for mh.harr[parent(i)] > mh.harr[i] {
		swap(mh.harr, parent(i), i)
		i = parent(i)
	}
}

func (mh *minHeap) extractMin() int {
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
	mh.minHeapify(0)
	return t
}

func (mh *minHeap) minHeapify(i int) {
	l := left(i)
	r := right(i)
	smallest := i
	if l < mh.heapSize && mh.harr[l] < mh.harr[i] {
		smallest = l
	}
	if r < mh.heapSize && mh.harr[r] < mh.harr[smallest] {
		smallest = r
	}
	if smallest != i {
		swap(mh.harr, i, smallest)
		mh.minHeapify(smallest)
	}
}

func (mh *minHeap) getMin() int {
	return mh.harr[0]
}

func (mh *minHeap) printHeap() {
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
	mh.insertKey(3)
	mh.insertKey(2)
	fmt.Println("---------------------------")
	mh.printHeap()
	mh.deleteKey(1)
	fmt.Println("---------------------------")
	mh.printHeap()
	mh.insertKey(15)
	mh.insertKey(5)
	mh.insertKey(4)
	mh.insertKey(45)
	fmt.Println("---------------------------")
	mh.printHeap()
	fmt.Printf("extract min in heap %d\n", mh.extractMin())
	fmt.Printf("min in heap %d\n", mh.getMin())
	fmt.Println("---------------------------")
	mh.printHeap()
	mh.decreaseKey(2, 1)
	fmt.Println("---------------------------")
	mh.printHeap()
	fmt.Printf("min in heap %d\n", mh.getMin())
}
