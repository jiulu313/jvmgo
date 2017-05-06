package classpath

import (
	"os"
	"strings"
)

//目录分隔符
const pathListSeparator = string(os.PathListSeparator)

//Entry ,表示类路径
type Entry interface {
	readClass(classname string) ([]byte, Entry, error)
	String() string
}

/**
	生成4种 Entry
	1 目录分隔符的
	2 通配符的
	3 zip,jar包的
	4 目录的
 */
func newEntry(path string) Entry {
	//生成路径组合的Entry
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}

	//生成通配符的Entry
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}

	//生成zip的Entry
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}

	//生成目录的Entry
	return newDirEntry(path)
}
