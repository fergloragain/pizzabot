package parsing

import (
	"testing"
)

func TestParseGridSize1(t *testing.T) {
	gridSize := "5x5"
	x, y, err := ParseGridSize(gridSize)

	if x != 5 || y != 5 {
		t.Fatalf("Expected >%d, %d<, got >%d, %d<", 5, 5, x, y)
	}

	if err != nil {
		t.Fatalf("Got >%s<", err.Error())
	}
}

func TestParseGridSize2(t *testing.T) {
	gridSize := "1x512"
	x, y, err := ParseGridSize(gridSize)

	if x != 1 || y != 512 {
		t.Fatalf("Expected >%d, %d<, got >%d, %d<", 1, 512, x, y)
	}

	if err != nil {
		t.Fatalf("Got >%s<", err.Error())
	}
}

func TestParseGridSize3(t *testing.T) {
	gridSize := "6x12.8"
	x, y, err := ParseGridSize(gridSize)

	if x != -1 || y != -1 {
		t.Fatalf("Expected >%d, %d<, got >%d, %d<", -1, -1, x, y)
	}

	if err.Error() != "Invalid integer string '12.8', must be a valid integer" {
		t.Fatalf("Got >%s<", err.Error())
	}
}

func TestParseGridSize4(t *testing.T) {
	gridSize := "6x1x2"
	x, y, err := ParseGridSize(gridSize)

	if x != -1 || y != -1 {
		t.Fatalf("Expected >%d, %d<, got >%d, %d<", -1, -1, x, y)
	}

	if err.Error() != "Must specify two grid dimensions in the form XxY, e.g. 5x5" {
		t.Fatalf("Got >%s<", err.Error())
	}
}

func TestParseGridSize5(t *testing.T) {
	gridSize := "x1"
	x, y, err := ParseGridSize(gridSize)

	if x != -1 || y != -1 {
		t.Fatalf("Expected >%d, %d<, got >%d, %d<", -1, -1, x, y)
	}

	if err.Error() != "Must specify two grid dimensions in the form XxY, e.g. 5x5" {
		t.Fatalf("Got >%s<", err.Error())
	}
}

func TestParseGridSize6(t *testing.T) {
	gridSize := "-5x1"
	x, y, err := ParseGridSize(gridSize)

	if x != -1 || y != -1 {
		t.Fatalf("Expected >%d, %d<, got >%d, %d<", -1, -1, x, y)
	}

	if err.Error() != "Invalid integer '-5', must be a positive integer" {
		t.Fatalf("Got >%s<", err.Error())
	}
}

func TestParsePositiveInteger1(t *testing.T) {
	integerString := "5"
	x, err := parsePositiveInteger(integerString)

	if x != 5 {
		t.Fatalf("Expected >%d<, got >%d<", 5, x)
	}

	if err != nil {
		t.Fatalf("Got >%s<", err.Error())
	}
}

func TestParsePositiveInteger2(t *testing.T) {
	integerString := "5.0"
	x, err := parsePositiveInteger(integerString)

	if x != -1 {
		t.Fatalf("Expected >%d<, got >%d<", -1, x)
	}

	if err.Error() != "Invalid integer string '5.0', must be a valid integer" {
		t.Fatalf("Got >%s<", err.Error())
	}
}

func TestParsePositiveInteger3(t *testing.T) {
	integerString := "-9"
	x, err := parsePositiveInteger(integerString)

	if x != -1 {
		t.Fatalf("Expected >%d<, got >%d<", -1, x)
	}

	if err.Error() != "Invalid integer '-9', must be a positive integer" {
		t.Fatalf("Got >%s<", err.Error())
	}
}

func TestParseCoordinates1(t *testing.T) {
	coordinatesString := []string{"(1,", " 2)"}

	coords, err := ParseCoordinates(coordinatesString, 2, 2)

	if len(coords) != 1 {
		t.Fatalf("Expected length >%d<, got >%d<", 1, len(coords))
	}

	if coords[0].X != 1 && coords[0].Y != 2 {
		t.Fatalf("Expected >%d, %d<, got >%d, %d<", 1, 2, coords[0].X, coords[0].Y)
	}

	if err != nil {
		t.Fatalf("Got >%s<", err.Error())
	}
}

func TestParseCoordinates2(t *testing.T) {
	coordinatesString := []string{"(-1,", " 2)"}

	coords, err := ParseCoordinates(coordinatesString, 2, 2)

	if len(coords) != 0 {
		t.Fatalf("Expected length >%d<, got >%d<", 0, len(coords))
	}

	if err.Error() != "Invalid integer '-1', must be a positive integer" {
		t.Fatalf("Got >%s<", err.Error())
	}
}

func TestGenerateDirections1(t *testing.T) {
	coordinates := []*Point{&Point{1, 2}, &Point{4, 5}}

	directions := GenerateDirections(coordinates)

	if directions != "ENNDEEENNND" {
		t.Fatalf("Got >%s<", directions)
	}
}

func TestGenerateDirections2(t *testing.T) {
	coordinates := []*Point{&Point{1, 2}, &Point{4, 5}, &Point{0, 5}}

	directions := GenerateDirections(coordinates)

	if directions != "ENNDEEENNNDWWWWD" {
		t.Fatalf("Got >%s<", directions)
	}
}

func TestMove1(t *testing.T) {

	current, steps := move(0, 2, "A", "B")

	if current != 2 {
		t.Fatalf("Got >%d<", current)
	}

	if steps != "AA" {
		t.Fatalf("Got >%s<", steps)
	}
}

func TestMove2(t *testing.T) {

	current, steps := move(5, 2, "A", "B")

	if current != 2 {
		t.Fatalf("Got >%d<", current)
	}

	if steps != "BBB" {
		t.Fatalf("Got >%s<", steps)
	}
}

func TestMove3(t *testing.T) {

	current, steps := move(10, 2, "A", "Back")

	if current != 2 {
		t.Fatalf("Got >%d<", current)
	}

	if steps != "BackBackBackBackBackBackBackBack" {
		t.Fatalf("Got >%s<", steps)
	}
}

func TestMove4(t *testing.T) {

	current, steps := move(2, 2, "Y", "Z")

	if current != 2 {
		t.Fatalf("Got >%d<", current)
	}

	if steps != "" {
		t.Fatalf("Got >%s<", steps)
	}
}