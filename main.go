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
type game struct {
	id string
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

func generateGame(nplayers int) {  //Because only one deck per game needs to be generated.
	rand.Seed(time.Now().UnixNano())
	game := game{id:"1",players: []player{}}

	//for i := 0;i<nplayers;i++{
	//	fmt.Println("whats your name player ",i+1)
	//	var name string
	//	fmt.Scanln(&name)
	//	game.players=append(game.players,player{name:name})
	//}
	game.players=[]player{player{name:"Julian"},player{name:"pablo"}}
	fmt.Println(game)


	// VARIABLRES
	numberCards := int8(108)
	colors := []string{"r", "y", "b", "g", "x"}
	colorsNum := []int8{25, 25, 25, 25, 8}
	numbersR := []card{	card{"0", 1}, card{"1", 2}, card{"2", 3},card{"3", 4}, card{"4", 2}, card{"5", 2}, card{"6", 2},
				card{"7", 2}, card{"8", 2}, card{"9", 2}, card{"turno", 2}, card{"or", 2}, card{"joder", 2}}
	numbersB := make([]card, len(numbersR))
	numbersG := make([]card, len(numbersR))
	numbersY := make([]card, len(numbersR))
	copy(numbersB, numbersR)
	copy(numbersG, numbersR)
	copy(numbersY, numbersR)

	//Deal cards
	for _,i := range game.players{
		fmt.Println(i)
		for i:=0;i<7;i++ {
			Azar :=rand.Float64()
			fmt.Println(Azar)
			switch {
			case Azar<(float64(colorsNum[0])/float64(numberCards)):
				colors[0] = "red"
				//fmt.Println(colors[0])
				//fmt.Println(Azar)



			}

		}
	}



	//drawCard()
	// numbers = numbers
	// Rcards=Gcards=

	//rand.Intn(max - min
}

func main() {
	generateGame(2)

	// draw()

}
