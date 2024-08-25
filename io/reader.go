package io

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func ReadFileBytes(fileName string, bufSize uint) ([]byte, error) {
	if bufSize == 0 {
		bufSize = 10240
	}

	file, errOpen := os.Open(fileName)
	if errOpen != nil {
		return nil, errOpen
	}
	defer func(file *os.File) {
		errClose := file.Close()
		if errClose != nil {
			fmt.Println("close file fail", fileName, errClose)
		}
	}(file)

	reader := bufio.NewReader(file)
	return ReadBytes(reader, bufSize)
}

func ReadBytes(reader io.Reader, bufSize uint) ([]byte, error) {
	buf := make([]byte, bufSize)
	bytes := make([]byte, 0)
	for {
		size, err := reader.Read(buf)
		if size > 0 {
			bytes = append(bytes, buf[:size]...)
		}

		if err == io.EOF || size == 0 {
			break
		} else if err != nil {
			return nil, err
		}
	}
	return bytes, nil
}
