package uploader

import (
	"io/ioutil"
	"mime/multipart"
	"os"
)

// Uploader Upload Image in server
func Uploader(file multipart.File, header *multipart.FileHeader, path string) (bool, error) {
	newFile, err := os.Create(path)
	if err != nil {
		return false, err
	}

	defer newFile.Close()

	fileByte, err := ioutil.ReadAll(file)
	if err != nil {
		return false, err
	}
	_, err = newFile.Write(fileByte)
	if err != nil {
		return false, err
	}
	return true, nil
}
