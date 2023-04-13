package main

import (
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/p2p/protocol/circuitv2/relay"
	"log"
)

func main() {
	run()
}

func run() {
	relay1, err := libp2p.New()
	if err != nil {
		log.Printf("Unable to create node,<%v>", err)
		return
	}
	_, err = relay.New(relay1)
	if err != nil {
		log.Printf("Fail to instantiate the realy server <%v>", err)
		return
	}

	log.Printf("relay server info is :  ID: %v , Address : %v", relay1.ID(), relay1.Addrs())

	//select {}
}
