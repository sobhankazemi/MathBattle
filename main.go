package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var (
	fileName          string
	totalTime         int
	score             int
	numberOfQuestions int
	lives             int
)

func init() {
	flag.StringVar(&fileName, "file", "problems.csv", "questions starting file")
	flag.IntVar(&totalTime, "time", 30, "total time for answering the questions")
	flag.IntVar(&lives, "lives", 2, "number of lives in the game")
	flag.Parse()
}

func main() {
	fmt.Println("press any key to start")
	var startPhrase string
	fmt.Scanln(&startPhrase)
	timer := time.NewTimer(time.Second * time.Duration(totalTime))
	go func() {
		<-timer.C
		fmt.Printf("you scored %d out of %d\n", score, numberOfQuestions)
		os.Exit(0)
	}()
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("file does not exist")
	}
	defer file.Close()
	problems := csv.NewReader(file)
	records, err := problems.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	numberOfQuestions = len(records)
	finished := true
	for _, record := range records {
		check := true
		fmt.Printf("what is %s sir ?\n", record[0])
		var answer string
		fmt.Scan(&answer)
		if strings.Compare(answer, record[1]) != 0 {
			lives--
			if lives == 0 {
				fmt.Printf("you scored %d out of %d\n", score, numberOfQuestions)
				finished = false
				break
			}
			check = false
		}
		if check {
			score++
		}
	}
	if finished {
		fmt.Println("congrats you answered all questions")
		fmt.Printf("total number of questions was %d\n", numberOfQuestions)
	}
}
