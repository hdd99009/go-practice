package main

import (
	"fmt"
	"practice/homework02"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	var count int64
	count = 1
	for i := 0; i < 10; i++ {

		go homework02.AddOnethousand2(&wg, &count)
	}
	wg.Wait()
	fmt.Println(count)
}
