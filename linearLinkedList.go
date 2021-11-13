package main

import (
	"fmt"
)

type LinearLinkedList struct {
	Data int
	Next *LinearLinkedList
}

func Length(l *LinearLinkedList) int {
	p := l
	i := 0
	for ; p != nil; p = p.Next {
		i++
	}
	return i
}
	
func main() {
	p2 := LinearLinkedList{10, nil}
	p1 := LinearLinkedList{5, &p2}
	head := &p1
	fmt.Println(Length(head))
}
