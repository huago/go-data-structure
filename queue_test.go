package tboxs

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	// 初始化queue
	q := InitQueue(4)

	// 循环入队列4个数
	for i := 10; i < 14; i++ {
		err := q.Enter(i)
		if err != nil {
			_ = fmt.Errorf("enter error=%v\n", err)
		}
	}

	// 获取队首元素
	head := q.Head()
	fmt.Printf("head=%v\n", head)

	// 出队列一个元素
	element := q.Out()
	fmt.Printf("element=%v\n", element)

	// 获取队列的长度
	length := q.Length()
	fmt.Printf("length=%v\n", length)

	// 清空队列
	fmt.Println("clear queue start...")
	q.Clear()
	fmt.Println("clear queue end...")

	// 判断队列是否为空
	fmt.Printf("empty=%v\n", q.IsEmpty())
}
