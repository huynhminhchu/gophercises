package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var data []phoneEntry
var index map[string]int

type phoneEntry struct {
	Name       string
	Surname    string
	Tel        string
	LastAccess string
}

func initEntry(N, S, T, L string) *phoneEntry {
	temp := phoneEntry{N, S, T, L}
	return &temp
}

func search(key string) *phoneEntry {
	i, ok := index[key]
	if !ok {
		return nil
	}
	return &data[i]
}

func createIndex() {
	index = make(map[string]int)
	for i, v := range data {
		fmt.Println(v.Tel)
		fmt.Println(i)
		index[v.Tel] = i
	}
}

func list() {
	for i := range data {
		fmt.Printf("%d.Name: %s\n", i+1, data[i].Name)
		fmt.Printf("Surname: %s\n", data[i].Surname)
		fmt.Printf("Tel: %s\n", data[i].Tel)
	}
}

func insert() {

	return
}
func delete() {
	return
}
func validateTel(tel string) bool {
	return true
}

func checkCSVFile(CSVFILE string) {
	_, err := os.Stat(CSVFILE)

	if os.IsNotExist(err) {
		fmt.Println("Creating new file ...")
		f, err := os.Create(CSVFILE)
		if err != nil {
			fmt.Println("Error creating new csv file. Exit")
			return
		}
		f.Close()
	}

	fileInfo, err := os.Stat(CSVFILE)
	if err == nil {
		mode := fileInfo.Mode()
		if !mode.IsRegular() {
			fmt.Println(CSVFILE, " is not a regular file. Exiting")
			return
		}
	}
}
func readCSVFile(CSVFILE string) {
	file, err := os.Open(CSVFILE)
	if err != nil {
		fmt.Println("Err open csv file: ", err)
		return
	}
	r := csv.NewReader(file)
	lines, _ := r.ReadAll()

	for _, line := range lines {
		data = append(data, phoneEntry{
			Name:       line[0],
			Surname:    line[1],
			Tel:        line[2],
			LastAccess: line[3],
		})
	}
}
func saveCSVFile(CSVFILE string) {
	file, err := os.Open(CSVFILE)
	if err != nil {
		fmt.Println("Err open csv file: ", err)
		return
	}

}

func help() {
	fmt.Println("Usage: ./phoneBook search|list|delete|insert <arguments.")
}
func main() {
	rand.Seed(time.Now().UnixNano())
	CSVFILE := "./data.csv"

	checkCSVFile(CSVFILE)
	readCSVFile(CSVFILE)
	fmt.Println("Data: ", data)

	createIndex()
	fmt.Println("Map index: ", index)

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
