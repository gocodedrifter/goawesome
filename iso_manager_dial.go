package main

import (
	"fmt"
	"log"
	"net"

	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/config"
)

// IsoManagerDial : iso manager dial
type IsoManagerDial struct {
	socket net.Conn
}

// Receive : receive the incoming message from iso server (gsp/pln/etc ...)
func (manager *IsoManagerDial) Receive() {
	log.Println("IsoManagerDial[Receive()] : starting to receive ")
	defer wg.Done()
	for {
		message := make([]byte, 4096)
		length, err := manager.socket.Read(message)
		if err != nil {
			log.Println("[Receive()] : the connection to dial is closed : ",
				config.Get().Iso.Server.Dial.IP, config.Get().Iso.Server.Dial.Port)

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
	log.Println("IsoManagerDial[GetListenerMessage()] : get listener message  ")
	defer wg.Done()
	for {
		select {
		case message := <-ServerDialOut:
			if len(message) > 0 {
				manager.socket.Write(message)
			}
		}
	}
}

func handleDialConnection() net.Conn {
	connection, err := net.Dial("tcp", fmt.Sprintf("%s:%s",
		config.Get().Iso.Server.Dial.IP, config.Get().Iso.Server.Dial.Port))

	if err != nil {
		log.Println("IsoManagerDial[handleDialConnection()] : unable to dial to the server : ",
			config.Get().Iso.Server.Dial.IP, config.Get().Iso.Server.Dial.Port)
		panic(err.Error())
	}

	err = connection.(*net.TCPConn).SetKeepAlive(true)
	if err != nil {
		log.Println("IsoManagerDial[handleDialConnection()] : unable to keep the server dial always live : ",
			config.Get().Iso.Server.Dial.IP, config.Get().Iso.Server.Dial.Port)
		panic(err.Error())
	}

	// connection.(*net.TCPConn).SetDeadline(time.Time{})

	return connection
}

// StartDialManager : start dial manager
func StartDialManager() {
	defer wg.Done()
	log.Println("[StartDialManager()] : start to dial manager")

	connection := handleDialConnection()

	manager := &IsoManagerDial{
		socket: connection,
	}

	wg.Add(2)
	go manager.Receive()
	go manager.GetListenerMessage()
	wg.Wait()
}
