package farm

import (
	"errors"
	"fmt"
)

type FodderCalculator interface {
	FodderAmount(int) (float64, error)
	FatteningFactor() (float64, error)
}

type InvalidCowsError struct {
	NumCows int
	Message string
}

func (err *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", err.NumCows, err.Message)
}

func DivideFood(fc FodderCalculator, numCows int) (_ float64, err error) {
	fodderAmount, err := fc.FodderAmount(numCows)
	if err != nil {
		return 0, err
	}
	factor, err := fc.FatteningFactor()
	if err != nil {
		return 0, err
	}
	return fodderAmount * factor / float64(numCows), nil
}

func ValidateInputAndDivideFood(fc FodderCalculator, numCows int) (float64, error) {
	if numCows < 1 {
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(fc, numCows)
}

func ValidateNumberOfCows(numCows int) error {
	if numCows < 0 {
		return &InvalidCowsError{NumCows: numCows, Message: "there are no negative cows"}
	}
	if numCows == 0 {
		return &InvalidCowsError{NumCows: numCows, Message: "no cows don't need food"}
	}
	return nil
}
