package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readFile(name string, ch chan string) {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		ch <- line
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	close(ch)
}

func splitAndCount(lines chan string, words chan int) {
	for {
		line, ok := <-lines

		if !ok {
			close(words)
			break
		}

		if line != "" {
			words <- len(strings.Fields(line))
		}
	}
}

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Usage: word-count <file>")
		return
	}

	lines := make(chan string)
	words := make(chan int)

	go readFile(args[0], lines)
	go splitAndCount(lines, words)

	result := 0
	for c := range words {
		result += c
	}

	fmt.Println(result)
}
