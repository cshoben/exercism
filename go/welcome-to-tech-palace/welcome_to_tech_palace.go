package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	welcome := "Welcome to the Tech Palace, "
	return welcome + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	var stars string
	stars = strings.Repeat("*", numStarsPerLine)
	return stars + "\n" + welcomeMsg + "\n" + stars
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	remStarNewLine := strings.Trim(oldMsg, "*\n")
	removeSpace := strings.TrimSpace(remStarNewLine)
	return removeSpace
}
