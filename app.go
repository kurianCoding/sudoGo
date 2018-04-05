package main

import (
	"fmt"
	"os"

	"github.com/kurianCoding/sudoGo/solver"
)

func main() {
	err, board := solver.GetBoardFromInput(os.Stdin)
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%d ", board.Cells[i][j])
		}
		fmt.Println()
	}
	//check all the columns for digits to exclude
	//check all the rows for digits to exclude
	//check each 9x9 cell for digits to exclude
	//cycle through the remaining digits
	//if solution is possible, print solved
	//if solution is not possible, print unsolvable
	return
}
