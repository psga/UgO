package main

import (
	"fmt"
	"math/rand"
	"time"
)

type card struct {
	color  string
	number int8
}
type deck struct {
	cards []card
}
type player struct {
	name string
	deck deck
}

// func (obj *player) Default() {
// 	if obj.deck == nil {
// 		obj.Name = "John Doe"
// 	}
// }

type game struct {
	id      string
	players []player
}

//func draw() {
//	var mainDeck deck
//
//	for _, c := range []string{"r", "g", "y", "b"} {
//		for _, n := range []int8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, -10, -11, -2} {
//			mainDeck.cards = append(mainDeck.cards, card{c, n})
//		}
//	}
//	fmt.Println(mainDeck.cards)
//	randomIndex := rand.Intn(len(mainDeck.cards))
//	pick := mainDeck.cards[randomIndex]
//	fmt.Println(pick)
//
//	// deck mainDeck :=
//}

func generatePlayer(name string) player {

	fmt.Println("Player " + name + " generated")
	// var
	var newPlayer player = player{name: name}
	return newPlayer
}

func pickCard(azar float64, colorsNum []int8, numberCards int8, colors []string) card {

	switch {
	case azar < (float64(colorsNum[0]) / float64(numberCards)):
		return card{"r", 1}
	case azar < (float64(colorsNum[0]+colorsNum[1]) / float64(numberCards)):
		return card{"y", 1}
	case azar < (float64(colorsNum[0]+colorsNum[1]+colorsNum[2]) / float64(numberCards)):
		colors[0] = "blue"
		return card{"b", 1}
	case azar < (float64(colorsNum[0]+colorsNum[1]+colorsNum[2]+colorsNum[3]) / float64(numberCards)):
		colors[0] = "green" // ...
		return card{"g", 1}
	case azar < (float64(colorsNum[0]+colorsNum[1]+colorsNum[2]+colorsNum[3]+colorsNum[4]) / float64(numberCards)):
		colors[0] = "black" // .
		return card{"x", 1}
	default:
		return card{"nunca", 1}
	}

}

func generateGame(nplayers int) {
	//Because only one deck per game needs to be generated.
	rand.Seed(time.Now().UnixNano())
	id := "1"
	game := game{id: id}
	fmt.Println("Game with id " + id + " started")

	//for i := 0;i<nplayers;i++{
	//	fmt.Println("whats your name player ",i+1)
	//	var name string
	//	fmt.Scanln(&name)
	//	game.players=append(game.players,player{name:name})
	//}
	game.players = []player{}
	game.players = append(game.players, generatePlayer("Julian"))
	game.players = append(game.players, generatePlayer("Pablo"))

	// Initial Deck variables to deal cards
	numberCards := int8(108)
	colors := []string{"r", "y", "b", "g", "x"}
	colorsNum := []int8{25, 25, 25, 25, 8}
	numbersR := []card{{"0", 1}, {"1", 2}, {"2", 3}, {"3", 4}, {"4", 2}, {"5", 2}, {"6", 2},
		{"7", 2}, {"8", 2}, {"9", 2}, {"turno", 2}, {"or", 2}, {"joder", 2}}
	numbersB := make([]card, len(numbersR))
	numbersG := make([]card, len(numbersR))
	numbersY := make([]card, len(numbersR))
	copy(numbersB, numbersR)
	copy(numbersG, numbersR)
	copy(numbersY, numbersR)

	//Deal cards for each player
	for i := range game.players {

		for c := 0; c < 7; c++ {
			//Deal 7 cards for each player
			azar := rand.Float64()
			game.players[i].deck.cards = append(game.players[i].deck.cards, pickCard(azar, colorsNum, numberCards, colors))

		}
	}
	fmt.Println("===== Game generated ======")
	fmt.Println(game.players)
	fmt.Println("===== ============== ======")

	//drawCard()
	// numbers = numbers
	// Rcards=Gcards=

	//rand.Intn(max - min
}

func main() {
	generateGame(2)

	// draw()

}
