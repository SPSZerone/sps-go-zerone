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

	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = file.Close()
		if err != nil {
			fmt.Printf("Error closing file %s err: %+v\n", fileName, err)
		}
	}()

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
