package main

import (
	"errors"
	"fmt"
	"strconv"
)

/**
栈的计算表达式原理
1.创建两个栈：numberStack栈(数字栈)存放数字  operatorStack栈(运算符栈)存放运算符。
2.exp是计算表达式，是一个字符串。
3.index是字符串的下标，从index=0开始扫描计算表达式，每次扫描exp的其中一个字符，扫描完一个字符index++一次扫描下一位。
4.如果扫描到的是数字，就直接入数字栈。
5.如果扫描到的是运算符：A:如果operatorStack栈为空就直接入栈。
					  B:如果operatorStack栈不为空：a:如果operatorStack栈栈顶的运算符优先级大于等于当前扫描到的运算符优先级，就从运算符栈pop出栈顶的运算符，并从数字栈pop出两个数字，
													用后出的数字去运算先出的数字；运算后将结果push到数字栈中，并将当前扫描到的运算符push到运算符栈中。
												  b:如果operatorStack栈栈顶的运算符优先级小于当前扫描到的运算符优先级，直接将当前扫描到的运算符push到运算符栈中。
6.exp计算表达式扫描完毕：[先从数字栈pop出两个数字，再从运算符栈pop出栈顶的运算符，用后出的数字去运算先出的数字；运算后将结果push到数字栈中]
					   不断重复以上操作，直到运算符栈为空；最后数字栈中栈顶的元素就是计算表达式的运算结果。
*/
type Stack struct {
	maxTop int     //栈顶最大值
	top    int     //栈顶
	bottom int     //栈底
	slice  [20]int //用数组模拟栈
}

func (stack *Stack) push(val int) {
	if stack.top >= stack.maxTop-1 {
		fmt.Println("stack full")
		return
	}
	stack.top++
	stack.slice[stack.top] = val
}

func (stack *Stack) pop() (error, int) {
	if stack.top == stack.bottom {
		return errors.New("stack empty"), -1
	}
	val := stack.slice[stack.top]
	stack.top--
	return nil, val
}

func (stack *Stack) show() {
	if stack.top == stack.bottom {
		fmt.Println("stack empty")
		return
	}
	for i := stack.top; i > stack.bottom; i-- {fmt.Printf("stack：slice[%v] = %v\n", i, stack.slice[i])}
}

//判断传入值是不是一个运算符[+ - * /]
func (stack *Stack) isOperator(val int) bool {
	if val == 42 || val == 43 || val == 45 || val == 47 {
		return true
	} else {
		return false
	}
}

//判断运算符的优先级
func (stack *Stack) operatorGrade(operator int) int {
	if operator == 42 || operator == 47 {return 1}
	if operator == 43 || operator == 45 {return 0}
	fmt.Printf("无法识别运算符[%v]\n", operator)
	return -1
}

//接收数字和运算符 返回运算结果
func (stack *Stack) cal(num2 int, operator int, num1 int) int {
	res := -1
	switch operator {
	case 42:
		res = num2 * num1
	case 43:
		res = num2 + num1
	case 45:
		res = num2 - num1
	case 47:
		res = num2 / num1
	default:
		fmt.Printf("无法识别运算符[%v]\n", operator)
	}
	return res
}

func main() {
	numberStack := &Stack{maxTop: 20, top: -1, bottom: -1}   //数字栈
	operatorStack := &Stack{maxTop: 20, top: -1, bottom: -1} //运算符栈
	exp := "30+30*6-4-5+-"
	index := 0
	number1 := -1
	number2 := -1
	operator := -1
	keepNumber := ""
	//开始扫描计算表达式 当exp中的所有元素扫描完就退出循环
	for index < len(exp) {
		//获取exp的第index个元素(字符)
		ch := exp[index : index+1]
		//先将ch转为切片类型 切片类型中就存储ch的int值 取切片的第一位因为编译器不知道ch转换后的切片中有几个元素 最后再将切片的第一个元素转为int
		intCh := int([]byte(ch)[0])
		//判断intCh是不是运算符 不是运算符就是数字
		if operatorStack.isOperator(intCh) { //是运算符
			if operatorStack.top == operatorStack.bottom { //先判断运算符栈是不是空栈 是就入栈
				operatorStack.push(intCh) //入栈
			} else { //如果运算符栈不是空栈
				//比较当前运算符和栈顶运算符的优先级
				if operatorStack.operatorGrade(operatorStack.slice[operatorStack.top]) >= operatorStack.operatorGrade(intCh) { //如果栈顶的运算符的优先级大于等于当前运算符优先级
					//就从数字栈里pop两个数 用后出的数去运算先出的数(number2 operator number1)
					_, number1 = numberStack.pop()
					_, number2 = numberStack.pop()
					_, operator = operatorStack.pop()
					calRes := operatorStack.cal(number2, operator, number1)
					//运算完成后
					numberStack.push(calRes)  //将运算结果push到数字栈中
					operatorStack.push(intCh) //将当前运算符push到运算符栈中
				} else { //如果栈顶的运算符的优先级小于当前运算符优先级
					//就将当前运算符push到运算符栈中
					operatorStack.push(intCh) //将当前运算符push到运算符栈中
				}
			}
		} else { //是数字 直接将当前keepNumber的int类型push到数字栈
			keepNumber += ch         //拼接keepNumber  将当前ch追加到keepNumber中
			if index == len(exp)-1 { //先判断当前ch是否是exp的最后一位 如果是就将当前keepNumber push到数字栈
				//因为此时的keepNumber不是具体int类型的数 而是字符串 需要将他转换为int类型的数再push
				val, _ := strconv.ParseInt(keepNumber, 10, 64)
				//返回的val是int64类型的 还要转换为int类型再传入
				numberStack.push(int(val))
			} else { //如果不是最后一位 就判断下一位是不是运算符 如果是就push当前的keepNumber 不是就继续扫描并拼接keepNumber
				if operatorStack.isOperator(int([]byte(exp[index+1 : index+2])[0])) {
					val, _ := strconv.ParseInt(keepNumber, 10, 64)
					numberStack.push(int(val))
					//push完keepNumber记得要清空keepNumber 防止下次使用时keepNumber残留数据
					keepNumber = ""
				}
			}
		}
		//完成以上操作后 index 前移一位继续处理
		index++
	}
	//如果扫描完毕 依次从符号栈取出符号 从数字栈取两个数 用后出的数去运算先出的数(number2 operator number1)
	//继续进行运算 运算完成后结果入数字栈 直到运算符栈为空
	for operatorStack.bottom < operatorStack.top {
		_, number1 = numberStack.pop()
		_, number2 = numberStack.pop()
		_, operator = operatorStack.pop()
		calRes := operatorStack.cal(number2, operator, number1)
		numberStack.push(calRes) //将运算结果push到数字栈中
	}
	//以上操作如果无误 数字栈中只剩一个栈顶的元素 就是运算结果
	_, res := numberStack.pop()
	fmt.Printf("%v = %v \n", exp, res)
}
