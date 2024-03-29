package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type Calc int

type Args struct {
	A, B int
}

type Reply struct {
	C int
}

func (c *Calc) Sum(args Args, reply *Reply) error {
	reply.C = args.A + args.B
	return nil
}

func main() {
	rpc.Register(new(Calc))
	listen, err := net.Listen("tcp", ":5000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer conn.Close()

		go rpc.ServeConn(conn)
	}
}
