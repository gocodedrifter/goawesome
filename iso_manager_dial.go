package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

// IsoManagerDial : iso manager dial
type IsoManagerDial struct {
	socket net.Conn
}

// Receive : receive the incoming message from iso server (gsp/pln/etc ...)
func (manager *IsoManagerDial) Receive() {
	defer wg.Done()
	for {
		message := make([]byte, 4096)
		length, err := manager.socket.Read(message)
		if err != nil {
			log.Println("[Receive()] : the connection to dial is closed : ",
				GetConfig().Iso.Server.Dial.IP, GetConfig().Iso.Server.Dial.Port)

			log.Println("[Receive()] : try to redial ")
			manager.socket = handleDialConnection()
		}

		if length > 0 {
			log.Println("[Receive()] : Received message from iso server (gsp/pln/etc ...) : ",
				string(message))
			ServerDialIn <- message
		}
	}
}

// GetListenerMessage : get listener message
func (manager *IsoManagerDial) GetListenerMessage() {
	defer wg.Done()
	for {
		select {
		case message := <-ServerDialOut:
			manager.socket.Write(message)
		}
	}
}

func handleDialConnection() net.Conn {
	connection, err := net.Dial("tcp", fmt.Sprintf("%s:%s",
		GetConfig().Iso.Server.Dial.IP, GetConfig().Iso.Server.Dial.Port))

	if err != nil {
		log.Println("[handleDialConnection()] : unable to dial to the server : ",
			GetConfig().Iso.Server.Dial.IP, GetConfig().Iso.Server.Dial.Port)
		panic(err.Error())
	}

	err = connection.(*net.TCPConn).SetKeepAlive(true)
	if err != nil {
		log.Println("[handleDialConnection()] : unable to keep the server dial always live : ",
			GetConfig().Iso.Server.Dial.IP, GetConfig().Iso.Server.Dial.Port)
		panic(err.Error())
	}

	connection.(*net.TCPConn).SetDeadline(time.Time{})

	return connection
}

// StartDialManager : start dial manager
func StartDialManager() {
	log.Println("[StartDialManager()] : start to dial manager")

	defer wg.Done()
	manager := &IsoManagerDial{
		socket: handleDialConnection(),
	}

	wg.Add(2)
	go manager.Receive()
	go manager.GetListenerMessage()
	wg.Wait()
}
