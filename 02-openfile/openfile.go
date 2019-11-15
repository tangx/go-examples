package openfile

import "io/ioutil"

import "fmt"

import "os"

func MustOpenByIoutils(filename string) []byte {

	body, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
	return body
}

// IsFileExist 判断文件是否存在
func IsFileExist(filename string) bool {

	_, err := os.Stat(filename)
	if err != nil {

		// fmt.Println(err)
		if !os.IsExist(err) {
			return false
		}

	}
	return true
}

func MustOpenByOsOpen(filename string) []byte {
	body, err := OpenByOsOpen(filename)
	if err != nil {
		panic(err)
	}

	PrintFile(body)
	return body
}

func OpenByOsOpen(filename string) ([]byte, error) {
	fobj, err := os.Open(filename)
	if err != nil {
		panic(fobj)
	}
	return ioutil.ReadAll(fobj)
}

func PrintFile(body []byte) {
	fmt.Println(string(body))
}

func OpenByOsOpenfile(filename string) ([]byte, error) {
	fobj, err := os.OpenFile(filename, os.O_RDONLY, 0644)

	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(fobj)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func MustOpenByOsOpenfile(filename string) []byte {
	body, err := OpenByOsOpenfile(filename)
	if err != nil {
		panic(err)
	}
	return body
}
