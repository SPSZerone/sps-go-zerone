package coscli

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/tencentyun/cos-go-sdk-v5"

	spsrand "github.com/SPSZerone/sps-go-zerone/math/rand"
)

func Put(opts *Options) (*cos.Response, error) {
	rsp, err := DoPut(opts)
	if err == nil {
		return rsp, nil
	}
	if opts.Args.ErrRetryNum <= 0 {
		return rsp, err
	}

	for i := 0; i < opts.Args.ErrRetryNum; i++ {
		time.Sleep(time.Duration(spsrand.RandomInt(1, 5)) * time.Second)
		rsp, err = DoPut(opts)
		if err == nil {
			return rsp, nil
		}
	}

	return nil, fmt.Errorf("put failed")
}

func DoPut(opts *Options) (*cos.Response, error) {
	if opts.Args.SrcFilePath != "" {
		return DoPutFromFile(opts.Client, opts.Args.FileName, opts.Args.SrcFilePath)
	}

	pipeReader, err := opts.NewPipeReader()
	if err != nil {
		return nil, err
	}
	return DoPutFromReader(opts.Client, opts.Args.FileName, pipeReader)
}

func DoPutFromFile(client *cos.Client, fileName string, srcFilePath string) (*cos.Response, error) {
	opt := &cos.ObjectPutOptions{}
	return client.Object.PutFromFile(context.Background(), fileName, srcFilePath, opt)
}

func DoPutFromReader(client *cos.Client, fileName string, reader io.Reader) (*cos.Response, error) {
	opt := &cos.ObjectPutOptions{}
	return client.Object.Put(context.Background(), fileName, reader, opt)
}
