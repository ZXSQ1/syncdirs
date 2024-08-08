package files

import (
	"io"
	"os"
)

const BufferSize = 1024 * 10

/*
description: copies the source file to the destination file
arguments:
  - source: the path to the source file
  - destination: the path to the destination file

return: an error if any problem
*/
func Copy(source, destination string) error {
	sourceObj, err := GetFile(source, 0)
	statSource, _ := os.Stat(source)
	sourcePerm := statSource.Mode()

	defer sourceObj.Close()

	if err != nil {
		return err
	}

	destinationObj, err := GetFile(destination, sourcePerm)
	defer destinationObj.Close()

	if err != nil {
		return err
	}

	_, err = io.Copy(destinationObj, sourceObj)

	return err
}
