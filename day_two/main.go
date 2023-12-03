package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"unicode"
)

const RED int = 12
const GREEN int = 13
const BLUE int = 14

type Set struct {
	red   int
	green int
	blue  int
}

type Game struct {
	name int
	sets []Set
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
	games, err := getGameDetails(input)
	if err != nil {
		return 0, fmt.Errorf("error getting game details: %v", err)
	}
	return 0, fmt.Errorf("error in partOne")
}

func partTwo(input []string) (int, error) {
	return 0, fmt.Errorf("error in partTwo")
}

func getGameDetails(input []string) ([]Game, error) {
	var games []Game
	for _, gameData := range input {
		var game Game
		sp := strings.Split(gameData, ":")
		g := sp[0]
		d := sp[1]

		// Grab game name
		f := func(c rune) bool {
			return !unicode.IsNumber(c)
		}
		spl := strings.FieldsFunc(g, f)
		n, err := strconv.Atoi(spl[0])
		if err != nil {
			return []Game{}, fmt.Errorf("could not convert %v to int", spl[0])
		}
		game.name = n

		// Grab each set's details
		r := strings.Split(d, ";")
		for _, s := range r {
			f := func(c rune) bool {
				return !unicode.IsLetter(c) && !unicode.IsNumber(c) && !unicode.IsSpace(c)
			}
			setSlice := strings.FieldsFunc(s, f)

			var set Set
			// Range through the details of the set
			for _, hand := range setSlice {
				// Get the number and the color
				f := func(c rune) bool {
					return !unicode.IsNumber(c)
				}
				num := strings.FieldsFunc(hand, f)[0]
				n, err := strconv.Atoi(num)
				if err != nil {
					return []Game{}, fmt.Errorf("could not convert %v to int", num)
				}

				f = func(c rune) bool {
					return !unicode.IsLetter(c)
				}
				color := strings.FieldsFunc(hand, f)[0]
				switch color {
				case "red":
					set.red = n
				case "green":
					set.green = n
				case "blue":
					set.blue = n
				default:
					return []Game{}, fmt.Errorf("encountered unknown color %v", color)
				}
			}

			game.sets = append(game.sets, set)
		}
		games = append(games, game)
	}
	return games, nil
}
