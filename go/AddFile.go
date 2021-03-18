package main

import (
	"context"
	"io/ioutil"
	"os"

	"service.demo/gen-go/filesafe"
)

func (e *FileServer) AddFile(ctx context.Context, req *filesafe.AddFileReq) (*filesafe.AddFileRes, error) {
	filedir := req.FilePath

	newFile, err := os.Create(filedir)
	defer newFile.Close()
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filedir, req.Buff, 0666)
	if err != nil {
		panic(err)
	}

	res := &filesafe.AddFileRes{
		StatusCode: 0,
		Message:    "success",
	}

	return res, nil
}
