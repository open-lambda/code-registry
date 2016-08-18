package main

import (
	"fmt"

	r "github.com/open-lambda/code-registry/registry"
	"github.com/open-lambda/load-balancer/balancer/inspect/codegen"
)

const (
	SERVER_ADDR = "127.0.0.1:10000"
	SERVER_PORT = 10000
	CHUNK_SIZE  = 1024
	DATABASE    = "registry"

	NAME         = "TEST"
	PROTO_PUSH   = "proto.in"
	PROTO_PULL   = "proto.out"
	HANDLER_PUSH = "handler.in"
	HANDLER_PULL = "handler.out"
)

func generateParser(name string, file []byte) ([]byte, error) {
	return []byte("This is a fake parsing libary"), nil
}

type LBFileProcessor struct{}

func (p LBFileProcessor) Process(name string, files map[string][]byte) ([]r.DBInsert, error) {
	ret := make([]r.DBInsert, 0)
	pb, err := codegen.Generate(files[r.PROTO], name)
	if err != nil {
		return ret, err
	}

	parser, err := generateParser(name, files[r.PROTO])
	if err != nil {
		return ret, err
	}

	sfiles := map[string]interface{}{
		"id":      name,
		"handler": files[r.HANDLER],
		"pb":      pb,
	}
	sinsert := r.DBInsert{
		Table: r.SERVER,
		Data:  &sfiles,
	}
	ret = append(ret, sinsert)

	lbfiles := map[string]interface{}{
		"id":     name,
		"parser": parser,
	}
	lbinsert := r.DBInsert{
		Table: r.BALANCER,
		Data:  &lbfiles,
	}
	ret = append(ret, lbinsert)

	return ret, nil
}

func main() {
	CLUSTER := []string{"127.0.0.1:28015"}
	proc := LBFileProcessor{}
	pushs := r.InitPushServer(CLUSTER, DATABASE, proc, SERVER_PORT, CHUNK_SIZE, SERVER, BALANCER)
	fmt.Println("Running pushserver...")
	pushs.Run()
}
