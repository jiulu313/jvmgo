package classfile

import (
	"fmt"
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

func (self *ClassFile)read(reader *ClassReader)  {
	self.readAndCheckMagic(reader)
	self.readAndCheckVersion(reader)
	self.constantPool = readConstantPool(reader)
	self.accessFlag = reader.readUint16()
	self.thisClass = reader.readUint16()
	self.superClass = reader.readUint16()
	self.interfaces = reader.readUint16s()
	self.fields = readMembers(reader,self.constantPool)
	self.methods = readMembers(reader,self.constantPool)
	self.attributes = readAttributes(reader,self.constantPool)
}

/*
 *	魔数，文件的格式都会以某种字节开头，.class文件魔数是 0xCAFEBABE
 */
func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	self.magic = reader.readUint32()
	if self.magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError")
	}
}

/*
 *	魔数之后是class文件的次版本号和主版本号，都是u2类型。
 * 	假如主版本是M，次版本是m，那么完整的版本号可以表示成 M.m
 *  次版本号在 J2SE 1.2 之前用过，1.2开始就基本没有什么用了（都是0）
 *  主版本号在 J2SE 1.2 之前都是45，每次有大的Java版本发布，都会加 1
 *	下面列出了到目前为止，使用过的class文件版本号
 *	JDK 1.0.2 	45.0 -- 45.3
 *	JDK 1.1 	45.0 -- 45.65535
 *	J2SE 1.2 	46.0
 *	J2SE 1.3 	47.0
 *	J2SE 1.4 	48.0
 *	J2SE 5.0 	49.0
 *	J2SE 6	 	50.0
 *	J2SE 7	 	51.0
 *	J2SE 8	 	52
 */
func (self *ClassFile)readAndCheckVersion(reader *ClassReader) {
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()
	
	switch self.majorVersion {
	case 45:
		return
	case 46,47,48,49,50,51,52:
		return
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

//返回次版本号
func (self *ClassFile)MinorVersion() uint16 {
	return self.minorVersion
}

//返回主版本号
func (self *ClassFile)MajorVersion() uint16 {
	return self.majorVersion
}

func (self *ClassFile)AccessFlag() uint16 {
	
	
	
	
	return self.accessFlag
}

func (self *ClassFile)ThisClass() uint16{
	return self.thisClass
}

func (self *ClassFile)SuperClass() uint16  {
	return self.superClass
}

//获取本类类名
func (self *ClassFile)ClassName() string  {
	return self.constantPool.getClassName(self.thisClass)
}

//获取超类类名
func (self *ClassFile)SuperClassName() string  {
	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)
	}
	return "" //只有java.lang.Object，没有超类
}

//从常量池中查找接口名
func (self *ClassFile)InterfaceNames() []string{
	interfaceNames := make([]string,len(self.interfaces))
	for i,cpIndex := range self.interfaces{
		interfaceNames[i] = self.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}













