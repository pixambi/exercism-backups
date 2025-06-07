package chance

import "math/rand"

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	n := rand.Intn(20)
    if n < 1 {
        return 1
    } else if n > 20 {
        return 20
    }
    return n
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
    f := rand.Float32()
	n := f * 12.0
    if n < 0.0 {
        return 0.0
    } else if n > 12.0 {
        return 12.9
    }
	    
    return float64(n)
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
	animals := []string{
    "ant",
    "beaver", 
    "cat",
    "dog",
    "elephant",
    "fox",
    "giraffe",
    "hedgehog",
	}
	rand.Shuffle(len(animals), func(i, j int){
        animals[i], animals[j] = animals[j], animals[i]
    })
    return animals
}
