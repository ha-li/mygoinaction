package main

import (
	"math/rand"
	"sync"
	"fmt"
	"time"
)

var (
	wg3 sync.WaitGroup
    wgCount = 4
)

func makeTask(c chan string, taskNum int) {
	defer wg3.Done()

	taskInfo, ok := <- c

	if !ok {
		fmt.Printf("worker %d shutting down\n", taskNum)
		return
	}

	fmt.Printf ("worker %d starting task %v\n", taskNum, taskInfo)

    sleep := rand.Int63n(100)
    time.Sleep(time.Duration(sleep)*time.Millisecond)
    fmt.Printf("worker %d completed task %v\n", taskNum, taskInfo)
}

func main() {

	wg3.Add(wgCount)
	c := make (chan string, 10)

	for i := 0; i < wgCount; i++ {
		go makeTask(c, i)
	}

	for i := 0; i < wgCount; i++ {
		c <- fmt.Sprintf("Task %d", i)
	}

	close(c)
	wg3.Wait()
}
