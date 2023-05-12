package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {

	fileName := flag.String("csv", "problems.csv", "CSV File that conatins quiz questions")
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

	var userAnswer string
	var correctAnswers int

	for i, record := range Questions { // record[0] -> Question, record[1] -> Correct Answer
		fmt.Printf("%d. %s\n", i+1, string(record[0]))
		fmt.Scan(&userAnswer)

		if err != nil {
			log.Fatal(err)
		}

		if record[1] == userAnswer {
			correctAnswers++
		}
	}

	fmt.Printf("There were %d Question, you got %d correctly!\n", len(Questions), correctAnswers)

}
