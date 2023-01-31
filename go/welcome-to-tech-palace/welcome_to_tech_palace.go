package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	welcome := "Welcom to the Tech Palace, "
	return welcome + customer
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	var stars string
	stars = "*"
	stars = "*" * numStarsPerLine
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	panic("Please implement the CleanupMessage() function")
}
