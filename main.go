package main

import (
	"fmt"
	"log"
	"net"
)

const (
	fromPort = 5200
	toPort   = 5500
)

func main() {

	scan()

}

func scan() {

	for i := fromPort; i <= toPort; i++ {
		conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", i))
		if err != nil {
			log.Printf("%d CLOSED (%s)\n", i, err)
			continue
		}
		conn.Close()
		log.Printf("%d OPEN\n", i)
	}

}