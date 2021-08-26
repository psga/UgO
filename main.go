package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"ugo.com/functions"
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
	id      string
	players []player
}

func generatePlayer(name string) player {

	fmt.Println("Player " + name + " generated")
	// var
	var newPlayer player = player{name: name}
	return newPlayer
}

func generateGame(nplayers int) {

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
	game.players = append(game.players, generatePlayer("Juan"))

	// Initial Deck variables to deal cards
	numberCards := int8(108)
	colors := []string{"r", "y", "b", "g", "x"}
	colorsNum := []int8{25, 25, 25, 25, 8}
	// "0" = 1 Card
	// "1-9" = 2 Cards
	// "Turn, Reverse, +2" = 2 Cards

	cardsRed := []int8{1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}
	cardsBlue := make([]int8, len(cardsRed))
	cardsGreen := make([]int8, len(cardsRed))
	cardsYellow := make([]int8, len(cardsRed))
	copy(cardsBlue, cardsRed)
	copy(cardsGreen, cardsRed)
	copy(cardsYellow, cardsRed)
	cardsBlack := []int8{4, 4}

	//Deal cards for each player
	for i := range game.players {

		for c := 0; c < 7; c++ {
			//Deal 7 cards for each player
			azar := rand.Float64()
			cardColor := functions.PickCardColor(azar, colorsNum, numberCards, colors)
			// cardColor := pickCardColor(azar, colorsNum, numberCards, colors)
			numberCards--

			var cardNumber int8
			azar = rand.Float64()
			fmt.Println(cardColor)

			switch {
			case cardColor == "r":
				cardNumber = functions.PickCardNumber(azar, cardsRed, colorsNum[0])
				colorsNum[0]--
				cardsRed[cardNumber]--
			case cardColor == "b":
				cardNumber = functions.PickCardNumber(azar, cardsBlue, colorsNum[1])
				colorsNum[1]--
				cardsBlue[cardNumber]--
			case cardColor == "g":
				cardNumber = functions.PickCardNumber(azar, cardsGreen, colorsNum[2])
				colorsNum[2]--
				cardsGreen[cardNumber]--
			case cardColor == "y":
				cardNumber = functions.PickCardNumber(azar, cardsYellow, colorsNum[3])
				colorsNum[3]--
				cardsYellow[cardNumber]--
			case cardColor == "x":
				cardNumber = functions.PickCardNumber(azar, cardsBlack, colorsNum[4])
				colorsNum[4]--
				cardsBlack[cardNumber]--
			}

			game.players[i].deck.cards = append(game.players[i].deck.cards, card{cardColor, cardNumber})

		}
	}
	fmt.Println("===== Game generated ======")
	fmt.Println(game.players)
	fmt.Println("===== ============== ======")

	fmt.Println("Number of Cards of each color:", colorsNum)
	fmt.Println("===== ============== ======")
	//drawCard()
	// numbers = numbers
	// Rcards=Gcards=

	//rand.Intn(max - min
}

func main() {
	generateGame(2)

	//Handle requests

	fmt.Println("Now Listening on 8081")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)

		p := "./client" + r.URL.Path
		fmt.Println(p)
		if p == "./client/" {
			p = "./client/index.html"
		}
		http.ServeFile(w, r, p)
	})

	//Fatal is equivalent to Print() followed by a call to os.Exit(1).
	log.Fatal(http.ListenAndServe(":8081", nil))

}
