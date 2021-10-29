package main

import "fmt"

type People struct {
	id   int
	name string
	next *People //下个一元素指针
	prev *People //上一个元素指针
}

//向链表末尾添加一个元素
func addPeople(head *People, newPeople *People) {
	temp := head //将链表头给temp
	for {
		//从temp开始 遍历链表 当其中一个temp的next为nil时退出循环
		if temp.next == nil {
			break
		}
		temp = temp.next //若temp.next不为空 就将temp的值更新为当前temp指向的下一个people
	}
	temp.next = newPeople //将要添加的newPeople赋值给next为空的people.next
	newPeople.prev = temp //将temp赋值给newPeople.prev  此时就形成了双向列表：temp.next指向newPeople newPeople的prev指向temp
}

//正向遍历双链表
func showPeopleAlong(head *People) {
	temp := head
	fmt.Print("正向双链表：head-->")
	for {
		if temp.next == nil {
			break
		} //如果temp.next表示到链表末尾 就退出循环
		fmt.Printf("[%v %v]-->", temp.next.id, temp.next.name)
		temp = temp.next //若temp.next不为空 就将temp的值更新为当前temp指向的下一个people
	}
	fmt.Println()
}

//逆向遍历双链表
func showPeopleInverse(head *People) {
	temp := head
	for {
		if temp.next == nil {break}
		temp = temp.next
	}
	fmt.Print("逆向双链表：")
	for {
		if temp.prev == nil {break} //如果temp.prev为空表示到链表表头 就退出循环
		fmt.Printf("<--[%v %v]", temp.id, temp.name)
		temp = temp.prev //若temp.prev不为空 就将temp的值更新为当前temp指向的上一个people
	}
	fmt.Print("<--head")
	fmt.Println()
}

//删除
func deletePeople(head *People, id int) {
	temp := head //将链表头给temp
	for {
		if temp.next == nil {
			fmt.Printf("链表中没有id为[%v]的元素 删除失败！\n", id)
			return //如果链表到末尾表示链表中没有id为传入值的cat 退出函数
		} else if temp.next.id == id { //找到id和传入值id相等的people  此时temp.next所指向的就是目标people
			break
		}
		temp = temp.next //若temp.next不为空 就将temp的值更新为当前temp指向的下一个people
	}
	if temp.next.next != nil { //如果temp.next.next不为空就说明删除的不是双链表的最后一个元素
		temp.next = temp.next.next //将temp.next里的next指针内容赋值给temp.next
		temp.next.prev = temp      //将temp.next的prev内容赋值为temp  此时就将temp.next从双链表中删除
	} else {
		//如果temp.next.next为空就说明删除的是双链表的最后一个元素
		temp.next = nil //直接给temp.next赋值nil即可
	}
}
func main() {
	head := &People{}
	people1 := &People{id: 1, name: "一"}
	people2 := &People{id: 2, name: "二"}
	people3 := &People{id: 3, name: "三"}
	addPeople(head, people1)
	addPeople(head, people2)
	addPeople(head, people3)
	showPeopleAlong(head)
	showPeopleInverse(head)
	deletePeople(head, 3)
	showPeopleAlong(head)
}
