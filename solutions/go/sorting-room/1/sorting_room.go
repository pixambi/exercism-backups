package sorting

import (
    "fmt"
    "strconv"
)
    

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %.1f", float32(nb.Number()))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	numBox, ok := fnb.(FancyNumber)
    if !ok{
        return 0
    }
    val, _ := strconv.Atoi(numBox.Value())
    return val
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
    val := ExtractFancyNumber(fnb) 
	if val == 0 {
        return fmt.Sprintf("This is a fancy box containing the number 0.0")
    }
    return fmt.Sprintf("This is a fancy box containing the number %.1f", float32(val))
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	switch v := i.(type){
        case int:
        	return DescribeNumber(float64(v))
        case float64:
        	return DescribeNumber(v)
    	case NumberBox:
    		return DescribeNumberBox(v)
    	case FancyNumberBox:
        	return DescribeFancyNumberBox(v)
        default:
        	return "Return to sender"
    }
}
