package main

import (
	"github.com/apache/thrift/lib/go/thrift"

	"service.demo/gen-go/filesafe"
)

type FileServer struct{}

func main() {
	transport, err := thrift.NewTServerSocket("localhost:12345")
	if err != nil {
		panic(err)
	}

	handler := &FileServer{}
	processor := filesafe.NewInternalServiceProcessor(handler)
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactory(false, false)
	server := thrift.NewTSimpleServer4(
		processor,
		transport,
		transportFactory,
		protocolFactory,
	)

	if err := server.Serve(); err != nil {
		panic(err)
	}
}
