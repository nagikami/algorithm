package main

import (
	"fmt"
	"strconv"
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
<<<<<<< HEAD
	fmt.Println("The length of linkedList is "+ strconv.Itoa(Length(head)))
=======
	fmt.Println(Length(head))
>>>>>>> c5a573e7470a493634fb7cbd3f1e2df9a65d6070
}
