package main

import (
	"fmt"
	"math/rand"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func randomStr(len int) string {
	temp := ""
	for i := 0; i < len; i++ {
		temp += string(byte(random(33, 126)))
	}
	return temp
}

func main() {
	rand.Seed(time.Now().UnixNano())
	test := randomStr(20)
	fmt.Println(test)
}
