package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type problem struct {
	question string
	answer   int
}

func main() {
	fmt.Println("🔖 quiz game 🎯")

	timePtr := flag.Int("time-limit", 30, "set the time limit for the quiz. default is 30")
	flag.Parse()
	records := readFile()
	problems := loadProblems(records)
	fmt.Println("Ready. Press any key to start")
	var input string
	fmt.Scan("%s", &input)
	if input != "" {
		startQuiz(problems, *timePtr)
	}
}

func startQuiz(problems []problem, duration int) {
	score := 0
	fmt.Println("Starting Quiz:", duration, "seconds left")
	f := func() {
		fmt.Println("Times Up!\nYou've scored", score, "/", len(problems))
		os.Exit(0)
	}
	timer := time.AfterFunc(time.Duration(duration)*time.Second, f)
	for _, problem := range problems {
		fmt.Print("\n", problem.question, " = ?\n")
		fmt.Println("Enter your answer: ")

		var input string
		var answer int
		fmt.Scanln(&input)

		if input == "" {
			fmt.Println("Too hard? Try the next one!")
		} else {
			// clean up input
			var err error
			input = strings.Trim(input, " ")
			answer, err = strconv.Atoi(input)
			if err != nil {
				fmt.Println("Only numeric answers are accepted, try again on the next one!")
			}
		}

		if answer == problem.answer {
			fmt.Println("Correct!")
			score += 1
		}
	}

	scoreMessage := fmt.Sprint("You've scored: ", score, "/", len(problems))
	fmt.Println(scoreMessage)
	<-timer.C
}

func loadProblems(records [][]string) []problem {
	fmt.Println("Loading Problems...")

	problems := []problem{}
	for _, record := range records {
		question := record[0]
		answer, err := strconv.Atoi(record[1])
		if err != nil {
			log.Fatalln("Solution must be an int", err)
		}
		problems = append(problems, problem{question, answer})
	}
	return problems
}

func readFile() [][]string {
	fmt.Println("Reading File...")

	file, err := os.Open("problems.csv")
	if err != nil {
		log.Fatalln("Error opening file", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalln("Error reading csv", err)
	}
	return records
}
