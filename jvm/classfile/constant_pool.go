package classfile

//常量池
type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool {
	return nil
}

func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	return nil
}

func (self ConstantPool) getNameAndType(index uint16) (string,string) {
	return nil
}

func (self ConstantPool) getClassName(index uint16) string  {
	return nil
}

func (self ConstantPool) getUtf8(index uint16) string  {
	return nil
}


