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
		log.Println("[MessageExchange] : start the function")
		defer wg.Done()
		for {
			select {
			case message := <-MessageClientIn:
				log.Println("[MessageExchange] : received message from client ")
				ServerDialOut <- message
			case message := <-ServerDialIn:
				log.Println("[MessageExchange] : received message from server ")
				MessageClientOut <- message
			}
		}
	}()

	wg.Wait()
}
