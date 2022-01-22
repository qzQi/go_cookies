package forwardlist

import (
	"errors"
	"fmt"
)

// node internal used
type node struct {
	val  interface{}
	next *node
}

// newNode internal use
func newNode(val interface{}) *node {
	return &node{val, nil}
}

type Singly struct {
	Length int
	// 这个head最好是dummy head
	Head *node
}

func NewSingly() *Singly {
	return &Singly{0, newNode(nil)}
}

func (li *Singly) Print() {
	for cur := li.Head; cur != nil; cur = cur.next {
		fmt.Printf("%#v ", cur.val)
		// fmt.Println()
	}
	fmt.Println()
}

func (l *Singly) CheckRange(left, right int) error {
	if left > right {
		return errors.New("left must no greater than right")
	} else if left < 1 {
		return errors.New("left must no less than 1")
	} else if right > l.Length {
		return errors.New("right must no greater than Length")
	}
	return nil
}

func reverse(n *node) *node {
	if n == nil || n.next == nil {
		return n
	}
	newNode := reverse(n.next)
	n.next.next = n
	n.next = nil
	return newNode
}
func (l *Singly) Reverse() {
	l.Head = reverse(l.Head.next)
}

func (l *Singly) AddAtBeg(val interface{}) {
	n := newNode(val)
	n.next = l.Head.next
	l.Head.next = n
}

func (l *Singly) AddAtEnd(val interface{}) {
	cur := l.Head
	for ; cur.next != nil; cur = cur.next {
	}

	cur.next = newNode(val)
}
