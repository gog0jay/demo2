package main

import "pacgo/step01"

func main() {
	err := step01.LoadMaze("resources/maze01.txt")
	if err != nil {
		return
	}

	step01.PrintScreen(step01.Maze)

}
