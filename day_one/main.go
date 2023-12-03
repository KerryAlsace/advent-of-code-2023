package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var NumberStringsToInt = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

var NumberFirstLetters = map[string][]string{
	"o": {
		"one",
	},
	"t": {
		"two",
		"three",
	},
	"f": {
		"four",
		"five",
	},
	"s": {
		"six",
		"seven",
	},
	"e": {
		"eight",
	},
	"n": {
		"nine",
	},
}

func main() {
	i := getInput()
	partOneAns, err := partOne(i)
	if err != nil {
		fmt.Printf("Error in partOne: %v\n", err)
	}
	partTwoAns, err := partTwo(i)
	if err != nil {
		fmt.Printf("Error in partTwo: %v\n", err)
	}
	fmt.Printf("Part One answer is: %v\n", partOneAns)
	fmt.Printf("Part Two answer is: %v\n", partTwoAns)
}

func getInput() []string {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	return strings.Split(string(b), "\n")
}

func partOne(input []string) (int, error) {
	totalSum := 0
	for _, val := range input {
		a, err := calculateCalibrationValue(val)
		if err != nil {
			return 0, fmt.Errorf("Error for val %v: %v", val, err)
		}
		totalSum += a
	}
	return totalSum, nil
}

func partTwo(input []string) (int, error) {
	totalSum := 0
	for _, val := range input {
		a, err := calculateCalibrationValueWithStrings(val)
		if err != nil {
			return 0, fmt.Errorf("Error for val %v: %v\n", val, err)
		}
		totalSum += a
	}
	return totalSum, nil
}

func calculateCalibrationValue(i string) (int, error) {
	f, err := findFirstNumber(i)
	if err != nil {
		return 0, err
	}
	l, err := findLastNumber(i)
	if err != nil {
		return 0, err
	}
	fmtdAns := fmt.Sprintf("%d%d", f, l)
	n, err := strconv.Atoi(fmtdAns)
	if err != nil {
		return 0, err
	}
	return n, nil
}

func findFirstNumber(input string) (int, error) {
	for _, char := range input {
		fmtdChar := fmt.Sprintf("%c", char)
		n, err := strconv.Atoi(fmtdChar)
		if err != nil {
			continue
		}

		return n, nil
	}

	return 0, fmt.Errorf("no first number found")
}

func findLastNumber(input string) (int, error) {
	lastIndex := len(input) - 1
	for i := lastIndex; i >= 0; i-- {
		fmtdChar := fmt.Sprintf("%c", input[i])
		n, err := strconv.Atoi(fmtdChar)
		if err != nil {
			continue
		}

		return n, nil
	}

	return 0, fmt.Errorf("no last number found")
}

func calculateCalibrationValueWithStrings(i string) (int, error) {
	f, err := findFirstNumberWithStrings(i)
	if err != nil {
		return 0, err
	}
	l, err := findLastNumberWithStrings(i)
	if err != nil {
		return 0, err
	}
	fmtdAns := fmt.Sprintf("%d%d", f, l)
	n, err := strconv.Atoi(fmtdAns)
	if err != nil {
		return 0, err
	}
	return n, nil
}

func findFirstNumberWithStrings(input string) (int, error) {
	for i, char := range input {
		fmtdChar := fmt.Sprintf("%c", char)
		// First check if it is an int, and return it if so
		n, nilIfInt := strconv.Atoi(fmtdChar)
		if nilIfInt != nil {
			// fmtdChar is not an int, so check if it is the starting letter
			// of one of the number strings
			if NumberFirstLetters[fmtdChar] != nil {
				n, err := checkForNumString(input, i, fmtdChar)
				if err != nil {
					continue
				}
				return n, nil
			}

			continue
		}

		return n, nil
	}

	return 0, fmt.Errorf("no first number found")
}

func findLastNumberWithStrings(input string) (int, error) {
	lastIndex := len(input) - 1
	for i := lastIndex; i >= 0; i-- {
		fmtdChar := fmt.Sprintf("%c", input[i])
		// First check if it is an int, and return it if so
		n, nilIfInt := strconv.Atoi(fmtdChar)
		if nilIfInt != nil {
			// fmtdChar is not an int, so check if it is the starting letter
			// of one of the number strings
			if NumberFirstLetters[fmtdChar] != nil {
				n, err := checkForNumString(input, i, fmtdChar)
				if err != nil {
					continue
				}
				return n, nil
			}

			continue
		}

		return n, nil
	}
	return 0, fmt.Errorf("no last number found")
}

func checkForNumString(input string, index int, s string) (int, error) {
	possibleNumbers := NumberFirstLetters[s]
	for _, numNameString := range possibleNumbers {
		lenNumString := len(numNameString)
		inputEndIndex := lenNumString + index
		if inputEndIndex < index {
			continue
		}
		// Check that the input supports that last index
		if len(input) < inputEndIndex {
			continue
		}
		wordInInput := input[index:inputEndIndex]
		if wordInInput == numNameString {
			return NumberStringsToInt[wordInInput], nil
		}
	}
	return 0, fmt.Errorf("not found")
}
