package classpath

import (
	"os"
	"path/filepath"
)

//Classpath struct
type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

//Parse 解析
func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

//parseBootAndExtClasspath
func (itself *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)

	// jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	itself.bootClasspath = newWildcardEntry(jreLibPath)

	// jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	itself.extClasspath = newWildcardEntry(jreExtPath)
}

//解析用户自定义的,如果用户没有提供-cp,则使用当前目录作为用户类路径
func (itself *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	itself.userClasspath = newEntry(cpOption)
}

//ReadClass 注意 className 不能包含 .class
func (itself *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"

	if data, entry, err := itself.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	if data, entry, err := itself.extClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	return itself.userClasspath.readClass(className)
}

//String 返回用户类路径
func (itself *Classpath) String() string {
	return itself.userClasspath.String()
}

/*
	优先使用用户输入的 -Xjre 选项作为jre目录.
	如果没有输入该选项,则在当前目录下寻找jre目录
	如果找不到,尝试用JAVA_HOME环境变量
*/
func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}

	if exists("./jre") {
		return "./jre"
	}

	if javaHome := os.Getenv("JAVA_HOME"); javaHome != "" {
		return filepath.Join(javaHome, "jre")
	}

	panic("Can not find jre folder")
}

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
