package classfile

import (
	"fmt"
	"io"
)

//ClassFile  class文件的数据结构
type ClassFile struct {
	magic 			uint32
	minorVersion	uint16
	majorVersion	uint16
	constantPool	ConstantPool
	accessFlag		uint16
	thisClass		uint16
	superClass		uint16
	interfaces		[]uint16
	//fields			[]*MemberInfo
	//methods			[]*MemberInfo
	//attributes		[]AttributeInfo
}

//把[]byte解析成ClassFile结构体
func Parse(classData []byte) (cf *ClassFile,err error)  {
	defer func() {
		if r := recover(); r != nil{
			var ok bool
			err,ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v",r)
			}
		}
	}()
	
	classReader := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(classReader)
	return
}

func (itself *ClassFile)read(reader *ClassReader)  {
	itself.readAndCheckMagic(reader)
	itself.readAndCheckVersion(reader)
	itself.constantPool = readConstantPool(reader)
	itself.accessFlag = reader.readUint16()
	itself.thisClass = reader.readUint16()
	itself.superClass = reader.readUint16()
	itself.interfaces = reader.readUint16s()
	itself.fields = readMembers(reader,itself.constantPool)
	itself.methods = readMembers(reader,itself.constantPool)
	itself.attributes = readAttributes(reader,itself.constantPool)
}

func (itself *ClassFile)MajorVersion() uint16 {
	return itself.majorVersion
}

func (itself *ClassFile)MinorVersion() uint16 {
	return itself.minorVersion
}

func (itself *ClassFile)AccessFlag() uint16 {
	return itself.accessFlag
}

func (itself *ClassFile)ThisClass() uint16{
	return itself.thisClass
}

func (itself *ClassFile)SuperClass() uint16  {
	return itself.superClass
}

//获取本类类名
func (itself *ClassFile)ClassName() string  {
	return itself.constantPool.getClassName(itself.thisClass)
}

//获取超类类名
func (itself *ClassFile)SuperClassName() string  {
	if itself.superClass > 0 {
		return itself.constantPool.getClassName(itself.superClass)
	}
	return "" //只有java.lang.Object，没有超类
}













