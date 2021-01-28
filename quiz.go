package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

type quizItem struct {
	Question string
	Answer   string
}

func main() {
	csvFile, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	records, err := reader.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(records); i++ {

		quizItem := quizItem{records[i][0], records[i][1]}
		fmt.Println("Questions:", quizItem.Question)
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter your answer now:")
		text, _ := reader.ReadString('\n')
		fmt.Println("Your answer is :",text)
		text = strings.TrimSuffix(text, "\n")
		if text == quizItem.Answer {
			fmt.Println("Correct!")
		} else {
			fmt.Println("Wrong the answer is:", quizItem.Answer)
		}
		
	}
}
