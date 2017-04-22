package classpath

import (
	"errors"
	"strings"
)

//CompositeEntry Entry数组
type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	compsiteEntry := []Entry{}
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compsiteEntry = append(compsiteEntry, entry)
	}

	return compsiteEntry
}

//读取类文件
func (self CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range self {
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, nil
		}
	}

	return nil, nil, errors.New("class not found: " + className)
}

//把CompositeEntry中的每个entry中的String()返回的字符拼接起来即可
func (self CompositeEntry) String() string {
	strs := make([]string, len(self))
	for i, entry := range self {
		strs[i] = entry.String()
	}

	return strings.Join(strs, pathListSeparator)
}
