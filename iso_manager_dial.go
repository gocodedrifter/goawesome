package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/config"
)

// IsoManagerDial : iso manager dial
type IsoManagerDial struct {
	socket net.Conn
}

// Receive : receive the incoming message from iso server (gsp/pln/etc ...)
func (manager *IsoManagerDial) Receive(result chan []byte) {
	log.Println("IsoManagerDial[Receive(result chan []byte)] : start to receive message from socket")
	for {
		message := make([]byte, 4096)
		length, err := manager.socket.Read(message)
		if err != nil {
			log.Panic("error reading : ", err.Error())
		}

		if length > 0 {
			log.Println("[Receive(result chan []byte)] : Received message from iso server (gsp/pln/etc ...) : ",
				string(message))
			result <- message
			break
		}
	}

	log.Println("IsoManagerDial[Receive(result chan []byte)] : end receive message from socket")
}

func handleDialConnection() net.Conn {
	log.Println("IsoManagerDial[handleDialConnection()] : start connection")
	connection, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%s",
		config.Get().Iso.Server.Dial.IP, config.Get().Iso.Server.Dial.Port), 5*time.Second)

	if err != nil {
		log.Println("IsoManagerDial[handleDialConnection()] : unable to dial to the server : ",
			config.Get().Iso.Server.Dial.IP, config.Get().Iso.Server.Dial.Port)
		panic(err.Error())
	}

	log.Println("IsoManagerDial[handleDialConnection()] : end connection")

	return connection
}

// StartDialManager : start dial manager
func StartDialManager(message []byte, result chan []byte) {
	log.Println("IsoManagerDial.[StartDialManager()] : start to dial manager")

	log.Println("IsoManagerDial.[StartDialManager()] : starting dial connection")
	connection := handleDialConnection()

	manager := &IsoManagerDial{
		socket: connection,
	}

	log.Println("IsoManagerDial.[StartDialManager()] : start thread to recieve message")
	msg := make(chan []byte)
	go manager.Receive(msg)
	manager.socket.Write(message)

	// message exchange
	msgIn := <-msg
	result <- msgIn

	defer connection.Close()
	defer manager.socket.Close()
	log.Println("IsoManagerDial.[StartDialManager()] : dial manager closed")
}
