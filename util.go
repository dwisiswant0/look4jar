package main

import (
	"os"
	"regexp"

	"archive/zip"
	"path/filepath"
)

func isVuln(jar *zip.ReadCloser) bool {
	for _, class := range jar.File {
		i, s := 0, 0
		for s < len(SIGNATURES) {
			status, e := false, s+7
			if isStr(class.Name, FATAL_CLASS[i]) {
				status = !isByt(readClass(class), SIGNATURES[s:e])
				if i == 0 {
					status = isByt(readClass(class), SIGNATURES[s:e])
				}
			}

			if status {
				return true
			}

			i, s = i+1, s+7
		}
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
