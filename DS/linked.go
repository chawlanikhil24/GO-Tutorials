package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Node struct {
	data int
	next *Node
}

type linkedlist struct {
	head *Node
}

func (ob *linkedlist) add(item int) {
	node := &Node{data: item}
	if ob.head == nil {
		ob.head = node
	} else {
		node.next = ob.head
		ob.head = node
	}
}

func (ob *linkedlist) show() {
	s := "head->"
	p := ob.head
	if p != nil {
		for p.next != nil {
			s += strconv.Itoa(p.data) + "->"
			p = p.next
		}
		s += strconv.Itoa(p.data) + "->"
	}
	s += "nil"
	fmt.Println(s)
	// return s
}

func main() {
	fmt.Println("linked List")
	l := linkedlist{}
	l.show()
	for i := 1; i <= 100000; i++ {
		l.add(rand.Intn(1000))
		// fmt.Print(i)
	}
	// l.show()
}
