package main

import "fmt"

func main() {
	aMap := map[string]int{}
	aMap["test"] = 1

	fmt.Println(aMap)
	aMap = nil

	fmt.Println(aMap)
	if aMap == nil {
		fmt.Println("nil map!")
		aMap = map[string]int{}
	}
	fmt.Println(aMap)

	myMap := map[string]string{"chu": "dep"}

	for i, v := range myMap {
		fmt.Print("i:", i, " v:", v)
	}

}
