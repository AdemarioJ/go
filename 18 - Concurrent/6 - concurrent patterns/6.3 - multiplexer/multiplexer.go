package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := multiplexer(write("Hello World!"), write("Programming in Go!"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}
}

func multiplexer(channelEntries1, channelEntries2 <-chan string) <-chan string {
	channelExit := make(chan string)

	go func() {
		for {
			select {
			case message := <-channelEntries1:
				channelExit <- message
			case message := <-channelEntries2:
				channelExit <- message
			}
		}
	}()

	return channelExit
}

func write(text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Value received: %s", text)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return channel
}
