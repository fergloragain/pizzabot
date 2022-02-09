package parsing

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	InputDelimiter = " "
	North          = "N"
	East           = "E"
	South          = "S"
	West           = "W"
	Drop           = "D"
)

// Point struct holds the X and Y coordinates of a point
type Point struct {
	X int
	Y int
}

// parseGridSize takes a string containing grid dimensions, e.g. 5x5, and returns the individual X and Y dimensions of
// the grid as integers. If there is a problem parsing the grid size, -1 is returned for both dimensions, along with
// an error message
func ParseGridSize(gridSize string) (int, int, error) {

	// split the gridSize string e.g. 5x5 using "x" as a separator
	xyDimensions := strings.Split(gridSize, "x")

	// if more than two, or less than two, dimensions have been specified, return an error message
	if len(xyDimensions) != 2 {
		return -1, -1, errors.New("Must specify two grid dimensions in the form XxY, e.g. 5x5")
	}

	// if two dimensions have been specified, and either one is a zero-length string, return an error message
	if len(xyDimensions[0]) == 0 || len(xyDimensions[1]) == 0 {
		return -1, -1, errors.New("Must specify two grid dimensions in the form XxY, e.g. 5x5")
	}

	// parse the X dimension into an int
	x, err := parsePositiveInteger(xyDimensions[0])

	// if an error ocurred during parsing, return the error
	if err != nil {
		return -1, -1, err
	}

	// parse the Y dimension into an int
	y, err := parsePositiveInteger(xyDimensions[1])

	// if an error ocurred during parsing, return the error
	if err != nil {
		return -1, -1, err
	}

	// return the successfully parsed dimensions as ints
	return x, y, nil
}

// move accepts a current position, a target position, and strings representing the forward and backward directions.
// for the given positions, the number of steps in a particular direction are generated.
func move(current, target int, forward, backward string) (int, string) {

	// determine the number of steps between the current and target positions
	steps := current - target

	// if the current position is ahead of the target position, we will get a negative value, but we want an absolute
	// value, so we negate a negative count
	if steps < 0 {
		steps = -steps
	}

	if target > current {
		// if the target is ahead of the current position, repeat the forward direction for the number of steps that we have
		return current + steps, strings.Repeat(forward, steps)
	} else if target < current {
		// if the target is behind the current position, repeat the backward direction for the number of steps that we have
		return current - steps, strings.Repeat(backward, steps)
	}

	// if the current position and the target position are the same, just return the current position and an empty string
	return current, ""

}

// generateDirections accepts an array of Points, and generates the directions from origin (0, 0), to each point using
// N for north, E for east, S for south, W for west, and D for drop.
func GenerateDirections(allCoordinatePairs []*Point) string {

	currentXPosition := 0
	currentYPosition := 0

	directions := ""

	for _, coordinates := range allCoordinatePairs {
		newXDirections := ""
		newYDirections := ""

		// get the new X position, and the directions to it from the previous point. East and West are the forward and
		// backward directions, respectively
		currentXPosition, newXDirections = move(currentXPosition, coordinates.X, East, West)

		// get the new Y position, and the directions to it from the previous point. North and South are the forward and
		// backward directions, respectively
		currentYPosition, newYDirections = move(currentYPosition, coordinates.Y, North, South)

		// append the new X and Y directions respectively to the existing directions, followed by a pizza drop
		directions = fmt.Sprintf("%s%s%s%s", directions, newXDirections, newYDirections, Drop)
	}

	return directions
}

