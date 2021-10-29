package main

import (
	"fmt"
	"os"
)

type Queue struct {
	maxSize int
	arr     [4]int
	front   int
	rear    int
}

//添加
func (queue *Queue) add(val int) error {
	if queue.rear == queue.maxSize-1 {
		return fmt.Errorf("Queue full")
	}
	queue.rear += 1
	queue.arr[queue.rear] = val
	fmt.Println("添加成功！！！")
	return nil
}

//遍历
func (queue *Queue) show() error {
	fmt.Println("当前队列：")
	for i := queue.front + 1; i <= queue.rear; i++ {
		fmt.Printf("arr[%v]=%v   ", i, queue.arr[i])
	}
	fmt.Println()
	return nil
}

//获取
func (queue *Queue) get() (int, int, error) {
	if queue.front == queue.rear {
		return -1, -1, fmt.Errorf("Queue empty")
	}
	queue.front += 1
	return queue.front, queue.arr[queue.front], nil
}
func main() {
	queue := Queue{
		maxSize: 4,
		front:   -1,
		rear:    -1,
	}
	for {
		fmt.Println("[add]添加  [show]遍历  [get]获取  [exit]退出")
		fmt.Print("请输入需要功能：")
		var str string
		fmt.Scan(&str)
		switch str {
		case "add":
			fmt.Print("请输入要添加的数：")
			var val int
			fmt.Scan(&val)
			addErr := queue.add(val)
			if addErr != nil {
				fmt.Printf("addErr:%v\n", addErr)
			}
		case "show":
			queue.show()
		case "get":
			index, res, getErr := queue.get()
			if getErr != nil {
				fmt.Printf("addErr:%v\n", getErr)
			} else {
				fmt.Printf("arr[%v]=%v\n", index, res)
			}
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("输入有误！")
		}
		fmt.Println()
	}
}
