package gochannel

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 创建 wg
var wg sync.WaitGroup

func gochan() {
	ch := make(chan int, 4)

	// 增加 1 个锁
	wg.Add(1)

	go Sender(ch)
	go Reciver(ch)

	// 等待锁释放
	wg.Wait()
}

func Sender(ch chan int) {

	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(1)

		time.Sleep(time.Duration(n) * time.Second)

		ch <- i
		fmt.Printf("-> Send %d, Cooldown %d s\n", i, n)

	}
	// 关闭 ch
	// https://colobu.com/2016/04/14/Golang-Channels/#close
	close(ch)

}

func Reciver(ch chan int) {
	// 释放
	defer wg.Done()

	for i := range ch {

		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(4)
		fmt.Printf("<- Recived %d , CoolDonw %d s\n", i, n)

		time.Sleep(time.Duration(n) * time.Second)
	}
}
