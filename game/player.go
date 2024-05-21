package game

import (
	"fmt"
)

type Player struct {
	Name   string
	Points int
	Hand   []Card
}

func NewPlayer(Name string) *Player {
	return &Player{
		Name:   Name,
		Points: 0,
	}
}

func (p *Player) StartPlayers() []*Player {
	n := 0
	for {
		fmt.Println("\nNumero de Jogadores: ")
		fmt.Scan(&n)
		if n < 2 || n > 12 {
			fmt.Println("Digite um numero valido")
			continue
		}
		break
	}
	players := []*Player{}
	for i := 0; i < n; i++ {
		name := ""
		fmt.Print("Nome: ")
		fmt.Scan(&name)
		player := *NewPlayer(name)
		players = append(players, &player)
	}
	return players
}

func (p *Player) ShowInfo() {
	fmt.Printf("Jogador %v\n", p.Name)
	fmt.Print("Cartas: ")
	for _, card := range p.Hand {
		fmt.Printf("| %v %v", card.Number, card.Suit)
	}
	fmt.Printf("\nPontos: %d\n", p.Points)
}

func (p *Player) PlayerTurn(deck Deck) {
	resp := 1
	for resp == 1 {
		p.ShowInfo()
		fmt.Println("1 - Puxar mais uma Carta 2 - Parar")
		fmt.Scan(&resp)
		if resp != 1 && resp != 2 {
			fmt.Println("Escreve direito")
			resp = 1
			continue
		}
		if resp == 1 {
			deck.TakeNewCard(p)
		}
		if p.Points >= 22 {
			p.ShowInfo()
			fmt.Println("Perdeu")
			break
		}
	}
}

func (p *Player) DealerTurn(deck Deck) {
	for {
		if p.Points < 17 {
			deck.TakeNewCard(p)
		} else {
			break
		}
	}
	if p.Points >= 22 {
		fmt.Println("Perdeu")
	}
	// p.ShowInfo()
}

// func (p *Player) EndGame(players []*Player, dealer Player) {
// 	for _, player := range players {
// 		str := ""
// 		if player.Points > 22 {
// 			str = player.Name + " PERDEU - " + strconv.Itoa(player.Points) + " Pontos"
// 		} else if dealer.Points > 22 {
// 			str = player.Name + " GANHOU - " + strconv.Itoa(player.Points) + " Pontos"
// 		} else if player.Points > dealer.Points {
// 			str = player.Name + " GANHOU - " + strconv.Itoa(player.Points) + " Pontos"
// 		} else if player.Points < dealer.Points {
// 			str = player.Name + " PERDEU - " + strconv.Itoa(player.Points) + " Pontos"
// 		} else {
// 			str = player.Name + " EMPATOU - " + strconv.Itoa(player.Points) + " Pontos"
// 		}
// 		fmt.Println(str)
// 	}
// 	dealer.ShowInfo()
// }
