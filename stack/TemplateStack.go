package stack

// Stack 定义一个数组栈，设置泛型传入不同类型的对象
// 其实直接用一个切片去定义数组最好了，因为会自动扩容对吧。
type Stack[T any] struct {
	data []T
}

// NewStack 获取一个Stack类型的对象
// 返回一个Stack[T]的指针类型的对象
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

// Push 栈的压栈操作
func (this *Stack[T]) Push(x T) {
	this.data = append(this.data, x)
}

// Pop 栈的弹出操作
func (this *Stack[T]) Pop() T {
	if len(this.data) == 0 {
		panic("stack is empty")
	}
	x := this.data[this.Len()-1]
	this.data = this.data[:this.Len()-1]
	return x
}

// Len 返回切片的长度
func (this *Stack[T]) Len() uint {
	return uint(len(this.data))
}
