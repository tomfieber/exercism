package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	value := 0
	switch card {
	case "ace":
		value = 11
	case "two":
		value = 2
	case "three":
		value = 3
	case "four":
		value = 4
	case "five":
		value = 5
	case "six":
		value = 6
	case "seven":
		value = 7
	case "eight":
		value = 8
	case "nine":
		value = 9
	case "ten", "jack", "queen", "king":
		value = 10
	default:
		value = 0
	}
	return value
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	card1Value := ParseCard(card1)
	card2Value := ParseCard(card2)
	dealerCardValue := ParseCard(dealerCard)
	playerValue := card1Value + card2Value

	switch {
	case playerValue == 22:
		return "P"
	case playerValue == 21 && dealerCardValue < 10:
		return "W"
	case playerValue == 21 && dealerCardValue >= 10:
		return "S"
	case playerValue >= 17 && playerValue <= 20:
		return "S"
	case playerValue >= 12 && playerValue <= 16 && dealerCardValue < 7:
		return "S"
	case playerValue >= 12 && playerValue <= 16 && dealerCardValue >= 7:
		return "H"
	case playerValue <= 11:
		return "H"
	default:
		return "H"
	}
}
