package main

import (
	"fmt"
	"log"
	"net"

	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/config"
)

// IsoManagerListener : iso manager listener
type IsoManagerListener struct {
	clients    map[*Client]string
	register   chan *Client
	unregister chan *Client
}

// Start : listen to each connection from the client
func (manager *IsoManagerListener) Start() {
	for {
		select {
		case connection := <-manager.register:
			manager.clients[connection] = ""
			log.Println("Added new connection!")
		case connection := <-manager.unregister:
			if _, ok := manager.clients[connection]; ok {
				close(connection.data)
				delete(manager.clients, connection)
				log.Println("A Connection has terminated!")
			}
		case message := <-MessageClientOut:
			log.Println("send back to client : ", string(message))
			for connection := range manager.clients {
				select {
				case connection.data <- message:
				default:
					close(connection.data)
					delete(manager.clients, connection)
				}
			}
		}
	}
}

// Receive : receive incoming call from client
func (manager *IsoManagerListener) Receive(client *Client) {
	for {
		message := make([]byte, 4096)
		length, err := client.socket.Read(message)
		if err != nil {
			manager.unregister <- client
			client.socket.Close()
			break
		}
		if length > 0 {
			log.Println("Received call from client : " + string(message))
			MessageClientIn <- message
		}
	}
}

// Send : send the response from dial, back to the client
func (manager *IsoManagerListener) Send(client *Client) {
	defer client.socket.Close()
	for {
		select {
		case message, ok := <-client.data:
			if !ok {
				return
			}
			client.socket.Write(message)
		}
	}
}

// StartListenerServer : start listener server
func StartListenerServer() {
	log.Println("[startListenerServer()] : starting server ...")
	listener, error := net.Listen("tcp", fmt.Sprintf("%s:%s",
		config.Get().Iso.Server.Listener.IP, config.Get().Iso.Server.Listener.Port))

	if error != nil {
		log.Printf("[startListenerServer()]: unable to listen for the ip and port : %s:%s, error : %s",
			config.Get().Iso.Server.Listener.IP, config.Get().Iso.Server.Listener.Port, error.Error())
	}
	manager := IsoManagerListener{
		clients:    make(map[*Client]string),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}

	go manager.Start()

	for {
		connection, _ := listener.Accept()
		if error != nil {
			log.Println("[startListenerServer()]: unable to accept the connection for the client, error : ", error.Error())
		}
		client := &Client{socket: connection, data: make(chan []byte)}
		manager.register <- client
		go manager.Receive(client)
		go manager.Send(client)
	}
}
