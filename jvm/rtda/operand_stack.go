package rtda

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


