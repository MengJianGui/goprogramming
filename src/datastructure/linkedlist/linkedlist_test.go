package linkedlist

import (
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {
	// 初始化一个链表
	list := LinkedList{HeadNode:&Node{
		Value: 1,
		Next:  nil,
	}}
	// 链表添加元素
	list.Append(3)
	list.Append(5)
	list.Append(7)
	fmt.Println(list)

	list.Reverse()
	fmt.Println(list)
}