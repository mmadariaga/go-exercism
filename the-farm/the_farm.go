package thefarm

import (
	"errors"
	"fmt"
)

// TODO: define the 'DivideFood' function
func DivideFood(calculator FodderCalculator, cowNum int) (float64, error) {

	fodderAmount, err := calculator.FodderAmount(cowNum)
	if err != nil {
		return 0, err
	}

	fatteningFactor, err := calculator.FatteningFactor()
	if err != nil {
		return 0, err
	}

	return fodderAmount * fatteningFactor / float64(cowNum), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(calculator FodderCalculator, cowNum int) (float64, error) {

	if cowNum < 1 {
		return 0.0, errors.New("invalid number of cows")
	}

	return DivideFood(calculator, cowNum)
}

// TODO: define the 'ValidateNumberOfCows' function
type InvalidCowsError struct {
	cowNum   int
	errorMsg string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.cowNum, e.errorMsg)
}

func ValidateNumberOfCows(cowNum int) *InvalidCowsError {

	if cowNum < 0 {
		return &InvalidCowsError{
			cowNum:   cowNum,
			errorMsg: "there are no negative cows",
		}
	}

	if cowNum == 0 {
		return &InvalidCowsError{
			cowNum:   cowNum,
			errorMsg: "no cows don't need food",
		}
	}

	return nil
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
