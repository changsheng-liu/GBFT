package main

import (
	"os"

	"./network"
)

func main() {
	nodeID := os.Args[1]
	server := network.NewServer(nodeID, 0)

	server.Start()
}
