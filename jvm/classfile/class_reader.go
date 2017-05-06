package classfile

import "encoding/binary"


//ClassReader 读取class文件
type ClassReader struct {
	data []byte
}

//读取 1 个字节
func (self *ClassReader) readUint8() uint8 {
	val := self.data[0]
	self.data = self.data[1:]
	return val
}

//读取 2 个字节
func (self *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:]
	return val
}

//读取 4 个字节
func (self *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:]
	return val
}

//读取 8 个字节
func (self *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:]
	return val
}

//读取uint16表(就是读取uint16 数组)
func (self *ClassReader) readUint16s() []uint16 {
	n := self.readUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = self.readUint16()
	}
	return s
}

//读取指定数量的字节
func (self *ClassReader)readBytes(n uint32) []byte  {
	bytes := self.data[:n]
	self.data = self.data[n:]
	return bytes
}



