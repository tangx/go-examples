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
