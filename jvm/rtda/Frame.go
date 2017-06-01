package rtda

type Frame struct {
	lower 			*Frame			//用于实现链表数据结构
	localVars		LocalVars		//局部变量表指针
	operandStack	*OperandStack	//保存操作数栈指针
	
	/**
		执行方法所需要的局部变量表大小和操作数栈深度是由编译器预先计算好的
		存储在class文件method_info结构的Code属性中
 	*/
}


//创建栈实例
func NewFrame(maxLocals uint,maxStack uint) *Frame{
	return &Frame{
		localVars: newLocalVars(maxLocals),
		operandStack:newOperandStack(maxStack),
	}
}










