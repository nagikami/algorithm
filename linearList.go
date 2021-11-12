package main

import (
	"fmt"
	"strconv"
)

const (
	maxSize = 10
)

type LinearList struct {
	Data [maxSize]int
	Last int
}

func (l *LinearList) Find(x int) int {
	i := 0
	for ; i <= l.Last && l.Data[i] != x; i++ {
	}
	if (i > l.Last) {
		return -1
	} else {
		return i
	}
}

func (l *LinearList) Length() int {
	return l.Last + 1
}

func (l *LinearList) Insert(i int, x int) { 
	if (l.Last == maxSize) {
		fmt.Println("List is full!")
		return
	}
	if (i < 1 || i > l.Last + 2) {
		fmt.Println("Invalid index!")
		return
	}
	for j := l.Last; j >= i - 1; j-- {
		 l.Data[j + 1] = l.Data[j]
	}
	l.Data[i - 1] = x
	l.Last++
}

func (l *LinearList) remove(i int) {
	if (i < 1 || i > l.Last + 1) {
		fmt.Println("Invalid index!")
		return
	}
	for ; i <= l.Last; i++ {
		l.Data[i -1] = l.Data[i]
	}
	l.Last--
}

func main() {
	l := LinearList{[maxSize]int{}, 2}
	l.Data[1] = 2
	fmt.Println(l.Data)
	fmt.Println(l.Find(2))
	fmt.Println(l.Length())
	l.Insert(1, 3)
	fmt.Println(l.Data)
	fmt.Println(strconv.Itoa(l.Data[1]) + " " + strconv.Itoa(l.Data[2]))
	l.remove(2)
	fmt.Println(l.Data)
}