// parseCoordinates accepts an array of strings, containing all coordinates, parses them into int pairs, and returns an
// array of Points
func ParseCoordinates(coordinates []string, gridX, gridY int) ([]*Point, error) {

	allCoordinates := []*Point{}

	// join the array of the coordinates with spaces, restoring the original format that was passed to the pizzaBot
	coordinatesAsString := strings.Join(coordinates, InputDelimiter)

	// remove whitespace from the coordinate string
	coordinatesAsString = strings.Replace(coordinatesAsString, " ", "", -1)

	// split the coordinate string with the closing and open parentheses
	coordinatesArray := strings.Split(coordinatesAsString, ")(")

	// for each coordinate pair in the coordinate array
	for _, coordinatePair := range coordinatesArray {

		// attempt to parse the coordinate pair into valid X and Y ints
		xCoordinate, yCoordinate, err := parseXYCoordinates(coordinatePair, gridX, gridY)

		// if an error occurred, return an empty Point array and an error message
		if err != nil {
			return []*Point{}, err
		}

		// otherwise, append a new Point containing the X and Y coordinates to the Point array
		allCoordinates = append(allCoordinates, &Point{xCoordinate, yCoordinate})
	}

	return allCoordinates, nil

}

// parseXYCoordinates takes a pair of coordinates in the form of a string, along with the grid dimensions within which
// the coordinates are expected to exist. The coordinate pair is parsed into a pair of ints, and provided the ints
// are within the bounds of the grid, they are returned to the calling function. Otherwise, -1 is returned for each
// coordinate, along with a relevant error message
func parseXYCoordinates(coordinatePair string, gridX int, gridY int) (int, int, error) {

	// parse the coordinate pair into an array of strings, e.g. "(1, 2)" becomes ["1", "2"]
	coordinatePairAsArray := parseCoordinatePair(coordinatePair)

	// if we have more or less than two coordinates, return an error message
	if len(coordinatePairAsArray) != 2 {
		return -1, -1, errors.New(fmt.Sprintf("Invalid coordinates '%s', must be a pair of positive integers, e.g. (1, 2)", coordinatePair))
	}

	// parse the X coordinate into an int
	xCoordinate, err := parsePositiveInteger(coordinatePairAsArray[0])

	// if an error ocurred during parsing, return an error
	if err != nil {
		return -1, -1, err
	}

	// parse the Y coordinate into an int
	yCoordinate, err := parsePositiveInteger(coordinatePairAsArray[1])

	// if an error ocurred during parsing, return an error
	if err != nil {
		return -1, -1, err
	}

	// if the coordinates are outside the bounds of the grid, return an error message
	if xCoordinate > gridX || yCoordinate > gridY {
		return -1, -1, errors.New(fmt.Sprintf("Point (%d, %d) must be within grid dimensions %dx%d", xCoordinate, yCoordinate, gridX, gridY))
	}

	return xCoordinate, yCoordinate, nil
}

// parseCoordinatePair removes the parentheses and other redundant runes from the coordinate pair string, e.g. "(1, 2)"
// becomes an array of the form ["1", "2"]
func parseCoordinatePair(coordinatePair string) []string {

	coordinatePairStripped := strings.Replace(coordinatePair, "(", "", 1)
	coordinatePairStripped = strings.Replace(coordinatePairStripped, ")", "", 1)
	coordinatePairStripped = strings.Replace(coordinatePairStripped, InputDelimiter, "", -1)

	return strings.Split(coordinatePairStripped, ",")
}

// parsePositiveInteger accepts a string, and attempts to parse it into an int. If the string contains a positive
// integer, the int is returned to the calling function. Otherwise, if the string does not contain an integer, of if the
// string contains a negative integer, an error message is returned to the calling function
func parsePositiveInteger(integerString string) (int, error) {

	// attempt to parse the string into an int
	parsedInteger, err := strconv.Atoi(integerString)

	// if an error occurred during parsing, i.e. if the string is not an integer, return -1 and an error message
	if err != nil {
		return -1, errors.New(fmt.Sprintf("Invalid integer string '%s', must be a valid integer", integerString))
	}

	// if the string is a negative integer, return -1 and an error message
	if parsedInteger < 0 {
		return -1, errors.New(fmt.Sprintf("Invalid integer '%d', must be a positive integer", parsedInteger))
	}

	// if the string is a positive integer, return the integer with no error message
	return parsedInteger, nil
}
