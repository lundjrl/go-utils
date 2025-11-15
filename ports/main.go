package main

import (
	"io"
	"net"
	"time"

	"github.com/charmbracelet/log"
)

func main() {
	network := "tcp"
	port := ":4000"
	timeInMilli := 100*time.Millisecond
	ln, err := net.DialTimeout(network, port, timeInMilli)
/// fix under
	
	if err != nil {
		panic(err)
	}

	defer ln.Close()

	for {
	conn, err := ln.Accept()

	if err != nil {
		log.Fatal(err)
	}

	go func(c net.Conn) {
		io.Copy(c, c)

		c.Close()
	}(conn)
	}

}
