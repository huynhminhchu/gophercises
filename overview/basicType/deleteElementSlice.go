package main

import (
	"fmt"
	"os"
	"strconv"
)
kubeadm join --token e9wfn6.9z5zri91vt6t7lko 10.100.6.48:6443 --discovery-token-ca-cert-hash sha256:99a3c4913fc6dc2cf4a3e73481721d1a551620a320e1899de8281574476d1219 --ignore-preflight-errors=All
// Use this function if u care about order in a slice

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	args := os.Args
	if len(args) == 1 {
		fmt.Println("Need argument")
		return
	}
	index, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	if index > len(a)-1 || index < 0 {
		fmt.Print("Element index need to be in the slice range")
		return
	}

	a1 := append(a[:index], a[index+1:]...)
	fmt.Println("Array after deletion 1: ", a1)

	//Not care about slice order
	a2 := a
	a2[index] = a2[len(a)-1]
	a2 = a2[:len(a2)-1]

	fmt.Println("Array after deletion 2: ", a2)
}
