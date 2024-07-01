package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type NumberWord struct {
	first int
	last  int
	word  string
}

func findFirstNumberInWord(word string) int {
	var result int
	for index := 0; index < len(word); index++ {
		letter := word[index : index+1]
		number, err := strconv.Atoi(letter)
		if err == nil {
			result = number
			break
		}
	}
	return result
}

func findLastNumberInWord(word string) int {
	var result int
	for index := len(word); index > 0; index-- {
		letter := word[index-1 : index]
		number, err := strconv.Atoi(letter)
		if err == nil {
			result = number
			break
		}
	}
	return result
}

func replaceAllNumberWords(word string) string {
	word = strings.ReplaceAll(word, "one", "o1e")
	word = strings.ReplaceAll(word, "two", "t2o")
	word = strings.ReplaceAll(word, "three", "t3e")
	word = strings.ReplaceAll(word, "four", "f4r")
	word = strings.ReplaceAll(word, "five", "f5e")
	word = strings.ReplaceAll(word, "six", "s6x")
	word = strings.ReplaceAll(word, "seven", "s7n")
	word = strings.ReplaceAll(word, "eight", "e8t")
	word = strings.ReplaceAll(word, "nine", "n9e")
	return word
}

func countNumbersInWords(words []NumberWord, dataWords []string) int {
	for dataIndex, dataWord := range dataWords {
		revisedWord := replaceAllNumberWords(dataWord)

		words[dataIndex].word = revisedWord
	}

	sum := 0
	for _, entry := range words {
		entry.first = findFirstNumberInWord(entry.word)
		entry.last = findLastNumberInWord(entry.word)

		fmt.Printf("%v: %v%v\r\n", entry.word, entry.first, entry.last)
		sum += (entry.first * 10) + entry.last
	}

	return sum
}

func main() {
	inputFile := "input.txt"
	dataAsByte, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	var words [1001]NumberWord
	dataAsString := bytes.NewBuffer(dataAsByte).String()
	dataWords := strings.Split(dataAsString, "\n")

	sum := countNumbersInWords(words[:], dataWords)

	fmt.Println("Final sum:", sum)

	// Part 1: 54927
	// Part 2: 54581
}
