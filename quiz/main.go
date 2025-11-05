package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type problem struct {
	q string
	a string
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer'")
	timeInSeconds := flag.Int("limit", 30, "time limit for the quiz")
	flag.Parse()
	file, err := os.Open(*csvFilename)

	if err != nil {
		exit(fmt.Sprintf("Failed to open the csv file: %s\n", *csvFilename))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()

	if err != nil {
		exit("Failed to parse the CSV file provided.")
	}
	problems := parseLines(lines)

	correct := 0

	timer := time.NewTimer(time.Duration(*timeInSeconds) * time.Second)
	// problemLoop:
	for i, problem := range problems {

		fmt.Printf("Problem %d: %s= ", i+1, problem.q)
		answerCh := make(chan string)

		go func() {
			var ans string
			fmt.Scanf("%s ", &ans)
			answerCh <- ans
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nYour correct answers were %d out of %d questions.\n", correct, len(problems))
			// break problemLoop // this is one way to break out of loop
			return
		case ans := <-answerCh:
			if ans == problem.a {
				fmt.Println("Correct!")
				correct++
			}
		}
	}

	fmt.Printf("Your correct answers were %d out of %d questions.\n", correct, len(problems))

}

func parseLines(lines [][]string) []problem {
	problems := make([]problem, len(lines))
	for i, line := range lines {
		problems[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}

	return problems
}

func exit(message string) {
	fmt.Println(message)
	os.Exit(1)
}
