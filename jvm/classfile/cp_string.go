package classfile

/*
	本身并不存储字符串数据，只是存储了常量池索引
*/
type ConstantStringInfo struct {
	cp 	ConstantPool
	stringIndex uint16
}

//读取常量池索引
func (self *ConstantStringInfo) readInfo(reader *ClassReader)  {
	self.stringIndex = reader.readUint16()
}

//从常量池中查找字符串
func (self *ConstantStringInfo) String() string  {
	return self.cp.getUtf8(self.stringIndex)
}




