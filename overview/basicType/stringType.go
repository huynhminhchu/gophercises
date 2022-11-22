package main

import (
	"fmt"
)

func main() {
	r := `€`
	fmt.Printf("%T", r)
	fmt.Println(r)
}
