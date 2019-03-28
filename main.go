package main

import (
	"fmt"
)

func main() {
	fmt.Printf("config: %s\n", GetConfig().Iso.Server.Listener.IP)
}
