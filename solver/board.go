package solver

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

const (
	Row = 9
	Col = 9
)

// invalid array is a hash that holds true if the value is already present in the row or col
type InvalidArray map[int]bool

// this struch will have an array that will store the size of the board
type Board struct {
	Cells        [Row][Col]int
	InvalidInRow [Row]InvalidArray
	InvalidInCol [Col]InvalidArray
}

func GetBoardFromInput(source io.Reader) (error, *Board) {

	board := Board{}
	scanner := bufio.NewScanner(source)
	row := 0
	for ; scanner.Scan() && row < Row; row++ {
		// splitgin with comma trimming for space
		curRow := strings.Split(scanner.Text(), ",")
		// function to catch if length is less than or greater than Row
		for col, val := range curRow {
			temp, err := getNumber(val)
			if err != nil {
				return err, nil
			} else {
				board.Cells[row][col] = temp
			}
		}
	}

	return nil, &board

}
func getNumber(value string) (int, error) {

	return strconv.Atoi(value)
}
