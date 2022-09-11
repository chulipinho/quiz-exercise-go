package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	defaultPath := "./data/problems.csv"
	usage := "Specifies custom file for questions."
	filePath := flag.String("p", defaultPath, usage)
	timerDuration := flag.Int("t", 30, "Timer duration in seconds")

	flag.Parse()

	problems := openFile(*filePath)
	timer := time.NewTimer(time.Duration(*timerDuration) * time.Second)

	answerCh := make(chan string)
	score := 0

	for i, problem := range problems {
		fmt.Printf("Question #%d - %s = ", i, problem.question)

		go func() {
			var a string
			fmt.Scanf("%s", &a)
			answerCh <- a
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nYou got %d out of %d\n", score, len(problems))
			return
		case a := <-answerCh:
			if a == problem.answer {
				score++
			}
		}
	}
	fmt.Printf("\nYou got %d out of %d\n", score, len(problems))
	return
}

type problem struct {
	question string
	answer   string
}

func parseProblems(file *csv.Reader) []problem {
	p, err := file.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	ret := make([]problem, len(p))
	for i, v := range p {
		ret[i] = problem{
			question: v[0],
			answer:   strings.TrimSpace(v[1]),
		}
	}

	return ret
}

func askQuestion(q []string) bool {
	question := q[0]
	correctAnswer := q[1]
	var answer string

	fmt.Printf("%s= ", question)
	fmt.Scanf("%s", &answer)

	return answer == correctAnswer
}

func openFile(path string) []problem {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	return parseProblems(reader)
}
