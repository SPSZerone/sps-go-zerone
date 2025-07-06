package coscli

import (
	"fmt"
	"os"

	"github.com/tencentyun/cos-go-sdk-v5"
)

func Exec(args *Args) (*cos.Response, error) {
	rsp, err := doExec(args)
	PrintResponse(rsp, err)
	if err != nil {
		os.Exit(1)
	}
	return rsp, err
}

func doExec(args *Args) (*cos.Response, error) {
	fmt.Println(Args2String(args))

	opts, err := NewOptions(args)
	if err != nil {
		return nil, err
	}

	if args.Mod == ModPut {
		return Put(opts)
	}

	return nil, fmt.Errorf("unknown mod %s", args.Mod)
}

func PrintResponse(rsp *cos.Response, err error) {
	if rsp != nil {
		fmt.Println(fmt.Sprintf("Status:%s", rsp.Status))
	}
	if err != nil {
		fmt.Println(fmt.Sprintf("%+v", err))
	}
}
