package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	onetoone := map[string]int{}
    for point, letters := range in {
        for _, letter := range letters {
            onetoone[strings.ToLower(letter)] = point
        }
    }
    return onetoone
}
