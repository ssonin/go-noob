package main

import (
	"fmt"
	"strings"
)

func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	border := strings.Repeat("*", numStarsPerLine)
	return border + "\n" + welcomeMsg + "\n" + border
}

func CleanupMessage(oldMsg string) string {
	return strings.TrimSpace(strings.ReplaceAll(oldMsg, "*", ""))
}

func main() {
	fmt.Println(WelcomeMessage("Monica"))
	fmt.Println(AddBorder("Whatever", 10))
	oldMsg := `
**************************
*    BUY NOW, SAVE 10%   *
**************************
`
	fmt.Println(CleanupMessage(oldMsg))
}
