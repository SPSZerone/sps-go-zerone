package coscli

import (
	"fmt"
	"io/fs"
)

const (
	ModPut = "put"
)

func NewArgs() *Args {
	return &Args{}
}

type Args struct {
	Mod         string
	FileName    string
	SrcFilePath string
	File        fs.File
	ErrRetryNum int
}

func Args2String(args *Args) string {
	return fmt.Sprintf(
		"Args:{Mod:%s FileName:%s SrcFilePath:%s ErrRetryNum:%d}",
		args.Mod, args.FileName, args.SrcFilePath, args.ErrRetryNum,
	)
}
