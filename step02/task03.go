package step02

import "os"

// reading from Stdin

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
