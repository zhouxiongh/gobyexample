package main

import "fmt"
import "time"

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)
	counter := 0

	go func() {
		for {
			select {
			case <- done:
				return
			case t := <-ticker.C:
				counter++
				fmt.Println("Tick at", t)
				fmt.Println("Counter:", counter)
			}
		}
	}()
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
