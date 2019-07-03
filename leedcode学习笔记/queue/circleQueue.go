package main

import "fmt"

func main() {

	obj := Constructor(3)

	fmt.Println(fmt.Sprintf("EnQueue1:%+v", obj.EnQueue(1)))
	fmt.Println(fmt.Sprintf("Obj:%+v", obj))
	fmt.Println(fmt.Sprintf("EnQueue2%+v", obj.EnQueue(2)))
	fmt.Println(fmt.Sprintf("Obj:%+v", obj))
	fmt.Println(fmt.Sprintf("Enqueue2%+v", obj.EnQueue(3)))
	fmt.Println(fmt.Sprintf("Obj:%+v", obj))
	fmt.Println(fmt.Sprintf("Enqueue3:%+v", obj.EnQueue(4)))
	fmt.Println(fmt.Sprintf("Obj:%+v", obj))
	fmt.Println(fmt.Sprintf("Dequeue:%+v", obj.DeQueue()))
	fmt.Println(fmt.Sprintf("Obj:%+v", obj))
	fmt.Println(fmt.Sprintf("Enqueue3:%+v", obj.EnQueue(5)))
	fmt.Println(fmt.Sprintf("Obj:%+v", obj))
	fmt.Println(fmt.Sprintf("Front:%+v", obj.Front()))
	fmt.Println(fmt.Sprintf("Rear%+v", obj.Rear()))
	fmt.Println(fmt.Sprintf("IsEmpty:%+v", obj.IsEmpty()))
	fmt.Println(fmt.Sprintf("IsFull:%+v", obj.IsFull()))

}

type MyCircularQueue struct {
	Size       int
	SizeMax    int
	Queue      []int
	IndexFirst int
	IndexLast  int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		Size:       0,
		SizeMax:    k,
		IndexFirst: 0,
		IndexLast:  -1,
	}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		fmt.Println("this queue is full!!!")
		return false
	}
	this.IndexLast++
	if this.IndexLast >= this.SizeMax {
		this.IndexLast = 0
	}
	if this.IndexLast >= len(this.Queue)-1 {
		this.Queue = append(this.Queue, value)
	} else {
		this.Queue[this.IndexLast] = value
	}
	this.Size++
	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		fmt.Println("this queue is empty!!!")
		return false
	}
	this.IndexFirst++
	if this.IndexFirst >= this.SizeMax {
		this.IndexFirst = 0
	}
	this.Size--
	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		fmt.Println("this queue is empty")
		return -1
	}
	return this.Queue[this.IndexFirst]
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		fmt.Println("this queue is empty")
		return -1
	}
	return this.Queue[this.IndexLast]
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	return this.Size == 0
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	return this.Size == this.SizeMax
}
