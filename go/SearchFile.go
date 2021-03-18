package main

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"service.demo/gen-go/filesafe"
)

type filemeta struct {
	Type string
	Size int64
	Name string
}

type dirmeta struct {
	Type    string
	Size    int64
	Name    string
	Content []filemeta
}

func (e *FileServer) SearchFile(ctx context.Context, req *filesafe.SearchFileReq) (*filesafe.SearchFileRes, error) {

	files, err := ioutil.ReadDir("./filesafe")
	if err != nil {
		panic(err)
	}

	filelist := dirmeta{
		Type:    "dir",
		Size:    0,
		Name:    "filesafe",
		Content: make([]filemeta, 0),
	}

	for _, file := range files {
		filelist.Content = append(filelist.Content, filemeta{
			Type: "file",
			Name: file.Name(),
			Size: file.Size(),
		})
		filelist.Size += file.Size()
	}
	str, err := json.Marshal(filelist)
	rsp := &filesafe.SearchFileRes{
		StatusCode: 0,
		Message:    "success",
		FileList:   filesafe.JsonString(str),
	}
	return rsp, err
}
