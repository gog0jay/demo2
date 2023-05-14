package step01

import "fmt"

func PrintScreen(Maze []string) {
	for _, line := range Maze {
		fmt.Println(line)
	}

}
