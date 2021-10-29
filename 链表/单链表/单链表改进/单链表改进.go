package main

import "fmt"

type Cat struct {
	id   int
	name string
	next *Cat
}

//将传入的cat 按照cat的id从小到大的顺序插入到链表中的指定位置
func addCat(head *Cat, newCat *Cat) {
	temp := head
	for {
		if temp.next == nil {
			break //如果链表到末尾就退出循环
		} else if temp.next.id > newCat.id {
			break //如果temp的下一个cat的id大于要插入的cat的id   就表示newCat应在当前temp和temp.next指向的cat之间
		} else if temp.next.id == newCat.id { //如果链表中有和newCat的id相等的cat 就不插入newCat
			fmt.Printf("id为[%v]的Cat已存在名为[%v]  当前%v添加失败！\n", temp.next.id, temp.next.name, *newCat)
			return
		}
		temp = temp.next //若temp.next不为空 就将temp的值更新为当前temp指向的下一个people
	}
	newCat.next = temp.next //将当前temp的next指向的cat 赋值给newCat的next字段
	temp.next = newCat      //再将newCat赋值给当前temp的next 即就完成了插入
}

//遍历
func showCat(head *Cat) {
	temp := head
	fmt.Print("当前链表：head-->")
	for {
		if temp.next == nil {
			break
		} //如果temp.next表示到链表末尾 就退出循环
		fmt.Printf("[%v %v]-->", temp.next.id, temp.next.name)
		temp = temp.next //若temp.next不为空 就将temp的值更新为当前temp指向的下一个people
	}
	fmt.Println()
}

//删除
func deleteCat(head *Cat, id int) {
	temp := head
	for {
		if temp.next == nil {
			fmt.Printf("链表中没有id为[%v]的元素 删除失败！\n", id)
			return //如果链表到末尾表示链表中没有id为传入值的cat 退出函数
		} else if temp.next.id == id {
			break //找到目标元素 退出循环
		}
		temp = temp.next //若temp.next不为空 就将temp的值更新为当前temp指向的下一个people
	}
	temp.next = temp.next.next //将当前temp.next指向的cat 修改为temp.next指向的cat的next指向的cat  即就从链表中删除了当前temp.next指向的cat
}

//修改
func updateCat(head *Cat, newCat *Cat) {
	temp := head
	for {
		if temp.next == nil {
			fmt.Printf("链表中没有id为[%v]的元素 修改失败！\n", newCat.id)
			return //如果链表到末尾表示链表中没有id为传入值的cat 退出函数
		} else if temp.next.id == newCat.id {
			break //找到目标元素 退出循环
		}
		temp = temp.next //若temp.next不为空 就将temp的值更新为当前temp指向的下一个people
	}
	newCat.next = temp.next.next //将当前temp.next指向的cat元素的next字段指向的cat赋值给newCat.next
	temp.next = newCat           //再将当前temp.next指向的cat元素修改为newCat  即就将链表中与newCat的id重复的元素修改为newCat
}
func main() {
	head := &Cat{}
	cat1 := &Cat{id: 1, name: "一"}
	cat2 := &Cat{id: 2, name: "二"}
	cat3 := &Cat{id: 3, name: "三"}
	cat4 := &Cat{id: 3, name: "四"}
	cat5 := &Cat{id: 1, name: "五"}
	cat6 := &Cat{id: 4, name: "六"}
	addCat(head, cat3)
	addCat(head, cat1)
	addCat(head, cat2)
	addCat(head, cat4) //cat4的id与cat3的id重复 检查addCat()查重代码逻辑
	showCat(head)

	deleteCat(head, 2) //删除id为2的元素
	deleteCat(head, 4) //删除链表中不存在的id为4的元素  检查addCat()查找代码逻辑
	showCat(head)

	updateCat(head, cat5) //将链表中与cat5的id相同的元素修改为cat5
	updateCat(head, cat6) //修改链表中不存在的id为4的元素  检查addCat()查找代码逻辑
	showCat(head)         //检查以上代码执行结果
}
