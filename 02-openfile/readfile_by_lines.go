package openfile

// 按行读取文件
// 关于使用 bufio 包中 ReadString 和 Scanner 的优劣
// https://stackoverflow.com/questions/47479564/go-bufio-readstring-in-loop-is-infinite

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func bufioReadline(path string) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	br := bufio.NewReader(f)
	for {
		line, isPrefix, err := br.ReadLine()
		if err == io.EOF {
			break
		}

		fmt.Println(isPrefix, string(line))
	}
}

func bufioScanner(path string) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		fmt.Println(string(scanner.Bytes()))
	}

}
