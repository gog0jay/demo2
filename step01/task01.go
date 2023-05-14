// read the maze01.txt

package main

import (
	"bufio"
	"fmt"
	"os"
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
	defer f.Close()

	// read the file line by line and appends it to the maze slice
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		maze = append(maze, line)
	}
	return nil
	// f.Close() is called implicitly

}

func main() {
	loadMaze("../resources/maze01.txt")
	for i, v := range maze {
		fmt.Printf("maze[%d] = %v \n", i, v)
	}

}
