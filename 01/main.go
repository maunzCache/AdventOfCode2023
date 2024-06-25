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

// NOTE: For this solution i'll focus on structuring the code and not go for performance.
func main() {
	// var words = [4]NumberWord{
	// 	{word: "1abc2"},
	// 	{word: "pqr3stu8vwx"},
	// 	{word: "a1b2c3d4e5f"},
	// 	{word: "treb7uchet"},
	// }

	inputFile := "input.txt"
	data_as_byte, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	var words [1001]NumberWord
	data_as_string := bytes.NewBuffer(data_as_byte).String()
	data_words := strings.Split(data_as_string, "\n")

	// TODO: There must be a better way - this feels oldschool
	for index := 0; index < len(data_words); index++ {
		words[index].word = data_words[index]
	}

	sum := 0
	for _, entry := range words {
		for index := 0; index < len(entry.word); index++ {
			letter := entry.word[index : index+1]
			number, err := strconv.Atoi(letter)
			if err == nil {
				entry.first = number
				break
			}
		}

		for index := len(entry.word); index > 0; index-- {
			letter := entry.word[index-1 : index]
			number, err := strconv.Atoi(letter)
			if err == nil {
				entry.last = number
				break
			}
		}

		fmt.Printf("%v%v\r\n", entry.first, entry.last)
		sum += (entry.first * 10) + entry.last
	}

	fmt.Println("Final sum:", sum)
}
