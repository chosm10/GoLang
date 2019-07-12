package main

import (
	"fmt"
	"net"
)

func requestHandler(conn net.Conn) {
	data := make([]byte, 4096)

	for {
		n, err := conn.Read(data)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(string(data[:n]))

		_, err = conn.Write(data[:n])
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func main() {
	listen, err := net.Listen("tcp", ":8100")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("err")
			continue
		}
		defer conn.Close()

		go requestHandler(conn)
	}
}
