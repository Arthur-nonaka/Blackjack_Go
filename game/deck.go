package game

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

type Deck struct {
	Cards []Card
}

func (d *Deck) Print() {
	str := "\n---------------------------------------------\n"
	for idx, card := range d.Cards {

		if idx != 0 && idx%13 == 0 {
			str += "\n"
		}
		str += fmt.Sprintf("| %v %v ", card.Number, card.Suit)

	}
	str += "\n---------------------------------------------"
	fmt.Print(str)
}

func (d *Deck) Create() {
	Suits := [4]string{"♠️", "♥️", "♦️", "♣️"}
	Numbers := [13]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Q", "J", "K"}

	for _, Number := range Numbers {
		for _, Suit := range Suits {
			d.Cards = append(d.Cards, Card{
				Number: Number,
				Suit:   Suit,
				// Played: false
			})
		}
	}
}

func (d *Deck) Shuffle() {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	random := r.Intn(300)
	random += 200
	i := 0
	for i < random {
		pos1 := r.Intn(52)
		pos2 := r.Intn(52)

		holder := d.Cards[pos1]
		d.Cards[pos1] = d.Cards[pos2]
		d.Cards[pos2] = holder
		i++
	}
}

func (d *Deck) TakeNewCard(player *Player) {
	takenCard := d.Cards[len(d.Cards)-1]
	d.Cards = d.Cards[:len(d.Cards)-1]

	player.Hand = append(player.Hand, takenCard)
	if takenCard.Number == "K" || takenCard.Number == "Q" || takenCard.Number == "J" {
		takenCard.Number = "10"
	}
	number, err := strconv.Atoi(takenCard.Number)
	if err != nil {
		log.Fatal(err)
	}
	player.Points += number
}
