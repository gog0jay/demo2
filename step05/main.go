package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"pacgo/utils/simpleansi"
)

var maze []string

type sprite struct {
	row int
	col int
}

var player sprite

var ghosts []*sprite

var score int
var numDots int
var lives = 1

func loadMaze(file string) error {

	// os.Open读取文件，有两个返回值
	f, err := os.Open(file)

	// error handlingmain
	if err != nil {
		return err
	}

	// put f.Close() in the Call stack
	defer f.Close()

	// read the file line by line and appends it to the maze slice
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		maze = append(maze, line)
	}

	// capture the player position
	for row, line := range maze {
		for col, char := range line {
			switch char {
			case 'P':
				player = sprite{row, col}
			case 'G':
				ghosts = append(ghosts, &sprite{row, col})
			case '.':
				numDots++
			}
		}
	}
	return nil
	// f.Close() is called implicitly

}

func printScreen() {
	simpleansi.ClearScreen()
	for _, line := range maze {
		for _, chr := range line {
			switch chr {
			case '#':
				//fmt.Printf("%c", chr)
				fallthrough
			case '.':
				fmt.Printf("%c", chr)
			default:
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	simpleansi.MoveCursor(player.row, player.col)
	fmt.Print("P")

	for _, g := range ghosts {
		simpleansi.MoveCursor(g.row, g.col)
		fmt.Print("G")
	}

	simpleansi.MoveCursor(len(maze)+1, 0)
	fmt.Println("Score: ", score, "\tLives: ", lives)
}

func initialise() {
	cbTerm := exec.Command("stty", "cbreak", "-echo")
	cbTerm.Stdin = os.Stdin

	err := cbTerm.Run()
	if err != nil {
		log.Fatalln("unable to activate cbreak mode :", err)
	}
}

func cleanup() {
	cookedTerm := exec.Command("stty", "-cbreak", "echo")
	cookedTerm.Stdin = os.Stdin

	err := cookedTerm.Run()
	if err != nil {
		log.Fatalln("unable to restore cooked mode: ", err)
	}
}

func readInput() (string, error) {
	// 创建一个buffer指针指向一个大小为100的字节切片
	buffer := make([]byte, 100)

	// os.Stdin.Read() 有两个返回值： 读到的字节数 和 错误值
	cnt, err := os.Stdin.Read(buffer)
	if err != nil {
		return "", err
	}

	// 只读到了一个字节，且这个字节是退出键 (0x1b表示Esc)
	if cnt == 1 && buffer[0] == 0x1b {
		return "ESC", nil
	} else if cnt >= 3 {
		// The escape sequence for the arrow keys are 3 bytes long,
		//starting with ESC+[ and then a letter from A to D.
		if buffer[0] == 0x1b && buffer[1] == '[' {
			switch buffer[2] {
			case 'A':
				return "UP", nil
			case 'B':
				return "DOWN", nil
			case 'C':
				return "RIGHT", nil
			case 'D':
				return "LEFT", nil

			}
		}
	}

	return "", nil

}

func makeMove(oldRow, oldCol int, dir string) (newRow, newCol int) {
	newRow, newCol = oldRow, oldCol
	switch dir {
	case "UP":
		newRow = newRow - 1
		if newRow < 0 {
			newRow = len(maze) - 1
		}
	case "DOWN":
		newRow = newRow + 1
		if newRow == len(maze) {
			newRow = 0
		}
	case "RIGHT":
		newCol = newCol + 1
		if newCol == len(maze[0]) {
			newCol = 0
		}
	case "LEFT":
		newCol = newCol - 1
		if newCol < 0 {
			newCol = len(maze[0]) - 1
		}

	}
	if maze[newRow][newCol] == '#' {
		newRow = oldRow
		newCol = oldCol
	}
	return newRow, newCol
}

func movePlayer(dir string) {
	player.row, player.col = makeMove(player.row, player.col, dir)
	switch maze[player.row][player.col] {
	case '.':
		numDots--
		score++
		// Remove dot from the maze
		// 不可以使用 maze[player.row][player.col] = " "
		// 因为maze的类型是[]string ， string是不可变类型
		maze[player.row] = maze[player.row][0:player.col] + " " + maze[player.row][player.col+1:]
	}
}

func drawDirection() string {
	dir := rand.Intn(4)
	move := map[int]string{
		0: "UP",
		1: "DOWN",
		2: "RIGHT",
		3: "LEFT",
	}
	return move[dir]

}

func moveGhost() {
	for _, g := range ghosts {
		dir := drawDirection()
		g.row, g.col = makeMove(g.row, g.col, dir)
	}
}

func main() {
	// initialise game
	initialise()
	defer cleanup()

	// load resources
	err := loadMaze("resources/maze01.txt")
	if err != nil {
		log.Println("failed to load maze:", err)
		return
	}

	// game loop
	for {
		// update screen
		printScreen()

		// process input
		// 在每次循环中调用ReadInput
		input, err := readInput()
		if err != nil {
			log.Println("error reading input: ", err)
			break
		}

		// process movement
		movePlayer(input)
		moveGhost()

		// process collisions

		// check game over
		if input == "ESC" || numDots == 0 || lives <= 0 {
			break
		}

		// Temp: break infinite loop

		// repeat
	}
}
