package classfile

type ConstantMemberrefInfo struct {
	cp					ConstantPool
	classIndex			uint16
	nameAndTypeIndex 	uint16
}

func (self *ConstantMemberrefInfo) readInfo(reader *ClassReader)  {
	self.classIndex = reader.readUint16()
	self.nameAndTypeIndex = reader.readUint16()
}

func (self *ConstantMemberrefInfo) ClassName() string {
		return self.cp.getUtf8(self.classIndex)
}

func (self *ConstantMemberrefInfo) NameAndDescriptor()string  {
	return self.cp.getUtf8(self.nameAndTypeIndex)
}

/**
	以下是定义了3个结构体，继承自ConstantMemberrefInfo
 */
type ConstantFieldrefInfo struct {
	ConstantMemberrefInfo
}

type ConstantMethodrefInfo struct {
	ConstantMemberrefInfo
}

type ConstantInterfaceMethodrefInfo struct {

}




