package gohttp

import "testing"
import "fmt"

func Test_httpGet(t *testing.T) {
	httpGet("ip.cip.cc")
}

func Test_print(t *testing.T){
	fmt.Println("http")
}