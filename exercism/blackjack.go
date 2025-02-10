package main

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	var result int
	switch card {
	case "two":
		result = 2
	case "three":
		result = 3
	case "four":
		result = 4
	case "five":
		result = 5
	case "six":
		result = 6
	case "seven":
		result = 7
	case "eight":
		result = 8
	case "nine":
		result = 9
	case "ten":
		result = 10
	case "jack":
		result = 10
	case "queen":
		result = 10
	case "king":
		result = 10
	case "ace":
		result = 11
	default:
		result = 0
	}
	return result
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	var result string
	sum := ParseCard(card1) + ParseCard(card2)
	dealer := ParseCard(dealerCard)
	switch {
	case card1 == "ace" && card2 == "ace":
		result = "P"
	case sum == 21 && dealer < 10:
		result = "W"
	case sum >= 17 && sum <= 21:
		result = "S"
	case (sum >= 12 && sum <= 16 && dealer >= 7) || sum <= 11:
		result = "H"
	default:
		result = "S"
	}
	return result
}
