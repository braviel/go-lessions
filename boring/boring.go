package boring

import (
	"fmt"
	"math/rand"
	"time"
)

//Boring func
func Boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

//FanIn func
func FanIn(in1, in2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-in1
		}
	}()
	go func() {
		for {
			c <- <-in2
		}
	}()
	return c
}
