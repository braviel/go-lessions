package main

import (
	"fmt"
	"runtime"
	"sync"
)

func g1(wg *sync.WaitGroup) { defer wg.Done(); fmt.Println("g1()") }
func g2(wg *sync.WaitGroup) { defer wg.Done(); fmt.Println("g2()") }
func g3(wg *sync.WaitGroup) { defer wg.Done(); fmt.Println("g3()") }
func g4(wg *sync.WaitGroup) { defer wg.Done(); fmt.Println("g4()") }

// Coroutine function to try out go routine
func Coroutine() {
	numberP := runtime.NumCPU()
	fmt.Println(numberP)
	maxProc := runtime.GOMAXPROCS(0)
	fmt.Println(maxProc)
	fmt.Println("start")
	var wg sync.WaitGroup
	wg.Add(4)
	go g1(&wg)
	go g2(&wg)
	go g3(&wg)
	go g4(&wg)
	wg.Wait()
	fmt.Println("end")
	// unbuffered channel
	fmt.Println("Unbuffered channel:")
	ch := make(chan int)
	go func() {
		fmt.Println("before store")
		ch <- 109
		fmt.Println("after store")
	}()
	fmt.Println("Before get")
	fmt.Println(<-ch)
	fmt.Println("After get")
}
