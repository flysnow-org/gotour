package main

import (
	"gotour/ch22/server"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	rpc.RegisterName("MathService", new(server.MathService))
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)
}
