package main

import (
	"fmt"
)

func main() {
	pool := make(chan chan int, 100)
	chi := make(chan int)
	chi2 := make(chan int)
	chi3 := make(chan int)
	go func() {
		for {
			pool <- chi
			pool <- chi2
			pool <- chi3
			select {
			case aa := <-chi:
				fmt.Println("aa:", aa)
			case bb := <-chi2:
				fmt.Println("bb:", bb)
			case cc := <-chi3:
				fmt.Println("cc:", cc)
			}
		}
	}()

	pp := <-pool
	chi <- 20
	pp <- 21
	chi2 <- 29
	cc := <-pool
	cc <- 30
	cc <- 33
	dd := <-pool
	dd <- 1
	kk := <-pool
	kk <- 99
	//time.Sleep(time.Second * 3)
}
