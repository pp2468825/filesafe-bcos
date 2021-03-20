package main

import (
	"context"
	"crypto/sha1"
	"io"
	"os"

	"service.demo/gen-go/filesafe"
)

func (e *FileServer) CheckFile(ctx context.Context, req *filesafe.CheckFileReq) (*filesafe.CheckFileRes, error) {
	digest := req.Digest

	file, err := os.Open(path)
	if err != nil {
		return nil, nil
	}
	defer file.Close()
	h := sha1.New()
	_, erro := io.Copy(h, file)
	if erro != nil {
		return nil, nil
	}
	result := h.Sum(nil)
	return nil, nil
}
