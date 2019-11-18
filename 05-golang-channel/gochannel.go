package gochannel

import (
	"fmt"
	"math/rand"
	"time"
)

func gochan() {
	ch := make(chan int, 4)
	go Sender(ch)
	Reciver(ch)

}

func Sender(ch chan int) {
	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(3)

		time.Sleep(time.Duration(n) * time.Second)

		ch <- i
		fmt.Printf("-> Send %d, Cooldown %d s\n", i, n)

	}
	// 关闭 ch
	// https://colobu.com/2016/04/14/Golang-Channels/#close
	close(ch)

}

func Reciver(ch chan int) {
	for i := range ch {

		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(3)
		fmt.Printf("<- Recived %d , CoolDonw %d s\n", i, n)

		// time.Sleep(time.Duration(n) * time.Second)

	}
}
