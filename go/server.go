package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/apache/thrift/lib/go/thrift"

	"service.demo/gen-go/demo"
)

type FileServer struct {
}

func (e *FileServer) UploadFile(ctx context.Context, req *demo.FileDataReq) (*demo.FileDataRsp, error) {
	fmt.Printf("message from client: %+v\n", req)

	newFile, err := os.Create("2.txt")
	defer newFile.Close()
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("2.txt", req.Buff, 0666)
	if err != nil {
		panic(err)
	}

	res := &demo.FileDataRsp{
		Code: 0,
		Msg:  "success",
	}

	return res, nil
}

func main() {
	transport, err := thrift.NewTServerSocket("localhost:12345")
	if err != nil {
		panic(err)
	}

	handler := &FileServer{}
	processor := demo.NewFileInfoExtractServiceProcessor(handler)
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
