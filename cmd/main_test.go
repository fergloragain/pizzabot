package main

import (
	"bytes"
	"fmt"
	"testing"
)

var output string

func init() {
	print = testPrinter
}

func testPrinter(format string, a ...interface{}) (int, error) {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, format, a...)
	output = buf.String()
	return 0, nil
}

func outputShouldEqual(expected string, t *testing.T) {
	if output != expected {
		t.Fatalf("Expected >%s<, got >%s<", expected, output)
	}
}

func TestPizzaBot1(t *testing.T) {
	inputStr := "5x5 (0, 0) (1, 3) (4, 4) (4, 2) (4, 2) (0, 1) (3, 2) (2, 3) (4, 1)"
	runPizzaBot(inputStr)
	outputShouldEqual("DENNNDEEENDSSDDWWWWSDEEENDWNDEESSD\n", t)
}

func TestPizzaBot2(t *testing.T) {
	inputStr := "5x5 (1, 3) (4, 4)"
	runPizzaBot(inputStr)
	outputShouldEqual("ENNNDEEEND\n", t)
}

func TestPizzaBot3(t *testing.T) {
	inputStr := "5x15 (0, 0) (1, 3) (4, 4) (4, 2) (4, 2) (0, 1) (3, 2) (2, 13) (5, 14)"
	runPizzaBot(inputStr)
	outputShouldEqual("DENNNDEEENDSSDDWWWWSDEEENDWNNNNNNNNNNNDEEEND\n", t)
}

func TestPizzaBot4(t *testing.T) {
	inputStr := "1x1 (0, 0)"
	runPizzaBot(inputStr)
	outputShouldEqual("D\n", t)
}

func TestPizzaBot5(t *testing.T) {
	inputStr := ""
	runPizzaBot(inputStr)
	outputShouldEqual("Error: Invalid input string\n", t)
}

func TestPizzaBot6(t *testing.T) {
	inputStr := "x15 (0, 0) (1, 3) (4, 4) (4, 2) (4, 2) (0, 1) (3, 2) (2, 13) (5, 14)"
	runPizzaBot(inputStr)
	outputShouldEqual("Error: Must specify two grid dimensions in the form XxY, e.g. 5x5\n", t)
}

func TestPizzaBot7(t *testing.T) {
	inputStr := "15 (0, 0) (1, 3) (4, 4) (4, 2) (4, 2) (0, 1) (3, 2) (2, 13) (5, 14)"
	runPizzaBot(inputStr)
	outputShouldEqual("Error: Must specify two grid dimensions in the form XxY, e.g. 5x5\n", t)
}

func TestPizzaBot8(t *testing.T) {
	inputStr := "x (0, 0) (1, 3) (4, 4) (4, 2) (4, 2) (0, 1) (3, 2) (2, 13) (5, 14)"
	runPizzaBot(inputStr)
	outputShouldEqual("Error: Must specify two grid dimensions in the form XxY, e.g. 5x5\n", t)
}

func TestPizzaBot9(t *testing.T) {
	inputStr := "-1x-1 ()"
	runPizzaBot(inputStr)
	outputShouldEqual("Error: Invalid integer '-1', must be a positive integer\n", t)
}

func TestPizzaBot10(t *testing.T) {
	inputStr := "1x1 ()"
	runPizzaBot(inputStr)
	outputShouldEqual("Error: Invalid coordinates '()', must be a pair of positive integers, e.g. (1, 2)\n", t)
}

func TestPizzaBot11(t *testing.T) {
	inputStr := "1x1 ("
	runPizzaBot(inputStr)
	outputShouldEqual("Error: Invalid coordinates '(', must be a pair of positive integers, e.g. (1, 2)\n", t)
}

func TestPizzaBot12(t *testing.T) {
	inputStr := "1x1 "
	runPizzaBot(inputStr)
	outputShouldEqual("Error: Invalid input string\n", t)
}

func TestPizzaBot13(t *testing.T) {
	inputStr := "1x1 ())"
	runPizzaBot(inputStr)
	outputShouldEqual("Error: Invalid coordinates '())', must be a pair of positive integers, e.g. (1, 2)\n", t)
}

func TestPizzaBot14(t *testing.T) {
	inputStr := "1x1 (5, 5)"
	runPizzaBot(inputStr)
	outputShouldEqual("Error: Point (5, 5) must be within grid dimensions 1x1\n", t)
}

func TestPizzaBot15(t *testing.T) {
	inputStr := "1x1 (-5, 5)"
	runPizzaBot(inputStr)
	outputShouldEqual("Error: Invalid integer '-5', must be a positive integer\n", t)
}

func TestPizzaBot16(t *testing.T) {
	inputStr := "1x1 (5e, 5)"
	runPizzaBot(inputStr)
	outputShouldEqual("Error: Invalid integer string '5e', must be a valid integer\n", t)
}

func TestPizzaBot17(t *testing.T) {
	inputStr := "           5x5 (0, 0      )          (1 , 3) (4,       4) (4, 2     ) (4, 2) (0, 1) (3, 2) (2, 3) (         4, 1)                 "
	runPizzaBot(inputStr)
	outputShouldEqual("DENNNDEEENDSSDDWWWWSDEEENDWNDEESSD\n", t)
}

func TestPizzaBot18(t *testing.T) {
	inputStr := "      5x5 {0, 0      )          (1 , 3) (4,    4) (4, 2     ) (4, 2) (0, 1) (3, 2) (2, 3) (  4, 1)       "
	runPizzaBot(inputStr)
	outputShouldEqual("Error: Invalid integer string '{0', must be a valid integer\n", t)
}
