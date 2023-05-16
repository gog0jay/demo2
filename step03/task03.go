package step03

import (
	"fmt"
	"pacgo/step01"
	"pacgo/utils/simpleansi"
)

// Updating the maze
// refactor the printScreen()
func printScreen() {
	simpleansi.ClearScreen()
	for _, line := range step01.Maze {
		for _, chr := range line {
			switch chr {
			case '#':
				fmt.Printf("%c", chr)
			default:
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	simpleansi.MoveCursor(player.row, player.col)
	fmt.Print("P")

	simpleansi.MoveCursor(len(step01.Maze)+1, 0)
}
