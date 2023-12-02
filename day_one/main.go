package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	i := getInput()
	partOne(i)
	partTwo(i)
}

func getInput() []string {
	b, err := ioutil.ReadFile("example_input.txt")
	if err != nil {
		panic(err)
	}

	return strings.Split(string(b), "\n")
}

func partOne(input []string) {
	for i, val := range input {
		a := calculateCalibrationValue(val)
		fmt.Printf("Index %v answer: %v\n", i, a)
	}
}

func partTwo(input []string) {}

func calculateCalibrationValue(i string) int {
	return 0
}
