/*
反转一个单链表。

示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL

如下方法使用的迭代的方式完成的反转联表
迭代时有两个要点
1. 需要一个临时的迭代变量
2. 需要一个临时的变量储存上一个节点
*/
package main

import "fmt"

func main() {
	node := newNode()
	p(node)
	nodeRev := revNode(node)
	p(nodeRev)
}

func p(n *Node) {
	for n != nil {
		fmt.Println(fmt.Sprintf("%+v", n))
		n = n.Next
	}
}

/*
  节点只能记录下一个节点，和本身的一些属性
*/
type Node struct {
	Next    *Node
	Current int
}

/*
  生成一个链表
  有两个要点
  1. 因为是单向链表，所以最早生成的应该是最后一个节点
  2. 单向链表需要使用地址，这样才可以反转
*/
func newNode() *Node {
	node5 := &Node{
		Current: 5,
		Next:    nil,
	}
	node4 := &Node{
		Current: 4,
		Next:    node5,
	}
	node3 := &Node{
		Current: 3,
		Next:    node4,
	}
	node2 := &Node{
		Current: 2,
		Next:    node3,
	}
	node1 := &Node{
		Current: 1,
		Next:    node2,
	}
	return node1
}

func revNode(n *Node) *Node {
	//记录临时变量
	var prev *Node
	//迭代
	for n != nil {
		//临时储存下一个节点
		tmp := n.Next
		//反转 将下一个节点改为指向前一个节点
		n.Next = prev
		//临时棉量储存该节点
		prev = n
		//本次迭代完成，指向下一个节点
		n = tmp
	}
	//迭代到最后 prev就是最后一个节点
	return prev
}
