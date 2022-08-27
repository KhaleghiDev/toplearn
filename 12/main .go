package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	//cmd  telnet chec  ->telnet localhos 8000
	l, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go func(c net.Conn) {
			c.Write([]byte("massage received ."))
			fmt.Println("massage received.")
			c.Close()
		}(conn)
	}
}
