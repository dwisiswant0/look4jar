package main

import (
	"os"
	"regexp"

	"archive/zip"
	"path/filepath"
)

func isVuln(jar *zip.ReadCloser) bool {
	var mgr *zip.File

	for _, class := range jar.File {
		if isStr(class.Name, FATAL_CLASS[0]) && isByt(readClass(class), SIGNATURES[:7]) {
			return true
		}

		if isStr(class.Name, FATAL_CLASS[1]) {
			mgr = class
		}
	}

	if mgr != nil && isByt(readClass(mgr), SIGNATURES[7:]) {
		return true
	}

	return false
}

func lookup(path string) ([]string, error) {
	var files []string

	file, err := os.Open(path)
	if err != nil {
		return files, err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return files, err
	}

	if !stat.IsDir() {
		path, _ = filepath.Abs(path)
		files = append(files, path)
	} else {
		jar, _ := regexp.Compile(JAR_PATTERN)
		err = filepath.Walk(path, func(fp string, fi os.FileInfo, fe error) error {
			if fe != nil {
				return fe
			}

			if !fi.IsDir() && jar.MatchString(fi.Name()) {
				fp, _ = filepath.Abs(fp)
				files = append(files, fp)
			}

			return nil
		})

		if err != nil {
			return files, err
		}
	}

	return files, nil
}
