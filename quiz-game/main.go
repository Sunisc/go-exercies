package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type problem struct {
	question string
	answer   int
}

func main() {
	fmt.Println("🔖 quiz game 🎯")

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

	problems := []problem{}
	for _, record := range records {
		question := record[0]
		answer, err := strconv.Atoi(record[01])
		if err != nil {
			log.Fatalln("Solution must be an int", err)
		}
		problems = append(problems, problem{question, answer})
	}

	fmt.Println("Problems loaded successfully ...")
	fmt.Println("Quiz starting ...")

	score := 0

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

	scoreString := fmt.Sprint("You've scored: ", score, "/", len(problems))
	fmt.Println(scoreString)
}
