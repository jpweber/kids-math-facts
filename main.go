package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type problem struct {
	value1 int
	value2 int
	answer int
}

func (p *problem) new() {
	p.value1 = rand.Intn(10)
	p.value2 = rand.Intn(10)
	p.answer = p.value1 * p.value2
}

func makeAnswer(input []byte) int {
	str := string(input)
	retVal, err := strconv.Atoi(strings.Trim(str, "\n"))
	if err != nil {
		fmt.Println(err)
	}
	return retVal
}

func showResults(results map[string]int) {
	answered := results["correct"] + results["incorrect"]
	fmt.Println("You answered", answered, "Questions")
	fmt.Println("you got", results["correct"], "answers right")
	fmt.Println("you got", results["incorrect"], "answers wrong")
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	testLength := flag.Int("l", 60, "How long the test lasts in seconds")
	// Once all flags are declared, call `flag.Parse()`
	// to execute the command-line parsing.
	flag.Parse()

	// create channel for ending main game loop after timer
	done := make(chan bool, 1)

	// create and start timer
	timer := time.NewTimer(time.Second * time.Duration(*testLength))
	go func(done chan bool) {
		<-timer.C
		done <- true
	}(done)

	results := make(map[string]int)
	// main game loop
loop:
	for {
		select {
		case _ = <-done:
			fmt.Println("Times up!")
			showResults(results)
			break loop
		default:
			p := problem{}
			p.new()
			buf := bufio.NewReader(os.Stdin)
			fmt.Printf("%d %s %d %s", p.value1, "*", p.value2, "= ")
			input, _ := buf.ReadBytes('\n')

			if p.answer == makeAnswer(input) {
				results["correct"]++
			} else {
				results["incorrect"]++
			}
		}
	}
	os.Exit(0)
}
