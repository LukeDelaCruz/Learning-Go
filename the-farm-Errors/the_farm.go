package thefarm

import (
	"errors"
	"fmt"
)

// Question 1:
func DivideFood(fodderCalc FodderCalculator, numOfCows int) (float64, error) {
	totalFodderForAllCows, fodderAmtErr := fodderCalc.FodderAmount(numOfCows)
	if fodderAmtErr != nil {
		return 0, fodderAmtErr
	}

	fatteningFactor, fatteningFactorErr := fodderCalc.FatteningFactor()
	if fatteningFactorErr != nil {
		return 0, fatteningFactorErr
	}

	return (totalFodderForAllCows * fatteningFactor) / float64(numOfCows), nil
}

// Question 2:
func ValidateInputAndDivideFood(fodderCalc FodderCalculator, numOfCows int) (float64, error) {
	if numOfCows <= 0 {
		return 0, errors.New("invalid number of cows")
	}

	return DivideFood(fodderCalc, numOfCows)
}

// Question 3:
type InvalidCowsError struct {
	numOfCows int
	message   string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.numOfCows, e.message)
}

func ValidateNumberOfCows(numOfCows int) *InvalidCowsError {
	if numOfCows < 0 {
		return &InvalidCowsError{
			numOfCows: numOfCows,
			message:   "there are no negative cows",
		}
	}

	if numOfCows == 0 {
		return &InvalidCowsError{
			numOfCows: numOfCows,
			message:   "no cows don't need food",
		}
	}

	return nil
}
