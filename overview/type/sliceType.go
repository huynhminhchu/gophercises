package main

import (
	"fmt"
)

func main() {
	aSlice := []string{"chu", "dep", "trai"}
	for i, v := range aSlice {
		fmt.Println(i, ":", v)
	}
}
