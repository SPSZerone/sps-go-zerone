package main

import (
	"flag"
	"os"

	"github.com/SPSZerone/sps-go-zerone/cloud/tencent/coscli"
)

func main() {
	args := coscli.NewArgs()
	flag.StringVar(&args.Mod, "mod", "", "Mod e.g. --mod=put")
	flag.StringVar(&args.FileName, "file-name", "", "FileName e.g. --file-name=test-000.txt")
	flag.StringVar(&args.SrcFilePath, "src-file-path", "", "SrcFilePath e.g. --src-file-path=test.txt")
	flag.IntVar(&args.ErrRetryNum, "err-retry-num", 0, "ErrRetryNum e.g. --err-retry-num=3")
	flag.Parse()

	args.File = os.Stdin

	_, _ = coscli.Exec(args)
}
