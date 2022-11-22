package main

import "fmt"

type aStructure struct {
	field1 complex128
	field2 int
}

func processPointer(x *float64) {
	*x = *x * *x
}

func returnPointer(x float64) *float64 {
	temp := x * 2
	return &temp
}

func bothPointer(x *float64) *float64 {
	temp := *x * 2
	return &temp
}

func main() {
	var f float64 = 12.123
	fP := &f
	fmt.Println("Memory address of f: ", fP)
	fmt.Println("value of f : ", *fP)

	processPointer(fP)
	fmt.Println("Memory address of f: ", fP)
	fmt.Printf("value of f : %0.2f\n", *fP)

	x := returnPointer(f)
	fmt.Printf("Value of x: %.2f\n", *x)

}
