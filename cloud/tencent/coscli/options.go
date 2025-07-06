package coscli

import (
	"bytes"
	"fmt"
	"io"

	"github.com/tencentyun/cos-go-sdk-v5"

	spsos "github.com/SPSZerone/sps-go-zerone/os"
)

func NewOptions(args *Args) (*Options, error) {
	client, err := NewClient()
	if err != nil {
		return nil, err
	}
	o := &Options{
		Args:   args,
		Client: client,
	}
	o.Init()
	return o, nil
}

type Options struct {
	Args   *Args
	Client *cos.Client

	isPipe      bool
	pipeContent []byte
}

func (o *Options) Init() {
	o.isPipe = spsos.IsPipeFile(o.Args.File)
	if o.isPipe {
		content, err := io.ReadAll(o.Args.File)
		if err == nil {
			o.pipeContent = content
			fmt.Println(fmt.Sprintf("PipeReader bytes: %d", len(content)))
		} else {
			fmt.Println(fmt.Sprintf("PipeReader err: %+v", err))
		}
	}
}

func (o *Options) NewPipeReader() (io.Reader, error) {
	if !o.isPipe {
		return nil, fmt.Errorf("not pipe")
	}
	if o.pipeContent == nil {
		return nil, fmt.Errorf("pipe content is nil")
	}
	return bytes.NewReader(o.pipeContent), nil
}
