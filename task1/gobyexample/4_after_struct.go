package gobyexample

import (
	"fmt"
	"iter"
)

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func (d Weekday) String() string {
	return []string{"星期日", "星期一", "星期二", "星期三", "星期四", "星期五", "星期六"}[d]
}

func enums() {
	var today Weekday = Tuesday
	next_day := Wednesday

	fmt.Println("今天是:", today)
	fmt.Println("它的底层值是:", int(today))

	fmt.Println("明天是:", next_day)
	fmt.Println("它的底层值是:", int(next_day))
}

func PrintSlice[T any](s []T) {
	fmt.Print("[")
	for i, v := range s {
		if i > 0 {
			fmt.Print(",")
		}
		fmt.Print(v)
	}
	fmt.Println("]")
	fmt.Println(s)
}

func Generics() {
	intSlice := []int{1, 23, 4}
	PrintSlice(intSlice)

	stringSlice := []string{"欧萌", "书记", "大丽", "小埋"}
	PrintSlice(stringSlice)
}

func SliceIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

func Generics1() {
	s := []int{1, 2, 3, 4}
	fmt.Println(SliceIndex(s, 3))
}

type element[T any] struct {
	val  T
	next *element[T]
}

type List[T any] struct {
	head, tail *element[T]
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) AllElements() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func List_Test() {
	lst := List[string]{}
	lst.Push("塞西莉亚")
	lst.Push("车✌ ")
	lst.Push("lxp")
	lst.Push("ShaddockNH3")
	fmt.Println("list:", lst.AllElements())
}

func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 1, 1
		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func fibtest() {
	for n := range genFib() {
		fmt.Println(n)
		if n >= 10 {
			break
		}
	}
}

func test4() {
	// enums()
	// Generics()
	// Generics1()
	// List_Test()
	fibtest()
}

// func main() {
// 	fmt.Println("Hello World!")
// 	test4()
// }
