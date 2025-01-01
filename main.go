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
	filename := flag.String("csv", "problems.csv", "a csv file in a format of 'question','answer'")
	flag.Parse()
	timelimit := flag.Int("timelimit", 30, "this is used for time limit in quiz")
	file, err := os.Open(*filename)
	if err != nil {
		exit(fmt.Sprintf("failed to open a file %s ", err))
	}
	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("failed to parsed an csv file %s", err))
	}
	pro := parseProblems(data)
	correct := 0
	timer := time.NewTimer(time.Duration(*timelimit) * time.Second)
	for i, p := range pro {
		select {
		case <-timer.C:
			fmt.Printf("you scored %d of %d problems\n", correct, len(pro))
			return
		default:
			fmt.Printf("problem #%d : %s \n", i, p.q)
			var ans string
			fmt.Scanf("%s\n", &ans)
			if ans == p.a {
				correct++
			}
		}
	}
	fmt.Printf("you scored %d of %d problems\n", correct, len(pro))
}

func parseProblems(p [][]string) []problem {
	var problems []problem
	for _, l := range p {
		problems = append(problems, problem{
			q: l[0],
			a: strings.TrimSpace(l[1]),
		})
	}
	return problems
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
