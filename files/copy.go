package files

import "io"

const BufferSize = 1024 * 10

/*
description: copies the source file to the destination file
arguments:
	- source: the path to the source file
	- destination: the path to the destination file
return: an error if any problem
*/
func Copy(source, destination string) error {
	sourceObj, err := GetFile(source)
	defer sourceObj.Close()

	if err != nil {
		return err
	}

	destinationObj, err := GetFile(destination)
	defer destinationObj.Close()

	if err != nil {
		return err
	}

	buffer := make([]byte, BufferSize)
	
	for {
		nRead, err := sourceObj.Read(buffer)

		if err != nil {
			if err == io.EOF {
				break
			}

			return err
		}

		_, err = destinationObj.Write(buffer[:nRead])

		if err != nil {
			return err
		}
	}

	return nil
}
