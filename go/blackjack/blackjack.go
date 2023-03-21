package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five" :
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten", "jack", "queen", "king":
		return 10
	default :
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	var turn string
	var dealerReveal bool
	playerValue := ParseCard(card1) + ParseCard(card2)
	if ParseCard(dealerCard) == 10 || ParseCard(dealerCard) == 11 {
		dealerReveal = true
	}

	switch {
	// if pair of aces, must always split(P) them
	case ParseCard(card1) == 11 && ParseCard(card2) == 11:
		turn = "P"
	// if player has black jack but dealer has Ace, face card, or 10, player must Stand(S)
	case playerValue == 21 && dealerReveal == true:
		turn = "S"
	// if player has black jack but dealer DOES NOT have Ace, face card, or 10, player Wins(W)!
	case playerValue == 21 && dealerReveal == false:
		turn = "W"
	// if player cards cume between 17 and 20, must Stand(S)
	case playerValue >=17 && playerValue <= 20:
		turn = "S"
	}
	return turn
}
