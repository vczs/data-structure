package main

import (
	"fmt"
	"os"
)

type Circle struct {
	maxSize int    //最大长度
	arr     [4]int //数组
	head    int    //头部
	tail    int    //尾部
}

//判断队列是否满
func (circle *Circle) isFull() bool {
	return (circle.tail+1)%circle.maxSize == circle.head
}

//判断队列是否空
func (circle *Circle) isEmpty() bool {
	return circle.head == circle.tail
}

//计算队列长度
func (circle *Circle) size() int {
	return (circle.tail + circle.maxSize - circle.head) % circle.maxSize
}

//放入元素
func (circle *Circle) push(val int) error {
	if circle.isFull() {
		return fmt.Errorf("Arr full\n")
	}
	circle.arr[circle.tail] = val
	circle.tail = (circle.tail + 1) % circle.maxSize
	fmt.Println("添加成功！！！")
	return nil
}

//取出元素
func (circle *Circle) pop() (int, int, error) {
	if circle.isEmpty() {return -1, -1, fmt.Errorf("Arr rmpty\n")}
	val := circle.arr[circle.head]
	index := circle.head
	circle.head = (circle.head + 1) % circle.maxSize
	return index, val, nil
}

//遍历队列
func (circle *Circle) list() {
	fmt.Print("环形队列遍历：")
	tempHead := circle.head
	for i := 0; i < circle.size(); i++ {
		fmt.Printf("arr[%v]=%v  ", tempHead, circle.arr[tempHead])
		tempHead = (tempHead + 1) % circle.maxSize
	}
	fmt.Println()
}
func main() {
	circle := &Circle{
		maxSize: 4,
		tail:    0,
		head:    0,
	}
	for {
		fmt.Println("[push]添加  [list]遍历  [pop]获取  [exit]退出")
		fmt.Print("请输入需要功能：")
		var str string
		fmt.Scan(&str)
		switch str {
		case "push":
			fmt.Print("请输入要添加的数：")
			var val int
			fmt.Scan(&val)
			addErr := circle.push(val)
			if addErr != nil {fmt.Printf("addErr:%v\n", addErr)}
		case "list":
			circle.list()
		case "pop":
			index, res, getErr := circle.pop()
			if getErr != nil {
				fmt.Printf("addErr:%v\n", getErr)
			} else {
				fmt.Printf("已取出：arr[%v]=%v\n", index, res)
			}
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("输入有误！")
		}
		fmt.Println()
	}
}
