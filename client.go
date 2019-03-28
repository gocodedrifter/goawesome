package main

import "net"

// Client : save the incoming call
type Client struct {
	socket net.Conn
	data   chan []byte
}
