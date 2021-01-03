package main

import (
	"gotour/ch22/server"
	"io"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	rpc.RegisterName("MathService", new(server.MathService))

	//注册一个path，用于提供基于http的json rpc服务
	http.HandleFunc(rpc.DefaultRPCPath, func(rw http.ResponseWriter, r *http.Request) {
		conn, _, err := rw.(http.Hijacker).Hijack()
		if err != nil {
			log.Print("rpc hijacking ", r.RemoteAddr, ": ", err.Error())
			return
		}
		var connected = "200 Connected to JSON RPC"
		io.WriteString(conn, "HTTP/1.0 "+connected+"\n\n")
		jsonrpc.ServeConn(conn)
	})

	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}

	http.Serve(l, nil)//换成http的服务
}
