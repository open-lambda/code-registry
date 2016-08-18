package main

import (
	"fmt"
	r "github.com/open-lambda/code-registry/registry"
)

const (
	SERVER_ADDR = "127.0.0.1:10000"
	SERVER_PORT = 10000
	CHUNK_SIZE  = 1024

	NAME         = "TEST"
	PROTO_PUSH   = "proto.in"
	PROTO_PULL   = "proto.out"
	HANDLER_PUSH = "handler.in"
	HANDLER_PULL = "handler.out"
)

func main() {
	pushc := r.InitPushClient(SERVER_ADDR, CHUNK_SIZE)
	fmt.Println("Pushing from client...")
	proto := r.File{Name: PROTO_PUSH, Type: r.PROTO}
	handler := r.File{Name: HANDLER_PUSH, Type: r.HANDLER}
	pushc.Push(NAME, proto, handler)
}
