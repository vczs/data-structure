package main

import "fmt"

type Cat struct {
	id   int
	name string
	next *Cat
}

func addCat(head *Cat, newCat *Cat) {
	//如果头链接的next为空 就说明链表中暂无元素
	if head.next == nil {
		head.id = newCat.id
		head.name = newCat.name
		head.next = head //先将链表头head的next指向自身head形成闭环
		return
	}
	temp := head
	for {
		if temp.id == newCat.id {
			fmt.Printf("id为[%v]的Cat已存在名为[%v]  当前%v添加失败！\n", temp.id, temp.name, *newCat)
			return
		}
		if temp.next == head {
			break
		} //当遍历到temp.next等于表头head时说明环形链表遍历到末尾
		temp = temp.next
	}
	temp.next = newCat //将末尾元素的temp.next指向newCat
	newCat.next = head //将newCat的next指向表头形成闭环
}

//遍历链表
func showCat(head *Cat) {
	temp := head
	if temp.next == nil {
		fmt.Println("链表为空！！！")
		return
	}
	fmt.Print("环形列表：-->")
	for {
		fmt.Printf("[%v %v]-->", temp.id, temp.name)
		if temp.next == head {
			break
		} //当遍历到temp.next等于表头head时说明环形链表遍历到末尾
		temp = temp.next
	}
	fmt.Println()
}

//删除环形链表的某一元素  因为有可能会修改链表头head  所以返回一个*Cat
func deleteCat(head *Cat, id int) *Cat {
	temp := head //让temp指向head
	tail := head //先让tail指向head
	if temp.next == nil {
		fmt.Println("链表为空！！！")
		return head
	}
	//如果temp.next为head说明链表中只有一个元素 就是链表头
	if temp.next == head {
		if temp.id == id {
			temp.next = nil
			fmt.Println("只有一个链表头 已删除！")
		}
		return head
	}
	//将tail指向链表的最后一位
	for {
		if tail.next == head {
			break
		}
		tail = tail.next
	}
	flag := true
	for {
		if temp.next == head {
			break
		} //已经遍历到最后一个元素（最后一个还未比较）
		if temp.id == id {
			if temp == head {head = head.next} //如果删除的是链表头 就要将链表头交给它的下一个元素
			tail.next = temp.next //让tail.next指向temp的next字段指向的cat 即就从链表中删除了temp
			fmt.Printf("[%v]被删除！！！\n", temp.id)
			flag = false
			break //删除完成就退出循环
		}
		temp = temp.next //temp指向下一个元素
		tail = tail.next //tail指向下一个元素
	}
	//如果上面的代码没有找到目标元素 就对比看最后一个是不是目标元素  因为前面的代码还剩链表中的最后一个元素未比较
	if flag {
		if temp.id == id { //如果最后一个元素id是目标元素就删除他
			tail.next = temp.next //让tail.next指向temp的next字段指向的cat 即就从链表中删除了temp
			fmt.Printf("[%v]被删除！！！\n", temp.id)
		} else {
			//如果最后一个元素id不是目标元素就提示元素不存在
			if flag {fmt.Printf("[%v]元素不存在！！！\n", id)}
		}
	}
	return head
}
func main() {
	head := &Cat{}
	cat1 := &Cat{id: 1, name: "一"}
	cat2 := &Cat{id: 2, name: "二"}
	cat3 := &Cat{id: 3, name: "三"}
	addCat(head, cat1)
	addCat(head, cat2)
	addCat(head, cat3)
	showCat(head)
	head = deleteCat(head, 1)
	showCat(head)
}
