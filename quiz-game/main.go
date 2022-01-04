package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'Q&A'")
	timeLimit := flag.Int("limit", 30, "The time limit for the quiz in seconds")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s", *csvFilename))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()

	if err != nil {
		exit("Failed to parse the provided CSV file")
	}
	problems := parseLines(lines)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	counter := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		answerCh := make(chan string)
		go func(){
			var ans string
			fmt.Scanf("%s\n", &ans)
			answerCh <- ans
		}()
		select {
		case <-timer.C:
			fmt.Printf("\nYou scored %v out of %d", counter, len(problems))
			return
		case ans := <-answerCh:
			if ans == p.a {
				fmt.Println("Correct!")
				counter++
			}
		}
	}
	fmt.Printf("You scored %v out of %d", counter, len(problems))
}

type problem struct {
	q string
	a string
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
