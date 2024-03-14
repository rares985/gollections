package gollections

import "fmt"

type ListNode[T comparable] struct {
	Val  T
	Next *ListNode[T]
}

type List[T comparable] struct {
	head *ListNode[T]
}

func (this *List[T]) IsEmpty() bool {
	return this == nil || this.head == nil
}

func (this *List[T]) Length() int {
	if this == nil {
		return 0
	}

	n := 0
	for it := this.head; it != nil; it = it.Next {
		n += 1
	}
	return n
}

func (this *List[T]) Front() *ListNode[T] {
	if this == nil {
		return nil
	}
	return this.head
}

func (this *List[T]) Back() *ListNode[T] {
	if this == nil {
		return nil
	}

	it := this.head

	for it != nil && it.Next != nil {
		it = it.Next
	}

	return it
}

func (this *List[T]) HasCycle() bool {
	if this == nil {
		return false
	}

	slow := this.head
	fast := this.head

	for slow != nil && fast != nil {
		slow = slow.Next

		if fast.Next == nil {
			return false
		}

		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}

func (this *List[T]) PushFront(data T) {
	if this == nil {
		return
	}
	newListNode := &ListNode[T]{Val: data, Next: this.head}
	this.head = newListNode
}

func (this *List[T]) PushBack(data T) {
	if this == nil {
		return
	}
	newListNode := &ListNode[T]{Val: data, Next: nil}
	if this.head == nil {
		this.head = newListNode
	} else {
		back := this.Back()
		back.Next = newListNode
	}
}

func (this *List[T]) PopFront() (*T, error) {
	if this == nil || this.head == nil {
		return nil, fmt.Errorf("can not pop empty list")
	}

	val := this.head.Val
	this.head = this.head.Next

	return &val, nil
}

func (this *List[T]) PopBack() (*T, error) {
	if this == nil || this.head == nil {
		return nil, fmt.Errorf("can not pop empty list")
	}

	if this.head.Next == nil {
		val := this.head.Val
		this.head = nil
		return &val, nil
	}

	it := this.head
	var prev *ListNode[T]

	for it.Next != nil {
		prev = it
		it = it.Next
	}

	val := it.Val
	prev.Next = nil
	return &val, nil
}

func (this *List[T]) String() string {
	result := ""
	if this == nil {
		return "nil"
	}

	for it := this.head; it != nil; it = it.Next {
		result += fmt.Sprintf("%v", it.Val)
		result += "->"
	}
	result += "nil"
	return result
}

func (this *List[T]) Equals(other *List[T]) bool {
	if this == nil && other == nil {
		return true
	}

	if this != nil && other == nil || this == nil && other != nil {
		return false
	}

	if this.head == nil && other.head == nil {
		return true
	}

	if this.head != nil && other.head == nil || this.head == nil && other.head != nil {
		return false
	}

	it1 := this.head
	it2 := other.head

	for it1 != nil && it2 != nil {
		if it1.Val != it2.Val {
			return false
		}
		it1 = it1.Next
		it2 = it2.Next
	}

	return true
}

func (this *List[T]) Reverse() {
	if this == nil || this.head == nil {
		return
	}

	it := this.head
	var prev, next *ListNode[T]

	for it != nil {
		next = it.Next
		it.Next = prev
		prev = it
		it = next
	}

	this.head = prev
}
