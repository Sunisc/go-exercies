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
	csvFilePtr := flag.String("csv", "problems.csv", "csv to read problems from")
	flag.Parse()

	records := readFile(csvFilePtr)
	problems := loadProblems(records)
	startQuiz(problems, *timePtr)

}

func startQuiz(problems []problem, duration int) {
	score := 0
	timer := time.NewTimer(time.Duration(duration) * time.Second)

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.question)
		answerCh := make(chan int)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answer = strings.Trim(answer, " ")
			answerInt, _ := strconv.Atoi(answer)
			answerCh <- answerInt
		}()
		select {
		case <-timer.C:
			fmt.Printf("You scored %d of %d", score, len(problems))
			return
		case answer := <-answerCh:
			if answer == p.answer {
				score++
			}
		}
	}
	fmt.Printf("You scored %d out of %d.\n", score, len(problems))
}

func loadProblems(records [][]string) []problem {
	fmt.Println("Loading Problems...")

	problems := make([]problem, len(records))
	for i, record := range records {
		question := record[0]
		answer, err := strconv.Atoi(record[1])
		if err != nil {
			log.Fatalln("Solution must be an int", err)
		}
		problems[i] = problem{question, answer}
	}
	return problems
}

func readFile(csvFileName *string) [][]string {
	fmt.Println("Reading File...")

	file, err := os.Open(*csvFileName)
	if err != nil {
		log.Fatalln(fmt.Sprintf("Failed to open file: %s\n", *csvFileName), err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalln("Error reading csv", err)
	}
	return records
}
