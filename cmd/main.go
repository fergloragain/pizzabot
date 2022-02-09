package main

import (
	"fmt"
	"os"
	"strings"
	"github.com/fergloragain/pizzabot/parsing"
)

// assign the function fmt.Printf to the variable print, allowing us to later specify a custom function for capturing output in unit tests
var print = fmt.Printf

// If the user specifies an argument string with grid dimensions and coordinates, run the pizza bot
func main() {
	if len(os.Args) > 1 {
		runPizzaBot(os.Args[1])
	} else {
		printHelp("No input specified")
	}
}

// pizzaBot accepts a string containing grid dimensions and coordinates, parses them into ints, validates them, and then
// generates directions to all coordinates within the grid
func runPizzaBot(input string) {

	// gather all fields within the input string. This prevents redundant spacing causing confusion when parsing the
	// input string
	inputComponents := strings.Fields(input)

	// before continuing, check that we have at least two input fields, i.e. grid dimensions and at least one
	// coordinate set
	if len(inputComponents) < 2 {
		printHelp("Invalid input string")
		return
	}

	gridSize := inputComponents[0]
	coordinates := inputComponents[1:]

	// parse the grid size into ints
	gridX, gridY, err := parsing.ParseGridSize(gridSize)

	// if an error occurred during parsing, print the error and return
	if err != nil {
		printHelp(err.Error())
		return
	}

	// parse the coordinates into int pairs
	allCoordinatePairs, err := parsing.ParseCoordinates(coordinates, gridX, gridY)

	// if an error occurred during parsing, print the error and return
	if err != nil {
		printHelp(err.Error())
		return
	}

	// with the parsed grid dimensions and coordinates, generate directions
	directions := parsing.GenerateDirections(allCoordinatePairs)

	// print the generated directions
	print("%s\n", directions)
}

// printHelp takes an error message, prints it, and then prints an example usage of pizzaBot
func printHelp(message string) {
	print("Error: %s\n", message)
	fmt.Println("Usage")
	fmt.Println("./pizzabot \"5x5 (1, 2) (3, 4)\"")
}
