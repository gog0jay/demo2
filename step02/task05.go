package step02

import "fmt"

import "github.com/danicat/simpleansi"

// update printScreen() in stop01--task02

func PrintScreen(Maze []string) {
	simpleansi.ClearScreen()
	for _, line := range Maze {
		fmt.Println(line)
	}

}
