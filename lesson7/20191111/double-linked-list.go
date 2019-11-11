package main

import (
	"fmt"
	"sync"
)

// 节点数据
type DoubleLinked interface{}

// 双向链表节点
type DoubleLinkedNode struct {
	Data DoubleLinked
	Prev *DoubleLinkedNode
	Next *DoubleLinkedNode
}

type DoubleLinkedList struct {
	mutex *sync.RWMutex
	Size  uint
	Head  *DoubleLinkedNode
	Tail  *DoubleLinkedNode
}

// 双链表初始化
func (list *DoubleLinkedList) Init() {
	list.mutex = new(sync.RWMutex)
	list.Size = 0
	list.Head = nil
	list.Tail = nil
}

// Get 获取指定位置的节点
func (list *DoubleLinkedList) Get(index uint) *DoubleNode {
	if list.Size == 0 || index > list.Size-1 {
		return nil
	}
	if index == 0 {
		return list.Head
	}
	node := list.Head
	var i uint
	for i = 1; i <= index; i++ {
		node = node.Next
	}
	return node
}

// Append 向双链表后面追加节点
func (list *DoubleLinkedList) Append(node *DoubleNode) bool {
	if node == nil {
		return false
	}
	list.mutex.Lock()
	defer list.mutex.Unlock()
	if list.Size == 0 {
		list.Head = node
		list.Tail = node
		node.Next = nil
		node.Prev = nil
	} else {
		node.Prev = list.Tail
		node.Next = nil
		list.Tail.Next = node
		list.Tail = node
	}
	list.Size++
	return true
}

// Insert 向双链表指定位置插入节点
func (list *DoubleLinkedList) Insert(index uint, node *DoubleNode) bool {
	if index > list.Size || node == nil {
		return false
	}

	if index == list.Size {
		return list.Append(node)
	}

	list.mutex.Lock()
	defer list.mutex.Unlock()
	if index == 0 {
		node.Next = list.Head
		list.Head = node
		list.Head.Prev = nil
		list.Size++
		return true
	}

	nextNode := list.Get(index)
	node.Prev = nextNode.Prev
	node.Next = nextNode
	nextNode.Prev.Next = node
	nextNode.Prev = node
	list.Size++
	return true
}

// Delete 删除指定位置的节点
func (list *DoubleLinkedList) Delete(index uint) bool {
	if index > list.Size-1 {
		return false
	}

	list.mutex.Lock()
	defer list.mutex.Unlock()
	if index == 0 {
		if list.Size == 1 {
			list.Head = nil
			list.Tail = nil
		} else {
			list.Head.Next.Prev = nil
			list.Head = list.Head.Next
		}
		list.Size--
		return true
	}
	if index == list.Size-1 {
		list.Tail.Prev.Next = nil
		list.Tail = list.Tail.Prev
		list.Size--
		return true
	}

	node := list.Get(index)
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	list.Size--
	return true
}

// Display 打印双链表信息
func (list *DoubleLinkedList) Display() {
	if list == nil || list.Size == 0 {
		fmt.Println("this double list is nil or empty")
		return
	}
	list.mutex.RLock()
	defer list.mutex.RUnlock()
	fmt.Printf("this double list size is %d \n", list.Size)
	ptr := list.Head
	for ptr != nil {
		fmt.Printf("data is %v\n", ptr.Data)
		ptr = ptr.Next
	}
}

// Reverse 倒序打印双链表信息
func (list *DoubleLinkedList) Reverse() {
	if list == nil || list.Size == 0 {
		fmt.Println("this double list is nil or empty")
		return
	}
	list.mutex.RLock()
	defer list.mutex.RUnlock()
	fmt.Printf("this double list size is %d \n", list.Size)
	ptr := list.Tail
	for ptr != nil {
		fmt.Printf("data is %v\n", ptr.Data)
		ptr = ptr.Prev
	}
}
