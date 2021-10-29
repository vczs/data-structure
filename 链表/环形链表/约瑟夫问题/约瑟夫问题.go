package main

import "fmt"

type Boy struct {
	id   int
	next *Boy
}

//给链表添加指定个数的元素
func addBoy(num int) *Boy {
	first := &Boy{}
	curBoy := &Boy{}
	if num < 1 {
		fmt.Println("num不能小于1")
		return first
	}
	for i := 1; i <= num; i++ {
		boy := &Boy{id: i} //实例化一个boy
		if i == 1 {        //链表的头
			first = boy         //先让first指向boy 因为是第一个元素所以他是表头
			curBoy = boy        //curBoy也指向boy  curBoy始终要指向链表的最后一个元素  因为此时的boy是第一个元素也是最后一个元素所以指向boy
			curBoy.next = first //末尾元素curBoy的next指向链表头frist 构成环形链表
		} else {
			curBoy.next = boy   //要新增加一个元素boy 让此时链表中的最后一个元素curBoy的next指向要新加的boy 将他链接到链表中
			curBoy = boy        //链接好boy后 将curBoy指向它    因为将他链接后 他就成了链表中的最后一个元素  因为curBoy始终要指向链表的最后一个元素
			curBoy.next = first //再让末尾元素curBoy的next指向链表头frist 构成环形链表
		}
	}
	return first //返回表头
}

//遍历链表
func showBoy(first *Boy) {
	temp := first
	if temp.next == nil {
		fmt.Println("链表为空！！！")
		return
	}
	fmt.Print("环形列表：-->")
	for {
		fmt.Printf("[%v]-->", temp.id)
		if temp.next == first {break} //当遍历到temp.next等于表头head时说明环形链表遍历到末尾
		temp = temp.next
	}
	fmt.Println()
}

/**
开始玩约瑟夫游戏
first：链表表头
start：从链表表头开始数到第几个元素开始
count：每次间隔几个元素
*/
func playGame(first *Boy, start int, count int) {
	//如果curBoy为空说明是个空链表
	if first.next == nil {
		fmt.Println("链表为空！！！")
		return
	}
	curBoy := first
	tail := first
	//让tail先指向链表最后一个元素
	for {
		if tail.next == first {break} //此时curBoy指向链表最后一个元素
		tail = tail.next
	}
	//让curBoy指向start目标
	for i := 1; i < start; i++ {
		//让curBoy从表头移动start次 找到开始游戏的元素
		curBoy = curBoy.next
		tail = tail.next //curBoy移动n次 tail也跟着移动n次 始终指向curBoy指向的前一个元素
	}
	//开始计数并删除目标
	for {
		//计数count次定位到要删除的元素
		for i := 1; i < count; i++ {
			curBoy = curBoy.next
			tail = tail.next //curBoy移动n次 tail也跟着移动n次 始终指向curBoy指向的前一个元素
		}
		//定位到目标 快开始删除curBoy指向的元素
		fmt.Printf("%v已出列-->\n", curBoy.id)
		curBoy = curBoy.next //让curBoy指向他的下一个元素
		tail.next = curBoy   //让curBoy后面的元素tail的next指向 更新后的curBoy  此时就将curBoy之前指向的元素从链表中删除
		if curBoy == tail {
			break
		} //大明湖curBoy和tail指向的是同一个元素 就说明链表中只剩下一个元素 就退出循环
	}
	//操作剩下的最后一个元素 此时curBoy和tail都指向他 用他俩谁操作该元素都可以
	fmt.Printf("最后出列：%v已出列-->\n", curBoy.id)
}
func main() {
	first := addBoy(100)    //向链表中添加100个boy 返回链表表头
	showBoy(first)          //遍历链表
	playGame(first, 20, 15) //开始约瑟夫游戏  从链表中第20个人开始  每隔15次出列一个boy
}
