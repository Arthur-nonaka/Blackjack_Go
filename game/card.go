package game

type Card struct {
	Number string
	Suit   string
	// Played bool
	// Owner  string
}

func NewCard(Number string, Suit string) *Card {
	return &Card{
		Number: Number,
		Suit:   Suit,
		// Played: false,
		// Owner:  "None",
	}
}
