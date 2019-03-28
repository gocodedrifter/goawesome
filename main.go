package main

import (
	"log"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	log.Println("starting ...")

	wg.Add(3)
	go StartListenerServer()
	go StartDialManager()

	go func() {
		defer wg.Done()
		for {
			select {
			case message := <-MessageClientIn:
				ServerDialOut <- message
			case message := <-ServerDialIn:
				MessageClientOut <- message
			}
		}
	}()

	wg.Wait()
}
