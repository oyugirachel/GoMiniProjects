package quiz

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type problem struct {
	question string
	answer   string
}

func main() {
	filename := flag.String("csv", "problems.csv", "csv filename with format 'question,answer' (default problems.csv)")
	limit := flag.Int("limit", 30, "limit time for the quiz in seconds (default 30s)")
	flag.Parse()

	f, err := os.Open(*filename)
	if err != nil {
		log.Printf("Error opening csv filename: %s\n", *filename)
		os.Exit(1)
	}

	r := csv.NewReader(f)
	rows, err := r.ReadAll()
	if err != nil {
		log.Printf("Error parsing csv filename: %s\n", err)
		os.Exit(1)
	}

	problems := parseCSV(rows)
	timer := time.NewTimer(time.Duration(*limit) * time.Second)
	correct := 0
	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, problem.question)
		answerChannel := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s", &answer)
			answerChannel <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\n%d answers correct of %d questions.\n", correct, len(problems))
			return
		case answer := <-answerChannel:
			if answer == problem.answer {
				correct++
			}
		}
	}

	fmt.Printf("%d answers correct of %d questions.\n", correct, len(problems))
}

func parseCSV(rows [][]string) []problem {
	problems := make([]problem, len(rows))

	for i, row := range rows {
		problems[i] = problem{
			question: row[0],
			answer:   strings.TrimSpace(row[1]),
		}
	}

	return problems
}