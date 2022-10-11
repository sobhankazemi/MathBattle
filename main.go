package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var fileName string

func init() {
	flag.StringVar(&fileName, "file", "problems.csv", "questions starting file")
	flag.Parse()
}

func main() {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("file does not exists")
	}
	defer file.Close()
	problems := csv.NewReader(file)
	records, err := problems.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	numberOfQuestions := len(records)
	finished := true
	var score int
	for _, record := range records {
		fmt.Printf("what is %s sir ?\n", record[0])
		var answer string
		fmt.Scan(&answer)
		if strings.Compare(answer, record[1]) != 0 {
			fmt.Printf("you answered %d out of %d questions\n", score, numberOfQuestions)
			finished = false
			break
		}
		score++
	}
	if finished {
		fmt.Println("congrats you answered all questions")
		fmt.Printf("total number of questions was %d\n", numberOfQuestions)
	}
}
