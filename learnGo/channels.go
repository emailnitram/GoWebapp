package main

import (
	"fmt"
	// "time"
)

type Sushi string

func main() {
	// ch := make(chan string)
	// go func() {
	// 	ch <- "Hello!"
	// 	ch <- "hi!"
	// 	ch <- "haha!"
	// 	ch <- "ok no!"
	// 	close(ch)
	// }()
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// v, ok := <-ch
	// fmt.Println(v)
	// fmt.Println(ok)

	// sushi
	// 	var ch <-chan Sushi
	// 	ch = Producer()
	// 	for s := range ch {
	// 		fmt.Println("Consumed", s)
	// 	}
	// }
	// func Producer() <-chan Sushi {
	// 	ch := make(chan Sushi)
	// 	go func() {
	// 		ch <- Sushi("hello")
	// 		ch <- Sushi("hola")
	// 		close(ch)
	// 	}()
	// 	return ch

	// Synchronization
	// wait := Publish("Channels let goroutines communicate.", 5*time.Second)
	// fmt.Println("Waiting for the news...")
	// <-wait
	// fmt.Println("The news is out, time to leave.")

	// race()
	sharingIsCaring()
}

// Synchronization
// Publish prints text to stdout after the given time has expired.
// It closes the wait channel when the text has been published.
// func Publish(text string, delay time.Duration) (wait <-chan struct{}) {
// 	ch := make(chan struct{})
// 	go func() {
// 		time.Sleep(delay)
// 		fmt.Println("BREAKING NEWS: ", text)
// 		close(ch) // broadcast - a closed channel sends a zero value forever
// 	}()
// 	return ch
// }

func race() {
	wait := make(chan struct{})
	n := 0
	go func() {
		n++
		close(wait)
	}()
	n++
	<-wait
	fmt.Println(n)
}

func sharingIsCaring() {
	ch := make(chan int)
	go func() {
		n := 0
		n++
		ch <- n
	}()
	n := <-ch
	n++
	fmt.Println(n)
}
