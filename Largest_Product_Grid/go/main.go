package main

import "fmt"

var (
	row1 = [20]int{
		8,
		2,
		22,
		97,
		38,
		15,
		0,
		40,
		0,
		75,
		4,
		5,
		7,
		78,
		52,
		12,
		50,
		77,
		91,
		8,
	}

	row2 = [20]int{
		49,
		49,
		99,
		40,
		17,
		81,
		18,
		57,
		60,
		87,
		17,
		40,
		98,
		43,
		69,
		48,
		4,
		56,
		62,
		0,
	}

	row3 = [20]int{
		81,
		49,
		31,
		73,
		55,
		79,
		14,
		29,
		93,
		71,
		40,
		67,
		53,
		88,
		30,
		3,
		49,
		13,
		36,
		65,
	}

	row4 = [20]int{
		52,
		70,
		95,
		23,
		4,
		60,
		11,
		42,
		69,
		24,
		68,
		56,
		1,
		32,
		56,
		71,
		37,
		2,
		36,
		91,
	}
	row5 = [20]int{
		22,
		31,
		16,
		71,
		51,
		67,
		63,
		89,
		41,
		92,
		36,
		54,
		22,
		40,
		40,
		28,
		66,
		33,
		13,
		80,
	}

	row6 = [20]int{
		24,
		47,
		32,
		60,
		99,
		3,
		45,
		2,
		44,
		75,
		33,
		53,
		78,
		36,
		84,
		20,
		35,
		17,
		12,
		50,
	}
	row7 = [20]int{
		32,
		98,
		81,
		28,
		64,
		23,
		67,
		10,
		26,
		38,
		40,
		67,
		59,
		54,
		70,
		66,
		18,
		38,
		64,
		70,
	}

	row8 = [20]int{67, 26, 20, 68, 2, 62, 12, 20, 95, 63, 94, 39, 63, 8, 40, 91, 66, 49, 94, 21}

	row9 = [20]int{
		24,
		55,
		58,
		5,
		66,
		73,
		99,
		26,
		97,
		17,
		78,
		78,
		96,
		83,
		14,
		88,
		34,
		89,
		63,
		72,
	}

	row10 = [20]int{
		21,
		36,
		23,
		9,
		75,
		0,
		76,
		44,
		20,
		45,
		35,
		14,
		0,
		61,
		33,
		97,
		34,
		31,
		33,
		95,
	}
	row11 = [20]int{
		78,
		17,
		53,
		28,
		22,
		75,
		31,
		67,
		15,
		94,
		3,
		80,
		4,
		62,
		16,
		14,
		9,
		53,
		56,
		92,
	}

	row12 = [20]int{
		16,
		39,
		5,
		42,
		96,
		35,
		31,
		47,
		55,
		58,
		88,
		24,
		0,
		17,
		54,
		24,
		36,
		29,
		85,
		57,
	}
	row13 = [20]int{
		86,
		56,
		0,
		48,
		35,
		71,
		89,
		7,
		5,
		44,
		44,
		37,
		44,
		60,
		21,
		58,
		51,
		54,
		17,
		58,
	}

	row14 = [20]int{
		19,
		80,
		81,
		68,
		5,
		94,
		47,
		69,
		28,
		73,
		92,
		13,
		86,
		52,
		17,
		77,
		4,
		89,
		55,
		40,
	}
	row15 = [20]int{
		4,
		52,
		8,
		83,
		97,
		35,
		99,
		16,
		7,
		97,
		57,
		32,
		16,
		26,
		26,
		79,
		33,
		27,
		98,
		66,
	}

	row16 = [20]int{
		88,
		36,
		68,
		87,
		57,
		62,
		20,
		72,
		3,
		46,
		33,
		67,
		46,
		55,
		12,
		32,
		63,
		93,
		53,
		69,
	}
	row17 = [20]int{
		4,
		42,
		16,
		73,
		38,
		25,
		39,
		11,
		24,
		94,
		72,
		18,
		8,
		46,
		29,
		32,
		40,
		62,
		76,
		36,
	}
	row18 = [20]int{
		20,
		69,
		36,
		41,
		72,
		30,
		23,
		88,
		34,
		62,
		99,
		69,
		82,
		67,
		59,
		85,
		74,
		4,
		36,
		16,
	}
	row19 = [20]int{
		20,
		73,
		35,
		29,
		78,
		31,
		90,
		1,
		74,
		31,
		49,
		71,
		48,
		86,
		81,
		16,
		23,
		57,
		5,
		54,
	}
	row20 = [20]int{
		1,
		70,
		54,
		71,
		83,
		51,
		54,
		69,
		16,
		92,
		33,
		48,
		61,
		43,
		52,
		1,
		89,
		19,
		67,
		48,
	}
	matrix [20][20]int = [20][20]int{
		row1,
		row2,
		row3,
		row4,
		row5,
		row6,
		row7,
		row8,
		row9,
		row10,
		row11,
		row12,
		row13,
		row14,
		row15,
		row16,
		row17,
		row18,
		row19,
		row20,
	}
)

const (
	consecutiveNum = 4
	matrixLen      = 20
)

func findMax(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func travelRight(currentMax *int) {
	for _, row := range matrix {
		for col := 0; col < matrixLen-consecutiveNum; col++ {
			product := 1

			for index := 0; index < 4; index++ {
				product *= row[col+index]
			}
			*currentMax = findMax(product, *currentMax)
		}
	}
}

func travelDown(currentMax *int) {
	for col := 0; col < matrixLen; col++ {
		for row := 0; row < matrixLen-consecutiveNum; row++ {
			product := 1

			for index := 0; index < consecutiveNum; index++ {
				product *= matrix[row+index][col]
			}

			*currentMax = findMax(product, *currentMax)
		}
	}
}

func travelDiagonalRight(currentMax *int) {
	for row := 0; row < matrixLen-consecutiveNum; row++ {
		for col := 0; col < matrixLen-consecutiveNum; col++ {
			product := 1
			for index := 0; index < consecutiveNum; index++ {
				product *= matrix[row+index][col+index]
			}

			*currentMax = findMax(product, *currentMax)
		}
	}
}

func travelDiagonalLeft(currentMax *int) {
	for row := 0; row < matrixLen-consecutiveNum; row++ {
		for col := consecutiveNum - 1; col < matrixLen; col++ {
			product := 1

			for index := 0; index < consecutiveNum; index++ {
				product *= matrix[row+index][col-index]
			}

			*currentMax = findMax(product, *currentMax)
		}
	}
}

func main() {
	var currentMax int

	travelRight(&currentMax)
	travelDown(&currentMax)
	travelDiagonalRight(&currentMax)
	travelDiagonalLeft(&currentMax)

	fmt.Printf("Result: %d\n", currentMax)
}
