package main

import (
	"fmt"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Reply struct {
	C int
}

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:5000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close()

	args := &Args{2, 3}
	reply := new(Reply)

	err = client.Call("Calc.Sum", args, reply)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(reply.C)

	args.A = 5
	args.B = 10
	sumCall := client.Go("Calc.Sum", args, reply, nil)
	<-sumCall.Done
	fmt.Println(reply.C)
}
