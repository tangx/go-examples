package forloop

import (
	"fmt"
	"testing"
)

type Frame struct {
	frame int
}

func NewFrame() *Frame {
	return &Frame{
		frame: 0,
	}
}

func (f *Frame) Next() (int, bool) {
	f.frame++

	if f.frame < 10 {
		return f.frame, true
	}

	return f.frame, false
}

func Test_ForLoop(t *testing.T) {
	// for frame,more=frames.Next(); more;  frame,more=frames.Next() {
	// 	// statement
	// 	if !more{
	// 		break
	// 	}
	// }

	frames := NewFrame()

	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		fmt.Println(frame)

		// if !more {
		// 	break
		// }
	}

}
