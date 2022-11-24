package main

import "fmt"

type Entry struct {
	Name    string
	Surname string
	Tel     string
}

func initS(N, S string, T string) *Entry {
	return &Entry{N, S, T}
}

func main() {
	S := Entry{Name: "chu", Surname: "minh", Tel: "01234567"}
	S2 := new(Entry)

	fmt.Println(S)
	fmt.Println(*S2)

	S3 := initS("chu", "dep", "trai")
	fmt.Println(*S3)
}
