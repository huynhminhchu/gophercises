package main

import "fmt"

func change(s []string) {
	s[0] = "Change_function"
}

func main() {
	a := [4]string{"zero", "one", "two", "three"}
	fmt.Println("a:", a)

	var S0 = a[0:1]
	fmt.Println(S0)
	S0[0] = "S0"
	fmt.Printf("S0 %T\n", S0)
	fmt.Printf("a %T\n", a)

	var S12 = a[1:3]
	fmt.Println(S12)
	S12[0] = "S12_0"
	S12[1] = "S12_1"

	fmt.Println("a:", a)

	change(S12)
	fmt.Println("a", a)

	fmt.Println("Cap of S0:", cap(S0), " , Length of S0: ", len(S0))
	S0 = append(S0, "four")
	S0 = append(S0, "five")
	S0 = append(S0, "six")
	S0 = append(S0, "seven")

	fmt.Println("a", a)
	change(S0)
	fmt.Println("a", a)
}
