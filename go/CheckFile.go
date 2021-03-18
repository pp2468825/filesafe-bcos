package main

import (
	"context"

	"service.demo/gen-go/filesafe"
)

func (e *FileServer) CheckFile(ctx context.Context, req *filesafe.CheckFileReq) (*filesafe.CheckFileRes, error) {
	return nil, nil
}
