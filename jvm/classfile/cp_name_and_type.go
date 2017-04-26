package classfile

//字段或者方法的名称和描述符
type ConstantNameAndTypeInfo struct {
	nameIndex 	uint16
	descriptorIndex uint16
}

func (self *ConstantNameAndTypeInfo) readInfo(reader *ClassReader)  {
	self.nameIndex = reader.readUint16()
	self.descriptorIndex = reader.readUint16()
}