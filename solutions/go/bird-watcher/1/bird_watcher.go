package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	totalBirds := 0
    for _, v := range birdsPerDay {
        totalBirds += v
    }
    return totalBirds
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
    start := (week - 1) * 7  // Convert to 0-indexed
    end := start + 7
    
    if start >= len(birdsPerDay) {
        return 0
    }
    if end > len(birdsPerDay) {
        end = len(birdsPerDay)
    }
    
    birdsPerWeek := 0
    for _, v := range birdsPerDay[start:end] {
        birdsPerWeek += v
    }
    return birdsPerWeek
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
    for i := range birdsPerDay {
        if i % 2 == 0 {
            birdsPerDay[i] += 1 
        }
    }
    return birdsPerDay
}
