package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	i := getInput()
	fmt.Printf("Part One answer is: %v\n", partOne(i))
	partTwo(i)
}

func getInput() []string {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	return strings.Split(string(b), "\n")
}

func partOne(input []string) int {
	totalSum := 0
	for _, val := range input {
		a := calculateCalibrationValue(val)
		totalSum += a
	}
	return totalSum
}

func partTwo(input []string) {}

func calculateCalibrationValue(i string) int {
	f, err := findFirstNumber(i)
	if err != nil {
		panic(err)
	}
	l, err := findLastNumber(i)
	if err != nil {
		panic(err)
	}
	fmtdAns := fmt.Sprintf("%d%d", f, l)
	n, err := strconv.Atoi(fmtdAns)
	if err != nil {
		panic(err)
	}
	return n
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

	return 0, fmt.Errorf("No first number found")
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

	return 0, fmt.Errorf("No last number found")
}
