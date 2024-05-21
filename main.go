package main

import (
	"fmt"
	"strconv"

	"github.com/Arthur-Nonaka/blackjack/game"
)

func main() {
	deck := game.Deck{}
	deck.Create()
	deck.Print()
	deck.Shuffle()
	deck.Print()

	player := game.Player{}
	players := player.StartPlayers()
	dealer := game.Player{Name: "Dealer", Points: 0}
	deck.TakeNewCard(&dealer)

	for _, player := range players {
		deck.TakeNewCard(player)
		player.PlayerTurn(deck)
	}

	dealer.DealerTurn(deck)

	for _, player := range players {
		str := ""
		if player.Points > 22 {
			str = player.Name + " PERDEU - " + strconv.Itoa(player.Points) + " Pontos"
		} else if dealer.Points > 21 {
			str = player.Name + " GANHOU - " + strconv.Itoa(player.Points) + " Pontos"
		} else if player.Points > dealer.Points {
			str = player.Name + " GANHOU - " + strconv.Itoa(player.Points) + " Pontos"
		} else if player.Points < dealer.Points {
			str = player.Name + " PERDEU - " + strconv.Itoa(player.Points) + " Pontos"
		} else {
			str = player.Name + " EMPATOU - " + strconv.Itoa(player.Points) + " Pontos"
		}
		fmt.Println(str)
	}
	dealer.ShowInfo()
}
