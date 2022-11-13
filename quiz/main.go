package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("only 1 argument")
		return
	} else {
		for i:=1;i<len(arguments);i++ {
			fmt.Println("index:", i, "value: ",arguments[i])
		}
	}
	
}