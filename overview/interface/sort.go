package main

import (
	"fmt"
	"sort"
)

//sort.Interface interface
type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

type S1 struct {
	F1 int
	F2 string
	F3 int
}

type S2 struct {
	F1 int
	F2 string
	F3 S1
}

type S2slice []S2

func (a S2slice) Len() int {
	return len(a)
}
func (a S2slice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a S2slice) Less(i, j int) bool {
	return a[i].F3.F1 < a[j].F3.F1
}

func main() {
	data := []S2{
		S2{1, "one", S1{1, "S1_1", 10}},
		S2{2, "two", S1{2, "S1_1", 20}},
		S2{-1, "two", S1{-1, "S1_1", -20}},
	}
	fmt.Println("Before: ", data)
	sort.Sort(S2slice(data))
	fmt.Println("After: ", data)
}
