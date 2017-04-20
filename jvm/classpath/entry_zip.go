package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

//ZipEntry struct
type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}

func (itself *ZipEntry) String() string {
	return itself.absPath
}

//从zip文件中读取
func (itself *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(itself.absPath)
	if err != nil {
		panic(err)
	}

	defer r.Close()
	for _, f := range r.File {
		if f.Name == className {
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}

			defer rc.Close()
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}

			return data, itself, nil
		}
	}

	return nil, nil, errors.New("class not found: " + className)
}
