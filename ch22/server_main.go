package main

import (
	"gotour/ch22/server"
	"log"
	"net"
	"net/rpc"
)

func main()  {
	rpc.RegisterName("MathService",new(server.MathService))

	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	rpc.Accept(l)
}


