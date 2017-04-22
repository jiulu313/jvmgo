package classpath

import (
	"io/ioutil"
	"path/filepath"
)

//DirEntry struct
type DirEntry struct {
	absDir string
}

//把参数转换成绝对路径
func newDirEntry(path string) *DirEntry {
	absDir, error := filepath.Abs(path)
	if error != nil {
		panic(error)
	}
	return &DirEntry{absDir}
}

func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	filename := filepath.Join(self.absDir, className)
	data, err := ioutil.ReadFile(filename)
	return data, self, err
}

func (self *DirEntry) String() string {
	return self.absDir
}
