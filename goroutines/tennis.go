package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)


// a global, accessible in all methods and user context
var wg sync.WaitGroup

func init () {
	rand.Seed(time.Now().UnixNano())
}

func play(player string, court chan int, hit int){
	defer wg.Done()
    for {

    	// ok is a boolean that is false when the channel
    	// is closed, otherwise it is true
    	hit, ok := <- court

    	// use ok to test
		if !ok {
			fmt.Printf("%s won!\n", player)
			return
		}

    	if rand.Intn(100) == 0 {
			fmt.Printf ("%s missed %d\n", player, hit)
			close(court)
			return
		}
		hit++
		fmt.Printf("%s play %d\n", player, hit)
		court <- hit
	}
}


func main () {
	//var wg sync.WaitGroup
	wg.Add(2)

	court := make (chan int)
	hit := 0

	go play("Nordic", court, hit)
	go play("Teflon", court, hit)

	// send the first serve, who does it go to?
	// not deterministic
	court <- hit
	wg.Wait()
}
