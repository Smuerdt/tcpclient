package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
	log.Println("Hello Client!")
	log.Println("Host and Port of Server:")

	// var str string
	// fmt.Scanf("%s", &str)

	str := "127.0.0.1:9988"

	tcpAddr, err := net.ResolveTCPAddr("tcp", str)
	if err != nil {
		log.Println("Error: Could not resolve Address")
	} else {
		log.Println("Connecting to: ", tcpAddr.String())
		conn, err := net.Dial("tcp", tcpAddr.String())
		if err != nil {
			log.Println("Error: Could not connect")
		} else {
			log.Println("Successfully connected")
			go ClientSender(conn)
			ClientReader(conn)
		}
	}
}

func ClientReader(conn net.Conn) {

	for {
		buffer := make([]byte, 2048)
		bytesRead, err := conn.Read(buffer)
		if err != nil {
			conn.Close()
			log.Println(err)
			break
		}

		log.Println("Read", bytesRead, "bytes")
		log.Println("ClientReader received >", string(buffer[0:bytesRead]))
	}
}

func ClientSender(conn net.Conn) {
	log.Println("Client sending Name:")

	var name string
	fmt.Scanf("%s", &name)

	conn.Write([]byte(name))

	var send string
	for {
		fmt.Scanf("%s\n", &send)
		count := 0
		for i := 0; i < len(send); i++ {
			if send[i] == 0x00 {
				break
			}
			count++
		}
		log.Println("Send size: ", count)
		conn.Write([]byte(send)[0:count])
	}
}
