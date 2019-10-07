package main

import (
	"encoding/hex"
	"fmt"
	"net"
	"strings"

	log "gitlab.com/kasku/kasku-2pay/2pay-billerpayment/log"

	"github.com/tidwall/sjson"
	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/config"
	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/gsp-pln/messaging/util"
	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/processor"
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
			log.Get().Println("Added new connection!")
		case connection := <-manager.unregister:
			if _, ok := manager.clients[connection]; ok {
				close(connection.data)
				delete(manager.clients, connection)
				log.Get().Println("A Connection has terminated!")
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
			log.Get().Println("Received call from client : " + string(message))
			// convert iso byte to json and process the data
			producer := &processor.Message{}
			producer.SetBuilder(&processor.IsoProcessor{})
			packetIso := string(message[strings.Index(string(message), "2"):])
			packetIsoLength := util.Bcd([]byte(fmt.Sprintf("%04x", len(packetIso)+2)))
			encapsulatedWithLength := append(packetIsoLength, []byte(packetIso)...)
			jsonRequest := producer.DecodeMessage(encapsulatedWithLength)
			jsonRequest, _ = sjson.Set(string(jsonRequest), "payload", string(message))
			jsonResult := Process([]byte(jsonRequest))

			// convert the result from json to byte
			log.Get().Println("response after processed : ", jsonResult)
			producer.SetBuilder(&processor.JSONProcessor{})
			isoResult := producer.Process([]byte(jsonResult))

			// client to send again the data result
			isoMessageNoHeader := string(isoResult[2:])
			idx0, _ := hex.DecodeString(fmt.Sprintf("%02x", len(isoMessageNoHeader)/256))
			idx1, _ := hex.DecodeString(fmt.Sprintf("%02x", len(isoMessageNoHeader)%256))
			n, err := client.socket.Write([]byte(strings.Join([]string{string(idx0), string(idx1), isoMessageNoHeader}, "")))
			log.Get().Println("success : ", n)
			if err != nil {
				log.Get().Println("error : ", err)
			}
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
	log.Get().Println("[startListenerServer()] : starting server ...")
	listener, error := net.Listen("tcp4", fmt.Sprintf("%s:%s",
		config.Get().Iso.Server.Listener.IP, *port))

	if error != nil {
		log.Get().Printf("[startListenerServer()]: unable to listen for the ip and port : %s:%s, error : %s",
			config.Get().Iso.Server.Listener.IP, *port, error.Error())
	}
	manager := IsoManagerListener{
		clients:    make(map[*Client]string),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}

	go manager.Start()

	for {
		connection, _ := listener.Accept()
		log.Get().Println("remote : ", connection.RemoteAddr())
		log.Get().Println("local : ", connection.LocalAddr())
		if error != nil {
			log.Get().Println("[startListenerServer()]: unable to accept the connection for the client, error : ", error.Error())
		}
		client := &Client{socket: connection, data: make(chan []byte)}

		manager.register <- client
		go manager.Receive(client)
		go manager.Send(client)

	}
}
