package main

import (
	"fmt"
)

type LinearLinkedList struct {
	Data int
	Next *LinearLinkedList
}

func (l *LinearLinkedList) Length() int {
	p := l.Next
	i := 0
	for ; p != nil; p = p.Next {
		i++
	}
	return i
}
	
func main() {
	p2 := LinearLinkedList{10, nil}
	p1 := LinearLinkedList{5, &p2}
	l := LinearLinkedList{0, &p1}
	fmt.Println(l.Length())
}
