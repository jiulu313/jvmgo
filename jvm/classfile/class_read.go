package classfile

import "encoding/binary"

type ClassReader struct {
	data []byte;
}

//读取 1 个字节
func (itself *ClassReader) readUint8() uint8 {
	val := itself.data[0]
	itself.data = itself.data[1:]
	return val
}

//读取 2 个字节
func (itself *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(itself.data)
	itself.data = itself.data[2:]
	return val
}

//读取 4 个字节
func (itself *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(itself.data)
	itself.data = itself.data[4:]
	return val
}

//读取 8 个字节
func (itself *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(itself.data)
	itself.data = itself.data[8:]
	return val
}

//读取uint16表
func (itself *ClassReader) readUint16s() []uint16 {
	n := itself.readUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = itself.readUint16()
	}
	return s
}

//读取指定数量的字节
func (itself *ClassReader)readBytes(n uint32) []byte  {
	bytes := itself.data[:n]
	itself.data = itself.data[n:]
	return bytes
}



