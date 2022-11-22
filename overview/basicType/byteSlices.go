//Convert byte slice to string

package main

import "fmt"

func main() {
	b := make([]byte, 12)
	fmt.Println("Byte slice: ", b)

	b = []byte("Byte slice €")
	fmt.Println("Byte slice: ", b)

	fmt.Println("Byte slice as text: \n", string(b))

}
