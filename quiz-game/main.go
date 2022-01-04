package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'Q&A'")
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
	counter := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		var ans string
		fmt.Scanf("%s\n", &ans)
		if ans == p.a {
			fmt.Println("Correct!")
			counter++
		}else{
			fmt.Println("OOPS, you're wrong!!! Try again")
			fmt.Scanf("%s\n", &ans)
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
