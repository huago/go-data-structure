package datastructure

import (
	"errors"
	"sync"
)

type Queue struct {
	data         []interface{}
	length       int
	capacity     int
	initCapacity int
	sync.Mutex
}

// 初始化队列
func InitQueue(cap int) *Queue {
	return &Queue{
		data:         make([]interface{}, cap),
		length:       0,
		capacity:     cap,
		initCapacity: cap,
	}
}

// 入队列
func (q *Queue) Enter(data interface{}) error {
	q.Lock()
	defer q.Unlock()

	if data == nil {
		return errors.New("data invalid")
	}

	if q.length+1 >= q.capacity {
		q.capacity <<= 1
		t := q.data
		q.data = make([]interface{}, q.capacity)
		copy(q.data, t)
	}

	q.data[q.length] = data
	q.length++

	return nil
}

// 出队列
func (q *Queue) Out() interface{} {
	q.Lock()
	defer q.Unlock()

	if q.length <= 0 {
		return nil
	}

	t := q.data[0]
	q.data = q.data[1:]
	q.length--

	return t
}

// 获取队列头元素
func (q *Queue) Head() interface{} {
	q.Lock()
	defer q.Unlock()

	if q.length <= 0 {
		return nil
	}

	return q.data[0]
}

// 获取队列长度
func (q *Queue) Length() int {
	q.Lock()
	defer q.Unlock()

	return q.length
}

// 清空队列
func (q *Queue) Clear() {
	q.Lock()
	defer q.Unlock()

	q.data = make([]interface{}, q.initCapacity)
	q.length = 0
	q.capacity = q.initCapacity
}

// 判断队列是否为空
func (q *Queue) IsEmpty() bool {
	q.Lock()
	defer q.Unlock()

	return q.length == 0
}
