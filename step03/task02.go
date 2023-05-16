package step03

import (
	"os"
	"pacgo/step01"
)

// Handling arrow key presses

// modify the readInput() in step02--task03

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
			newRow = len(step01.Maze) - 1
		}
	case "DOWN":
		newRow = newRow + 1
		if newRow == len(step01.Maze) {
			newRow = 0
		}
	case "RIGHT":
		newCol = newCol + 1
		if newCol == len(step01.Maze[0]) {
			newCol = 0
		}
	case "LEFT":
		newCol = newCol - 1
		if newCol < 0 {
			newCol = len(step01.Maze[0]) - 1
		}

	}
	if step01.Maze[newRow][newCol] == '#' {
		newRow = oldRow
		newCol = oldCol
	}
	return newRow, newCol
}

func movePlayer(dir string) {
	player.row, player.col = makeMove(player.row, player.col, dir)
}
