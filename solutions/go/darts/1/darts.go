package darts

import "math"

func Score(x, y float64) int {
	position := math.Sqrt((x * x) + (y * y))
   	if position > 10 {
        return 0
    } else if position > 5 {
        return 1
    } else if position > 1 {
        return 5
    }
    return 10
}