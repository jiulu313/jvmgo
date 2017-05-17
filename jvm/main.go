package main

import (
	"fmt"
	"jvmgo/jvm/classpath"
	"strings"
	"jvmgo/jvm/classfile"
)

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption,cmd.cpOption)
	className := strings.Replace(cmd.class,".","/",-1)
	cf := loadClass(className,cp)
	fmt.Println(cmd.class)
	printClassInfo(cf)
}


//读取并解析 class 文件
func loadClass(className string,cp *classpath.Classpath) *classfile.ClassFile  {
	classData, _ , err := cp.ReadClass(className)
	if err != nil{
		panic(err)
	}
	
	cf , err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}
	
	return cf
}

//打印一些重要的class文件的信息
func printClassInfo(cf *classfile.ClassFile) {
	fmt.Printf("version:%v.%v\n",cf.MajorVersion(),cf.MinorVersion())
	fmt.Printf("constants count: %v\n",len(cf.ConstantPool()))
	fmt.Printf("access flags: 0x%x\n",cf.AccessFlag())
	fmt.Printf("this class: %v\n",cf.ClassName())
	fmt.Printf("super class: %v\n",cf.SuperClassName())
	fmt.Printf("interfaces: %v\n",cf.InterfaceNames())
	fmt.Printf("fields count: %v\n",len(cf.Fields()))
	
	for _,m := range cf.Methods() {
		fmt.Printf("  %s\n",m.Name)
	}
}