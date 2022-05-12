package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, `usage:simpleDict WORD example: simpleDict hello`)
		os.Exit(1)
	}
	word := os.Args[1]
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		Query1(word)
	}()
	go func() {
		defer wg.Done()
		Query2(word)
	}()
	wg.Wait()
}
