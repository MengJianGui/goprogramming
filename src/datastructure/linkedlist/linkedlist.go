package linkedlist

import (
	"fmt"
)

// Node 定义链表节点
type Node struct {
	Value int   // Value 定义节点值
	Next  *Node // Next 链表的指向
}

// LinkedList 定义链表
type LinkedList struct {
	HeadNode *Node // HeadNode 头节点，指针类型变量（指针变量可以为空，但是结构体不行）
}

// IsEmpty()方法，判断链表是否为空链表
func (list *LinkedList) IsEmpty() bool {
	if list.HeadNode == nil {
		return true
	}
	return false
}

// Len()方法判断链表的长度
func (list *LinkedList) Len() int {
	size := 0
	temp := list.HeadNode
	for temp != nil {
		size++
		temp = temp.Next
	}
	return size
}

// Append() 方法往链表尾部添加新节点,
// 当不传入指针变量时，对HeadNode做直接赋值是无效的，
// 可以对HeadNode结构体中的内容进行修改，因为HeadNode在LinkedList中是指针形式。
func (list *LinkedList) Append(data int) {
	temp := &Node{Value: data}
	if list.HeadNode == nil {
		list.HeadNode = temp
	} else {
		current := list.HeadNode
		for current.Next != nil {
			current = current.Next
		}
		current.Next = temp
	}
}

// Reverse() 反转链表方法：递归实现, 而且对于头节点的修改如果不传入指针变量是修改不了的。
// 当链表不是以指针形式传入时，1357为什么就只剩1了。:因为执行reverse(list.HeadNode)函数后，链表的指向已经变了，这个时候HeadNode已经指向nil了。
func (list *LinkedList) Reverse() {
	// 第一种方法：递归形式反转链表
	//list.HeadNode = reverseByRecursion(list.HeadNode)

	// 第二种方法：非递归形式反转链表
	/*
	 * 不使用递归反转链表的节点指向, 重点在于节点解耦
	 * nil<-1, 3->5->7; 第一步：1和3解耦，并且1指向nil
	 * nil<-1<-3, 5->7; 第二步：开始循环，首先解耦3和5，并使3指向1.
	 * nil<-1<-3<-5, 7;                            然后解耦5和7，并使5指向3
	 * nil<-1<-3<-5<-7;
	 */
	head := list.HeadNode
	if head == nil || head.Next == nil {
		return
	}
	current := head.Next
	head.Next = nil
	for current != nil {
		temp := current.Next
		current.Next = head
		head = current
		current = temp
	}
	list.HeadNode = head
}

/*
 * 递归反转链表的节点指向
 * 1->3->5->7;
 * null<-7;
 * null<-5<-7;
 * null<-3<-5<-7;
 * null<-1<-3<-5<-7;返回7然后在用head.next=Node(7)即可
 */
func reverseByRecursion(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	current := head.Next
	head.Next = nil
	temp := reverseByRecursion(current) // 递归要避免死循环
	current.Next = head
	return temp
}

// 方便链表输出，实现String()方法
func (list LinkedList) String() string {
	current := list.HeadNode
	if current == nil {
		return "LinkedList is empty!"
	}
	var str string
	for current != nil {
		str = fmt.Sprintf("%s	%d", str, current.Value)
		current = current.Next
	}
	return str
}
