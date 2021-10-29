package main

import (
	"fmt"
	"os"
)

type Emp struct {
	id   int
	name string
	next *Emp
}

//展示当前调用者emp的信息
func (emp *Emp) show() {
	fmt.Printf("链表[%v]：{id:%v name:%v}\n", emp.id%7, emp.id, emp.name)
}

type EmpLink struct {
	head *Emp
}

func (empLink *EmpLink) insert(emp *Emp) {
	var pre *Emp        //pre是个辅助指针 始终在cur的前一位
	cur := empLink.head //将链表表头赋值给cur
	if cur == nil {     //如果cur为空 说明当前empLink是个空链表
		empLink.head = emp //就将emp赋值给empLink的head  head做链表表头
		return
	}
	//如果当前empLink不是空链表 就给emp找合适位置添加
	for {
		if cur.id > emp.id {
			break
		} //找到emp的位置了
		//pre和cur都向前移动一位
		pre = cur      //pre指向他下一位cur
		cur = cur.next //cur指向他的下一位  保证pre始终在cur的前一位 保持同步
		if cur == nil {
			break
		}
	}
	//当emp找到合适他的位置
	pre.next = emp //当前pre的下一位指向emp
	emp.next = cur //emp的下一位指向本来在pre前面的cur  这样就把emp插入pre和cur之间
}
func (empLink *EmpLink) show(id int) {
	//如果当前empLink的head为nil 就说明当前empLink为空链表
	if empLink.head == nil {
		fmt.Printf("链表[%v]为空！\n", id)
		return
	}
	temp := empLink.head
	fmt.Printf("链表[%v]-->", id)
	for {
		if temp != nil {
			//如果temp不为空就先输出当前temp的信息 再前移一位
			fmt.Printf("{id:%v name:%v}-->", temp.id, temp.name)
			temp = temp.next
		} else {
			break
		} //如果temp为空就退出循环 遍历结束
	}
	fmt.Println()
}
func (empLink *EmpLink) findById(id int) *Emp { //该方法返回一个Emp指针
	temp := empLink.head //先将表头赋值给temp
	for {
		if temp == nil {break} //如果temp为nil就退出循环
		if temp.id == id {return temp} //找到与传入的id相等的元素temp 并返回目标元素temp
		temp = temp.next //temp前移一位
	}
	return nil //返回nil
}

type HashTable struct {
	linkArr [7]EmpLink
}

func (hashTable *HashTable) insert(emp *Emp) {
	//先根据emp的id找到他对应的linkArr分组
	linkNumber := hashTable.hashFun(emp.id)
	//再将emp添加到他对应的linkArr分组中
	hashTable.linkArr[linkNumber].insert(emp)
}
func (hashTable *HashTable) hashFun(id int) int {
	//判断该id属于哪个linkArr分组
	return id % len(hashTable.linkArr)
}
func (hashTable *HashTable) show() {
	//遍历所有linkArr分组
	for i := 0; i < len(hashTable.linkArr); i++ {
		hashTable.linkArr[i].show(i) //第i个linkArr分组调用它的show方法
	}
}
func (hashTable *HashTable) findById(id int) *Emp {
	//先根据id找到他对应的linkArr分组
	linkNumber := hashTable.hashFun(id)
	return hashTable.linkArr[linkNumber].findById(id) //第i个linkArr分组调用它的findById方法 传入当前方法接收的id
}
func main() {
	input := 4
	id := -1
	name := "nil"
	var hashTable HashTable
	for {
		fmt.Println("***雇员系统管理***")
		fmt.Println("*     1.添加     *")
		fmt.Println("*     2.显示     *")
		fmt.Println("*     3.查找     *")
		fmt.Println("*     4.退出     *")
		fmt.Println("*****************")
		fmt.Println("请选择(1-4):")
		fmt.Scanln(&input)
		switch input {
		case 1:
			fmt.Println("请输入id:")
			fmt.Scanln(&id)
			fmt.Println("请输入name:")
			fmt.Scanln(&name)
			emp := &Emp{id: id, name: name}
			hashTable.insert(emp)
		case 2:
			hashTable.show()
		case 3:
			fmt.Println("请输入要查找雇员的id:")
			fmt.Scanln(&id)
			emp := hashTable.findById(id)
			if emp != nil {
				emp.show() //当前emp调用它的show方法
			} else {
				fmt.Printf("未找到[%v]号雇员的信息！\n", id)
			}
		case 4:
			os.Exit(0)
		default:
			fmt.Println("输入有误，请重新输入。")
		}
	}
}
