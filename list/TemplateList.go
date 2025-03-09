package main

import "fmt"

// Node 代表链表中的一个节点，包含数据和指向下一个节点的指针
type Node[T comparable] struct {
	value T
	next  *Node[T]
}

// LinkedList 是一个泛型链表，支持任意类型的元素
type LinkedList[T comparable] struct {
	head *Node[T]
}

// NewLinkedList 创建一个空链表
func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// Append 向链表尾部添加元素
func (ll *LinkedList[T]) Append(value T) {
	newNode := &Node[T]{value: value}
	if ll.head == nil {
		ll.head = newNode
		return
	}

	// 找到链表的最后一个节点
	current := ll.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

// Prepend 向链表头部添加元素
func (ll *LinkedList[T]) Prepend(value T) {
	newNode := &Node[T]{value: value}
	newNode.next = ll.head
	ll.head = newNode
}

// Remove 删除链表中第一个匹配的值
func (ll *LinkedList[T]) Remove(value T) {
	if ll.head == nil {
		return
	}

	// 如果头节点就是要删除的元素
	if ll.head.value == value {
		ll.head = ll.head.next
		return
	}

	// 查找要删除的节点
	current := ll.head
	for current.next != nil && current.next.value != value {
		current = current.next
	}

	// 如果找到了匹配的节点
	if current.next != nil {
		current.next = current.next.next
	}
}

// Print 打印链表中的所有元素
func (ll *LinkedList[T]) Print() {
	current := ll.head
	for current != nil {
		fmt.Print(current.value, " ")
		current = current.next
	}
	fmt.Println()
}

func main() {
	// 创建一个整数类型的链表
	intList := NewLinkedList[int]()
	intList.Append(10)
	intList.Append(20)
	intList.Append(30)
	intList.Prepend(5)

	fmt.Print("Integer List: ")
	intList.Print() // 输出: 5 10 20 30

	intList.Remove(20)
	fmt.Print("After removing 20: ")
	intList.Print() // 输出: 5 10 30

	// 创建一个字符串类型的链表
	stringList := NewLinkedList[string]()
	stringList.Append("apple")
	stringList.Append("banana")
	stringList.Append("cherry")

	fmt.Print("String List: ")
	stringList.Print() // 输出: apple banana cherry
}
