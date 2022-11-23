package main

import "fmt"

func TwoArrayToSlice(firstArray, secondArray []string) {
	return
}

func testSlice(slice []string) {
	fmt.Println(slice)
}
func testArray(array [100]string) {
	fmt.Println(array)
}
func main() {
	var aSlice = []string{}
	var anArray = [100]string{"one", "two", "three", "four"}

	fmt.Printf("Slice type: %T\n", aSlice)
	fmt.Printf("Array type: %T\n", anArray)
	testSlice(aSlice)
	testArray(aSlice)
}
