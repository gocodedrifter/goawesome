package main

var (
	// MessageClientIn : message from client listen by the server
	MessageClientIn = make(chan []byte)
	// MessageClientOut : response message to client after receipt from ServerDialIn after dial to the iso server
	MessageClientOut = make(chan []byte)
	// ServerDialOut : receipt message from MessageClientIn to prosess for dial
	ServerDialOut = make(chan []byte)
	// ServerDialIn : response message after dial to the iso server
	ServerDialIn = make(chan []byte)
)
