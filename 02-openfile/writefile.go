package openfile

import (
	"fmt"
	"io/ioutil"
	"os"
)

func osWritefile(filename string, content string) {
	fobj, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer fobj.Close()

	n, err := fobj.WriteString(content)

	if err != nil {
		panic(err)
	}
	fmt.Println(n)
}

// WriteFileString 使用 ioutil.WriteFile 编辑文件
func WriteFileString(filename string, data string, perm os.FileMode) error {
	return ioutil.WriteFile(filename, []byte(data), perm)
}

// FileInfo 显示文件属性
func FileInfo(filename string) {
	f, err := os.Stat(filename)
	if err != nil {
		panic(err)
	}

	fmt.Println("f.IsDir()=", f.IsDir())
	fmt.Println(f.ModTime())
	fmt.Println(f.Mode())
	fmt.Println(f.Size())

	// 可以使用反射获取 f.Sys() 的信息
	// https://studygolang.com/topics/287
	fmt.Println(f.Sys())
}
