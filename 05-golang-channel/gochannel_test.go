package gochannel

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func Test_Gochannel(t *testing.T) {
	gochan()
}

func Test_RandInt(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(10))
}
