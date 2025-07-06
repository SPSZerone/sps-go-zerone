package os

import (
	"io/fs"
	"os"
)

func IsPipeFile(file fs.File) bool {
	if file == nil {
		return false
	}
	fileInfo, err := file.Stat()
	if err != nil {
		return false
	}
	return IsPipe(fileInfo)
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
