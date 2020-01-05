package openfile

import (
	"fmt"
	"testing"
)

func Test_readfile_by_lines(t *testing.T) {
	bufioReadline("/tmp/1.txt")
	fmt.Println("+++++++")
	bufioScanner("/tmp/1.txt")
}
