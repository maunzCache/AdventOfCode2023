package main

import (
	"testing"
)

func TestCountNumbersInWordsPart1(t *testing.T) {
	// Part 1
	var words [5]NumberWord
	var dataWords = [4]string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}

	sum := countNumbersInWords(words[:], dataWords[:])

	if sum != 142 {
		t.Errorf("Wanted 142, got %d", sum)
	}
}

func TestCountNumbersInWordsPart2(t *testing.T) {
	// Part 2
	var words [8]NumberWord
	var dataWords = [7]string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}

	sum := countNumbersInWords(words[:], dataWords[:])

	if sum != 281 {
		t.Errorf("Wanted 281, got %d", sum)
	}
}

func TestCountNumbersInWordsStruggeling(t *testing.T) {
	// Part 2
	var words [9]NumberWord
	var dataWords = [8]string{
		"oneight",
		"eightwo",
		"1oneight8",
		"8eightwo2",
		"one18eight",
		"eight82two",
		"oneighteightone",
		"eightwotwoeight",
	}

	sum := countNumbersInWords(words[:], dataWords[:])

	if sum != 399 {
		t.Errorf("Wanted 399, got %d", sum)
	}
}
