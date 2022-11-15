package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("only 1 argument")
		return
	}
	fileName := arguments[1]
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		return
	}
	r := csv.NewReader(file)
	records, _ := r.ReadAll()

	totalQuestion := 0
	totalCorrectAns := 0
	//fmt.Println(records)
	for {
		fmt.Println("Enter s to start the quiz game: ")
		var start string
		fmt.Scan(&start)
		if strings.Compare(start, "s") == 0 {
			break
		}
	}
	timer := time.NewTimer(30 * time.Second)
	go func() {
		<-timer.C
		fmt.Println("Time out 30")
	}()
	for _, v := range records {
		totalQuestion++
		temp, _ := strconv.Atoi(v[1])
		// re := regexp.MustCompile(`(?P<first>\d+)\+(?P<last>\d+)`)
		// subStr := re.FindStringSubmatch(v[0])
		// firstNum, _ := strconv.Atoi(subStr[1])
		// secondNum, _ := strconv.Atoi(subStr[2])
		// fmt.Println(firstNum, "+", secondNum)
		fmt.Println(v[0])
		fmt.Println("Please input your answer:")
		var answer string
		fmt.Scan(&answer)
		temp2, _ := strconv.Atoi(answer)
		if temp == temp2 {
			totalCorrectAns++
		}
	}
	fmt.Println("Total questions: ", totalQuestion)
	fmt.Println("Total correct answers: ", totalCorrectAns)
}
