package functions

//Las funciones y variables necesitan tener la primera letra
//mayuscula para poder llamarlas en donde se importen
func PickCardColor(azar float64, colorsNum []int8, numberCards int8, colors []string) string {

	switch {
	case azar < (float64(colorsNum[0]) / float64(numberCards)):
		return "r"
	case azar < (float64(colorsNum[0]+colorsNum[1]) / float64(numberCards)):
		return "y"
	case azar < (float64(colorsNum[0]+colorsNum[1]+colorsNum[2]) / float64(numberCards)):
		return "b"
	case azar < (float64(colorsNum[0]+colorsNum[1]+colorsNum[2]+colorsNum[3]) / float64(numberCards)):
		colors[0] = "green" // ...
		return "g"
	case azar < (float64(colorsNum[0]+colorsNum[1]+colorsNum[2]+colorsNum[3]+colorsNum[4]) / float64(numberCards)):
		return "x"
	default:
		return "neverland"
	}

}

// This function should receive the number of cards based
//on previous function call, "pickCardColor"

//	cardslist := []int8{1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}
func PickCardNumber(azar float64, cardList []int8, numberCards int8) int8 {
	porcentages := []float64{}
	for i, x := range cardList {
		if i == 0 {
			porcentages = append(porcentages, (float64(x) / float64(numberCards)))
		} else {
			porcentages = append(porcentages, (float64(x)/float64(numberCards))+float64(porcentages[len(porcentages)-1]))
		}
	}
	//fmt.Println("porcentage of",porcentages)
	//fmt.Println("azar",azar)

	for i, x := range porcentages {
		if azar < x {
			return int8(i)
		}

	}
	return int8(len(porcentages) - 1)
}
