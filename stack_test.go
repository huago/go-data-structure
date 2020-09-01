package tboxs

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	// 初始化stack
	s := InitStack(4)

	// 循环入队列4个数
	for i := 10; i < 14; i++ {
		err := s.Push(i)
		if err != nil {
			_ = fmt.Errorf("push error=%v\n", err)
		}
	}

	// 获取栈顶元素
	head := s.Peek()
	fmt.Printf("head=%v\n", head)

	// 出栈一个元素
	element := s.Pop()
	fmt.Printf("element=%v\n", element)

	// 获取栈的长度
	length := s.Length()
	fmt.Printf("length=%v\n", length)

	// 清空栈
	fmt.Println("clear stack start...")
	s.Clear()
	fmt.Println("clear stack end...")

	// 判断栈是否为空
	fmt.Printf("empty=%v\n", s.IsEmpty())
}
