package main

import (
	"archive/zip"
	"io/ioutil"
)

func readClass(file *zip.File) string {
	read, err := file.Open()
	if err != nil {
		return ""
	}
	defer read.Close()

	content, err := ioutil.ReadAll(read)
	if err != nil {
		return ""
	}

	return string(content)
}

func readJar(file string) (*zip.ReadCloser, error) {
	jar, err := zip.OpenReader(file)
	if err != nil {
		return nil, err
	}
	defer jar.Close()

	return jar, nil
}
