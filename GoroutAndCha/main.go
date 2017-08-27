package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup
var data chan string

func main() { }

func worker() {
	fmt.Println("Goroutine worker is now starting...")
	defer func() {
		fmt.Println("Destroying the worker...")
		waitGroup.Done()
	}()
	for {
		value, ok := <-data
		if !ok {
			fmt.Println("The channel is closed!")
			break
		}
		fmt.Println(value)
		time.Sleep(time.Second * 1)
	}
}