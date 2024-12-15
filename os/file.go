package os

import "os"

func FileExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func MakeDir(dir string) error {
	if FileExist(dir) {
		return nil
	}
	return os.MkdirAll(dir, os.ModePerm)
}
