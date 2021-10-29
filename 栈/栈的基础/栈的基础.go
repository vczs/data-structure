package main

import (
	"errors"
	"fmt"
)

type MyStack struct {
	maxTop int    //栈顶最大值
	top    int    //栈顶
	bottom int    //栈底
	slice  [5]int //用数组模拟栈
}

func (myStack *MyStack) push(val int) {
	if myStack.top >= myStack.maxTop-1 {
		fmt.Println("myStack full")
		return
	}
	myStack.top++
	myStack.slice[myStack.top] = val
}

func (myStack *MyStack) pop() (error, int) {
	if myStack.top == myStack.bottom {
		return errors.New("myStack empty"), -1
	}
	val := myStack.slice[myStack.top]
	myStack.top--
	return nil, val
}

func (myStack *MyStack) show() {
	if myStack.top == myStack.bottom {
		fmt.Println("myStack empty")
		return
	}
	for i := myStack.top; i > myStack.bottom; i-- {
		fmt.Printf("myStack：slice[%v] = %v\n", i, myStack.slice[i])
	}
}

func main() {
	myStack := &MyStack{maxTop: 5, top: -1, bottom: -1}
	myStack.push(1)
	myStack.push(2)
	myStack.push(3)
	myStack.push(4)
	_, val := myStack.pop()
	fmt.Printf("pop : %v\n", val)
	myStack.show()
}
