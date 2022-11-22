package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var data []phoneEntry

func random(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func randomStr(len int) string {
	temp := ""
	for i := 0; i < len; i++ {
		temp += string(byte(random(65, 90)))
	}
	return temp
}

type phoneEntry struct {
	Name    string
	Surname string
	Tel     string
}

func search(key string) *phoneEntry {
	for i := range data {
		if data[i].Tel == key {
			return &data[i]
		}
	}
	return nil
}
func list() {
	for i := range data {
		fmt.Printf("%d.Name: %s\n", i+1, data[i].Name)
		fmt.Printf("Surname: %s\n", data[i].Surname)
		fmt.Printf("Tel: %s\n", data[i].Tel)
	}
}
func help() {
	fmt.Println("Usage: ./phoneBook search|list <arguments.")
}
func main() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 100; i++ {
		randomTel := string(byte(random(48, 57))) + string(byte(random(48, 57))) + string(byte(random(48, 57)))
		temp := phoneEntry{randomStr(4), randomStr(5), string(randomTel)}
		data = append(data, temp)
	}
	arguments := os.Args

	if len(arguments) < 2 {
		help()
		return
	}
	action := arguments[1]
	if action == "search" {
		if len(arguments) < 3 {
			help()
			return
		}
		t := search(arguments[2])
		if t != nil {
			fmt.Println("Found ", *t)
		} else {
			fmt.Println("Not found")
		}
	} else if action == "list" {
		list()
	}
}
