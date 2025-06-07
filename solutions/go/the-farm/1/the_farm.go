package thefarm

import (
 	"fmt"
    "errors"
)

// TODO: define the 'DivideFood' function
func DivideFood(fodderCalculator FodderCalculator, numCows int) (float64, error){
    totalFodder, err := fodderCalculator.FodderAmount(numCows)
    if err != nil{
        return 0.0, err
    }
    fatteningFactor, err := fodderCalculator.FatteningFactor()
    if err != nil{
        return 0.0, err
    }
    return (totalFodder * fatteningFactor) / float64(numCows), nil
}
// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fodderCalculator FodderCalculator, numCows int) (float64, error ){
     if numCows <= 0{
        return 0.0, errors.New("invalid number of cows")
    }
    food, err := DivideFood(fodderCalculator, numCows)
    if err != nil{
        return 0.0, err
    }
    return food, nil
}

type InvalidCowsError struct {
    numCows int
    message string
}
func (e *InvalidCowsError) Error() string {
    return fmt.Sprintf("%d cows are invalid: %s", e.numCows, e.message)
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(numCows int)error {
    if numCows < 0 {
        return &InvalidCowsError{
            numCows: numCows,
    		message: "there are no negative cows",
  		}
    }
    if numCows == 0 {
        return &InvalidCowsError{
            numCows: numCows,
    		message: "no cows don't need food",
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
