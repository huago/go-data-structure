package tboxs

import (
	"errors"
	"sync"
)

type Stack struct {
	data         []interface{}
	length       int
	capacity     int
	initCapacity int
	sync.Mutex
}

// 初始化栈
func InitStack(cap int) *Stack {
	return &Stack{
		data:         make([]interface{}, cap),
		length:       0,
		capacity:     cap,
		initCapacity: cap,
	}
}

// 入栈
func (s *Stack) Push(data interface{}) error {
	s.Lock()
	defer s.Unlock()

	if data == nil {
		return errors.New("data invalid")
	}

	if s.length+1 >= s.capacity {
		s.capacity <<= 1

		t := s.data
		s.data = make([]interface{}, s.capacity)
		copy(s.data, t)
	}

	s.data[s.length] = data
	s.length++

	return nil
}

// 出栈
func (s *Stack) Pop() interface{} {
	s.Lock()
	defer s.Unlock()

	if s.length <= 0 {
		return nil
	}

	t := s.data[s.length-1]
	s.data = s.data[:s.length-1]

	s.length--

	return t
}

// 获取栈顶元素
func (s *Stack) Peek() interface{} {
	s.Lock()
	defer s.Unlock()

	if s.length <= 0 {
		return nil
	}

	return s.data[s.length-1]
}

// 获取栈长度
func (s *Stack) Length() int {
	s.Lock()
	defer s.Unlock()

	return s.length
}

// 清空栈
func (s *Stack) Clear() {
	s.Lock()
	defer s.Unlock()

	s.data = make([]interface{}, s.initCapacity)
	s.length = 0
	s.capacity = s.initCapacity
}

// 判断栈是否为空
func (s *Stack) IsEmpty() bool {
	s.Lock()
	defer s.Unlock()

	return s.length == 0
}
