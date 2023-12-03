package main

import (
	"reflect"
	"testing"
)

func TestPartOne(t *testing.T) {
	tests := []struct {
		testName       string
		input          []string
		expectedOutput int
	}{
		{
			testName: "1",
			input: []string{
				"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
				"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
				"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
				"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
				"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			},
			expectedOutput: 8,
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

func TestGetGameDetails(t *testing.T) {
	tests := []struct {
		testName       string
		input          []string
		expectedOutput []Game
	}{
		{
			testName: "1",
			input: []string{
				"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
				"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
				"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
				"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
				"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			},
			expectedOutput: []Game{
				{
					name: 1,
					sets: []Set{
						{
							red:  4,
							blue: 3,
						},
						{
							red:   1,
							green: 2,
							blue:  6,
						},
						{
							green: 2,
						},
					},
				},
				{
					name: 2,
					sets: []Set{
						{
							green: 2,
							blue:  1,
						},
						{
							red:   1,
							green: 3,
							blue:  4,
						},
						{
							green: 1,
							blue:  1,
						},
					},
				},
				{
					name: 3,
					sets: []Set{
						{
							red:   20,
							green: 8,
							blue:  6,
						},
						{
							red:   4,
							green: 13,
							blue:  5,
						},
						{
							green: 5,
							red:   1,
						},
					},
				},
				{
					name: 4,
					sets: []Set{
						{
							red:   3,
							green: 1,
							blue:  6,
						},
						{
							red:   6,
							green: 3,
						},
						{
							red:   14,
							green: 3,
							blue:  15,
						},
					},
				},
				{
					name: 5,
					sets: []Set{
						{
							red:   6,
							green: 3,
							blue:  1,
						},
						{
							red:   1,
							green: 2,
							blue:  2,
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			calculatedOutput, err := getGameDetails(tt.input)
			if err != nil {
				t.Errorf("getGameDetails() returned an error: %v", err)
			}
			if !reflect.DeepEqual(calculatedOutput, tt.expectedOutput) {
				t.Errorf("getGameDetails() = [%v], expected [%v]", calculatedOutput, tt.expectedOutput)
			}
		})
	}
}
