package classfile


//MemberInfo 属性和方法结构体
type MemberInfo struct {
	cp			ConstantPool
	accessFlags	uint16
	nameIndex	uint16
	descriptorIndex	uint16
	attributes 		[]AttributeInfo
}

//读取属性和方法
func readMembers(reader *ClassReader,cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo,memberCount)
	for i := range members{
		members[i] = readMember(reader,cp)
	}
	return members
}

//读取单个属性
func readMember(reader *ClassReader,cp ConstantPool) *MemberInfo {
	return &MemberInfo {
		cp:	cp,
		accessFlags:reader.readUint16(),
		nameIndex:reader.readUint16(),
		attributes:readAttributes(reader,cp),
	}
}

//?
func (self *MemberInfo) AccessFlags() uint16 {
	return 0
}

//?
func (self *MemberInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}

//?
func (self *MemberInfo) Descriptor() string {
	return self.cp.getUtf8(self.descriptorIndex)
}