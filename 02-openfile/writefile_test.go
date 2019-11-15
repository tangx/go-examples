package openfile

import "testing"

// const (
// 	fileNotExist = `/tmp/fileNotExist.txt`
// )

func Test_Writefile(t *testing.T) {
	osWritefile("/tmp/fileExist.txt", "hello,world")
}

func Test_ioutilWriteFile(t *testing.T) {
	data := "hello, golang \n"
	WriteFileString("/tmp/fileExist.txt", data, 0644)
}

func Test_FileInfo(t *testing.T) {
	FileInfo("/tmp/fileExist.txt")
}
