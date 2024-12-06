package main

import (
	"github.com/RaresVirjoghe/AdventOfCode2024/day1/pkg/handleinput"
	"github.com/RaresVirjoghe/AdventOfCode2024/day1/pkg/solve"
	"fmt"
)

func main() {
	col1, col2 := handleinput.ReadInput("input.txt")

	fmt.Println(solve.CalculateDistance(col1, col2))
	fmt.Println(solve.CalculateSimilarity(col1, col2))
}
