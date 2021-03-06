package base

import "jvmgo/jvm/rtda"

//基本指令接口
type Instruction interface {
	//从字节码中提取操作数
	FetchOperands(reader *BytecodeReader)
	
	//执行指令逻辑
	Execute(frame *rtda.Frame)
}


//没有操作指令
type NoOperandsInstruction struct {

}

func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader)  {
	//nothing to do
}

//跳转指令
type BranchInstruction struct {
	Offset 	int
}

func (self *BranchInstruction) FetchOperands(reader *BytecodeReader)  {
	self.Offset = int(reader.ReadInt16())
}


type Index8Instruction struct {
	Index 	uint
}

func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
}

type Index16Instruction struct {
	Index 	uint
}

func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint16())
}






