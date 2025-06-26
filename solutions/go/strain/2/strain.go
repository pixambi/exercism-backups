package strain

// Implement the "Keep" and "Discard" function in this file.
func Keep[T any](collection []T, predicate func(v T)bool) []T{
	filtered := []T{}
    for _, v := range collection{
        if predicate(v) {
            filtered = append(filtered, v)
        }
    }
    return filtered
}

func Discard[T any](collection []T, predicate func(v T)bool) []T{
    return Keep(collection, func(v T)bool {return !predicate(v) })
}
// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1
