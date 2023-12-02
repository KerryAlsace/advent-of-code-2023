package main

import (
	"testing"
)

func TestCalculateCalibrationValue(t *testing.T) {
	tests := []struct {
		testName       string
		input          string
		expectedOutput int
	}{
		{
			testName:       "1",
			input:          "1abc2",
			expectedOutput: 12,
		},
		{
			testName:       "2",
			input:          "pqr3stu8vwx",
			expectedOutput: 38,
		},
		{
			testName:       "3",
			input:          "a1b2c3d4e5f",
			expectedOutput: 15,
		},
		{
			testName:       "4",
			input:          "treb7uchet",
			expectedOutput: 77,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			if calculatedOutput := calculateCalibrationValue(tt.input); calculatedOutput != tt.expectedOutput {
				t.Errorf("calculateCalibrationValue() = [%v], expected [%v]", calculatedOutput, tt.expectedOutput)
			}
		})
	}
}

func TestPartOne(t *testing.T) {
	tests := []struct {
		testName       string
		input          []string
		expectedOutput int
	}{
		{
			testName: "1",
			input: []string{
				"1abc2",
				"pqr3stu8vwx",
				"a1b2c3d4e5f",
				"treb7uchet",
			},
			expectedOutput: 142,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			if calculatedOutput := partOne(tt.input); calculatedOutput != tt.expectedOutput {
				t.Errorf("partOne() = [%v], expected [%v]", calculatedOutput, tt.expectedOutput)
			}
		})
	}
}

func TestPartTwo(t *testing.T) {
	tests := []struct {
		testName       string
		input          []string
		expectedOutput int
	}{
		{
			testName: "1",
			input: []string{
				"two1nine",
				"eightwothree",
				"abcone2threexyz",
				"xtwone3four",
				"4nineeightseven2",
				"zoneight234",
				"7pqrstsixteen",
			},
			expectedOutput: 281,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			if calculatedOutput := partOne(tt.input); calculatedOutput != tt.expectedOutput {
				t.Errorf("partOne() = [%v], expected [%v]", calculatedOutput, tt.expectedOutput)
			}
		})
	}
}
