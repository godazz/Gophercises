package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {

	fileName := flag.String("csv", "problems.csv", "CSV File that conatins quiz questions")
	timeLimit := flag.Int("limit", 30, "Time Limit of the quiz")
	suffle := flag.Bool("shuffle", false, "Shuffle the Question each run")
	flag.Parse()

	file, err := os.Open(*fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	r := csv.NewReader(file)
	Questions, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	if *suffle {
		rand.Shuffle(len(Questions), func(i, j int) {
			Questions[i], Questions[j] = Questions[j], Questions[i]
		})
	}

	var userAnswer string
	var correctAnswers int

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	for i, record := range Questions {
		fmt.Printf("%d. %s\n", i+1, string(record[0]))
		answerCh := make(chan string)
		go func() {
			fmt.Scan(&userAnswer)
			answerCh <- userAnswer
		}()

		select {
		case <-timer.C:
			fmt.Printf("There were %d Question, you got %d correctly!\n", len(Questions), correctAnswers)
			return
		case userAnswer := <-answerCh:
			if record[1] == userAnswer {
				correctAnswers++
			}
		}
	}

	fmt.Printf("There were %d Question, you got %d correctly!\n", len(Questions), correctAnswers)
}
