package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type problem struct {
	q string
	a string
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer'")

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

	for i, problem := range problems {
		fmt.Printf("Problem %d: %s= \n", i+1, problem.q)
		var ans string
		fmt.Scanf("%s ", &ans)
		if ans == problem.a {
			fmt.Println("Correct!")
			correct++
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
