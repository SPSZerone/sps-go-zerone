package io

import (
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"os"
)

func PipeReader(file fs.File) (io.Reader, error) {
	if file == nil {
		return nil, fmt.Errorf("file is nil")
	}
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}
	if !IsPipe(fileInfo) {
		return nil, fmt.Errorf("file is not a pipe")
	}
	content, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(content), nil
}

func IsPipe(fileInfo os.FileInfo) bool {
	if fileInfo == nil {
		return false
	}
	return IsFileMode(fileInfo.Mode(), os.ModeNamedPipe)
}

func IsFileMode(mode, check fs.FileMode) bool {
	return mode&check == check
}
