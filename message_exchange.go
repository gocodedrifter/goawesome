package main

var (
	// MessageClientIn : message from client listen by the server
	MessageClientIn chan []byte
	// MessageClientOut : response message to client after receipt from ServerDialIn after dial to the iso server
	MessageClientOut chan []byte
	// ServerDialOut : receipt message from MessageClientIn to prosess for dial
	ServerDialOut chan []byte
	// ServerDialIn : response message after dial to the iso server
	ServerDialIn chan []byte
)
