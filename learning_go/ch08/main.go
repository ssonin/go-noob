package main

import (
	"errors"
	"fmt"
)

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint64 | uintptr |
		~float32 | ~float64
}

func Double[T Number](n T) T {
	return n + n
}

type Printable interface {
	~int | ~float64
	fmt.Stringer
}

type MyPrintableInt int

func (mpi MyPrintableInt) String() string {
	return fmt.Sprintf("%d", mpi)
}

type MyPrintableFloat64 float64

func (mpf MyPrintableFloat64) String() string {
	return fmt.Sprintf("%f", mpf)
}

func PrintPrintable[T Printable](t T) {
	fmt.Println(t)
}

// ex3

type Node[T any] struct {
	Val  T
	Next *Node[T]
}

func (n *Node[T]) Add(val T) *Node[T] {
	if n == nil {
		return &Node[T]{Val: val}
	}
	curr := n
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = &Node[T]{Val: val}
	return n
}

func (n *Node[T]) Insert(val T, index int) *Node[T] {
	if n == nil {
		return &Node[T]{Val: val}
	}
	var prev *Node[T]
	curr := n
	for i := 0; i < index; i++ {
		if curr == nil {
			break
		}
		prev = curr
		curr = curr.Next
	}
	if prev == nil {
		return &Node[T]{Val: val, Next: curr}
	} else {
		prev.Next = &Node[T]{Val: val, Next: curr}
		return n
	}
}

type List[T comparable] struct {
	Head *Node[T]
}

func NewList[T comparable]() *List[T] {
	return &List[T]{}
}

func (l *List[T]) Add(val T) {
	l.Head = l.Head.Add(val)
}

func (l *List[T]) Insert(val T, index int) (*List[T], error) {
	if index < 0 {
		return l, errors.New("Negative index")
	}
	l.Head = l.Head.Insert(val, index)
	return l, nil
}

func (l *List[T]) Index(val T) int {
	var index int
	for curr := l.Head; curr != nil; curr = curr.Next {
		if curr.Val == val {
			return index
		}
		index++
	}
	return -1
}

func (l List[T]) Println() {
	curr := l.Head
	for curr != nil {
		fmt.Printf("%v -> ", curr.Val)
		curr = curr.Next
	}
	fmt.Println("nil")
}

func main() {
	// ex1
	fmt.Println(Double(1.15))
	// ex2
	mpi := MyPrintableInt(42)
	PrintPrintable(mpi)
	mpf := MyPrintableFloat64(42.0)
	PrintPrintable(mpf)
	// ex3
	l := NewList[int]()
	l.Println()
	l.Add(42)
	l.Add(4)
	l.Add(2)
	l.Add(1)
	l.Println()
	l.Insert(146, 4)
	l.Println()
}
