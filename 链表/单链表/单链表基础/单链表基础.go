package main

import "fmt"

type People struct {
	id   int
	name string
	next *People
}

//向链表末尾添加一个元素
func addPeople(head *People, newPeople *People) {
	temp := head //将链表头给temp
	for {
		//从temp开始 遍历链表 当其中一个temp的next为nil时说明遍历完毕 退出循环
		if temp.next == nil {
			break
		}
		temp = temp.next //若temp.next不为空 就将temp的值更新为当前temp指向的下一个people
	}
	temp.next = newPeople //将要添加的newPeople赋值给next为空的people.next
}
func showPeople(head *People) {
	temp := head
	fmt.Print("head-->")
	for {
		if temp.next == nil {
			break
		} //如果temp.next表示到链表末尾 就退出循环
		fmt.Printf("[%v %v]-->", temp.next.id, temp.next.name)
		temp = temp.next //若temp.next不为空 就将temp的值更新为当前temp指向的下一个people
	}
	fmt.Println()
}
func main() {
	head := &People{}
	people1 := &People{id: 1, name: "一"}
	people3 := &People{id: 3, name: "三"}
	people2 := &People{id: 2, name: "二"}
	addPeople(head, people1)
	addPeople(head, people3)
	addPeople(head, people2)
	showPeople(head)
}
