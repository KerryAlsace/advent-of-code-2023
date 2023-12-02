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
			calculatedOutput, err := calculateCalibrationValue(tt.input)
			if err != nil {
				t.Errorf("calculateCalibrationValue() returned an error: %v", err)
			}
			if calculatedOutput != tt.expectedOutput {
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
			calculatedOutput, err := partOne(tt.input)
			if err != nil {
				t.Errorf("partOne() returned an error: %v", err)
			}
			if calculatedOutput != tt.expectedOutput {
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
			calculatedOutput, err := partTwo(tt.input)
			if err != nil {
				t.Errorf("partTwo() returned an error: %v", err)
			}
			if calculatedOutput != tt.expectedOutput {
				t.Errorf("partTwo() = [%v], expected [%v]", calculatedOutput, tt.expectedOutput)
			}
		})
	}
}

func TestCheckForNumString(t *testing.T) {
	tests := []struct {
		testName       string
		input          string
		index          int
		s              string
		expectedOutput int
	}{
		{
			testName:       "1",
			input:          "two1nine",
			index:          0,
			s:              "t",
			expectedOutput: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			calculatedOutput, err := checkForNumString(tt.input, tt.index, tt.s)
			if err != nil {
				t.Errorf("checkForNumString() returned an error: %v", err)
			}
			if calculatedOutput != tt.expectedOutput {
				t.Errorf("checkForNumString() = [%v], expected [%v]", calculatedOutput, tt.expectedOutput)
			}
		})
	}
}

func TestCalculateCalibrationValueWithStrings(t *testing.T) {
	tests := []struct {
		testName       string
		input          string
		expectedOutput int
	}{
		{
			testName:       "1",
			input:          "two1nine",
			expectedOutput: 29,
		},
		{
			testName:       "2",
			input:          "eightwothree",
			expectedOutput: 83,
		},
		{
			testName:       "3",
			input:          "abcone2threexyz",
			expectedOutput: 13,
		},
		{
			testName:       "4",
			input:          "xtwone3four",
			expectedOutput: 24,
		},
		{
			testName:       "5",
			input:          "4nineeightseven2",
			expectedOutput: 42,
		},
		{
			testName:       "6",
			input:          "zoneight234",
			expectedOutput: 14,
		},
		{
			testName:       "7",
			input:          "7pqrstsixteen",
			expectedOutput: 76,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			calculatedOutput, err := calculateCalibrationValueWithStrings(tt.input)
			if err != nil {
				t.Errorf("calculateCalibrationValueWithStrings() returned an error: %v", err)
			}
			if calculatedOutput != tt.expectedOutput {
				t.Errorf("calculateCalibrationValueWithStrings() = [%v], expected [%v]", calculatedOutput, tt.expectedOutput)
			}
		})
	}
}

func TestFindFirstNumberWithStrings(t *testing.T) {
	tests := []struct {
		testName       string
		input          string
		expectedOutput int
	}{
		{
			testName:       "1",
			input:          "two1nine",
			expectedOutput: 2,
		},
		{
			testName:       "2",
			input:          "eightwothree",
			expectedOutput: 8,
		},
		{
			testName:       "3",
			input:          "abcone2threexyz",
			expectedOutput: 1,
		},
		{
			testName:       "4",
			input:          "xtwone3four",
			expectedOutput: 2,
		},
		{
			testName:       "5",
			input:          "4nineeightseven2",
			expectedOutput: 4,
		},
		{
			testName:       "6",
			input:          "zoneight234",
			expectedOutput: 1,
		},
		{
			testName:       "7",
			input:          "7pqrstsixteen",
			expectedOutput: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			calculatedOutput, err := findFirstNumberWithStrings(tt.input)
			if err != nil {
				t.Errorf("findFirstNumberWithStrings() returned an error: %v", err)
			}
			if calculatedOutput != tt.expectedOutput {
				t.Errorf("findFirstNumberWithStrings() = [%v], expected [%v]", calculatedOutput, tt.expectedOutput)
			}
		})
	}
}
