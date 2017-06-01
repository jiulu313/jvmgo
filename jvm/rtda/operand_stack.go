package rtda

import (
	"math"
	"os/signal"
)

//操作数栈
//操作数栈的大小是编译器已经确定的，所以可以用[]Slot 实现
// size 用于记录栈顶位置
type OperandStack struct {
	size 	uint
	slots 	[]Slot
}

func newOperandStack(maxStack uint) *OperandStack  {
	if maxStack > 0 {
		return &OperandStack{
			slots:make([]Slot,maxStack),
		}
	}
	
	return nil
}

func (self *OperandStack) PushInt(val int32)  {
	self.slots[self.size].num = val
	self.size++
}

func (self *OperandStack) PopInt() int32 {
	self.size--
	return self.slots[self.size].num
}

func (self *OperandStack) PushFloat(val float32) {
	bits := math.Float32bits(val)
	self.slots[self.size].num = int32(bits)
	self.size++
}

func (self *OperandStack) PopFloat() float32  {
	self.size--
	bits := uint32(self.slots[self.size].num)
	return math.Float32frombits(bits)
}














