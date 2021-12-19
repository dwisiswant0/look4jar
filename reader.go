package main

import (
	"archive/zip"
	"io/ioutil"
)

func readClass(file *zip.File) []byte {
	read, err := file.Open()
	if err != nil {
		panic(err)
	}

	content, err := ioutil.ReadAll(read)
	if err != nil {
		panic(err)
	}
	defer read.Close()

	return content
}

func readJar(file string) (*zip.ReadCloser, error) {
	jar, err := zip.OpenReader(file)
	if err != nil {
		return nil, err
	}

	return jar, nil
}
