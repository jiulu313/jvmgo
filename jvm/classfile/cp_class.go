package classfile

type ConstantClassInfo struct {
	cp        ConstantPool
	nameIndex uint16
}

//读取索引
func (self *ConstantClassInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
}

//从常量池中查找
func (self *ConstantClassInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}
