package techpalace

import (
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	strStars := strings.Repeat("*", numStarsPerLine)
	return strStars + "\n" + welcomeMsg + "\n" + strStars
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	lines := strings.Split(oldMsg, "\n")
	if len(lines) >= 3 {
		// Remove leading and trailing asterisks and trim whitespace
		cleanedMessage := strings.TrimSpace(strings.Trim(lines[1], "*"))
		return cleanedMessage
	}
	return oldMsg
}
