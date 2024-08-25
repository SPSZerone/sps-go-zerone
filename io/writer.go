package io

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	spsos "github.com/SPSZerone/sps-go-zerone/os"
)

func WriteBytes(dataBytes []byte, fileName string, overwriteFile bool) error {
	if !overwriteFile {
		if spsos.FileExist(fileName) {
			return errors.New(fmt.Sprintf("FileExist %s", fileName))
		}
	}

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {

		}
	}(file)

	writer := bufio.NewWriter(file)

	_, err = writer.Write(dataBytes)
	if err != nil {
		return err
	}

	err = writer.Flush()
	if err != nil {
		return err
	}

	return nil
}
