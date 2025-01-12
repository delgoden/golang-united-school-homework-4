package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {

	if len(input) == strings.Count(input, " ") {
		return "", fmt.Errorf("nothing was found in the passed string: %w", errorEmptyInput)
	}

	input = strings.ReplaceAll(input, " ", "")
	operands := make([]string, 0, 2)
	j := 0
	for i := len(input); i > 0; {
		j = strings.LastIndexAny(input, "+-")
		if j == -1 {
			operands = append(operands, input)
			break
		}
		operands = append(operands, input[j:i])
		i = j
		input = input[:i]
	}

	if len(operands) < 2 {
		return "", fmt.Errorf("less than two operands: %w", errorNotTwoOperands)
	}
	if len(operands) > 2 {
		return "", fmt.Errorf("more than two operands: %w", errorNotTwoOperands)
	}

	operandFirst, err := strconv.Atoi(operands[0])
	if err != nil {
		return "", fmt.Errorf("operand is not number: %w", err)
	}
	operandSecond, err := strconv.Atoi(operands[1])
	if err != nil {
		return "", fmt.Errorf("operand is not number: %w", err)
	}
	output = strconv.Itoa(operandFirst + operandSecond)

	return
}
