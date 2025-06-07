package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	const welcomeMessage string = "Welcome to the Tech Palace, "
	return welcomeMessage + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	var border string = strings.Repeat("*", numStarsPerLine)
	return border + "\n" + welcomeMsg + "\n" + border
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {

	return strings.TrimSpace(strings.ReplaceAll(oldMsg, "*", ""))
}
