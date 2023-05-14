package main

import (
	"bufio"
	"fmt"
	"github.com/danicat/simpleansi"
	"log"
	"os"
	"os/exec"
)

var maze []string

func loadMaze(file string) error {

	// os.Open读取文件，有两个返回值
	f, err := os.Open(file)

	// error handling
	if err != nil {
		return err
	}

	// put f.Close() in the Call stack
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	// read the file line by line and appends it to the maze slice
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		maze = append(maze, line)
	}
	return nil
	// f.Close() is called implicitly

}

func printScreen() {
	simpleansi.ClearScreen()
	for _, line := range maze {
		fmt.Println(line)
	}

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
	}

	return "", nil

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
		if input == "Esc" {
			break
		}

		// process movement

		// process collisions

		// check game over

		// Temp: break infinite loop
		break

		// repeat
	}
}
