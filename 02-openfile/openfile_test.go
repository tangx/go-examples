package openfile

import (
	"fmt"
	"os"
	"testing"
)

const (
	fileExist    = `/tmp/fileExist.txt`
	fileNotExist = `/tmp/fileNotExist.txt`
)

func Test_MustIsExist(t *testing.T) {

	if !IsFileExist(fileExist) {
		os.Create(fileExist)
	}

	fmt.Println(IsFileExist(fileExist))
	fmt.Println(IsFileExist(fileNotExist))
}

func Test_MustOpenByOsOpen(t *testing.T) {
	// fmt.Println(MustOpenByOsOpen(fileExist))
	fmt.Printf("%s", MustOpenByOsOpen(fileExist))
}

func Test_MustOpenByOsOpenfile(t *testing.T) {
	fmt.Printf("%s", MustOpenByOsOpenfile(fileExist))
}

func Test_MustOpenByIoutils(t *testing.T) {
	fmt.Printf("%s", MustOpenByIoutils(fileExist))
}
