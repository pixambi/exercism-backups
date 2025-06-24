package space

type Planet string

func Age(seconds float64, planet Planet) float64 {
	earth := 365.25 * 24 * 60 * 60 
	earthYear := seconds / earth
	
	switch planet {
	case "Mercury":
		return earthYear / 0.2408467
	case "Venus":
		return earthYear / 0.61519726
	case "Earth":
		return earthYear / 1.0
	case "Mars":
		return earthYear / 1.8808158
	case "Jupiter":
		return earthYear / 11.862615
	case "Saturn":
		return earthYear / 29.447498
	case "Uranus":
		return earthYear / 84.016846
	case "Neptune":
		return earthYear / 164.79132
	default:
		return -1.000000
	}
}